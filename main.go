package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

// Hasher is the interface implemented by objects having a Hash method.
type Hasher interface {
	// Hash generates the hash value of v.
	Hash(v string) string
}

// KVS is key value store.
type KVS interface {

	// store stores a key value pair.
	store(k, v string) error

	// load returns the value associated with a given key.
	load(k string) (string, error)
}

type server struct {
	kvs    KVS
	hasher Hasher
}

func newServer(kvs KVS, hasher Hasher) *server {
	return &server{
		kvs:    kvs,
		hasher: hasher,
	}
}

func (s *server) shorten(w http.ResponseWriter, r *http.Request) {
	// extract original URL.
	org := r.URL.Query().Get("url")
	if org == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "empty url field")
		return
	}

	// hash original url.
	short := s.hasher.Hash(org)
	log.Printf("/shorten org=%s short=%s", org, short)

	// store short url/original url pair in the key-value store.
	err := s.kvs.store(short, org)
	if err != nil {
		// handle error
		log.Printf("/shorten error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "interval server error: %v", err)
		return
	}

	// print hostname and shortened url in the response.
	fullshort := path.Join(r.Host, short)
	fmt.Fprintf(w, "shortened URL: %s\n", fullshort)
}

func (s *server) redirect(w http.ResponseWriter, r *http.Request) {
	// consider we received a short url, extract it.
	short := strings.Replace(r.URL.RequestURI(), "/", "", 1)

	// load original url from the key-value store.
	org, err := s.kvs.load(short)
	if err != nil {
		log.Printf("/ error: loading short=%s", short)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "URL not found: %v", err)
		return
	}

	// forge the redirection URL.
	log.Printf("/ redirecting short=%s org=%s", short, org)
	http.Redirect(w, r, org, 307)
}

func main() {
	addr := flag.String("addr", "localhost:7070", "listening address")
	flag.Parse()

	// create url key-value store.
	kvs := newBasicKVS()

	// create url hasher.
	hasher := newBasicHasher()

	// create the server.
	s := &server{
		kvs:    kvs,
		hasher: hasher,
	}

	// install HTTP handlers.
	http.HandleFunc("/shorten/", s.shorten)
	http.HandleFunc("/", s.redirect)

	log.Println("listening on", *addr)
	log.Fatalln(http.ListenAndServe(*addr, nil))
}
