package http

import "net/http"

type Router struct{ Mux *http.ServeMux }

func NewRouter() *Router { r:=http.NewServeMux(); return &Router{Mux:r} }