package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	serverBindIp := flag.String("bind", "0.0.0.0", "Server ip to bind")
	serverPort := flag.Int("port", 5000, "Server port to listen")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		hostname, err := os.Hostname()
		now := time.Now().UTC()

		if err != nil {
			panic(err)
		}

		log.Println("Info Request from", req.RemoteAddr)

		out := now.Format("2006-01-02 15:04:05")
		out += " Hostname: " + hostname + " requested from " + req.RemoteAddr

		_, _ = io.WriteString(w, out)
	})

	log.Println("Docker Node Hostname Server running on", *serverBindIp+":"+strconv.Itoa(*serverPort))
	log.Fatal(http.ListenAndServe(*serverBindIp+":"+strconv.Itoa(*serverPort), nil))
}
