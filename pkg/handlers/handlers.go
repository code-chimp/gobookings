package handlers

import (
	"github.com/code-chimp/gobookings/pkg/config"
	"github.com/code-chimp/gobookings/pkg/models"
	"github.com/code-chimp/gobookings/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home display the landing page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About display About page - test template data
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	sm := make(map[string]string)
	sm["test"] = "Hodor, hodor hodor."
	sm["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: sm,
	})
}

// Contact display company contact information
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.gohtml", &models.TemplateData{})
}

// GeneralsQuarters display details of the General's Quarters
func (m *Repository) GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.gohtml", &models.TemplateData{})
}

// ColonelsSuite display details of the Colonel's Suite
func (m *Repository) ColonelsSuite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "colonels.page.gohtml", &models.TemplateData{})
}

// MakeReservation allow user to confirm a booking
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.gohtml", &models.TemplateData{})
}

// SearchAvailability allow the user to determine room availability
func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.gohtml", &models.TemplateData{})
}
