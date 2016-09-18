package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"io"
	"github.com/NYTimes/gziphandler"
)

func main() {
	m := mux.NewRouter()
	m.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {
		file, _ := os.Open("/tmp/dataBasic.jpg")
		io.Copy(w, file)
		file.Close()
	})
	http.ListenAndServe(":3000", gziphandler.GzipHandler(m))
}
