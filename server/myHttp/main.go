package myHttp

import (
	"fmt"
	"net/http"
)

func RunHttp() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":9999", nil)
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
