package html

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed static/**
var staticFS embed.FS

//go:embed templates/**
var templatesFS embed.FS

var Templates = template.Must(template.ParseFS(templatesFS, "templates/*.html"))
var StaticHandler = http.FileServerFS(staticFS)
