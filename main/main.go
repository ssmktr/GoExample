package main

import (
	"net/http"
	"github.com/unrolled/render"
)

var renderer render.Render

func main() {

	http.HandleFunc("/", func (res http.ResponseWriter, req *http.Request) {
		renderer.Text(res, http.StatusOK, "SSMKTR")
	})

	http.ListenAndServe(":2305", nil)
}


