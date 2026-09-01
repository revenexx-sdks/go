package shipping_carriers

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// ShippingCarriers service
type ShippingCarriers struct {
	client client.Client
}

func New(clt client.Client) *ShippingCarriers {
	return &ShippingCarriers{
		client: clt,
	}
}

type ShippingCarriersListOptions struct {
	Limit int
	Offset int
	Order string
	Code string
	Status string
	ServiceLevel string
	enabledSetters map[string]bool
}
func (options ShippingCarriersListOptions) New() *ShippingCarriersListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Code": false,
		"Status": false,
		"ServiceLevel": false,
	}
	return &options
}
type ShippingCarriersListOption func(*ShippingCarriersListOptions)
func (srv *ShippingCarriers) WithShippingCarriersListLimit(v int) ShippingCarriersListOption {
	return func(o *ShippingCarriersListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersListOffset(v int) ShippingCarriersListOption {
	return func(o *ShippingCarriersListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersListOrder(v string) ShippingCarriersListOption {
	return func(o *ShippingCarriersListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersListCode(v string) ShippingCarriersListOption {
	return func(o *ShippingCarriersListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersListStatus(v string) ShippingCarriersListOption {
	return func(o *ShippingCarriersListOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersListServiceLevel(v string) ShippingCarriersListOption {
	return func(o *ShippingCarriersListOptions) {
		o.ServiceLevel = v
		o.enabledSetters["ServiceLevel"] = true
	}
}
	
// ShippingCarriersList filterable by exact column value — `?code=`,
// `?status=` and `?service_level=` are applied as equalities and echoed back
// in `filter`. A query key that names no column of this entity is SILENTLY
// IGNORED: the page comes back unfiltered, 200, with an empty `filter`, so
// compare the echo against what you sent rather than trusting the status.
func (srv *ShippingCarriers) ShippingCarriersList(optionalSetters ...ShippingCarriersListOption)(*models.Error, error) {
	path := "/v1/shipping/carriers"
	options := ShippingCarriersListOptions{}.New()
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
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["ServiceLevel"] {
		params["service_level"] = options.ServiceLevel
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
type ShippingCarriersCreateOptions struct {
	Countries []string
	CutoffTime string
	EtaDaysMax int
	EtaDaysMin int
	HandlingDays int
	Labels interface{}
	Metadata interface{}
	Position int
	ServiceLevel string
	Status string
	TrackingUrlTemplate string
	enabledSetters map[string]bool
}
func (options ShippingCarriersCreateOptions) New() *ShippingCarriersCreateOptions {
	options.enabledSetters = map[string]bool{
		"Countries": false,
		"CutoffTime": false,
		"EtaDaysMax": false,
		"EtaDaysMin": false,
		"HandlingDays": false,
		"Labels": false,
		"Metadata": false,
		"Position": false,
		"ServiceLevel": false,
		"Status": false,
		"TrackingUrlTemplate": false,
	}
	return &options
}
type ShippingCarriersCreateOption func(*ShippingCarriersCreateOptions)
func (srv *ShippingCarriers) WithShippingCarriersCreateCountries(v []string) ShippingCarriersCreateOption {
	return func(o *ShippingCarriersCreateOptions) {
		o.Countries = v
		o.enabledSetters["Countries"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersCreateCutoffTime(v string) ShippingCarriersCreateOption {
	return func(o *ShippingCarriersCreateOptions) {
		o.CutoffTime = v
		o.enabledSetters["CutoffTime"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersCreateEtaDaysMax(v int) ShippingCarriersCreateOption {
	return func(o *ShippingCarriersCreateOptions) {
		o.EtaDaysMax = v
		o.enabledSetters["EtaDaysMax"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersCreateEtaDaysMin(v int) ShippingCarriersCreateOption {
	return func(o *ShippingCarriersCreateOptions) {
		o.EtaDaysMin = v
		o.enabledSetters["EtaDaysMin"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersCreateHandlingDays(v int) ShippingCarriersCreateOption {
	return func(o *ShippingCarriersCreateOptions) {
		o.HandlingDays = v
		o.enabledSetters["HandlingDays"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersCreateLabels(v interface{}) ShippingCarriersCreateOption {
	return func(o *ShippingCarriersCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersCreateMetadata(v interface{}) ShippingCarriersCreateOption {
	return func(o *ShippingCarriersCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersCreatePosition(v int) ShippingCarriersCreateOption {
	return func(o *ShippingCarriersCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersCreateServiceLevel(v string) ShippingCarriersCreateOption {
	return func(o *ShippingCarriersCreateOptions) {
		o.ServiceLevel = v
		o.enabledSetters["ServiceLevel"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersCreateStatus(v string) ShippingCarriersCreateOption {
	return func(o *ShippingCarriersCreateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersCreateTrackingUrlTemplate(v string) ShippingCarriersCreateOption {
	return func(o *ShippingCarriersCreateOptions) {
		o.TrackingUrlTemplate = v
		o.enabledSetters["TrackingUrlTemplate"] = true
	}
}
					
// ShippingCarriersCreate a carrier row is one company shipping one class of
// service: it owns the tracking-URL template, the service level, the transit
// days, the pickup cut-off and the handling days, and every method that ships
// with it inherits all of those unless it states its own. A carrier selling
// both a parcel and an express product is two rows. Reach for it for a
// carrier this app does not describe — a regional courier, a forwarder, an
// own fleet; for the DACH networks read GET /shipping/carriers/catalog and
// let POST /shipping/carriers/defaults write them. A create cannot omit
// `code` and `name`; every other column is optional or defaulted by the
// database. Two rows of this tenant may not share `code` — that is the 409.
// `service_level` has to name one of the tenant's own levels and
// `cutoff_time` has to be HH:MM in 24-hour UTC — both are refused rather
// than stored, because a cut-off the estimator cannot read would be dropped
// in silence and the shop would keep promising a ship date nobody computed.
// Creating a carrier quotes nothing on its own: a method has to reference it
// (`carrier_id`, or a `carrier` text equal to this code) before any of it is
// inherited.
func (srv *ShippingCarriers) ShippingCarriersCreate(Code string, Name string, optionalSetters ...ShippingCarriersCreateOption)(*models.Error, error) {
	path := "/v1/shipping/carriers"
	options := ShippingCarriersCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["name"] = Name
	if options.enabledSetters["Countries"] {
		params["countries"] = options.Countries
	}
	if options.enabledSetters["CutoffTime"] {
		params["cutoff_time"] = options.CutoffTime
	}
	if options.enabledSetters["EtaDaysMax"] {
		params["eta_days_max"] = options.EtaDaysMax
	}
	if options.enabledSetters["EtaDaysMin"] {
		params["eta_days_min"] = options.EtaDaysMin
	}
	if options.enabledSetters["HandlingDays"] {
		params["handling_days"] = options.HandlingDays
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["ServiceLevel"] {
		params["service_level"] = options.ServiceLevel
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["TrackingUrlTemplate"] {
		params["tracking_url_template"] = options.TrackingUrlTemplate
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

// ShippingCarriersCatalog the DACH set — the three German parcel networks,
// the express carriers, the AT/CH incumbents and the pallet forwarders —
// each with the tracking template, service level, transit time and pickup
// cut-off it would be created with. `seeded` marks the four a fresh install
// already has. Adding a carrier is a data change, never a code change, and a
// merchant may of course create one that is not in here at all.
func (srv *ShippingCarriers) ShippingCarriersCatalog()(*interface{}, error) {
	path := "/v1/shipping/carriers/catalog"
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

// ShippingCarriersDefaults the four networks a DACH shop is expected to have
// — DHL, DPD, GLS and UPS — created by code, and only the ones that are
// missing. The app runs this itself on `app.installed`, so a fresh install
// already has them; calling it by hand afterwards is how a tenant that
// predates a catalog entry catches up, and calling it twice costs nothing,
// because it reconciles rather than seeds. An existing row belongs to the
// merchant: only columns that are genuinely EMPTY are filled in (a tracking
// template added to the catalog after their install), never a value they set.
// Nothing is deleted.
func (srv *ShippingCarriers) ShippingCarriersDefaults()(*interface{}, error) {
	path := "/v1/shipping/carriers/defaults"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
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
	
// ShippingCarriersDelete deleting one clears `shipping_methods.carrier_id`
// rather than deleting those rows — the foreign keys decide that, not this
// route. So a method that referenced this carrier keeps working and resolves
// through its `carrier` code instead, which is also why this never answers a
// conflict — and it is the reason to prefer `status: 'retired'` where the
// carrier is merely finished. What the method silently LOSES is everything it
// was inheriting: the tracking template, the pickup cut-off, the handling
// days and the transit days. Unless its `carrier` text still matches another
// carrier, its ship date is recomputed on the market's own cut-off and
// handling settings, and a method that stated no `eta_days_min`/`max` of its
// own stops carrying a `delivery` estimate altogether. Nothing errors; the
// promise in the checkout just changes.
func (srv *ShippingCarriers) ShippingCarriersDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/carriers/{id}")
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
	
// ShippingCarriersGet a carrier row is one company shipping one class of
// service: it owns the tracking-URL template, the service level, the transit
// days, the pickup cut-off and the handling days, and every method that ships
// with it inherits all of those unless it states its own. A carrier selling
// both a parcel and an express product is two rows. Read it when you need to
// know what a method's delivery promise really is: `cutoff_time`,
// `handling_days` and `eta_days_min`/`max` are inherited from here, so a shop
// that seems to promise the wrong ship date is usually explained by this row
// rather than by the method. It does NOT say which methods ship with it —
// that is GET /shipping/methods?carrier_id=… for the ones holding a
// reference and ?carrier=… for the ones still resolving through the legacy
// code text.
func (srv *ShippingCarriers) ShippingCarriersGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/carriers/{id}")
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
type ShippingCarriersUpdateOptions struct {
	Code string
	Countries []string
	CutoffTime string
	EtaDaysMax int
	EtaDaysMin int
	HandlingDays int
	Labels interface{}
	Metadata interface{}
	Name string
	Position int
	ServiceLevel string
	Status string
	TrackingUrlTemplate string
	enabledSetters map[string]bool
}
func (options ShippingCarriersUpdateOptions) New() *ShippingCarriersUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Countries": false,
		"CutoffTime": false,
		"EtaDaysMax": false,
		"EtaDaysMin": false,
		"HandlingDays": false,
		"Labels": false,
		"Metadata": false,
		"Name": false,
		"Position": false,
		"ServiceLevel": false,
		"Status": false,
		"TrackingUrlTemplate": false,
	}
	return &options
}
type ShippingCarriersUpdateOption func(*ShippingCarriersUpdateOptions)
func (srv *ShippingCarriers) WithShippingCarriersUpdateCode(v string) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdateCountries(v []string) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.Countries = v
		o.enabledSetters["Countries"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdateCutoffTime(v string) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.CutoffTime = v
		o.enabledSetters["CutoffTime"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdateEtaDaysMax(v int) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.EtaDaysMax = v
		o.enabledSetters["EtaDaysMax"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdateEtaDaysMin(v int) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.EtaDaysMin = v
		o.enabledSetters["EtaDaysMin"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdateHandlingDays(v int) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.HandlingDays = v
		o.enabledSetters["HandlingDays"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdateLabels(v interface{}) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdateMetadata(v interface{}) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdateName(v string) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdatePosition(v int) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdateServiceLevel(v string) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.ServiceLevel = v
		o.enabledSetters["ServiceLevel"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdateStatus(v string) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *ShippingCarriers) WithShippingCarriersUpdateTrackingUrlTemplate(v string) ShippingCarriersUpdateOption {
	return func(o *ShippingCarriersUpdateOptions) {
		o.TrackingUrlTemplate = v
		o.enabledSetters["TrackingUrlTemplate"] = true
	}
}
			
// ShippingCarriersUpdate a carrier row is one company shipping one class of
// service: it owns the tracking-URL template, the service level, the transit
// days, the pickup cut-off and the handling days, and every method that ships
// with it inherits all of those unless it states its own. A carrier selling
// both a parcel and an express product is two rows. A partial update — send
// only what changes, which is where a carrier is paused, given a different
// tracking template, or moved to another pickup cut-off or transit time. This
// is the one switch that acts on several methods at once, in both directions.
// Moving `status` off 'active' takes every method that ships with this
// carrier out of POST /shipping/rates with a reason, which beats disabling
// each of them and forgetting one; tracking links are deliberately not gated
// on it, so a retired carrier's old shipments stay resolvable. Editing
// `cutoff_time`, `handling_days` or `eta_days_min`/`max` MOVES THE PROMISED
// SHIP DATE of every method that states none of its own: the estimator adds
// the handling days, then one further day when the cut-off has already passed
// at the instant being evaluated — compared at or after, in UTC, and as
// calendar days that do not skip a weekend. Two rows of this tenant may not
// share `code` — that is the 409.
func (srv *ShippingCarriers) ShippingCarriersUpdate(Id string, optionalSetters ...ShippingCarriersUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/shipping/carriers/{id}")
	options := ShippingCarriersUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Countries"] {
		params["countries"] = options.Countries
	}
	if options.enabledSetters["CutoffTime"] {
		params["cutoff_time"] = options.CutoffTime
	}
	if options.enabledSetters["EtaDaysMax"] {
		params["eta_days_max"] = options.EtaDaysMax
	}
	if options.enabledSetters["EtaDaysMin"] {
		params["eta_days_min"] = options.EtaDaysMin
	}
	if options.enabledSetters["HandlingDays"] {
		params["handling_days"] = options.HandlingDays
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
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
	if options.enabledSetters["ServiceLevel"] {
		params["service_level"] = options.ServiceLevel
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["TrackingUrlTemplate"] {
		params["tracking_url_template"] = options.TrackingUrlTemplate
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
type ShippingTrackingOptions struct {
	Country string
	PostalCode string
	TrackingCode string
	enabledSetters map[string]bool
}
func (options ShippingTrackingOptions) New() *ShippingTrackingOptions {
	options.enabledSetters = map[string]bool{
		"Country": false,
		"PostalCode": false,
		"TrackingCode": false,
	}
	return &options
}
type ShippingTrackingOption func(*ShippingTrackingOptions)
func (srv *ShippingCarriers) WithShippingTrackingCountry(v string) ShippingTrackingOption {
	return func(o *ShippingTrackingOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *ShippingCarriers) WithShippingTrackingPostalCode(v string) ShippingTrackingOption {
	return func(o *ShippingTrackingOptions) {
		o.PostalCode = v
		o.enabledSetters["PostalCode"] = true
	}
}
func (srv *ShippingCarriers) WithShippingTrackingTrackingCode(v string) ShippingTrackingOption {
	return func(o *ShippingTrackingOptions) {
		o.TrackingCode = v
		o.enabledSetters["TrackingCode"] = true
	}
}
			
// ShippingTracking hand in a carrier code and the tracking number printed on
// the label, and this answers the URL a buyer follows. The carrier owns the
// URL format, so nobody else has to. `order_shipments` stores a tracking_url
// per shipment today, which is one carrier's URL shape copied into every row
// — the day it changes, every historic link is wrong. Ask here instead.
// Tracking is NOT gated on carrier status: a retired carrier's old shipments
// stay resolvable.
func (srv *ShippingCarriers) ShippingTracking(Carrier string, optionalSetters ...ShippingTrackingOption)(*models.Error, error) {
	path := "/v1/shipping/tracking"
	options := ShippingTrackingOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["carrier"] = Carrier
	if options.enabledSetters["Country"] {
		params["country"] = options.Country
	}
	if options.enabledSetters["PostalCode"] {
		params["postal_code"] = options.PostalCode
	}
	if options.enabledSetters["TrackingCode"] {
		params["tracking_code"] = options.TrackingCode
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
