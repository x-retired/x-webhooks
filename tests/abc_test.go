package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func TestABC(t *testing.T) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)
	// Use client
	var query struct {
		Repository struct {
			Refs struct {
				Nodes struct {
					id     githubv4.String
					name   githubv4.String
					prefix githubv4.String
				}
			} `graphql:"refs(first: 10, orderBy: {direction:DESC, field:TAG_COMMIT_DATE}, refPrefix:\"refs/tags/\")"`
		} `graphql:"repository(owner: \"tencent\", name: \"bk-cmdb\")"`
	}
	/*{
		repository(owner: "tencent", name: "bk-cmdb") {
		refs(first: 10, orderBy: {direction:DESC, field:TAG_COMMIT_DATE}, refPrefix:"refs/tags/") {
			nodes {
				id
				name
				prefix
				target{
					id
				}
			}
			pageInfo {
				startCursor
				endCursor
				hasPreviousPage
				hasNextPage
			}
		}
	}
	}*/
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
		fmt.Println("err:", err)
	}
	//for k, v := range query.Repository.Refs.Node {
	//	fmt.Println("    k:", k)
	//	fmt.Println("    v:", v)
	//}
	fmt.Println(query.Repository.Refs.Nodes)
}
