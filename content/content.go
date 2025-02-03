package content

import "embed"

//go:embed content
var contentFS embed.FS

func GetContentFS() embed.FS {
	return contentFS
}
