package webapp

import (
	"net/http"

	"github.com/gobuffalo/packr"
)

type Server struct {
	box packr.Box
}

func NewServer() *Server {
	box := packr.NewBox("./public")
	return &Server{box}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(s.box)
	fs.ServeHTTP(w, r)
}
