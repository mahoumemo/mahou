package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"

	"github.com/mahoumemo/mahou/auth"
)

var isMahouServiceURL = regexp.MustCompile(`^\/ds\/v2-[a-z]{2}\/(.*)$`).MatchString

type Server struct{}

func main() {
	proto := flag.String("proto", "tcp", "protocol to use (\"tcp\", \"unix\", etc)")
	addr := flag.String("addr", "0.0.0.0:8101", "address to listen on")
	flag.Parse()

	if *proto == "unix" {
		os.Remove(*addr)
	}

	listener, err := net.Listen(*proto, *addr)
	if err != nil {
		log.Fatalf("failed to create %s listener on %s: %s", *proto, *addr, err)
	}

	if *proto == "unix" {
		os.Chmod(*addr, 0777)
	}

	http.Serve(listener, &Server{})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !isMahouServiceURL(r.URL.Path) {
		http.Error(w, "unsupported URL", http.StatusBadRequest)
		return
	}

	switch r.URL.Path[len("/ds/v2-xx/"):] {
	case "auth":
		auth.HandleRequest(w, r)
	default:
		http.Error(w, "unknown endpoint", http.StatusNotFound)
		return
	}
}
