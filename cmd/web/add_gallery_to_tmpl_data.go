package main

import (
	"fmt"
	"path/filepath"

	"github.com/adrg/frontmatter"
	"github.com/elkcityhazard/pc-building-company/content"
)

type GalleryFM struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Images      []GalleryImage
}

type GalleryImage struct {
	Id          int    `yaml:"id"`
	Thumbnail   string `yaml:"thumb"`
	ThumbWidth  int    `yaml:"thumb_width"`
	ThumbHeight int    `yaml:"thumb_height"`
	Fullsize    string `yaml:"fulls"`
	FullsWidth  int    `yaml:"fulls_width"`
	FullsHeight int    `yaml:"fulls_height"`
	Description string `yaml:"description"`
}

func addGalleryToTmplData(pathToTmpl string) (*GalleryFM, error) {
	contentFS := content.GetContentFS()

	g, err := contentFS.Open(fmt.Sprintf("%s", filepath.Clean(pathToTmpl)))
	if err != nil {
		return nil, err
	}

	var matter GalleryFM

	_, err = frontmatter.Parse(g, &matter)
	if err != nil {
		return nil, err
	}

	return &matter, nil
}
