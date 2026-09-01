package customers_segments

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// CustomersSegments service
type CustomersSegments struct {
	client client.Client
}

func New(clt client.Client) *CustomersSegments {
	return &CustomersSegments{
		client: clt,
	}
}

type CustomersSegmentMembersListOptions struct {
	Id string
	SegmentId string
	OrganizationId string
	Source string
	CreatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options CustomersSegmentMembersListOptions) New() *CustomersSegmentMembersListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"SegmentId": false,
		"OrganizationId": false,
		"Source": false,
		"CreatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type CustomersSegmentMembersListOption func(*CustomersSegmentMembersListOptions)
func (srv *CustomersSegments) WithCustomersSegmentMembersListId(v string) CustomersSegmentMembersListOption {
	return func(o *CustomersSegmentMembersListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentMembersListSegmentId(v string) CustomersSegmentMembersListOption {
	return func(o *CustomersSegmentMembersListOptions) {
		o.SegmentId = v
		o.enabledSetters["SegmentId"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentMembersListOrganizationId(v string) CustomersSegmentMembersListOption {
	return func(o *CustomersSegmentMembersListOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentMembersListSource(v string) CustomersSegmentMembersListOption {
	return func(o *CustomersSegmentMembersListOptions) {
		o.Source = v
		o.enabledSetters["Source"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentMembersListCreatedAt(v string) CustomersSegmentMembersListOption {
	return func(o *CustomersSegmentMembersListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentMembersListLimit(v int) CustomersSegmentMembersListOption {
	return func(o *CustomersSegmentMembersListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentMembersListOffset(v int) CustomersSegmentMembersListOption {
	return func(o *CustomersSegmentMembersListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentMembersListOrder(v string) CustomersSegmentMembersListOption {
	return func(o *CustomersSegmentMembersListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// CustomersSegmentMembersList one organization inside one segment, plus the
// record of how it got there: `source: "manual"` for a company somebody put
// in, `source: "rule"` for one the rule engine matched. That distinction is
// what lets a recompute rewrite its own rows and leave every hand-picked one
// alone. The membership rows themselves — the answer to "which companies
// are in this segment" (`segment_id`) and to "which segments is this company
// in" (`organization_id`). Paged with `limit`/`offset`/`order`.
func (srv *CustomersSegments) CustomersSegmentMembersList(optionalSetters ...CustomersSegmentMembersListOption)(*interface{}, error) {
	path := "/v1/customers/segment_members"
	options := CustomersSegmentMembersListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Id"] {
		params["id"] = options.Id
	}
	if options.enabledSetters["SegmentId"] {
		params["segment_id"] = options.SegmentId
	}
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["Source"] {
		params["source"] = options.Source
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
type CustomersSegmentMembersCreateOptions struct {
	Source string
	enabledSetters map[string]bool
}
func (options CustomersSegmentMembersCreateOptions) New() *CustomersSegmentMembersCreateOptions {
	options.enabledSetters = map[string]bool{
		"Source": false,
	}
	return &options
}
type CustomersSegmentMembersCreateOption func(*CustomersSegmentMembersCreateOptions)
func (srv *CustomersSegments) WithCustomersSegmentMembersCreateSource(v string) CustomersSegmentMembersCreateOption {
	return func(o *CustomersSegmentMembersCreateOptions) {
		o.Source = v
		o.enabledSetters["Source"] = true
	}
}
					
// CustomersSegmentMembersCreate one organization inside one segment, plus the
// record of how it got there: `source: "manual"` for a company somebody put
// in, `source: "rule"` for one the rule engine matched. That distinction is
// what lets a recompute rewrite its own rows and leave every hand-picked one
// alone. Adds a company to a segment BY HAND. The row is `source: "manual"`,
// which is what protects it: a rule recompute rewrites the rule-derived rows
// of that segment and never touches this one. A create cannot omit
// `segment_id` and `organization_id`; everything else is optional or
// defaulted by the database. Two rows of this tenant may not share the
// combination of `segment_id` + `organization_id`.
func (srv *CustomersSegments) CustomersSegmentMembersCreate(OrganizationId string, SegmentId string, optionalSetters ...CustomersSegmentMembersCreateOption)(*models.Error, error) {
	path := "/v1/customers/segment_members"
	options := CustomersSegmentMembersCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["organization_id"] = OrganizationId
	params["segment_id"] = SegmentId
	if options.enabledSetters["Source"] {
		params["source"] = options.Source
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
	
// CustomersSegmentMembersDelete one organization inside one segment, plus the
// record of how it got there: `source: "manual"` for a company somebody put
// in, `source: "rule"` for one the rule engine matched. That distinction is
// what lets a recompute rewrite its own rows and leave every hand-picked one
// alone. Takes the company out of the segment. If the segment carries rules
// and the company still matches them, the next recompute puts it back; remove
// it from the rule, not from the list. Nothing else in this app points at it,
// so nothing else goes with it.
func (srv *CustomersSegments) CustomersSegmentMembersDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/segment_members/{id}")
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
	
// CustomersSegmentMembersGet one organization inside one segment, plus the
// record of how it got there: `source: "manual"` for a company somebody put
// in, `source: "rule"` for one the rule engine matched. That distinction is
// what lets a recompute rewrite its own rows and leave every hand-picked one
// alone. One membership row by id, with the `source` that says how it came
// about.
func (srv *CustomersSegments) CustomersSegmentMembersGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/segment_members/{id}")
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
type CustomersSegmentMembersUpdateOptions struct {
	OrganizationId string
	SegmentId string
	Source string
	enabledSetters map[string]bool
}
func (options CustomersSegmentMembersUpdateOptions) New() *CustomersSegmentMembersUpdateOptions {
	options.enabledSetters = map[string]bool{
		"OrganizationId": false,
		"SegmentId": false,
		"Source": false,
	}
	return &options
}
type CustomersSegmentMembersUpdateOption func(*CustomersSegmentMembersUpdateOptions)
func (srv *CustomersSegments) WithCustomersSegmentMembersUpdateOrganizationId(v string) CustomersSegmentMembersUpdateOption {
	return func(o *CustomersSegmentMembersUpdateOptions) {
		o.OrganizationId = v
		o.enabledSetters["OrganizationId"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentMembersUpdateSegmentId(v string) CustomersSegmentMembersUpdateOption {
	return func(o *CustomersSegmentMembersUpdateOptions) {
		o.SegmentId = v
		o.enabledSetters["SegmentId"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentMembersUpdateSource(v string) CustomersSegmentMembersUpdateOption {
	return func(o *CustomersSegmentMembersUpdateOptions) {
		o.Source = v
		o.enabledSetters["Source"] = true
	}
}
			
// CustomersSegmentMembersUpdate one organization inside one segment, plus the
// record of how it got there: `source: "manual"` for a company somebody put
// in, `source: "rule"` for one the rule engine matched. That distinction is
// what lets a recompute rewrite its own rows and leave every hand-picked one
// alone. A partial update. In practice there is little to change — a
// membership is a pair of ids — so this exists for the `source` correction
// rather than as the normal path. Two rows of this tenant may not share the
// combination of `segment_id` + `organization_id`.
func (srv *CustomersSegments) CustomersSegmentMembersUpdate(Id string, optionalSetters ...CustomersSegmentMembersUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/segment_members/{id}")
	options := CustomersSegmentMembersUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["OrganizationId"] {
		params["organization_id"] = options.OrganizationId
	}
	if options.enabledSetters["SegmentId"] {
		params["segment_id"] = options.SegmentId
	}
	if options.enabledSetters["Source"] {
		params["source"] = options.Source
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
type CustomersSegmentsListOptions struct {
	Id string
	Code string
	Position int
	RuleMatch string
	RulesComputedAt string
	CreatedAt string
	UpdatedAt string
	Limit int
	Offset int
	Order string
	enabledSetters map[string]bool
}
func (options CustomersSegmentsListOptions) New() *CustomersSegmentsListOptions {
	options.enabledSetters = map[string]bool{
		"Id": false,
		"Code": false,
		"Position": false,
		"RuleMatch": false,
		"RulesComputedAt": false,
		"CreatedAt": false,
		"UpdatedAt": false,
		"Limit": false,
		"Offset": false,
		"Order": false,
	}
	return &options
}
type CustomersSegmentsListOption func(*CustomersSegmentsListOptions)
func (srv *CustomersSegments) WithCustomersSegmentsListId(v string) CustomersSegmentsListOption {
	return func(o *CustomersSegmentsListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsListCode(v string) CustomersSegmentsListOption {
	return func(o *CustomersSegmentsListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsListPosition(v int) CustomersSegmentsListOption {
	return func(o *CustomersSegmentsListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsListRuleMatch(v string) CustomersSegmentsListOption {
	return func(o *CustomersSegmentsListOptions) {
		o.RuleMatch = v
		o.enabledSetters["RuleMatch"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsListRulesComputedAt(v string) CustomersSegmentsListOption {
	return func(o *CustomersSegmentsListOptions) {
		o.RulesComputedAt = v
		o.enabledSetters["RulesComputedAt"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsListCreatedAt(v string) CustomersSegmentsListOption {
	return func(o *CustomersSegmentsListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsListUpdatedAt(v string) CustomersSegmentsListOption {
	return func(o *CustomersSegmentsListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsListLimit(v int) CustomersSegmentsListOption {
	return func(o *CustomersSegmentsListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsListOffset(v int) CustomersSegmentsListOption {
	return func(o *CustomersSegmentsListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsListOrder(v string) CustomersSegmentsListOption {
	return func(o *CustomersSegmentsListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
	
// CustomersSegmentsList a segment is a named group of ORGANIZATIONS — never
// of people — built by hand, by rule, or both at once. It is what a price
// list, a campaign or a shipping option is pointed at when the answer is
// "these customers, not those". Every segment this tenant keeps, with its
// stored rules. Any column filters and the page is `limit`/`offset`/`order`.
// Which companies are actually IN one is `segment_members`, because the rule
// half is materialized rather than evaluated on read.
func (srv *CustomersSegments) CustomersSegmentsList(optionalSetters ...CustomersSegmentsListOption)(*interface{}, error) {
	path := "/v1/customers/segments"
	options := CustomersSegmentsListOptions{}.New()
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
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["RuleMatch"] {
		params["rule_match"] = options.RuleMatch
	}
	if options.enabledSetters["RulesComputedAt"] {
		params["rules_computed_at"] = options.RulesComputedAt
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
type CustomersSegmentsCreateOptions struct {
	Labels interface{}
	Position int
	RuleMatch string
	Rules interface{}
	enabledSetters map[string]bool
}
func (options CustomersSegmentsCreateOptions) New() *CustomersSegmentsCreateOptions {
	options.enabledSetters = map[string]bool{
		"Labels": false,
		"Position": false,
		"RuleMatch": false,
		"Rules": false,
	}
	return &options
}
type CustomersSegmentsCreateOption func(*CustomersSegmentsCreateOptions)
func (srv *CustomersSegments) WithCustomersSegmentsCreateLabels(v interface{}) CustomersSegmentsCreateOption {
	return func(o *CustomersSegmentsCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsCreatePosition(v int) CustomersSegmentsCreateOption {
	return func(o *CustomersSegmentsCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsCreateRuleMatch(v string) CustomersSegmentsCreateOption {
	return func(o *CustomersSegmentsCreateOptions) {
		o.RuleMatch = v
		o.enabledSetters["RuleMatch"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsCreateRules(v interface{}) CustomersSegmentsCreateOption {
	return func(o *CustomersSegmentsCreateOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
			
// CustomersSegmentsCreate a segment is a named group of ORGANIZATIONS —
// never of people — built by hand, by rule, or both at once. It is what a
// price list, a campaign or a shipping option is pointed at when the answer
// is "these customers, not those". Creates the group. Rules are optional:
// leave them out for a hand-picked list, or store a rule document and let the
// recompute keep the membership up to date. The `code` is what other apps
// point at, so pick it deliberately. `code` is the only field a create cannot
// omit; everything else is optional or defaulted by the database. Two rows of
// this tenant may not share `code`.
func (srv *CustomersSegments) CustomersSegmentsCreate(Code string, optionalSetters ...CustomersSegmentsCreateOption)(*models.Error, error) {
	path := "/v1/customers/segments"
	options := CustomersSegmentsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["RuleMatch"] {
		params["rule_match"] = options.RuleMatch
	}
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
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
	
// CustomersSegmentsRulesRecomputeAll same sync as the single-segment
// recompute, applied to every segment with non-null rules. A failing segment
// is reported in its result entry instead of aborting the run. The run shares
// one budget: a segment that does not fit reports done:false (or
// skipped:true) and keeps rules_computed_at null, so the next call resumes it
// from its own data. Repeat until the top-level done is true.
func (srv *CustomersSegments) CustomersSegmentsRulesRecomputeAll(Data interface{})(*models.Error, error) {
	path := "/v1/customers/segments/rules/recompute-all"
	params := map[string]interface{}{}
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
	
// CustomersSegmentsDelete a segment is a named group of ORGANIZATIONS —
// never of people — built by hand, by rule, or both at once. It is what a
// price list, a campaign or a shipping option is pointed at when the answer
// is "these customers, not those". Removes the segment. Anything in another
// app that points at its `code` — a price list, a campaign — is left
// pointing at nothing, because no app may hold a foreign key into another
// (ADR-0055). Deleting one takes every `segment_members` row that points at
// it with it — the foreign keys decide, not this route.
func (srv *CustomersSegments) CustomersSegmentsDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/segments/{id}")
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
	
// CustomersSegmentsGet a segment is a named group of ORGANIZATIONS — never
// of people — built by hand, by rule, or both at once. It is what a price
// list, a campaign or a shipping option is pointed at when the answer is
// "these customers, not those". One segment by id, including the rule
// document it carries. A segment with no rules is hand-picked and completely
// valid.
func (srv *CustomersSegments) CustomersSegmentsGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/segments/{id}")
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
type CustomersSegmentsUpdateOptions struct {
	Code string
	Labels interface{}
	Position int
	RuleMatch string
	Rules interface{}
	enabledSetters map[string]bool
}
func (options CustomersSegmentsUpdateOptions) New() *CustomersSegmentsUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Labels": false,
		"Position": false,
		"RuleMatch": false,
		"Rules": false,
	}
	return &options
}
type CustomersSegmentsUpdateOption func(*CustomersSegmentsUpdateOptions)
func (srv *CustomersSegments) WithCustomersSegmentsUpdateCode(v string) CustomersSegmentsUpdateOption {
	return func(o *CustomersSegmentsUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsUpdateLabels(v interface{}) CustomersSegmentsUpdateOption {
	return func(o *CustomersSegmentsUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsUpdatePosition(v int) CustomersSegmentsUpdateOption {
	return func(o *CustomersSegmentsUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsUpdateRuleMatch(v string) CustomersSegmentsUpdateOption {
	return func(o *CustomersSegmentsUpdateOptions) {
		o.RuleMatch = v
		o.enabledSetters["RuleMatch"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsUpdateRules(v interface{}) CustomersSegmentsUpdateOption {
	return func(o *CustomersSegmentsUpdateOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
			
// CustomersSegmentsUpdate a segment is a named group of ORGANIZATIONS —
// never of people — built by hand, by rule, or both at once. It is what a
// price list, a campaign or a shipping option is pointed at when the answer
// is "these customers, not those". A partial update — send only what
// changes. Editing the rules does NOT re-evaluate them: that is `POST
// /customers/segments/{segment_id}/rules/recompute`, so a half-typed rule
// never silently empties a live segment. Two rows of this tenant may not
// share `code`.
func (srv *CustomersSegments) CustomersSegmentsUpdate(Id string, optionalSetters ...CustomersSegmentsUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/customers/segments/{id}")
	options := CustomersSegmentsUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["RuleMatch"] {
		params["rule_match"] = options.RuleMatch
	}
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
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
type CustomersSegmentsRulesPreviewOptions struct {
	RuleMatch string
	Target string
	enabledSetters map[string]bool
}
func (options CustomersSegmentsRulesPreviewOptions) New() *CustomersSegmentsRulesPreviewOptions {
	options.enabledSetters = map[string]bool{
		"RuleMatch": false,
		"Target": false,
	}
	return &options
}
type CustomersSegmentsRulesPreviewOption func(*CustomersSegmentsRulesPreviewOptions)
func (srv *CustomersSegments) WithCustomersSegmentsRulesPreviewRuleMatch(v string) CustomersSegmentsRulesPreviewOption {
	return func(o *CustomersSegmentsRulesPreviewOptions) {
		o.RuleMatch = v
		o.enabledSetters["RuleMatch"] = true
	}
}
func (srv *CustomersSegments) WithCustomersSegmentsRulesPreviewTarget(v string) CustomersSegmentsRulesPreviewOption {
	return func(o *CustomersSegmentsRulesPreviewOptions) {
		o.Target = v
		o.enabledSetters["Target"] = true
	}
}
					
// CustomersSegmentsRulesPreview a dry run: it answers how many organizations
// the rule would select, with a handful of them by name, and writes nothing
// at all. Evaluates the rule document in the REQUEST BODY (not the stored
// segments.rules), so the cockpit can preview an unsaved rule. Costs a single
// count query for the common single-query rule; 'any' rules and rules
// repeating a column are combined in the app and capped at 5000 ids, in which
// case 'capped' is true and 'count' is a LOWER bound. Membership is never
// touched.
func (srv *CustomersSegments) CustomersSegmentsRulesPreview(SegmentId string, Conditions []models.SegmentRuleCondition, optionalSetters ...CustomersSegmentsRulesPreviewOption)(*models.Error, error) {
	r := strings.NewReplacer("{segment_id}", SegmentId)
	path := r.Replace("/v1/customers/segments/{segment_id}/rules/preview")
	options := CustomersSegmentsRulesPreviewOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["segment_id"] = SegmentId
	params["conditions"] = Conditions
	if options.enabledSetters["RuleMatch"] {
		params["rule_match"] = options.RuleMatch
	}
	if options.enabledSetters["Target"] {
		params["target"] = options.Target
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
type CustomersSegmentsRulesRecomputeOptions struct {
	Cursor string
	enabledSetters map[string]bool
}
func (options CustomersSegmentsRulesRecomputeOptions) New() *CustomersSegmentsRulesRecomputeOptions {
	options.enabledSetters = map[string]bool{
		"Cursor": false,
	}
	return &options
}
type CustomersSegmentsRulesRecomputeOption func(*CustomersSegmentsRulesRecomputeOptions)
func (srv *CustomersSegments) WithCustomersSegmentsRulesRecomputeCursor(v string) CustomersSegmentsRulesRecomputeOption {
	return func(o *CustomersSegmentsRulesRecomputeOptions) {
		o.Cursor = v
		o.enabledSetters["Cursor"] = true
	}
}
			
// CustomersSegmentsRulesRecompute evaluates segments.rules (NOT the request
// body), then inserts the newly matching organizations as source='rule' rows
// and deletes the rule rows that no longer match. Manual (source='manual')
// memberships are never inserted, deleted or shadowed. Bounded by a
// wall-clock budget below the gateway's upstream timeout: when 'done' is
// false, POST again with the returned 'cursor' until it is true.
// added/removed/processed count THIS call only. Omitting 'cursor' resumes an
// unfinished pass and starts a fresh one after a completed pass; an explicit
// null always restarts. segments.rules_computed_at is stamped only when the
// pass completes.
func (srv *CustomersSegments) CustomersSegmentsRulesRecompute(SegmentId string, optionalSetters ...CustomersSegmentsRulesRecomputeOption)(*models.Error, error) {
	r := strings.NewReplacer("{segment_id}", SegmentId)
	path := r.Replace("/v1/customers/segments/{segment_id}/rules/recompute")
	options := CustomersSegmentsRulesRecomputeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["segment_id"] = SegmentId
	if options.enabledSetters["Cursor"] {
		params["cursor"] = options.Cursor
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
