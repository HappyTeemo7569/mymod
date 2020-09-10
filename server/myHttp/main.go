package myHttp

import (
	"fmt"
	"github.com/HappyTeemo7569/mymod/base"
	"net/http"
)

func RunHttp() {
	r := New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	addr := fmt.Sprintf(":%d", base.ConfigServer.HttpPort)
	err := r.Run(addr)
	if err != nil {
		panic("服务启动失败:HTTP监听失败:" + err.Error())
	}

}
