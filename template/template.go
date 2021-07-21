package template

import (
	"bytes"
	"embed"
	"html/template"
	"log"
	"time"
)

// please be careful about the file names
//go:embed *.tmpl
var tpls embed.FS

// templates are the embeded templates
var templates *template.Template

var funcMap = template.FuncMap{
	// The name "inc" is what the function will be called in the template text.
	"inc": func(i int) int {
		return i + 1
	},
}

func init() {
	t, err := template.New("test").Funcs(funcMap).ParseFS(tpls, "*")
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}
	templates = t
}

type Setting struct {
	Name          string
	ID            int
	RangeKey      string
	PartitionName string
	Reservations  []*Reservation
}

func (s *Setting) Render() string {
	template := "setting.tmpl"
	buf := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(buf, template, s); err != nil {
		panic("failed to render template: " + err.Error())
	}
	return buf.String()
}

type Reservation struct {
	ID         int
	Department string
	Date       time.Time
}

func (r *Reservation) Render() string {
	template := "reservation.tmpl"
	buf := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(buf, template, r); err != nil {
		panic("failed to render template: " + err.Error())
	}
	return buf.String()
}

type Availability struct {
	ID                int
	Department        string
	Date              time.Time
	NumOfAvailability int
}

type Availabilities struct {
	Availabilities []*Availability
	LoggedIn       bool
}

func (a *Availabilities) Render() string {
	template := "availability.tmpl"
	buf := &bytes.Buffer{}
	if err := templates.ExecuteTemplate(buf, template, a); err != nil {
		panic("failed to render template: " + err.Error())
	}
	return buf.String()
}
