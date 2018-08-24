package api

import (
	"fmt"
	"log"
	"net/http"
)

type LightSwitcher interface {
	TurnOn() error
	TurnOff() error
}

type Server struct {
	switcher LightSwitcher
	mux      *http.ServeMux
}

// TODO: Handle auth
func NewServer(switcher LightSwitcher) *Server {
	s := &Server{
		switcher: switcher,
	}

	// Setup routes
	mux := http.NewServeMux()

	mux.HandleFunc("/turn_on_lights", s.TurnOnLights)
	mux.HandleFunc("/turn_off_lights", s.TurnOffLights)

	s.mux = mux

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) TurnOnLights(w http.ResponseWriter, r *http.Request) {
	if err := s.switcher.TurnOn(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func (s *Server) TurnOffLights(w http.ResponseWriter, r *http.Request) {
	if err := s.switcher.TurnOff(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
