package httpServer

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Ip       string `json:"ip"`
	HttpCode int    `json:"http_code"`
}

func ServerStart() {
	routeRegister()
	http.ListenAndServe(":8080", nil)
}

func routeRegister() {
	http.HandleFunc("/GetSettings", settings)
	http.HandleFunc("/healthz", healthz)
}

func settings(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	//fmt.Println(header)

	for k, v := range header {
		w.Header().Set(k, strings.Join(v, ""))
	}

	version := os.Getenv("VERSION")
	if version == "" {
		version = "0.0.0"
	}
	w.Header().Set("VERSION", version)

	ip := ClientIP(r)
	resp := Response{
		Ip:       ip,
		HttpCode: http.StatusOK,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(res)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
	w.WriteHeader(http.StatusOK)
}
