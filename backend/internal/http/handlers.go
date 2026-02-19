package http

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {}
func GenerateNotes(w http.ResponseWriter, r *http.Request) {}
func ExportNotes(w http.ResponseWriter, r *http.Request) {}