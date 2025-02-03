package templates

import "embed"

//go:embed templates

var templatesFS embed.FS

func GetTemplatesFS() embed.FS {
	return templatesFS
}
