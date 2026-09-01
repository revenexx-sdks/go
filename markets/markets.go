package markets

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Markets service
type Markets struct {
	client client.Client
}

func New(clt client.Client) *Markets {
	return &Markets{
		client: clt,
	}
}

type MarketsListOptions struct {
	Id string
	Code string
	Name string
	Labels string
	Currency string
	Status string
	IsDefault bool
	Position int
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options MarketsListOptions) New() *MarketsListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Code": false,
		"Name": false,
		"Labels": false,
		"Currency": false,
		"Status": false,
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
type MarketsListOption func(*MarketsListOptions)
func (srv *Markets) WithMarketsListId(v string) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Markets) WithMarketsListCode(v string) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsListName(v string) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Markets) WithMarketsListLabels(v string) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Markets) WithMarketsListCurrency(v string) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Markets) WithMarketsListStatus(v string) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Markets) WithMarketsListIsDefault(v bool) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsListPosition(v int) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsListCreatedAt(v string) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Markets) WithMarketsListUpdatedAt(v string) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *Markets) WithMarketsListLimit(v int) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Markets) WithMarketsListOffset(v int) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Markets) WithMarketsListOrder(v string) MarketsListOption {
	return func(o *MarketsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// MarketsList every column is an exact-match filter and they combine with AND
// (?code=northwind); each one is declared as a query parameter above. A
// `?column=value` this entity does not have is DROPPED rather than refused
// — the call answers 200 with the unfiltered list — and `filter` echoes
// what was actually applied, which is the only way to tell that apart from a
// filter that matched nothing.
func (srv *Markets) MarketsList(optionalSetters ...MarketsListOption)(*models.Error, error) {
	path := "/v1/markets"
	options := MarketsListOptions{}.New()
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
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
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
type MarketsCreateOptions struct {
	Currency string
	IsDefault bool
	Labels interface{}
	Position int
	Status string
	enabledSetters map[string]bool
}
func (options MarketsCreateOptions) New() *MarketsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Currency": false,
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Status": false,
	}
	return &options
}
type MarketsCreateOption func(*MarketsCreateOptions)
func (srv *Markets) WithMarketsCreateCurrency(v string) MarketsCreateOption {
	return func(o *MarketsCreateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Markets) WithMarketsCreateIsDefault(v bool) MarketsCreateOption {
	return func(o *MarketsCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsCreateLabels(v interface{}) MarketsCreateOption {
	return func(o *MarketsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Markets) WithMarketsCreatePosition(v int) MarketsCreateOption {
	return func(o *MarketsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsCreateStatus(v string) MarketsCreateOption {
	return func(o *MarketsCreateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
					
// MarketsCreate a market needs a 'code' and a 'name' — currency defaults to
// EUR, status to active. To get a market that can actually trade, clone an
// existing one instead: POST /markets/{id}/clone.
func (srv *Markets) MarketsCreate(Code string, Name string, optionalSetters ...MarketsCreateOption)(*models.Error, error) {
	path := "/v1/markets"
	options := MarketsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	params["name"] = Name
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
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
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
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

// MarketsLocalePolicy how this tenant keys its translations, resolved for a
// surface that stands in no market at all. The Cockpit edits a tenant
// BASELINE when no market is selected, and a baseline value has to be
// readable by every market — so the locale set answered here is the UNION
// of every market's locales, each one already resolved to the key it is
// written under, not one market's list and not a pair of setting names to
// re-implement. Each entry names the markets that asked for that locale: an
// editor listing six inputs without saying who needs them invites
// translations nobody will ever read. Write/read keys follow the same two
// settings as the per-market answer, so a baseline and a market value can
// never be keyed differently.
func (srv *Markets) MarketsLocalePolicy()(*models.TenantLocalePolicy, error) {
	path := "/v1/markets/locale-policy"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.TenantLocalePolicy{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.TenantLocalePolicy
	parsed, ok := resp.Result.(models.TenantLocalePolicy)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// MarketsVocabularies every closed value set this app owns, listed by name
// with its title and its description but WITHOUT its values — enough to
// build a menu of them, and a name to fetch one by when a select box actually
// needs the values. Static per app version; nothing about a tenant changes
// it. It reads no table and takes no parameter, so 200 is the only answer it
// has beyond the gateway's own.
func (srv *Markets) MarketsVocabularies()(*models.MarketsVocabularyIndex, error) {
	path := "/v1/markets/vocabularies"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MarketsVocabularyIndex{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MarketsVocabularyIndex
	parsed, ok := resp.Result.(models.MarketsVocabularyIndex)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// MarketsVocabulary one value set in full: every value the column may hold,
// in the order it may hold them, with the copy and the badge tone a client
// renders each one as. The values are not kept in a list beside the database,
// they are parsed out of the CHECK constraint in this app's own schema.json
// — so the set served here IS the set enforced on a write, and a select box
// built from it cannot offer a value the write would then refuse. A name
// outside the declared enum is a 404 rather than an empty list — an empty
// vocabulary and an unknown one mean different things to a select box.
func (srv *Markets) MarketsVocabulary(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/markets/vocabularies/{name}")
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
	
// MarketsDelete deleting a market takes its locales, currencies and tax
// classes with it: all three carry an ON DELETE CASCADE onto markets.id, so
// this is never refused for having children.
func (srv *Markets) MarketsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}")
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
	
// MarketsGet resolved by uuid only — unlike /readiness, /clone, /backfill
// and /make-default, a market CODE here is a 400 rather than a lookup.
func (srv *Markets) MarketsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}")
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
type MarketsUpdateOptions struct {
	Code string
	Currency string
	IsDefault bool
	Labels interface{}
	Name string
	Position int
	Status string
	enabledSetters map[string]bool
}
func (options MarketsUpdateOptions) New() *MarketsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Currency": false,
		"IsDefault": false,
		"Labels": false,
		"Name": false,
		"Position": false,
		"Status": false,
	}
	return &options
}
type MarketsUpdateOption func(*MarketsUpdateOptions)
func (srv *Markets) WithMarketsUpdateCode(v string) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsUpdateCurrency(v string) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Markets) WithMarketsUpdateIsDefault(v bool) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsUpdateLabels(v interface{}) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Markets) WithMarketsUpdateName(v string) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Markets) WithMarketsUpdatePosition(v int) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsUpdateStatus(v string) MarketsUpdateOption {
	return func(o *MarketsUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
			
// MarketsUpdate partial: omitted fields keep their value.
func (srv *Markets) MarketsUpdate(Id string, optionalSetters ...MarketsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}")
	options := MarketsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
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
type MarketsBackfillOptions struct {
	Currencies bool
	Locales bool
	TaxClasses bool
	enabledSetters map[string]bool
}
func (options MarketsBackfillOptions) New() *MarketsBackfillOptions {
	options.enabledSetters = map[string]bool{
		"Currencies": false,
		"Locales": false,
		"TaxClasses": false,
	}
	return &options
}
type MarketsBackfillOption func(*MarketsBackfillOptions)
func (srv *Markets) WithMarketsBackfillCurrencies(v bool) MarketsBackfillOption {
	return func(o *MarketsBackfillOptions) {
		o.Currencies = v
		o.enabledSetters["Currencies"] = true
	}
}
func (srv *Markets) WithMarketsBackfillLocales(v bool) MarketsBackfillOption {
	return func(o *MarketsBackfillOptions) {
		o.Locales = v
		o.enabledSetters["Locales"] = true
	}
}
func (srv *Markets) WithMarketsBackfillTaxClasses(v bool) MarketsBackfillOption {
	return func(o *MarketsBackfillOptions) {
		o.TaxClasses = v
		o.enabledSetters["TaxClasses"] = true
	}
}
					
// MarketsBackfill repairs the market in the path out of a source market that
// is already right. The two are compared by CODE, collection by collection,
// and only the codes this market does not already carry are added — so a
// locale, a currency or a tax class it already holds is left exactly as the
// merchant left it, rate included, and is never overwritten. Both the path id
// and `source` are resolved by uuid OR by market code. Idempotent: running it
// twice adds nothing the second time.
func (srv *Markets) MarketsBackfill(Id string, Source string, optionalSetters ...MarketsBackfillOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}/backfill")
	options := MarketsBackfillOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	params["source"] = Source
	if options.enabledSetters["Currencies"] {
		params["currencies"] = options.Currencies
	}
	if options.enabledSetters["Locales"] {
		params["locales"] = options.Locales
	}
	if options.enabledSetters["TaxClasses"] {
		params["tax_classes"] = options.TaxClasses
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
type MarketsCloneOptions struct {
	CopyCurrencies bool
	CopyLocales bool
	CopyTaxClasses bool
	Currency string
	Name string
	Status string
	enabledSetters map[string]bool
}
func (options MarketsCloneOptions) New() *MarketsCloneOptions {
	options.enabledSetters = map[string]bool{
		"CopyCurrencies": false,
		"CopyLocales": false,
		"CopyTaxClasses": false,
		"Currency": false,
		"Name": false,
		"Status": false,
	}
	return &options
}
type MarketsCloneOption func(*MarketsCloneOptions)
func (srv *Markets) WithMarketsCloneCopyCurrencies(v bool) MarketsCloneOption {
	return func(o *MarketsCloneOptions) {
		o.CopyCurrencies = v
		o.enabledSetters["CopyCurrencies"] = true
	}
}
func (srv *Markets) WithMarketsCloneCopyLocales(v bool) MarketsCloneOption {
	return func(o *MarketsCloneOptions) {
		o.CopyLocales = v
		o.enabledSetters["CopyLocales"] = true
	}
}
func (srv *Markets) WithMarketsCloneCopyTaxClasses(v bool) MarketsCloneOption {
	return func(o *MarketsCloneOptions) {
		o.CopyTaxClasses = v
		o.enabledSetters["CopyTaxClasses"] = true
	}
}
func (srv *Markets) WithMarketsCloneCurrency(v string) MarketsCloneOption {
	return func(o *MarketsCloneOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *Markets) WithMarketsCloneName(v string) MarketsCloneOption {
	return func(o *MarketsCloneOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Markets) WithMarketsCloneStatus(v string) MarketsCloneOption {
	return func(o *MarketsCloneOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
					
// MarketsClone creates a NEW market out of an existing one, taking its
// locales, its traded currencies and its tax classes with it in a single
// call. That is the difference between this and POST /markets: a plain create
// leaves a row that cannot serve anybody, while what comes back here is a
// market with a language to render in, a currency to price in and a rate to
// tax with. The path id is the SOURCE market, resolved by uuid OR by market
// code.
func (srv *Markets) MarketsClone(Id string, Code string, optionalSetters ...MarketsCloneOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}/clone")
	options := MarketsCloneOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	params["code"] = Code
	if options.enabledSetters["CopyCurrencies"] {
		params["copy_currencies"] = options.CopyCurrencies
	}
	if options.enabledSetters["CopyLocales"] {
		params["copy_locales"] = options.CopyLocales
	}
	if options.enabledSetters["CopyTaxClasses"] {
		params["copy_tax_classes"] = options.CopyTaxClasses
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
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
	
// MarketsContext the storefront bootstrap: everything a frontend needs to
// render one market, resolved server-side so no client re-derives it — the
// market row, its locales, the currencies it trades in and its tax classes;
// WHICH locale to actually render in and where that answer came from; which
// key to read and write a translation under; whether the prices it will be
// handed are gross or net; and whether any of it is trustworthy. One call
// rather than five, and — more to the point — one place the resolution
// rules live, instead of a slightly different copy of them in every
// storefront. This one resolves the market by id only: unlike /readiness,
// /clone and /backfill, a market CODE here is a 400, not a lookup.
func (srv *Markets) MarketsContext(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}/context")
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
			
// MarketsMakeDefault a tenant has ONE default market: it is what every call
// naming none falls back to. Moving the flag from a client was
// promote-then-demote, two PATCHes that leave two defaults when the second
// does not land and none when the first does. This is the one call instead
// — it promotes the market in the path and demotes whoever held the flag in
// the same operation, writing once per row that was actually wrong and not
// touching the rest. Accepts an id or a market CODE. Answers the market plus
// the codes it demoted; repeating the call writes nothing.
func (srv *Markets) MarketsMakeDefault(Id string, Data interface{})(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}/make-default")
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
	
// MarketsReadiness whether this market can actually trade, and if not, what
// is missing. Every check runs on every call and comes back with its own
// severity, so the answer is a diagnosis rather than a yes or a no: a market
// with no currency registered has nothing to price in and a market with no
// tax class has nothing to tax with, and both of those fail BLOCKING, which
// is what turns `ready` false. A check that is merely degraded — no locale
// of its own, while the tenant declares a fallback_locale that covers for it
// — fails as a warning and leaves the market serviceable. Resolves the
// market by uuid OR by market code.
func (srv *Markets) MarketsReadiness(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/markets/{id}/readiness")
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
type MarketsCurrenciesListOptions struct {
	Id string
	Code string
	IsDefault bool
	Position int
	CreatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options MarketsCurrenciesListOptions) New() *MarketsCurrenciesListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Code": false,
		"IsDefault": false,
		"Position": false,
		"CreatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type MarketsCurrenciesListOption func(*MarketsCurrenciesListOptions)
func (srv *Markets) WithMarketsCurrenciesListId(v string) MarketsCurrenciesListOption {
	return func(o *MarketsCurrenciesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesListCode(v string) MarketsCurrenciesListOption {
	return func(o *MarketsCurrenciesListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesListIsDefault(v bool) MarketsCurrenciesListOption {
	return func(o *MarketsCurrenciesListOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesListPosition(v int) MarketsCurrenciesListOption {
	return func(o *MarketsCurrenciesListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesListCreatedAt(v string) MarketsCurrenciesListOption {
	return func(o *MarketsCurrenciesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesListLimit(v int) MarketsCurrenciesListOption {
	return func(o *MarketsCurrenciesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesListOffset(v int) MarketsCurrenciesListOption {
	return func(o *MarketsCurrenciesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesListOrder(v string) MarketsCurrenciesListOption {
	return func(o *MarketsCurrenciesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
			
// MarketsCurrenciesList every column is an exact-match filter and they
// combine with AND (?code=EUR); each one is declared as a query parameter
// above. A `?column=value` this entity does not have is DROPPED rather than
// refused — the call answers 200 with the unfiltered list — and `filter`
// echoes what was actually applied, which is the only way to tell that apart
// from a filter that matched nothing. `market_id` is not among them: the
// owning market comes from the path and overwrites anything the query says.
// An unknown but well-formed market lists empty rather than 404 — the
// parent is filtered on, not verified.
func (srv *Markets) MarketsCurrenciesList(MarketId string, optionalSetters ...MarketsCurrenciesListOption)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/currencies")
	options := MarketsCurrenciesListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
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
type MarketsCurrenciesCreateOptions struct {
	IsDefault bool
	Position int
	enabledSetters map[string]bool
}
func (options MarketsCurrenciesCreateOptions) New() *MarketsCurrenciesCreateOptions {
	options.enabledSetters = map[string]bool{
		"IsDefault": false,
		"Position": false,
	}
	return &options
}
type MarketsCurrenciesCreateOption func(*MarketsCurrenciesCreateOptions)
func (srv *Markets) WithMarketsCurrenciesCreateIsDefault(v bool) MarketsCurrenciesCreateOption {
	return func(o *MarketsCurrenciesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesCreatePosition(v int) MarketsCurrenciesCreateOption {
	return func(o *MarketsCurrenciesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
					
// MarketsCurrenciesCreate the owning market comes from the path and overrides
// anything in the body.
func (srv *Markets) MarketsCurrenciesCreate(MarketId string, Code string, optionalSetters ...MarketsCurrenciesCreateOption)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/currencies")
	options := MarketsCurrenciesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	params["code"] = Code
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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
			
// MarketsCurrenciesDelete scoped to the market in the path — a row
// belonging to another market is a 404 here, and is never deleted.
func (srv *Markets) MarketsCurrenciesDelete(MarketId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/currencies/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
			
// MarketsCurrenciesGet scoped strictly to the market in the path: a row
// belonging to another market is a 404 here, never a 200.
func (srv *Markets) MarketsCurrenciesGet(MarketId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/currencies/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
type MarketsCurrenciesUpdateOptions struct {
	Code string
	IsDefault bool
	Position int
	enabledSetters map[string]bool
}
func (options MarketsCurrenciesUpdateOptions) New() *MarketsCurrenciesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"IsDefault": false,
		"Position": false,
	}
	return &options
}
type MarketsCurrenciesUpdateOption func(*MarketsCurrenciesUpdateOptions)
func (srv *Markets) WithMarketsCurrenciesUpdateCode(v string) MarketsCurrenciesUpdateOption {
	return func(o *MarketsCurrenciesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesUpdateIsDefault(v bool) MarketsCurrenciesUpdateOption {
	return func(o *MarketsCurrenciesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsCurrenciesUpdatePosition(v int) MarketsCurrenciesUpdateOption {
	return func(o *MarketsCurrenciesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
					
// MarketsCurrenciesUpdate partial: omitted fields keep their value.
func (srv *Markets) MarketsCurrenciesUpdate(MarketId string, Id string, optionalSetters ...MarketsCurrenciesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/currencies/{id}")
	options := MarketsCurrenciesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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
type MarketsLocalesListOptions struct {
	Id string
	Code string
	Language string
	Country string
	IsDefault bool
	Position int
	CreatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options MarketsLocalesListOptions) New() *MarketsLocalesListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Code": false,
		"Language": false,
		"Country": false,
		"IsDefault": false,
		"Position": false,
		"CreatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type MarketsLocalesListOption func(*MarketsLocalesListOptions)
func (srv *Markets) WithMarketsLocalesListId(v string) MarketsLocalesListOption {
	return func(o *MarketsLocalesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Markets) WithMarketsLocalesListCode(v string) MarketsLocalesListOption {
	return func(o *MarketsLocalesListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsLocalesListLanguage(v string) MarketsLocalesListOption {
	return func(o *MarketsLocalesListOptions) {
		o.Language = v
		o.enabledSetters["Language"] = true
	}
}
func (srv *Markets) WithMarketsLocalesListCountry(v string) MarketsLocalesListOption {
	return func(o *MarketsLocalesListOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *Markets) WithMarketsLocalesListIsDefault(v bool) MarketsLocalesListOption {
	return func(o *MarketsLocalesListOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsLocalesListPosition(v int) MarketsLocalesListOption {
	return func(o *MarketsLocalesListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsLocalesListCreatedAt(v string) MarketsLocalesListOption {
	return func(o *MarketsLocalesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Markets) WithMarketsLocalesListLimit(v int) MarketsLocalesListOption {
	return func(o *MarketsLocalesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Markets) WithMarketsLocalesListOffset(v int) MarketsLocalesListOption {
	return func(o *MarketsLocalesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Markets) WithMarketsLocalesListOrder(v string) MarketsLocalesListOption {
	return func(o *MarketsLocalesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
			
// MarketsLocalesList every column is an exact-match filter and they combine
// with AND (?code=de-DE); each one is declared as a query parameter above. A
// `?column=value` this entity does not have is DROPPED rather than refused
// — the call answers 200 with the unfiltered list — and `filter` echoes
// what was actually applied, which is the only way to tell that apart from a
// filter that matched nothing. `market_id` is not among them: the owning
// market comes from the path and overwrites anything the query says. An
// unknown but well-formed market lists empty rather than 404 — the parent
// is filtered on, not verified.
func (srv *Markets) MarketsLocalesList(MarketId string, optionalSetters ...MarketsLocalesListOption)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/locales")
	options := MarketsLocalesListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Language"] {
		params["language"] = options.Language
	}
	if options.enabledSetters["Country"] {
		params["country"] = options.Country
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
type MarketsLocalesCreateOptions struct {
	IsDefault bool
	Position int
	enabledSetters map[string]bool
}
func (options MarketsLocalesCreateOptions) New() *MarketsLocalesCreateOptions {
	options.enabledSetters = map[string]bool{
		"IsDefault": false,
		"Position": false,
	}
	return &options
}
type MarketsLocalesCreateOption func(*MarketsLocalesCreateOptions)
func (srv *Markets) WithMarketsLocalesCreateIsDefault(v bool) MarketsLocalesCreateOption {
	return func(o *MarketsLocalesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsLocalesCreatePosition(v int) MarketsLocalesCreateOption {
	return func(o *MarketsLocalesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
									
// MarketsLocalesCreate the owning market comes from the path and overrides
// anything in the body.
func (srv *Markets) MarketsLocalesCreate(MarketId string, Code string, Country string, Language string, optionalSetters ...MarketsLocalesCreateOption)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/locales")
	options := MarketsLocalesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	params["code"] = Code
	params["country"] = Country
	params["language"] = Language
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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
			
// MarketsLocalesDelete scoped to the market in the path — a row belonging
// to another market is a 404 here, and is never deleted.
func (srv *Markets) MarketsLocalesDelete(MarketId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/locales/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
			
// MarketsLocalesGet scoped strictly to the market in the path: a row
// belonging to another market is a 404 here, never a 200.
func (srv *Markets) MarketsLocalesGet(MarketId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/locales/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
type MarketsLocalesUpdateOptions struct {
	Code string
	Country string
	IsDefault bool
	Language string
	Position int
	enabledSetters map[string]bool
}
func (options MarketsLocalesUpdateOptions) New() *MarketsLocalesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Country": false,
		"IsDefault": false,
		"Language": false,
		"Position": false,
	}
	return &options
}
type MarketsLocalesUpdateOption func(*MarketsLocalesUpdateOptions)
func (srv *Markets) WithMarketsLocalesUpdateCode(v string) MarketsLocalesUpdateOption {
	return func(o *MarketsLocalesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsLocalesUpdateCountry(v string) MarketsLocalesUpdateOption {
	return func(o *MarketsLocalesUpdateOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *Markets) WithMarketsLocalesUpdateIsDefault(v bool) MarketsLocalesUpdateOption {
	return func(o *MarketsLocalesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsLocalesUpdateLanguage(v string) MarketsLocalesUpdateOption {
	return func(o *MarketsLocalesUpdateOptions) {
		o.Language = v
		o.enabledSetters["Language"] = true
	}
}
func (srv *Markets) WithMarketsLocalesUpdatePosition(v int) MarketsLocalesUpdateOption {
	return func(o *MarketsLocalesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
					
// MarketsLocalesUpdate partial: omitted fields keep their value.
func (srv *Markets) MarketsLocalesUpdate(MarketId string, Id string, optionalSetters ...MarketsLocalesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/locales/{id}")
	options := MarketsLocalesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Country"] {
		params["country"] = options.Country
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Language"] {
		params["language"] = options.Language
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
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
type MarketsTaxClassesListOptions struct {
	Id string
	Code string
	Name string
	Labels string
	Rate float64
	IsDefault bool
	Position int
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options MarketsTaxClassesListOptions) New() *MarketsTaxClassesListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Code": false,
		"Name": false,
		"Labels": false,
		"Rate": false,
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
type MarketsTaxClassesListOption func(*MarketsTaxClassesListOptions)
func (srv *Markets) WithMarketsTaxClassesListId(v string) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesListCode(v string) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesListName(v string) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesListLabels(v string) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesListRate(v float64) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.Rate = v
		o.enabledSetters["Rate"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesListIsDefault(v bool) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesListPosition(v int) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesListCreatedAt(v string) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesListUpdatedAt(v string) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesListLimit(v int) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesListOffset(v int) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesListOrder(v string) MarketsTaxClassesListOption {
	return func(o *MarketsTaxClassesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
			
// MarketsTaxClassesList every column is an exact-match filter and they
// combine with AND (?code=standard); each one is declared as a query
// parameter above. A `?column=value` this entity does not have is DROPPED
// rather than refused — the call answers 200 with the unfiltered list —
// and `filter` echoes what was actually applied, which is the only way to
// tell that apart from a filter that matched nothing. `market_id` is not
// among them: the owning market comes from the path and overwrites anything
// the query says. An unknown but well-formed market lists empty rather than
// 404 — the parent is filtered on, not verified.
func (srv *Markets) MarketsTaxClassesList(MarketId string, optionalSetters ...MarketsTaxClassesListOption)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/tax_classes")
	options := MarketsTaxClassesListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
	if options.enabledSetters["Rate"] {
		params["rate"] = options.Rate
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
type MarketsTaxClassesCreateOptions struct {
	IsDefault bool
	Labels interface{}
	Position int
	Rate float64
	enabledSetters map[string]bool
}
func (options MarketsTaxClassesCreateOptions) New() *MarketsTaxClassesCreateOptions {
	options.enabledSetters = map[string]bool{
		"IsDefault": false,
		"Labels": false,
		"Position": false,
		"Rate": false,
	}
	return &options
}
type MarketsTaxClassesCreateOption func(*MarketsTaxClassesCreateOptions)
func (srv *Markets) WithMarketsTaxClassesCreateIsDefault(v bool) MarketsTaxClassesCreateOption {
	return func(o *MarketsTaxClassesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesCreateLabels(v interface{}) MarketsTaxClassesCreateOption {
	return func(o *MarketsTaxClassesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesCreatePosition(v int) MarketsTaxClassesCreateOption {
	return func(o *MarketsTaxClassesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesCreateRate(v float64) MarketsTaxClassesCreateOption {
	return func(o *MarketsTaxClassesCreateOptions) {
		o.Rate = v
		o.enabledSetters["Rate"] = true
	}
}
							
// MarketsTaxClassesCreate the owning market comes from the path and overrides
// anything in the body.
func (srv *Markets) MarketsTaxClassesCreate(MarketId string, Code string, Name string, optionalSetters ...MarketsTaxClassesCreateOption)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId)
	path := r.Replace("/v1/markets/{market_id}/tax_classes")
	options := MarketsTaxClassesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
	if options.enabledSetters["Rate"] {
		params["rate"] = options.Rate
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
			
// MarketsTaxClassesDelete refused with a 409 for as long as another app still
// points at this tax class by its code. A tax class is the source of record
// for a rate, and other apps name it by CODE with no foreign key behind it
// — a cross-app FK is what ADR-0055 forbids. So this asks the shipping app
// what still uses the code (shipping.tax-classes.usage) and answers 409 with
// the count and the first few names rather than leaving methods quoting a
// rate nobody defines. The check FAILS OPEN: a tenant without the shipping
// app, or an unreachable one, deletes as before, and the answer says which
// happened in 'usage_checked'. Matched on the code, which is shared across
// markets — the refusal message says so.
func (srv *Markets) MarketsTaxClassesDelete(MarketId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/tax_classes/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
			
// MarketsTaxClassesGet scoped strictly to the market in the path: a row
// belonging to another market is a 404 here, never a 200.
func (srv *Markets) MarketsTaxClassesGet(MarketId string, Id string)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/tax_classes/{id}")
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
type MarketsTaxClassesUpdateOptions struct {
	Code string
	IsDefault bool
	Labels interface{}
	Name string
	Position int
	Rate float64
	enabledSetters map[string]bool
}
func (options MarketsTaxClassesUpdateOptions) New() *MarketsTaxClassesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"IsDefault": false,
		"Labels": false,
		"Name": false,
		"Position": false,
		"Rate": false,
	}
	return &options
}
type MarketsTaxClassesUpdateOption func(*MarketsTaxClassesUpdateOptions)
func (srv *Markets) WithMarketsTaxClassesUpdateCode(v string) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesUpdateIsDefault(v bool) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesUpdateLabels(v interface{}) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesUpdateName(v string) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesUpdatePosition(v int) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *Markets) WithMarketsTaxClassesUpdateRate(v float64) MarketsTaxClassesUpdateOption {
	return func(o *MarketsTaxClassesUpdateOptions) {
		o.Rate = v
		o.enabledSetters["Rate"] = true
	}
}
					
// MarketsTaxClassesUpdate partial: omitted fields keep their value.
func (srv *Markets) MarketsTaxClassesUpdate(MarketId string, Id string, optionalSetters ...MarketsTaxClassesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{market_id}", MarketId, "{id}", Id)
	path := r.Replace("/v1/markets/{market_id}/tax_classes/{id}")
	options := MarketsTaxClassesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["market_id"] = MarketId
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
	if options.enabledSetters["Rate"] {
		params["rate"] = options.Rate
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
