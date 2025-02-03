package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/elkcityhazard/pc-building-company/content"
)

func (hr *HandlerRepo) HandleGetRobots(w http.ResponseWriter, r *http.Request) {

	contentFS := content.GetContentFS()

	file, err := contentFS.Open("content/robots.txt")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, string(content))

}
