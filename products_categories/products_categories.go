package products_categories

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// ProductsCategories service
type ProductsCategories struct {
	client client.Client
}

func New(clt client.Client) *ProductsCategories {
	return &ProductsCategories{
		client: clt,
	}
}

type ProductsCategoriesListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	Code string
	ParentId string
	Path string
	Position int
	Labels string
	Values string
	Rules string
	RuleMatch string
	RulesComputedAt string
	CreatedAt string
	UpdatedAt string
	enabledSetters map[string]bool
}
func (options ProductsCategoriesListOptions) New() *ProductsCategoriesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"Code": false,
		"ParentId": false,
		"Path": false,
		"Position": false,
		"Labels": false,
		"Values": false,
		"Rules": false,
		"RuleMatch": false,
		"RulesComputedAt": false,
		"CreatedAt": false,
		"UpdatedAt": false,
	}
	return &options
}
type ProductsCategoriesListOption func(*ProductsCategoriesListOptions)
func (srv *ProductsCategories) WithProductsCategoriesListLimit(v int) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListOffset(v int) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListOrder(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListId(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListCode(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListParentId(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.ParentId = v
		o.enabledSetters["ParentId"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListPath(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.Path = v
		o.enabledSetters["Path"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListPosition(v int) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListLabels(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListValues(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.Values = v
		o.enabledSetters["Values"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListRules(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListRuleMatch(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.RuleMatch = v
		o.enabledSetters["RuleMatch"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListRulesComputedAt(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.RulesComputedAt = v
		o.enabledSetters["RulesComputedAt"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListCreatedAt(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesListUpdatedAt(v string) ProductsCategoriesListOption {
	return func(o *ProductsCategoriesListOptions) {
		o.UpdatedAt = v
		o.enabledSetters["UpdatedAt"] = true
	}
}
	
// ProductsCategoriesList one node of the category tree. `parent_id` is the
// structure this app navigates — null is a root — while `path` is kept
// only for importers that carry one and nothing here reads or writes it. A
// category is hand-picked or RULE-DRIVEN: a non-null `rules` selector makes
// every matching product a `product_categories` row with source `rule`,
// alongside the hand-picked ones, and `rules_computed_at` says when that last
// completed.
// 
// Every column of `categories` is an exact-match query parameter, `order`
// sorts by one column, and `limit`/`offset` page through `page.total`. A
// query key that is NOT a column is dropped rather than refused, and the
// `filter` object echoes the ones that were understood — that echo is the
// only way to tell an unfiltered answer from an empty one. It reads rows
// exactly as they are stored: no join is resolved, no jsonb value is
// unpacked.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsCategories) ProductsCategoriesList(optionalSetters ...ProductsCategoriesListOption)(*interface{}, error) {
	path := "/v1/products/categories"
	options := ProductsCategoriesListOptions{}.New()
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
	if options.enabledSetters["Code"] {
		params["code"] = options.Code
	}
	if options.enabledSetters["ParentId"] {
		params["parent_id"] = options.ParentId
	}
	if options.enabledSetters["Path"] {
		params["path"] = options.Path
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["Values"] {
		params["values"] = options.Values
	}
	if options.enabledSetters["Rules"] {
		params["rules"] = options.Rules
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
type ProductsCategoriesCreateOptions struct {
	Labels interface{}
	ParentId string
	Path string
	Position int
	RuleMatch string
	Rules interface{}
	RulesComputedAt string
	Values interface{}
	enabledSetters map[string]bool
}
func (options ProductsCategoriesCreateOptions) New() *ProductsCategoriesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Labels": false,
		"ParentId": false,
		"Path": false,
		"Position": false,
		"RuleMatch": false,
		"Rules": false,
		"RulesComputedAt": false,
		"Values": false,
	}
	return &options
}
type ProductsCategoriesCreateOption func(*ProductsCategoriesCreateOptions)
func (srv *ProductsCategories) WithProductsCategoriesCreateLabels(v interface{}) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesCreateParentId(v string) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.ParentId = v
		o.enabledSetters["ParentId"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesCreatePath(v string) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.Path = v
		o.enabledSetters["Path"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesCreatePosition(v int) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesCreateRuleMatch(v string) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.RuleMatch = v
		o.enabledSetters["RuleMatch"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesCreateRules(v interface{}) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesCreateRulesComputedAt(v string) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.RulesComputedAt = v
		o.enabledSetters["RulesComputedAt"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesCreateValues(v interface{}) ProductsCategoriesCreateOption {
	return func(o *ProductsCategoriesCreateOptions) {
		o.Values = v
		o.enabledSetters["Values"] = true
	}
}
			
// ProductsCategoriesCreate creates one category and answers 201 with the
// stored row, including the id and the timestamps the database filled in —
// a client never sends an id, it reads one back and uses it in the path of
// every later call.
// 
// One node of the category tree. `parent_id` is the structure this app
// navigates — null is a root — while `path` is kept only for importers
// that carry one and nothing here reads or writes it. A category is
// hand-picked or RULE-DRIVEN: a non-null `rules` selector makes every
// matching product a `product_categories` row with source `rule`, alongside
// the hand-picked ones, and `rules_computed_at` says when that last
// completed.
// 
// `code` is the only column the database refuses the row without; everything
// else has a default or is nullable. A second row with the same `code`
// answers 409.
func (srv *ProductsCategories) ProductsCategoriesCreate(Code string, optionalSetters ...ProductsCategoriesCreateOption)(*models.Error, error) {
	path := "/v1/products/categories"
	options := ProductsCategoriesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["Labels"] {
		params["labels"] = options.Labels
	}
	if options.enabledSetters["ParentId"] {
		params["parent_id"] = options.ParentId
	}
	if options.enabledSetters["Path"] {
		params["path"] = options.Path
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
	if options.enabledSetters["RulesComputedAt"] {
		params["rules_computed_at"] = options.RulesComputedAt
	}
	if options.enabledSetters["Values"] {
		params["values"] = options.Values
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
	
// ProductsCategoriesRulesRecomputeAll what the nightly
// `recompute-category-rules` schedule calls, and the call to reach for after
// a bulk import has changed what the rules select. Same sync as the
// single-category recompute, applied to every category with non-null rules.
// The whole run shares ONE budget: a category the budget no longer reaches is
// reported as `skipped` and picked up by the next run, and a failing category
// is reported in its result entry instead of aborting the run.
func (srv *ProductsCategories) ProductsCategoriesRulesRecomputeAll(Data interface{})(*models.Error, error) {
	path := "/v1/products/categories/rules/recompute-all"
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
type ProductsCategoriesRulesPreviewOptions struct {
	RuleMatch string
	enabledSetters map[string]bool
}
func (options ProductsCategoriesRulesPreviewOptions) New() *ProductsCategoriesRulesPreviewOptions {
	options.enabledSetters = map[string]bool{
		"RuleMatch": false,
	}
	return &options
}
type ProductsCategoriesRulesPreviewOption func(*ProductsCategoriesRulesPreviewOptions)
func (srv *ProductsCategories) WithProductsCategoriesRulesPreviewRuleMatch(v string) ProductsCategoriesRulesPreviewOption {
	return func(o *ProductsCategoriesRulesPreviewOptions) {
		o.RuleMatch = v
		o.enabledSetters["RuleMatch"] = true
	}
}
					
// ProductsCategoriesRulesPreview dry-runs a rule: how many products it
// selects, plus a sample of up to ten, and it WRITES NOTHING. Evaluates the
// rule in the request body against the live catalog WITHOUT touching
// product_categories — this powers the cockpit's "matches N products"
// preview while an operator edits a rule. Soft-deleted products are excluded.
// Counting is delegated to the database, never enumerated: a rule that
// compiles to a single query is answered by one exact-count request whatever
// its match set. A rule that needs several queries (rule_match "any", or a
// repeated column such as a range) is combined in the app and stops at `cap`
// ids — check `capped` before showing `count` as a total.
func (srv *ProductsCategories) ProductsCategoriesRulesPreview(CategoryId string, Conditions []models.CategoryRuleCondition, optionalSetters ...ProductsCategoriesRulesPreviewOption)(*models.Error, error) {
	r := strings.NewReplacer("{category_id}", CategoryId)
	path := r.Replace("/v1/products/categories/{category_id}/rules/preview")
	options := ProductsCategoriesRulesPreviewOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["category_id"] = CategoryId
	params["conditions"] = Conditions
	if options.enabledSetters["RuleMatch"] {
		params["rule_match"] = options.RuleMatch
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
type ProductsCategoriesRulesRecomputeOptions struct {
	Cursor string
	enabledSetters map[string]bool
}
func (options ProductsCategoriesRulesRecomputeOptions) New() *ProductsCategoriesRulesRecomputeOptions {
	options.enabledSetters = map[string]bool{
		"Cursor": false,
	}
	return &options
}
type ProductsCategoriesRulesRecomputeOption func(*ProductsCategoriesRulesRecomputeOptions)
func (srv *ProductsCategories) WithProductsCategoriesRulesRecomputeCursor(v string) ProductsCategoriesRulesRecomputeOption {
	return func(o *ProductsCategoriesRulesRecomputeOptions) {
		o.Cursor = v
		o.enabledSetters["Cursor"] = true
	}
}
			
// ProductsCategoriesRulesRecompute syncs one category's rule-derived
// memberships to what its stored rule selects today. Evaluates
// categories.rules (NOT the request body), then inserts the newly matching
// products as source='rule' rows and deletes the rule rows that no longer
// match. Manual (source='manual') memberships are never inserted, deleted or
// shadowed. Stamps categories.rules_computed_at.
// 
// A large category does NOT finish in one call: the run stops when its
// wall-clock budget is spent and answers `done: false` with the `cursor` to
// send back, so drive it in a loop until `done` is true.
func (srv *ProductsCategories) ProductsCategoriesRulesRecompute(CategoryId string, optionalSetters ...ProductsCategoriesRulesRecomputeOption)(*models.Error, error) {
	r := strings.NewReplacer("{category_id}", CategoryId)
	path := r.Replace("/v1/products/categories/{category_id}/rules/recompute")
	options := ProductsCategoriesRulesRecomputeOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["category_id"] = CategoryId
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
	
// ProductsCategoriesDelete deletes one category by id. It is a hard delete
// — the row is gone, and the answer is a confirmation rather than a result
// to branch on.
// 
// It takes what hangs off it: product category memberships (`category_id`)
// are deleted with it. `categories.parent_id` is set to null instead, so the
// rows that pointed at it survive the delete rather than going with it.
// 
// An id no category of this tenant carries answers 404; there is no 409,
// because every foreign key pointing at this entity resolves itself on delete
// rather than blocking one.
func (srv *ProductsCategories) ProductsCategoriesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/categories/{id}")
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
	
// ProductsCategoriesGet reads one category by its id — the whole row, every
// column, as it is stored.
// 
// One node of the category tree. `parent_id` is the structure this app
// navigates — null is a root — while `path` is kept only for importers
// that carry one and nothing here reads or writes it. A category is
// hand-picked or RULE-DRIVEN: a non-null `rules` selector makes every
// matching product a `product_categories` row with source `rule`, alongside
// the hand-picked ones, and `rules_computed_at` says when that last
// completed.
// 
// An id no category of this tenant carries answers 404, and so does one
// belonging to another tenant: row-level security makes that row invisible
// rather than forbidden. A malformed id answers 400 before the route is
// reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsCategories) ProductsCategoriesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/categories/{id}")
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
type ProductsCategoriesUpdateOptions struct {
	Code string
	Labels interface{}
	ParentId string
	Path string
	Position int
	RuleMatch string
	Rules interface{}
	RulesComputedAt string
	Values interface{}
	enabledSetters map[string]bool
}
func (options ProductsCategoriesUpdateOptions) New() *ProductsCategoriesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Code": false,
		"Labels": false,
		"ParentId": false,
		"Path": false,
		"Position": false,
		"RuleMatch": false,
		"Rules": false,
		"RulesComputedAt": false,
		"Values": false,
	}
	return &options
}
type ProductsCategoriesUpdateOption func(*ProductsCategoriesUpdateOptions)
func (srv *ProductsCategories) WithProductsCategoriesUpdateCode(v string) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.Code = v
		o.enabledSetters["Code"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesUpdateLabels(v interface{}) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.Labels = v
		o.enabledSetters["Labels"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesUpdateParentId(v string) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.ParentId = v
		o.enabledSetters["ParentId"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesUpdatePath(v string) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.Path = v
		o.enabledSetters["Path"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesUpdatePosition(v int) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesUpdateRuleMatch(v string) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.RuleMatch = v
		o.enabledSetters["RuleMatch"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesUpdateRules(v interface{}) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.Rules = v
		o.enabledSetters["Rules"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesUpdateRulesComputedAt(v string) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.RulesComputedAt = v
		o.enabledSetters["RulesComputedAt"] = true
	}
}
func (srv *ProductsCategories) WithProductsCategoriesUpdateValues(v interface{}) ProductsCategoriesUpdateOption {
	return func(o *ProductsCategoriesUpdateOptions) {
		o.Values = v
		o.enabledSetters["Values"] = true
	}
}
			
// ProductsCategoriesUpdate updates one category by id. A partial patch: the
// body names only the columns to change and every column it leaves out keeps
// its current value, so there is no read-modify-write and no way to blank a
// field by forgetting it.
// 
// One node of the category tree. `parent_id` is the structure this app
// navigates — null is a root — while `path` is kept only for importers
// that carry one and nothing here reads or writes it. A category is
// hand-picked or RULE-DRIVEN: a non-null `rules` selector makes every
// matching product a `product_categories` row with source `rule`, alongside
// the hand-picked ones, and `rules_computed_at` says when that last
// completed.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `code` answers 409.
func (srv *ProductsCategories) ProductsCategoriesUpdate(Id string, optionalSetters ...ProductsCategoriesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/categories/{id}")
	options := ProductsCategoriesUpdateOptions{}.New()
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
	if options.enabledSetters["ParentId"] {
		params["parent_id"] = options.ParentId
	}
	if options.enabledSetters["Path"] {
		params["path"] = options.Path
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
	if options.enabledSetters["RulesComputedAt"] {
		params["rules_computed_at"] = options.RulesComputedAt
	}
	if options.enabledSetters["Values"] {
		params["values"] = options.Values
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
type ProductsProductCategoriesListOptions struct {
	Limit int
	Offset int
	Order string
	Id string
	ProductId string
	CategoryId string
	Position int
	Source string
	CreatedAt string
	enabledSetters map[string]bool
}
func (options ProductsProductCategoriesListOptions) New() *ProductsProductCategoriesListOptions {
	options.enabledSetters = map[string]bool{
		"Limit": false,
		"Offset": false,
		"Order": false,
		"Id": false,
		"ProductId": false,
		"CategoryId": false,
		"Position": false,
		"Source": false,
		"CreatedAt": false,
	}
	return &options
}
type ProductsProductCategoriesListOption func(*ProductsProductCategoriesListOptions)
func (srv *ProductsCategories) WithProductsProductCategoriesListLimit(v int) ProductsProductCategoriesListOption {
	return func(o *ProductsProductCategoriesListOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesListOffset(v int) ProductsProductCategoriesListOption {
	return func(o *ProductsProductCategoriesListOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesListOrder(v string) ProductsProductCategoriesListOption {
	return func(o *ProductsProductCategoriesListOptions) {
		o.Order = v
		o.enabledSetters["Order"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesListId(v string) ProductsProductCategoriesListOption {
	return func(o *ProductsProductCategoriesListOptions) {
		o.Id = v
		o.enabledSetters["Id"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesListProductId(v string) ProductsProductCategoriesListOption {
	return func(o *ProductsProductCategoriesListOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesListCategoryId(v string) ProductsProductCategoriesListOption {
	return func(o *ProductsProductCategoriesListOptions) {
		o.CategoryId = v
		o.enabledSetters["CategoryId"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesListPosition(v int) ProductsProductCategoriesListOption {
	return func(o *ProductsProductCategoriesListOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesListSource(v string) ProductsProductCategoriesListOption {
	return func(o *ProductsProductCategoriesListOptions) {
		o.Source = v
		o.enabledSetters["Source"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesListCreatedAt(v string) ProductsProductCategoriesListOption {
	return func(o *ProductsProductCategoriesListOptions) {
		o.CreatedAt = v
		o.enabledSetters["CreatedAt"] = true
	}
}
	
// ProductsProductCategoriesList one membership: this product is filed in this
// category. `source` says how it got there — `manual` is hand-picked,
// `rule` was materialized by a category rule — and the two never touch each
// other: a recompute only ever inserts and deletes `rule` rows, so a
// hand-picked membership survives every pass. `POST
// /products/{id}/categories` is the friendlier way to create one, because it
// takes the product from the path and answers with the category code and the
// SKU.
// 
// Every column of `product_categories` is an exact-match query parameter,
// `order` sorts by one column, and `limit`/`offset` page through
// `page.total`. A query key that is NOT a column is dropped rather than
// refused, and the `filter` object echoes the ones that were understood —
// that echo is the only way to tell an unfiltered answer from an empty one.
// It reads rows exactly as they are stored: no join is resolved, no jsonb
// value is unpacked.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsCategories) ProductsProductCategoriesList(optionalSetters ...ProductsProductCategoriesListOption)(*interface{}, error) {
	path := "/v1/products/product_categories"
	options := ProductsProductCategoriesListOptions{}.New()
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
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
	}
	if options.enabledSetters["CategoryId"] {
		params["category_id"] = options.CategoryId
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["Source"] {
		params["source"] = options.Source
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
type ProductsProductCategoriesCreateOptions struct {
	Position int
	Source string
	enabledSetters map[string]bool
}
func (options ProductsProductCategoriesCreateOptions) New() *ProductsProductCategoriesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Position": false,
		"Source": false,
	}
	return &options
}
type ProductsProductCategoriesCreateOption func(*ProductsProductCategoriesCreateOptions)
func (srv *ProductsCategories) WithProductsProductCategoriesCreatePosition(v int) ProductsProductCategoriesCreateOption {
	return func(o *ProductsProductCategoriesCreateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesCreateSource(v string) ProductsProductCategoriesCreateOption {
	return func(o *ProductsProductCategoriesCreateOptions) {
		o.Source = v
		o.enabledSetters["Source"] = true
	}
}
					
// ProductsProductCategoriesCreate creates one product category membership and
// answers 201 with the stored row, including the id and the timestamps the
// database filled in — a client never sends an id, it reads one back and
// uses it in the path of every later call.
// 
// One membership: this product is filed in this category. `source` says how
// it got there — `manual` is hand-picked, `rule` was materialized by a
// category rule — and the two never touch each other: a recompute only ever
// inserts and deletes `rule` rows, so a hand-picked membership survives every
// pass. `POST /products/{id}/categories` is the friendlier way to create one,
// because it takes the product from the path and answers with the category
// code and the SKU.
// 
// `product_id` and `category_id` are the only columns the database refuses
// the row without; everything else has a default or is nullable. A second row
// with the same `product_id` and `category_id` answers 409.
func (srv *ProductsCategories) ProductsProductCategoriesCreate(CategoryId string, ProductId string, optionalSetters ...ProductsProductCategoriesCreateOption)(*models.Error, error) {
	path := "/v1/products/product_categories"
	options := ProductsProductCategoriesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["category_id"] = CategoryId
	params["product_id"] = ProductId
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
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
	
// ProductsProductCategoriesDelete deletes one product category membership by
// id. It is a hard delete — the row is gone, and the answer is a
// confirmation rather than a result to branch on.
// 
// Nothing in this schema references it, so nothing else changes.
// 
// An id no product category membership of this tenant carries answers 404;
// there is no 409, because every foreign key pointing at this entity resolves
// itself on delete rather than blocking one.
func (srv *ProductsCategories) ProductsProductCategoriesDelete(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/product_categories/{id}")
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
	
// ProductsProductCategoriesGet reads one product category membership by its
// id — the whole row, every column, as it is stored.
// 
// One membership: this product is filed in this category. `source` says how
// it got there — `manual` is hand-picked, `rule` was materialized by a
// category rule — and the two never touch each other: a recompute only ever
// inserts and deletes `rule` rows, so a hand-picked membership survives every
// pass. `POST /products/{id}/categories` is the friendlier way to create one,
// because it takes the product from the path and answers with the category
// code and the SKU.
// 
// An id no product category membership of this tenant carries answers 404,
// and so does one belonging to another tenant: row-level security makes that
// row invisible rather than forbidden. A malformed id answers 400 before the
// route is reached.
// 
// Answered from the gateway's tenant cache for up to 30 minutes and dropped
// the moment this entity is written, because the data model changes weekly at
// most and every product page asks the same question.
func (srv *ProductsCategories) ProductsProductCategoriesGet(Id string)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/product_categories/{id}")
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
type ProductsProductCategoriesUpdateOptions struct {
	CategoryId string
	Position int
	ProductId string
	Source string
	enabledSetters map[string]bool
}
func (options ProductsProductCategoriesUpdateOptions) New() *ProductsProductCategoriesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"CategoryId": false,
		"Position": false,
		"ProductId": false,
		"Source": false,
	}
	return &options
}
type ProductsProductCategoriesUpdateOption func(*ProductsProductCategoriesUpdateOptions)
func (srv *ProductsCategories) WithProductsProductCategoriesUpdateCategoryId(v string) ProductsProductCategoriesUpdateOption {
	return func(o *ProductsProductCategoriesUpdateOptions) {
		o.CategoryId = v
		o.enabledSetters["CategoryId"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesUpdatePosition(v int) ProductsProductCategoriesUpdateOption {
	return func(o *ProductsProductCategoriesUpdateOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesUpdateProductId(v string) ProductsProductCategoriesUpdateOption {
	return func(o *ProductsProductCategoriesUpdateOptions) {
		o.ProductId = v
		o.enabledSetters["ProductId"] = true
	}
}
func (srv *ProductsCategories) WithProductsProductCategoriesUpdateSource(v string) ProductsProductCategoriesUpdateOption {
	return func(o *ProductsProductCategoriesUpdateOptions) {
		o.Source = v
		o.enabledSetters["Source"] = true
	}
}
			
// ProductsProductCategoriesUpdate updates one product category membership by
// id. A partial patch: the body names only the columns to change and every
// column it leaves out keeps its current value, so there is no
// read-modify-write and no way to blank a field by forgetting it.
// 
// One membership: this product is filed in this category. `source` says how
// it got there — `manual` is hand-picked, `rule` was materialized by a
// category rule — and the two never touch each other: a recompute only ever
// inserts and deletes `rule` rows, so a hand-picked membership survives every
// pass. `POST /products/{id}/categories` is the friendlier way to create one,
// because it takes the product from the path and answers with the category
// code and the SKU.
// 
// A body that names nothing writable is refused with 400 rather than answered
// as a no-op, an id nobody carries answers 404, and a value that collides on
// `product_id` and `category_id` answers 409.
func (srv *ProductsCategories) ProductsProductCategoriesUpdate(Id string, optionalSetters ...ProductsProductCategoriesUpdateOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/product_categories/{id}")
	options := ProductsProductCategoriesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["CategoryId"] {
		params["category_id"] = options.CategoryId
	}
	if options.enabledSetters["Position"] {
		params["position"] = options.Position
	}
	if options.enabledSetters["ProductId"] {
		params["product_id"] = options.ProductId
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
type ProductsCategoriesAssignOptions struct {
	Position int
	enabledSetters map[string]bool
}
func (options ProductsCategoriesAssignOptions) New() *ProductsCategoriesAssignOptions {
	options.enabledSetters = map[string]bool{
		"Position": false,
	}
	return &options
}
type ProductsCategoriesAssignOption func(*ProductsCategoriesAssignOptions)
func (srv *ProductsCategories) WithProductsCategoriesAssignPosition(v int) ProductsCategoriesAssignOption {
	return func(o *ProductsCategoriesAssignOptions) {
		o.Position = v
		o.enabledSetters["Position"] = true
	}
}
					
// ProductsCategoriesAssign files one product into one category by hand, and
// the membership is always `source: 'manual'` — a rule recompute never
// deletes or shadows it. product_categories holds 28 758 rows and had no
// write surface that named the product it was filing. This takes the product
// from the route and the category from the body, which is what a bulk 'add
// the selected products to …' needs. The membership is always
// source='manual', so a rule recompute never deletes or shadows it.
func (srv *ProductsCategories) ProductsCategoriesAssign(Id string, CategoryId string, optionalSetters ...ProductsCategoriesAssignOption)(*models.Error, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/products/{id}/categories")
	options := ProductsCategoriesAssignOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	params["category_id"] = CategoryId
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
