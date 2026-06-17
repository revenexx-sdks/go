package search

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"strings"
)

// Search service
type Search struct {
	client client.Client
}

func New(clt client.Client) *Search {
	return &Search{
		client: clt,
	}
}


// SearchListCollections the collections the tenant's installed apps have
// provisioned.
func (srv *Search) SearchListCollections()(*interface{}, error) {
	path := "/v1/search/collections"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type SearchSearchDocumentsGetOptions struct {
	Q string
	QueryBy string
	FilterBy string
	SortBy string
	Page int
	PerPage int
	enabledSetters map[string]bool
}
func (options SearchSearchDocumentsGetOptions) New() *SearchSearchDocumentsGetOptions {
	options.enabledSetters = map[string]bool{
		"Q": false,
		"QueryBy": false,
		"FilterBy": false,
		"SortBy": false,
		"Page": false,
		"PerPage": false,
	}
	return &options
}
type SearchSearchDocumentsGetOption func(*SearchSearchDocumentsGetOptions)
func (srv *Search) WithSearchSearchDocumentsGetQ(v string) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.Q = v
		o.enabledSetters["Q"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetQueryBy(v string) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.QueryBy = v
		o.enabledSetters["QueryBy"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetFilterBy(v string) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.FilterBy = v
		o.enabledSetters["FilterBy"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetSortBy(v string) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.SortBy = v
		o.enabledSetters["SortBy"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetPage(v int) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.Page = v
		o.enabledSetters["Page"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetPerPage(v int) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.PerPage = v
		o.enabledSetters["PerPage"] = true
	}
}
			
// SearchSearchDocumentsGet full-text search within one collection using
// Typesense query parameters as the query string.
func (srv *Search) SearchSearchDocumentsGet(Collection string, optionalSetters ...SearchSearchDocumentsGetOption)(*interface{}, error) {
	r := strings.NewReplacer("{collection}", Collection)
	path := r.Replace("/v1/search/collections/{collection}/documents/search")
	options := SearchSearchDocumentsGetOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["collection"] = Collection
	if options.enabledSetters["Q"] {
		params["q"] = options.Q
	}
	if options.enabledSetters["QueryBy"] {
		params["query_by"] = options.QueryBy
	}
	if options.enabledSetters["FilterBy"] {
		params["filter_by"] = options.FilterBy
	}
	if options.enabledSetters["SortBy"] {
		params["sort_by"] = options.SortBy
	}
	if options.enabledSetters["Page"] {
		params["page"] = options.Page
	}
	if options.enabledSetters["PerPage"] {
		params["per_page"] = options.PerPage
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type SearchSearchDocumentsOptions struct {
	FacetBy string
	FilterBy string
	Page int
	PerPage int
	Q string
	QueryBy string
	SortBy string
	enabledSetters map[string]bool
}
func (options SearchSearchDocumentsOptions) New() *SearchSearchDocumentsOptions {
	options.enabledSetters = map[string]bool{
		"FacetBy": false,
		"FilterBy": false,
		"Page": false,
		"PerPage": false,
		"Q": false,
		"QueryBy": false,
		"SortBy": false,
	}
	return &options
}
type SearchSearchDocumentsOption func(*SearchSearchDocumentsOptions)
func (srv *Search) WithSearchSearchDocumentsFacetBy(v string) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.FacetBy = v
		o.enabledSetters["FacetBy"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsFilterBy(v string) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.FilterBy = v
		o.enabledSetters["FilterBy"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsPage(v int) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.Page = v
		o.enabledSetters["Page"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsPerPage(v int) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.PerPage = v
		o.enabledSetters["PerPage"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsQ(v string) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.Q = v
		o.enabledSetters["Q"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsQueryBy(v string) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.QueryBy = v
		o.enabledSetters["QueryBy"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsSortBy(v string) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.SortBy = v
		o.enabledSetters["SortBy"] = true
	}
}
			
// SearchSearchDocuments full-text search within one collection. The body
// holds Typesense search parameters.
func (srv *Search) SearchSearchDocuments(Collection string, optionalSetters ...SearchSearchDocumentsOption)(*interface{}, error) {
	r := strings.NewReplacer("{collection}", Collection)
	path := r.Replace("/v1/search/collections/{collection}/documents/search")
	options := SearchSearchDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["collection"] = Collection
	if options.enabledSetters["FacetBy"] {
		params["facet_by"] = options.FacetBy
	}
	if options.enabledSetters["FilterBy"] {
		params["filter_by"] = options.FilterBy
	}
	if options.enabledSetters["Page"] {
		params["page"] = options.Page
	}
	if options.enabledSetters["PerPage"] {
		params["per_page"] = options.PerPage
	}
	if options.enabledSetters["Q"] {
		params["q"] = options.Q
	}
	if options.enabledSetters["QueryBy"] {
		params["query_by"] = options.QueryBy
	}
	if options.enabledSetters["SortBy"] {
		params["sort_by"] = options.SortBy
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// SearchGetDocument fetch a single document by id from a collection the
// tenant has installed.
func (srv *Search) SearchGetDocument(Collection string, DocumentId string)(*interface{}, error) {
	r := strings.NewReplacer("{collection}", Collection, "{documentId}", DocumentId)
	path := r.Replace("/v1/search/collections/{collection}/documents/{documentId}")
	params := map[string]interface{}{}
	params["collection"] = Collection
	params["documentId"] = DocumentId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// SearchMultiSearch run several searches in one request (the InstantSearch
// adapter uses this). Each entry names its collection.
func (srv *Search) SearchMultiSearch(Searches []interface{})(*interface{}, error) {
	path := "/v1/search/multi_search"
	params := map[string]interface{}{}
	params["searches"] = Searches
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		var parsed interface{}

		err = json.Unmarshal(bytes, &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	var parsed interface{}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
