package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var pwd string
var port string = ":8035"
var dir string = "www"

func parseArgs() {
	if len(os.Args) > 2 {
		c := os.Args[1]
		d := os.Args[2]
		a, _ := strconv.Atoi(c)
		b, _ := strconv.Atoi(d)
		if a > 0 {
			port = ":" + c
			if a == 0 {
				dir = c
			} else if b == 0 {
				dir = d
			}
		} else if b > 0 {
			port = ":" + d
			if a == 0 {
				dir = c
			} else if b == 0 {
				dir = d
			}
		}
	} else if len(os.Args) == 2 {
		c := os.Args[1]
		a, _ := strconv.Atoi(c)
		if a > 0 {
			port = ":" + c
		} else if a == 0 {
			dir = c
		}
	} else {
		return
	}
}

func serve(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path
	http.ServeFile(w, r, dir+"/"+file)
	log.Print(r.Method, " ", file)
}

func main() {
	wd, _ := os.Getwd()
	parseArgs()
	pwd = filepath.Join(wd, dir)

	log.Printf("SERVE %s on %s ", pwd, port)

	http.HandleFunc("/", serve)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
