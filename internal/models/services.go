package models

import "html/template"

type Service struct {
	Name        string        `json:"name"`
	Url         string        `json:"url"`
	Description string        `json:"description"`
	Icon        template.HTML `json:"icon"`
}

type HomeServices struct {
	Services []Service `json:"services"`
}
