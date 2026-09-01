package pages_editor

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// PagesEditor service
type PagesEditor struct {
	client client.Client
}

func New(clt client.Client) *PagesEditor {
	return &PagesEditor{
		client: clt,
	}
}

type PagesEditorEditStatesOptions struct {
	Status string
	Limit int
	Offset int
	enabledSetters map[string]bool
}
func (options PagesEditorEditStatesOptions) New() *PagesEditorEditStatesOptions {
	options.enabledSetters = map[string]bool{
		"Status": false,
		"Limit": false,
		"Offset": false,
	}
	return &options
}
type PagesEditorEditStatesOption func(*PagesEditorEditStatesOptions)
func (srv *PagesEditor) WithPagesEditorEditStatesStatus(v string) PagesEditorEditStatesOption {
	return func(o *PagesEditorEditStatesOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *PagesEditor) WithPagesEditorEditStatesLimit(v int) PagesEditorEditStatesOption {
	return func(o *PagesEditorEditStatesOptions) {
		o.Limit = v
		o.enabledSetters["Limit"] = true
	}
}
func (srv *PagesEditor) WithPagesEditorEditStatesOffset(v int) PagesEditorEditStatesOption {
	return func(o *PagesEditorEditStatesOptions) {
		o.Offset = v
		o.enabledSetters["Offset"] = true
	}
}
	
// PagesEditorEditStates the drafts overview — the "what is unpublished
// right now" list, across every page: who holds it, since when, and whether
// it is parked for a date. Always newest-first — this route does not read
// `order`. An edit state whose page has been deleted is dropped from `items`
// but still counted in `total`.
func (srv *PagesEditor) PagesEditorEditStates(optionalSetters ...PagesEditorEditStatesOption)(*interface{}, error) {
	path := "/v1/pages/editor/edit-states"
	options := PagesEditorEditStatesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
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
type PagesEditorTranslateOptions struct {
	Items []interface{}
	enabledSetters map[string]bool
}
func (options PagesEditorTranslateOptions) New() *PagesEditorTranslateOptions {
	options.enabledSetters = map[string]bool{
		"Items": false,
	}
	return &options
}
type PagesEditorTranslateOption func(*PagesEditorTranslateOptions)
func (srv *PagesEditor) WithPagesEditorTranslateItems(v []interface{}) PagesEditorTranslateOption {
	return func(o *PagesEditorTranslateOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
	
// PagesEditorTranslate the translation is the tenant's provider's, not this
// app's, and a tenant that has configured none gets no translation at all.
// The endpoint comes from the tenant setting `translate_endpoint`
// (PAGES_TRANSLATE_ENDPOINT remains a fallback). The bearer token does NOT:
// the gateway masks every setting flagged `sensitive`, so a key stored as one
// could never be read back — it stays the PAGES_TRANSLATE_KEY function
// secret. This app does not translate anything itself; it forwards `items`
// and hands the answer back.
func (srv *PagesEditor) PagesEditorTranslate(optionalSetters ...PagesEditorTranslateOption)(*models.Error, error) {
	path := "/v1/pages/editor/translate"
	options := PagesEditorTranslateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Items"] {
		params["items"] = options.Items
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

// PagesEditorUserSettingsGet per-user editor preferences — one row per
// user, scoped to this app. Not tenant configuration: nothing here changes
// what the API does, only how one person's editor looks.
func (srv *PagesEditor) PagesEditorUserSettingsGet()(*interface{}, error) {
	path := "/v1/pages/editor/user-settings"
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
type PagesEditorUserSettingsPutOptions struct {
	Settings interface{}
	enabledSetters map[string]bool
}
func (options PagesEditorUserSettingsPutOptions) New() *PagesEditorUserSettingsPutOptions {
	options.enabledSetters = map[string]bool{
		"Settings": false,
	}
	return &options
}
type PagesEditorUserSettingsPutOption func(*PagesEditorUserSettingsPutOptions)
func (srv *PagesEditor) WithPagesEditorUserSettingsPutSettings(v interface{}) PagesEditorUserSettingsPutOption {
	return func(o *PagesEditorUserSettingsPutOptions) {
		o.Settings = v
		o.enabledSetters["Settings"] = true
	}
}
	
// PagesEditorUserSettingsPut replaces the caller's preferences wholesale —
// this is not a merge, so send the whole bag.
func (srv *PagesEditor) PagesEditorUserSettingsPut(optionalSetters ...PagesEditorUserSettingsPutOption)(*interface{}, error) {
	path := "/v1/pages/editor/user-settings"
	options := PagesEditorUserSettingsPutOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Settings"] {
		params["settings"] = options.Settings
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
type PagesEditorHistoryOptions struct {
	Langcode string
	enabledSetters map[string]bool
}
func (options PagesEditorHistoryOptions) New() *PagesEditorHistoryOptions {
	options.enabledSetters = map[string]bool{
		"Langcode": false,
	}
	return &options
}
type PagesEditorHistoryOption func(*PagesEditorHistoryOptions)
func (srv *PagesEditor) WithPagesEditorHistoryLangcode(v string) PagesEditorHistoryOption {
	return func(o *PagesEditorHistoryOptions) {
		o.Langcode = v
		o.enabledSetters["Langcode"] = true
	}
}
					
// PagesEditorHistory undo and redo. The pointer is the edit state's
// `current_index`, the position in the mutation log the page is materialized
// at, and this route is the only thing that moves it — `GET
// …/state?index=` looks at another position without going there. The log
// itself is never rewritten — only the pointer moves — so redo stays
// available until the next change is appended.
func (srv *PagesEditor) PagesEditorHistory(PageId string, Index int, optionalSetters ...PagesEditorHistoryOption)(*models.MutationResponse, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/history")
	options := PagesEditorHistoryOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["page_id"] = PageId
	params["index"] = Index
	if options.enabledSetters["Langcode"] {
		params["langcode"] = options.Langcode
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

		parsed := models.MutationResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MutationResponse
	parsed, ok := resp.Result.(models.MutationResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PagesEditorLastChanged the cheap poll behind "someone else is editing this
// page": one integer, the moment the open edit state last moved, in epoch
// seconds rather than as a timestamp so a comparison is a subtraction.
// Compare it with the `updatedAt` you last saw and re-fetch the state only
// when it moved.
func (srv *PagesEditor) PagesEditorLastChanged(PageId string)(*interface{}, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/last-changed")
	params := map[string]interface{}{}
	params["page_id"] = PageId
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
type PagesEditorMutationStatusOptions struct {
	Langcode string
	enabledSetters map[string]bool
}
func (options PagesEditorMutationStatusOptions) New() *PagesEditorMutationStatusOptions {
	options.enabledSetters = map[string]bool{
		"Langcode": false,
	}
	return &options
}
type PagesEditorMutationStatusOption func(*PagesEditorMutationStatusOptions)
func (srv *PagesEditor) WithPagesEditorMutationStatusLangcode(v string) PagesEditorMutationStatusOption {
	return func(o *PagesEditorMutationStatusOptions) {
		o.Langcode = v
		o.enabledSetters["Langcode"] = true
	}
}
							
// PagesEditorMutationStatus take one change out of the replay without
// deleting it — "what would the page look like without this edit". The
// entry stays in the history and can be switched back on.
func (srv *PagesEditor) PagesEditorMutationStatus(PageId string, Enabled bool, Index int, optionalSetters ...PagesEditorMutationStatusOption)(*models.MutationResponse, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/mutation-status")
	options := PagesEditorMutationStatusOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["page_id"] = PageId
	params["enabled"] = Enabled
	params["index"] = Index
	if options.enabledSetters["Langcode"] {
		params["langcode"] = options.Langcode
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

		parsed := models.MutationResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MutationResponse
	parsed, ok := resp.Result.(models.MutationResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PagesEditorMutateOptions struct {
	Langcode string
	Payload interface{}
	enabledSetters map[string]bool
}
func (options PagesEditorMutateOptions) New() *PagesEditorMutateOptions {
	options.enabledSetters = map[string]bool{
		"Langcode": false,
		"Payload": false,
	}
	return &options
}
type PagesEditorMutateOption func(*PagesEditorMutateOptions)
func (srv *PagesEditor) WithPagesEditorMutateLangcode(v string) PagesEditorMutateOption {
	return func(o *PagesEditorMutateOptions) {
		o.Langcode = v
		o.enabledSetters["Langcode"] = true
	}
}
func (srv *PagesEditor) WithPagesEditorMutatePayload(v interface{}) PagesEditorMutateOption {
	return func(o *PagesEditorMutateOptions) {
		o.Payload = v
		o.enabledSetters["Payload"] = true
	}
}
					
// PagesEditorMutate the one way page CONTENT changes. Each call appends one
// entry to the append-only log and answers the whole re-materialized state,
// so a client never re-fetches. A page nobody has opened yet needs no
// separate call to open it: the first mutation creates the edit state and
// takes ownership of it, and every later one asks for that ownership, so a
// second person editing the same page is refused until they take it over.
// Appending while the pointer sits mid-history discards the redo branch,
// exactly as an editor expects.
func (srv *PagesEditor) PagesEditorMutate(PageId string, Plugin string, optionalSetters ...PagesEditorMutateOption)(*models.MutationResponse, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/mutations")
	options := PagesEditorMutateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["page_id"] = PageId
	params["plugin"] = Plugin
	if options.enabledSetters["Langcode"] {
		params["langcode"] = options.Langcode
	}
	if options.enabledSetters["Payload"] {
		params["payload"] = options.Payload
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

		parsed := models.MutationResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MutationResponse
	parsed, ok := resp.Result.(models.MutationResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PagesEditorPreviewGrantOptions struct {
	TtlHours int
	enabledSetters map[string]bool
}
func (options PagesEditorPreviewGrantOptions) New() *PagesEditorPreviewGrantOptions {
	options.enabledSetters = map[string]bool{
		"TtlHours": false,
	}
	return &options
}
type PagesEditorPreviewGrantOption func(*PagesEditorPreviewGrantOptions)
func (srv *PagesEditor) WithPagesEditorPreviewGrantTtlHours(v int) PagesEditorPreviewGrantOption {
	return func(o *PagesEditorPreviewGrantOptions) {
		o.TtlHours = v
		o.enabledSetters["TtlHours"] = true
	}
}
			
// PagesEditorPreviewGrant mints a link that shows this page's current edit
// state — the UNPUBLISHED one — to somebody without an editor account.
// The token is the whole credential — anyone holding it sees the page —
// so it expires, and a new one is cheap.
func (srv *PagesEditor) PagesEditorPreviewGrant(PageId string, optionalSetters ...PagesEditorPreviewGrantOption)(*interface{}, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/preview-grant")
	options := PagesEditorPreviewGrantOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["page_id"] = PageId
	if options.enabledSetters["TtlHours"] {
		params["ttlHours"] = options.TtlHours
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
type PagesEditorPublishOptions struct {
	Force bool
	Label string
	enabledSetters map[string]bool
}
func (options PagesEditorPublishOptions) New() *PagesEditorPublishOptions {
	options.enabledSetters = map[string]bool{
		"Force": false,
		"Label": false,
	}
	return &options
}
type PagesEditorPublishOption func(*PagesEditorPublishOptions)
func (srv *PagesEditor) WithPagesEditorPublishForce(v bool) PagesEditorPublishOption {
	return func(o *PagesEditorPublishOptions) {
		o.Force = v
		o.enabledSetters["Force"] = true
	}
}
func (srv *PagesEditor) WithPagesEditorPublishLabel(v string) PagesEditorPublishOption {
	return func(o *PagesEditorPublishOptions) {
		o.Label = v
		o.enabledSetters["Label"] = true
	}
}
			
// PagesEditorPublish four things in one call: the mutation log is replayed
// into a finished block tree, that tree is snapshotted into a new revision,
// the page's canonical blocks are replaced by it, and the edit state is
// archived — so the page comes out of this with nothing unpublished and the
// working copy behind it closed rather than deleted. The revision is written
// FIRST and the canonical blocks replaced after, so a failure mid-way leaves
// the page recoverable. Block uuids survive, which is why comments anchored
// to a block outlive the publish.
func (srv *PagesEditor) PagesEditorPublish(PageId string, optionalSetters ...PagesEditorPublishOption)(*models.Error, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/publish")
	options := PagesEditorPublishOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["page_id"] = PageId
	if options.enabledSetters["Force"] {
		params["force"] = options.Force
	}
	if options.enabledSetters["Label"] {
		params["label"] = options.Label
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
	
// PagesEditorRevert throws the whole working copy away: the edit state row is
// deleted and its mutation log with it, so the history goes too — this is
// not an undo and cannot itself be undone. Unlike publishing, which archives
// the edit state, nothing of it survives to be reopened. The published page
// is untouched.
func (srv *PagesEditor) PagesEditorRevert(PageId string)(*models.MutationResponse, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/revert")
	params := map[string]interface{}{}
	params["page_id"] = PageId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MutationResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MutationResponse
	parsed, ok := resp.Result.(models.MutationResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// PagesEditorSchedule gated on the tenant setting
// `enable_scheduled_publishing`, which is off by default: nothing in the
// platform publishes a scheduled edit state yet, so a date accepted here
// would be a promise the app cannot keep. Every editor state carries
// `features.scheduledPublishing` so the control can be hidden rather than the
// refusal discovered.
func (srv *PagesEditor) PagesEditorSchedule(PageId string, ScheduledAt string)(*models.Error, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/schedule")
	params := map[string]interface{}{}
	params["page_id"] = PageId
	params["scheduledAt"] = ScheduledAt
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
type PagesEditorStateOptions struct {
	Langcode string
	Index int
	enabledSetters map[string]bool
}
func (options PagesEditorStateOptions) New() *PagesEditorStateOptions {
	options.enabledSetters = map[string]bool{
		"Langcode": false,
		"Index": false,
	}
	return &options
}
type PagesEditorStateOption func(*PagesEditorStateOptions)
func (srv *PagesEditor) WithPagesEditorStateLangcode(v string) PagesEditorStateOption {
	return func(o *PagesEditorStateOptions) {
		o.Langcode = v
		o.enabledSetters["Langcode"] = true
	}
}
func (srv *PagesEditor) WithPagesEditorStateIndex(v int) PagesEditorStateOption {
	return func(o *PagesEditorStateOptions) {
		o.Index = v
		o.enabledSetters["Index"] = true
	}
}
			
// PagesEditorState the one call the visual editor boots on, and the only
// place the UNPUBLISHED page can be seen whole: the canonical blocks with
// every enabled mutation of the log replayed over them, the resulting field
// lists, the mutation history itself, who owns the edit state and where the
// undo pointer sits, and the tenant's editor feature flags. `langcode`
// decides which language the props resolve in, falling back to the page's
// source language. `index` replays the log up to a given position instead of
// the current one, which is how the editor previews an undo without
// performing it — it changes nothing, so it is safe to call at any
// position. Reading this creates nothing either: a page nobody has opened
// answers with a null `editState`, an empty history, and the published blocks
// as they stand.
func (srv *PagesEditor) PagesEditorState(PageId string, optionalSetters ...PagesEditorStateOption)(*models.EditorState, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/state")
	options := PagesEditorStateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["page_id"] = PageId
	if options.enabledSetters["Langcode"] {
		params["langcode"] = options.Langcode
	}
	if options.enabledSetters["Index"] {
		params["index"] = options.Index
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.EditorState{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.EditorState
	parsed, ok := resp.Result.(models.EditorState)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PagesEditorTakeOwnership one page has one writer. This is how the second
// person gets the pen — the previous owner is notified rather than silently
// locked out.
func (srv *PagesEditor) PagesEditorTakeOwnership(PageId string)(*models.MutationResponse, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/take-ownership")
	params := map[string]interface{}{}
	params["page_id"] = PageId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MutationResponse{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MutationResponse
	parsed, ok := resp.Result.(models.MutationResponse)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PagesEditorTemplatesCreateOptions struct {
	Description string
	FieldName string
	IsDefault bool
	PageBundle string
	enabledSetters map[string]bool
}
func (options PagesEditorTemplatesCreateOptions) New() *PagesEditorTemplatesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"FieldName": false,
		"IsDefault": false,
		"PageBundle": false,
	}
	return &options
}
type PagesEditorTemplatesCreateOption func(*PagesEditorTemplatesCreateOptions)
func (srv *PagesEditor) WithPagesEditorTemplatesCreateDescription(v string) PagesEditorTemplatesCreateOption {
	return func(o *PagesEditorTemplatesCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *PagesEditor) WithPagesEditorTemplatesCreateFieldName(v string) PagesEditorTemplatesCreateOption {
	return func(o *PagesEditorTemplatesCreateOptions) {
		o.FieldName = v
		o.enabledSetters["FieldName"] = true
	}
}
func (srv *PagesEditor) WithPagesEditorTemplatesCreateIsDefault(v bool) PagesEditorTemplatesCreateOption {
	return func(o *PagesEditorTemplatesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *PagesEditor) WithPagesEditorTemplatesCreatePageBundle(v string) PagesEditorTemplatesCreateOption {
	return func(o *PagesEditorTemplatesCreateOptions) {
		o.PageBundle = v
		o.enabledSetters["PageBundle"] = true
	}
}
							
// PagesEditorTemplatesCreate freezes a selection into a reusable starting
// point. The blocks are read out of the page's CURRENT edit state rather than
// out of what is published, so a template can be cut from work in progress
// and the uuids you send are the ones the editor is showing. Unlike making a
// block reusable, this COPIES: pages later made from the template are
// independent of it and of each other.
func (srv *PagesEditor) PagesEditorTemplatesCreate(PageId string, Label string, Uuids []string, optionalSetters ...PagesEditorTemplatesCreateOption)(*models.Error, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/templates")
	options := PagesEditorTemplatesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["page_id"] = PageId
	params["label"] = Label
	params["uuids"] = Uuids
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["FieldName"] {
		params["fieldName"] = options.FieldName
	}
	if options.enabledSetters["IsDefault"] {
		params["isDefault"] = options.IsDefault
	}
	if options.enabledSetters["PageBundle"] {
		params["pageBundle"] = options.PageBundle
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
	
// PagesEditorUnschedule takes a parked edit state back to `active` and clears
// its date, so the scheduled publication simply does not happen. The work is
// not touched — the mutation log, the undo position and the owner all stay
// as they were — and the page can then be published by hand or scheduled
// again for a different date. Like every other write to an edit state it asks
// for ownership, and a page with no open edit state answers 404 rather than
// pretending to have cancelled something.
func (srv *PagesEditor) PagesEditorUnschedule(PageId string)(*interface{}, error) {
	r := strings.NewReplacer("{page_id}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/unschedule")
	params := map[string]interface{}{}
	params["page_id"] = PageId
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
