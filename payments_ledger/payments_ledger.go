package payments_ledger

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// PaymentsLedger service
type PaymentsLedger struct {
	client client.Client
}

func New(clt client.Client) *PaymentsLedger {
	return &PaymentsLedger{
		client: clt,
	}
}

type PaymentsListOptions struct {
	Limit int
	Offset int
	Order string
	CartId string
	ContactId string
	Status string
	OrderRef string
	MethodCode string
	Kind string
	Provider string
	DunningStage string
	IdempotencyKey string
	enabledSetters map[string]bool
}
func (options PaymentsListOptions) New() *PaymentsListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"CartId": false,
		"ContactId": false,
		"Status": false,
		"OrderRef": false,
		"MethodCode": false,
		"Kind": false,
		"Provider": false,
		"DunningStage": false,
		"IdempotencyKey": false,
	}
	return &options
}
type PaymentsListOption func(*PaymentsListOptions)
func (srv *PaymentsLedger) WithPaymentsListLimit(v int) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsListOffset(v int) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsListOrder(v string) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsListCartId(v string) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.CartId = v
		o.enabledSetters["CartId"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsListContactId(v string) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsListStatus(v string) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsListOrderRef(v string) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.OrderRef = v
		o.enabledSetters["OrderRef"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsListMethodCode(v string) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.MethodCode = v
		o.enabledSetters["MethodCode"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsListKind(v string) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.Kind = v
		o.enabledSetters["Kind"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsListProvider(v string) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.Provider = v
		o.enabledSetters["Provider"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsListDunningStage(v string) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.DunningStage = v
		o.enabledSetters["DunningStage"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsListIdempotencyKey(v string) PaymentsListOption {
	return func(o *PaymentsListOptions) {
		o.IdempotencyKey = v
		o.enabledSetters["IdempotencyKey"] = true
	}
}
	
// PaymentsList the ledger, paged and filtered — the Payments screen, the
// reconciliation query and the way an order or a cart finds out what has been
// paid against it. Every column of the entity is an exact-match filter, which
// is what makes it useful: `?cart_id=` and `?contact_id=` are indexed,
// `?status=authorized&kind=self_managed` is the awaiting-payment queue the
// dunning scan classifies, and `?order_ref=` is the only way to resolve a
// payment by its external reference. Rows come back in the database's own
// order, so a newest-first list needs `?order=created_at.desc`.
// `error_message` is answered from the failure taxonomy rather than echoed
// out of the column, so what a driver or a PSP actually wrote is never
// serialized here.
func (srv *PaymentsLedger) PaymentsList(optionalSetters ...PaymentsListOption)(*interface{}, error) {
	path := "/v1/payments"
	options := PaymentsListOptions{}.New()
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
	if options.enabledSetters["CartId"] {
		params["cart_id"] = options.CartId
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["OrderRef"] {
		params["order_ref"] = options.OrderRef
	}
	if options.enabledSetters["MethodCode"] {
		params["method_code"] = options.MethodCode
	}
	if options.enabledSetters["Kind"] {
		params["kind"] = options.Kind
	}
	if options.enabledSetters["Provider"] {
		params["provider"] = options.Provider
	}
	if options.enabledSetters["DunningStage"] {
		params["dunning_stage"] = options.DunningStage
	}
	if options.enabledSetters["IdempotencyKey"] {
		params["idempotency_key"] = options.IdempotencyKey
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
type PaymentsCreateOptions struct {
	CartId string
	ContactId string
	Country string
	Currency string
	IdempotencyKey string
	Metadata interface{}
	OrderRef string
	ReturnUrl string
	enabledSetters map[string]bool
}
func (options PaymentsCreateOptions) New() *PaymentsCreateOptions {
	options.enabledSetters = map[string]bool{
		"CartId": false,
		"ContactId": false,
		"Country": false,
		"Currency": false,
		"IdempotencyKey": false,
		"Metadata": false,
		"OrderRef": false,
		"ReturnUrl": false,
	}
	return &options
}
type PaymentsCreateOption func(*PaymentsCreateOptions)
func (srv *PaymentsLedger) WithPaymentsCreateCartId(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.CartId = v
		o.enabledSetters["CartId"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsCreateContactId(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.ContactId = v
		o.enabledSetters["ContactId"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsCreateCountry(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.Country = v
		o.enabledSetters["Country"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsCreateCurrency(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.Currency = v
		o.enabledSetters["Currency"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsCreateIdempotencyKey(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.IdempotencyKey = v
		o.enabledSetters["IdempotencyKey"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsCreateMetadata(v interface{}) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.Metadata = v
		o.enabledSetters["Metadata"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsCreateOrderRef(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.OrderRef = v
		o.enabledSetters["OrderRef"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsCreateReturnUrl(v string) PaymentsCreateOption {
	return func(o *PaymentsCreateOptions) {
		o.ReturnUrl = v
		o.enabledSetters["ReturnUrl"] = true
	}
}
					
// PaymentsCreate the checkout's write: it opens the ledger row and takes it
// as far as the named method allows, in one call. A create cannot omit
// `method_code` and `amount`; every other column is optional or defaulted by
// the database. Nothing else about the money is the caller's to choose:
// `kind`, `provider` and `fee_amount` are read off the method that
// `method_code` names, so a caller can neither pick an acquirer nor discount
// its own fee. `amount: 0` is legal (free orders); negative is 400.
// Eligibility is enforced HERE and not only in the checkout UI — the same
// country and order-value rules POST /payments/methods/eligible applies
// answer 422 if the method does not apply to this buyer. What comes back
// depends on the method: a self-managed one (invoice, prepayment) is
// `authorized` at once with the dunning clock already started, and a PSP one
// is `captured` or `authorized`, or `requires_action` with `next_action` —
// the instruction the storefront must carry out, typically a redirect, set at
// that status and at no other. Send an `idempotency_key` and a repeat of the
// same call answers 200 with the payment that key already named, unchanged
// and not re-authorized. What is never stored: the `instrument`, `token` or
// `card` is handed to the driver in-process and no token or PAN is written to
// the row.
func (srv *PaymentsLedger) PaymentsCreate(Amount float64, MethodCode string, optionalSetters ...PaymentsCreateOption)(*models.Error, error) {
	path := "/v1/payments"
	options := PaymentsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["amount"] = Amount
	params["method_code"] = MethodCode
	if options.enabledSetters["CartId"] {
		params["cart_id"] = options.CartId
	}
	if options.enabledSetters["ContactId"] {
		params["contact_id"] = options.ContactId
	}
	if options.enabledSetters["Country"] {
		params["country"] = options.Country
	}
	if options.enabledSetters["Currency"] {
		params["currency"] = options.Currency
	}
	if options.enabledSetters["IdempotencyKey"] {
		params["idempotency_key"] = options.IdempotencyKey
	}
	if options.enabledSetters["Metadata"] {
		params["metadata"] = options.Metadata
	}
	if options.enabledSetters["OrderRef"] {
		params["order_ref"] = options.OrderRef
	}
	if options.enabledSetters["ReturnUrl"] {
		params["return_url"] = options.ReturnUrl
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

// PaymentsDunningScan classifies every unpaid self-managed payment (invoice,
// prepayment) as on time / reminder due / overdue from
// payment_reminder_after_days and overdue_after_days, writes the stage and
// the next due date, and reports PSP payments still waiting on a callback
// longer than webhook_stale_after_minutes. Pure function of each payment's
// age, so it is idempotent — it also runs daily as the 'dunning-scan'
// schedule. It classifies and does not send: a stage change emits
// payment.updated, and what a reminder looks like is the merchant's workflow.
func (srv *PaymentsLedger) PaymentsDunningScan()(*interface{}, error) {
	path := "/v1/payments/dunning/scan"
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
type PaymentsErrorsRedactOptions struct {
	Apply bool
	Limit int
	enabledSetters map[string]bool
}
func (options PaymentsErrorsRedactOptions) New() *PaymentsErrorsRedactOptions {
	options.enabledSetters = map[string]bool{
		"Apply": false,
		"Limit": false,
	}
	return &options
}
type PaymentsErrorsRedactOption func(*PaymentsErrorsRedactOptions)
func (srv *PaymentsLedger) WithPaymentsErrorsRedactApply(v bool) PaymentsErrorsRedactOption {
	return func(o *PaymentsErrorsRedactOptions) {
		o.Apply = v
		o.enabledSetters["Apply"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsErrorsRedactLimit(v int) PaymentsErrorsRedactOption {
	return func(o *PaymentsErrorsRedactOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
	
// PaymentsErrorsRedact rows written before the failure taxonomy still store
// the provider's/runtime's raw text in error_message. API responses never
// repeat it (the read path projects), but the column is also read directly
// through Baseline, so it needs rewriting once per tenant. Dry-run by default
// — reports what it would touch and changes nothing until apply:true.
// Idempotent: rows already carrying a taxonomy message are skipped.
func (srv *PaymentsLedger) PaymentsErrorsRedact(optionalSetters ...PaymentsErrorsRedactOption)(*interface{}, error) {
	path := "/v1/payments/errors/redact"
	options := PaymentsErrorsRedactOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Apply"] {
		params["apply"] = options.Apply
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
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
	
// PaymentsOrdersCapture this is the hook the tenant's `auto_capture_policy:
// 'on_ship'` was written for: fulfilment knows the order it shipped and not
// the payment ids behind it, so the shipment calls this one route with the
// reference it already holds and the money for that order is collected in a
// single request. Resolves payments by their order_ref (the same key the PSP
// webhooks fall back to), captures every authorized one and reports the rest
// instead of failing — an order whose payment was already captured is a
// successful no-op, and a provider that refuses one payment lands in
// `skipped` rather than failing the call. Note that payments.order_ref is
// nullable with no foreign key: this route is exactly as good as the
// reference the checkout writes onto the payment.
func (srv *PaymentsLedger) PaymentsOrdersCapture(OrderRef string)(*models.Error, error) {
	r := strings.NewReplacer("{order_ref}", OrderRef)
	path := r.Replace("/v1/payments/orders/{order_ref}/capture")
	params := map[string]interface{}{}
	params["order_ref"] = OrderRef
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

// PaymentsVocabulariesList the enums this app owns, four of them: statuses,
// method kinds, fee types and dunning stages. This is the index and carries a
// name and a title per set and nothing more — the values themselves, with
// their labels and badge tones, are one call further down at GET
// /payments/vocabularies/{name}, so a client that only needs to know which
// sets exist does not pay for all of them. Values come out of the CHECK
// constraints, so what is served is what the database enforces — a client
// renders a status this app adds without a release of its own.
func (srv *PaymentsLedger) PaymentsVocabulariesList()(*interface{}, error) {
	path := "/v1/payments/vocabularies"
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
	
// PaymentsVocabulariesGet one set in full: every value it permits, the label
// to show for each and the badge tone to render it in, which is what a client
// needs to draw a status chip without hard-coding this app's enums. The value
// set is parsed out of the CHECK constraint in schema.json, so what is served
// IS what the database enforces. Labels are curated on top and can only add
// words and colour — a permitted value nobody labelled still appears,
// titled from its own key, which is why `title` and `description` are a
// locale map on a labelled value and a plain string on an unlabelled one.
func (srv *PaymentsLedger) PaymentsVocabulariesGet(Name string)(*models.Error, error) {
	r := strings.NewReplacer("{name}", Name)
	path := r.Replace("/v1/payments/vocabularies/{name}")
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
type PaymentsWebhooksIngestOptions struct {
	Id interface{}
	Request interface{}
	Verified interface{}
	enabledSetters map[string]bool
}
func (options PaymentsWebhooksIngestOptions) New() *PaymentsWebhooksIngestOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Request": false,
		"Verified": false,
	}
	return &options
}
type PaymentsWebhooksIngestOption func(*PaymentsWebhooksIngestOptions)
func (srv *PaymentsLedger) WithPaymentsWebhooksIngestId(v interface{}) PaymentsWebhooksIngestOption {
	return func(o *PaymentsWebhooksIngestOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsWebhooksIngestRequest(v interface{}) PaymentsWebhooksIngestOption {
	return func(o *PaymentsWebhooksIngestOptions) {
		o.Request = v
		o.enabledSetters["Request"] = true
	}
}
func (srv *PaymentsLedger) WithPaymentsWebhooksIngestVerified(v interface{}) PaymentsWebhooksIngestOption {
	return func(o *PaymentsWebhooksIngestOptions) {
		o.Verified = v
		o.enabledSetters["Verified"] = true
	}
}
			
// PaymentsWebhooksIngest the sink a PSP callback ends up in, and an inbound
// ingress endpoint in the sense of ADR-0066: the provider never posts here
// directly, it posts to webhooks.revenexx.com, which verifies and captures
// the delivery and dispatches its envelope to this route through the gateway.
// That indirection is also what makes this the one override point for PSP
// callback handling — everything a callback does to the ledger happens here
// and nowhere else, so a deployment that needs a provider's callbacks
// normalized differently replaces this operation instead of touching the
// lifecycle routes. Consumes the dispatch envelope from
// webhooks.revenexx.com: normalizes the provider callback (stripe payment
// intents + a generic shape), resolves the payment by psp_payment_id or
// order_ref and moves the ledger. Facts only move forward — provider
// retries and redeliveries are idempotent no-ops; unverified envelopes are
// refused.
func (srv *PaymentsLedger) PaymentsWebhooksIngest(Provider string, optionalSetters ...PaymentsWebhooksIngestOption)(*models.Error, error) {
	r := strings.NewReplacer("{provider}", Provider)
	path := r.Replace("/v1/payments/webhooks/{provider}")
	options := PaymentsWebhooksIngestOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["provider"] = Provider
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["Request"] {
		params["request"] = options.Request
	}
	if options.enabledSetters["Verified"] {
		params["verified"] = options.Verified
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
	
// PaymentsGet one ledger row in full: the amount and the fee that were
// computed at creation, the method code and PSP it was made through, where it
// stands in the lifecycle, the timestamp of each transition it has been
// through (`authorized_at`, `captured_at`, `failed_at`, `refunded_at`), the
// dunning columns the daily scan maintains and, while the buyer still has
// something to do, `next_action`. This is the call to poll after sending a
// buyer to a PSP redirect. Two things it does not do: `error_message` is
// answered from the failure taxonomy and never carries the provider's or the
// runtime's own words, and there is no route that resolves a payment by
// `order_ref` — that column is nullable and not unique, so it is a filter
// on the list (`GET /payments?order_ref=…`) which may legitimately answer
// several rows.
func (srv *PaymentsLedger) PaymentsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/{id}")
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
type PaymentsCancelOptions struct {
	Reason string
	enabledSetters map[string]bool
}
func (options PaymentsCancelOptions) New() *PaymentsCancelOptions {
	options.enabledSetters = map[string]bool{
		"Reason": false,
	}
	return &options
}
type PaymentsCancelOption func(*PaymentsCancelOptions)
func (srv *PaymentsLedger) WithPaymentsCancelReason(v string) PaymentsCancelOption {
	return func(o *PaymentsCancelOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
			
// PaymentsCancel drops the claim before any money has been taken — the
// abandoned basket, the buyer who never came back from the redirect, the
// invoice an operator writes off. It is the only transition that starts from
// three statuses rather than one, because everything short of captured can
// still be released. A captured payment is not cancellable at all: that is a
// refund, and the lattice answers 400 rather than pretending. Unlike capture
// and refund this transition has no time window — the merchant's
// `capture_expiry_days` and `refund_window_days` do not apply, so a stale
// authorization can always be released even once it is too old to collect. On
// a PSP payment the provider is called and the `reason` in the body is passed
// to it, so it reaches the PSP's own cancellation-reason field as well as
// being stored under `metadata.cancel_reason`. Cancelling stops the dunning
// clock: the stage goes back to `none` and the due date is cleared.
func (srv *PaymentsLedger) PaymentsCancel(Id string, optionalSetters ...PaymentsCancelOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/{id}/cancel")
	options := PaymentsCancelOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
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
	
// PaymentsCapture collects money that is currently only reserved. It starts
// from `authorized` and from nothing else — under `auto_capture_policy:
// 'immediate'` a payment is captured in the same request that created it and
// never passes through here, so this is the route for the 'manual' and
// 'on_ship' policies, and POST /payments/orders/{order_ref}/capture is the
// same operation addressed by the order reference a warehouse actually holds.
// There is no request body and no amount: the ledger carries one amount and
// one status, so a capture is the whole authorization or nothing. On a
// self-managed payment it takes no PSP anywhere near it — it records that
// an invoice or a prepayment was paid, and stops the dunning clock. Refused
// with 422 once the authorization is older than the tenant's
// `capture_expiry_days` (the message carries both numbers), because an
// expired authorization is declined by the provider anyway and a 422 here is
// the cheap version of finding out later.
func (srv *PaymentsLedger) PaymentsCapture(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/{id}/capture")
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
	
// PaymentsConfirm the other half of a redirect. POST /payments answered
// `requires_action` with a `next_action` the storefront carried out — a 3-D
// Secure step, a wallet approval, a bank login — and this is the call that
// asks the PSP how it went and writes the answer to the ledger. It starts
// from `requires_action` and from nothing else, so a payment that already
// came back authorized needs no confirm and the lattice answers 400 rather
// than repeating one. `next_action` is cleared by this call whatever the
// outcome. Where the tenant's `auto_capture_policy` is 'immediate' the money
// is taken straight after the authorization, in the same request, so a
// successful confirm can come back `captured` rather than `authorized`; a
// failed auto-capture does not fail the confirm, because a good authorization
// is worth more than a tidy status.
func (srv *PaymentsLedger) PaymentsConfirm(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/{id}/confirm")
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
type PaymentsRefundOptions struct {
	Reason string
	enabledSetters map[string]bool
}
func (options PaymentsRefundOptions) New() *PaymentsRefundOptions {
	options.enabledSetters = map[string]bool{
		"Reason": false,
	}
	return &options
}
type PaymentsRefundOption func(*PaymentsRefundOptions)
func (srv *PaymentsLedger) WithPaymentsRefundReason(v string) PaymentsRefundOption {
	return func(o *PaymentsRefundOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
			
// PaymentsRefund gives captured money back. It starts from `captured` and
// from nothing else — money that was only authorized is cancelled, not
// refunded, and the lattice answers 400 rather than guessing which was meant.
// All or nothing: the ledger carries one amount and one status, so there is
// no partial refund and no second one to express — a refunded payment is
// refunded in full, and a repeat is a 400 because `refunded` is not a status
// a refund starts from. The `reason` in the body is handed to the driver in
// the same call, so it reaches the PSP's own refund-reason field rather than
// being a note only this database ever sees, and it is stored under
// `metadata.refund_reason`. On a self-managed payment nothing is sent
// anywhere: it records that the merchant paid the buyer back by their own
// means. Refused with 422 once the capture is older than the tenant's
// `refund_window_days` (the message carries both numbers) — past that the
// provider stops accepting a refund against the transaction and it has to be
// made by bank transfer.
func (srv *PaymentsLedger) PaymentsRefund(Id string, optionalSetters ...PaymentsRefundOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/payments/{id}/refund")
	options := PaymentsRefundOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
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
