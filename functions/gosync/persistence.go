package gosync

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/xiexianbin/webhooks/drivers/aliyun"
	"github.com/xiexianbin/webhooks/utils"

	"github.com/astaxie/beego/logs"
)

func ReadDir(filesMap map[string]interface{}, sourceDir, subDir string) {
	if strings.HasSuffix(sourceDir, "/") == false {
		sourceDir += "/"
	}

	currentDir := sourceDir
	if subDir != "" {
		if strings.HasSuffix(subDir, "/") == false {
			subDir += "/"
		}
		currentDir += subDir
	}

	dirInfos, err := ioutil.ReadDir(currentDir)
	if err != nil {
		logs.Info("Read Source Dir error:", err)
	}

	for _, dirInfo := range dirInfos {
		if dirInfo.IsDir() {
			newSubDir := subDir + dirInfo.Name()
			ReadDir(filesMap, sourceDir, newSubDir)
		} else {
			filePath := currentDir + dirInfo.Name()
			fileByte, err := ioutil.ReadFile(filePath)
			if err != nil {
				logs.Info("read file err:", err)
			}

			fileContent := string(fileByte)
			re, err := regexp.Compile("<li>Build <small>&copy; .*</small></li>")
			if err != nil {
				logs.Info("init regexp err:", err)
			}

			fileContent = re.ReplaceAllString(fileContent, "")
			md5sum := utils.Md5sum(string(fileContent))

			filesMap[strings.Replace(filePath, sourceDir, "", 1)] = md5sum
		}
	}
}

func CacheWrite(m map[string]interface{}, cacheFile string) error {
	_, err := os.Stat(cacheFile)
	if err != nil {
		f, _ := os.Create(cacheFile)
		defer f.Close()
	}

	j, err := json.Marshal(m)
	if err != nil {
		logs.Info("json.Marshal failed:", err)
		return err
	}

	err = ioutil.WriteFile(cacheFile, []byte(j), 0644)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func CacheRead(filename string) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadFile(filename)
	var m map[string]interface{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		logs.Info("Unmarshal failed, ", err)
		return nil, err
	}
	return m, nil
}

func SyncLocalToOSS(aliyunOSSConfig *utils.AliyunOSSConfig, sourceDir, metaKey string) error {
	if metaKey == "" {
		metaKey = "Content-Md5sum"
	}
	cacheFile := "./" + aliyunOSSConfig.BucketName + ".js"

	// read local files
	filesMap := make(map[string]interface{})
	ReadDir(filesMap, sourceDir, "")

	// list oss object metadata
	objectsMap := make(map[string]interface{})
	_, err := os.Stat(cacheFile)
	if err != nil {
		objectsMap, err = aliyun.ListObjects(aliyunOSSConfig, metaKey)
		if err != nil {
			aliyun.HandleError(err)
		}
	} else {
		objectsMap, err = CacheRead(cacheFile)
		if err != nil {
			aliyun.HandleError(err)
		}
	}

	// get diff map
	justM1, justM2, diffM1AndM2, err := utils.DiffMap(filesMap, objectsMap)
	if err != nil {
		aliyun.HandleError(err)
	}

	// do upload
	logs.Info("new file Map:")
	for k, v := range justM1 {
		logs.Info(k, v)
		metasMap := make(map[string]interface{})
		metasMap[metaKey] = v
		err := aliyun.PutObjectFromFile(aliyunOSSConfig, k, sourceDir + "/" + k, metasMap)
		if err != nil {
			aliyun.HandleError(err)
			logs.Info("Upload OSS Object", k, "Error:", err)
		}
		logs.Info("Upload OSS Object", k)
	}

	logs.Info("update file Map:")
	for k, v := range diffM1AndM2 {
		logs.Info(k, v)
		metasMap := make(map[string]interface{})
		metasMap[metaKey] = v
		err := aliyun.PutObjectFromFile(aliyunOSSConfig, k, sourceDir + "/" + k, metasMap)
		if err != nil {
			aliyun.HandleError(err)
			logs.Info("Update OSS Object", k, "Error:", err)
		}
		logs.Info("Update OSS Object", k)
	}
	logs.Info("delete file Map:")
	for k, v := range justM2 {
		logs.Info(k, v)
		err := aliyun.DeleteObject(aliyunOSSConfig, k)
		if err != nil {
			logs.Info("Delete OSS Object", k, "Error:", err)
		}
		logs.Info("Delete OSS Object", k)
	}

	// cache new map to file
	_, err = os.Stat(cacheFile)
	if err == nil {
		_ = os.Truncate(cacheFile, 0)
	}

	err = CacheWrite(filesMap, cacheFile)
	if err != nil {
		logs.Info("cache file map to file fail.")
	}
	return nil
}
