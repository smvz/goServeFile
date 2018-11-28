package main

import (
	"os"
	"path/filepath"
	"fmt"
	"strconv"
	"net/http"
	"net/url"
)

func exitWithError(err string) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(1)
}

func usage() {
	fmt.Println(os.Args[0], "file [port]")
	os.Exit(1)
}

func urlPathStringForFilePath(path string, encoded bool) string {
	u, _ := url.Parse("/")
	_, file := filepath.Split(path)
	u.Path += file
	if encoded {
		return u.String()
	} else {
		return u.Path
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, urlPathStringForFilePath(os.Args[1], true), 302)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, os.Args[1])
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		usage()
	} else if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		exitWithError("File does not exist at path")
	} else {
		file, err := os.Open(os.Args[1])
		file.Close()
		if err != nil {
			exitWithError("Cannot read file at path")
		}
	}
	
	port := ":8080"
	if len(os.Args) == 3 {
		if _, err := strconv.Atoi(os.Args[2]); err == nil {
			port = fmt.Sprintf(":%s", os.Args[2])
		} else {
			exitWithError("Invalid port number")
		}
	}

	http.HandleFunc("/", redirect)
	http.HandleFunc(urlPathStringForFilePath(os.Args[1], false), serveFile)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		exitWithError(err.Error())
	}
}
