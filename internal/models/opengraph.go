package models

type OpenGraph struct {
	AppID         string
	URL           string
	Type          string
	Title         string
	Image         string
	ImageAlt      string
	Description   string
	SiteName      string
	Locale        string
	ArticleAuthor string
}

func NewOpenGraph() *OpenGraph {
	return &OpenGraph{
		AppID:         "",
		URL:           "",
		Type:          "website",
		Title:         "",
		Image:         "",
		ImageAlt:      "",
		Description:   "",
		SiteName:      "",
		Locale:        "en_US",
		ArticleAuthor: "",
	}
}

type TwitterCard struct {
	Card        string
	Site        string
	Creator     string
	URL         string
	Title       string
	Description string
	Image       string
	ImageAlt    string
}

func NewTwitterCard() *TwitterCard {
	return &TwitterCard{
		Card:        "",
		Site:        "",
		Creator:     "",
		URL:         "",
		Title:       "",
		Description: "",
		Image:       "",
		ImageAlt:    "",
	}
}
