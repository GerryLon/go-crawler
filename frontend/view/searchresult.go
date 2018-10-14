package view

import (
	"github.com/GerryLon/go-crawler/frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

// create searchResultView using supplied filename(template file)
func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template.Must(template.ParseFiles(filename)),
	}
}

// render search result to template
func (s *SearchResultView) Render(w io.Writer, result model.SearchResult) error {
	return s.template.Execute(w, result)
}
