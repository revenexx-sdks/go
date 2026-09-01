package messaging

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Messaging service
type Messaging struct {
	client client.Client
}

func New(clt client.Client) *Messaging {
	return &Messaging{
		client: clt,
	}
}

type AuditIndexOptions struct {
	ResourceType string
	ResourceId string
	Subject string
	Limit int
	enabledSetters map[string]bool
}
func (options AuditIndexOptions) New() *AuditIndexOptions {
	options.enabledSetters = map[string]bool{
		"ResourceType": false,
		"ResourceId": false,
		"Subject": false,
		"Limit": false,
	}
	return &options
}
type AuditIndexOption func(*AuditIndexOptions)
func (srv *Messaging) WithAuditIndexResourceType(v string) AuditIndexOption {
	return func(o *AuditIndexOptions) {
		o.ResourceType = v
		o.enabledSetters["ResourceType"] = true
	}
}
func (srv *Messaging) WithAuditIndexResourceId(v string) AuditIndexOption {
	return func(o *AuditIndexOptions) {
		o.ResourceId = v
		o.enabledSetters["ResourceId"] = true
	}
}
func (srv *Messaging) WithAuditIndexSubject(v string) AuditIndexOption {
	return func(o *AuditIndexOptions) {
		o.Subject = v
		o.enabledSetters["Subject"] = true
	}
}
func (srv *Messaging) WithAuditIndexLimit(v int) AuditIndexOption {
	return func(o *AuditIndexOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
	
// AuditIndex filterable by `resource_type`, `resource_id` and `subject` —
// the last one
// being the human-readable name a row was recorded under (a template's key,
// a layout's name), which is what an operator has to hand six weeks later
// when the id means nothing to them.
// 
// There is no write route and no delete route: an append-only log with an
// editor is a log that says whatever the last editor wanted.
func (srv *Messaging) AuditIndex(optionalSetters ...AuditIndexOption)(*models.Error, error) {
	path := "/v1/messaging/audit"
	options := AuditIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["ResourceType"] {
		params["resource_type"] = options.ResourceType
	}
	if options.enabledSetters["ResourceId"] {
		params["resource_id"] = options.ResourceId
	}
	if options.enabledSetters["Subject"] {
		params["subject"] = options.Subject
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
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
type BindingIndexOptions struct {
	EventTopic string
	enabledSetters map[string]bool
}
func (options BindingIndexOptions) New() *BindingIndexOptions {
	options.enabledSetters = map[string]bool{
		"EventTopic": false,
	}
	return &options
}
type BindingIndexOption func(*BindingIndexOptions)
func (srv *Messaging) WithBindingIndexEventTopic(v string) BindingIndexOption {
	return func(o *BindingIndexOptions) {
		o.EventTopic = v
		o.enabledSetters["EventTopic"] = true
	}
}
	
// BindingIndex `?event_topic=` narrows to one topic, which is the question
// worth asking
// of this list: "what does this event actually do".
func (srv *Messaging) BindingIndex(optionalSetters ...BindingIndexOption)(*models.Error, error) {
	path := "/v1/messaging/bindings"
	options := BindingIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["EventTopic"] {
		params["event_topic"] = options.EventTopic
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
type BindingStoreOptions struct {
	Enabled bool
	FallbackOrder int
	Locale string
	enabledSetters map[string]bool
}
func (options BindingStoreOptions) New() *BindingStoreOptions {
	options.enabledSetters = map[string]bool{
		"Enabled": false,
		"FallbackOrder": false,
		"Locale": false,
	}
	return &options
}
type BindingStoreOption func(*BindingStoreOptions)
func (srv *Messaging) WithBindingStoreEnabled(v bool) BindingStoreOption {
	return func(o *BindingStoreOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithBindingStoreFallbackOrder(v int) BindingStoreOption {
	return func(o *BindingStoreOptions) {
		o.FallbackOrder = v
		o.enabledSetters["FallbackOrder"] = true
	}
}
func (srv *Messaging) WithBindingStoreLocale(v string) BindingStoreOption {
	return func(o *BindingStoreOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
									
// BindingStore `recipient` is a template, not an address: `{{ customer.email
// }}` is
// rendered against the event payload when the event arrives, which is the
// only way one binding can serve every customer. An event that renders it
// empty is skipped and logged rather than sent to nobody.
// 
// `locale` is what the OPERATOR said this route speaks, and it outranks the
// tenant's default. Leave it null when nobody has made that decision, so
// that the recipient's own language is still allowed to decide.
func (srv *Messaging) BindingStore(Channel string, EventTopic string, Recipient string, TemplateKey string, optionalSetters ...BindingStoreOption)(*models.Error, error) {
	path := "/v1/messaging/bindings"
	options := BindingStoreOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["channel"] = Channel
	params["event_topic"] = EventTopic
	params["recipient"] = Recipient
	params["template_key"] = TemplateKey
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FallbackOrder"] {
		params["fallback_order"] = options.FallbackOrder
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
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
	
// BindingDestroy the event it answered goes back to doing nothing. Prefer
// `enabled: false`
// when the intent is to pause rather than to forget.
func (srv *Messaging) BindingDestroy(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/bindings/{id}")
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
	
// BindingShow 404 for a binding belonging to another tenant, not 403 — an
// id that
// answered differently would say whether it exists.
func (srv *Messaging) BindingShow(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/bindings/{id}")
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
type BindingUpdatePatchOptions struct {
	Channel string
	Enabled bool
	EventTopic string
	FallbackOrder int
	Locale string
	Recipient string
	TemplateKey string
	enabledSetters map[string]bool
}
func (options BindingUpdatePatchOptions) New() *BindingUpdatePatchOptions {
	options.enabledSetters = map[string]bool{
		"Channel": false,
		"Enabled": false,
		"EventTopic": false,
		"FallbackOrder": false,
		"Locale": false,
		"Recipient": false,
		"TemplateKey": false,
	}
	return &options
}
type BindingUpdatePatchOption func(*BindingUpdatePatchOptions)
func (srv *Messaging) WithBindingUpdatePatchChannel(v string) BindingUpdatePatchOption {
	return func(o *BindingUpdatePatchOptions) {
		o.Channel = v
		o.enabledSetters["Channel"] = true
	}
}
func (srv *Messaging) WithBindingUpdatePatchEnabled(v bool) BindingUpdatePatchOption {
	return func(o *BindingUpdatePatchOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithBindingUpdatePatchEventTopic(v string) BindingUpdatePatchOption {
	return func(o *BindingUpdatePatchOptions) {
		o.EventTopic = v
		o.enabledSetters["EventTopic"] = true
	}
}
func (srv *Messaging) WithBindingUpdatePatchFallbackOrder(v int) BindingUpdatePatchOption {
	return func(o *BindingUpdatePatchOptions) {
		o.FallbackOrder = v
		o.enabledSetters["FallbackOrder"] = true
	}
}
func (srv *Messaging) WithBindingUpdatePatchLocale(v string) BindingUpdatePatchOption {
	return func(o *BindingUpdatePatchOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Messaging) WithBindingUpdatePatchRecipient(v string) BindingUpdatePatchOption {
	return func(o *BindingUpdatePatchOptions) {
		o.Recipient = v
		o.enabledSetters["Recipient"] = true
	}
}
func (srv *Messaging) WithBindingUpdatePatchTemplateKey(v string) BindingUpdatePatchOption {
	return func(o *BindingUpdatePatchOptions) {
		o.TemplateKey = v
		o.enabledSetters["TemplateKey"] = true
	}
}
			
// BindingUpdatePatch every field is optional; only what is sent is written.
// `enabled: false`
// is how a binding is taken out of service without losing what it said —
// the alternative is deleting it and typing the payload path back in
// correctly from memory later.
// 
// This path answers on `PUT` and `PATCH`, both routed to the same action.
func (srv *Messaging) BindingUpdatePatch(Id string, optionalSetters ...BindingUpdatePatchOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/bindings/{id}")
	options := BindingUpdatePatchOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Channel"] {
		params["channel"] = options.Channel
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["EventTopic"] {
		params["event_topic"] = options.EventTopic
	}
	if options.enabledSetters["FallbackOrder"] {
		params["fallback_order"] = options.FallbackOrder
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["Recipient"] {
		params["recipient"] = options.Recipient
	}
	if options.enabledSetters["TemplateKey"] {
		params["template_key"] = options.TemplateKey
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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
type BindingUpdateOptions struct {
	Channel string
	Enabled bool
	EventTopic string
	FallbackOrder int
	Locale string
	Recipient string
	TemplateKey string
	enabledSetters map[string]bool
}
func (options BindingUpdateOptions) New() *BindingUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Channel": false,
		"Enabled": false,
		"EventTopic": false,
		"FallbackOrder": false,
		"Locale": false,
		"Recipient": false,
		"TemplateKey": false,
	}
	return &options
}
type BindingUpdateOption func(*BindingUpdateOptions)
func (srv *Messaging) WithBindingUpdateChannel(v string) BindingUpdateOption {
	return func(o *BindingUpdateOptions) {
		o.Channel = v
		o.enabledSetters["Channel"] = true
	}
}
func (srv *Messaging) WithBindingUpdateEnabled(v bool) BindingUpdateOption {
	return func(o *BindingUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithBindingUpdateEventTopic(v string) BindingUpdateOption {
	return func(o *BindingUpdateOptions) {
		o.EventTopic = v
		o.enabledSetters["EventTopic"] = true
	}
}
func (srv *Messaging) WithBindingUpdateFallbackOrder(v int) BindingUpdateOption {
	return func(o *BindingUpdateOptions) {
		o.FallbackOrder = v
		o.enabledSetters["FallbackOrder"] = true
	}
}
func (srv *Messaging) WithBindingUpdateLocale(v string) BindingUpdateOption {
	return func(o *BindingUpdateOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Messaging) WithBindingUpdateRecipient(v string) BindingUpdateOption {
	return func(o *BindingUpdateOptions) {
		o.Recipient = v
		o.enabledSetters["Recipient"] = true
	}
}
func (srv *Messaging) WithBindingUpdateTemplateKey(v string) BindingUpdateOption {
	return func(o *BindingUpdateOptions) {
		o.TemplateKey = v
		o.enabledSetters["TemplateKey"] = true
	}
}
			
// BindingUpdate every field is optional; only what is sent is written.
// `enabled: false`
// is how a binding is taken out of service without losing what it said —
// the alternative is deleting it and typing the payload path back in
// correctly from memory later.
// 
// This path answers on `PUT` and `PATCH`, both routed to the same action.
func (srv *Messaging) BindingUpdate(Id string, optionalSetters ...BindingUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/bindings/{id}")
	options := BindingUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Channel"] {
		params["channel"] = options.Channel
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["EventTopic"] {
		params["event_topic"] = options.EventTopic
	}
	if options.enabledSetters["FallbackOrder"] {
		params["fallback_order"] = options.FallbackOrder
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["Recipient"] {
		params["recipient"] = options.Recipient
	}
	if options.enabledSetters["TemplateKey"] {
		params["template_key"] = options.TemplateKey
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
type ChannelCredentialIndexOptions struct {
	Market string
	Markets string
	enabledSetters map[string]bool
}
func (options ChannelCredentialIndexOptions) New() *ChannelCredentialIndexOptions {
	options.enabledSetters = map[string]bool{
		"Market": false,
		"Markets": false,
	}
	return &options
}
type ChannelCredentialIndexOption func(*ChannelCredentialIndexOptions)
func (srv *Messaging) WithChannelCredentialIndexMarket(v string) ChannelCredentialIndexOption {
	return func(o *ChannelCredentialIndexOptions) {
		o.Market = v
		o.enabledSetters["Market"] = true
	}
}
func (srv *Messaging) WithChannelCredentialIndexMarkets(v string) ChannelCredentialIndexOption {
	return func(o *ChannelCredentialIndexOptions) {
		o.Markets = v
		o.enabledSetters["Markets"] = true
	}
}
	
// ChannelCredentialIndex answers per channel with: which fields the chosen
// provider wants and
// which of them are SET (never their values — secrets go in and do not come
// back), which markets hold an override, which providers this build offers,
// whether the deployment has the channel switched on at all, the URL to
// paste into the provider's own console so bounces and opens come back, and
// whether callbacks are actually arriving.
// 
// Admin tier on the read as well as the write: the identifiers alone —
// which Twilio account, which sender number — are more than a read-only
// operator has reason to see, and the webhook URL served here contains the
// tenant's callback token.
func (srv *Messaging) ChannelCredentialIndex(optionalSetters ...ChannelCredentialIndexOption)(*models.Error, error) {
	path := "/v1/messaging/channel-credentials"
	options := ChannelCredentialIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Market"] {
		params["market"] = options.Market
	}
	if options.enabledSetters["Markets"] {
		params["markets"] = options.Markets
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
type ChannelCredentialDestroyOptions struct {
	Market string
	enabledSetters map[string]bool
}
func (options ChannelCredentialDestroyOptions) New() *ChannelCredentialDestroyOptions {
	options.enabledSetters = map[string]bool{
		"Market": false,
	}
	return &options
}
type ChannelCredentialDestroyOption func(*ChannelCredentialDestroyOptions)
func (srv *Messaging) WithChannelCredentialDestroyMarket(v string) ChannelCredentialDestroyOption {
	return func(o *ChannelCredentialDestroyOptions) {
		o.Market = v
		o.enabledSetters["Market"] = true
	}
}
			
// ChannelCredentialDestroy with `?market=`, only that market's override goes
// and the global
// credentials stand — the market then sends over the global provider again,
// which is what it did before anybody configured it. Without a market the
// channel goes entirely, overrides and all: a caller asking for a channel
// to hold no credentials means all of them.
// 
// 204 whether or not anything was there. The caller wants this channel to
// hold no credentials, and it does.
func (srv *Messaging) ChannelCredentialDestroy(Channel string, optionalSetters ...ChannelCredentialDestroyOption)(*models.Error, error) {
	r := strings.NewReplacer("{channel}", Channel)
	path := r.Replace("/v1/messaging/channel-credentials/{channel}")
	options := ChannelCredentialDestroyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["channel"] = Channel
	if options.enabledSetters["Market"] {
		params["market"] = options.Market
	}
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
type ChannelCredentialUpdatePatchOptions struct {
	Market string
	Driver string
	enabledSetters map[string]bool
}
func (options ChannelCredentialUpdatePatchOptions) New() *ChannelCredentialUpdatePatchOptions {
	options.enabledSetters = map[string]bool{
		"Market": false,
		"Driver": false,
	}
	return &options
}
type ChannelCredentialUpdatePatchOption func(*ChannelCredentialUpdatePatchOptions)
func (srv *Messaging) WithChannelCredentialUpdatePatchMarket(v string) ChannelCredentialUpdatePatchOption {
	return func(o *ChannelCredentialUpdatePatchOptions) {
		o.Market = v
		o.enabledSetters["Market"] = true
	}
}
func (srv *Messaging) WithChannelCredentialUpdatePatchDriver(v string) ChannelCredentialUpdatePatchOption {
	return func(o *ChannelCredentialUpdatePatchOptions) {
		o.Driver = v
		o.enabledSetters["Driver"] = true
	}
}
			
// ChannelCredentialUpdatePatch a PATCH in spirit whichever verb is used: only
// the fields present in the
// body are written, and the answer says which of them actually CHANGED, so
// a form that resent everything it had on screen does not report a change
// that did not happen.
// 
// Three refusals, all 422 and all deliberate rather than ignored. A field
// the channel's provider does not have (`unknown_credential_field`) — a
// typo sitting in the bag looking like configuration fails later with a
// message about a MISSING field the operator can see they filled in. A
// field the platform issues (`managed_credential`) — ignoring it would have
// the caller believe they set something. A channel with nothing to
// configure (`channel_not_configurable`), which is push: its VAPID keypair
// is generated at provisioning, and pasting a new one would orphan every
// browser registration the tenant has collected.
// 
// Switching provider is `driver`, and the fields in the same request are
// validated against the provider being switched TO — validating Postmark's
// key against Mailgun's field list is how a switch loses everything the
// operator just typed.
// 
// This path answers on `PUT` and `PATCH`, both routed to the same action.
func (srv *Messaging) ChannelCredentialUpdatePatch(Channel string, optionalSetters ...ChannelCredentialUpdatePatchOption)(*models.Error, error) {
	r := strings.NewReplacer("{channel}", Channel)
	path := r.Replace("/v1/messaging/channel-credentials/{channel}")
	options := ChannelCredentialUpdatePatchOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["channel"] = Channel
	if options.enabledSetters["Market"] {
		params["market"] = options.Market
	}
	if options.enabledSetters["Driver"] {
		params["driver"] = options.Driver
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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
type ChannelCredentialUpdateOptions struct {
	Market string
	Driver string
	enabledSetters map[string]bool
}
func (options ChannelCredentialUpdateOptions) New() *ChannelCredentialUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Market": false,
		"Driver": false,
	}
	return &options
}
type ChannelCredentialUpdateOption func(*ChannelCredentialUpdateOptions)
func (srv *Messaging) WithChannelCredentialUpdateMarket(v string) ChannelCredentialUpdateOption {
	return func(o *ChannelCredentialUpdateOptions) {
		o.Market = v
		o.enabledSetters["Market"] = true
	}
}
func (srv *Messaging) WithChannelCredentialUpdateDriver(v string) ChannelCredentialUpdateOption {
	return func(o *ChannelCredentialUpdateOptions) {
		o.Driver = v
		o.enabledSetters["Driver"] = true
	}
}
			
// ChannelCredentialUpdate a PATCH in spirit whichever verb is used: only the
// fields present in the
// body are written, and the answer says which of them actually CHANGED, so
// a form that resent everything it had on screen does not report a change
// that did not happen.
// 
// Three refusals, all 422 and all deliberate rather than ignored. A field
// the channel's provider does not have (`unknown_credential_field`) — a
// typo sitting in the bag looking like configuration fails later with a
// message about a MISSING field the operator can see they filled in. A
// field the platform issues (`managed_credential`) — ignoring it would have
// the caller believe they set something. A channel with nothing to
// configure (`channel_not_configurable`), which is push: its VAPID keypair
// is generated at provisioning, and pasting a new one would orphan every
// browser registration the tenant has collected.
// 
// Switching provider is `driver`, and the fields in the same request are
// validated against the provider being switched TO — validating Postmark's
// key against Mailgun's field list is how a switch loses everything the
// operator just typed.
// 
// This path answers on `PUT` and `PATCH`, both routed to the same action.
func (srv *Messaging) ChannelCredentialUpdate(Channel string, optionalSetters ...ChannelCredentialUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{channel}", Channel)
	path := r.Replace("/v1/messaging/channel-credentials/{channel}")
	options := ChannelCredentialUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["channel"] = Channel
	if options.enabledSetters["Market"] {
		params["market"] = options.Market
	}
	if options.enabledSetters["Driver"] {
		params["driver"] = options.Driver
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
type ChannelCredentialVerifyOptions struct {
	Market string
	enabledSetters map[string]bool
}
func (options ChannelCredentialVerifyOptions) New() *ChannelCredentialVerifyOptions {
	options.enabledSetters = map[string]bool{
		"Market": false,
	}
	return &options
}
type ChannelCredentialVerifyOption func(*ChannelCredentialVerifyOptions)
func (srv *Messaging) WithChannelCredentialVerifyMarket(v string) ChannelCredentialVerifyOption {
	return func(o *ChannelCredentialVerifyOptions) {
		o.Market = v
		o.enabledSetters["Market"] = true
	}
}
			
// ChannelCredentialVerify the one thing that turns this screen from a form
// into a tool. Credentials
// that only fail at send time cost a customer their first order
// confirmation, and by then nobody connects the failure to the afternoon
// somebody pasted a key with a trailing space.
// 
// **Always 200.** The answer is `{ok, message}` in the body, including when
// the credentials are wrong: the REQUEST was fine, the credentials are not,
// and a 4xx here would have the cockpit's own error handling swallow the
// one sentence worth reading. A channel that asks for no credentials at all
// (push, in-app) answers `ok: true` — "nothing to verify" is a finished
// check, not a failed one, and reporting it as an error painted a channel
// that has worked since provisioning in the same red as a wrong token.
func (srv *Messaging) ChannelCredentialVerify(Channel string, optionalSetters ...ChannelCredentialVerifyOption)(*models.Error, error) {
	r := strings.NewReplacer("{channel}", Channel)
	path := r.Replace("/v1/messaging/channel-credentials/{channel}/verify")
	options := ChannelCredentialVerifyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["channel"] = Channel
	if options.enabledSetters["Market"] {
		params["market"] = options.Market
	}
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

// ChannelIndex each entry says whether the channel is switched on and which
// provider
// carries it by default. A channel that is off will refuse a send, so a UI
// that offers a channel picker should build it from this rather than from a
// list of its own — a channel added to the service then appears without a
// release of the client.
func (srv *Messaging) ChannelIndex()(*models.Error, error) {
	path := "/v1/messaging/channels"
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

// ConfigShow a tenant that was never provisioned has no row and still gets an
// answer:
// an empty shape rather than a 404, so the Cockpit's panels open on
// editable blanks instead of an error.
// 
// `meta.push_public_key` is the VAPID public key, and only the public one.
// A storefront cannot call `PushManager.subscribe()` without it, so it has
// to leave the service; the private half and every provider secret stay
// hidden on the model, where they are protected on every route rather than
// on this one.
func (srv *Messaging) ConfigShow()(*models.Error, error) {
	path := "/v1/messaging/config"
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
type ConfigUpdatePatchOptions struct {
	DefaultLocale string
	Defaults []string
	Product string
	QuietHours []string
	SupportEmail string
	enabledSetters map[string]bool
}
func (options ConfigUpdatePatchOptions) New() *ConfigUpdatePatchOptions {
	options.enabledSetters = map[string]bool{
		"DefaultLocale": false,
		"Defaults": false,
		"Product": false,
		"QuietHours": false,
		"SupportEmail": false,
	}
	return &options
}
type ConfigUpdatePatchOption func(*ConfigUpdatePatchOptions)
func (srv *Messaging) WithConfigUpdatePatchDefaultLocale(v string) ConfigUpdatePatchOption {
	return func(o *ConfigUpdatePatchOptions) {
		o.DefaultLocale = v
		o.enabledSetters["DefaultLocale"] = true
	}
}
func (srv *Messaging) WithConfigUpdatePatchDefaults(v []string) ConfigUpdatePatchOption {
	return func(o *ConfigUpdatePatchOptions) {
		o.Defaults = v
		o.enabledSetters["Defaults"] = true
	}
}
func (srv *Messaging) WithConfigUpdatePatchProduct(v string) ConfigUpdatePatchOption {
	return func(o *ConfigUpdatePatchOptions) {
		o.Product = v
		o.enabledSetters["Product"] = true
	}
}
func (srv *Messaging) WithConfigUpdatePatchQuietHours(v []string) ConfigUpdatePatchOption {
	return func(o *ConfigUpdatePatchOptions) {
		o.QuietHours = v
		o.enabledSetters["QuietHours"] = true
	}
}
func (srv *Messaging) WithConfigUpdatePatchSupportEmail(v string) ConfigUpdatePatchOption {
	return func(o *ConfigUpdatePatchOptions) {
		o.SupportEmail = v
		o.enabledSetters["SupportEmail"] = true
	}
}
	
// ConfigUpdatePatch reaches every message this tenant sends, including
// templates saved months
// ago — content placeholders resolve at send time, not at save time —
// which
// is why writing is admin tier while reading is not.
// 
// Two refusals worth knowing about. `defaults.brand` is 422, not ignored:
// the letterhead moved to /v1/layouts when a tenant gained more than one of
// them, and a letterhead edit that appears to save and changes nothing is
// the worst of the three possible behaviours. A half-written `quiet_hours`
// is 422 as well — a tenant that typed a start and forgot the end has an
// opinion about when not to message people, and silently sending through
// the night is the one answer that is definitely wrong.
// 
// Provider credentials cannot be written here. That path is
// /v1/channel-credentials, so the one route that handles secrets stays the
// one that was built for it.
// 
// This path answers on `PUT` and `PATCH`, both routed to the same action.
func (srv *Messaging) ConfigUpdatePatch(optionalSetters ...ConfigUpdatePatchOption)(*models.Error, error) {
	path := "/v1/messaging/config"
	options := ConfigUpdatePatchOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["DefaultLocale"] {
		params["default_locale"] = options.DefaultLocale
	}
	if options.enabledSetters["Defaults"] {
		params["defaults"] = options.Defaults
	}
	if options.enabledSetters["Product"] {
		params["product"] = options.Product
	}
	if options.enabledSetters["QuietHours"] {
		params["quiet_hours"] = options.QuietHours
	}
	if options.enabledSetters["SupportEmail"] {
		params["support_email"] = options.SupportEmail
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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
type ConfigUpdateOptions struct {
	DefaultLocale string
	Defaults []string
	Product string
	QuietHours []string
	SupportEmail string
	enabledSetters map[string]bool
}
func (options ConfigUpdateOptions) New() *ConfigUpdateOptions {
	options.enabledSetters = map[string]bool{
		"DefaultLocale": false,
		"Defaults": false,
		"Product": false,
		"QuietHours": false,
		"SupportEmail": false,
	}
	return &options
}
type ConfigUpdateOption func(*ConfigUpdateOptions)
func (srv *Messaging) WithConfigUpdateDefaultLocale(v string) ConfigUpdateOption {
	return func(o *ConfigUpdateOptions) {
		o.DefaultLocale = v
		o.enabledSetters["DefaultLocale"] = true
	}
}
func (srv *Messaging) WithConfigUpdateDefaults(v []string) ConfigUpdateOption {
	return func(o *ConfigUpdateOptions) {
		o.Defaults = v
		o.enabledSetters["Defaults"] = true
	}
}
func (srv *Messaging) WithConfigUpdateProduct(v string) ConfigUpdateOption {
	return func(o *ConfigUpdateOptions) {
		o.Product = v
		o.enabledSetters["Product"] = true
	}
}
func (srv *Messaging) WithConfigUpdateQuietHours(v []string) ConfigUpdateOption {
	return func(o *ConfigUpdateOptions) {
		o.QuietHours = v
		o.enabledSetters["QuietHours"] = true
	}
}
func (srv *Messaging) WithConfigUpdateSupportEmail(v string) ConfigUpdateOption {
	return func(o *ConfigUpdateOptions) {
		o.SupportEmail = v
		o.enabledSetters["SupportEmail"] = true
	}
}
	
// ConfigUpdate reaches every message this tenant sends, including templates
// saved months
// ago — content placeholders resolve at send time, not at save time —
// which
// is why writing is admin tier while reading is not.
// 
// Two refusals worth knowing about. `defaults.brand` is 422, not ignored:
// the letterhead moved to /v1/layouts when a tenant gained more than one of
// them, and a letterhead edit that appears to save and changes nothing is
// the worst of the three possible behaviours. A half-written `quiet_hours`
// is 422 as well — a tenant that typed a start and forgot the end has an
// opinion about when not to message people, and silently sending through
// the night is the one answer that is definitely wrong.
// 
// Provider credentials cannot be written here. That path is
// /v1/channel-credentials, so the one route that handles secrets stays the
// one that was built for it.
// 
// This path answers on `PUT` and `PATCH`, both routed to the same action.
func (srv *Messaging) ConfigUpdate(optionalSetters ...ConfigUpdateOption)(*models.Error, error) {
	path := "/v1/messaging/config"
	options := ConfigUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["DefaultLocale"] {
		params["default_locale"] = options.DefaultLocale
	}
	if options.enabledSetters["Defaults"] {
		params["defaults"] = options.Defaults
	}
	if options.enabledSetters["Product"] {
		params["product"] = options.Product
	}
	if options.enabledSetters["QuietHours"] {
		params["quiet_hours"] = options.QuietHours
	}
	if options.enabledSetters["SupportEmail"] {
		params["support_email"] = options.SupportEmail
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
type LayoutIndexOptions struct {
	Markets string
	enabledSetters map[string]bool
}
func (options LayoutIndexOptions) New() *LayoutIndexOptions {
	options.enabledSetters = map[string]bool{
		"Markets": false,
	}
	return &options
}
type LayoutIndexOption func(*LayoutIndexOptions)
func (srv *Messaging) WithLayoutIndexMarkets(v string) LayoutIndexOption {
	return func(o *LayoutIndexOptions) {
		o.Markets = v
		o.enabledSetters["Markets"] = true
	}
}
	
// LayoutIndex the order is the list's purpose: it is a picker, and the entry
// most
// templates are actually on belongs at the top of it.
// 
// Market-scoped as a browsing filter — see the parameters. `GET
// /layouts/{id}`
// deliberately is not: somebody holding an id may read it.
func (srv *Messaging) LayoutIndex(optionalSetters ...LayoutIndexOption)(*models.Error, error) {
	path := "/v1/messaging/layouts"
	options := LayoutIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Markets"] {
		params["markets"] = options.Markets
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

// LayoutStore a tenant's FIRST layout becomes the default whatever the
// request says: a
// tenant with no default cannot compile a template that does not name one.
// 
// The default may hold neither a validity window nor `enabled: false`, and
// asking for both in one request is refused with 422
// `layout_default_always_in_force`. There is no fallback behind the default
// — every template that names no layout is framed by it — so a window set
// today would take a tenant's whole letterhead away on a morning months
// from now, with nobody left who remembers typing the date.
func (srv *Messaging) LayoutStore()(*models.Error, error) {
	path := "/v1/messaging/layouts"
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
	
// LayoutDestroy answers 200 with a body rather than the 204 the other
// resources use: the
// count of reassigned templates is the part an operator needs, and a
// deletion that silently moved eleven templates onto another letterhead is
// one they would only discover from the next mail that went out.
func (srv *Messaging) LayoutDestroy(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/layouts/{id}")
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
	
// LayoutShow not market-filtered, deliberately: market scoping is a browsing
// concern,
// and somebody holding an id may read the row. A template pinned to a
// layout keeps mailing on it whatever market the reader is looking at.
func (srv *Messaging) LayoutShow(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/layouts/{id}")
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
	
// LayoutUpdate the change reaches every template on this layout, including
// ones saved
// months ago and never opened since — which is exactly the change nobody
// remembers making when the mails start looking wrong. It is audited for
// that reason, and only when something actually changed: an audit line on
// every save teaches its readers to ignore the log.
// 
// Two 422s. Clearing `is_default` on the current default is
// `layout_default_required` — promoting another layout is the operation
// that exists for this, and it clears this one as a side effect, which is
// the only way the count stays at exactly one. Giving the default a
// validity window or switching it off is `layout_default_always_in_force`,
// and the check is made of the OUTCOME, so promoting a layout and dating it
// in the same request is caught.
// 
// The structural half of a layout — colours, width, font — is baked into
// each template's compiled body, so templates already on it keep the old
// one until they are recompiled.
func (srv *Messaging) LayoutUpdate(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/layouts/{id}")
	params := map[string]interface{}{}
	params["id"] = Id
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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
type LibraryIndexOptions struct {
	Channel string
	Locale string
	enabledSetters map[string]bool
}
func (options LibraryIndexOptions) New() *LibraryIndexOptions {
	options.enabledSetters = map[string]bool{
		"Channel": false,
		"Locale": false,
	}
	return &options
}
type LibraryIndexOption func(*LibraryIndexOptions)
func (srv *Messaging) WithLibraryIndexChannel(v string) LibraryIndexOption {
	return func(o *LibraryIndexOptions) {
		o.Channel = v
		o.enabledSetters["Channel"] = true
	}
}
func (srv *Messaging) WithLibraryIndexLocale(v string) LibraryIndexOption {
	return func(o *LibraryIndexOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
	
// LibraryIndex what the Cockpit's "start from a template" gallery is built
// from. These
// are not the tenant's rows and cannot be edited here: provisioning clones
// them into `/v1/templates`, and it is the clone that a tenant owns.
func (srv *Messaging) LibraryIndex(optionalSetters ...LibraryIndexOption)(*models.Error, error) {
	path := "/v1/messaging/library"
	options := LibraryIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Channel"] {
		params["channel"] = options.Channel
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
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
type MessageIndexOptions struct {
	Channel string
	Status string
	enabledSetters map[string]bool
}
func (options MessageIndexOptions) New() *MessageIndexOptions {
	options.enabledSetters = map[string]bool{
		"Channel": false,
		"Status": false,
	}
	return &options
}
type MessageIndexOption func(*MessageIndexOptions)
func (srv *Messaging) WithMessageIndexChannel(v string) MessageIndexOption {
	return func(o *MessageIndexOptions) {
		o.Channel = v
		o.enabledSetters["Channel"] = true
	}
}
func (srv *Messaging) WithMessageIndexStatus(v string) MessageIndexOption {
	return func(o *MessageIndexOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
	
// MessageIndex `?channel=` and `?status=` narrow it; `?limit=` is clamped to
// 200 and
// defaults to 50. `?channel=inapp` is the tenant's in-app inbox — the
// Message row IS the inbox item, so there is no second store for it.
// 
// Rows are subject to the deployment's retention window and to erasure
// requests, so this is not an archive.
func (srv *Messaging) MessageIndex(optionalSetters ...MessageIndexOption)(*models.Error, error) {
	path := "/v1/messaging/messages"
	options := MessageIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Channel"] {
		params["channel"] = options.Channel
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
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
	
// MessageShow carries the render model it was sent with, so "why did this
// mail say
// * that" is answerable after the fact. That is also why the row is personal
// data and why it can be erased — see POST /v1/privacy/erasures.
func (srv *Messaging) MessageShow(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/messages/{id}")
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
type SendPreviewOptions struct {
	Data interface{}
	Locale string
	enabledSetters map[string]bool
}
func (options SendPreviewOptions) New() *SendPreviewOptions {
	options.enabledSetters = map[string]bool{
		"Data": false,
		"Locale": false,
	}
	return &options
}
type SendPreviewOption func(*SendPreviewOptions)
func (srv *Messaging) WithSendPreviewData(v interface{}) SendPreviewOption {
	return func(o *SendPreviewOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *Messaging) WithSendPreviewLocale(v string) SendPreviewOption {
	return func(o *SendPreviewOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
					
// SendPreview answers with the resolved subject, HTML and text exactly as a
// real send
// would produce them, so an editor can show a faithful preview without a
// message row, a provider call or a suppression check.
// 
// Takes no `market`, deliberately: rendering picks no provider, so there is
// nothing here for a market to change. Nor `send_at`, `draft` or
// `attachments` — all of them are properties of a dispatch, not of a
// render.
func (srv *Messaging) SendPreview(Channel string, Template string, optionalSetters ...SendPreviewOption)(*models.Error, error) {
	path := "/v1/messaging/preview"
	options := SendPreviewOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["channel"] = Channel
	params["template"] = Template
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
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
			
// ErasureStore per (channel, address), because an address is channel-shaped
// and the rows
// it has to line up with are keyed that way. Matching is done on the
// normalised form on both sides, so a request for `ada@acme.test` finds a
// log written for `Ada@Acme.test` — an erasure that misses on
// capitalisation is an erasure that did not happen and reports success.
// 
// Message rows and unsubscribe tokens are DELETED. Suppressions are KEPT
// with the clear-text address nulled: matching runs on a keyed hash, so the
// row can still block and can no longer identify. Deleting it instead is
// the obvious reading of "erase everything about them", and it is the
// reading that mails a dead address again next week — or mails somebody who
// complained, which is how a sending domain gets blocked.
// 
// Answers with the counts, `suppressions_kept` among them, so the design is
// stated in the response rather than only in this paragraph.
func (srv *Messaging) ErasureStore(Address string, Channel string)(*models.Error, error) {
	path := "/v1/messaging/privacy/erasures"
	params := map[string]interface{}{}
	params["address"] = Address
	params["channel"] = Channel
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
	
// PushSubscriptionDestroy by endpoint and not by id, because the browser
// knows its endpoint and has
// never seen our id — this is called from a service worker reacting to
// `pushsubscriptionchange`, or from a "turn off notifications" button.
func (srv *Messaging) PushSubscriptionDestroy(Endpoint string)(*models.Error, error) {
	path := "/v1/messaging/push/subscriptions"
	params := map[string]interface{}{}
	params["endpoint"] = Endpoint
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
	
// PushSubscriptionIndex `subscriber_id` is required: this is not a list of
// everybody, and there
// is no route that is. The caller is a storefront acting for one visitor
// and has no business enumerating the rest.
// 
// The client key material is never returned — see the `$hidden` list on the
// model. A registration that can be read back is a registration somebody
// else can push with.
func (srv *Messaging) PushSubscriptionIndex(SubscriberId string)(*models.Error, error) {
	path := "/v1/messaging/push/subscriptions"
	params := map[string]interface{}{}
	params["subscriber_id"] = SubscriberId
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
type PushSubscriptionStoreOptions struct {
	UserAgent string
	enabledSetters map[string]bool
}
func (options PushSubscriptionStoreOptions) New() *PushSubscriptionStoreOptions {
	options.enabledSetters = map[string]bool{
		"UserAgent": false,
	}
	return &options
}
type PushSubscriptionStoreOption func(*PushSubscriptionStoreOptions)
func (srv *Messaging) WithPushSubscriptionStoreUserAgent(v string) PushSubscriptionStoreOption {
	return func(o *PushSubscriptionStoreOptions) {
		o.UserAgent = v
		o.enabledSetters["UserAgent"] = true
	}
}
							
// PushSubscriptionStore send what `PushManager.subscribe()` handed back —
// the endpoint and the
// two keys — plus the id you know that person by. The VAPID public key the
// browser needs to produce it comes from `GET /v1/config`
// (`meta.push_public_key`).
// 
// **Idempotent by endpoint**, and the two statuses say which happened: 201
// for a browser seen for the first time, 200 for one already registered. A
// browser calls `subscribe()` on every page load and hands back the same
// endpoint each time; treating that as a new device would give one laptop a
// thousand rows and push to it a thousand times.
func (srv *Messaging) PushSubscriptionStore(Endpoint string, Keys interface{}, SubscriberId string, optionalSetters ...PushSubscriptionStoreOption)(*models.Error, error) {
	path := "/v1/messaging/push/subscriptions"
	options := PushSubscriptionStoreOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["endpoint"] = Endpoint
	params["keys"] = Keys
	params["subscriber_id"] = SubscriberId
	if options.enabledSetters["UserAgent"] {
		params["user_agent"] = options.UserAgent
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
type SendSendOptions struct {
	Attachments []interface{}
	Data interface{}
	Draft bool
	Locale string
	Market string
	SendAt string
	enabledSetters map[string]bool
}
func (options SendSendOptions) New() *SendSendOptions {
	options.enabledSetters = map[string]bool{
		"Attachments": false,
		"Data": false,
		"Draft": false,
		"Locale": false,
		"Market": false,
		"SendAt": false,
	}
	return &options
}
type SendSendOption func(*SendSendOptions)
func (srv *Messaging) WithSendSendAttachments(v []interface{}) SendSendOption {
	return func(o *SendSendOptions) {
		o.Attachments = v
		o.enabledSetters["Attachments"] = true
	}
}
func (srv *Messaging) WithSendSendData(v interface{}) SendSendOption {
	return func(o *SendSendOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *Messaging) WithSendSendDraft(v bool) SendSendOption {
	return func(o *SendSendOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithSendSendLocale(v string) SendSendOption {
	return func(o *SendSendOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Messaging) WithSendSendMarket(v string) SendSendOption {
	return func(o *SendSendOptions) {
		o.Market = v
		o.enabledSetters["Market"] = true
	}
}
func (srv *Messaging) WithSendSendSendAt(v string) SendSendOption {
	return func(o *SendSendOptions) {
		o.SendAt = v
		o.enabledSetters["SendAt"] = true
	}
}
							
// SendSend renders a tenant template and dispatches it — now, at `send_at`,
// or at
// the end of the tenant's quiet hours.
// 
// The first line is deliberately a title, not a sentence about the
// mechanism: Scramble takes it as the operation's `summary`, and a summary
// is what an API explorer prints in its route list. The paragraph that used
// to be here ran to 119 characters across two lines, which the gateway's
// fragment tests reject for exactly that reason.
// 
// Retry-safe when the caller sends an `Idempotency-Key` header. The two
// answers are deliberately different:
// 
// 201 — a message was created by THIS call
// 200 — this key was already used; here is the message it produced
// 
// A caller has to be able to tell those apart. "Your mail went out" and
// "your mail had already gone out" are the same outcome and different
// facts, and a client reconciling its own records needs the second one.
// Same key with a different body is a 422 — see IdempotencyConflict.
// 
// A recipient on the tenant's suppression list is not sent to, and that is
// reported as a refusal rather than as a silent success.
func (srv *Messaging) SendSend(Channel string, Template string, To string, optionalSetters ...SendSendOption)(*models.Error, error) {
	path := "/v1/messaging/send"
	options := SendSendOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["channel"] = Channel
	params["template"] = Template
	params["to"] = To
	if options.enabledSetters["Attachments"] {
		params["attachments"] = options.Attachments
	}
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["Market"] {
		params["market"] = options.Market
	}
	if options.enabledSetters["SendAt"] {
		params["send_at"] = options.SendAt
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
type StatsIndexOptions struct {
	Days int
	From string
	To string
	enabledSetters map[string]bool
}
func (options StatsIndexOptions) New() *StatsIndexOptions {
	options.enabledSetters = map[string]bool{
		"Days": false,
		"From": false,
		"To": false,
	}
	return &options
}
type StatsIndexOption func(*StatsIndexOptions)
func (srv *Messaging) WithStatsIndexDays(v int) StatsIndexOption {
	return func(o *StatsIndexOptions) {
		o.Days = v
		o.enabledSetters["Days"] = true
	}
}
func (srv *Messaging) WithStatsIndexFrom(v string) StatsIndexOption {
	return func(o *StatsIndexOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
func (srv *Messaging) WithStatsIndexTo(v string) StatsIndexOption {
	return func(o *StatsIndexOptions) {
		o.To = v
		o.enabledSetters["To"] = true
	}
}
	
// StatsIndex either `days` (a window ending now, default 30) or an explicit
// `from`/`to`
// span. Both ends of the span or neither: `from` alone would be an open
// range and the service would have to guess which end was meant.
// 
// Three numbers are deliberately not the naive ones, and the `window` block
// says so rather than leaving a chart to imply otherwise. The window is
// CLAMPED to the tenant's retention, and `clamped_by_retention` says when
// that happened — 90 days on a 30-day retention is 30 days of data wearing
// a 90-day label, and the trend line it draws invents a collapse that never
// happened. Opens are counted only over channels that can report them; SMS
// and push have no such thing, so dividing opens by all messages would
// quietly halve every open rate the moment a tenant adds a second channel.
// The delivery rate is sent ÷ (sent + failed): suppressed is the service
// doing what it was told, and counting it as a failure would punish a
// tenant for having a working unsubscribe list.
// 
// `previous` is the same window again immediately before this one, which is
// what turns a figure into a direction. **It is null** whenever the
// preceding window is not entirely inside retention: the query would answer
// zero rather than fail, and zero against 1,337 renders as a triumphant
// +100 % beside every tile on the screen. Show no trend rather than a
// flattering one.
// 
// Nothing here names a recipient. That is the delivery log, which is a
// different endpoint with a different question.
func (srv *Messaging) StatsIndex(optionalSetters ...StatsIndexOption)(*models.Error, error) {
	path := "/v1/messaging/stats"
	options := StatsIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Days"] {
		params["days"] = options.Days
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	if options.enabledSetters["To"] {
		params["to"] = options.To
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
type SuppressionIndexOptions struct {
	Channel string
	Scope string
	Reason string
	Address string
	Limit int
	enabledSetters map[string]bool
}
func (options SuppressionIndexOptions) New() *SuppressionIndexOptions {
	options.enabledSetters = map[string]bool{
		"Channel": false,
		"Scope": false,
		"Reason": false,
		"Address": false,
		"Limit": false,
	}
	return &options
}
type SuppressionIndexOption func(*SuppressionIndexOptions)
func (srv *Messaging) WithSuppressionIndexChannel(v string) SuppressionIndexOption {
	return func(o *SuppressionIndexOptions) {
		o.Channel = v
		o.enabledSetters["Channel"] = true
	}
}
func (srv *Messaging) WithSuppressionIndexScope(v string) SuppressionIndexOption {
	return func(o *SuppressionIndexOptions) {
		o.Scope = v
		o.enabledSetters["Scope"] = true
	}
}
func (srv *Messaging) WithSuppressionIndexReason(v string) SuppressionIndexOption {
	return func(o *SuppressionIndexOptions) {
		o.Reason = v
		o.enabledSetters["Reason"] = true
	}
}
func (srv *Messaging) WithSuppressionIndexAddress(v string) SuppressionIndexOption {
	return func(o *SuppressionIndexOptions) {
		o.Address = v
		o.enabledSetters["Address"] = true
	}
}
func (srv *Messaging) WithSuppressionIndexLimit(v int) SuppressionIndexOption {
	return func(o *SuppressionIndexOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
	
// SuppressionIndex filterable by `channel`, `scope`, `reason` and `address`.
// The address
// filter is looked up by FINGERPRINT rather than against the address
// column, which is what makes "why did this person stop getting our mail"
// answerable for somebody who has since been erased: the row has no
// address left to match on, and the question is still the same question.
func (srv *Messaging) SuppressionIndex(optionalSetters ...SuppressionIndexOption)(*models.Error, error) {
	path := "/v1/messaging/suppressions"
	options := SuppressionIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Channel"] {
		params["channel"] = options.Channel
	}
	if options.enabledSetters["Scope"] {
		params["scope"] = options.Scope
	}
	if options.enabledSetters["Reason"] {
		params["reason"] = options.Reason
	}
	if options.enabledSetters["Address"] {
		params["address"] = options.Address
	}
	if options.enabledSetters["Limit"] {
		params["limit"] = options.Limit
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
type SuppressionStoreOptions struct {
	ExpiresAt string
	Note string
	Scope string
	enabledSetters map[string]bool
}
func (options SuppressionStoreOptions) New() *SuppressionStoreOptions {
	options.enabledSetters = map[string]bool{
		"ExpiresAt": false,
		"Note": false,
		"Scope": false,
	}
	return &options
}
type SuppressionStoreOption func(*SuppressionStoreOptions)
func (srv *Messaging) WithSuppressionStoreExpiresAt(v string) SuppressionStoreOption {
	return func(o *SuppressionStoreOptions) {
		o.ExpiresAt = v
		o.enabledSetters["ExpiresAt"] = true
	}
}
func (srv *Messaging) WithSuppressionStoreNote(v string) SuppressionStoreOption {
	return func(o *SuppressionStoreOptions) {
		o.Note = v
		o.enabledSetters["Note"] = true
	}
}
func (srv *Messaging) WithSuppressionStoreScope(v string) SuppressionStoreOption {
	return func(o *SuppressionStoreOptions) {
		o.Scope = v
		o.enabledSetters["Scope"] = true
	}
}
							
// SuppressionStore 201 for a row this call created, 200 for an address that
// was already on
// the list — so a client can tell whether it changed anything.
// 
// The `scope` follows from the `reason` for every reason but `manual`, and
// asking for a different one is 422 `suppression_scope_fixed` rather than
// being quietly corrected: a caller who asked for `marketing` on a hard
// bounce has the model wrong, and a silent upgrade to `all` would leave
// them believing transactional mail still flows to an address that does not
// exist.
func (srv *Messaging) SuppressionStore(Address string, Channel string, Reason string, optionalSetters ...SuppressionStoreOption)(*models.Error, error) {
	path := "/v1/messaging/suppressions"
	options := SuppressionStoreOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["address"] = Address
	params["channel"] = Channel
	params["reason"] = Reason
	if options.enabledSetters["ExpiresAt"] {
		params["expires_at"] = options.ExpiresAt
	}
	if options.enabledSetters["Note"] {
		params["note"] = options.Note
	}
	if options.enabledSetters["Scope"] {
		params["scope"] = options.Scope
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
	
// SuppressionDestroy audited, unlike most deletes in this service. Removing a
// row here is the
// one operation that makes the service mail an address something decided
// not to mail — if a complaint turns into a spam report later, "who took
// * this off the list, and when" is the whole investigation.
func (srv *Messaging) SuppressionDestroy(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/suppressions/{id}")
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
	
// SuppressionShow `address` may be null: that is a person who has been erased
// (POST /v1/privacy/erasures). The row survives as a hash, which is the
// point — the clear text is gone and the address is still blocked.
func (srv *Messaging) SuppressionShow(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/suppressions/{id}")
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
type TemplateIndexOptions struct {
	Channel string
	Markets string
	enabledSetters map[string]bool
}
func (options TemplateIndexOptions) New() *TemplateIndexOptions {
	options.enabledSetters = map[string]bool{
		"Channel": false,
		"Markets": false,
	}
	return &options
}
type TemplateIndexOption func(*TemplateIndexOptions)
func (srv *Messaging) WithTemplateIndexChannel(v string) TemplateIndexOption {
	return func(o *TemplateIndexOptions) {
		o.Channel = v
		o.enabledSetters["Channel"] = true
	}
}
func (srv *Messaging) WithTemplateIndexMarkets(v string) TemplateIndexOption {
	return func(o *TemplateIndexOptions) {
		o.Markets = v
		o.enabledSetters["Markets"] = true
	}
}
	
// TemplateIndex `?channel=` narrows to one channel. Market-scoped as a
// BROWSING filter:
// with `X-Revenexx-Market` the list is the global rows plus that market's,
// without it the global rows only, and `?markets=all` is the unscoped read.
// Never a boundary — the tenant is fixed by the credential and by row-level
// security, and no value of either parameter reaches another tenant's rows.
func (srv *Messaging) TemplateIndex(optionalSetters ...TemplateIndexOption)(*models.Error, error) {
	path := "/v1/messaging/templates"
	options := TemplateIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Channel"] {
		params["channel"] = options.Channel
	}
	if options.enabledSetters["Markets"] {
		params["markets"] = options.Markets
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
type TemplateStoreOptions struct {
	BodyHtml string
	BodyText string
	ContentSid string
	Design []string
	Enabled bool
	LayoutId string
	Locale string
	Markets []string
	MessageClass string
	Subject string
	TestMode bool
	Title string
	ValidFrom string
	ValidUntil string
	VariableDefaults []string
	Variables []string
	WhatsappCategory string
	enabledSetters map[string]bool
}
func (options TemplateStoreOptions) New() *TemplateStoreOptions {
	options.enabledSetters = map[string]bool{
		"BodyHtml": false,
		"BodyText": false,
		"ContentSid": false,
		"Design": false,
		"Enabled": false,
		"LayoutId": false,
		"Locale": false,
		"Markets": false,
		"MessageClass": false,
		"Subject": false,
		"TestMode": false,
		"Title": false,
		"ValidFrom": false,
		"ValidUntil": false,
		"VariableDefaults": false,
		"Variables": false,
		"WhatsappCategory": false,
	}
	return &options
}
type TemplateStoreOption func(*TemplateStoreOptions)
func (srv *Messaging) WithTemplateStoreBodyHtml(v string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.BodyHtml = v
		o.enabledSetters["BodyHtml"] = true
	}
}
func (srv *Messaging) WithTemplateStoreBodyText(v string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.BodyText = v
		o.enabledSetters["BodyText"] = true
	}
}
func (srv *Messaging) WithTemplateStoreContentSid(v string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.ContentSid = v
		o.enabledSetters["ContentSid"] = true
	}
}
func (srv *Messaging) WithTemplateStoreDesign(v []string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.Design = v
		o.enabledSetters["Design"] = true
	}
}
func (srv *Messaging) WithTemplateStoreEnabled(v bool) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithTemplateStoreLayoutId(v string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.LayoutId = v
		o.enabledSetters["LayoutId"] = true
	}
}
func (srv *Messaging) WithTemplateStoreLocale(v string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Messaging) WithTemplateStoreMarkets(v []string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.Markets = v
		o.enabledSetters["Markets"] = true
	}
}
func (srv *Messaging) WithTemplateStoreMessageClass(v string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.MessageClass = v
		o.enabledSetters["MessageClass"] = true
	}
}
func (srv *Messaging) WithTemplateStoreSubject(v string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.Subject = v
		o.enabledSetters["Subject"] = true
	}
}
func (srv *Messaging) WithTemplateStoreTestMode(v bool) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.TestMode = v
		o.enabledSetters["TestMode"] = true
	}
}
func (srv *Messaging) WithTemplateStoreTitle(v string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *Messaging) WithTemplateStoreValidFrom(v string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.ValidFrom = v
		o.enabledSetters["ValidFrom"] = true
	}
}
func (srv *Messaging) WithTemplateStoreValidUntil(v string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.ValidUntil = v
		o.enabledSetters["ValidUntil"] = true
	}
}
func (srv *Messaging) WithTemplateStoreVariableDefaults(v []string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.VariableDefaults = v
		o.enabledSetters["VariableDefaults"] = true
	}
}
func (srv *Messaging) WithTemplateStoreVariables(v []string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.Variables = v
		o.enabledSetters["Variables"] = true
	}
}
func (srv *Messaging) WithTemplateStoreWhatsappCategory(v string) TemplateStoreOption {
	return func(o *TemplateStoreOptions) {
		o.WhatsappCategory = v
		o.enabledSetters["WhatsappCategory"] = true
	}
}
					
// TemplateStore send a `design` document and the service compiles it against
// the
// template's layout — or send `body_html` and `body_text` yourself and skip
// compilation entirely.
// 
// A design that the compiler refuses is 422 and NOTHING is written, with
// `error.details` naming the offending block. That order is deliberate: a
// save whose compile failed must leave the row alone, because storing the
// design while keeping a stale body would hand the next send a mail that no
// longer matches the document it claims to be built from, and nothing would
// ever surface it. A sidecar that is down is 503 `mjml_unavailable`, which
// is worth retrying; a rejected design is not.
// 
// The row this creates is a DRAFT and sends nothing until it is published.
func (srv *Messaging) TemplateStore(Channel string, Key string, optionalSetters ...TemplateStoreOption)(*models.Error, error) {
	path := "/v1/messaging/templates"
	options := TemplateStoreOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["channel"] = Channel
	params["key"] = Key
	if options.enabledSetters["BodyHtml"] {
		params["body_html"] = options.BodyHtml
	}
	if options.enabledSetters["BodyText"] {
		params["body_text"] = options.BodyText
	}
	if options.enabledSetters["ContentSid"] {
		params["content_sid"] = options.ContentSid
	}
	if options.enabledSetters["Design"] {
		params["design"] = options.Design
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["LayoutId"] {
		params["layout_id"] = options.LayoutId
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["Markets"] {
		params["markets"] = options.Markets
	}
	if options.enabledSetters["MessageClass"] {
		params["message_class"] = options.MessageClass
	}
	if options.enabledSetters["Subject"] {
		params["subject"] = options.Subject
	}
	if options.enabledSetters["TestMode"] {
		params["test_mode"] = options.TestMode
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["ValidFrom"] {
		params["valid_from"] = options.ValidFrom
	}
	if options.enabledSetters["ValidUntil"] {
		params["valid_until"] = options.ValidUntil
	}
	if options.enabledSetters["VariableDefaults"] {
		params["variable_defaults"] = options.VariableDefaults
	}
	if options.enabledSetters["Variables"] {
		params["variables"] = options.Variables
	}
	if options.enabledSetters["WhatsappCategory"] {
		params["whatsapp_category"] = options.WhatsappCategory
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
	
// TemplateDestroy any binding still naming this template's key will find
// nothing when its
// event next arrives. Audited under the KEY as well as the id: after the
// delete the id resolves to nothing, and "deleted tmpl_01J…" is not
// something an operator can act on six weeks later.
func (srv *Messaging) TemplateDestroy(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/templates/{id}")
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
	
// TemplateShow what customers are receiving is the published snapshot; see
// `GET /v1/templates/{id}/versions`, whose `meta.has_unpublished_changes`
// says whether the two differ.
// 
// Not market-filtered, deliberately: market scoping is a browsing concern
// and somebody holding an id may read the row.
func (srv *Messaging) TemplateShow(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/templates/{id}")
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
type TemplateUpdatePatchOptions struct {
	BodyHtml string
	BodyText string
	ContentSid string
	Design []string
	Enabled bool
	LayoutId string
	Markets []string
	MessageClass string
	Subject string
	TestMode bool
	Title string
	ValidFrom string
	ValidUntil string
	VariableDefaults []string
	Variables []string
	WhatsappCategory string
	enabledSetters map[string]bool
}
func (options TemplateUpdatePatchOptions) New() *TemplateUpdatePatchOptions {
	options.enabledSetters = map[string]bool{
		"BodyHtml": false,
		"BodyText": false,
		"ContentSid": false,
		"Design": false,
		"Enabled": false,
		"LayoutId": false,
		"Markets": false,
		"MessageClass": false,
		"Subject": false,
		"TestMode": false,
		"Title": false,
		"ValidFrom": false,
		"ValidUntil": false,
		"VariableDefaults": false,
		"Variables": false,
		"WhatsappCategory": false,
	}
	return &options
}
type TemplateUpdatePatchOption func(*TemplateUpdatePatchOptions)
func (srv *Messaging) WithTemplateUpdatePatchBodyHtml(v string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.BodyHtml = v
		o.enabledSetters["BodyHtml"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchBodyText(v string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.BodyText = v
		o.enabledSetters["BodyText"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchContentSid(v string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.ContentSid = v
		o.enabledSetters["ContentSid"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchDesign(v []string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.Design = v
		o.enabledSetters["Design"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchEnabled(v bool) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchLayoutId(v string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.LayoutId = v
		o.enabledSetters["LayoutId"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchMarkets(v []string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.Markets = v
		o.enabledSetters["Markets"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchMessageClass(v string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.MessageClass = v
		o.enabledSetters["MessageClass"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchSubject(v string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.Subject = v
		o.enabledSetters["Subject"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchTestMode(v bool) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.TestMode = v
		o.enabledSetters["TestMode"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchTitle(v string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchValidFrom(v string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.ValidFrom = v
		o.enabledSetters["ValidFrom"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchValidUntil(v string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.ValidUntil = v
		o.enabledSetters["ValidUntil"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchVariableDefaults(v []string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.VariableDefaults = v
		o.enabledSetters["VariableDefaults"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchVariables(v []string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.Variables = v
		o.enabledSetters["Variables"] = true
	}
}
func (srv *Messaging) WithTemplateUpdatePatchWhatsappCategory(v string) TemplateUpdatePatchOption {
	return func(o *TemplateUpdatePatchOptions) {
		o.WhatsappCategory = v
		o.enabledSetters["WhatsappCategory"] = true
	}
}
			
// TemplateUpdatePatch only the fields sent are written, and the change is
// audited only when
// something actually changed — a PATCH that resent the same values records
// nothing, because an audit line on every save teaches its readers to
// ignore the log.
// 
// Moving a template to another layout recompiles it against the NEW one,
// even when nothing else changed: colours, width and font come from the
// layout and are already inlined, so a template that merely changed hands
// would otherwise keep showing the old letterhead until somebody happened
// to press save on it again.
// 
// Changes nothing customers receive until the template is published.
// 
// This path answers on `PUT` and `PATCH`, both routed to the same action.
func (srv *Messaging) TemplateUpdatePatch(Id string, optionalSetters ...TemplateUpdatePatchOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/templates/{id}")
	options := TemplateUpdatePatchOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["BodyHtml"] {
		params["body_html"] = options.BodyHtml
	}
	if options.enabledSetters["BodyText"] {
		params["body_text"] = options.BodyText
	}
	if options.enabledSetters["ContentSid"] {
		params["content_sid"] = options.ContentSid
	}
	if options.enabledSetters["Design"] {
		params["design"] = options.Design
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["LayoutId"] {
		params["layout_id"] = options.LayoutId
	}
	if options.enabledSetters["Markets"] {
		params["markets"] = options.Markets
	}
	if options.enabledSetters["MessageClass"] {
		params["message_class"] = options.MessageClass
	}
	if options.enabledSetters["Subject"] {
		params["subject"] = options.Subject
	}
	if options.enabledSetters["TestMode"] {
		params["test_mode"] = options.TestMode
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["ValidFrom"] {
		params["valid_from"] = options.ValidFrom
	}
	if options.enabledSetters["ValidUntil"] {
		params["valid_until"] = options.ValidUntil
	}
	if options.enabledSetters["VariableDefaults"] {
		params["variable_defaults"] = options.VariableDefaults
	}
	if options.enabledSetters["Variables"] {
		params["variables"] = options.Variables
	}
	if options.enabledSetters["WhatsappCategory"] {
		params["whatsapp_category"] = options.WhatsappCategory
	}
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("PATCH", path, headers, params)
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
type TemplateUpdateOptions struct {
	BodyHtml string
	BodyText string
	ContentSid string
	Design []string
	Enabled bool
	LayoutId string
	Markets []string
	MessageClass string
	Subject string
	TestMode bool
	Title string
	ValidFrom string
	ValidUntil string
	VariableDefaults []string
	Variables []string
	WhatsappCategory string
	enabledSetters map[string]bool
}
func (options TemplateUpdateOptions) New() *TemplateUpdateOptions {
	options.enabledSetters = map[string]bool{
		"BodyHtml": false,
		"BodyText": false,
		"ContentSid": false,
		"Design": false,
		"Enabled": false,
		"LayoutId": false,
		"Markets": false,
		"MessageClass": false,
		"Subject": false,
		"TestMode": false,
		"Title": false,
		"ValidFrom": false,
		"ValidUntil": false,
		"VariableDefaults": false,
		"Variables": false,
		"WhatsappCategory": false,
	}
	return &options
}
type TemplateUpdateOption func(*TemplateUpdateOptions)
func (srv *Messaging) WithTemplateUpdateBodyHtml(v string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.BodyHtml = v
		o.enabledSetters["BodyHtml"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateBodyText(v string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.BodyText = v
		o.enabledSetters["BodyText"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateContentSid(v string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.ContentSid = v
		o.enabledSetters["ContentSid"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateDesign(v []string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.Design = v
		o.enabledSetters["Design"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateEnabled(v bool) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateLayoutId(v string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.LayoutId = v
		o.enabledSetters["LayoutId"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateMarkets(v []string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.Markets = v
		o.enabledSetters["Markets"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateMessageClass(v string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.MessageClass = v
		o.enabledSetters["MessageClass"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateSubject(v string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.Subject = v
		o.enabledSetters["Subject"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateTestMode(v bool) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.TestMode = v
		o.enabledSetters["TestMode"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateTitle(v string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateValidFrom(v string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.ValidFrom = v
		o.enabledSetters["ValidFrom"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateValidUntil(v string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.ValidUntil = v
		o.enabledSetters["ValidUntil"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateVariableDefaults(v []string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.VariableDefaults = v
		o.enabledSetters["VariableDefaults"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateVariables(v []string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.Variables = v
		o.enabledSetters["Variables"] = true
	}
}
func (srv *Messaging) WithTemplateUpdateWhatsappCategory(v string) TemplateUpdateOption {
	return func(o *TemplateUpdateOptions) {
		o.WhatsappCategory = v
		o.enabledSetters["WhatsappCategory"] = true
	}
}
			
// TemplateUpdate only the fields sent are written, and the change is audited
// only when
// something actually changed — a PATCH that resent the same values records
// nothing, because an audit line on every save teaches its readers to
// ignore the log.
// 
// Moving a template to another layout recompiles it against the NEW one,
// even when nothing else changed: colours, width and font come from the
// layout and are already inlined, so a template that merely changed hands
// would otherwise keep showing the old letterhead until somebody happened
// to press save on it again.
// 
// Changes nothing customers receive until the template is published.
// 
// This path answers on `PUT` and `PATCH`, both routed to the same action.
func (srv *Messaging) TemplateUpdate(Id string, optionalSetters ...TemplateUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/messaging/templates/{id}")
	options := TemplateUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["BodyHtml"] {
		params["body_html"] = options.BodyHtml
	}
	if options.enabledSetters["BodyText"] {
		params["body_text"] = options.BodyText
	}
	if options.enabledSetters["ContentSid"] {
		params["content_sid"] = options.ContentSid
	}
	if options.enabledSetters["Design"] {
		params["design"] = options.Design
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["LayoutId"] {
		params["layout_id"] = options.LayoutId
	}
	if options.enabledSetters["Markets"] {
		params["markets"] = options.Markets
	}
	if options.enabledSetters["MessageClass"] {
		params["message_class"] = options.MessageClass
	}
	if options.enabledSetters["Subject"] {
		params["subject"] = options.Subject
	}
	if options.enabledSetters["TestMode"] {
		params["test_mode"] = options.TestMode
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["ValidFrom"] {
		params["valid_from"] = options.ValidFrom
	}
	if options.enabledSetters["ValidUntil"] {
		params["valid_until"] = options.ValidUntil
	}
	if options.enabledSetters["VariableDefaults"] {
		params["variable_defaults"] = options.VariableDefaults
	}
	if options.enabledSetters["Variables"] {
		params["variables"] = options.Variables
	}
	if options.enabledSetters["WhatsappCategory"] {
		params["whatsapp_category"] = options.WhatsappCategory
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
type TemplateVersionStoreOptions struct {
	Note string
	enabledSetters map[string]bool
}
func (options TemplateVersionStoreOptions) New() *TemplateVersionStoreOptions {
	options.enabledSetters = map[string]bool{
		"Note": false,
	}
	return &options
}
type TemplateVersionStoreOption func(*TemplateVersionStoreOptions)
func (srv *Messaging) WithTemplateVersionStoreNote(v string) TemplateVersionStoreOption {
	return func(o *TemplateVersionStoreOptions) {
		o.Note = v
		o.enabledSetters["Note"] = true
	}
}
			
// TemplateVersionStore answers 200 with the version already live when there
// was nothing to
// publish, and 201 when a new one was written — so a client can tell
// whether its press did anything without diffing the payload.
func (srv *Messaging) TemplateVersionStore(TemplateId string, optionalSetters ...TemplateVersionStoreOption)(*models.Error, error) {
	r := strings.NewReplacer("{templateId}", TemplateId)
	path := r.Replace("/v1/messaging/templates/{templateId}/publish")
	options := TemplateVersionStoreOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["templateId"] = TemplateId
	if options.enabledSetters["Note"] {
		params["note"] = options.Note
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
	
// TemplateVersionIndex summaries only: version, subject, message class,
// layout, who published it
// and when, and their note. The BODIES are deliberately absent — a compiled
// `body_html` runs to tens of kilobytes, and a template with forty versions
// would make this a several-megabyte download that nobody scrolls to the
// end of. `GET /v1/templates/{id}/versions/{version}` serves the full
// snapshot for the one somebody actually opened.
// 
// `meta.published_version_id` says which of them is live — a property of
// the template, said once, rather than a flag repeated on every row that
// two rows could then claim. `meta.has_unpublished_changes` says whether
// the draft has moved on since.
func (srv *Messaging) TemplateVersionIndex(TemplateId string)(*models.Error, error) {
	r := strings.NewReplacer("{templateId}", TemplateId)
	path := r.Replace("/v1/messaging/templates/{templateId}/versions")
	params := map[string]interface{}{}
	params["templateId"] = TemplateId
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
			
// TemplateVersionShow addressed by its VERSION NUMBER — the small integer
// on the history row,
// not the snapshot's id — because that is the number an author has in front
// of them.
// 
// This is what sends actually rendered while that version was live, so it
// is the thing to read when the question is "what did the mail we sent in
// * March say".
func (srv *Messaging) TemplateVersionShow(TemplateId string, Version string)(*models.Error, error) {
	r := strings.NewReplacer("{templateId}", TemplateId, "{version}", Version)
	path := r.Replace("/v1/messaging/templates/{templateId}/versions/{version}")
	params := map[string]interface{}{}
	params["templateId"] = TemplateId
	params["version"] = Version
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
type TemplateVersionRestoreOptions struct {
	Publish bool
	enabledSetters map[string]bool
}
func (options TemplateVersionRestoreOptions) New() *TemplateVersionRestoreOptions {
	options.enabledSetters = map[string]bool{
		"Publish": false,
	}
	return &options
}
type TemplateVersionRestoreOption func(*TemplateVersionRestoreOptions)
func (srv *Messaging) WithTemplateVersionRestorePublish(v bool) TemplateVersionRestoreOption {
	return func(o *TemplateVersionRestoreOptions) {
		o.Publish = v
		o.enabledSetters["Publish"] = true
	}
}
					
// TemplateVersionRestore `publish: true` makes it live in the same
// transaction — see
// TemplatePublisher::restore for why that flag exists rather than asking
// the caller for a second round trip.
func (srv *Messaging) TemplateVersionRestore(TemplateId string, Version string, optionalSetters ...TemplateVersionRestoreOption)(*models.Error, error) {
	r := strings.NewReplacer("{templateId}", TemplateId, "{version}", Version)
	path := r.Replace("/v1/messaging/templates/{templateId}/versions/{version}/restore")
	options := TemplateVersionRestoreOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["templateId"] = TemplateId
	params["version"] = Version
	if options.enabledSetters["Publish"] {
		params["publish"] = options.Publish
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
