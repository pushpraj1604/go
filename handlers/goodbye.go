package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServHTTP(rw http.ResponseWriter, r *http.request) {
	rw.Write([]byte("byee"))
}
