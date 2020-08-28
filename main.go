package main

import (
	"net/http"
	"os"
	"time"

	"github.com/ZYunH/pgoproxy/goproxy"
	"github.com/ZYunH/pgoproxy/goproxy/cacher"
)

func main() {
	instance := goproxy.New()
	instance.GoBinFetchTimeout = time.Minute * 3
	instance.GoBinEnv = append(os.Environ(), "GOPROXY=https://goproxy.cn|https://mirrors.aliyun.com/goproxy|https://goproxy.io", "GOSUMDB=sum.golang.google.cn")
	instance.ProxiedSUMDBs = []string{"sum.golang.org https://sum.golang.google.cn/", "sum.golang.google.cn https://sum.golang.google.cn/"}

	// Add disk cache.
	diskCacher := &cacher.Disk{Root: "modcache"}
	instance.Cacher = diskCacher

	http.ListenAndServe("0.0.0.0:7617", instance)
}
