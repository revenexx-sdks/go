package orderlists

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Orderlists service
type Orderlists struct {
	client client.Client
}

func New(clt client.Client) *Orderlists {
	return &Orderlists{
		client: clt,
	}
}

type OrderlistsListOptions struct {
	OwnerId string
	OrganizationId string
	Kind string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options OrderlistsListOptions) New() *OrderlistsListOptions {
	options.enabledSetters = map[string]bool{
		"OwnerId": false,
		"OrganizationId": false,
		"Kind": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type OrderlistsListOption func(*OrderlistsListOptions)
func (srv *Orderlists) WithOrderlistsListOwnerId(v string) OrderlistsListOption {
	return func(o *OrderlistsListOptions) {
		o.OwnerId = v
		o.enabledSetters["OwnerId"] = true
	}
}
func (srv *Orderlists) WithOrderlistsListOrganizationId(v string) OrderlistsListOption {
	return func(o *OrderlistsListOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Orderlists) WithOrderlistsListKind(v string) OrderlistsListOption {
	return func(o *OrderlistsListOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *Orderlists) WithOrderlistsListLimit(v int) OrderlistsListOption {
	return func(o *OrderlistsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Orderlists) WithOrderlistsListOffset(v int) OrderlistsListOption {
	return func(o *OrderlistsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Orderlists) WithOrderlistsListOrder(v string) OrderlistsListOption {
	return func(o *OrderlistsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// OrderlistsList what a caller may see is a UNION, not an intersection: the
// lists this contact owns, plus the lists their organization shares —
// `owner_id = X OR (organization_id = Y AND shared)`. A list that satisfies
// both sides is merged by id and counted once. Where the gateway resolved an
// acting contact, that contact and their organization ARE the scope and
// neither `owner_id` nor `organization_id` in the query can widen it; without
// a resolved principal — a back-office caller holding the tenant key —
// the two are read from the query, and a call that names neither sees every
// list the tenant keeps. Three filters are read in all — `owner_id`,
// `organization_id`, `kind` — and any OTHER query key is ignored rather
// than refused, which is what the `filter` echo makes visible: a key that is
// missing there was not applied. When only one side of the predicate is in
// play the database pages the rows and reports the true total; when both are,
// each side is read separately and bounded at a thousand rows, merged, and
// paged after the merge, so `total` is the size of the merged set rather than
// a database count. The default sort is `updated_at.desc`, which is why
// adding a position moves its list to the front of the page. Every row
// carries `item_count`. Without it the only way to render a per-list badge
// was to read the positions of every list on the page — thousands of rows
// to draw twenty numbers. The count is bounded the way the page is: at most
// 200 lists, each capped by the tenant's max_items_per_list.
func (srv *Orderlists) OrderlistsList(optionalSetters ...OrderlistsListOption)(*models.Error, error) {
	path := "/v1/orderlists"
	options := OrderlistsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["OwnerId"] {
		params["owner_id"] = options.OwnerId
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
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
type OrderlistsCreateOptions struct {
	Items []models.OrderListItemInput
	Kind string
	Metadata interface{}
	OrganizationId string
	Shared bool
	enabledSetters map[string]bool
}
func (options OrderlistsCreateOptions) New() *OrderlistsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Items": false,
		"Kind": false,
		"Metadata": false,
		"OrganizationId": false,
		"Shared": false,
	}
	return &options
}
type OrderlistsCreateOption func(*OrderlistsCreateOptions)
func (srv *Orderlists) WithOrderlistsCreateItems(v []models.OrderListItemInput) OrderlistsCreateOption {
	return func(o *OrderlistsCreateOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
func (srv *Orderlists) WithOrderlistsCreateKind(v string) OrderlistsCreateOption {
	return func(o *OrderlistsCreateOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *Orderlists) WithOrderlistsCreateMetadata(v interface{}) OrderlistsCreateOption {
	return func(o *OrderlistsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Orderlists) WithOrderlistsCreateOrganizationId(v string) OrderlistsCreateOption {
	return func(o *OrderlistsCreateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *Orderlists) WithOrderlistsCreateShared(v bool) OrderlistsCreateOption {
	return func(o *OrderlistsCreateOptions) {
		o.Shared = v
		o.enabledSetters["Shared"] = true
	}
}
							
// OrderlistsCreate three fields are required, and they are exactly the
// columns the database will not fill in: `name`, `owner_id` and `owner_name`.
// Everything else has an answer already — `kind` resolves to the caller's
// value, else the market's `default_kind` setting, else the kind the tenant
// flagged; `shared` is false; `organization_id` is null, which makes `shared`
// meaningless because there is then nobody to share with. Nothing about a
// list is unique: one owner may keep two lists with the same name, and the
// same article may appear in as many lists as the buyer wants. The list may
// be created empty or pre-filled in the same call: an optional `items` array
// is written as the list's positions with the row, so a twenty-line list is
// one request rather than a create followed by twenty adds, and the array
// order is the position order. Those initial `items` are normalized and
// article-checked BEFORE the list row is written, and both caps are checked
// first as well — the tenant's `max_items_per_list` against the array, and
// its `max_lists_per_owner` against what this contact already keeps — so a
// rejected position never leaves an empty list behind and a contact at their
// limit is refused before anything is inserted. The owner is set once — no
// route moves a list to another contact.
func (srv *Orderlists) OrderlistsCreate(Name string, OwnerId string, OwnerName string, optionalSetters ...OrderlistsCreateOption)(*models.Error, error) {
	path := "/v1/orderlists"
	options := OrderlistsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["owner_id"] = OwnerId
	params["owner_name"] = OwnerName
	if options.enabledSetters["Items"] {
		params["items"] = options.Items
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Shared"] {
		params["shared"] = options.Shared
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

// OrderlistsDefaults seeds the two kinds a fresh tenant starts with —
// `shopping` and `label` — and gives `shopping` the default flag.
// Idempotent by code: `created` names the kinds this call wrote, `existing`
// the ones that were already there and were left exactly as the tenant keeps
// them, renamed, retoned and reordered included. On a settled tenant
// `created` is empty. It is rarely the call you need — the `app.installed`
// event runs the same seed, and the first read of GET /orderlists/kinds on an
// empty table seeds before it answers. It never removes a kind and never
// restores one a merchant deleted.
func (srv *Orderlists) OrderlistsDefaults()(*models.Error, error) {
	path := "/v1/orderlists/defaults"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
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
type OrderlistsKindsListOptions struct {
	Limit int
	Offset int
	enabledSetters map[string]bool
}
func (options OrderlistsKindsListOptions) New() *OrderlistsKindsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
	}
	return &options
}
type OrderlistsKindsListOption func(*OrderlistsKindsListOptions)
func (srv *Orderlists) WithOrderlistsKindsListLimit(v int) OrderlistsKindsListOption {
	return func(o *OrderlistsKindsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsListOffset(v int) OrderlistsKindsListOption {
	return func(o *OrderlistsKindsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
	
// OrderlistsKindsList what a saved list may be FOR — the tenant's own
// taxonomy, and the set every `kind` on a list is drawn from. This used to be
// a CHECK constraint, which meant a merchant who keeps reagent lists or
// sample lists needed a release of this app to say so — and the app never
// branched on the value, it only checked membership. The set is the tenant's
// rows now. Reading this route on a tenant that has none seeds them, so it
// never answers an empty set on a fresh install and a client may treat the
// first read as the install step it no longer has to make. Rows come back in
// `position` order, ascending, which is the order a select should offer them
// in, and each carries the `is_default` flag that decides what a create with
// no `kind` falls back to. It takes NO filters: `limit` and `offset` are the
// only query keys it reads, and any other is ignored rather than refused —
// which is also why this collection alone answers no `filter` echo, since
// echoing an empty one would be noise. The `code` on each row, not the `id`,
// is what `lists.kind` stores and what `?kind=` on GET /orderlists matches.
func (srv *Orderlists) OrderlistsKindsList(optionalSetters ...OrderlistsKindsListOption)(*interface{}, error) {
	path := "/v1/orderlists/kinds"
	options := OrderlistsKindsListOptions{}.New()
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
type OrderlistsKindsCreateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Tone string
	enabledSetters map[string]bool
}
func (options OrderlistsKindsCreateOptions) New() *OrderlistsKindsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"Descriptions": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Tone": false,
	}
	return &options
}
type OrderlistsKindsCreateOption func(*OrderlistsKindsCreateOptions)
func (srv *Orderlists) WithOrderlistsKindsCreateDescription(v string) OrderlistsKindsCreateOption {
	return func(o *OrderlistsKindsCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsCreateDescriptions(v interface{}) OrderlistsKindsCreateOption {
	return func(o *OrderlistsKindsCreateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsCreateIsDefault(v bool) OrderlistsKindsCreateOption {
	return func(o *OrderlistsKindsCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsCreateLabels(v interface{}) OrderlistsKindsCreateOption {
	return func(o *OrderlistsKindsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsCreatePosition(v int) OrderlistsKindsCreateOption {
	return func(o *OrderlistsKindsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsCreateTone(v string) OrderlistsKindsCreateOption {
	return func(o *OrderlistsKindsCreateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
					
// OrderlistsKindsCreate adds a kind to the tenant's own taxonomy — reagent
// lists, sample lists, whatever a merchant sorts their saved lists by —
// without a release of this app, because nothing here branches on the value.
// `code` and `title` are required, and they are exactly the two columns of
// `list_kinds` the database will not fill in. The code is lowercased on the
// way in and immutable afterwards: renaming it would orphan every list
// carrying it, since a list stores the code and not the id. `is_default:
// true` promotes the new kind and demotes whoever held the flag. Creating a
// kind changes no existing list.
func (srv *Orderlists) OrderlistsKindsCreate(Code string, Title string, optionalSetters ...OrderlistsKindsCreateOption)(*models.Error, error) {
	path := "/v1/orderlists/kinds"
	options := OrderlistsKindsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["title"] = Title
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Descriptions"] {
		params["descriptions"] = options.Descriptions
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Tone"] {
		params["tone"] = options.Tone
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
	
// OrderlistsKindsDelete there is no foreign key behind `lists.kind` — it is
// a plain text column holding a code, and nothing in the database points at
// `list_kinds` — so this route's own 409 is the whole of the referential
// integrity. It reads whether any list still carries the code and refuses if
// one does, and refuses again when this is the last kind left, because a list
// must have one. Nothing cascades and no list is rewritten. Two gaps the
// guard leaves: it is a read followed by a delete with no lock between them,
// so a list written with the code in that window survives it; and the
// market-scoped `default_kind` SETTING is neither consulted nor cleared, so
// deleting the kind it names leaves the setting pointing at nothing while
// creates fall through to whichever kind holds the default flag. A list that
// does end up naming a code nothing defines is not broken, only stranded: it
// is still returned by GET /orderlists and GET /orderlists/{id} carrying the
// bare code, the vocabulary no longer offers that value so a UI renders the
// code itself, `?kind=` refuses it with a 400 naming the codes that remain,
// and the way back is PUT /orderlists/{id} with a kind the tenant keeps.
// Deleting the flag-holder hands the flag to the first remaining kind. The
// answer is the `code`, not the `{deleted, id}` the other deletes here
// return.
func (srv *Orderlists) OrderlistsKindsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orderlists/kinds/{id}")
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
	
// OrderlistsKindsGet one kind, by the id this route takes. The `code` is the
// OTHER identity and the one that matters to the data: `lists.kind` stores
// the code and never this id, so a list is joined to its kind by code while
// every /orderlists/kinds/{id} route is addressed by uuid. A fresh tenant
// starts with two — `shopping` and `label`, seeded on install — and
// everything beyond them is the merchant's own. A kind seeded before 0.15.0
// may hold a serialized locale map in `title` and `description` where plain
// text belongs; those rows were left as they stand, because repairing them is
// a data change.
func (srv *Orderlists) OrderlistsKindsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orderlists/kinds/{id}")
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
type OrderlistsKindsUpdateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Title string
	Tone string
	enabledSetters map[string]bool
}
func (options OrderlistsKindsUpdateOptions) New() *OrderlistsKindsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"Descriptions": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Title": false,
		"Tone": false,
	}
	return &options
}
type OrderlistsKindsUpdateOption func(*OrderlistsKindsUpdateOptions)
func (srv *Orderlists) WithOrderlistsKindsUpdateDescription(v string) OrderlistsKindsUpdateOption {
	return func(o *OrderlistsKindsUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsUpdateDescriptions(v interface{}) OrderlistsKindsUpdateOption {
	return func(o *OrderlistsKindsUpdateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsUpdateIsDefault(v bool) OrderlistsKindsUpdateOption {
	return func(o *OrderlistsKindsUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsUpdateLabels(v interface{}) OrderlistsKindsUpdateOption {
	return func(o *OrderlistsKindsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsUpdatePosition(v int) OrderlistsKindsUpdateOption {
	return func(o *OrderlistsKindsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsUpdateTitle(v string) OrderlistsKindsUpdateOption {
	return func(o *OrderlistsKindsUpdateOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *Orderlists) WithOrderlistsKindsUpdateTone(v string) OrderlistsKindsUpdateOption {
	return func(o *OrderlistsKindsUpdateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
			
// OrderlistsKindsUpdate everything a kind has except its code: the title a
// person reads, the sentence underneath it, the localized forms of both, the
// badge tone, and where it sits in a select. The code is not among them and
// cannot be reached from here at all: sending a different one is a 400 rather
// than a silent no-op, because `lists.kind` stores the code and a rename
// would orphan every list that carries it with no foreign key to stop it. So
// a rename is never how a list comes to name a code nothing defines — only
// a delete can do that. Renaming the TITLE touches no list, for the same
// reason. A blank title is ignored rather than stored; an explicit null
// clears the description; `labels` and `descriptions` replace the whole map
// rather than merging into it. `is_default: true` makes the same move POST
// /orderlists/kinds/{id}/make-default makes on its own. A system kind is
// editable like any other.
func (srv *Orderlists) OrderlistsKindsUpdate(Id string, optionalSetters ...OrderlistsKindsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orderlists/kinds/{id}")
	options := OrderlistsKindsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["Descriptions"] {
		params["descriptions"] = options.Descriptions
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["Tone"] {
		params["tone"] = options.Tone
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
			
// OrderlistsKindsMakeDefault one call MOVES the flag: the kind in the path is
// promoted and whoever held the flag before is demoted in the same request,
// because the flag is a single answer and not a per-row opinion. It is what a
// list created without a kind falls back to, so two defaults leave the result
// to row order and none leaves it to whatever sorts first — which is
// exactly why promotion and demotion cannot be two calls a client makes in
// sequence. PUT with is_default already moved it, but only as a side effect
// of an edit, and a client promoting and then demoting by hand produces those
// two broken states whenever one of the pair does not land. Every kind the
// tenant keeps is walked, and only the rows whose flag is wrong are written
// — the new default if it was not already set, the old one if it was — so
// the call costs at most two writes and repeating it costs none, which makes
// it safe to retry. The kind's other fields are untouched and no existing
// list is rewritten: lists that already name a kind keep it, since the flag
// decides only what a FUTURE create with no `kind` resolves to. The
// market-scoped `default_kind` setting still wins where it is set; this flag
// is the tenant-wide answer underneath it.
func (srv *Orderlists) OrderlistsKindsMakeDefault(Id string, Data interface{})(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orderlists/kinds/{id}/make-default")
	params := map[string]interface{}{}
	params["id"] = Id
	params["data"] = Data
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

// OrderlistsVocabulariesList discovery for the vocabulary routes, and nothing
// more: every enum this app publishes, each as a name plus the words a person
// reads for it — its title and its description — and never the values,
// which are one call further down at GET /orderlists/vocabularies/{name}. It
// exists so that a client holding a qualified pair like 'orderlists.kinds'
// can build that URL from the pair alone and keep no copy of an enum of its
// own. Names: kinds. The split is deliberate rather than an economy: the set
// of NAMES is fixed by a release of this app, so a client may cache this
// answer for as long as it caches the contract, while the values under
// 'kinds' are the tenant's own rows and change without a release — which is
// why this route says nothing about them and why a UI building a select must
// make the second call rather than read the values off here. Title and
// description come back either as a plain string or as a locale map keyed by
// language tag, so a client reads the tag it wants and falls back to `en` —
// the same shape every localized field in this app carries.
func (srv *Orderlists) OrderlistsVocabulariesList()(*models.OrderListVocabularyIndex, error) {
	path := "/v1/orderlists/vocabularies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.OrderListVocabularyIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.OrderListVocabularyIndex
	parsed, ok := resp.Result.(models.OrderListVocabularyIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// OrderlistsVocabulariesGet one named enum with every value it permits, and
// enough about each value to render it without a second source: the `key` the
// database stores and enforces, the title and the description a person reads,
// and the semantic badge `tone` a UI colours it with — which is why no
// client needs a colour map of its own, and why the Cockpit's hand-kept one
// could go. A value that names no tone of its own inherits the vocabulary's
// `default_tone`, so the field is never empty. 'kinds' is table-backed: the
// tenant's own rows ARE the value set, so a value they added appears here
// without a release of this app, and each value carries its `labels`,
// `descriptions` and the `is_default` flag besides. Values come back in
// `position` order, which is the order a select should offer. 'closed' says
// the set is exhaustive at this moment, so a value outside it is stale data
// rather than a missing label — what changed with the move to a table is
// WHO may extend it, not whether the set is closed. `source` says which:
// 'schema' where a CHECK constraint owns the values, 'table' where the
// tenant's rows do. Names: kinds.
func (srv *Orderlists) OrderlistsVocabulariesGet(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/orderlists/vocabularies/{name}")
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
	
// OrderlistsDelete takes every position with it, in the database:
// `items.list_id` is the app's only foreign key and it is ON DELETE CASCADE,
// and the handler removes the positions explicitly first besides. Nothing
// survives the list, there is no soft delete and no undo — and the answer
// carries no count, so read the list (or its `item_count`) BEFORE the call if
// you need to know how much went. What it does NOT take is what the list has
// already produced: a cart line or an order position built by the conversions
// carries `order_list_id`, `order_list_name` and `order_list_item_id` in its
// snapshot, and those are jsonb values inside another app rather than foreign
// keys — ADR-0055 forbids a cross-app FK, so nothing cascades there and
// nothing is nulled. The cart and the order are unharmed, because every
// position was copied as a snapshot rather than referenced; the provenance
// link is what dangles, permanently.
func (srv *Orderlists) OrderlistsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orderlists/{id}")
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
	
// OrderlistsGet the whole list in one call: the row plus every position
// inline, in `position` order, up to a thousand of them. The nested positions
// collection exists to CHANGE the positions, not to page them, so this is the
// read a detail view makes. Reading is wider than writing here — an acting
// contact sees their own lists and their organization's shared ones, and a
// list that is neither answers 404 rather than 403, so an outsider learns
// nothing from the difference. The row carries the dead `public` column next
// to `shared`; read `shared`.
func (srv *Orderlists) OrderlistsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orderlists/{id}")
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
type OrderlistsUpdateOptions struct {
	Kind string
	Metadata interface{}
	Name string
	Shared bool
	enabledSetters map[string]bool
}
func (options OrderlistsUpdateOptions) New() *OrderlistsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Kind": false,
		"Metadata": false,
		"Name": false,
		"Shared": false,
	}
	return &options
}
type OrderlistsUpdateOption func(*OrderlistsUpdateOptions)
func (srv *Orderlists) WithOrderlistsUpdateKind(v string) OrderlistsUpdateOption {
	return func(o *OrderlistsUpdateOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *Orderlists) WithOrderlistsUpdateMetadata(v interface{}) OrderlistsUpdateOption {
	return func(o *OrderlistsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Orderlists) WithOrderlistsUpdateName(v string) OrderlistsUpdateOption {
	return func(o *OrderlistsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Orderlists) WithOrderlistsUpdateShared(v bool) OrderlistsUpdateOption {
	return func(o *OrderlistsUpdateOptions) {
		o.Shared = v
		o.enabledSetters["Shared"] = true
	}
}
			
// OrderlistsUpdate rename, share or reclassify — the whole of what a list
// says about itself, plus `metadata`. Positions go through the items routes
// and the owner cannot be changed by anything. `shared` is what the column
// `public` was renamed to in June 2026; `public` is still on the wire because
// the provisioner is additive, is false on every row written since, and says
// nothing about who may see the list. One trap: a `kind` this tenant does not
// keep is IGNORED rather than refused, so the list quietly keeps the kind it
// had and a client that cares must read the answer back. An empty body is a
// 400 rather than a no-op.
func (srv *Orderlists) OrderlistsUpdate(Id string, optionalSetters ...OrderlistsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orderlists/{id}")
	options := OrderlistsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Shared"] {
		params["shared"] = options.Shared
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
type OrderlistsToCartOptions struct {
	CartId string
	Currency string
	Mode string
	enabledSetters map[string]bool
}
func (options OrderlistsToCartOptions) New() *OrderlistsToCartOptions {
	options.enabledSetters = map[string]bool{
		"CartId": false,
		"Currency": false,
		"Mode": false,
	}
	return &options
}
type OrderlistsToCartOption func(*OrderlistsToCartOptions)
func (srv *Orderlists) WithOrderlistsToCartCartId(v string) OrderlistsToCartOption {
	return func(o *OrderlistsToCartOptions) {
		o.CartId = v
		o.enabledSetters["CartId"] = true
	}
}
func (srv *Orderlists) WithOrderlistsToCartCurrency(v string) OrderlistsToCartOption {
	return func(o *OrderlistsToCartOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Orderlists) WithOrderlistsToCartMode(v string) OrderlistsToCartOption {
	return func(o *OrderlistsToCartOptions) {
		o.Mode = v
		o.enabledSetters["Mode"] = true
	}
}
			
// OrderlistsToCart the reason a buyer keeps a list at all: every position of
// the list goes into a cart in one call. The cart is either one the caller
// names or one this call makes. Sending 'cart_id' adds to that existing cart;
// omitting it creates a cart for the LIST'S OWNER — not for whoever called
// — names it after the list, and makes it that owner's current cart,
// because a cart the buyer cannot see is not 'added to cart'. Which of the
// two happened is not left to be inferred: `cart_created` says so and
// `cart_id` names the cart either way. 'append' (the default,
// tenant-configurable through `cart_merge_mode`) lets the carts app merge
// each line by product and price so quantities accumulate, and is sent one
// line at a time precisely because that merge happens on add; 'replace' makes
// the list the cart's whole contents in one call. What the cart has no column
// for — cost centre, custom SKU, position texts — rides in each line's
// snapshot together with the list it came from. The list itself is never
// touched: it is read, not emptied, so the same list converts again next
// month. Cross-app: carts.create, carts.items.create, carts.items.replace.
func (srv *Orderlists) OrderlistsToCart(Id string, optionalSetters ...OrderlistsToCartOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orderlists/{id}/cart")
	options := OrderlistsToCartOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["CartId"] {
		params["cart_id"] = options.CartId
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["Mode"] {
		params["mode"] = options.Mode
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
type OrderlistsToOrderOptions struct {
	Currency string
	CustomerOrderNumber string
	enabledSetters map[string]bool
}
func (options OrderlistsToOrderOptions) New() *OrderlistsToOrderOptions {
	options.enabledSetters = map[string]bool{
		"Currency": false,
		"CustomerOrderNumber": false,
	}
	return &options
}
type OrderlistsToOrderOption func(*OrderlistsToOrderOptions)
func (srv *Orderlists) WithOrderlistsToOrderCurrency(v string) OrderlistsToOrderOption {
	return func(o *OrderlistsToOrderOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Orderlists) WithOrderlistsToOrderCustomerOrderNumber(v string) OrderlistsToOrderOption {
	return func(o *OrderlistsToOrderOptions) {
		o.CustomerOrderNumber = v
		o.enabledSetters["CustomerOrderNumber"] = true
	}
}
			
// OrderlistsToOrder the other half of the reason a list exists — and it is
// the ORDERS app that does it, over the gateway rather than over a shared
// table, so everything an order means is that app's answer and not this
// one's. Places the list's positions as an order: buyer and organization come
// from the list, the cost centre and the position texts land on the order's
// own columns, and the list is left exactly as it stands so it can be ordered
// again next month. The acting contact is re-asserted on the call, so the
// orders app applies ITS rules to the BUYER rather than to this app — a
// contact holding only orders.request, or an order above the tenant's
// approval threshold, comes back with status 'pending' and no placed_at
// instead of being refused. That pending order is the platform's nearest
// thing to a draft; the orders app owns the state and this one cannot
// override it, which is why `status` is reported rather than chosen and why
// the created order is handed back verbatim under `order` beside the three
// fields lifted out of it. Cross-app: orders.place.
func (srv *Orderlists) OrderlistsToOrder(Id string, optionalSetters ...OrderlistsToOrderOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/orderlists/{id}/order")
	options := OrderlistsToOrderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["CustomerOrderNumber"] {
		params["customer_order_number"] = options.CustomerOrderNumber
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
type OrderlistsItemsListOptions struct {
	Id string
	ProductId string
	Sku string
	Name string
	Image string
	Quantity float64
	Unit string
	Price float64
	TaxRate float64
	CostCenterId string
	PositionTexts string
	CustomSku string
	CategorySlug string
	SubcategorySlug string
	Position int
	Metadata string
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options OrderlistsItemsListOptions) New() *OrderlistsItemsListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"ProductId": false,
		"Sku": false,
		"Name": false,
		"Image": false,
		"Quantity": false,
		"Unit": false,
		"Price": false,
		"TaxRate": false,
		"CostCenterId": false,
		"PositionTexts": false,
		"CustomSku": false,
		"CategorySlug": false,
		"SubcategorySlug": false,
		"Position": false,
		"Metadata": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type OrderlistsItemsListOption func(*OrderlistsItemsListOptions)
func (srv *Orderlists) WithOrderlistsItemsListId(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListProductId(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListSku(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListName(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListImage(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListQuantity(v float64) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListUnit(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListPrice(v float64) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Price = v
		o.enabledSetters["Price"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListTaxRate(v float64) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.TaxRate = v
		o.enabledSetters["TaxRate"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListCostCenterId(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.CostCenterId = v
		o.enabledSetters["CostCenterId"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListPositionTexts(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.PositionTexts = v
		o.enabledSetters["PositionTexts"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListCustomSku(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.CustomSku = v
		o.enabledSetters["CustomSku"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListCategorySlug(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.CategorySlug = v
		o.enabledSetters["CategorySlug"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListSubcategorySlug(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.SubcategorySlug = v
		o.enabledSetters["SubcategorySlug"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListPosition(v int) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListMetadata(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListCreatedAt(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListUpdatedAt(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListLimit(v int) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListOffset(v int) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsListOrder(v string) OrderlistsItemsListOption {
	return func(o *OrderlistsItemsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
			
// OrderlistsItemsList every column of a position is an exact-match filter —
// eighteen of them, which is the whole row — and they combine as AND.
// `list_id` is not among them: it comes from the path and overwrites anything
// the query says. The default sort is `position.asc`, and `position` is
// neither dense nor unique: removing a position leaves its number behind
// while the next add takes the list's current COUNT, so a delete from the
// middle followed by an add produces two rows sharing a number and the tie
// falls to whatever the database returns first. Sort by `created_at` where
// the order has to be unambiguous.
func (srv *Orderlists) OrderlistsItemsList(ListId string, optionalSetters ...OrderlistsItemsListOption)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId)
	path := r.Replace("/v1/orderlists/{list_id}/items")
	options := OrderlistsItemsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["list_id"] = ListId
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Image"] {
		params["image"] = options.Image
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Unit"] {
		params["unit"] = options.Unit
	}
	if options.enabledSetters["Price"] {
		params["price"] = options.Price
	}
	if options.enabledSetters["TaxRate"] {
		params["tax_rate"] = options.TaxRate
	}
	if options.enabledSetters["CostCenterId"] {
		params["cost_center_id"] = options.CostCenterId
	}
	if options.enabledSetters["PositionTexts"] {
		params["position_texts"] = options.PositionTexts
	}
	if options.enabledSetters["CustomSku"] {
		params["custom_sku"] = options.CustomSku
	}
	if options.enabledSetters["CategorySlug"] {
		params["category_slug"] = options.CategorySlug
	}
	if options.enabledSetters["SubcategorySlug"] {
		params["subcategory_slug"] = options.SubcategorySlug
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["CreatedAt"] {
		params["created_at"] = options.CreatedAt
	}
	if options.enabledSetters["UpdatedAt"] {
		params["updated_at"] = options.UpdatedAt
	}
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
type OrderlistsItemsCreateOptions struct {
	CategorySlug string
	CostCenterId string
	CustomSku string
	Image string
	Metadata interface{}
	Position int
	PositionTexts []string
	Price float64
	ProductId string
	Quantity float64
	Sku string
	SubcategorySlug string
	TaxRate float64
	Unit string
	enabledSetters map[string]bool
}
func (options OrderlistsItemsCreateOptions) New() *OrderlistsItemsCreateOptions {
	options.enabledSetters = map[string]bool{
		"CategorySlug": false,
		"CostCenterId": false,
		"CustomSku": false,
		"Image": false,
		"Metadata": false,
		"Position": false,
		"PositionTexts": false,
		"Price": false,
		"ProductId": false,
		"Quantity": false,
		"Sku": false,
		"SubcategorySlug": false,
		"TaxRate": false,
		"Unit": false,
	}
	return &options
}
type OrderlistsItemsCreateOption func(*OrderlistsItemsCreateOptions)
func (srv *Orderlists) WithOrderlistsItemsCreateCategorySlug(v string) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.CategorySlug = v
		o.enabledSetters["CategorySlug"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreateCostCenterId(v string) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.CostCenterId = v
		o.enabledSetters["CostCenterId"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreateCustomSku(v string) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.CustomSku = v
		o.enabledSetters["CustomSku"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreateImage(v string) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreateMetadata(v interface{}) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreatePosition(v int) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreatePositionTexts(v []string) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.PositionTexts = v
		o.enabledSetters["PositionTexts"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreatePrice(v float64) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.Price = v
		o.enabledSetters["Price"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreateProductId(v string) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreateQuantity(v float64) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreateSku(v string) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreateSubcategorySlug(v string) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.SubcategorySlug = v
		o.enabledSetters["SubcategorySlug"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreateTaxRate(v float64) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.TaxRate = v
		o.enabledSetters["TaxRate"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsCreateUnit(v string) OrderlistsItemsCreateOption {
	return func(o *OrderlistsItemsCreateOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
					
// OrderlistsItemsCreate a position is a whole saved line, not a pointer at a
// product. `name` is required and one of `product_id` / `sku` must be set —
// the two things the database itself insists on — and everything else is a
// snapshot of what the buyer saw. Nothing here deduplicates: adding the same
// article twice makes two positions, because it is the CART that merges lines
// by product and price, not the list. The new row takes the list's current
// position COUNT unless the payload names a `position` of its own, so it
// collides with an existing number whenever an earlier position was deleted
// from the middle. The list's `updated_at` is touched, which is what the
// default sort of GET /orderlists reads.
func (srv *Orderlists) OrderlistsItemsCreate(ListId string, Name string, optionalSetters ...OrderlistsItemsCreateOption)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId)
	path := r.Replace("/v1/orderlists/{list_id}/items")
	options := OrderlistsItemsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["list_id"] = ListId
	params["name"] = Name
	if options.enabledSetters["CategorySlug"] {
		params["category_slug"] = options.CategorySlug
	}
	if options.enabledSetters["CostCenterId"] {
		params["cost_center_id"] = options.CostCenterId
	}
	if options.enabledSetters["CustomSku"] {
		params["custom_sku"] = options.CustomSku
	}
	if options.enabledSetters["Image"] {
		params["image"] = options.Image
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["PositionTexts"] {
		params["position_texts"] = options.PositionTexts
	}
	if options.enabledSetters["Price"] {
		params["price"] = options.Price
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["SubcategorySlug"] {
		params["subcategory_slug"] = options.SubcategorySlug
	}
	if options.enabledSetters["TaxRate"] {
		params["tax_rate"] = options.TaxRate
	}
	if options.enabledSetters["Unit"] {
		params["unit"] = options.Unit
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
			
// OrderlistsItemsReplace set semantics: what you send becomes the list's
// positions and everything else is deleted. Ids are NOT preserved — every
// row is dropped and rewritten, so a client holding position ids must re-read
// them — and an empty array empties the list. Both guards run before the
// first delete, so an oversized or unknown-article replace answers 400 with
// the list still holding exactly what it held. It is not a renumbering call:
// an entry that names no `position` takes its array index, one that names its
// own keeps it, so the array order is the default rather than an override.
// Writing is narrower than reading: the owner may always replace, and anyone
// else only when the list is shared with their own organization AND the
// tenant turned `shared_lists_editable` on — otherwise a caller who can
// READ the list through the sharing rule is answered 403 here. The
// delete-then-insert is not wrapped in a transaction of its own, so a client
// should treat a failed replace as a list of unknown contents and re-read it
// rather than retry blind. The answer is the whole new set in the same paged
// envelope every other collection uses, with `limit`, `offset` and `total`
// describing exactly what was written; the list's `updated_at` is touched,
// which moves it to the front of the default GET /orderlists page.
func (srv *Orderlists) OrderlistsItemsReplace(ListId string, Items []models.OrderListItemInput)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId)
	path := r.Replace("/v1/orderlists/{list_id}/items")
	params := map[string]interface{}{}
	params["list_id"] = ListId
	params["items"] = Items
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
			
// OrderlistsItemsDelete removes one saved line and takes nothing with it —
// no foreign key in this app points at a position. What it leaves behind is
// the gap: every remaining row keeps the number it had, and the next add
// takes the list's COUNT as its `position`, so a removal from the middle sets
// up a later collision. A bulk replace is the only call that rewrites the
// sequence. Outside this app, a cart line or order position built from this
// row still carries `order_list_item_id` in its snapshot — a jsonb value,
// not a reference — so it is simply left naming a row that is gone. The
// list's `updated_at` is touched.
func (srv *Orderlists) OrderlistsItemsDelete(ListId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId, "{id}", Id)
	path := r.Replace("/v1/orderlists/{list_id}/items/{id}")
	params := map[string]interface{}{}
	params["list_id"] = ListId
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
			
// OrderlistsItemsGet one saved line by its own id, in exactly the shape the
// collection returns — there is nothing here the collection does not
// already give you, so this is the read for a client that holds a position id
// and nothing else. The list in the path is enforced rather than decorative:
// a position that belongs to a different list answers 404 rather than the
// row, which is what stops an id lifting a position out of a list the caller
// may not read. An unknown or unreadable list is a 404 before the position is
// looked at.
func (srv *Orderlists) OrderlistsItemsGet(ListId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId, "{id}", Id)
	path := r.Replace("/v1/orderlists/{list_id}/items/{id}")
	params := map[string]interface{}{}
	params["list_id"] = ListId
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
type OrderlistsItemsUpdateOptions struct {
	CategorySlug string
	CostCenterId string
	CustomSku string
	Image string
	Metadata interface{}
	Name string
	Position int
	PositionTexts []string
	Price float64
	ProductId string
	Quantity float64
	Sku string
	SubcategorySlug string
	TaxRate float64
	Unit string
	enabledSetters map[string]bool
}
func (options OrderlistsItemsUpdateOptions) New() *OrderlistsItemsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"CategorySlug": false,
		"CostCenterId": false,
		"CustomSku": false,
		"Image": false,
		"Metadata": false,
		"Name": false,
		"Position": false,
		"PositionTexts": false,
		"Price": false,
		"ProductId": false,
		"Quantity": false,
		"Sku": false,
		"SubcategorySlug": false,
		"TaxRate": false,
		"Unit": false,
	}
	return &options
}
type OrderlistsItemsUpdateOption func(*OrderlistsItemsUpdateOptions)
func (srv *Orderlists) WithOrderlistsItemsUpdateCategorySlug(v string) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.CategorySlug = v
		o.enabledSetters["CategorySlug"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdateCostCenterId(v string) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.CostCenterId = v
		o.enabledSetters["CostCenterId"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdateCustomSku(v string) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.CustomSku = v
		o.enabledSetters["CustomSku"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdateImage(v string) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdateMetadata(v interface{}) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdateName(v string) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdatePosition(v int) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdatePositionTexts(v []string) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.PositionTexts = v
		o.enabledSetters["PositionTexts"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdatePrice(v float64) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.Price = v
		o.enabledSetters["Price"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdateProductId(v string) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdateQuantity(v float64) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.Quantity = v
		o.enabledSetters["Quantity"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdateSku(v string) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.Sku = v
		o.enabledSetters["Sku"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdateSubcategorySlug(v string) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.SubcategorySlug = v
		o.enabledSetters["SubcategorySlug"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdateTaxRate(v float64) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.TaxRate = v
		o.enabledSetters["TaxRate"] = true
	}
}
func (srv *Orderlists) WithOrderlistsItemsUpdateUnit(v string) OrderlistsItemsUpdateOption {
	return func(o *OrderlistsItemsUpdateOptions) {
		o.Unit = v
		o.enabledSetters["Unit"] = true
	}
}
					
// OrderlistsItemsUpdate a partial update: omitted fields keep the value they
// have, and an explicit null is the only way to clear one. `quantity` is
// re-checked (> 0), and where `reject_unknown_articles` is on the article is
// re-checked against the MERGED row rather than the payload — so changing
// only the name cannot smuggle an unknown article past the guard that the
// create applied. `position` is set, not shifted: writing 3 puts this row at
// 3 and moves nothing else, which is the other way two positions come to
// share a number. The list's `updated_at` is touched.
func (srv *Orderlists) OrderlistsItemsUpdate(ListId string, Id string, optionalSetters ...OrderlistsItemsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{list_id}", ListId, "{id}", Id)
	path := r.Replace("/v1/orderlists/{list_id}/items/{id}")
	options := OrderlistsItemsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["list_id"] = ListId
	params["id"] = Id
	if options.enabledSetters["CategorySlug"] {
		params["category_slug"] = options.CategorySlug
	}
	if options.enabledSetters["CostCenterId"] {
		params["cost_center_id"] = options.CostCenterId
	}
	if options.enabledSetters["CustomSku"] {
		params["custom_sku"] = options.CustomSku
	}
	if options.enabledSetters["Image"] {
		params["image"] = options.Image
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["PositionTexts"] {
		params["position_texts"] = options.PositionTexts
	}
	if options.enabledSetters["Price"] {
		params["price"] = options.Price
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["Quantity"] {
		params["quantity"] = options.Quantity
	}
	if options.enabledSetters["Sku"] {
		params["sku"] = options.Sku
	}
	if options.enabledSetters["SubcategorySlug"] {
		params["subcategory_slug"] = options.SubcategorySlug
	}
	if options.enabledSetters["TaxRate"] {
		params["tax_rate"] = options.TaxRate
	}
	if options.enabledSetters["Unit"] {
		params["unit"] = options.Unit
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
