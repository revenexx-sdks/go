package pages_delivery

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// PagesDelivery service
type PagesDelivery struct {
	client client.Client
}

func New(clt client.Client) *PagesDelivery {
	return &PagesDelivery{
		client: clt,
	}
}

type PagesDeliveryMenusOptions struct {
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options PagesDeliveryMenusOptions) New() *PagesDeliveryMenusOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type PagesDeliveryMenusOption func(*PagesDeliveryMenusOptions)
func (srv *PagesDelivery) WithPagesDeliveryMenusLimit(v int) PagesDeliveryMenusOption {
	return func(o *PagesDeliveryMenusOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *PagesDelivery) WithPagesDeliveryMenusOffset(v int) PagesDeliveryMenusOption {
	return func(o *PagesDeliveryMenusOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *PagesDelivery) WithPagesDeliveryMenusOrder(v string) PagesDeliveryMenusOption {
	return func(o *PagesDeliveryMenusOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// PagesDeliveryMenus one call gives a theme its whole chrome: header, footer
// and account navigation, each under the key the theme looks it up by. This
// route reads no filter — fetch all of them once and index by `id`.
func (srv *PagesDelivery) PagesDeliveryMenus(optionalSetters ...PagesDeliveryMenusOption)(*interface{}, error) {
	path := "/v1/pages/delivery/menus"
	options := PagesDeliveryMenusOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Order"] {
		params["order"] = options.Order
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
type PagesDeliveryPageOptions struct {
	Slug string
	Id string
	Langcode string
	enabledSetters map[string]bool
}
func (options PagesDeliveryPageOptions) New() *PagesDeliveryPageOptions {
	options.enabledSetters = map[string]bool{
		"Slug": false,
		"Id": false,
		"Langcode": false,
	}
	return &options
}
type PagesDeliveryPageOption func(*PagesDeliveryPageOptions)
func (srv *PagesDelivery) WithPagesDeliveryPageSlug(v string) PagesDeliveryPageOption {
	return func(o *PagesDeliveryPageOptions) {
		o.Slug = v
		o.enabledSetters["Slug"] = true
	}
}
func (srv *PagesDelivery) WithPagesDeliveryPageId(v string) PagesDeliveryPageOption {
	return func(o *PagesDeliveryPageOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *PagesDelivery) WithPagesDeliveryPageLangcode(v string) PagesDeliveryPageOption {
	return func(o *PagesDeliveryPageOptions) {
		o.Langcode = v
		o.enabledSetters["Langcode"] = true
	}
}
	
// PagesDeliveryPage what a storefront calls to render a URL: `GET
// /pages/delivery/page?slug=about-us&langcode=de`. Send exactly one selector
// — `slug` or `id`. `slug` is matched against the page and then against its
// translations, so a localized URL resolves to its page. Only the PUBLISHED
// revision is served, so an edit in progress never leaks. What comes back is
// finished rather than raw: `langcode` is resolved field by field with the
// page's source language behind it, blocks whose publish window has not
// opened or has already closed are left out, and every library reference is
// expanded into the subtree it points at — so a renderer walks the tree it
// is given and makes no second call for any of it.
func (srv *PagesDelivery) PagesDeliveryPage(optionalSetters ...PagesDeliveryPageOption)(*models.Error, error) {
	path := "/v1/pages/delivery/page"
	options := PagesDeliveryPageOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Slug"] {
		params["slug"] = options.Slug
	}
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Langcode"] {
		params["langcode"] = options.Langcode
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
type PagesDeliveryPagesOptions struct {
	Limit int
	Offset int
	Order string
	Bundle string
	enabledSetters map[string]bool
}
func (options PagesDeliveryPagesOptions) New() *PagesDeliveryPagesOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Bundle": false,
	}
	return &options
}
type PagesDeliveryPagesOption func(*PagesDeliveryPagesOptions)
func (srv *PagesDelivery) WithPagesDeliveryPagesLimit(v int) PagesDeliveryPagesOption {
	return func(o *PagesDeliveryPagesOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *PagesDelivery) WithPagesDeliveryPagesOffset(v int) PagesDeliveryPagesOption {
	return func(o *PagesDeliveryPagesOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *PagesDelivery) WithPagesDeliveryPagesOrder(v string) PagesDeliveryPagesOption {
	return func(o *PagesDeliveryPagesOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *PagesDelivery) WithPagesDeliveryPagesBundle(v string) PagesDeliveryPagesOption {
	return func(o *PagesDeliveryPagesOptions) {
		o.Bundle = v
		o.enabledSetters["Bundle"] = true
	}
}
	
// PagesDeliveryPages the route a sitemap, a static build or a link picker is
// generated from. Only published pages, never a soft-deleted one — `filter`
// echoes both predicates the route applies on its own. A `?status=` of your
// own is ignored: this route is the published view by definition.
func (srv *PagesDelivery) PagesDeliveryPages(optionalSetters ...PagesDeliveryPagesOption)(*interface{}, error) {
	path := "/v1/pages/delivery/pages"
	options := PagesDeliveryPagesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Order"] {
		params["order"] = options.Order
	}
	if options.enabledSetters["Bundle"] {
		params["bundle"] = options.Bundle
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
type PagesDeliveryPreviewOptions struct {
	Langcode string
	enabledSetters map[string]bool
}
func (options PagesDeliveryPreviewOptions) New() *PagesDeliveryPreviewOptions {
	options.enabledSetters = map[string]bool{
		"Langcode": false,
	}
	return &options
}
type PagesDeliveryPreviewOption func(*PagesDeliveryPreviewOptions)
func (srv *PagesDelivery) WithPagesDeliveryPreviewLangcode(v string) PagesDeliveryPreviewOption {
	return func(o *PagesDeliveryPreviewOptions) {
		o.Langcode = v
		o.enabledSetters["Langcode"] = true
	}
}
			
// PagesDeliveryPreview the same shape `GET /pages/delivery/page` answers,
// built from the UNPUBLISHED working copy instead of the published revision
// — so a reviewer without an editor account sees exactly what the
// storefront would render.
func (srv *PagesDelivery) PagesDeliveryPreview(Token string, optionalSetters ...PagesDeliveryPreviewOption)(*models.Error, error) {
	r := strings.NewReplacer("{token}", Token)
	path := r.Replace("/v1/pages/delivery/preview/{token}")
	options := PagesDeliveryPreviewOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["token"] = Token
	if options.enabledSetters["Langcode"] {
		params["langcode"] = options.Langcode
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
