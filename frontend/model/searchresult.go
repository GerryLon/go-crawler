package model

type SearchResult struct {
	Hits  int64         // total
	Items []interface{} // current page's items
}
