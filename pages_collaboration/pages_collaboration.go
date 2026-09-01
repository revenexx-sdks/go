package pages_collaboration

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// PagesCollaboration service
type PagesCollaboration struct {
	client client.Client
}

func New(clt client.Client) *PagesCollaboration {
	return &PagesCollaboration{
		client: clt,
	}
}

type PagesEditorNotificationsListOptions struct {
	After string
	MarkAsRead string
	enabledSetters map[string]bool
}
func (options PagesEditorNotificationsListOptions) New() *PagesEditorNotificationsListOptions {
	options.enabledSetters = map[string]bool{
		"After": false,
		"MarkAsRead": false,
	}
	return &options
}
type PagesEditorNotificationsListOption func(*PagesEditorNotificationsListOptions)
func (srv *PagesCollaboration) WithPagesEditorNotificationsListAfter(v string) PagesEditorNotificationsListOption {
	return func(o *PagesEditorNotificationsListOptions) {
		o.After = v
		o.enabledSetters["After"] = true
	}
}
func (srv *PagesCollaboration) WithPagesEditorNotificationsListMarkAsRead(v string) PagesEditorNotificationsListOption {
	return func(o *PagesEditorNotificationsListOptions) {
		o.MarkAsRead = v
		o.enabledSetters["MarkAsRead"] = true
	}
}
	
// PagesEditorNotificationsList the caller's own notifications, newest first,
// 20 at a time. Paged by an opaque cursor rather than by offset, so new
// arrivals never shift a page under the reader. It is also the one read in
// this app that writes: `?markAsRead=true` flags the notifications on the
// page it just returned as read, which is how a feed that has been looked at
// empties its badge without a second call — leave it off and reading
// changes nothing.
func (srv *PagesCollaboration) PagesEditorNotificationsList(optionalSetters ...PagesEditorNotificationsListOption)(*interface{}, error) {
	path := "/v1/pages/editor/notifications"
	options := PagesEditorNotificationsListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["After"] {
		params["after"] = options.After
	}
	if options.enabledSetters["MarkAsRead"] {
		params["markAsRead"] = options.MarkAsRead
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

// PagesEditorNotificationsMarkAllRead empties the badge in one call. Every
// unread notification of the CURRENT user is flagged read — the user is the
// one the request's context token names and there is no body with which to
// name another. Nothing is deleted: `GET /pages/editor/notifications` still
// returns the same feed, just with `read` set. The answer is the new unread
// count, so a client can set the badge straight from it without a second
// read.
func (srv *PagesCollaboration) PagesEditorNotificationsMarkAllRead()(*interface{}, error) {
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

// PagesEditorNotificationsUnreadCount the cheap poll behind the badge.
func (srv *PagesCollaboration) PagesEditorNotificationsUnreadCount()(*interface{}, error) {
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

// PagesEditorUsers what the @mention picker is filled from. When the identity
// service cannot be reached this degrades to the authors who have already
// commented on this tenant's pages rather than answering an error — a
// mention list that is short is more useful than one that is missing.
func (srv *PagesCollaboration) PagesEditorUsers()(*interface{}, error) {
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
	
// PagesEditorCommentsList every comment on the page in one flat list, oldest
// first, roots and replies together and resolved threads included — there
// is no filter and no paging, because the editor nests and filters them
// itself from `parentUuid` and pins each root to its blocks with
// `blockUuids`. Comments hang off the PAGE, not off a revision or an edit
// state, so publishing and reverting leave them standing; that is what makes
// them usable as a review trail across several rounds of edits.
func (srv *PagesCollaboration) PagesEditorCommentsList(PageId string)(*models.PageCommentList, error) {
	r := strings.NewReplacer("{page_id}", PageId)
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

		parsed := models.PageCommentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PageCommentList
	parsed, ok := resp.Result.(models.PageCommentList)
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
func (srv *PagesCollaboration) WithPagesEditorCommentsCreateBlockUuids(v []string) PagesEditorCommentsCreateOption {
	return func(o *PagesEditorCommentsCreateOptions) {
		o.BlockUuids = v
		o.enabledSetters["BlockUuids"] = true
	}
}
func (srv *PagesCollaboration) WithPagesEditorCommentsCreateParentUuid(v string) PagesEditorCommentsCreateOption {
	return func(o *PagesEditorCommentsCreateOptions) {
		o.ParentUuid = v
		o.enabledSetters["ParentUuid"] = true
	}
}
					
// PagesEditorCommentsCreate the same route writes both kinds, and which one
// you get is decided by the body: `blockUuids` starts a new thread pinned to
// those blocks, `parentUuid` hangs a reply under an existing root. Everyone
// named with an @mention in the body is notified, and on a reply so is
// everybody already in the thread — the actor never notifies themselves.
func (srv *PagesCollaboration) PagesEditorCommentsCreate(PageId string, Body string, optionalSetters ...PagesEditorCommentsCreateOption)(*models.PageCommentList, error) {
	r := strings.NewReplacer("{page_id}", PageId)
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

		parsed := models.PageCommentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PageCommentList
	parsed, ok := resp.Result.(models.PageCommentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// PagesEditorCommentsDelete a hard delete, and deleting a root takes its
// replies with it.
func (srv *PagesCollaboration) PagesEditorCommentsDelete(PageId string, Uuid string)(*models.PageCommentList, error) {
	r := strings.NewReplacer("{page_id}", PageId, "{uuid}", Uuid)
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

		parsed := models.PageCommentList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.PageCommentList
	parsed, ok := resp.Result.(models.PageCommentList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// PagesEditorCommentsUpdate rewrites what a comment says, and only its author
// may — a comment carries an `author_id` and anybody else is refused with
// 403. Only the body moves: what the comment is pinned to, whether the thread
// is resolved and who wrote it are all fixed when it is created. Rewriting a
// body does NOT re-run the @mention notifications, so mentioning somebody new
// by editing will not reach them. Answers the page's whole comment list
// rather than the one row, so a client can re-render from the response.
func (srv *PagesCollaboration) PagesEditorCommentsUpdate(PageId string, Uuid string, Body string)(*models.Error, error) {
	r := strings.NewReplacer("{page_id}", PageId, "{uuid}", Uuid)
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
			
// PagesEditorCommentsResolve marks a thread handled, so the editor stops
// surfacing it on the block it is pinned to. Only a ROOT can be resolved —
// resolved-ness is a property of the thread and not of a message in it, so
// pointing this at a reply is refused with 400 rather than quietly resolving
// its parent. Nothing is deleted, nobody is notified, and the thread stays in
// the list; `.../unresolve` is the way back. Answers the page's whole comment
// list.
func (srv *PagesCollaboration) PagesEditorCommentsResolve(PageId string, Uuid string)(*models.Error, error) {
	r := strings.NewReplacer("{page_id}", PageId, "{uuid}", Uuid)
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
					
// PagesEditorCommentsToggleTask a comment body may carry a task list. This
// flips one checkbox by rewriting the body's markup, and answers the single
// comment rather than the whole list. A `taskIndex` that names no checkbox is
// refused and nothing is written — the comment's `updated_at` is the
// editor's "edited" marker, so a call that changes nothing must not move it.
func (srv *PagesCollaboration) PagesEditorCommentsToggleTask(PageId string, Uuid string, TaskIndex int)(*models.Error, error) {
	r := strings.NewReplacer("{page_id}", PageId, "{uuid}", Uuid)
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
			
// PagesEditorCommentsUnresolve clears the resolved flag and puts the thread
// back in front of whoever is editing — the mirror of `.../resolve` in
// every respect, including that only a root can be reopened and that a reply
// answers 400. A thread that was already open is accepted and stays open.
// Answers the page's whole comment list.
func (srv *PagesCollaboration) PagesEditorCommentsUnresolve(PageId string, Uuid string)(*models.Error, error) {
	r := strings.NewReplacer("{page_id}", PageId, "{uuid}", Uuid)
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
