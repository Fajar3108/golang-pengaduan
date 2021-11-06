package controllers

import (
	"pengaduan/helpers"
)

// Used to compact data to view
type compact map[string]interface{}

// Aliasing helpers
var view = helpers.View

// Must used components
var components = []string{"views/components/header.html"}
