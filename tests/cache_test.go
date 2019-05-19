package test

import (
	"testing"
	"time"

	"github.com/astaxie/beego/cache"
)

func TestCache(t *testing.T) {
	bm, _ := cache.NewCache("memory", `{"interval":60}`)
	_ = bm.Put("user", "foo", 10*time.Second) //
	t.Log(bm.Get("user"))
	t.Log(bm.IsExist("user"))
	_ = bm.Delete("user")
}
