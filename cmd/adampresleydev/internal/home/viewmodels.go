package home

import (
	"html/template"

	"github.com/adampresley/adamgokit/rendering"
)

type HomePage struct {
	rendering.BaseViewModel
	NumYears int
}

type Experience struct {
	rendering.BaseViewModel
	Jobs []Job
}

type Job struct {
	Logo        string
	LogoAlt     string
	Initials    string
	Title       string
	Description template.HTML
	YearStarted string
}
