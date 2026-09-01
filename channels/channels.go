package channels

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Channels service
type Channels struct {
	client client.Client
}

func New(clt client.Client) *Channels {
	return &Channels{
		client: clt,
	}
}

type ChannelsListOptions struct {
	Id string
	Code string
	Name string
	Labels string
	Type string
	Status string
	UnassignedVisibility string
	IsDefault bool
	Position int
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options ChannelsListOptions) New() *ChannelsListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Code": false,
		"Name": false,
		"Labels": false,
		"Type": false,
		"Status": false,
		"UnassignedVisibility": false,
		"IsDefault": false,
		"Position": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type ChannelsListOption func(*ChannelsListOptions)
func (srv *Channels) WithChannelsListId(v string) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Channels) WithChannelsListCode(v string) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Channels) WithChannelsListName(v string) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Channels) WithChannelsListLabels(v string) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Channels) WithChannelsListType(v string) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Channels) WithChannelsListStatus(v string) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Channels) WithChannelsListUnassignedVisibility(v string) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.UnassignedVisibility = v
		o.enabledSetters["UnassignedVisibility"] = true
	}
}
func (srv *Channels) WithChannelsListIsDefault(v bool) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Channels) WithChannelsListPosition(v int) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Channels) WithChannelsListCreatedAt(v string) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Channels) WithChannelsListUpdatedAt(v string) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *Channels) WithChannelsListLimit(v int) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Channels) WithChannelsListOffset(v int) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Channels) WithChannelsListOrder(v string) ChannelsListOption {
	return func(o *ChannelsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// ChannelsList the filters are what make this list usable: `?code=` turns a
// scope slug another app stored into the channel row that owns it,
// `?is_default=true` finds the fallback channel without resolving a context,
// and `?unassigned_visibility=assigned_only` finds the channels that closed
// their assortment. Every filter is an exact-column equality — there is no
// contains, prefix or range form — and the honoured set is exactly this
// entity's 11 columns, because the generic list mount matches any query key
// that names one. Each of them is declared as a query parameter with the
// column's own CHECK behind it, so the 11 that work are the 11 the document
// offers rather than a list a caller has to keep somewhere. An unknown column
// is dropped rather than refused, so `?stauts=active` returns the unfiltered
// page; `filter` echoes what was understood, which is the only way to catch
// that. Paging is `limit`/`offset` over whatever survived the filters, and
// `?order=` sorts by one column with an optional `.asc`/`.desc`; ask for no
// order and the page comes back in insertion order. `order` is the one input
// here that is refused rather than ignored — a malformed value, or one
// naming a column this entity does not have, is a 400 where the same mistake
// in a filter key passes silently.
func (srv *Channels) ChannelsList(optionalSetters ...ChannelsListOption)(*models.Error, error) {
	path := "/v1/channels"
	options := ChannelsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["UnassignedVisibility"] {
		params["unassigned_visibility"] = options.UnassignedVisibility
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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
type ChannelsCreateOptions struct {
	IsDefault bool
	Labels interface{}
	Position int
	Status string
	Type string
	UnassignedVisibility string
	enabledSetters map[string]bool
}
func (options ChannelsCreateOptions) New() *ChannelsCreateOptions {
	options.enabledSetters = map[string]bool{
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Status": false,
		"Type": false,
		"UnassignedVisibility": false,
	}
	return &options
}
type ChannelsCreateOption func(*ChannelsCreateOptions)
func (srv *Channels) WithChannelsCreateIsDefault(v bool) ChannelsCreateOption {
	return func(o *ChannelsCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Channels) WithChannelsCreateLabels(v interface{}) ChannelsCreateOption {
	return func(o *ChannelsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Channels) WithChannelsCreatePosition(v int) ChannelsCreateOption {
	return func(o *ChannelsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Channels) WithChannelsCreateStatus(v string) ChannelsCreateOption {
	return func(o *ChannelsCreateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Channels) WithChannelsCreateType(v string) ChannelsCreateOption {
	return func(o *ChannelsCreateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Channels) WithChannelsCreateUnassignedVisibility(v string) ChannelsCreateOption {
	return func(o *ChannelsCreateOptions) {
		o.UnassignedVisibility = v
		o.enabledSetters["UnassignedVisibility"] = true
	}
}
					
// ChannelsCreate two fields are yours and everything else has an answer
// already: `code` and `name` are the only columns the database will not fill
// in, and the rest arrive from their defaults — `status` active,
// `unassigned_visibility` inherit, `is_default` false, `position` 0. `type`
// is the exception the app makes for itself: omitted, it becomes the type the
// tenant FLAGGED as their default rather than the column default, so a
// merchant who retired the seeded `storefront` does not get channels carrying
// a type they no longer keep. `code` is the load-bearing one. It is the scope
// slug Baseline matches every channel assignment on, which is why it is held
// to Baseline's own shape here rather than to the column's `length > 0`, and
// why it is unique per tenant — a second channel claiming a code another
// already holds is a 409 off the `(tenant_id, code)` index. Treat it as
// permanent: the API will let you change it later and nothing follows it (see
// PUT /channels/{id}). Creating a channel assigns nothing to it. Products,
// categories and everything else scopeable stay exactly as visible as they
// were — until rows are assigned, what this channel shows is whatever
// `unassigned_channel_visibility` says, which on the shipped default is the
// entire catalogue. And a code is only free in THIS app: assignments made
// against a code that a since-deleted channel used are still in Baseline, so
// re-using the code adopts them.
func (srv *Channels) ChannelsCreate(Code string, Name string, optionalSetters ...ChannelsCreateOption)(*models.Error, error) {
	path := "/v1/channels"
	options := ChannelsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["name"] = Name
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["UnassignedVisibility"] {
		params["unassigned_visibility"] = options.UnassignedVisibility
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
type ChannelsContextOptions struct {
	Channel string
	enabledSetters map[string]bool
}
func (options ChannelsContextOptions) New() *ChannelsContextOptions {
	options.enabledSetters = map[string]bool{
		"Channel": false,
	}
	return &options
}
type ChannelsContextOption func(*ChannelsContextOptions)
func (srv *Channels) WithChannelsContextChannel(v string) ChannelsContextOption {
	return func(o *ChannelsContextOptions) {
		o.Channel = v
		o.enabledSetters["Channel"] = true
	}
}
	
// ChannelsContext the storefront/punchout bootstrap: one call tells a shop
// front, a punchout front-end or a feed builder which channel it is in and
// what an unassigned row means there, so it can apply the policy itself
// instead of hardcoding one. Resolution order is body/query, then the
// x-revenexx-channel header, then the scope_context.channel claim, then the
// channel flagged is_default — header before claim, the same order
// baseline.is_visible() uses. Through api.revenexx.com the header step is
// inert (the gateway does not forward it), so in practice it is `?channel=`,
// then the claim, then the default. Never errors on an unknown or inactive
// channel: it answers resolved:false with a reason, so a caller can tell "no
// such channel" from "the service is down". That is why this operation
// declares no 4xx of its own — a tenant with no channels at all answers 200
// with reason no_default_channel. Two things come back, not one: the channel
// that was resolved, and the visibility policy in force for it — the
// tenant-wide unassigned_channel_visibility answer, or the channel's own
// override where it has one. The policy travels with the channel because a
// caller that has one and not the other still cannot render anything: knowing
// you are in the punchout channel says nothing about whether an unassigned
// product belongs in its catalogue. With both, a client reproduces the
// decision itself and calls POST /channels/visibility only when it wants the
// app to decide row by row.
func (srv *Channels) ChannelsContext(optionalSetters ...ChannelsContextOption)(*models.ChannelContext, error) {
	path := "/v1/channels/context"
	options := ChannelsContextOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Channel"] {
		params["channel"] = options.Channel
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ChannelContext{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ChannelContext
	parsed, ok := resp.Result.(models.ChannelContext)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// ChannelsDefaults the repair call. A tenant installed before `channel_types`
// existed, or one that deleted its way into a state where nothing works, gets
// the shipped starting point back — the 5 seeded types first, because the
// seeded channel carries one of them, then the `shop` channel. Most tenants
// never call it: the platform invokes the same routine itself on
// `app.installed`, so a fresh install already has its 5 types and its shop
// channel before anyone asks, and this route exists for the tenant whose
// install predates them or who has since removed them. Calling it on a
// settled tenant is safe and cheap for the same reason it is safe to fire on
// every install: it is idempotent, keyed on the code, so a second call writes
// nothing. Everything a merchant added themselves is left alone, and a row
// that already exists is reported under `existing` rather than rewritten —
// the values you edited on a seeded type survive this call. It RESTORES THE
// WHOLE SEED SET, including a seeded type the merchant deliberately deleted.
// Idempotency here is keyed on the code and nothing else, and there is
// nowhere to remember a retirement: retirement is not a state this app can
// represent. Retiring a type IS deleting the row; `channel_types` has no
// retired flag and these tables carry no foreign keys, so nothing anywhere
// distinguishes a code a merchant removed on purpose from one they never had.
// Honouring the retirement would mean inventing a tombstone rather than
// reading one. Given that, restoring all 5 is the better half of the trade:
// this is the call a tenant makes when something is missing, and a repair
// that silently skips part of what it repairs, with no way to ask for the
// rest, is worse than one that says plainly what it puts back. It is also
// never a surprise. The only automatic seeding elsewhere in the app fires
// when the type table is completely EMPTY, which cannot happen once installed
// because the last remaining type cannot be deleted — so a retired type
// comes back exactly when somebody calls this route or the app is installed
// again, and never as a side effect of an unrelated read. Deleting it a
// second time costs one DELETE, and is refused only if a channel has since
// started carrying it. What it does not do: it creates no assignments, it
// does not repair a channel whose own code you deleted (only `shop` comes
// back), and it does not restore the seeded VALUES of a type that still
// exists — a renamed `storefront` stays renamed.
func (srv *Channels) ChannelsDefaults()(*models.Error, error) {
	path := "/v1/channels/defaults"
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
type ChannelsTypesListOptions struct {
	Limit int
	Offset int
	enabledSetters map[string]bool
}
func (options ChannelsTypesListOptions) New() *ChannelsTypesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
	}
	return &options
}
type ChannelsTypesListOption func(*ChannelsTypesListOptions)
func (srv *Channels) WithChannelsTypesListLimit(v int) ChannelsTypesListOption {
	return func(o *ChannelsTypesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Channels) WithChannelsTypesListOffset(v int) ChannelsTypesListOption {
	return func(o *ChannelsTypesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
	
// ChannelsTypesList what a channel may BE. This used to be a CHECK constraint
// over five values, which meant the merchant who runs a feed channel or a
// print channel needed a release of this app to say so — and nothing in the
// app ever branched on the value, only on membership. The set is the tenant's
// rows now. Seeds itself on first read, so the list is never empty and a
// channel can always carry a type. Rows come back in `position` order,
// always: this route is not the generic list mount and takes no `order` —
// `limit` and `offset` are the whole of its query, and it takes no filters,
// so a caller looking for one code reads the list and matches. The set is
// bounded: a tenant keeps at most 200 types, which is the size this app can
// check a channel's type against in one query, and POST /channels/types
// refuses the 201st rather than build a set it could not read back.
// `page.total` counts the rows that exist, not the ones this answer carries,
// and the order is total — `position` then `code`, because `position` is
// not unique and an order that leaves rows tied is one the database is free
// to answer differently on the next page, which is how a walk serves a row
// twice and skips another.
func (srv *Channels) ChannelsTypesList(optionalSetters ...ChannelsTypesListOption)(*interface{}, error) {
	path := "/v1/channels/types"
	options := ChannelsTypesListOptions{}.New()
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
type ChannelsTypesCreateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Tone string
	enabledSetters map[string]bool
}
func (options ChannelsTypesCreateOptions) New() *ChannelsTypesCreateOptions {
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
type ChannelsTypesCreateOption func(*ChannelsTypesCreateOptions)
func (srv *Channels) WithChannelsTypesCreateDescription(v string) ChannelsTypesCreateOption {
	return func(o *ChannelsTypesCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Channels) WithChannelsTypesCreateDescriptions(v interface{}) ChannelsTypesCreateOption {
	return func(o *ChannelsTypesCreateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *Channels) WithChannelsTypesCreateIsDefault(v bool) ChannelsTypesCreateOption {
	return func(o *ChannelsTypesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Channels) WithChannelsTypesCreateLabels(v interface{}) ChannelsTypesCreateOption {
	return func(o *ChannelsTypesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Channels) WithChannelsTypesCreatePosition(v int) ChannelsTypesCreateOption {
	return func(o *ChannelsTypesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Channels) WithChannelsTypesCreateTone(v string) ChannelsTypesCreateOption {
	return func(o *ChannelsTypesCreateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
					
// ChannelsTypesCreate what lets a merchant name a kind of channel this app
// never thought of — a feed, a print catalogue, a kiosk — without waiting
// for a release. `code` and `title` are the only two the database will not
// fill in; everything else has a default. The code is trimmed and lowercased
// and becomes exactly what `channels.type` stores, and it is fixed from then
// on, because there is no foreign key behind that column to carry a rename:
// every channel holding the old string would be left pointing at nothing. The
// title is the part a merchant renames later. A duplicate code is a 409, and
// it is worth knowing that the collision is wider than this tenant —
// `channel_types.code` is unique on the column alone, so a code held by
// another tenant collides too and the read this route does before inserting
// cannot see it. A tenant keeps at most 200 types; the 201st is a 409
// `type_limit_reached` rather than a row the app would then be unable to read
// back. Creating a type changes nothing about existing channels: it is a name
// that becomes available, not one that gets applied. Adding a type does not
// make it the default either — pass `is_default: true` for that, which
// demotes the current holder.
func (srv *Channels) ChannelsTypesCreate(Code string, Title string, optionalSetters ...ChannelsTypesCreateOption)(*models.Error, error) {
	path := "/v1/channels/types"
	options := ChannelsTypesCreateOptions{}.New()
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
	
// ChannelsTypesDelete retiring a type IS deleting the row — there is no
// retired flag on `channel_types` — which is why the two things that would
// make a deletion destructive are refused instead of allowed: a type at least
// one channel still carries is a 409, and so is the last remaining type.
// There is no foreign key behind `channels.type`, so those two checks are not
// a convenience on top of the database, they ARE the integrity. Move the
// channels to another type first and the delete goes through. Nothing else
// goes with it. A type has no dependents once no channel names it: no rows in
// this app point at it and none in Baseline do either, since assignments are
// made against a channel `code`, never a type. Deleting the type the tenant
// had flagged as default is allowed, and the flag is handed to the next type
// by position rather than left unheld, so a channel created afterwards still
// has something to fall back to. Because the guard is a read followed by a
// write with no transaction between them, and no constraint underneath it, a
// channel created against this type in the same instant can survive it. Worth
// knowing what that leaves, since it is not what "orphaned" usually means:
// the channel keeps working. `channels.type` is a stored string that nothing
// joins on, so the channel still reads, still filters under `?type=` by that
// same string, and still resolves in /channels/context and POST
// /channels/visibility — neither of which consults `type` at all. What it
// loses is its label, because the types vocabulary is built from the rows and
// there is no longer one to render a badge from. An update that does not
// mention `type` leaves the value alone; naming it is refused, which is how
// the channel is moved to a type that exists. One thing the deletion frees is
// wider than the tenant: `channel_types.code` is unique on the column alone,
// so the code becomes available platform-wide, not just here. And the seed
// does not know the row is gone — POST /channels/defaults and a re-install
// both put a deleted SEEDED type back, by design; see that operation.
func (srv *Channels) ChannelsTypesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/channels/types/{id}")
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
	
// ChannelsTypesGet one type row, by its uuid — the handle PUT and DELETE
// take, and the reason to hold on to what the list gave you. It is NOT the
// `code`: the code is what `channels.type` stores, and this route will not
// look one up. Neither will the list, which takes no filters at all, so a
// caller holding only a code pages `GET /channels/types` and matches
// client-side. Since the whole set is bounded and small that is one call, not
// a search. Unlike the list, this route does not seed. The list is
// hand-written so that a tenant whose table is still empty is given the 5
// shipped types instead of being told they have none; this is the generic
// item route, so on that same tenant it answers 404 for every id — which is
// the correct answer, since there is genuinely no such row yet. Read the list
// first. Nothing here is cached: the type list changes when a merchant edits
// it and this route always reflects that. Rows seeded before 0.7.0 may hold a
// serialized locale map in `title` or `description` rather than plain text
// (PE-452); `labels` and `descriptions` are the columns that carry the
// per-locale copy now.
func (srv *Channels) ChannelsTypesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/channels/types/{id}")
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
type ChannelsTypesUpdateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Title string
	Tone string
	enabledSetters map[string]bool
}
func (options ChannelsTypesUpdateOptions) New() *ChannelsTypesUpdateOptions {
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
type ChannelsTypesUpdateOption func(*ChannelsTypesUpdateOptions)
func (srv *Channels) WithChannelsTypesUpdateDescription(v string) ChannelsTypesUpdateOption {
	return func(o *ChannelsTypesUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Channels) WithChannelsTypesUpdateDescriptions(v interface{}) ChannelsTypesUpdateOption {
	return func(o *ChannelsTypesUpdateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *Channels) WithChannelsTypesUpdateIsDefault(v bool) ChannelsTypesUpdateOption {
	return func(o *ChannelsTypesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Channels) WithChannelsTypesUpdateLabels(v interface{}) ChannelsTypesUpdateOption {
	return func(o *ChannelsTypesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Channels) WithChannelsTypesUpdatePosition(v int) ChannelsTypesUpdateOption {
	return func(o *ChannelsTypesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Channels) WithChannelsTypesUpdateTitle(v string) ChannelsTypesUpdateOption {
	return func(o *ChannelsTypesUpdateOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *Channels) WithChannelsTypesUpdateTone(v string) ChannelsTypesUpdateOption {
	return func(o *ChannelsTypesUpdateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
			
// ChannelsTypesUpdate everything but the code. This is where a merchant
// renames a seeded type into their own words, gives it its German, moves it
// in the list a person picks from, or hands it the default flag. Seeded types
// are as editable as ones the merchant added — `is_system` records where a
// row came from and grants it nothing. Sending a different `code` is a 400
// rather than a silent no-op: it is what `channels.type` stores, there is no
// foreign key behind that column to carry the change — the database has
// none at all on these tables — and a rename would therefore move nothing.
// Every channel holding the old string would keep holding it, still working
// but with no type row to draw its name from. This refusal is the whole of
// the protection; to move channels to a new code, create the type and update
// the channels, in that order. Two fields are quietly forgiving rather than
// strict — a blank `title` and a `tone` outside the palette are both
// ignored and the stored value kept, so a client that sends a half-filled
// form does not clear what is there. `is_default` is one-way: true promotes
// this type and demotes the previous holder, false does nothing at all,
// because some type has to be the one a channel created without one gets.
func (srv *Channels) ChannelsTypesUpdate(Id string, optionalSetters ...ChannelsTypesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/channels/types/{id}")
	options := ChannelsTypesUpdateOptions{}.New()
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
type ChannelsVisibilityOptions struct {
	Channel string
	ChannelBody string
	enabledSetters map[string]bool
}
func (options ChannelsVisibilityOptions) New() *ChannelsVisibilityOptions {
	options.enabledSetters = map[string]bool{
		"Channel": false,
		"ChannelBody": false,
	}
	return &options
}
type ChannelsVisibilityOption func(*ChannelsVisibilityOptions)
func (srv *Channels) WithChannelsVisibilityChannel(v string) ChannelsVisibilityOption {
	return func(o *ChannelsVisibilityOptions) {
		o.Channel = v
		o.enabledSetters["Channel"] = true
	}
}
func (srv *Channels) WithChannelsVisibilityChannelBody(v string) ChannelsVisibilityOption {
	return func(o *ChannelsVisibilityOptions) {
		o.ChannelBody = v
		o.enabledSetters["ChannelBody"] = true
	}
}
			
// ChannelsVisibility the gate. A row WITH channel assignments is decided
// exactly as baseline.is_visible() decides it — visible iff the active
// channel is among them. A row WITHOUT assignments is the case
// unassigned_channel_visibility owns: 'all' shows it (Baseline's
// open-by-default, unchanged) and 'assigned_only' hides it, which the
// generated _scoped view has no way to express. A channel may override the
// tenant answer for itself, so the shop can stay open while a punchout
// channel serves only its negotiated assortment.
func (srv *Channels) ChannelsVisibility(Items []models.ChannelVisibilityItem, optionalSetters ...ChannelsVisibilityOption)(*models.Error, error) {
	path := "/v1/channels/visibility"
	options := ChannelsVisibilityOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["items"] = Items
	if options.enabledSetters["Channel"] {
		params["channel"] = options.Channel
	}
	if options.enabledSetters["ChannelBody"] {
		params["channel"] = options.ChannelBody
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

// ChannelsVocabulariesList discovery for the vocabulary routes: which enums
// this app publishes, not what is in them. An entry carries the name and the
// localised title and description a UI would put above a select, and stops
// there — the permitted values, their labels and their badge tones are the
// other route's answer. The split is deliberate. This index is a fixed, tiny
// answer a client can hold onto, while a vocabulary's contents are not fixed
// at all: `types` is backed by the tenant's own rows, so its values change
// whenever a merchant adds or retires one, and folding them in here would
// make every consumer re-fetch the whole set to learn a title. Names:
// statuses, types, unassigned-visibility. Fetch one with GET
// /channels/vocabularies/{name}; a client holding the qualified pair
// 'channels.<name>' builds that URL from the pair alone.
func (srv *Channels) ChannelsVocabulariesList()(*models.ChannelVocabularyIndex, error) {
	path := "/v1/channels/vocabularies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ChannelVocabularyIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ChannelVocabularyIndex
	parsed, ok := resp.Result.(models.ChannelVocabularyIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ChannelsVocabulariesGet one vocabulary with every permitted value in it,
// and a value here is more than the string the column stores: it arrives with
// the localised title and description a select puts in front of a person, and
// with a badge tone for rendering it as a status chip — `default_tone` is
// what a value carrying none falls back to, so there is always something to
// render. That is the whole reason this route exists rather than a client
// hardcoding the list. Two sources, one guarantee: what is served is what is
// in force, so no UI keeps a second copy. 'source' says which — 'schema'
// means the values are read out of the column's CHECK constraint (a value
// added to the constraint appears here even before anyone labels it, titled
// from its own key); 'table' means they are the tenant's own rows, which a
// merchant may add to, rename and retire without a release of this app.
// Values come back in author order, which is the order a select should offer.
// 'closed' says the set is exhaustive at this moment, so a value outside it
// is stale data rather than a missing label. Names: statuses, types,
// unassigned-visibility.
func (srv *Channels) ChannelsVocabulariesGet(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/channels/vocabularies/{name}")
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
	
// ChannelsDelete nothing cascades from here, and that is a statement about
// the schema rather than a reassurance: this app declares no foreign key in
// either direction, so there is nothing to cascade TO. The channel
// ASSIGNMENTS other apps hold live in Baseline, keyed by the scope slug, and
// deleting the channel does not remove them. A slug that no longer names a
// channel simply stops resolving. The consequence is that the assignments
// OUTLIVE the row. Create a channel again under a code a deleted one used and
// it silently adopts every assignment ever made against that code — which
// is the opposite of the fresh channel the call looks like it produces. If
// that is not what you want, choose a new code. The other half is the default
// flag, which nothing here protects. There is no rule that a tenant keeps at
// least one channel and none reserving the one flagged `is_default` — both
// of which the channel TYPES do have — so deleting the default is permitted
// and leaves the tenant without one. From that moment every request that
// names no channel resolves to nothing: `GET /channels/context` answers
// resolved:false with reason no_default_channel, and `POST
// /channels/visibility` hides every row that carries assignments
// (no_channel_context) while rows carrying none still follow the tenant
// policy. Promote another channel first, or restore the seeded `shop` with
// POST /channels/defaults — which brings back `shop`, never the code you
// deleted.
func (srv *Channels) ChannelsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/channels/{id}")
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
	
// ChannelsGet one row, by its uuid. The `code` is the handle everything else
// in the platform uses — it is the scope slug Baseline stores assignments
// against — and this route does not accept it: to go from a slug to the
// channel that owns it, use `GET /channels?code=…`, which answers the same
// row inside the list envelope. What this does NOT tell you is whether the
// request is in this channel. It returns an inactive channel as readily as an
// active one and applies no policy: which channel a caller is in, and what an
// unassigned row means there, is `GET /channels/context`. Answers are cached
// per tenant for 30 minutes and invalidated on any write to `channels`, so a
// read that follows someone else's write within that window can be stale by
// exactly one revision.
func (srv *Channels) ChannelsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/channels/{id}")
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
type ChannelsUpdateOptions struct {
	Code string
	IsDefault bool
	Labels interface{}
	Name string
	Position int
	Status string
	Type string
	UnassignedVisibility string
	enabledSetters map[string]bool
}
func (options ChannelsUpdateOptions) New() *ChannelsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"IsDefault": false,
		"Labels": false,
		"Name": false,
		"Position": false,
		"Status": false,
		"Type": false,
		"UnassignedVisibility": false,
	}
	return &options
}
type ChannelsUpdateOption func(*ChannelsUpdateOptions)
func (srv *Channels) WithChannelsUpdateCode(v string) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Channels) WithChannelsUpdateIsDefault(v bool) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Channels) WithChannelsUpdateLabels(v interface{}) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Channels) WithChannelsUpdateName(v string) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Channels) WithChannelsUpdatePosition(v int) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Channels) WithChannelsUpdateStatus(v string) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Channels) WithChannelsUpdateType(v string) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.Type = v
		o.enabledSetters["Type"] = true
	}
}
func (srv *Channels) WithChannelsUpdateUnassignedVisibility(v string) ChannelsUpdateOption {
	return func(o *ChannelsUpdateOptions) {
		o.UnassignedVisibility = v
		o.enabledSetters["UnassignedVisibility"] = true
	}
}
			
// ChannelsUpdate a partial write: send the fields you are changing, keep the
// rest. An empty body is a 400 rather than a no-op, so a client that computed
// no diff hears about it. Two of these fields do more than they look like
// they do, and neither is guarded the way its counterpart on the channel
// TYPES is. Sending `code` is accepted — it is only checked for scope-slug
// shape — and nothing follows it: the assignments other apps made are held
// by Baseline against the OLD slug, there is no foreign key to cascade, so a
// rename silently detaches every one of them and the channel filters as if it
// had just been created. The types route refuses the same edit outright for
// the same reason; here it is permitted, so do it deliberately or not at all.
// And `is_default` is a two-way switch here. Setting it true demotes whoever
// held it, which is what you want; setting it FALSE on the only holder leaves
// the tenant with no default channel at all, and every request that names
// none then resolves to nothing — `GET /channels/context` answers
// resolved:false with reason no_default_channel. Promote another channel in
// the same breath. On the types route sending false does nothing, precisely
// because some row must hold that flag; channels have no such rule.
func (srv *Channels) ChannelsUpdate(Id string, optionalSetters ...ChannelsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/channels/{id}")
	options := ChannelsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Type"] {
		params["type"] = options.Type
	}
	if options.enabledSetters["UnassignedVisibility"] {
		params["unassigned_visibility"] = options.UnassignedVisibility
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
