package home

import (
	"html/template"

	"github.com/adampresley/rendering"
)

type HomePage struct {
	rendering.BaseViewModel
	AssetVersion string
	NumYears     int
}

type Experience struct {
	rendering.BaseViewModel
	AssetVersion string
	Jobs         []Job
}

type Job struct {
	Logo        string
	LogoAlt     string
	Initials    string
	Title       string
	Description template.HTML
	YearStarted string
	Highlights  []template.HTML
}
