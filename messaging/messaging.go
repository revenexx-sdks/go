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

type MessagingListMessagesOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options MessagingListMessagesOptions) New() *MessagingListMessagesOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type MessagingListMessagesOption func(*MessagingListMessagesOptions)
func (srv *Messaging) WithMessagingListMessagesQueries(v []string) MessagingListMessagesOption {
	return func(o *MessagingListMessagesOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithMessagingListMessagesSearch(v string) MessagingListMessagesOption {
	return func(o *MessagingListMessagesOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Messaging) WithMessagingListMessagesTotal(v bool) MessagingListMessagesOption {
	return func(o *MessagingListMessagesOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// MessagingListMessages get a list of all messages from the current Revenexx
// project.
func (srv *Messaging) MessagingListMessages(optionalSetters ...MessagingListMessagesOption)(*models.MessageList, error) {
	path := "/v1/messaging/messages"
	options := MessagingListMessagesOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.MessageList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.MessageList
	parsed, ok := resp.Result.(models.MessageList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingCreateEmailOptions struct {
	Attachments []string
	Bcc []string
	Cc []string
	Draft bool
	Html bool
	ScheduledAt string
	Targets []string
	Topics []string
	Users []string
	enabledSetters map[string]bool
}
func (options MessagingCreateEmailOptions) New() *MessagingCreateEmailOptions {
	options.enabledSetters = map[string]bool{
		"Attachments": false,
		"Bcc": false,
		"Cc": false,
		"Draft": false,
		"Html": false,
		"ScheduledAt": false,
		"Targets": false,
		"Topics": false,
		"Users": false,
	}
	return &options
}
type MessagingCreateEmailOption func(*MessagingCreateEmailOptions)
func (srv *Messaging) WithMessagingCreateEmailAttachments(v []string) MessagingCreateEmailOption {
	return func(o *MessagingCreateEmailOptions) {
		o.Attachments = v
		o.enabledSetters["Attachments"] = true
	}
}
func (srv *Messaging) WithMessagingCreateEmailBcc(v []string) MessagingCreateEmailOption {
	return func(o *MessagingCreateEmailOptions) {
		o.Bcc = v
		o.enabledSetters["Bcc"] = true
	}
}
func (srv *Messaging) WithMessagingCreateEmailCc(v []string) MessagingCreateEmailOption {
	return func(o *MessagingCreateEmailOptions) {
		o.Cc = v
		o.enabledSetters["Cc"] = true
	}
}
func (srv *Messaging) WithMessagingCreateEmailDraft(v bool) MessagingCreateEmailOption {
	return func(o *MessagingCreateEmailOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithMessagingCreateEmailHtml(v bool) MessagingCreateEmailOption {
	return func(o *MessagingCreateEmailOptions) {
		o.Html = v
		o.enabledSetters["Html"] = true
	}
}
func (srv *Messaging) WithMessagingCreateEmailScheduledAt(v string) MessagingCreateEmailOption {
	return func(o *MessagingCreateEmailOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
func (srv *Messaging) WithMessagingCreateEmailTargets(v []string) MessagingCreateEmailOption {
	return func(o *MessagingCreateEmailOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithMessagingCreateEmailTopics(v []string) MessagingCreateEmailOption {
	return func(o *MessagingCreateEmailOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithMessagingCreateEmailUsers(v []string) MessagingCreateEmailOption {
	return func(o *MessagingCreateEmailOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
							
// MessagingCreateEmail create a new email message.
func (srv *Messaging) MessagingCreateEmail(Content string, MessageId string, Subject string, optionalSetters ...MessagingCreateEmailOption)(*models.Message, error) {
	path := "/v1/messaging/messages/email"
	options := MessagingCreateEmailOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["content"] = Content
	params["messageId"] = MessageId
	params["subject"] = Subject
	if options.enabledSetters["Attachments"] {
		params["attachments"] = options.Attachments
	}
	if options.enabledSetters["Bcc"] {
		params["bcc"] = options.Bcc
	}
	if options.enabledSetters["Cc"] {
		params["cc"] = options.Cc
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["Html"] {
		params["html"] = options.Html
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
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

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingUpdateEmailOptions struct {
	Attachments []string
	Bcc []string
	Cc []string
	Content string
	Draft bool
	Html bool
	ScheduledAt string
	Subject string
	Targets []string
	Topics []string
	Users []string
	enabledSetters map[string]bool
}
func (options MessagingUpdateEmailOptions) New() *MessagingUpdateEmailOptions {
	options.enabledSetters = map[string]bool{
		"Attachments": false,
		"Bcc": false,
		"Cc": false,
		"Content": false,
		"Draft": false,
		"Html": false,
		"ScheduledAt": false,
		"Subject": false,
		"Targets": false,
		"Topics": false,
		"Users": false,
	}
	return &options
}
type MessagingUpdateEmailOption func(*MessagingUpdateEmailOptions)
func (srv *Messaging) WithMessagingUpdateEmailAttachments(v []string) MessagingUpdateEmailOption {
	return func(o *MessagingUpdateEmailOptions) {
		o.Attachments = v
		o.enabledSetters["Attachments"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateEmailBcc(v []string) MessagingUpdateEmailOption {
	return func(o *MessagingUpdateEmailOptions) {
		o.Bcc = v
		o.enabledSetters["Bcc"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateEmailCc(v []string) MessagingUpdateEmailOption {
	return func(o *MessagingUpdateEmailOptions) {
		o.Cc = v
		o.enabledSetters["Cc"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateEmailContent(v string) MessagingUpdateEmailOption {
	return func(o *MessagingUpdateEmailOptions) {
		o.Content = v
		o.enabledSetters["Content"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateEmailDraft(v bool) MessagingUpdateEmailOption {
	return func(o *MessagingUpdateEmailOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateEmailHtml(v bool) MessagingUpdateEmailOption {
	return func(o *MessagingUpdateEmailOptions) {
		o.Html = v
		o.enabledSetters["Html"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateEmailScheduledAt(v string) MessagingUpdateEmailOption {
	return func(o *MessagingUpdateEmailOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateEmailSubject(v string) MessagingUpdateEmailOption {
	return func(o *MessagingUpdateEmailOptions) {
		o.Subject = v
		o.enabledSetters["Subject"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateEmailTargets(v []string) MessagingUpdateEmailOption {
	return func(o *MessagingUpdateEmailOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateEmailTopics(v []string) MessagingUpdateEmailOption {
	return func(o *MessagingUpdateEmailOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateEmailUsers(v []string) MessagingUpdateEmailOption {
	return func(o *MessagingUpdateEmailOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
			
// MessagingUpdateEmail update an email message by its unique ID. This
// endpoint only works on messages that are in draft status. Messages that are
// already processing, sent, or failed cannot be updated.
func (srv *Messaging) MessagingUpdateEmail(MessageId string, optionalSetters ...MessagingUpdateEmailOption)(*models.Message, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/v1/messaging/messages/email/{messageId}")
	options := MessagingUpdateEmailOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Attachments"] {
		params["attachments"] = options.Attachments
	}
	if options.enabledSetters["Bcc"] {
		params["bcc"] = options.Bcc
	}
	if options.enabledSetters["Cc"] {
		params["cc"] = options.Cc
	}
	if options.enabledSetters["Content"] {
		params["content"] = options.Content
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["Html"] {
		params["html"] = options.Html
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
	}
	if options.enabledSetters["Subject"] {
		params["subject"] = options.Subject
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
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

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingCreatePushOptions struct {
	Action string
	Badge int
	Body string
	Color string
	ContentAvailable bool
	Critical bool
	Data interface{}
	Draft bool
	Icon string
	Image string
	Priority string
	ScheduledAt string
	Sound string
	Tag string
	Targets []string
	Title string
	Topics []string
	Users []string
	enabledSetters map[string]bool
}
func (options MessagingCreatePushOptions) New() *MessagingCreatePushOptions {
	options.enabledSetters = map[string]bool{
		"Action": false,
		"Badge": false,
		"Body": false,
		"Color": false,
		"ContentAvailable": false,
		"Critical": false,
		"Data": false,
		"Draft": false,
		"Icon": false,
		"Image": false,
		"Priority": false,
		"ScheduledAt": false,
		"Sound": false,
		"Tag": false,
		"Targets": false,
		"Title": false,
		"Topics": false,
		"Users": false,
	}
	return &options
}
type MessagingCreatePushOption func(*MessagingCreatePushOptions)
func (srv *Messaging) WithMessagingCreatePushAction(v string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Action = v
		o.enabledSetters["Action"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushBadge(v int) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Badge = v
		o.enabledSetters["Badge"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushBody(v string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Body = v
		o.enabledSetters["Body"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushColor(v string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Color = v
		o.enabledSetters["Color"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushContentAvailable(v bool) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.ContentAvailable = v
		o.enabledSetters["ContentAvailable"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushCritical(v bool) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Critical = v
		o.enabledSetters["Critical"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushData(v interface{}) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushDraft(v bool) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushIcon(v string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Icon = v
		o.enabledSetters["Icon"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushImage(v string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushPriority(v string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushScheduledAt(v string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushSound(v string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Sound = v
		o.enabledSetters["Sound"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushTag(v string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Tag = v
		o.enabledSetters["Tag"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushTargets(v []string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushTitle(v string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushTopics(v []string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithMessagingCreatePushUsers(v []string) MessagingCreatePushOption {
	return func(o *MessagingCreatePushOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
			
// MessagingCreatePush create a new push notification.
func (srv *Messaging) MessagingCreatePush(MessageId string, optionalSetters ...MessagingCreatePushOption)(*models.Message, error) {
	path := "/v1/messaging/messages/push"
	options := MessagingCreatePushOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Action"] {
		params["action"] = options.Action
	}
	if options.enabledSetters["Badge"] {
		params["badge"] = options.Badge
	}
	if options.enabledSetters["Body"] {
		params["body"] = options.Body
	}
	if options.enabledSetters["Color"] {
		params["color"] = options.Color
	}
	if options.enabledSetters["ContentAvailable"] {
		params["contentAvailable"] = options.ContentAvailable
	}
	if options.enabledSetters["Critical"] {
		params["critical"] = options.Critical
	}
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["Icon"] {
		params["icon"] = options.Icon
	}
	if options.enabledSetters["Image"] {
		params["image"] = options.Image
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
	}
	if options.enabledSetters["Sound"] {
		params["sound"] = options.Sound
	}
	if options.enabledSetters["Tag"] {
		params["tag"] = options.Tag
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
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

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingUpdatePushOptions struct {
	Action string
	Badge int
	Body string
	Color string
	ContentAvailable bool
	Critical bool
	Data interface{}
	Draft bool
	Icon string
	Image string
	Priority string
	ScheduledAt string
	Sound string
	Tag string
	Targets []string
	Title string
	Topics []string
	Users []string
	enabledSetters map[string]bool
}
func (options MessagingUpdatePushOptions) New() *MessagingUpdatePushOptions {
	options.enabledSetters = map[string]bool{
		"Action": false,
		"Badge": false,
		"Body": false,
		"Color": false,
		"ContentAvailable": false,
		"Critical": false,
		"Data": false,
		"Draft": false,
		"Icon": false,
		"Image": false,
		"Priority": false,
		"ScheduledAt": false,
		"Sound": false,
		"Tag": false,
		"Targets": false,
		"Title": false,
		"Topics": false,
		"Users": false,
	}
	return &options
}
type MessagingUpdatePushOption func(*MessagingUpdatePushOptions)
func (srv *Messaging) WithMessagingUpdatePushAction(v string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Action = v
		o.enabledSetters["Action"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushBadge(v int) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Badge = v
		o.enabledSetters["Badge"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushBody(v string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Body = v
		o.enabledSetters["Body"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushColor(v string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Color = v
		o.enabledSetters["Color"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushContentAvailable(v bool) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.ContentAvailable = v
		o.enabledSetters["ContentAvailable"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushCritical(v bool) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Critical = v
		o.enabledSetters["Critical"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushData(v interface{}) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Data = v
		o.enabledSetters["Data"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushDraft(v bool) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Draft = v
		o.enabledSetters["Draft"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushIcon(v string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Icon = v
		o.enabledSetters["Icon"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushImage(v string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Image = v
		o.enabledSetters["Image"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushPriority(v string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Priority = v
		o.enabledSetters["Priority"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushScheduledAt(v string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.ScheduledAt = v
		o.enabledSetters["ScheduledAt"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushSound(v string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Sound = v
		o.enabledSetters["Sound"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushTag(v string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Tag = v
		o.enabledSetters["Tag"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushTargets(v []string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Targets = v
		o.enabledSetters["Targets"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushTitle(v string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Title = v
		o.enabledSetters["Title"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushTopics(v []string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Topics = v
		o.enabledSetters["Topics"] = true
	}
}
func (srv *Messaging) WithMessagingUpdatePushUsers(v []string) MessagingUpdatePushOption {
	return func(o *MessagingUpdatePushOptions) {
		o.Users = v
		o.enabledSetters["Users"] = true
	}
}
			
// MessagingUpdatePush update a push notification by its unique ID. This
// endpoint only works on messages that are in draft status. Messages that are
// already processing, sent, or failed cannot be updated.
func (srv *Messaging) MessagingUpdatePush(MessageId string, optionalSetters ...MessagingUpdatePushOption)(*models.Message, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/v1/messaging/messages/push/{messageId}")
	options := MessagingUpdatePushOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Action"] {
		params["action"] = options.Action
	}
	if options.enabledSetters["Badge"] {
		params["badge"] = options.Badge
	}
	if options.enabledSetters["Body"] {
		params["body"] = options.Body
	}
	if options.enabledSetters["Color"] {
		params["color"] = options.Color
	}
	if options.enabledSetters["ContentAvailable"] {
		params["contentAvailable"] = options.ContentAvailable
	}
	if options.enabledSetters["Critical"] {
		params["critical"] = options.Critical
	}
	if options.enabledSetters["Data"] {
		params["data"] = options.Data
	}
	if options.enabledSetters["Draft"] {
		params["draft"] = options.Draft
	}
	if options.enabledSetters["Icon"] {
		params["icon"] = options.Icon
	}
	if options.enabledSetters["Image"] {
		params["image"] = options.Image
	}
	if options.enabledSetters["Priority"] {
		params["priority"] = options.Priority
	}
	if options.enabledSetters["ScheduledAt"] {
		params["scheduledAt"] = options.ScheduledAt
	}
	if options.enabledSetters["Sound"] {
		params["sound"] = options.Sound
	}
	if options.enabledSetters["Tag"] {
		params["tag"] = options.Tag
	}
	if options.enabledSetters["Targets"] {
		params["targets"] = options.Targets
	}
	if options.enabledSetters["Title"] {
		params["title"] = options.Title
	}
	if options.enabledSetters["Topics"] {
		params["topics"] = options.Topics
	}
	if options.enabledSetters["Users"] {
		params["users"] = options.Users
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

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// MessagingDelete delete a message. If the message is not a draft or
// scheduled, but has been sent, this will not recall the message.
func (srv *Messaging) MessagingDelete(MessageId string)(*interface{}, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/v1/messaging/messages/{messageId}")
	params := map[string]interface{}{}
	params["messageId"] = MessageId
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
	
// MessagingGetMessage get a message by its unique ID.
func (srv *Messaging) MessagingGetMessage(MessageId string)(*models.Message, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/v1/messaging/messages/{messageId}")
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Message{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Message
	parsed, ok := resp.Result.(models.Message)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingListMessageLogsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options MessagingListMessageLogsOptions) New() *MessagingListMessageLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type MessagingListMessageLogsOption func(*MessagingListMessageLogsOptions)
func (srv *Messaging) WithMessagingListMessageLogsQueries(v []string) MessagingListMessageLogsOption {
	return func(o *MessagingListMessageLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithMessagingListMessageLogsTotal(v bool) MessagingListMessageLogsOption {
	return func(o *MessagingListMessageLogsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// MessagingListMessageLogs get the message activity logs listed by its unique
// ID.
func (srv *Messaging) MessagingListMessageLogs(MessageId string, optionalSetters ...MessagingListMessageLogsOption)(*models.LogList, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/v1/messaging/messages/{messageId}/logs")
	options := MessagingListMessageLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.LogList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LogList
	parsed, ok := resp.Result.(models.LogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingListTargetsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options MessagingListTargetsOptions) New() *MessagingListTargetsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type MessagingListTargetsOption func(*MessagingListTargetsOptions)
func (srv *Messaging) WithMessagingListTargetsQueries(v []string) MessagingListTargetsOption {
	return func(o *MessagingListTargetsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithMessagingListTargetsTotal(v bool) MessagingListTargetsOption {
	return func(o *MessagingListTargetsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// MessagingListTargets get a list of the targets associated with a message.
func (srv *Messaging) MessagingListTargets(MessageId string, optionalSetters ...MessagingListTargetsOption)(*models.TargetList, error) {
	r := strings.NewReplacer("{messageId}", MessageId)
	path := r.Replace("/v1/messaging/messages/{messageId}/targets")
	options := MessagingListTargetsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["messageId"] = MessageId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.TargetList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.TargetList
	parsed, ok := resp.Result.(models.TargetList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingListProvidersOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options MessagingListProvidersOptions) New() *MessagingListProvidersOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type MessagingListProvidersOption func(*MessagingListProvidersOptions)
func (srv *Messaging) WithMessagingListProvidersQueries(v []string) MessagingListProvidersOption {
	return func(o *MessagingListProvidersOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithMessagingListProvidersSearch(v string) MessagingListProvidersOption {
	return func(o *MessagingListProvidersOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Messaging) WithMessagingListProvidersTotal(v bool) MessagingListProvidersOption {
	return func(o *MessagingListProvidersOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// MessagingListProviders get a list of all providers from the current
// Revenexx project.
func (srv *Messaging) MessagingListProviders(optionalSetters ...MessagingListProvidersOption)(*models.ProviderList, error) {
	path := "/v1/messaging/providers"
	options := MessagingListProvidersOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ProviderList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ProviderList
	parsed, ok := resp.Result.(models.ProviderList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingCreateMailgunProviderOptions struct {
	ApiKey string
	Domain string
	Enabled bool
	FromEmail string
	FromName string
	IsEuRegion bool
	ReplyToEmail string
	ReplyToName string
	enabledSetters map[string]bool
}
func (options MessagingCreateMailgunProviderOptions) New() *MessagingCreateMailgunProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"Domain": false,
		"Enabled": false,
		"FromEmail": false,
		"FromName": false,
		"IsEuRegion": false,
		"ReplyToEmail": false,
		"ReplyToName": false,
	}
	return &options
}
type MessagingCreateMailgunProviderOption func(*MessagingCreateMailgunProviderOptions)
func (srv *Messaging) WithMessagingCreateMailgunProviderApiKey(v string) MessagingCreateMailgunProviderOption {
	return func(o *MessagingCreateMailgunProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingCreateMailgunProviderDomain(v string) MessagingCreateMailgunProviderOption {
	return func(o *MessagingCreateMailgunProviderOptions) {
		o.Domain = v
		o.enabledSetters["Domain"] = true
	}
}
func (srv *Messaging) WithMessagingCreateMailgunProviderEnabled(v bool) MessagingCreateMailgunProviderOption {
	return func(o *MessagingCreateMailgunProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingCreateMailgunProviderFromEmail(v string) MessagingCreateMailgunProviderOption {
	return func(o *MessagingCreateMailgunProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithMessagingCreateMailgunProviderFromName(v string) MessagingCreateMailgunProviderOption {
	return func(o *MessagingCreateMailgunProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithMessagingCreateMailgunProviderIsEuRegion(v bool) MessagingCreateMailgunProviderOption {
	return func(o *MessagingCreateMailgunProviderOptions) {
		o.IsEuRegion = v
		o.enabledSetters["IsEuRegion"] = true
	}
}
func (srv *Messaging) WithMessagingCreateMailgunProviderReplyToEmail(v string) MessagingCreateMailgunProviderOption {
	return func(o *MessagingCreateMailgunProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithMessagingCreateMailgunProviderReplyToName(v string) MessagingCreateMailgunProviderOption {
	return func(o *MessagingCreateMailgunProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
					
// MessagingCreateMailgunProvider create a new Mailgun provider.
func (srv *Messaging) MessagingCreateMailgunProvider(Name string, ProviderId string, optionalSetters ...MessagingCreateMailgunProviderOption)(*models.Provider, error) {
	path := "/v1/messaging/providers/mailgun"
	options := MessagingCreateMailgunProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Domain"] {
		params["domain"] = options.Domain
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["IsEuRegion"] {
		params["isEuRegion"] = options.IsEuRegion
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingUpdateMailgunProviderOptions struct {
	ApiKey string
	Domain string
	Enabled bool
	FromEmail string
	FromName string
	IsEuRegion bool
	Name string
	ReplyToEmail string
	ReplyToName string
	enabledSetters map[string]bool
}
func (options MessagingUpdateMailgunProviderOptions) New() *MessagingUpdateMailgunProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"Domain": false,
		"Enabled": false,
		"FromEmail": false,
		"FromName": false,
		"IsEuRegion": false,
		"Name": false,
		"ReplyToEmail": false,
		"ReplyToName": false,
	}
	return &options
}
type MessagingUpdateMailgunProviderOption func(*MessagingUpdateMailgunProviderOptions)
func (srv *Messaging) WithMessagingUpdateMailgunProviderApiKey(v string) MessagingUpdateMailgunProviderOption {
	return func(o *MessagingUpdateMailgunProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMailgunProviderDomain(v string) MessagingUpdateMailgunProviderOption {
	return func(o *MessagingUpdateMailgunProviderOptions) {
		o.Domain = v
		o.enabledSetters["Domain"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMailgunProviderEnabled(v bool) MessagingUpdateMailgunProviderOption {
	return func(o *MessagingUpdateMailgunProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMailgunProviderFromEmail(v string) MessagingUpdateMailgunProviderOption {
	return func(o *MessagingUpdateMailgunProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMailgunProviderFromName(v string) MessagingUpdateMailgunProviderOption {
	return func(o *MessagingUpdateMailgunProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMailgunProviderIsEuRegion(v bool) MessagingUpdateMailgunProviderOption {
	return func(o *MessagingUpdateMailgunProviderOptions) {
		o.IsEuRegion = v
		o.enabledSetters["IsEuRegion"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMailgunProviderName(v string) MessagingUpdateMailgunProviderOption {
	return func(o *MessagingUpdateMailgunProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMailgunProviderReplyToEmail(v string) MessagingUpdateMailgunProviderOption {
	return func(o *MessagingUpdateMailgunProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMailgunProviderReplyToName(v string) MessagingUpdateMailgunProviderOption {
	return func(o *MessagingUpdateMailgunProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
			
// MessagingUpdateMailgunProvider update a Mailgun provider by its unique ID.
func (srv *Messaging) MessagingUpdateMailgunProvider(ProviderId string, optionalSetters ...MessagingUpdateMailgunProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/v1/messaging/providers/mailgun/{providerId}")
	options := MessagingUpdateMailgunProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Domain"] {
		params["domain"] = options.Domain
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["IsEuRegion"] {
		params["isEuRegion"] = options.IsEuRegion
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingCreateMsg91ProviderOptions struct {
	AuthKey string
	Enabled bool
	SenderId string
	TemplateId string
	enabledSetters map[string]bool
}
func (options MessagingCreateMsg91ProviderOptions) New() *MessagingCreateMsg91ProviderOptions {
	options.enabledSetters = map[string]bool{
		"AuthKey": false,
		"Enabled": false,
		"SenderId": false,
		"TemplateId": false,
	}
	return &options
}
type MessagingCreateMsg91ProviderOption func(*MessagingCreateMsg91ProviderOptions)
func (srv *Messaging) WithMessagingCreateMsg91ProviderAuthKey(v string) MessagingCreateMsg91ProviderOption {
	return func(o *MessagingCreateMsg91ProviderOptions) {
		o.AuthKey = v
		o.enabledSetters["AuthKey"] = true
	}
}
func (srv *Messaging) WithMessagingCreateMsg91ProviderEnabled(v bool) MessagingCreateMsg91ProviderOption {
	return func(o *MessagingCreateMsg91ProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingCreateMsg91ProviderSenderId(v string) MessagingCreateMsg91ProviderOption {
	return func(o *MessagingCreateMsg91ProviderOptions) {
		o.SenderId = v
		o.enabledSetters["SenderId"] = true
	}
}
func (srv *Messaging) WithMessagingCreateMsg91ProviderTemplateId(v string) MessagingCreateMsg91ProviderOption {
	return func(o *MessagingCreateMsg91ProviderOptions) {
		o.TemplateId = v
		o.enabledSetters["TemplateId"] = true
	}
}
					
// MessagingCreateMsg91Provider create a new MSG91 provider.
func (srv *Messaging) MessagingCreateMsg91Provider(Name string, ProviderId string, optionalSetters ...MessagingCreateMsg91ProviderOption)(*models.Provider, error) {
	path := "/v1/messaging/providers/msg91"
	options := MessagingCreateMsg91ProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["providerId"] = ProviderId
	if options.enabledSetters["AuthKey"] {
		params["authKey"] = options.AuthKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["SenderId"] {
		params["senderId"] = options.SenderId
	}
	if options.enabledSetters["TemplateId"] {
		params["templateId"] = options.TemplateId
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingUpdateMsg91ProviderOptions struct {
	AuthKey string
	Enabled bool
	Name string
	SenderId string
	TemplateId string
	enabledSetters map[string]bool
}
func (options MessagingUpdateMsg91ProviderOptions) New() *MessagingUpdateMsg91ProviderOptions {
	options.enabledSetters = map[string]bool{
		"AuthKey": false,
		"Enabled": false,
		"Name": false,
		"SenderId": false,
		"TemplateId": false,
	}
	return &options
}
type MessagingUpdateMsg91ProviderOption func(*MessagingUpdateMsg91ProviderOptions)
func (srv *Messaging) WithMessagingUpdateMsg91ProviderAuthKey(v string) MessagingUpdateMsg91ProviderOption {
	return func(o *MessagingUpdateMsg91ProviderOptions) {
		o.AuthKey = v
		o.enabledSetters["AuthKey"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMsg91ProviderEnabled(v bool) MessagingUpdateMsg91ProviderOption {
	return func(o *MessagingUpdateMsg91ProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMsg91ProviderName(v string) MessagingUpdateMsg91ProviderOption {
	return func(o *MessagingUpdateMsg91ProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMsg91ProviderSenderId(v string) MessagingUpdateMsg91ProviderOption {
	return func(o *MessagingUpdateMsg91ProviderOptions) {
		o.SenderId = v
		o.enabledSetters["SenderId"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateMsg91ProviderTemplateId(v string) MessagingUpdateMsg91ProviderOption {
	return func(o *MessagingUpdateMsg91ProviderOptions) {
		o.TemplateId = v
		o.enabledSetters["TemplateId"] = true
	}
}
			
// MessagingUpdateMsg91Provider update a MSG91 provider by its unique ID.
func (srv *Messaging) MessagingUpdateMsg91Provider(ProviderId string, optionalSetters ...MessagingUpdateMsg91ProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/v1/messaging/providers/msg91/{providerId}")
	options := MessagingUpdateMsg91ProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["AuthKey"] {
		params["authKey"] = options.AuthKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["SenderId"] {
		params["senderId"] = options.SenderId
	}
	if options.enabledSetters["TemplateId"] {
		params["templateId"] = options.TemplateId
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingCreateResendProviderOptions struct {
	ApiKey string
	Enabled bool
	FromEmail string
	FromName string
	ReplyToEmail string
	ReplyToName string
	enabledSetters map[string]bool
}
func (options MessagingCreateResendProviderOptions) New() *MessagingCreateResendProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"Enabled": false,
		"FromEmail": false,
		"FromName": false,
		"ReplyToEmail": false,
		"ReplyToName": false,
	}
	return &options
}
type MessagingCreateResendProviderOption func(*MessagingCreateResendProviderOptions)
func (srv *Messaging) WithMessagingCreateResendProviderApiKey(v string) MessagingCreateResendProviderOption {
	return func(o *MessagingCreateResendProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingCreateResendProviderEnabled(v bool) MessagingCreateResendProviderOption {
	return func(o *MessagingCreateResendProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingCreateResendProviderFromEmail(v string) MessagingCreateResendProviderOption {
	return func(o *MessagingCreateResendProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithMessagingCreateResendProviderFromName(v string) MessagingCreateResendProviderOption {
	return func(o *MessagingCreateResendProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithMessagingCreateResendProviderReplyToEmail(v string) MessagingCreateResendProviderOption {
	return func(o *MessagingCreateResendProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithMessagingCreateResendProviderReplyToName(v string) MessagingCreateResendProviderOption {
	return func(o *MessagingCreateResendProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
					
// MessagingCreateResendProvider create a new Resend provider.
func (srv *Messaging) MessagingCreateResendProvider(Name string, ProviderId string, optionalSetters ...MessagingCreateResendProviderOption)(*models.Provider, error) {
	path := "/v1/messaging/providers/resend"
	options := MessagingCreateResendProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingUpdateResendProviderOptions struct {
	ApiKey string
	Enabled bool
	FromEmail string
	FromName string
	Name string
	ReplyToEmail string
	ReplyToName string
	enabledSetters map[string]bool
}
func (options MessagingUpdateResendProviderOptions) New() *MessagingUpdateResendProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"Enabled": false,
		"FromEmail": false,
		"FromName": false,
		"Name": false,
		"ReplyToEmail": false,
		"ReplyToName": false,
	}
	return &options
}
type MessagingUpdateResendProviderOption func(*MessagingUpdateResendProviderOptions)
func (srv *Messaging) WithMessagingUpdateResendProviderApiKey(v string) MessagingUpdateResendProviderOption {
	return func(o *MessagingUpdateResendProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateResendProviderEnabled(v bool) MessagingUpdateResendProviderOption {
	return func(o *MessagingUpdateResendProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateResendProviderFromEmail(v string) MessagingUpdateResendProviderOption {
	return func(o *MessagingUpdateResendProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateResendProviderFromName(v string) MessagingUpdateResendProviderOption {
	return func(o *MessagingUpdateResendProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateResendProviderName(v string) MessagingUpdateResendProviderOption {
	return func(o *MessagingUpdateResendProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateResendProviderReplyToEmail(v string) MessagingUpdateResendProviderOption {
	return func(o *MessagingUpdateResendProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateResendProviderReplyToName(v string) MessagingUpdateResendProviderOption {
	return func(o *MessagingUpdateResendProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
			
// MessagingUpdateResendProvider update a Resend provider by its unique ID.
func (srv *Messaging) MessagingUpdateResendProvider(ProviderId string, optionalSetters ...MessagingUpdateResendProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/v1/messaging/providers/resend/{providerId}")
	options := MessagingUpdateResendProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingCreateSendgridProviderOptions struct {
	ApiKey string
	Enabled bool
	FromEmail string
	FromName string
	ReplyToEmail string
	ReplyToName string
	enabledSetters map[string]bool
}
func (options MessagingCreateSendgridProviderOptions) New() *MessagingCreateSendgridProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"Enabled": false,
		"FromEmail": false,
		"FromName": false,
		"ReplyToEmail": false,
		"ReplyToName": false,
	}
	return &options
}
type MessagingCreateSendgridProviderOption func(*MessagingCreateSendgridProviderOptions)
func (srv *Messaging) WithMessagingCreateSendgridProviderApiKey(v string) MessagingCreateSendgridProviderOption {
	return func(o *MessagingCreateSendgridProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingCreateSendgridProviderEnabled(v bool) MessagingCreateSendgridProviderOption {
	return func(o *MessagingCreateSendgridProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingCreateSendgridProviderFromEmail(v string) MessagingCreateSendgridProviderOption {
	return func(o *MessagingCreateSendgridProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithMessagingCreateSendgridProviderFromName(v string) MessagingCreateSendgridProviderOption {
	return func(o *MessagingCreateSendgridProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithMessagingCreateSendgridProviderReplyToEmail(v string) MessagingCreateSendgridProviderOption {
	return func(o *MessagingCreateSendgridProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithMessagingCreateSendgridProviderReplyToName(v string) MessagingCreateSendgridProviderOption {
	return func(o *MessagingCreateSendgridProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
					
// MessagingCreateSendgridProvider create a new Sendgrid provider.
func (srv *Messaging) MessagingCreateSendgridProvider(Name string, ProviderId string, optionalSetters ...MessagingCreateSendgridProviderOption)(*models.Provider, error) {
	path := "/v1/messaging/providers/sendgrid"
	options := MessagingCreateSendgridProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingUpdateSendgridProviderOptions struct {
	ApiKey string
	Enabled bool
	FromEmail string
	FromName string
	Name string
	ReplyToEmail string
	ReplyToName string
	enabledSetters map[string]bool
}
func (options MessagingUpdateSendgridProviderOptions) New() *MessagingUpdateSendgridProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"Enabled": false,
		"FromEmail": false,
		"FromName": false,
		"Name": false,
		"ReplyToEmail": false,
		"ReplyToName": false,
	}
	return &options
}
type MessagingUpdateSendgridProviderOption func(*MessagingUpdateSendgridProviderOptions)
func (srv *Messaging) WithMessagingUpdateSendgridProviderApiKey(v string) MessagingUpdateSendgridProviderOption {
	return func(o *MessagingUpdateSendgridProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateSendgridProviderEnabled(v bool) MessagingUpdateSendgridProviderOption {
	return func(o *MessagingUpdateSendgridProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateSendgridProviderFromEmail(v string) MessagingUpdateSendgridProviderOption {
	return func(o *MessagingUpdateSendgridProviderOptions) {
		o.FromEmail = v
		o.enabledSetters["FromEmail"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateSendgridProviderFromName(v string) MessagingUpdateSendgridProviderOption {
	return func(o *MessagingUpdateSendgridProviderOptions) {
		o.FromName = v
		o.enabledSetters["FromName"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateSendgridProviderName(v string) MessagingUpdateSendgridProviderOption {
	return func(o *MessagingUpdateSendgridProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateSendgridProviderReplyToEmail(v string) MessagingUpdateSendgridProviderOption {
	return func(o *MessagingUpdateSendgridProviderOptions) {
		o.ReplyToEmail = v
		o.enabledSetters["ReplyToEmail"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateSendgridProviderReplyToName(v string) MessagingUpdateSendgridProviderOption {
	return func(o *MessagingUpdateSendgridProviderOptions) {
		o.ReplyToName = v
		o.enabledSetters["ReplyToName"] = true
	}
}
			
// MessagingUpdateSendgridProvider update a Sendgrid provider by its unique
// ID.
func (srv *Messaging) MessagingUpdateSendgridProvider(ProviderId string, optionalSetters ...MessagingUpdateSendgridProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/v1/messaging/providers/sendgrid/{providerId}")
	options := MessagingUpdateSendgridProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["FromEmail"] {
		params["fromEmail"] = options.FromEmail
	}
	if options.enabledSetters["FromName"] {
		params["fromName"] = options.FromName
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["ReplyToEmail"] {
		params["replyToEmail"] = options.ReplyToEmail
	}
	if options.enabledSetters["ReplyToName"] {
		params["replyToName"] = options.ReplyToName
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingCreateTelesignProviderOptions struct {
	ApiKey string
	CustomerId string
	Enabled bool
	From string
	enabledSetters map[string]bool
}
func (options MessagingCreateTelesignProviderOptions) New() *MessagingCreateTelesignProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"CustomerId": false,
		"Enabled": false,
		"From": false,
	}
	return &options
}
type MessagingCreateTelesignProviderOption func(*MessagingCreateTelesignProviderOptions)
func (srv *Messaging) WithMessagingCreateTelesignProviderApiKey(v string) MessagingCreateTelesignProviderOption {
	return func(o *MessagingCreateTelesignProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingCreateTelesignProviderCustomerId(v string) MessagingCreateTelesignProviderOption {
	return func(o *MessagingCreateTelesignProviderOptions) {
		o.CustomerId = v
		o.enabledSetters["CustomerId"] = true
	}
}
func (srv *Messaging) WithMessagingCreateTelesignProviderEnabled(v bool) MessagingCreateTelesignProviderOption {
	return func(o *MessagingCreateTelesignProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingCreateTelesignProviderFrom(v string) MessagingCreateTelesignProviderOption {
	return func(o *MessagingCreateTelesignProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
					
// MessagingCreateTelesignProvider create a new Telesign provider.
func (srv *Messaging) MessagingCreateTelesignProvider(Name string, ProviderId string, optionalSetters ...MessagingCreateTelesignProviderOption)(*models.Provider, error) {
	path := "/v1/messaging/providers/telesign"
	options := MessagingCreateTelesignProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["CustomerId"] {
		params["customerId"] = options.CustomerId
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingUpdateTelesignProviderOptions struct {
	ApiKey string
	CustomerId string
	Enabled bool
	From string
	Name string
	enabledSetters map[string]bool
}
func (options MessagingUpdateTelesignProviderOptions) New() *MessagingUpdateTelesignProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"CustomerId": false,
		"Enabled": false,
		"From": false,
		"Name": false,
	}
	return &options
}
type MessagingUpdateTelesignProviderOption func(*MessagingUpdateTelesignProviderOptions)
func (srv *Messaging) WithMessagingUpdateTelesignProviderApiKey(v string) MessagingUpdateTelesignProviderOption {
	return func(o *MessagingUpdateTelesignProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTelesignProviderCustomerId(v string) MessagingUpdateTelesignProviderOption {
	return func(o *MessagingUpdateTelesignProviderOptions) {
		o.CustomerId = v
		o.enabledSetters["CustomerId"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTelesignProviderEnabled(v bool) MessagingUpdateTelesignProviderOption {
	return func(o *MessagingUpdateTelesignProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTelesignProviderFrom(v string) MessagingUpdateTelesignProviderOption {
	return func(o *MessagingUpdateTelesignProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTelesignProviderName(v string) MessagingUpdateTelesignProviderOption {
	return func(o *MessagingUpdateTelesignProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
			
// MessagingUpdateTelesignProvider update a Telesign provider by its unique
// ID.
func (srv *Messaging) MessagingUpdateTelesignProvider(ProviderId string, optionalSetters ...MessagingUpdateTelesignProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/v1/messaging/providers/telesign/{providerId}")
	options := MessagingUpdateTelesignProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["CustomerId"] {
		params["customerId"] = options.CustomerId
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingCreateTextmagicProviderOptions struct {
	ApiKey string
	Enabled bool
	From string
	Username string
	enabledSetters map[string]bool
}
func (options MessagingCreateTextmagicProviderOptions) New() *MessagingCreateTextmagicProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"Enabled": false,
		"From": false,
		"Username": false,
	}
	return &options
}
type MessagingCreateTextmagicProviderOption func(*MessagingCreateTextmagicProviderOptions)
func (srv *Messaging) WithMessagingCreateTextmagicProviderApiKey(v string) MessagingCreateTextmagicProviderOption {
	return func(o *MessagingCreateTextmagicProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingCreateTextmagicProviderEnabled(v bool) MessagingCreateTextmagicProviderOption {
	return func(o *MessagingCreateTextmagicProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingCreateTextmagicProviderFrom(v string) MessagingCreateTextmagicProviderOption {
	return func(o *MessagingCreateTextmagicProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
func (srv *Messaging) WithMessagingCreateTextmagicProviderUsername(v string) MessagingCreateTextmagicProviderOption {
	return func(o *MessagingCreateTextmagicProviderOptions) {
		o.Username = v
		o.enabledSetters["Username"] = true
	}
}
					
// MessagingCreateTextmagicProvider create a new Textmagic provider.
func (srv *Messaging) MessagingCreateTextmagicProvider(Name string, ProviderId string, optionalSetters ...MessagingCreateTextmagicProviderOption)(*models.Provider, error) {
	path := "/v1/messaging/providers/textmagic"
	options := MessagingCreateTextmagicProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	if options.enabledSetters["Username"] {
		params["username"] = options.Username
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingUpdateTextmagicProviderOptions struct {
	ApiKey string
	Enabled bool
	From string
	Name string
	Username string
	enabledSetters map[string]bool
}
func (options MessagingUpdateTextmagicProviderOptions) New() *MessagingUpdateTextmagicProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"Enabled": false,
		"From": false,
		"Name": false,
		"Username": false,
	}
	return &options
}
type MessagingUpdateTextmagicProviderOption func(*MessagingUpdateTextmagicProviderOptions)
func (srv *Messaging) WithMessagingUpdateTextmagicProviderApiKey(v string) MessagingUpdateTextmagicProviderOption {
	return func(o *MessagingUpdateTextmagicProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTextmagicProviderEnabled(v bool) MessagingUpdateTextmagicProviderOption {
	return func(o *MessagingUpdateTextmagicProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTextmagicProviderFrom(v string) MessagingUpdateTextmagicProviderOption {
	return func(o *MessagingUpdateTextmagicProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTextmagicProviderName(v string) MessagingUpdateTextmagicProviderOption {
	return func(o *MessagingUpdateTextmagicProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTextmagicProviderUsername(v string) MessagingUpdateTextmagicProviderOption {
	return func(o *MessagingUpdateTextmagicProviderOptions) {
		o.Username = v
		o.enabledSetters["Username"] = true
	}
}
			
// MessagingUpdateTextmagicProvider update a Textmagic provider by its unique
// ID.
func (srv *Messaging) MessagingUpdateTextmagicProvider(ProviderId string, optionalSetters ...MessagingUpdateTextmagicProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/v1/messaging/providers/textmagic/{providerId}")
	options := MessagingUpdateTextmagicProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Username"] {
		params["username"] = options.Username
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingCreateTwilioProviderOptions struct {
	AccountSid string
	AuthToken string
	Enabled bool
	From string
	enabledSetters map[string]bool
}
func (options MessagingCreateTwilioProviderOptions) New() *MessagingCreateTwilioProviderOptions {
	options.enabledSetters = map[string]bool{
		"AccountSid": false,
		"AuthToken": false,
		"Enabled": false,
		"From": false,
	}
	return &options
}
type MessagingCreateTwilioProviderOption func(*MessagingCreateTwilioProviderOptions)
func (srv *Messaging) WithMessagingCreateTwilioProviderAccountSid(v string) MessagingCreateTwilioProviderOption {
	return func(o *MessagingCreateTwilioProviderOptions) {
		o.AccountSid = v
		o.enabledSetters["AccountSid"] = true
	}
}
func (srv *Messaging) WithMessagingCreateTwilioProviderAuthToken(v string) MessagingCreateTwilioProviderOption {
	return func(o *MessagingCreateTwilioProviderOptions) {
		o.AuthToken = v
		o.enabledSetters["AuthToken"] = true
	}
}
func (srv *Messaging) WithMessagingCreateTwilioProviderEnabled(v bool) MessagingCreateTwilioProviderOption {
	return func(o *MessagingCreateTwilioProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingCreateTwilioProviderFrom(v string) MessagingCreateTwilioProviderOption {
	return func(o *MessagingCreateTwilioProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
					
// MessagingCreateTwilioProvider create a new Twilio provider.
func (srv *Messaging) MessagingCreateTwilioProvider(Name string, ProviderId string, optionalSetters ...MessagingCreateTwilioProviderOption)(*models.Provider, error) {
	path := "/v1/messaging/providers/twilio"
	options := MessagingCreateTwilioProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["providerId"] = ProviderId
	if options.enabledSetters["AccountSid"] {
		params["accountSid"] = options.AccountSid
	}
	if options.enabledSetters["AuthToken"] {
		params["authToken"] = options.AuthToken
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingUpdateTwilioProviderOptions struct {
	AccountSid string
	AuthToken string
	Enabled bool
	From string
	Name string
	enabledSetters map[string]bool
}
func (options MessagingUpdateTwilioProviderOptions) New() *MessagingUpdateTwilioProviderOptions {
	options.enabledSetters = map[string]bool{
		"AccountSid": false,
		"AuthToken": false,
		"Enabled": false,
		"From": false,
		"Name": false,
	}
	return &options
}
type MessagingUpdateTwilioProviderOption func(*MessagingUpdateTwilioProviderOptions)
func (srv *Messaging) WithMessagingUpdateTwilioProviderAccountSid(v string) MessagingUpdateTwilioProviderOption {
	return func(o *MessagingUpdateTwilioProviderOptions) {
		o.AccountSid = v
		o.enabledSetters["AccountSid"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTwilioProviderAuthToken(v string) MessagingUpdateTwilioProviderOption {
	return func(o *MessagingUpdateTwilioProviderOptions) {
		o.AuthToken = v
		o.enabledSetters["AuthToken"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTwilioProviderEnabled(v bool) MessagingUpdateTwilioProviderOption {
	return func(o *MessagingUpdateTwilioProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTwilioProviderFrom(v string) MessagingUpdateTwilioProviderOption {
	return func(o *MessagingUpdateTwilioProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTwilioProviderName(v string) MessagingUpdateTwilioProviderOption {
	return func(o *MessagingUpdateTwilioProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
			
// MessagingUpdateTwilioProvider update a Twilio provider by its unique ID.
func (srv *Messaging) MessagingUpdateTwilioProvider(ProviderId string, optionalSetters ...MessagingUpdateTwilioProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/v1/messaging/providers/twilio/{providerId}")
	options := MessagingUpdateTwilioProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["AccountSid"] {
		params["accountSid"] = options.AccountSid
	}
	if options.enabledSetters["AuthToken"] {
		params["authToken"] = options.AuthToken
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingCreateVonageProviderOptions struct {
	ApiKey string
	ApiSecret string
	Enabled bool
	From string
	enabledSetters map[string]bool
}
func (options MessagingCreateVonageProviderOptions) New() *MessagingCreateVonageProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"ApiSecret": false,
		"Enabled": false,
		"From": false,
	}
	return &options
}
type MessagingCreateVonageProviderOption func(*MessagingCreateVonageProviderOptions)
func (srv *Messaging) WithMessagingCreateVonageProviderApiKey(v string) MessagingCreateVonageProviderOption {
	return func(o *MessagingCreateVonageProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingCreateVonageProviderApiSecret(v string) MessagingCreateVonageProviderOption {
	return func(o *MessagingCreateVonageProviderOptions) {
		o.ApiSecret = v
		o.enabledSetters["ApiSecret"] = true
	}
}
func (srv *Messaging) WithMessagingCreateVonageProviderEnabled(v bool) MessagingCreateVonageProviderOption {
	return func(o *MessagingCreateVonageProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingCreateVonageProviderFrom(v string) MessagingCreateVonageProviderOption {
	return func(o *MessagingCreateVonageProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
					
// MessagingCreateVonageProvider create a new Vonage provider.
func (srv *Messaging) MessagingCreateVonageProvider(Name string, ProviderId string, optionalSetters ...MessagingCreateVonageProviderOption)(*models.Provider, error) {
	path := "/v1/messaging/providers/vonage"
	options := MessagingCreateVonageProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["ApiSecret"] {
		params["apiSecret"] = options.ApiSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingUpdateVonageProviderOptions struct {
	ApiKey string
	ApiSecret string
	Enabled bool
	From string
	Name string
	enabledSetters map[string]bool
}
func (options MessagingUpdateVonageProviderOptions) New() *MessagingUpdateVonageProviderOptions {
	options.enabledSetters = map[string]bool{
		"ApiKey": false,
		"ApiSecret": false,
		"Enabled": false,
		"From": false,
		"Name": false,
	}
	return &options
}
type MessagingUpdateVonageProviderOption func(*MessagingUpdateVonageProviderOptions)
func (srv *Messaging) WithMessagingUpdateVonageProviderApiKey(v string) MessagingUpdateVonageProviderOption {
	return func(o *MessagingUpdateVonageProviderOptions) {
		o.ApiKey = v
		o.enabledSetters["ApiKey"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateVonageProviderApiSecret(v string) MessagingUpdateVonageProviderOption {
	return func(o *MessagingUpdateVonageProviderOptions) {
		o.ApiSecret = v
		o.enabledSetters["ApiSecret"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateVonageProviderEnabled(v bool) MessagingUpdateVonageProviderOption {
	return func(o *MessagingUpdateVonageProviderOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateVonageProviderFrom(v string) MessagingUpdateVonageProviderOption {
	return func(o *MessagingUpdateVonageProviderOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateVonageProviderName(v string) MessagingUpdateVonageProviderOption {
	return func(o *MessagingUpdateVonageProviderOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
			
// MessagingUpdateVonageProvider update a Vonage provider by its unique ID.
func (srv *Messaging) MessagingUpdateVonageProvider(ProviderId string, optionalSetters ...MessagingUpdateVonageProviderOption)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/v1/messaging/providers/vonage/{providerId}")
	options := MessagingUpdateVonageProviderOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["ApiKey"] {
		params["apiKey"] = options.ApiKey
	}
	if options.enabledSetters["ApiSecret"] {
		params["apiSecret"] = options.ApiSecret
	}
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["From"] {
		params["from"] = options.From
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
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

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// MessagingDeleteProvider delete a provider by its unique ID.
func (srv *Messaging) MessagingDeleteProvider(ProviderId string)(*interface{}, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/v1/messaging/providers/{providerId}")
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
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
	
// MessagingGetProvider get a provider by its unique ID.
func (srv *Messaging) MessagingGetProvider(ProviderId string)(*models.Provider, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/v1/messaging/providers/{providerId}")
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Provider{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Provider
	parsed, ok := resp.Result.(models.Provider)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingListProviderLogsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options MessagingListProviderLogsOptions) New() *MessagingListProviderLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type MessagingListProviderLogsOption func(*MessagingListProviderLogsOptions)
func (srv *Messaging) WithMessagingListProviderLogsQueries(v []string) MessagingListProviderLogsOption {
	return func(o *MessagingListProviderLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithMessagingListProviderLogsTotal(v bool) MessagingListProviderLogsOption {
	return func(o *MessagingListProviderLogsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// MessagingListProviderLogs get the provider activity logs listed by its
// unique ID.
func (srv *Messaging) MessagingListProviderLogs(ProviderId string, optionalSetters ...MessagingListProviderLogsOption)(*models.LogList, error) {
	r := strings.NewReplacer("{providerId}", ProviderId)
	path := r.Replace("/v1/messaging/providers/{providerId}/logs")
	options := MessagingListProviderLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["providerId"] = ProviderId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.LogList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LogList
	parsed, ok := resp.Result.(models.LogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingListSubscriberLogsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options MessagingListSubscriberLogsOptions) New() *MessagingListSubscriberLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type MessagingListSubscriberLogsOption func(*MessagingListSubscriberLogsOptions)
func (srv *Messaging) WithMessagingListSubscriberLogsQueries(v []string) MessagingListSubscriberLogsOption {
	return func(o *MessagingListSubscriberLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithMessagingListSubscriberLogsTotal(v bool) MessagingListSubscriberLogsOption {
	return func(o *MessagingListSubscriberLogsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// MessagingListSubscriberLogs get the subscriber activity logs listed by its
// unique ID.
func (srv *Messaging) MessagingListSubscriberLogs(SubscriberId string, optionalSetters ...MessagingListSubscriberLogsOption)(*models.LogList, error) {
	r := strings.NewReplacer("{subscriberId}", SubscriberId)
	path := r.Replace("/v1/messaging/subscribers/{subscriberId}/logs")
	options := MessagingListSubscriberLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["subscriberId"] = SubscriberId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.LogList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LogList
	parsed, ok := resp.Result.(models.LogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingListTopicsOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options MessagingListTopicsOptions) New() *MessagingListTopicsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type MessagingListTopicsOption func(*MessagingListTopicsOptions)
func (srv *Messaging) WithMessagingListTopicsQueries(v []string) MessagingListTopicsOption {
	return func(o *MessagingListTopicsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithMessagingListTopicsSearch(v string) MessagingListTopicsOption {
	return func(o *MessagingListTopicsOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Messaging) WithMessagingListTopicsTotal(v bool) MessagingListTopicsOption {
	return func(o *MessagingListTopicsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
	
// MessagingListTopics get a list of all topics from the current Revenexx
// project.
func (srv *Messaging) MessagingListTopics(optionalSetters ...MessagingListTopicsOption)(*models.TopicList, error) {
	path := "/v1/messaging/topics"
	options := MessagingListTopicsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.TopicList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.TopicList
	parsed, ok := resp.Result.(models.TopicList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingCreateTopicOptions struct {
	Subscribe []string
	enabledSetters map[string]bool
}
func (options MessagingCreateTopicOptions) New() *MessagingCreateTopicOptions {
	options.enabledSetters = map[string]bool{
		"Subscribe": false,
	}
	return &options
}
type MessagingCreateTopicOption func(*MessagingCreateTopicOptions)
func (srv *Messaging) WithMessagingCreateTopicSubscribe(v []string) MessagingCreateTopicOption {
	return func(o *MessagingCreateTopicOptions) {
		o.Subscribe = v
		o.enabledSetters["Subscribe"] = true
	}
}
					
// MessagingCreateTopic create a new topic.
func (srv *Messaging) MessagingCreateTopic(Name string, TopicId string, optionalSetters ...MessagingCreateTopicOption)(*models.Topic, error) {
	path := "/v1/messaging/topics"
	options := MessagingCreateTopicOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	params["topicId"] = TopicId
	if options.enabledSetters["Subscribe"] {
		params["subscribe"] = options.Subscribe
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

		parsed := models.Topic{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Topic
	parsed, ok := resp.Result.(models.Topic)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// MessagingDeleteTopic delete a topic by its unique ID.
func (srv *Messaging) MessagingDeleteTopic(TopicId string)(*interface{}, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/v1/messaging/topics/{topicId}")
	params := map[string]interface{}{}
	params["topicId"] = TopicId
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
	
// MessagingGetTopic get a topic by its unique ID.
func (srv *Messaging) MessagingGetTopic(TopicId string)(*models.Topic, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/v1/messaging/topics/{topicId}")
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Topic{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Topic
	parsed, ok := resp.Result.(models.Topic)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingUpdateTopicOptions struct {
	Name string
	Subscribe []string
	enabledSetters map[string]bool
}
func (options MessagingUpdateTopicOptions) New() *MessagingUpdateTopicOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Subscribe": false,
	}
	return &options
}
type MessagingUpdateTopicOption func(*MessagingUpdateTopicOptions)
func (srv *Messaging) WithMessagingUpdateTopicName(v string) MessagingUpdateTopicOption {
	return func(o *MessagingUpdateTopicOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Messaging) WithMessagingUpdateTopicSubscribe(v []string) MessagingUpdateTopicOption {
	return func(o *MessagingUpdateTopicOptions) {
		o.Subscribe = v
		o.enabledSetters["Subscribe"] = true
	}
}
			
// MessagingUpdateTopic update a topic by its unique ID.
func (srv *Messaging) MessagingUpdateTopic(TopicId string, optionalSetters ...MessagingUpdateTopicOption)(*models.Topic, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/v1/messaging/topics/{topicId}")
	options := MessagingUpdateTopicOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Subscribe"] {
		params["subscribe"] = options.Subscribe
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

		parsed := models.Topic{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Topic
	parsed, ok := resp.Result.(models.Topic)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingListTopicLogsOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options MessagingListTopicLogsOptions) New() *MessagingListTopicLogsOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type MessagingListTopicLogsOption func(*MessagingListTopicLogsOptions)
func (srv *Messaging) WithMessagingListTopicLogsQueries(v []string) MessagingListTopicLogsOption {
	return func(o *MessagingListTopicLogsOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithMessagingListTopicLogsTotal(v bool) MessagingListTopicLogsOption {
	return func(o *MessagingListTopicLogsOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// MessagingListTopicLogs get the topic activity logs listed by its unique ID.
func (srv *Messaging) MessagingListTopicLogs(TopicId string, optionalSetters ...MessagingListTopicLogsOption)(*models.LogList, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/v1/messaging/topics/{topicId}/logs")
	options := MessagingListTopicLogsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.LogList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.LogList
	parsed, ok := resp.Result.(models.LogList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type MessagingListSubscribersOptions struct {
	Queries []string
	Search string
	Total bool
	enabledSetters map[string]bool
}
func (options MessagingListSubscribersOptions) New() *MessagingListSubscribersOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Search": false,
		"Total": false,
	}
	return &options
}
type MessagingListSubscribersOption func(*MessagingListSubscribersOptions)
func (srv *Messaging) WithMessagingListSubscribersQueries(v []string) MessagingListSubscribersOption {
	return func(o *MessagingListSubscribersOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Messaging) WithMessagingListSubscribersSearch(v string) MessagingListSubscribersOption {
	return func(o *MessagingListSubscribersOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
func (srv *Messaging) WithMessagingListSubscribersTotal(v bool) MessagingListSubscribersOption {
	return func(o *MessagingListSubscribersOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
			
// MessagingListSubscribers get a list of all subscribers from the current
// Revenexx project.
func (srv *Messaging) MessagingListSubscribers(TopicId string, optionalSetters ...MessagingListSubscribersOption)(*models.SubscriberList, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/v1/messaging/topics/{topicId}/subscribers")
	options := MessagingListSubscribersOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	if options.enabledSetters["Queries"] {
		params["queries"] = options.Queries
	}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
	}
	if options.enabledSetters["Total"] {
		params["total"] = options.Total
	}
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.SubscriberList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.SubscriberList
	parsed, ok := resp.Result.(models.SubscriberList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
					
// MessagingCreateSubscriber create a new subscriber.
func (srv *Messaging) MessagingCreateSubscriber(TopicId string, SubscriberId string, TargetId string)(*models.Subscriber, error) {
	r := strings.NewReplacer("{topicId}", TopicId)
	path := r.Replace("/v1/messaging/topics/{topicId}/subscribers")
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	params["subscriberId"] = SubscriberId
	params["targetId"] = TargetId
	headers := map[string]interface{}{
		"content-type": "application/json",
	}

	resp, err := srv.client.Call("POST", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Subscriber{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Subscriber
	parsed, ok := resp.Result.(models.Subscriber)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
			
// MessagingDeleteSubscriber delete a subscriber by its unique ID.
func (srv *Messaging) MessagingDeleteSubscriber(TopicId string, SubscriberId string)(*interface{}, error) {
	r := strings.NewReplacer("{topicId}", TopicId, "{subscriberId}", SubscriberId)
	path := r.Replace("/v1/messaging/topics/{topicId}/subscribers/{subscriberId}")
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	params["subscriberId"] = SubscriberId
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
			
// MessagingGetSubscriber get a subscriber by its unique ID.
func (srv *Messaging) MessagingGetSubscriber(TopicId string, SubscriberId string)(*models.Subscriber, error) {
	r := strings.NewReplacer("{topicId}", TopicId, "{subscriberId}", SubscriberId)
	path := r.Replace("/v1/messaging/topics/{topicId}/subscribers/{subscriberId}")
	params := map[string]interface{}{}
	params["topicId"] = TopicId
	params["subscriberId"] = SubscriberId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.Subscriber{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.Subscriber
	parsed, ok := resp.Result.(models.Subscriber)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
