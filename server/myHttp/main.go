package myHttp

import (
	"fmt"
	"github.com/HappyTeemo7569/mymod/base"
	"net/http"
)

func RunHttp() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	addr := fmt.Sprintf(":%d", base.ConfigServer.HttpPort)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic("服务启动失败:HTTP监听失败:" + err.Error())
	}
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
