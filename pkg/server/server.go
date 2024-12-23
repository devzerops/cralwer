package server

import (
	
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Collect(urls []string) error {
	// URL 수집 로직
	return nil
}

func (s *Server) Resolve(ips []string) ([]string, error) {
	// IP 해석 로직
	return nil, nil
}
