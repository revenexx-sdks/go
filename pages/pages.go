package pages

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Pages service
type Pages struct {
	client client.Client
}

func New(clt client.Client) *Pages {
	return &Pages{
		client: clt,
	}
}


// PagesDeliveryMenus
func (srv *Pages) PagesDeliveryMenus()(*interface{}, error) {
	path := "/v1/pages/delivery/menus"
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

// PagesDeliveryPage
func (srv *Pages) PagesDeliveryPage()(*models.DeliveryPage, error) {
	path := "/v1/pages/delivery/page"
	params := map[string]interface{}{}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.DeliveryPage{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DeliveryPage
	parsed, ok := resp.Result.(models.DeliveryPage)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// PagesDeliveryPages
func (srv *Pages) PagesDeliveryPages()(*interface{}, error) {
	path := "/v1/pages/delivery/pages"
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
	
// PagesDeliveryPreview
func (srv *Pages) PagesDeliveryPreview(Token string)(*models.DeliveryPage, error) {
	r := strings.NewReplacer("{token}", Token)
	path := r.Replace("/v1/pages/delivery/preview/{token}")
	params := map[string]interface{}{}
	params["token"] = Token
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.DeliveryPage{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.DeliveryPage
	parsed, ok := resp.Result.(models.DeliveryPage)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// PagesEditorEditStates
func (srv *Pages) PagesEditorEditStates()(*interface{}, error) {
	path := "/v1/pages/editor/edit-states"
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

// PagesEditorNotificationsList
func (srv *Pages) PagesEditorNotificationsList()(*interface{}, error) {
	path := "/v1/pages/editor/notifications"
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

// PagesEditorNotificationsMarkAllRead
func (srv *Pages) PagesEditorNotificationsMarkAllRead()(*interface{}, error) {
	path := "/v1/pages/editor/notifications/mark-all-read"
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

// PagesEditorNotificationsUnreadCount
func (srv *Pages) PagesEditorNotificationsUnreadCount()(*interface{}, error) {
	path := "/v1/pages/editor/notifications/unread-count"
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
func (srv *Pages) WithPagesEditorTranslateItems(v []interface{}) PagesEditorTranslateOption {
	return func(o *PagesEditorTranslateOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
	
// PagesEditorTranslate
func (srv *Pages) PagesEditorTranslate(optionalSetters ...PagesEditorTranslateOption)(*interface{}, error) {
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

// PagesEditorUserSettingsGet
func (srv *Pages) PagesEditorUserSettingsGet()(*interface{}, error) {
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
func (srv *Pages) WithPagesEditorUserSettingsPutSettings(v interface{}) PagesEditorUserSettingsPutOption {
	return func(o *PagesEditorUserSettingsPutOptions) {
		o.Settings = v
		o.enabledSetters["Settings"] = true
	}
}
	
// PagesEditorUserSettingsPut
func (srv *Pages) PagesEditorUserSettingsPut(optionalSetters ...PagesEditorUserSettingsPutOption)(*interface{}, error) {
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

// PagesEditorUsers
func (srv *Pages) PagesEditorUsers()(*interface{}, error) {
	path := "/v1/pages/editor/users"
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
	
// PagesEditorCommentsList
func (srv *Pages) PagesEditorCommentsList(PageId string)(*interface{}, error) {
	r := strings.NewReplacer("{pageId}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/comments")
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
type PagesEditorCommentsCreateOptions struct {
	BlockUuids []string
	ParentUuid string
	enabledSetters map[string]bool
}
func (options PagesEditorCommentsCreateOptions) New() *PagesEditorCommentsCreateOptions {
	options.enabledSetters = map[string]bool{
		"BlockUuids": false,
		"ParentUuid": false,
	}
	return &options
}
type PagesEditorCommentsCreateOption func(*PagesEditorCommentsCreateOptions)
func (srv *Pages) WithPagesEditorCommentsCreateBlockUuids(v []string) PagesEditorCommentsCreateOption {
	return func(o *PagesEditorCommentsCreateOptions) {
		o.BlockUuids = v
		o.enabledSetters["BlockUuids"] = true
	}
}
func (srv *Pages) WithPagesEditorCommentsCreateParentUuid(v string) PagesEditorCommentsCreateOption {
	return func(o *PagesEditorCommentsCreateOptions) {
		o.ParentUuid = v
		o.enabledSetters["ParentUuid"] = true
	}
}
					
// PagesEditorCommentsCreate
func (srv *Pages) PagesEditorCommentsCreate(PageId string, Body string, optionalSetters ...PagesEditorCommentsCreateOption)(*interface{}, error) {
	r := strings.NewReplacer("{pageId}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/comments")
	options := PagesEditorCommentsCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["page_id"] = PageId
	params["body"] = Body
	if options.enabledSetters["BlockUuids"] {
		params["blockUuids"] = options.BlockUuids
	}
	if options.enabledSetters["ParentUuid"] {
		params["parentUuid"] = options.ParentUuid
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
			
// PagesEditorCommentsDelete
func (srv *Pages) PagesEditorCommentsDelete(PageId string, Uuid string)(*interface{}, error) {
	r := strings.NewReplacer("{pageId}", PageId, "{uuid}", Uuid)
	path := r.Replace("/v1/pages/editor/{page_id}/comments/{uuid}")
	params := map[string]interface{}{}
	params["page_id"] = PageId
	params["uuid"] = Uuid
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("DELETE", path, headers, params)
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
					
// PagesEditorCommentsUpdate
func (srv *Pages) PagesEditorCommentsUpdate(PageId string, Uuid string, Body string)(*interface{}, error) {
	r := strings.NewReplacer("{pageId}", PageId, "{uuid}", Uuid)
	path := r.Replace("/v1/pages/editor/{page_id}/comments/{uuid}")
	params := map[string]interface{}{}
	params["page_id"] = PageId
	params["uuid"] = Uuid
	params["body"] = Body
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
			
// PagesEditorCommentsResolve
func (srv *Pages) PagesEditorCommentsResolve(PageId string, Uuid string)(*interface{}, error) {
	r := strings.NewReplacer("{pageId}", PageId, "{uuid}", Uuid)
	path := r.Replace("/v1/pages/editor/{page_id}/comments/{uuid}/resolve")
	params := map[string]interface{}{}
	params["page_id"] = PageId
	params["uuid"] = Uuid
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
					
// PagesEditorCommentsToggleTask
func (srv *Pages) PagesEditorCommentsToggleTask(PageId string, Uuid string, TaskIndex int)(*models.Comment, error) {
	r := strings.NewReplacer("{pageId}", PageId, "{uuid}", Uuid)
	path := r.Replace("/v1/pages/editor/{page_id}/comments/{uuid}/toggle-task")
	params := map[string]interface{}{}
	params["page_id"] = PageId
	params["uuid"] = Uuid
	params["taskIndex"] = TaskIndex
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Comment{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Comment
	parsed, ok := resp.Result.(models.Comment)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// PagesEditorCommentsUnresolve
func (srv *Pages) PagesEditorCommentsUnresolve(PageId string, Uuid string)(*interface{}, error) {
	r := strings.NewReplacer("{pageId}", PageId, "{uuid}", Uuid)
	path := r.Replace("/v1/pages/editor/{page_id}/comments/{uuid}/unresolve")
	params := map[string]interface{}{}
	params["page_id"] = PageId
	params["uuid"] = Uuid
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
func (srv *Pages) WithPagesEditorHistoryLangcode(v string) PagesEditorHistoryOption {
	return func(o *PagesEditorHistoryOptions) {
		o.Langcode = v
		o.enabledSetters["Langcode"] = true
	}
}
					
// PagesEditorHistory
func (srv *Pages) PagesEditorHistory(PageId string, Index int, optionalSetters ...PagesEditorHistoryOption)(*models.MutationResponse, error) {
	r := strings.NewReplacer("{pageId}", PageId)
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
	
// PagesEditorLastChanged
func (srv *Pages) PagesEditorLastChanged(PageId string)(*interface{}, error) {
	r := strings.NewReplacer("{pageId}", PageId)
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
func (srv *Pages) WithPagesEditorMutationStatusLangcode(v string) PagesEditorMutationStatusOption {
	return func(o *PagesEditorMutationStatusOptions) {
		o.Langcode = v
		o.enabledSetters["Langcode"] = true
	}
}
							
// PagesEditorMutationStatus
func (srv *Pages) PagesEditorMutationStatus(PageId string, Enabled bool, Index int, optionalSetters ...PagesEditorMutationStatusOption)(*models.MutationResponse, error) {
	r := strings.NewReplacer("{pageId}", PageId)
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
func (srv *Pages) WithPagesEditorMutateLangcode(v string) PagesEditorMutateOption {
	return func(o *PagesEditorMutateOptions) {
		o.Langcode = v
		o.enabledSetters["Langcode"] = true
	}
}
func (srv *Pages) WithPagesEditorMutatePayload(v interface{}) PagesEditorMutateOption {
	return func(o *PagesEditorMutateOptions) {
		o.Payload = v
		o.enabledSetters["Payload"] = true
	}
}
					
// PagesEditorMutate
func (srv *Pages) PagesEditorMutate(PageId string, Plugin string, optionalSetters ...PagesEditorMutateOption)(*models.MutationResponse, error) {
	r := strings.NewReplacer("{pageId}", PageId)
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
func (srv *Pages) WithPagesEditorPreviewGrantTtlHours(v int) PagesEditorPreviewGrantOption {
	return func(o *PagesEditorPreviewGrantOptions) {
		o.TtlHours = v
		o.enabledSetters["TtlHours"] = true
	}
}
			
// PagesEditorPreviewGrant
func (srv *Pages) PagesEditorPreviewGrant(PageId string, optionalSetters ...PagesEditorPreviewGrantOption)(*interface{}, error) {
	r := strings.NewReplacer("{pageId}", PageId)
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
func (srv *Pages) WithPagesEditorPublishForce(v bool) PagesEditorPublishOption {
	return func(o *PagesEditorPublishOptions) {
		o.Force = v
		o.enabledSetters["Force"] = true
	}
}
func (srv *Pages) WithPagesEditorPublishLabel(v string) PagesEditorPublishOption {
	return func(o *PagesEditorPublishOptions) {
		o.Label = v
		o.enabledSetters["Label"] = true
	}
}
			
// PagesEditorPublish
func (srv *Pages) PagesEditorPublish(PageId string, optionalSetters ...PagesEditorPublishOption)(*models.MutationResponse, error) {
	r := strings.NewReplacer("{pageId}", PageId)
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
	
// PagesEditorRevert
func (srv *Pages) PagesEditorRevert(PageId string)(*models.MutationResponse, error) {
	r := strings.NewReplacer("{pageId}", PageId)
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
			
// PagesEditorSchedule
func (srv *Pages) PagesEditorSchedule(PageId string, ScheduledAt string)(*interface{}, error) {
	r := strings.NewReplacer("{pageId}", PageId)
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
	
// PagesEditorState
func (srv *Pages) PagesEditorState(PageId string)(*models.EditorState, error) {
	r := strings.NewReplacer("{pageId}", PageId)
	path := r.Replace("/v1/pages/editor/{page_id}/state")
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
	
// PagesEditorTakeOwnership
func (srv *Pages) PagesEditorTakeOwnership(PageId string)(*models.MutationResponse, error) {
	r := strings.NewReplacer("{pageId}", PageId)
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
func (srv *Pages) WithPagesEditorTemplatesCreateDescription(v string) PagesEditorTemplatesCreateOption {
	return func(o *PagesEditorTemplatesCreateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Pages) WithPagesEditorTemplatesCreateFieldName(v string) PagesEditorTemplatesCreateOption {
	return func(o *PagesEditorTemplatesCreateOptions) {
		o.FieldName = v
		o.enabledSetters["FieldName"] = true
	}
}
func (srv *Pages) WithPagesEditorTemplatesCreateIsDefault(v bool) PagesEditorTemplatesCreateOption {
	return func(o *PagesEditorTemplatesCreateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Pages) WithPagesEditorTemplatesCreatePageBundle(v string) PagesEditorTemplatesCreateOption {
	return func(o *PagesEditorTemplatesCreateOptions) {
		o.PageBundle = v
		o.enabledSetters["PageBundle"] = true
	}
}
							
// PagesEditorTemplatesCreate
func (srv *Pages) PagesEditorTemplatesCreate(PageId string, Label string, Uuids []string, optionalSetters ...PagesEditorTemplatesCreateOption)(*models.Template, error) {
	r := strings.NewReplacer("{pageId}", PageId)
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

		parsed := models.Template{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Template
	parsed, ok := resp.Result.(models.Template)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PagesEditorUnschedule
func (srv *Pages) PagesEditorUnschedule(PageId string)(*interface{}, error) {
	r := strings.NewReplacer("{pageId}", PageId)
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

// PagesLibraryList
func (srv *Pages) PagesLibraryList()(*interface{}, error) {
	path := "/v1/pages/library"
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
	
// PagesLibraryDelete
func (srv *Pages) PagesLibraryDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/library/{id}")
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
	
// PagesLibraryGet
func (srv *Pages) PagesLibraryGet(Id string)(*models.LibraryItem, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/library/{id}")
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

		parsed := models.LibraryItem{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LibraryItem
	parsed, ok := resp.Result.(models.LibraryItem)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PagesLibraryUpdateOptions struct {
	Bundle string
	Label string
	Tree interface{}
	enabledSetters map[string]bool
}
func (options PagesLibraryUpdateOptions) New() *PagesLibraryUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Bundle": false,
		"Label": false,
		"Tree": false,
	}
	return &options
}
type PagesLibraryUpdateOption func(*PagesLibraryUpdateOptions)
func (srv *Pages) WithPagesLibraryUpdateBundle(v string) PagesLibraryUpdateOption {
	return func(o *PagesLibraryUpdateOptions) {
		o.Bundle = v
		o.enabledSetters["Bundle"] = true
	}
}
func (srv *Pages) WithPagesLibraryUpdateLabel(v string) PagesLibraryUpdateOption {
	return func(o *PagesLibraryUpdateOptions) {
		o.Label = v
		o.enabledSetters["Label"] = true
	}
}
func (srv *Pages) WithPagesLibraryUpdateTree(v interface{}) PagesLibraryUpdateOption {
	return func(o *PagesLibraryUpdateOptions) {
		o.Tree = v
		o.enabledSetters["Tree"] = true
	}
}
			
// PagesLibraryUpdate
func (srv *Pages) PagesLibraryUpdate(Id string, optionalSetters ...PagesLibraryUpdateOption)(*models.LibraryItem, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/library/{id}")
	options := PagesLibraryUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Bundle"] {
		params["bundle"] = options.Bundle
	}
	if options.enabledSetters["Label"] {
		params["label"] = options.Label
	}
	if options.enabledSetters["Tree"] {
		params["tree"] = options.Tree
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

		parsed := models.LibraryItem{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LibraryItem
	parsed, ok := resp.Result.(models.LibraryItem)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// PagesMenusList
func (srv *Pages) PagesMenusList()(*interface{}, error) {
	path := "/v1/pages/menus"
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
type PagesMenusUpsertOptions struct {
	Items []interface{}
	enabledSetters map[string]bool
}
func (options PagesMenusUpsertOptions) New() *PagesMenusUpsertOptions {
	options.enabledSetters = map[string]bool{
		"Items": false,
	}
	return &options
}
type PagesMenusUpsertOption func(*PagesMenusUpsertOptions)
func (srv *Pages) WithPagesMenusUpsertItems(v []interface{}) PagesMenusUpsertOption {
	return func(o *PagesMenusUpsertOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
					
// PagesMenusUpsert
func (srv *Pages) PagesMenusUpsert(Label string, MenuKey string, optionalSetters ...PagesMenusUpsertOption)(*models.Menu, error) {
	path := "/v1/pages/menus"
	options := PagesMenusUpsertOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["label"] = Label
	params["menuKey"] = MenuKey
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

		parsed := models.Menu{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Menu
	parsed, ok := resp.Result.(models.Menu)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PagesMenusDelete
func (srv *Pages) PagesMenusDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/menus/{id}")
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
	
// PagesMenusGet
func (srv *Pages) PagesMenusGet(Id string)(*models.Menu, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/menus/{id}")
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

		parsed := models.Menu{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Menu
	parsed, ok := resp.Result.(models.Menu)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PagesMenusUpdateOptions struct {
	Items []interface{}
	Label string
	enabledSetters map[string]bool
}
func (options PagesMenusUpdateOptions) New() *PagesMenusUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Items": false,
		"Label": false,
	}
	return &options
}
type PagesMenusUpdateOption func(*PagesMenusUpdateOptions)
func (srv *Pages) WithPagesMenusUpdateItems(v []interface{}) PagesMenusUpdateOption {
	return func(o *PagesMenusUpdateOptions) {
		o.Items = v
		o.enabledSetters["Items"] = true
	}
}
func (srv *Pages) WithPagesMenusUpdateLabel(v string) PagesMenusUpdateOption {
	return func(o *PagesMenusUpdateOptions) {
		o.Label = v
		o.enabledSetters["Label"] = true
	}
}
			
// PagesMenusUpdate
func (srv *Pages) PagesMenusUpdate(Id string, optionalSetters ...PagesMenusUpdateOption)(*models.Menu, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/menus/{id}")
	options := PagesMenusUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Items"] {
		params["items"] = options.Items
	}
	if options.enabledSetters["Label"] {
		params["label"] = options.Label
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

		parsed := models.Menu{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Menu
	parsed, ok := resp.Result.(models.Menu)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}

// PagesPagesList
func (srv *Pages) PagesPagesList()(*interface{}, error) {
	path := "/v1/pages/pages"
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
type PagesPagesCreateOptions struct {
	Bundle string
	HostOptions interface{}
	Meta interface{}
	Slug string
	SourceLanguage string
	enabledSetters map[string]bool
}
func (options PagesPagesCreateOptions) New() *PagesPagesCreateOptions {
	options.enabledSetters = map[string]bool{
		"Bundle": false,
		"HostOptions": false,
		"Meta": false,
		"Slug": false,
		"SourceLanguage": false,
	}
	return &options
}
type PagesPagesCreateOption func(*PagesPagesCreateOptions)
func (srv *Pages) WithPagesPagesCreateBundle(v string) PagesPagesCreateOption {
	return func(o *PagesPagesCreateOptions) {
		o.Bundle = v
		o.enabledSetters["Bundle"] = true
	}
}
func (srv *Pages) WithPagesPagesCreateHostOptions(v interface{}) PagesPagesCreateOption {
	return func(o *PagesPagesCreateOptions) {
		o.HostOptions = v
		o.enabledSetters["HostOptions"] = true
	}
}
func (srv *Pages) WithPagesPagesCreateMeta(v interface{}) PagesPagesCreateOption {
	return func(o *PagesPagesCreateOptions) {
		o.Meta = v
		o.enabledSetters["Meta"] = true
	}
}
func (srv *Pages) WithPagesPagesCreateSlug(v string) PagesPagesCreateOption {
	return func(o *PagesPagesCreateOptions) {
		o.Slug = v
		o.enabledSetters["Slug"] = true
	}
}
func (srv *Pages) WithPagesPagesCreateSourceLanguage(v string) PagesPagesCreateOption {
	return func(o *PagesPagesCreateOptions) {
		o.SourceLanguage = v
		o.enabledSetters["SourceLanguage"] = true
	}
}
			
// PagesPagesCreate
func (srv *Pages) PagesPagesCreate(Title string, optionalSetters ...PagesPagesCreateOption)(*models.Page, error) {
	path := "/v1/pages/pages"
	options := PagesPagesCreateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["title"] = Title
	if options.enabledSetters["Bundle"] {
		params["bundle"] = options.Bundle
	}
	if options.enabledSetters["HostOptions"] {
		params["hostOptions"] = options.HostOptions
	}
	if options.enabledSetters["Meta"] {
		params["meta"] = options.Meta
	}
	if options.enabledSetters["Slug"] {
		params["slug"] = options.Slug
	}
	if options.enabledSetters["SourceLanguage"] {
		params["sourceLanguage"] = options.SourceLanguage
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

		parsed := models.Page{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Page
	parsed, ok := resp.Result.(models.Page)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PagesPagesDelete
func (srv *Pages) PagesPagesDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/pages/{id}")
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
	
// PagesPagesGet
func (srv *Pages) PagesPagesGet(Id string)(*models.Page, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/pages/{id}")
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

		parsed := models.Page{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Page
	parsed, ok := resp.Result.(models.Page)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PagesPagesUpdateOptions struct {
	Bundle string
	Meta interface{}
	Slug string
	Status string
	Title string
	enabledSetters map[string]bool
}
func (options PagesPagesUpdateOptions) New() *PagesPagesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Bundle": false,
		"Meta": false,
		"Slug": false,
		"Status": false,
		"Title": false,
	}
	return &options
}
type PagesPagesUpdateOption func(*PagesPagesUpdateOptions)
func (srv *Pages) WithPagesPagesUpdateBundle(v string) PagesPagesUpdateOption {
	return func(o *PagesPagesUpdateOptions) {
		o.Bundle = v
		o.enabledSetters["Bundle"] = true
	}
}
func (srv *Pages) WithPagesPagesUpdateMeta(v interface{}) PagesPagesUpdateOption {
	return func(o *PagesPagesUpdateOptions) {
		o.Meta = v
		o.enabledSetters["Meta"] = true
	}
}
func (srv *Pages) WithPagesPagesUpdateSlug(v string) PagesPagesUpdateOption {
	return func(o *PagesPagesUpdateOptions) {
		o.Slug = v
		o.enabledSetters["Slug"] = true
	}
}
func (srv *Pages) WithPagesPagesUpdateStatus(v string) PagesPagesUpdateOption {
	return func(o *PagesPagesUpdateOptions) {
		o.Status = v
		o.enabledSetters["Status"] = true
	}
}
func (srv *Pages) WithPagesPagesUpdateTitle(v string) PagesPagesUpdateOption {
	return func(o *PagesPagesUpdateOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
			
// PagesPagesUpdate
func (srv *Pages) PagesPagesUpdate(Id string, optionalSetters ...PagesPagesUpdateOption)(*models.Page, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/pages/{id}")
	options := PagesPagesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Bundle"] {
		params["bundle"] = options.Bundle
	}
	if options.enabledSetters["Meta"] {
		params["meta"] = options.Meta
	}
	if options.enabledSetters["Slug"] {
		params["slug"] = options.Slug
	}
	if options.enabledSetters["Status"] {
		params["status"] = options.Status
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
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

		parsed := models.Page{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Page
	parsed, ok := resp.Result.(models.Page)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// PagesPagesRevisions
func (srv *Pages) PagesPagesRevisions(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/pages/{id}/revisions")
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
type PagesSeedOptions struct {
	Menus []interface{}
	Pages []interface{}
	enabledSetters map[string]bool
}
func (options PagesSeedOptions) New() *PagesSeedOptions {
	options.enabledSetters = map[string]bool{
		"Menus": false,
		"Pages": false,
	}
	return &options
}
type PagesSeedOption func(*PagesSeedOptions)
func (srv *Pages) WithPagesSeedMenus(v []interface{}) PagesSeedOption {
	return func(o *PagesSeedOptions) {
		o.Menus = v
		o.enabledSetters["Menus"] = true
	}
}
func (srv *Pages) WithPagesSeedPages(v []interface{}) PagesSeedOption {
	return func(o *PagesSeedOptions) {
		o.Pages = v
		o.enabledSetters["Pages"] = true
	}
}
	
// PagesSeed
func (srv *Pages) PagesSeed(optionalSetters ...PagesSeedOption)(*interface{}, error) {
	path := "/v1/pages/seed"
	options := PagesSeedOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Menus"] {
		params["menus"] = options.Menus
	}
	if options.enabledSetters["Pages"] {
		params["pages"] = options.Pages
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

// PagesTemplatesList
func (srv *Pages) PagesTemplatesList()(*interface{}, error) {
	path := "/v1/pages/templates"
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
	
// PagesTemplatesDelete
func (srv *Pages) PagesTemplatesDelete(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/templates/{id}")
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
	
// PagesTemplatesGet
func (srv *Pages) PagesTemplatesGet(Id string)(*models.Template, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/templates/{id}")
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

		parsed := models.Template{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Template
	parsed, ok := resp.Result.(models.Template)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type PagesTemplatesUpdateOptions struct {
	Description string
	FieldName string
	IsDefault bool
	Label string
	PageBundle string
	Tree []interface{}
	enabledSetters map[string]bool
}
func (options PagesTemplatesUpdateOptions) New() *PagesTemplatesUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Description": false,
		"FieldName": false,
		"IsDefault": false,
		"Label": false,
		"PageBundle": false,
		"Tree": false,
	}
	return &options
}
type PagesTemplatesUpdateOption func(*PagesTemplatesUpdateOptions)
func (srv *Pages) WithPagesTemplatesUpdateDescription(v string) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Pages) WithPagesTemplatesUpdateFieldName(v string) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.FieldName = v
		o.enabledSetters["FieldName"] = true
	}
}
func (srv *Pages) WithPagesTemplatesUpdateIsDefault(v bool) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.IsDefault = v
		o.enabledSetters["IsDefault"] = true
	}
}
func (srv *Pages) WithPagesTemplatesUpdateLabel(v string) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.Label = v
		o.enabledSetters["Label"] = true
	}
}
func (srv *Pages) WithPagesTemplatesUpdatePageBundle(v string) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.PageBundle = v
		o.enabledSetters["PageBundle"] = true
	}
}
func (srv *Pages) WithPagesTemplatesUpdateTree(v []interface{}) PagesTemplatesUpdateOption {
	return func(o *PagesTemplatesUpdateOptions) {
		o.Tree = v
		o.enabledSetters["Tree"] = true
	}
}
			
// PagesTemplatesUpdate
func (srv *Pages) PagesTemplatesUpdate(Id string, optionalSetters ...PagesTemplatesUpdateOption)(*models.Template, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/pages/templates/{id}")
	options := PagesTemplatesUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["FieldName"] {
		params["field_name"] = options.FieldName
	}
	if options.enabledSetters["IsDefault"] {
		params["is_default"] = options.IsDefault
	}
	if options.enabledSetters["Label"] {
		params["label"] = options.Label
	}
	if options.enabledSetters["PageBundle"] {
		params["page_bundle"] = options.PageBundle
	}
	if options.enabledSetters["Tree"] {
		params["tree"] = options.Tree
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

		parsed := models.Template{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Template
	parsed, ok := resp.Result.(models.Template)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
