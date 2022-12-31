package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var (
	err        error
	hm         string
	siteFlHdlr http.Handler
)

const (
	PORT       = ":3000"
	PUBLIC_DIR = "public"
)

func serve() error {
	hm, err = os.UserHomeDir()
	if err != nil {
		return err
	}

	siteFlHdlr = http.FileServer(http.Dir(filepath.Join(hm, PUBLIC_DIR)))
	http.Handle("/", siteFlHdlr)

	fmt.Printf("Listening on %s ...\n", PORT)
	return http.ListenAndServe(PORT, nil)
}

func main() {
	err = serve()
	if err != nil {
		panic(err)
	}
}
