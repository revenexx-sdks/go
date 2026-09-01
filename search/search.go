package search

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
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
// provisioned. Available on the API-gateway-trust path only — a `revx_` key
// authorises a single collection, so discovery is a gateway concern and a
// key-authenticated caller gets 403.
func (srv *Search) SearchListCollections()(*models.Error, error) {
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// SearchGetCollection returns the Typesense collection definition (fields,
// defaults, document count). Requires the `collections:read` action.
func (srv *Search) SearchGetCollection(Collection string)(*models.Error, error) {
	r := strings.NewReplacer("{collection}", Collection)
	path := r.Replace("/v1/search/collections/{collection}")
	params := map[string]interface{}{}
	params["collection"] = Collection
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
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
	FacetBy string
	MaxFacetValues int
	GroupBy string
	IncludeFields string
	ExcludeFields string
	HighlightFullFields string
	NumTypos int
	Prefix string
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
		"FacetBy": false,
		"MaxFacetValues": false,
		"GroupBy": false,
		"IncludeFields": false,
		"ExcludeFields": false,
		"HighlightFullFields": false,
		"NumTypos": false,
		"Prefix": false,
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
func (srv *Search) WithSearchSearchDocumentsGetFacetBy(v string) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.FacetBy = v
		o.enabledSetters["FacetBy"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetMaxFacetValues(v int) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.MaxFacetValues = v
		o.enabledSetters["MaxFacetValues"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetGroupBy(v string) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.GroupBy = v
		o.enabledSetters["GroupBy"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetIncludeFields(v string) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.IncludeFields = v
		o.enabledSetters["IncludeFields"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetExcludeFields(v string) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.ExcludeFields = v
		o.enabledSetters["ExcludeFields"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetHighlightFullFields(v string) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.HighlightFullFields = v
		o.enabledSetters["HighlightFullFields"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetNumTypos(v int) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.NumTypos = v
		o.enabledSetters["NumTypos"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsGetPrefix(v string) SearchSearchDocumentsGetOption {
	return func(o *SearchSearchDocumentsGetOptions) {
		o.Prefix = v
		o.enabledSetters["Prefix"] = true
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
			
// SearchSearchDocumentsGet full-text search within one collection. Typesense
// search parameters are passed through verbatim as the query string, so
// parameters not listed here still reach Typesense. Requires the
// `documents:search` action.
func (srv *Search) SearchSearchDocumentsGet(Collection string, optionalSetters ...SearchSearchDocumentsGetOption)(*models.Error, error) {
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
	if options.enabledSetters["FacetBy"] {
		params["facet_by"] = options.FacetBy
	}
	if options.enabledSetters["MaxFacetValues"] {
		params["max_facet_values"] = options.MaxFacetValues
	}
	if options.enabledSetters["GroupBy"] {
		params["group_by"] = options.GroupBy
	}
	if options.enabledSetters["IncludeFields"] {
		params["include_fields"] = options.IncludeFields
	}
	if options.enabledSetters["ExcludeFields"] {
		params["exclude_fields"] = options.ExcludeFields
	}
	if options.enabledSetters["HighlightFullFields"] {
		params["highlight_full_fields"] = options.HighlightFullFields
	}
	if options.enabledSetters["NumTypos"] {
		params["num_typos"] = options.NumTypos
	}
	if options.enabledSetters["Prefix"] {
		params["prefix"] = options.Prefix
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type SearchSearchDocumentsOptions struct {
	ExcludeFields string
	FacetBy string
	FilterBy string
	GroupBy string
	HighlightFullFields string
	IncludeFields string
	MaxFacetValues int
	NumTypos int
	Page int
	PerPage int
	Prefix string
	Q string
	QueryBy string
	SortBy string
	enabledSetters map[string]bool
}
func (options SearchSearchDocumentsOptions) New() *SearchSearchDocumentsOptions {
	options.enabledSetters = map[string]bool{
		"ExcludeFields": false,
		"FacetBy": false,
		"FilterBy": false,
		"GroupBy": false,
		"HighlightFullFields": false,
		"IncludeFields": false,
		"MaxFacetValues": false,
		"NumTypos": false,
		"Page": false,
		"PerPage": false,
		"Prefix": false,
		"Q": false,
		"QueryBy": false,
		"SortBy": false,
	}
	return &options
}
type SearchSearchDocumentsOption func(*SearchSearchDocumentsOptions)
func (srv *Search) WithSearchSearchDocumentsExcludeFields(v string) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.ExcludeFields = v
		o.enabledSetters["ExcludeFields"] = true
	}
}
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
func (srv *Search) WithSearchSearchDocumentsGroupBy(v string) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.GroupBy = v
		o.enabledSetters["GroupBy"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsHighlightFullFields(v string) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.HighlightFullFields = v
		o.enabledSetters["HighlightFullFields"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsIncludeFields(v string) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.IncludeFields = v
		o.enabledSetters["IncludeFields"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsMaxFacetValues(v int) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.MaxFacetValues = v
		o.enabledSetters["MaxFacetValues"] = true
	}
}
func (srv *Search) WithSearchSearchDocumentsNumTypos(v int) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.NumTypos = v
		o.enabledSetters["NumTypos"] = true
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
func (srv *Search) WithSearchSearchDocumentsPrefix(v string) SearchSearchDocumentsOption {
	return func(o *SearchSearchDocumentsOptions) {
		o.Prefix = v
		o.enabledSetters["Prefix"] = true
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
			
// SearchSearchDocuments full-text search within one collection, with the
// Typesense search parameters in the body. Requires the `documents:search`
// action.
func (srv *Search) SearchSearchDocuments(Collection string, optionalSetters ...SearchSearchDocumentsOption)(*models.Error, error) {
	r := strings.NewReplacer("{collection}", Collection)
	path := r.Replace("/v1/search/collections/{collection}/documents/search")
	options := SearchSearchDocumentsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["collection"] = Collection
	if options.enabledSetters["ExcludeFields"] {
		params["exclude_fields"] = options.ExcludeFields
	}
	if options.enabledSetters["FacetBy"] {
		params["facet_by"] = options.FacetBy
	}
	if options.enabledSetters["FilterBy"] {
		params["filter_by"] = options.FilterBy
	}
	if options.enabledSetters["GroupBy"] {
		params["group_by"] = options.GroupBy
	}
	if options.enabledSetters["HighlightFullFields"] {
		params["highlight_full_fields"] = options.HighlightFullFields
	}
	if options.enabledSetters["IncludeFields"] {
		params["include_fields"] = options.IncludeFields
	}
	if options.enabledSetters["MaxFacetValues"] {
		params["max_facet_values"] = options.MaxFacetValues
	}
	if options.enabledSetters["NumTypos"] {
		params["num_typos"] = options.NumTypos
	}
	if options.enabledSetters["Page"] {
		params["page"] = options.Page
	}
	if options.enabledSetters["PerPage"] {
		params["per_page"] = options.PerPage
	}
	if options.enabledSetters["Prefix"] {
		params["prefix"] = options.Prefix
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// SearchGetDocument fetch a single document by id. The document shape is the
// collection's own schema, so it is described as a free-form object. Requires
// the `documents:get` action.
func (srv *Search) SearchGetDocument(Collection string, DocumentId string)(*models.Error, error) {
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type GatewayFacetResyncOptions struct {
	App string
	Vendor string
	enabledSetters map[string]bool
}
func (options GatewayFacetResyncOptions) New() *GatewayFacetResyncOptions {
	options.enabledSetters = map[string]bool{
		"App": false,
		"Vendor": false,
	}
	return &options
}
type GatewayFacetResyncOption func(*GatewayFacetResyncOptions)
func (srv *Search) WithGatewayFacetResyncApp(v string) GatewayFacetResyncOption {
	return func(o *GatewayFacetResyncOptions) {
		o.App = v
		o.enabledSetters["App"] = true
	}
}
func (srv *Search) WithGatewayFacetResyncVendor(v string) GatewayFacetResyncOption {
	return func(o *GatewayFacetResyncOptions) {
		o.Vendor = v
		o.enabledSetters["Vendor"] = true
	}
}
	
// GatewayFacetResync idempotent, and bounded by the tenant's own
// configuration: it can add
// no field for an attribute the tenant has not marked `is_filterable`,
// and drops only fields whose attribute it has itself un-marked. A run
// that changes nothing makes zero calls to Typesense.
// 
// Body (optional) narrows the sweep to one app:
// 
// {"vendor": "revenexx", "app": "products"}
// 
// Omitted, every app the tenant has installed is swept. Apps outside the
// facet-sync allowlist are included in the response with
// `skipped: app_not_enabled` rather than silently dropped — a caller
// asking for an app that cannot have facets deserves to be told so.
// 
// The response shape below is DECLARED rather than inferred. Its entries
// are built by spreading AttributeFacetSyncer::syncForCollection()'s
// summary, and the generator cannot see through an array spread: left to
// itself it emits an unnamed property and a null in `required`, which
// Spectral rejects as `"1" property must be string`.
// AppController::resyncFacets() carries the same declaration for the same
// reason — keep both in step with syncForApp()'s return type.
func (srv *Search) GatewayFacetResync(optionalSetters ...GatewayFacetResyncOption)(*interface{}, error) {
	path := "/v1/search/facets/resync"
	options := GatewayFacetResyncOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["App"] {
		params["app"] = options.App
	}
	if options.enabledSetters["Vendor"] {
		params["vendor"] = options.Vendor
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
	
// SearchMultiSearch run several searches in one round trip — the endpoint
// the typesense-js `multiSearch` helper and the InstantSearch adapter use for
// every query. On the gateway-trust path each entry must name a collection
// the tenant owns. With a `revx_` key `collection_name` is optional and is
// forced to the key's own collection. Requires the `documents:search` action.
func (srv *Search) SearchMultiSearch(Searches []models.MultiSearchEntry)(*models.Error, error) {
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

		parsed := models.Error{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Error
	parsed, ok := resp.Result.(models.Error)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
