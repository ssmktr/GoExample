package httpservermanager

import (
	"net/http"
	"github.com/unrolled/render"
	"fmt"
	"encoding/json"
)

type signupPacket struct {
	Uid string
	Id string
	Pw string
	Nickname string
}

var renderer render.Render

func OnHttpServer() {

	http.HandleFunc("/signup", func(res http.ResponseWriter, req *http.Request) {
		data := make([]byte, 2048)
		n, _ := req.Body.Read(data)
		var res_pack signupPacket
		json.Unmarshal([]byte(string(data[:n])), &res_pack)

		fmt.Println(string(data[:n]))
		fmt.Println(res_pack.Id)

		bytes, _ := json.Marshal(res_pack)

		renderer.Data(res, http.StatusOK, bytes)
	})

	http.ListenAndServe(":2305", nil)
}