package shipping_value_lists

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// ShippingValueLists service
type ShippingValueLists struct {
	client client.Client
}

func New(clt client.Client) *ShippingValueLists {
	return &ShippingValueLists{
		client: clt,
	}
}

type ShippingServiceLevelsListOptions struct {
	Limit int
	Offset int
	enabledSetters map[string]bool
}
func (options ShippingServiceLevelsListOptions) New() *ShippingServiceLevelsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
	}
	return &options
}
type ShippingServiceLevelsListOption func(*ShippingServiceLevelsListOptions)
func (srv *ShippingValueLists) WithShippingServiceLevelsListLimit(v int) ShippingServiceLevelsListOption {
	return func(o *ShippingServiceLevelsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsListOffset(v int) ShippingServiceLevelsListOption {
	return func(o *ShippingServiceLevelsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
	
// ShippingServiceLevelsList what class of service a carrier row represents.
// This used to be a CHECK constraint, which meant a merchant with a
// night-courier tier or a two-man delivery service needed a release of this
// app to say so — and nothing in the app ever branched on the value, it
// only carried it. The set is the tenant's rows now, and the first read seeds
// it, so this never answers empty. Hand-rolled rather than a generic mount,
// because seeding is the point: it therefore honours limit/offset AND NOTHING
// ELSE. There is no `?code=` filter and no `order` — the rows always come
// back in `position` order, and a sort or a filter sent anyway is accepted,
// ignored, and answered 200.
func (srv *ShippingValueLists) ShippingServiceLevelsList(optionalSetters ...ShippingServiceLevelsListOption)(*interface{}, error) {
	path := "/v1/shipping/service-levels"
	options := ShippingServiceLevelsListOptions{}.New()
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
type ShippingServiceLevelsCreateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Tone string
	enabledSetters map[string]bool
}
func (options ShippingServiceLevelsCreateOptions) New() *ShippingServiceLevelsCreateOptions {
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
type ShippingServiceLevelsCreateOption func(*ShippingServiceLevelsCreateOptions)
func (srv *ShippingValueLists) WithShippingServiceLevelsCreateDescription(v string) ShippingServiceLevelsCreateOption {
	return func(o *ShippingServiceLevelsCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsCreateDescriptions(v interface{}) ShippingServiceLevelsCreateOption {
	return func(o *ShippingServiceLevelsCreateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsCreateIsDefault(v bool) ShippingServiceLevelsCreateOption {
	return func(o *ShippingServiceLevelsCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsCreateLabels(v interface{}) ShippingServiceLevelsCreateOption {
	return func(o *ShippingServiceLevelsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsCreatePosition(v int) ShippingServiceLevelsCreateOption {
	return func(o *ShippingServiceLevelsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsCreateTone(v string) ShippingServiceLevelsCreateOption {
	return func(o *ShippingServiceLevelsCreateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
					
// ShippingServiceLevelsCreate a service level is the class of service a
// carrier row represents, as one of the tenant's own codes. It is carried by
// `shipping_carriers.service_level` and reported on a rate as
// `carrier_service_level`; nothing in this app branches on it. A method never
// names one — it gets its level through the carrier it ships with. Reach
// for this when a merchant sells a class this app was not shipped with — a
// night courier, a two-man delivery, a same-day run. A create cannot omit
// `code` and `title`; every other column is optional or defaulted by the
// database. Two rows of this tenant may not share `code` — that is the 409.
// The code is lowercase and becomes what a carrier stores; it cannot be
// changed afterwards, because every carrier carrying it would be orphaned.
// Creating one changes nothing on its own: a carrier has to be moved onto it
// before it means anything.
func (srv *ShippingValueLists) ShippingServiceLevelsCreate(Code string, Title string, optionalSetters ...ShippingServiceLevelsCreateOption)(*models.Error, error) {
	path := "/v1/shipping/service-levels"
	options := ShippingServiceLevelsCreateOptions{}.New()
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
	
// ShippingServiceLevelsDelete there is no foreign key doing this: adding one
// to a table that starts empty would fail the migration of every existing
// tenant. The refusal lives in the handler instead.
func (srv *ShippingValueLists) ShippingServiceLevelsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/service-levels/{id}")
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
	
// ShippingServiceLevelsGet a service level is the class of service a carrier
// row represents, as one of the tenant's own codes. It is carried by
// `shipping_carriers.service_level` and reported on a rate as
// `carrier_service_level`; nothing in this app branches on it. A method never
// names one — it gets its level through the carrier it ships with. This
// reads one of them by ROW ID — which is what an editor holds after listing
// the set, and not what anything else in the platform stores. A caller
// holding the CODE (off a carrier row, or off a rate's
// `carrier_service_level`) cannot use this route: there is no `?code=` filter
// on the collection either, so read GET
// /shipping/vocabularies/service-levels, which is keyed the way the rest of
// the platform refers to these values.
func (srv *ShippingValueLists) ShippingServiceLevelsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/service-levels/{id}")
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
type ShippingServiceLevelsUpdateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Title string
	Tone string
	enabledSetters map[string]bool
}
func (options ShippingServiceLevelsUpdateOptions) New() *ShippingServiceLevelsUpdateOptions {
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
type ShippingServiceLevelsUpdateOption func(*ShippingServiceLevelsUpdateOptions)
func (srv *ShippingValueLists) WithShippingServiceLevelsUpdateDescription(v string) ShippingServiceLevelsUpdateOption {
	return func(o *ShippingServiceLevelsUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsUpdateDescriptions(v interface{}) ShippingServiceLevelsUpdateOption {
	return func(o *ShippingServiceLevelsUpdateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsUpdateIsDefault(v bool) ShippingServiceLevelsUpdateOption {
	return func(o *ShippingServiceLevelsUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsUpdateLabels(v interface{}) ShippingServiceLevelsUpdateOption {
	return func(o *ShippingServiceLevelsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsUpdatePosition(v int) ShippingServiceLevelsUpdateOption {
	return func(o *ShippingServiceLevelsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsUpdateTitle(v string) ShippingServiceLevelsUpdateOption {
	return func(o *ShippingServiceLevelsUpdateOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *ShippingValueLists) WithShippingServiceLevelsUpdateTone(v string) ShippingServiceLevelsUpdateOption {
	return func(o *ShippingServiceLevelsUpdateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
			
// ShippingServiceLevelsUpdate a service level is the class of service a
// carrier row represents, as one of the tenant's own codes. It is carried by
// `shipping_carriers.service_level` and reported on a rate as
// `carrier_service_level`; nothing in this app branches on it. A method never
// names one — it gets its level through the carrier it ships with. This
// edits the DISPLAY half of one — title, description, their locale maps,
// badge tone, position, and the default flag. Everything a carrier or a
// filter joins on stays put: the code is immutable (a different one in the
// payload is a 400, not a silent no-op), and no carrier is moved onto or off
// this level by renaming it. Moving a row's `position` does not renumber its
// neighbours — the collection is returned in position order and ties fall
// back to whatever the database returns, so a deliberate order means writing
// every row's position.
func (srv *ShippingValueLists) ShippingServiceLevelsUpdate(Id string, optionalSetters ...ShippingServiceLevelsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/service-levels/{id}")
	options := ShippingServiceLevelsUpdateOptions{}.New()
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
			
// ShippingServiceLevelsMakeDefault the flag is a single answer, not a per-row
// opinion: it is what every fallback lands on, so two defaults leave the
// result to row order and none leaves it to the seeded value. This row takes
// it and whoever was holding it is demoted in the same call — there is no
// separate write to clear the old one, and no window in which both carry it.
// Only the rows whose flag is wrong are written, so repeating the call is
// free.
func (srv *ShippingValueLists) ShippingServiceLevelsMakeDefault(Id string, Data interface{})(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/service-levels/{id}/make-default")
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

// ShippingVocabulariesList discovery for the vocabulary routes: every enum
// this app publishes, each with its name, its title and its description, and
// deliberately without its values — an index stays an index, and the set a
// value belongs to is one further call. Names: carrier-statuses,
// matrix-bases, pricing-types, service-levels, weight-units. Fetch one with
// GET /shipping/vocabularies/{name}; a client holding the qualified pair
// 'shipping.<name>' builds that URL from the pair alone. `title` and
// `description` are either one string or a locale map keyed by locale —
// every entry here carries the map, because every one of them is curated
// copy.
func (srv *ShippingValueLists) ShippingVocabulariesList()(*models.ShippingVocabularyIndex, error) {
	path := "/v1/shipping/vocabularies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ShippingVocabularyIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ShippingVocabularyIndex
	parsed, ok := resp.Result.(models.ShippingVocabularyIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// ShippingVocabulariesGet one vocabulary in full: every value it permits,
// each carrying the title to show, the description to explain it and the
// badge tone to draw it in — everything a select or a status chip needs, so
// nothing has to be labelled a second time in a client. Two sources, one
// guarantee: what is served is what is enforced, so no UI keeps a second
// copy. 'source: schema' means the values are read out of a CHECK constraint
// — a value added to the constraint appears here even before anyone labels
// it, titled from its own key, in constraint order. 'source: table' means the
// values are the TENANT's own rows (service-levels, weight-units), read per
// request and seeded on first use, so a merchant may add one without a
// release of this app; those values also carry labels/descriptions, is_system
// and is_default, and weight-units carries the conversion factor. 'closed'
// says the set is exhaustive either way, so a value outside it is stale data
// rather than a missing label. `title` and `description` — the vocabulary's
// and every value's — are either one string or a locale map keyed by
// locale: curated copy carries the map, a value titled from its own key
// carries the string. Names: carrier-statuses, matrix-bases, pricing-types,
// service-levels, weight-units.
func (srv *ShippingValueLists) ShippingVocabulariesGet(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/shipping/vocabularies/{name}")
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
type ShippingWeightUnitsListOptions struct {
	Limit int
	Offset int
	enabledSetters map[string]bool
}
func (options ShippingWeightUnitsListOptions) New() *ShippingWeightUnitsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
	}
	return &options
}
type ShippingWeightUnitsListOption func(*ShippingWeightUnitsListOptions)
func (srv *ShippingValueLists) WithShippingWeightUnitsListLimit(v int) ShippingWeightUnitsListOption {
	return func(o *ShippingWeightUnitsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsListOffset(v int) ShippingWeightUnitsListOption {
	return func(o *ShippingWeightUnitsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
	
// ShippingWeightUnitsList not a taxonomy: a unit is a code PLUS a factor, and
// the factor prices parcels. `factor` is how many kilograms one of this unit
// weighs, so a matrix keyed in one unit can price a request expressed in
// another. Exactly one row is the BASE (kg, factor 1) — the anchor every
// other factor and every stored rate tier is expressed in — and it is fixed
// at install. Seeded on first read, so this never answers empty. Like the
// service levels it is hand-rolled and honours limit/offset AND NOTHING ELSE:
// no column filter, no `order`, always `position` order, and a sort sent
// anyway is ignored rather than refused.
func (srv *ShippingValueLists) ShippingWeightUnitsList(optionalSetters ...ShippingWeightUnitsListOption)(*interface{}, error) {
	path := "/v1/shipping/weight-units"
	options := ShippingWeightUnitsListOptions{}.New()
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
type ShippingWeightUnitsCreateOptions struct {
	Description string
	Descriptions interface{}
	IsDefault bool
	Labels interface{}
	Position int
	Tone string
	enabledSetters map[string]bool
}
func (options ShippingWeightUnitsCreateOptions) New() *ShippingWeightUnitsCreateOptions {
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
type ShippingWeightUnitsCreateOption func(*ShippingWeightUnitsCreateOptions)
func (srv *ShippingValueLists) WithShippingWeightUnitsCreateDescription(v string) ShippingWeightUnitsCreateOption {
	return func(o *ShippingWeightUnitsCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsCreateDescriptions(v interface{}) ShippingWeightUnitsCreateOption {
	return func(o *ShippingWeightUnitsCreateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsCreateIsDefault(v bool) ShippingWeightUnitsCreateOption {
	return func(o *ShippingWeightUnitsCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsCreateLabels(v interface{}) ShippingWeightUnitsCreateOption {
	return func(o *ShippingWeightUnitsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsCreatePosition(v int) ShippingWeightUnitsCreateOption {
	return func(o *ShippingWeightUnitsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsCreateTone(v string) ShippingWeightUnitsCreateOption {
	return func(o *ShippingWeightUnitsCreateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
							
// ShippingWeightUnitsCreate reach for this when a merchant weighs goods in
// something this app was not shipped with — a tonne for pallet freight, a
// carat for jewellery — and wants a rate matrix keyed in it. `factor` is
// required and must be greater than 0: zero does not convert a weight, it
// divides by it, and a negative factor turns a parcel into a credit. The new
// unit is never the base — which unit anchors the others is decided at
// install, and moving it would silently reprice every weight matrix in the
// shop.
func (srv *ShippingValueLists) ShippingWeightUnitsCreate(Code string, Factor float64, Title string, optionalSetters ...ShippingWeightUnitsCreateOption)(*models.Error, error) {
	path := "/v1/shipping/weight-units"
	options := ShippingWeightUnitsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["factor"] = Factor
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
	
// ShippingWeightUnitsDelete the market check is best effort by design — the
// setting is per market and this request carries one, so another market may
// still name the unit. That case degrades to the market falling back to the
// flagged unit rather than failing its quotes.
func (srv *ShippingValueLists) ShippingWeightUnitsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/weight-units/{id}")
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
	
// ShippingWeightUnitsGet a weight unit is a code PLUS a factor — how many
// kilograms one of this unit weighs — and the factor is what prices
// parcels: a rate request expressed in one unit is converted through the two
// factors into the unit the market's tiers are keyed in. Exactly one row is
// the base (kg, factor 1), fixed at install. This reads one of them by ROW
// ID, which is what an editor holds after listing the set; a caller holding
// the CODE (a market's `weight_unit` setting, a rate request's `weight_unit`)
// has no filter for it here and should read GET
// /shipping/vocabularies/weight-units instead. Reading the factor back is NOT
// how a past quote is checked: a rate answer echoes the factors it applied in
// `basis.weight_unit_factor` and `basis.request_weight_unit_factor` precisely
// so it stays re-derivable after this row has been edited.
func (srv *ShippingValueLists) ShippingWeightUnitsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/weight-units/{id}")
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
type ShippingWeightUnitsUpdateOptions struct {
	Description string
	Descriptions interface{}
	Factor float64
	IsDefault bool
	Labels interface{}
	Position int
	Title string
	Tone string
	enabledSetters map[string]bool
}
func (options ShippingWeightUnitsUpdateOptions) New() *ShippingWeightUnitsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"Descriptions": false,
		"Factor": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Title": false,
		"Tone": false,
	}
	return &options
}
type ShippingWeightUnitsUpdateOption func(*ShippingWeightUnitsUpdateOptions)
func (srv *ShippingValueLists) WithShippingWeightUnitsUpdateDescription(v string) ShippingWeightUnitsUpdateOption {
	return func(o *ShippingWeightUnitsUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsUpdateDescriptions(v interface{}) ShippingWeightUnitsUpdateOption {
	return func(o *ShippingWeightUnitsUpdateOptions) {
		o.Descriptions = v
		o.enabledSetters["Descriptions"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsUpdateFactor(v float64) ShippingWeightUnitsUpdateOption {
	return func(o *ShippingWeightUnitsUpdateOptions) {
		o.Factor = v
		o.enabledSetters["Factor"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsUpdateIsDefault(v bool) ShippingWeightUnitsUpdateOption {
	return func(o *ShippingWeightUnitsUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsUpdateLabels(v interface{}) ShippingWeightUnitsUpdateOption {
	return func(o *ShippingWeightUnitsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsUpdatePosition(v int) ShippingWeightUnitsUpdateOption {
	return func(o *ShippingWeightUnitsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsUpdateTitle(v string) ShippingWeightUnitsUpdateOption {
	return func(o *ShippingWeightUnitsUpdateOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *ShippingValueLists) WithShippingWeightUnitsUpdateTone(v string) ShippingWeightUnitsUpdateOption {
	return func(o *ShippingWeightUnitsUpdateOptions) {
		o.Tone = v
		o.enabledSetters["Tone"] = true
	}
}
			
// ShippingWeightUnitsUpdate everything but the code and the base flag. A
// factor sent for the BASE unit is refused rather than silently ignored: it
// reads as 1 because every other factor is relative to it, so changing it
// would rescale the whole table without touching another row.
func (srv *ShippingValueLists) ShippingWeightUnitsUpdate(Id string, optionalSetters ...ShippingWeightUnitsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/weight-units/{id}")
	options := ShippingWeightUnitsUpdateOptions{}.New()
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
	if options.enabledSetters["Factor"] {
		params["factor"] = options.Factor
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
			
// ShippingWeightUnitsMakeDefault the flag is a single answer, not a per-row
// opinion: it is what every fallback lands on, so two defaults leave the
// result to row order and none leaves it to the seeded value. This row takes
// it and whoever was holding it is demoted in the same call — there is no
// separate write to clear the old one, and no window in which both carry it.
// Only the rows whose flag is wrong are written, so repeating the call is
// free.
func (srv *ShippingValueLists) ShippingWeightUnitsMakeDefault(Id string, Data interface{})(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/weight-units/{id}/make-default")
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
