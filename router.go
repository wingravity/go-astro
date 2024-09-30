package main

import (
	"net/http"
	"text/template"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) ListenAndServe(address string) error {
	return http.ListenAndServe(address, r.mux)
}

type ViewsHandler struct {
	Index *template.Template
}

func NewTemplatesHandler() *ViewsHandler {
	return &ViewsHandler{
		Index: template.Must(template.New("index").Delims("[[", "]]").ParseFiles("client/dist/index.html")),
	}
}

type IndexParams struct {
	Name string
}

func (h *ViewsHandler) IndexView(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		h.Index.ExecuteTemplate(w, "index.html", IndexParams{Name: "Go Astro"})
		return
	}
	http.ServeFile(w, r, "client/dist"+r.URL.Path)
}
