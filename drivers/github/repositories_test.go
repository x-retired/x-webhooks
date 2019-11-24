package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"testing"
)


func TestGithubSDK(t *testing.T) {
	client := github.NewClient(nil)

	//releases, resp, err := client.Repositories.ListReleases(
	//	context.Background(),
	//	"Tencent",
	//	"bk-cmdb",
	//	nil)
	//if err != nil {
	//	fmt.Println("err. resp is", resp)
	//	return
	//}
	//for _, release := range releases {
	//	fmt.Println(release.GetName())
	//	fmt.Println(release.GetCreatedAt())
	//	fmt.Println(release.GetID())
	//	fmt.Println(release.GetNodeID())
	//	fmt.Println(release.GetTagName())
	//	fmt.Println("")
	//}
	opt := &github.ListOptions{Page: 1, PerPage: 1}
	tags, resp, err := client.Repositories.ListTags(
		context.Background(),
		"Tencent",
		"bk-cmdb",
		opt)
	if err != nil {
		fmt.Println("err. resp is", resp)
		return
	}
	for _, tag := range tags {
		fmt.Println(tag.GetName())
		fmt.Println(tag.GetCommit())
	}
}
