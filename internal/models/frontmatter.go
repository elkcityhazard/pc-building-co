package models

type DefaultFrontMatter struct {
	Title           string `yaml:"title"`
	Description     string `yaml:"description"`
	URL             string `yaml:"url"`
	MetaImage       string `yaml:"metaImage"`
	MetaImageAlt    string `yaml:"metaImageAlt"`
	Author          string `yaml:"author"`
	PublishDate     string `yaml:"publishDate"`
	UpdateDate      string `yaml:"updateDate"`
	BgImage         string `yaml:"bgImage"`
	TwitterUsername string `yaml:"twitterUsername"`
}
