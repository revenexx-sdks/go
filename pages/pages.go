package pages

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Pages service
type Pages struct {
	client client.Client
}

func New(clt client.Client) *Pages {
	return &Pages{
		client: clt,
	}
}

type PagesLibraryListOptions struct {
	Limit int
	Offset int
	Order string
	Bundles string
	Text string
	enabledSetters map[string]bool
}
func (options PagesLibraryListOptions) New() *PagesLibraryListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Bundles": false,
		"Text": false,
	}
	return &options
}
type PagesLibraryListOption func(*PagesLibraryListOptions)
func (srv *Pages) WithPagesLibraryListLimit(v int) PagesLibraryListOption {
	return func(o *PagesLibraryListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Pages) WithPagesLibraryListOffset(v int) PagesLibraryListOption {
	return func(o *PagesLibraryListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Pages) WithPagesLibraryListOrder(v string) PagesLibraryListOption {
	return func(o *PagesLibraryListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *Pages) WithPagesLibraryListBundles(v string) PagesLibraryListOption {
	return func(o *PagesLibraryListOptions) {
		o.Bundles = v
		o.enabledSetters["Bundles"] = true
	}
}
func (srv *Pages) WithPagesLibraryListText(v string) PagesLibraryListOption {
	return func(o *PagesLibraryListOptions) {
		o.Text = v
		o.enabledSetters["Text"] = true
	}
}
	
// PagesLibraryList the pool an editor picks a reusable block from. A library
// item is ONE block subtree that many pages share BY REFERENCE — edit the
// item and every page using it changes — which is what separates it from a
// template, the other reusable thing here, which copies instead and is at
// `GET /pages/templates`. So the two filters are the two questions the picker
// asks: `bundles` narrows to the block types that fit the field being filled,
// `text` matches the label a person gave the item.
func (srv *Pages) PagesLibraryList(optionalSetters ...PagesLibraryListOption)(*interface{}, error) {
	path := "/v1/pages/library"
	options := PagesLibraryListOptions{}.New()
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
	if options.enabledSetters["Bundles"] {
		params["bundles"] = options.Bundles
	}
	if options.enabledSetters["Text"] {
		params["text"] = options.Text
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
	
// PagesLibraryDelete retires a reusable block. It leaves the picker and every
// list, but the blocks pointing at it keep their `library_item_id` — the
// FK's `set null` belongs to a hard delete, and this writes a tombstone.
// Delivery then skips the expansion for a struck item rather than failing on
// it, so a page that used it falls back to the block content stored in its
// own published revision: nothing breaks, but the pages quietly stop tracking
// each other. Nothing here tells you which pages those are, so establish that
// before striking it.
func (srv *Pages) PagesLibraryDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/library/{id}")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
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
	
// PagesLibraryGet the stored subtree behind one reusable block, so a picker
// can preview what dropping it into a page would produce. Because delivery
// expands the reference against THIS row at read time, what comes back is
// also what every page already using the item is currently rendering —
// which makes this the call to make before editing one.
func (srv *Pages) PagesLibraryGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/library/{id}")
	params := map[string]interface{}{}
	params["id"] = Id
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
type PagesLibraryUpdateOptions struct {
	Bundle string
	Label string
	Tree interface{}
	enabledSetters map[string]bool
}
func (options PagesLibraryUpdateOptions) New() *PagesLibraryUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Bundle": false,
		"Label": false,
		"Tree": false,
	}
	return &options
}
type PagesLibraryUpdateOption func(*PagesLibraryUpdateOptions)
func (srv *Pages) WithPagesLibraryUpdateBundle(v string) PagesLibraryUpdateOption {
	return func(o *PagesLibraryUpdateOptions) {
		o.Bundle = v
		o.enabledSetters["Bundle"] = true
	}
}
func (srv *Pages) WithPagesLibraryUpdateLabel(v string) PagesLibraryUpdateOption {
	return func(o *PagesLibraryUpdateOptions) {
		o.Label = v
		o.enabledSetters["Label"] = true
	}
}
func (srv *Pages) WithPagesLibraryUpdateTree(v interface{}) PagesLibraryUpdateOption {
	return func(o *PagesLibraryUpdateOptions) {
		o.Tree = v
		o.enabledSetters["Tree"] = true
	}
}
			
// PagesLibraryUpdate the one write in this app whose blast radius is not a
// single page. Delivery expands a library reference against this row every
// time it serves, so replacing `tree` re-renders every page that points at
// the item — published ones included — without any of them being edited,
// republished or even touched. Nothing warns you first and no revision
// records it, because the pages did not change; the item did. Changing
// `label` or `bundle` only moves the item around the picker. Detaching one
// page from the item, so it keeps a copy of its own, is an editor mutation
// and not this route.
func (srv *Pages) PagesLibraryUpdate(Id string, optionalSetters ...PagesLibraryUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/library/{id}")
	options := PagesLibraryUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Bundle"] {
		params["bundle"] = options.Bundle
	}
	if options.enabledSetters["Label"] {
		params["label"] = options.Label
	}
	if options.enabledSetters["Tree"] {
		params["tree"] = options.Tree
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
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
type PagesMenusListOptions struct {
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options PagesMenusListOptions) New() *PagesMenusListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type PagesMenusListOption func(*PagesMenusListOptions)
func (srv *Pages) WithPagesMenusListLimit(v int) PagesMenusListOption {
	return func(o *PagesMenusListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Pages) WithPagesMenusListOffset(v int) PagesMenusListOption {
	return func(o *PagesMenusListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Pages) WithPagesMenusListOrder(v string) PagesMenusListOption {
	return func(o *PagesMenusListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// PagesMenusList the management view of the menus a tenant keeps — `main`,
// `footer`, `account` and whatever else the theme asks for, each with the key
// it is looked up by. This route reads no filter at all — a `?menu_key=` is
// ignored, which the empty `filter` echo shows — so fetch a page and pick,
// or address one by id.
func (srv *Pages) PagesMenusList(optionalSetters ...PagesMenusListOption)(*interface{}, error) {
	path := "/v1/pages/menus"
	options := PagesMenusListOptions{}.New()
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
type PagesMenusUpsertOptions struct {
	Items []models.PageMenuItem
	enabledSetters map[string]bool
}
func (options PagesMenusUpsertOptions) New() *PagesMenusUpsertOptions {
	options.enabledSetters = map[string]bool{
		"Items": false,
	}
	return &options
}
type PagesMenusUpsertOption func(*PagesMenusUpsertOptions)
func (srv *Pages) WithPagesMenusUpsertItems(v []models.PageMenuItem) PagesMenusUpsertOption {
	return func(o *PagesMenusUpsertOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
					
// PagesMenusUpsert writes a menu by its KEY rather than by its id, which is
// what makes theme seeding safe to repeat: a key the tenant already has has
// its label and items replaced in place, a key it does not have is created.
// `items` is replaced wholesale and never merged, so sending an empty list
// empties the navigation. One caveat worth reading before you rely on the
// idempotence: the key's uniqueness is this route's doing and not the
// database's — `menu_key` carries an index but no unique constraint — so
// a duplicate key created any other way leaves this route updating whichever
// row it finds first.
func (srv *Pages) PagesMenusUpsert(Label string, MenuKey string, optionalSetters ...PagesMenusUpsertOption)(*models.Error, error) {
	path := "/v1/pages/menus"
	options := PagesMenusUpsertOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["label"] = Label
	params["menuKey"] = MenuKey
	if options.enabledSetters["Items"] {
		params["items"] = options.Items
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
	
// PagesMenusDelete writes the tombstone. The menu drops out of the management
// list and out of `GET /pages/delivery/menus` in the same moment, so a theme
// that reads its key gets nothing back and renders nothing — there is no
// fallback and no error a storefront could act on. The key is free
// immediately, which means re-seeding the theme is the way back. Check what
// reads the key before striking it.
func (srv *Pages) PagesMenusDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/menus/{id}")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
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
	
// PagesMenusGet one menu and its whole item tree — the ordered links a
// theme renders as its header, footer or account navigation. `items` is
// nested, not one level, so this is the entire navigation for that key in a
// single read. Addressed by ROW ID here; the key a theme knows it by is
// `menu_key` on the body, and the route that works by key is the upsert.
func (srv *Pages) PagesMenusGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/menus/{id}")
	params := map[string]interface{}{}
	params["id"] = Id
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
type PagesMenusUpdateOptions struct {
	Items []models.PageMenuItem
	Label string
	enabledSetters map[string]bool
}
func (options PagesMenusUpdateOptions) New() *PagesMenusUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Items": false,
		"Label": false,
	}
	return &options
}
type PagesMenusUpdateOption func(*PagesMenusUpdateOptions)
func (srv *Pages) WithPagesMenusUpdateItems(v []models.PageMenuItem) PagesMenusUpdateOption {
	return func(o *PagesMenusUpdateOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
func (srv *Pages) WithPagesMenusUpdateLabel(v string) PagesMenusUpdateOption {
	return func(o *PagesMenusUpdateOptions) {
		o.Label = v
		o.enabledSetters["Label"] = true
	}
}
			
// PagesMenusUpdate the same write as the upsert, for a caller that already
// holds the row id — use this when editing a menu a person picked from a
// list, and the upsert when reconciling a theme's defaults. `menu_key` is
// deliberately not editable here: the key is the handle every theme reads the
// menu by, so changing it would empty whatever is rendering that key without
// anything reporting an error.
func (srv *Pages) PagesMenusUpdate(Id string, optionalSetters ...PagesMenusUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/menus/{id}")
	options := PagesMenusUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Items"] {
		params["items"] = options.Items
	}
	if options.enabledSetters["Label"] {
		params["label"] = options.Label
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
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
type PagesPagesListOptions struct {
	Limit int
	Offset int
	Order string
	Bundle string
	Status string
	Q string
	enabledSetters map[string]bool
}
func (options PagesPagesListOptions) New() *PagesPagesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Bundle": false,
		"Status": false,
		"Q": false,
	}
	return &options
}
type PagesPagesListOption func(*PagesPagesListOptions)
func (srv *Pages) WithPagesPagesListLimit(v int) PagesPagesListOption {
	return func(o *PagesPagesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Pages) WithPagesPagesListOffset(v int) PagesPagesListOption {
	return func(o *PagesPagesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Pages) WithPagesPagesListOrder(v string) PagesPagesListOption {
	return func(o *PagesPagesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *Pages) WithPagesPagesListBundle(v string) PagesPagesListOption {
	return func(o *PagesPagesListOptions) {
		o.Bundle = v
		o.enabledSetters["Bundle"] = true
	}
}
func (srv *Pages) WithPagesPagesListStatus(v string) PagesPagesListOption {
	return func(o *PagesPagesListOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Pages) WithPagesPagesListQ(v string) PagesPagesListOption {
	return func(o *PagesPagesListOptions) {
		o.Q = v
		o.enabledSetters["Q"] = true
	}
}
	
// PagesPagesList the EDITORIAL index — every live page of the tenant,
// whatever its status, newest change first. This is the list the Cockpit
// shows a person: drafts and archived pages are in it, and a row here says
// nothing about whether a visitor can see the page, because a published
// status without a published revision still delivers nothing. A storefront
// wants `GET /pages/delivery/pages` instead, which answers only what is
// actually servable. Soft-deleted pages are never returned and the predicate
// is this route's own, not something a caller can switch off.
func (srv *Pages) PagesPagesList(optionalSetters ...PagesPagesListOption)(*interface{}, error) {
	path := "/v1/pages/pages"
	options := PagesPagesListOptions{}.New()
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
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Q"] {
		params["q"] = options.Q
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
type PagesPagesCreateOptions struct {
	Bundle string
	HostOptions interface{}
	Meta interface{}
	Slug string
	SourceLanguage string
	enabledSetters map[string]bool
}
func (options PagesPagesCreateOptions) New() *PagesPagesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Bundle": false,
		"HostOptions": false,
		"Meta": false,
		"Slug": false,
		"SourceLanguage": false,
	}
	return &options
}
type PagesPagesCreateOption func(*PagesPagesCreateOptions)
func (srv *Pages) WithPagesPagesCreateBundle(v string) PagesPagesCreateOption {
	return func(o *PagesPagesCreateOptions) {
		o.Bundle = v
		o.enabledSetters["Bundle"] = true
	}
}
func (srv *Pages) WithPagesPagesCreateHostOptions(v interface{}) PagesPagesCreateOption {
	return func(o *PagesPagesCreateOptions) {
		o.HostOptions = v
		o.enabledSetters["HostOptions"] = true
	}
}
func (srv *Pages) WithPagesPagesCreateMeta(v interface{}) PagesPagesCreateOption {
	return func(o *PagesPagesCreateOptions) {
		o.Meta = v
		o.enabledSetters["Meta"] = true
	}
}
func (srv *Pages) WithPagesPagesCreateSlug(v string) PagesPagesCreateOption {
	return func(o *PagesPagesCreateOptions) {
		o.Slug = v
		o.enabledSetters["Slug"] = true
	}
}
func (srv *Pages) WithPagesPagesCreateSourceLanguage(v string) PagesPagesCreateOption {
	return func(o *PagesPagesCreateOptions) {
		o.SourceLanguage = v
		o.enabledSetters["SourceLanguage"] = true
	}
}
			
// PagesPagesCreate writes two rows, not one: the page itself and the
// translation row for its source language, so a page is never without the
// language it was authored in and `GET /pages/delivery/page?slug=` can match
// a localized URL from the first moment. Everything the caller leaves out
// comes from the tenant's settings, not from a literal in this app: `bundle`
// from default_page_bundle, `sourceLanguage` from default_source_language
// (resolved for the request's market), and the status of both the page and
// its source translation from default_page_status (draft | published).
func (srv *Pages) PagesPagesCreate(Title string, optionalSetters ...PagesPagesCreateOption)(*models.Error, error) {
	path := "/v1/pages/pages"
	options := PagesPagesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["title"] = Title
	if options.enabledSetters["Bundle"] {
		params["bundle"] = options.Bundle
	}
	if options.enabledSetters["HostOptions"] {
		params["hostOptions"] = options.HostOptions
	}
	if options.enabledSetters["Meta"] {
		params["meta"] = options.Meta
	}
	if options.enabledSetters["Slug"] {
		params["slug"] = options.Slug
	}
	if options.enabledSetters["SourceLanguage"] {
		params["sourceLanguage"] = options.SourceLanguage
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
	
// PagesPagesDelete writes a tombstone. The page leaves every list, every read
// and all delivery at once, and its slug is immediately free for another page
// — the unique index counts live rows only. Nothing is erased: the
// translations, blocks, edit state, revisions, comments and preview grants
// that hang off the page all keep their rows, because their `on delete
// cascade` belongs to a hard delete and this is not one. So a page can be
// brought back intact by clearing `deleted_at` — but not through this app,
// which publishes no route that does it.
func (srv *Pages) PagesPagesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/pages/{id}")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
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
	
// PagesPagesGet one page RECORD: what it is called, where it routes, what
// type it is, which revision is live. Not its content — the blocks are not
// on this row and no expansion here returns them. The editor reads them with
// `GET /pages/editor/{page_id}/state`, a renderer with `GET
// /pages/delivery/page`. A soft-deleted page answers 404 exactly like one
// that never existed, so this is also the check for whether an id is still
// good.
func (srv *Pages) PagesPagesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/pages/{id}")
	params := map[string]interface{}{}
	params["id"] = Id
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
type PagesPagesUpdateOptions struct {
	Bundle string
	Meta interface{}
	Slug string
	Status string
	Title string
	enabledSetters map[string]bool
}
func (options PagesPagesUpdateOptions) New() *PagesPagesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Bundle": false,
		"Meta": false,
		"Slug": false,
		"Status": false,
		"Title": false,
	}
	return &options
}
type PagesPagesUpdateOption func(*PagesPagesUpdateOptions)
func (srv *Pages) WithPagesPagesUpdateBundle(v string) PagesPagesUpdateOption {
	return func(o *PagesPagesUpdateOptions) {
		o.Bundle = v
		o.enabledSetters["Bundle"] = true
	}
}
func (srv *Pages) WithPagesPagesUpdateMeta(v interface{}) PagesPagesUpdateOption {
	return func(o *PagesPagesUpdateOptions) {
		o.Meta = v
		o.enabledSetters["Meta"] = true
	}
}
func (srv *Pages) WithPagesPagesUpdateSlug(v string) PagesPagesUpdateOption {
	return func(o *PagesPagesUpdateOptions) {
		o.Slug = v
		o.enabledSetters["Slug"] = true
	}
}
func (srv *Pages) WithPagesPagesUpdateStatus(v string) PagesPagesUpdateOption {
	return func(o *PagesPagesUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Pages) WithPagesPagesUpdateTitle(v string) PagesPagesUpdateOption {
	return func(o *PagesPagesUpdateOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
			
// PagesPagesUpdate corrects the page RECORD — the five fields an editor
// changes without opening the visual editor, which are `title`, `slug`,
// `status`, `meta` and `bundle`, and no others. Anything else in the body is
// dropped rather than refused, and the block tree is unreachable from here by
// design: content moves only through the editor's mutation log, so a caller
// cannot half-edit a page behind the undo history's back. Two consequences
// worth knowing before you call it: a slug is unique among live pages, so
// claiming one that is held answers 409; and setting `status` to published
// does NOT put anything in front of a visitor — delivery needs a revision,
// which only `POST /pages/editor/{page_id}/publish` writes.
func (srv *Pages) PagesPagesUpdate(Id string, optionalSetters ...PagesPagesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/pages/{id}")
	options := PagesPagesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Bundle"] {
		params["bundle"] = options.Bundle
	}
	if options.enabledSetters["Meta"] {
		params["meta"] = options.Meta
	}
	if options.enabledSetters["Slug"] {
		params["slug"] = options.Slug
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
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
type PagesPagesRevisionsOptions struct {
	Limit int
	Offset int
	Order string
	Label string
	CreatedBy string
	CreatedByName string
	CreatedAt string
	enabledSetters map[string]bool
}
func (options PagesPagesRevisionsOptions) New() *PagesPagesRevisionsOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Label": false,
		"CreatedBy": false,
		"CreatedByName": false,
		"CreatedAt": false,
	}
	return &options
}
type PagesPagesRevisionsOption func(*PagesPagesRevisionsOptions)
func (srv *Pages) WithPagesPagesRevisionsLimit(v int) PagesPagesRevisionsOption {
	return func(o *PagesPagesRevisionsOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Pages) WithPagesPagesRevisionsOffset(v int) PagesPagesRevisionsOption {
	return func(o *PagesPagesRevisionsOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Pages) WithPagesPagesRevisionsOrder(v string) PagesPagesRevisionsOption {
	return func(o *PagesPagesRevisionsOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *Pages) WithPagesPagesRevisionsLabel(v string) PagesPagesRevisionsOption {
	return func(o *PagesPagesRevisionsOptions) {
		o.Label = v
		o.enabledSetters["Label"] = true
	}
}
func (srv *Pages) WithPagesPagesRevisionsCreatedBy(v string) PagesPagesRevisionsOption {
	return func(o *PagesPagesRevisionsOptions) {
		o.CreatedBy = v
		o.enabledSetters["CreatedBy"] = true
	}
}
func (srv *Pages) WithPagesPagesRevisionsCreatedByName(v string) PagesPagesRevisionsOption {
	return func(o *PagesPagesRevisionsOptions) {
		o.CreatedByName = v
		o.enabledSetters["CreatedByName"] = true
	}
}
func (srv *Pages) WithPagesPagesRevisionsCreatedAt(v string) PagesPagesRevisionsOption {
	return func(o *PagesPagesRevisionsOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
			
// PagesPagesRevisions one entry per publication, newest first, which is the
// order a history is read in and the one this route sorts by unless `order`
// says otherwise. The `snapshot` — the whole published page, in every
// language — is deliberately not in the index: it is page-sized, and
// nothing that renders a history needs it.
func (srv *Pages) PagesPagesRevisions(Id string, optionalSetters ...PagesPagesRevisionsOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/pages/{id}/revisions")
	options := PagesPagesRevisionsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
	}
	if options.enabledSetters["Offset"] {
		params["offset"] = options.Offset
	}
	if options.enabledSetters["Order"] {
		params["order"] = options.Order
	}
	if options.enabledSetters["Label"] {
		params["label"] = options.Label
	}
	if options.enabledSetters["CreatedBy"] {
		params["created_by"] = options.CreatedBy
	}
	if options.enabledSetters["CreatedByName"] {
		params["created_by_name"] = options.CreatedByName
	}
	if options.enabledSetters["CreatedAt"] {
		params["created_at"] = options.CreatedAt
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
type PagesSeedOptions struct {
	Menus []interface{}
	Pages []interface{}
	enabledSetters map[string]bool
}
func (options PagesSeedOptions) New() *PagesSeedOptions {
	options.enabledSetters = map[string]bool{
		"Menus": false,
		"Pages": false,
	}
	return &options
}
type PagesSeedOption func(*PagesSeedOptions)
func (srv *Pages) WithPagesSeedMenus(v []interface{}) PagesSeedOption {
	return func(o *PagesSeedOptions) {
		o.Menus = v
		o.enabledSetters["Menus"] = true
	}
}
func (srv *Pages) WithPagesSeedPages(v []interface{}) PagesSeedOption {
	return func(o *PagesSeedOptions) {
		o.Pages = v
		o.enabledSetters["Pages"] = true
	}
}
	
// PagesSeed the target of a theme activation hook: hand it the theme's
// default pages and menus and it creates whatever is missing. Idempotent by
// `slug` and by menu key — a slug or a key the tenant already holds is
// skipped rather than rewritten, so re-running after a theme update adds only
// the new ones and never overwrites what an editor has since changed. A
// seeded page is published on the spot, immediately servable by delivery: the
// default_page_status setting deliberately does not apply, because a theme
// that activates with invisible pages looks broken.
func (srv *Pages) PagesSeed(optionalSetters ...PagesSeedOption)(*models.SeedResult, error) {
	path := "/v1/pages/seed"
	options := PagesSeedOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Menus"] {
		params["menus"] = options.Menus
	}
	if options.enabledSetters["Pages"] {
		params["pages"] = options.Pages
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

		parsed := models.SeedResult{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.SeedResult
	parsed, ok := resp.Result.(models.SeedResult)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PagesTemplatesListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	Label string
	Description string
	PageBundle string
	FieldName string
	IsDefault bool
	CreatedBy string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options PagesTemplatesListOptions) New() *PagesTemplatesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"Label": false,
		"Description": false,
		"PageBundle": false,
		"FieldName": false,
		"IsDefault": false,
		"CreatedBy": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type PagesTemplatesListOption func(*PagesTemplatesListOptions)
func (srv *Pages) WithPagesTemplatesListLimit(v int) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Pages) WithPagesTemplatesListOffset(v int) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Pages) WithPagesTemplatesListOrder(v string) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *Pages) WithPagesTemplatesListId(v string) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Pages) WithPagesTemplatesListLabel(v string) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.Label = v
		o.enabledSetters["Label"] = true
	}
}
func (srv *Pages) WithPagesTemplatesListDescription(v string) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Pages) WithPagesTemplatesListPageBundle(v string) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.PageBundle = v
		o.enabledSetters["PageBundle"] = true
	}
}
func (srv *Pages) WithPagesTemplatesListFieldName(v string) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.FieldName = v
		o.enabledSetters["FieldName"] = true
	}
}
func (srv *Pages) WithPagesTemplatesListIsDefault(v bool) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Pages) WithPagesTemplatesListCreatedBy(v string) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.CreatedBy = v
		o.enabledSetters["CreatedBy"] = true
	}
}
func (srv *Pages) WithPagesTemplatesListCreatedAt(v string) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Pages) WithPagesTemplatesListUpdatedAt(v string) PagesTemplatesListOption {
	return func(o *PagesTemplatesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// PagesTemplatesList every column of a template is an exact-match filter
// here: `?page_bundle=standard&field_name=content` is how a picker asks for
// the templates offered in one place, and `?is_default=true` is how a "new
// page" flow finds the one to start from.
func (srv *Pages) PagesTemplatesList(optionalSetters ...PagesTemplatesListOption)(*interface{}, error) {
	path := "/v1/pages/templates"
	options := PagesTemplatesListOptions{}.New()
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
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Label"] {
		params["label"] = options.Label
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["PageBundle"] {
		params["page_bundle"] = options.PageBundle
	}
	if options.enabledSetters["FieldName"] {
		params["field_name"] = options.FieldName
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["CreatedBy"] {
		params["created_by"] = options.CreatedBy
	}
	if options.enabledSetters["CreatedAt"] {
		params["created_at"] = options.CreatedAt
	}
	if options.enabledSetters["UpdatedAt"] {
		params["updated_at"] = options.UpdatedAt
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
	
// PagesTemplatesDelete removes the template row outright. This is the one
// delete in the app that is not a tombstone — `templates` carries no
// `deleted_at` — so it cannot be undone and the id will not come back.
// Nothing else breaks by it: pages built from the template hold their own
// copy of the blocks and never referenced the row.
func (srv *Pages) PagesTemplatesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/templates/{id}")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
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
	
// PagesTemplatesGet the blocks a page would START from if an editor picked
// this template — read it to preview the insert. A template is a COPY
// source, the opposite of a library item: nothing links back from the pages
// already built from it, so this tells you what future pages get and nothing
// about existing ones.
func (srv *Pages) PagesTemplatesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/templates/{id}")
	params := map[string]interface{}{}
	params["id"] = Id
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
type PagesTemplatesUpdateOptions struct {
	Description string
	FieldName string
	IsDefault bool
	Label string
	PageBundle string
	Tree []models.PageBlockTree
	enabledSetters map[string]bool
}
func (options PagesTemplatesUpdateOptions) New() *PagesTemplatesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"FieldName": false,
		"IsDefault": false,
		"Label": false,
		"PageBundle": false,
		"Tree": false,
	}
	return &options
}
type PagesTemplatesUpdateOption func(*PagesTemplatesUpdateOptions)
func (srv *Pages) WithPagesTemplatesUpdateDescription(v string) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Pages) WithPagesTemplatesUpdateFieldName(v string) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.FieldName = v
		o.enabledSetters["FieldName"] = true
	}
}
func (srv *Pages) WithPagesTemplatesUpdateIsDefault(v bool) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Pages) WithPagesTemplatesUpdateLabel(v string) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.Label = v
		o.enabledSetters["Label"] = true
	}
}
func (srv *Pages) WithPagesTemplatesUpdatePageBundle(v string) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.PageBundle = v
		o.enabledSetters["PageBundle"] = true
	}
}
func (srv *Pages) WithPagesTemplatesUpdateTree(v []models.PageBlockTree) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.Tree = v
		o.enabledSetters["Tree"] = true
	}
}
			
// PagesTemplatesUpdate edits what a future page will start from. Because
// templates copy rather than share, this reaches nothing that already exists
// — pages built from it keep the blocks they were handed, which is exactly
// the property that makes a template safe to edit and a library item
// dangerous. `is_default` is the one field with an effect past the picker: it
// decides what a new page of `page_bundle` starts with, and nothing here
// stops two templates of the same bundle from both claiming it, so which one
// wins is left to whoever reads the list.
func (srv *Pages) PagesTemplatesUpdate(Id string, optionalSetters ...PagesTemplatesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/templates/{id}")
	options := PagesTemplatesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["FieldName"] {
		params["field_name"] = options.FieldName
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Label"] {
		params["label"] = options.Label
	}
	if options.enabledSetters["PageBundle"] {
		params["page_bundle"] = options.PageBundle
	}
	if options.enabledSetters["Tree"] {
		params["tree"] = options.Tree
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PUT", path, headers, params)
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

// PagesVocabulariesList discovery for the vocabulary routes: the enums this
// app publishes, each with its name, its title and what it is for, and none
// of them unpacked — the permitted values are not on this route, only on
// the one that serves a single vocabulary. Names: edit-state-statuses,
// page-statuses, translation-statuses. Fetch one with GET
// /pages/vocabularies/{name}; a client holding the qualified pair
// 'pages.<name>' builds that URL from the pair alone.
func (srv *Pages) PagesVocabulariesList()(*models.PagesVocabularyIndex, error) {
	path := "/v1/pages/vocabularies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.PagesVocabularyIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PagesVocabularyIndex
	parsed, ok := resp.Result.(models.PagesVocabularyIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PagesVocabulariesGet one vocabulary unpacked: every value the column
// permits, each with the title to show for it, the sentence explaining it and
// the badge tone to render it in — everything a select or a status pill
// needs, so nothing downstream keeps its own copy of the labels. The values
// are read out of the column's CHECK constraint, so the served set IS the
// enforced set and the two cannot drift — a value added to the constraint
// appears here even before anyone labels it, titled from its own key. Values
// come back in constraint order, which is the order a select should offer.
// 'closed' says the set is exhaustive, so a value outside it is stale data
// rather than a missing label. Names: edit-state-statuses, page-statuses,
// translation-statuses.
func (srv *Pages) PagesVocabulariesGet(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/pages/vocabularies/{name}")
	params := map[string]interface{}{}
	params["name"] = Name
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
