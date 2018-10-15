package controller

import (
	"context"
	"github.com/GerryLon/go-crawler/config"
	"github.com/GerryLon/go-crawler/engine"
	"github.com/GerryLon/go-crawler/frontend/model"
	"github.com/GerryLon/go-crawler/frontend/view"
	"github.com/olivere/elastic"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type SearchResultHandler struct {
	view          view.SearchResultView
	elasticClient *elastic.Client
}

func CreateSearchResultHandler(filename string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.ElasticUrl))

	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:          view.CreateSearchResultView(filename),
		elasticClient: client,
	}
}

// search condition form
type SearchCondition struct {
	q string
}

func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sc := SearchCondition{}
	sc.q = strings.TrimSpace(r.FormValue("q"))

	//result := model.SearchResult{}
	//
	//client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(config.ElasticUrl))
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//resp, err := client.Search(config.ElasticIndex).
	//	Query(elastic.NewQueryStringQuery("女")).
	//	Do(context.Background())
	//
	//result.Hits = resp.TotalHits()
	//result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	//
	//fmt.Printf("%+v", result.Items)
	s.view.Render(w, s.getSearchResult(&sc))
}

func (s SearchResultHandler) getSearchResult(sc *SearchCondition) model.SearchResult {
	result := model.SearchResult{}

	resp, err := s.elasticClient.Search(config.ElasticIndex).
		Query(elastic.NewQueryStringQuery(sc.q)).
		Do(context.Background())

	// empty result will returned when error occurred
	if err != nil {
		log.Printf("elastic search err: %v", err)
		return result
	}

	result.Hits = resp.TotalHits()
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.Query = sc.q

	return result
}
