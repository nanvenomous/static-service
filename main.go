package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

var (
	err    error
	hm     string
	flHdlr http.Handler
	header http.Header
)

const (
	PORT    = ":3000"
	BIN_DIR = "bin"
)

func addHeaders(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(path.Base(r.RequestURI))
		header = w.Header()
		header.Set("Content-Description", "File Transfer")
		header.Set(
			"Content-Disposition",
			"attachment; filename="+strconv.Quote(path.Base(r.RequestURI)),
		)
		header.Set("Content-Type", "application/octet-stream")
		handler.ServeHTTP(w, r)
		// Content-Length: [filesize]
	}
}

func serve() error {
	hm, err = os.UserHomeDir()
	if err != nil {
		return err
	}
	flHdlr = http.FileServer(http.Dir(filepath.Join(hm, BIN_DIR)))
	http.Handle("/", addHeaders(flHdlr))
	fmt.Printf("Listening on %s ...\n", PORT)
	return http.ListenAndServe(PORT, nil)
}

func main() {
	err = serve()
	if err != nil {
		panic(err)
	}
}
