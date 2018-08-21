package main

import (
		"net/http"
	"github.com/unrolled/render"
)

var renderer render.Render

func main() {

	//fmt.Println("SSMKTR")

	http.HandleFunc("/", func (res http.ResponseWriter, req *http.Request) {
		renderer.Text(res, http.StatusOK, "HI TONY?")
	})

	http.ListenAndServe(":2305", nil)
}


