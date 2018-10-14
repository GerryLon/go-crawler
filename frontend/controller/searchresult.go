package controller

import (
	"context"
	"github.com/GerryLon/go-crawler/config"
	"github.com/GerryLon/go-crawler/engine"
	"github.com/GerryLon/go-crawler/frontend/model"
	"github.com/GerryLon/go-crawler/frontend/view"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
)

type SearchResultHandler struct {
	view view.SearchResultView
}

func CreateSearchResultHandler(filename string) SearchResultHandler {
	return SearchResultHandler{
		view: view.CreateSearchResultView(filename),
	}
}

func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := model.SearchResult{}

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.ElasticUrl))

	if err != nil {
		panic(err)
	}

	resp, err := client.Search(config.ElasticIndex).
		Query(elastic.NewQueryStringQuery("女")).
		Do(context.Background())

	result.Hits = resp.TotalHits()
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	s.view.Render(w, result)
}
