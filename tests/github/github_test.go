package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"testing"
)


func TestCache(t *testing.T) {
	client := github.NewClient(nil)
	//tags, resp, err := client.Repositories.ListTags(context.Background(), "Tencent", "bk-cmdb", nil)
	//if err != nil {
	//	fmt.Println("err. resp is", resp)
	//	return
	//}
	//for _, tag := range tags {
	//	fmt.Println(tag.GetName())
	//}

	releases, resp, err := client.Repositories.ListReleases(context.Background(), "Tencent", "bk-cmdb", nil)
	if err != nil {
		fmt.Println("err. resp is", resp)
		return
	}
	for _, release := range releases {
		fmt.Println(release.GetName())
		fmt.Println(release.GetCreatedAt())
	}
}
