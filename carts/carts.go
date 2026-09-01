package carts

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Carts service
type Carts struct {
	client client.Client
}

func New(clt client.Client) *Carts {
	return &Carts{
		client: clt,
	}
}

type CartsListOptions struct {
	Id string
	Name string
	Status string
	ContactId string
	SessionKey string
	ChannelId string
	Currency string
	IsCurrent bool
	ItemCount int
	Subtotal float64
	AbandonedAt string
	OrderedAt string
	OrderRef string
	MergedIntoCartId string
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options CartsListOptions) New() *CartsListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Name": false,
		"Status": false,
		"ContactId": false,
		"SessionKey": false,
		"ChannelId": false,
		"Currency": false,
		"IsCurrent": false,
		"ItemCount": false,
		"Subtotal": false,
		"AbandonedAt": false,
		"OrderedAt": false,
		"OrderRef": false,
		"MergedIntoCartId": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type CartsListOption func(*CartsListOptions)
func (srv *Carts) WithCartsListId(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Carts) WithCartsListName(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Carts) WithCartsListStatus(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Carts) WithCartsListContactId(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Carts) WithCartsListSessionKey(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.SessionKey = v
		o.enabledSetters["SessionKey"] = true
	}
}
func (srv *Carts) WithCartsListChannelId(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Carts) WithCartsListCurrency(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Carts) WithCartsListIsCurrent(v bool) CartsListOption {
	return func(o *CartsListOptions) {
		o.IsCurrent = v
		o.enabledSetters["IsCurrent"] = true
	}
}
func (srv *Carts) WithCartsListItemCount(v int) CartsListOption {
	return func(o *CartsListOptions) {
		o.ItemCount = v
		o.enabledSetters["ItemCount"] = true
	}
}
func (srv *Carts) WithCartsListSubtotal(v float64) CartsListOption {
	return func(o *CartsListOptions) {
		o.Subtotal = v
		o.enabledSetters["Subtotal"] = true
	}
}
func (srv *Carts) WithCartsListAbandonedAt(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.AbandonedAt = v
		o.enabledSetters["AbandonedAt"] = true
	}
}
func (srv *Carts) WithCartsListOrderedAt(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.OrderedAt = v
		o.enabledSetters["OrderedAt"] = true
	}
}
func (srv *Carts) WithCartsListOrderRef(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.OrderRef = v
		o.enabledSetters["OrderRef"] = true
	}
}
func (srv *Carts) WithCartsListMergedIntoCartId(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.MergedIntoCartId = v
		o.enabledSetters["MergedIntoCartId"] = true
	}
}
func (srv *Carts) WithCartsListCreatedAt(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Carts) WithCartsListUpdatedAt(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *Carts) WithCartsListLimit(v int) CartsListOption {
	return func(o *CartsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Carts) WithCartsListOffset(v int) CartsListOption {
	return func(o *CartsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Carts) WithCartsListOrder(v string) CartsListOption {
	return func(o *CartsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// CartsList the cart index, and the route a storefront resumes a session
// with: `?contact_id=…` for a customer's carts, `?session_key=…` for a
// guest's, and `?is_current=true` alongside one of those two for the single
// cart carts.activate last marked — this list is the ONLY place that flag
// can be read back, and on its own the filter selects every current cart in
// the tenant. Filters are exact equality and never a search, unknown keys are
// dropped rather than refused, and `filter` echoes what was understood. Each
// row carries its own stored totals — `item_count` is the sum of the line
// QUANTITIES, not the number of lines — but never its lines: those are one
// call per cart. With no filter at all this is every cart the tenant holds,
// paged, which is a report rather than a session lookup.
func (srv *Carts) CartsList(optionalSetters ...CartsListOption)(*models.Error, error) {
	path := "/v1/carts"
	options := CartsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["SessionKey"] {
		params["session_key"] = options.SessionKey
	}
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["IsCurrent"] {
		params["is_current"] = options.IsCurrent
	}
	if options.enabledSetters["ItemCount"] {
		params["item_count"] = options.ItemCount
	}
	if options.enabledSetters["Subtotal"] {
		params["subtotal"] = options.Subtotal
	}
	if options.enabledSetters["AbandonedAt"] {
		params["abandoned_at"] = options.AbandonedAt
	}
	if options.enabledSetters["OrderedAt"] {
		params["ordered_at"] = options.OrderedAt
	}
	if options.enabledSetters["OrderRef"] {
		params["order_ref"] = options.OrderRef
	}
	if options.enabledSetters["MergedIntoCartId"] {
		params["merged_into_cart_id"] = options.MergedIntoCartId
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
type CartsCreateOptions struct {
	ChannelId string
	ContactId string
	Currency string
	IsCurrent bool
	Metadata interface{}
	Name string
	SessionKey string
	enabledSetters map[string]bool
}
func (options CartsCreateOptions) New() *CartsCreateOptions {
	options.enabledSetters = map[string]bool{
		"ChannelId": false,
		"ContactId": false,
		"Currency": false,
		"IsCurrent": false,
		"Metadata": false,
		"Name": false,
		"SessionKey": false,
	}
	return &options
}
type CartsCreateOption func(*CartsCreateOptions)
func (srv *Carts) WithCartsCreateChannelId(v string) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Carts) WithCartsCreateContactId(v string) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *Carts) WithCartsCreateCurrency(v string) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Carts) WithCartsCreateIsCurrent(v bool) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.IsCurrent = v
		o.enabledSetters["IsCurrent"] = true
	}
}
func (srv *Carts) WithCartsCreateMetadata(v interface{}) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Carts) WithCartsCreateName(v string) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Carts) WithCartsCreateSessionKey(v string) CartsCreateOption {
	return func(o *CartsCreateOptions) {
		o.SessionKey = v
		o.enabledSetters["SessionKey"] = true
	}
}
	
// CartsCreate opens an empty cart. The one thing it requires is an OWNER —
// `contact_id` for a signed-in customer or `session_key` for a guest, never
// neither: that is a database check on the table, and this route refuses it
// first with a 400 so the caller gets a sentence rather than a constraint
// name. Everything else is defaulted: the name 'Cart', currency EUR, status
// 'active', both totals 0. No column of a cart is unique, so one owner may
// hold as many carts as they like — unless the tenant's
// `multi_cart_enabled` is off, in which case a second ACTIVE cart for the
// same owner answers 409 naming the cart that already exists, because a
// storefront that hit that wants to fill THAT cart. Send `is_current: true`
// to have the new cart made current in the same call, which clears the flag
// on every sibling of the same owner. Lines are added afterwards, one call
// each or one bulk replace.
func (srv *Carts) CartsCreate(optionalSetters ...CartsCreateOption)(*models.Error, error) {
	path := "/v1/carts"
	options := CartsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["IsCurrent"] {
		params["is_current"] = options.IsCurrent
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["SessionKey"] {
		params["session_key"] = options.SessionKey
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
type CartsClaimOptions struct {
	Strategy string
	TargetCartId string
	enabledSetters map[string]bool
}
func (options CartsClaimOptions) New() *CartsClaimOptions {
	options.enabledSetters = map[string]bool{
		"Strategy": false,
		"TargetCartId": false,
	}
	return &options
}
type CartsClaimOption func(*CartsClaimOptions)
func (srv *Carts) WithCartsClaimStrategy(v string) CartsClaimOption {
	return func(o *CartsClaimOptions) {
		o.Strategy = v
		o.enabledSetters["Strategy"] = true
	}
}
func (srv *Carts) WithCartsClaimTargetCartId(v string) CartsClaimOption {
	return func(o *CartsClaimOptions) {
		o.TargetCartId = v
		o.enabledSetters["TargetCartId"] = true
	}
}
					
// CartsClaim the login call, and the one route that turns a guest into a
// customer: every ACTIVE cart of one session_key is handed to a contact_id,
// which is what a storefront fires the moment somebody signs in with a basket
// already filled. There are two ways it can land, and the body picks between
// them. Without a target_cart_id the session carts are ADOPTED as they stand
// — same carts, same lines, contact_id set and session_key cleared, nothing
// copied and nothing closed. With a target_cart_id they are instead folded
// into that cart, which survives while each session cart is closed as status
// merged; 'adopted' and 'merged' in the answer say which of the two happened
// to each one. With a target cart, cart_merge_strategy decides what happens
// to the target's OWN lines: 'merge' keeps them and folds the session lines
// in, 'replace' clears them first. 'strategy' overrides it for one call
// (merge | replace); the answer always echoes which one ran and how many
// lines a replace removed.
func (srv *Carts) CartsClaim(ContactId string, SessionKey string, optionalSetters ...CartsClaimOption)(*models.Error, error) {
	path := "/v1/carts/claim"
	options := CartsClaimOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["contact_id"] = ContactId
	params["session_key"] = SessionKey
	if options.enabledSetters["Strategy"] {
		params["strategy"] = options.Strategy
	}
	if options.enabledSetters["TargetCartId"] {
		params["target_cart_id"] = options.TargetCartId
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
type CartsMaintenanceRunOptions struct {
	DryRun bool
	enabledSetters map[string]bool
}
func (options CartsMaintenanceRunOptions) New() *CartsMaintenanceRunOptions {
	options.enabledSetters = map[string]bool{
		"DryRun": false,
	}
	return &options
}
type CartsMaintenanceRunOption func(*CartsMaintenanceRunOptions)
func (srv *Carts) WithCartsMaintenanceRunDryRun(v bool) CartsMaintenanceRunOption {
	return func(o *CartsMaintenanceRunOptions) {
		o.DryRun = v
		o.enabledSetters["DryRun"] = true
	}
}
	
// CartsMaintenanceRun two sweeps in one pass. abandon_after_minutes marks
// active carts that have sat untouched past the window as abandoned (stamping
// abandoned_at, which nothing else in the platform ever sets — without this
// the abandonment funnel is empty by construction, not empty because nobody
// abandons carts). cart_ttl_days / guest_cart_ttl_days then DELETE carts past
// their retention window, line items included; both default to 0 (never), and
// an 'ordered' cart is never touched at any setting because it is the source
// record of a sale. Send dry_run to get the same counts and cart ids while
// writing nothing. The platform runs this per installed tenant on the
// schedule; it is idempotent, so calling it by hand between ticks is safe.
func (srv *Carts) CartsMaintenanceRun(optionalSetters ...CartsMaintenanceRunOption)(*models.CartMaintenanceResult, error) {
	path := "/v1/carts/maintenance/run"
	options := CartsMaintenanceRunOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["DryRun"] {
		params["dry_run"] = options.DryRun
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

		parsed := models.CartMaintenanceResult{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CartMaintenanceResult
	parsed, ok := resp.Result.(models.CartMaintenanceResult)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// CartsMerge which of the two carts survives is the whole question, and the
// answer is the TARGET: the source's lines are COPIED into the target, the
// target keeps every line it already had, its totals are recomputed, and it
// is the cart the caller goes on using. Nothing is replaced and nothing is
// moved — the source keeps its own line rows and is closed with status
// 'merged' and `merged_into_cart_id` pointing at the target, so a merged cart
// stays readable as the record of what went where. On the way in, a plain
// product line with the same product/sku AND the same `unit_price` as a line
// already in the target adds its quantity to that line; configured and custom
// lines always land as new ones. Both carts must be active and must differ,
// and the tenant's line limits are enforced on the target as the copies land
// (422). Reach for carts.merge_into where the caller holds one cart id and
// not two.
func (srv *Carts) CartsMerge(SourceCartId string, TargetCartId string)(*models.Error, error) {
	path := "/v1/carts/merge"
	params := map[string]interface{}{}
	params["source_cart_id"] = SourceCartId
	params["target_cart_id"] = TargetCartId
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

// CartsVocabulariesList discovery for the vocabulary routes: every enum this
// app publishes, each as its name, its title and its description and nothing
// else. The VALUES are deliberately not here — this is the index a client
// builds a menu from, and one call per vocabulary fills it. Names:
// io-apply-modes, io-directions, io-entities, io-formats, item-types,
// statuses. Fetch one with GET /carts/vocabularies/{name}; a client holding
// the qualified pair 'carts.<name>' builds that URL from the pair alone.
func (srv *Carts) CartsVocabulariesList()(*models.CartVocabularyIndex, error) {
	path := "/v1/carts/vocabularies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.CartVocabularyIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.CartVocabularyIndex
	parsed, ok := resp.Result.(models.CartVocabularyIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// CartsVocabulariesGet one vocabulary with its values filled in — every
// value permitted by the column behind it, each carrying the key the database
// stores, a human title, a description where one was written and the badge
// tone a UI should render it in, which is everything a select or a status
// chip needs from one call. The values are read out of the column's CHECK
// constraint, so the served set IS the enforced set and the two cannot drift
// — a value added to the constraint appears here even before anyone labels
// it, titled from its own key. Values come back in constraint order, which is
// the order a select should offer. 'closed' says the set is exhaustive, so a
// value outside it is stale data rather than a missing label. Names:
// io-apply-modes, io-directions, io-entities, io-formats, item-types,
// statuses.
func (srv *Carts) CartsVocabulariesGet(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/carts/vocabularies/{name}")
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
	
// CartsDelete removes the cart row and, through the `on delete cascade` on
// `cart_items.cart_id`, every line in it. There is no soft delete and no
// undo. One status is protected and it is protected permanently: an 'ordered'
// cart is the source record of a sale — the order carries its id in
// `cart_id` and the order.placed event records it — so this route refuses
// it with 400 and there is no flag, no force and no lifecycle route that
// makes it deletable. Do not go looking for one. 'active', 'abandoned' and
// 'merged' are all deletable, which is deliberate and is the same set the
// cart-maintenance sweep removes on a retention window: clearing out
// abandoned guest carts is the main thing anyone deletes a cart for, and a
// merged cart's lines were COPIED into the target, which still holds them.
// What the delete does NOT take with it is the trail: `merged_into_cart_id`
// is a plain uuid column and not a foreign key, so deleting a cart that other
// carts were merged INTO leaves those carts pointing at a row that no longer
// exists, and nothing refuses the delete or clears the pointer — the
// retention sweep does the same, so this is a property of the column and not
// of this route. For a cart a buyer simply walked away from, carts.abandon
// keeps the row and the funnel; for deleting on a retention window, the
// cart-maintenance sweep does it per market and can be asked first with
// `dry_run`.
func (srv *Carts) CartsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}")
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
	
// CartsGet one cart with its owner, its totals and its lifecycle stamps —
// and none of its lines: those are a separate call (`GET
// /carts/{cart_id}/items`), because a cart row is small and a filled cart is
// not. The two totals are derived and stored, never taken from a caller:
// `item_count` is the sum of the line QUANTITIES rather than the number of
// lines (two lines of five pieces answer 10, not 2) and `subtotal` the sum of
// the line totals, net of shipping and tax; both are recomputed after every
// line write. `status` says what may still be done — only an 'active' cart
// accepts a write of any kind, 'abandoned' is the one reversible ending, and
// a 'merged' cart carries `merged_into_cart_id`, which is the trail to the
// cart its lines were copied into.
func (srv *Carts) CartsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}")
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
type CartsUpdateOptions struct {
	ChannelId string
	Currency string
	Metadata interface{}
	Name string
	enabledSetters map[string]bool
}
func (options CartsUpdateOptions) New() *CartsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"ChannelId": false,
		"Currency": false,
		"Metadata": false,
		"Name": false,
	}
	return &options
}
type CartsUpdateOption func(*CartsUpdateOptions)
func (srv *Carts) WithCartsUpdateChannelId(v string) CartsUpdateOption {
	return func(o *CartsUpdateOptions) {
		o.ChannelId = v
		o.enabledSetters["ChannelId"] = true
	}
}
func (srv *Carts) WithCartsUpdateCurrency(v string) CartsUpdateOption {
	return func(o *CartsUpdateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Carts) WithCartsUpdateMetadata(v interface{}) CartsUpdateOption {
	return func(o *CartsUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *Carts) WithCartsUpdateName(v string) CartsUpdateOption {
	return func(o *CartsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
			
// CartsUpdate the four columns a cart's own editing screen owns, and only
// those: `name`, `currency`, `channel_id` and `metadata`. Everything else
// about a cart is either derived or a lifecycle move, and both are
// deliberately out of reach here — `item_count` and `subtotal` are
// recomputed from the lines, `status` travels through the action routes
// (activate, abandon, reopen, order, merge) so that every transition is
// guarded, and `market_id` is the platform's scope on the row rather than a
// column this app writes. A payload carrying none of the four answers 400
// rather than storing nothing quietly, so a caller never believes an ignored
// field was saved. The owner is not updatable either: a guest cart becomes a
// customer's through carts.claim.
func (srv *Carts) CartsUpdate(Id string, optionalSetters ...CartsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}")
	options := CartsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["ChannelId"] {
		params["channel_id"] = options.ChannelId
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
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
	
// CartsAbandon the by-hand half of the abandonment funnel: an active cart
// becomes 'abandoned', `abandoned_at` is stamped, and `is_current` is cleared
// — so its owner is left with no current cart until another one is
// activated. Nothing else in the platform writes `abandoned_at`; the only
// other writer is the cart-maintenance sweep, which does exactly this once a
// cart has sat untouched past the market's `abandon_after_minutes`. This is
// the one reversible ending: the lines are untouched throughout and
// carts.reopen takes the cart back. Only an active cart can be abandoned —
// an ordered or merged cart is already finished and answers 400 naming the
// status it actually holds.
func (srv *Carts) CartsAbandon(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}/abandon")
	params := map[string]interface{}{}
	params["id"] = Id
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
	
// CartsActivate activate writes exactly one thing: `is_current` on this cart,
// cleared on every other cart of the same owner (the same contact_id, or the
// same session_key). It does NOT change the status — an active cart stays
// active, and only an active cart may be made current. Read it back with `GET
// /carts?is_current=true` plus the owner: that filter is the only way to see
// what this route wrote, and a storefront resuming a session is its main
// caller. The flag is cleared again by abandoning, ordering or merging the
// cart, so an owner can legitimately have no current cart at all.
func (srv *Carts) CartsActivate(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}/activate")
	params := map[string]interface{}{}
	params["id"] = Id
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
			
// CartsMergeInto identical to carts.merge, with the SOURCE taken from the
// path — which is what makes the merge reachable from anything holding one
// cart and only one: a Cockpit row action, a detail page, a storefront
// session. The cart in the path is therefore the one that ends: its lines are
// copied into the `target_cart_id` named in the body, that target keeps its
// own lines and survives, and the path cart is closed with status 'merged'
// and `merged_into_cart_id` pointing at it. Getting the two the wrong way
// round is the mistake this route exists to make hard, so read the path id as
// "the cart I am giving away". Both carts must be active and must differ.
func (srv *Carts) CartsMergeInto(Id string, TargetCartId string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}/merge-into")
	params := map[string]interface{}{}
	params["id"] = Id
	params["target_cart_id"] = TargetCartId
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
type CartsOrderOptions struct {
	OrderRef string
	enabledSetters map[string]bool
}
func (options CartsOrderOptions) New() *CartsOrderOptions {
	options.enabledSetters = map[string]bool{
		"OrderRef": false,
	}
	return &options
}
type CartsOrderOption func(*CartsOrderOptions)
func (srv *Carts) WithCartsOrderOrderRef(v string) CartsOrderOption {
	return func(o *CartsOrderOptions) {
		o.OrderRef = v
		o.enabledSetters["OrderRef"] = true
	}
}
			
// CartsOrder the hand-over to order management, and the end of the cart as a
// workspace: an ACTIVE cart becomes 'ordered', ordered_at is stamped, and the
// order_ref the call carries — order management's own number for the order
// this cart became — is stored on the cart, which is what lets anyone
// filter their way from an order number back to the cart behind it. Nothing
// moves out of 'ordered' afterwards, and no route will delete it. The
// conversion applies the two tenant decisions a cart cannot make for itself.
// price_snapshot_mode (snapshot | live) settles which of a line's two prices
// is charged — the snapshot the buyer was shown, or the current unit_price
// — and the cart's subtotal is rewritten to match, so cart and order can
// never disagree; 'pricing' reports the mode, the lines it rewrote and the
// subtotal on both sides. convert_reserves_stock (never | request | require)
// decides whether inventories is asked to hold the lines; at 'require' a
// refusal answers 409 and the cart stays active and unchanged. The
// reservation is attempted BEFORE anything is written.
func (srv *Carts) CartsOrder(Id string, optionalSetters ...CartsOrderOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}/order")
	options := CartsOrderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["OrderRef"] {
		params["order_ref"] = options.OrderRef
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
	
// CartsReopen takes an abandoned cart back to 'active' with its lines exactly
// as they were — what a storefront calls when a buyer follows a recovery
// mail, and the way out of the 400 a write gets on a cart the maintenance
// sweep closed while nobody was looking. It also CLEARS `abandoned_at`, so a
// cart that was abandoned and reopened leaves nothing behind in the funnel:
// the funnel counts carts that are still abandoned, not carts that ever were.
// It does not restore `is_current` — a reopened cart is active but not
// current until carts.activate says so. Only an abandoned cart may be
// reopened; 'ordered' and 'merged' are final and answer 400 naming the status
// the cart holds.
func (srv *Carts) CartsReopen(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/carts/{id}/reopen")
	params := map[string]interface{}{}
	params["id"] = Id
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
