package avatars

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"strings"
)

// Avatars service
type Avatars struct {
	client client.Client
}

func New(clt client.Client) *Avatars {
	return &Avatars{
		client: clt,
	}
}

type AvatarsGetBrowserOptions struct {
	Width int
	Height int
	Quality int
	enabledSetters map[string]bool
}
func (options AvatarsGetBrowserOptions) New() *AvatarsGetBrowserOptions {
	options.enabledSetters = map[string]bool{
		"Width": false,
		"Height": false,
		"Quality": false,
	}
	return &options
}
type AvatarsGetBrowserOption func(*AvatarsGetBrowserOptions)
func (srv *Avatars) WithAvatarsGetBrowserWidth(v int) AvatarsGetBrowserOption {
	return func(o *AvatarsGetBrowserOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithAvatarsGetBrowserHeight(v int) AvatarsGetBrowserOption {
	return func(o *AvatarsGetBrowserOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Avatars) WithAvatarsGetBrowserQuality(v int) AvatarsGetBrowserOption {
	return func(o *AvatarsGetBrowserOptions) {
		o.Quality = v
		o.enabledSetters["Quality"] = true
	}
}
			
// AvatarsGetBrowser you can use this endpoint to show different browser icons
// to your users. The code argument receives the browser code as it appears in
// your user [GET
// /account/sessions](https://app.revenexx.com/docs/references/cloud/client-web/account#getSessions)
// endpoint. Use width, height and quality arguments to change the output
// settings.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
func (srv *Avatars) AvatarsGetBrowser(Code string, optionalSetters ...AvatarsGetBrowserOption)(*interface{}, error) {
	r := strings.NewReplacer("{code}", Code)
	path := r.Replace("/v1/avatars/browsers/{code}")
	options := AvatarsGetBrowserOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["Width"] {
		params["width"] = options.Width
	}
	if options.enabledSetters["Height"] {
		params["height"] = options.Height
	}
	if options.enabledSetters["Quality"] {
		params["quality"] = options.Quality
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
type AvatarsGetCreditCardOptions struct {
	Width int
	Height int
	Quality int
	enabledSetters map[string]bool
}
func (options AvatarsGetCreditCardOptions) New() *AvatarsGetCreditCardOptions {
	options.enabledSetters = map[string]bool{
		"Width": false,
		"Height": false,
		"Quality": false,
	}
	return &options
}
type AvatarsGetCreditCardOption func(*AvatarsGetCreditCardOptions)
func (srv *Avatars) WithAvatarsGetCreditCardWidth(v int) AvatarsGetCreditCardOption {
	return func(o *AvatarsGetCreditCardOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithAvatarsGetCreditCardHeight(v int) AvatarsGetCreditCardOption {
	return func(o *AvatarsGetCreditCardOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Avatars) WithAvatarsGetCreditCardQuality(v int) AvatarsGetCreditCardOption {
	return func(o *AvatarsGetCreditCardOptions) {
		o.Quality = v
		o.enabledSetters["Quality"] = true
	}
}
			
// AvatarsGetCreditCard the credit card endpoint will return you the icon of
// the credit card provider you need. Use width, height and quality arguments
// to change the output settings.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
func (srv *Avatars) AvatarsGetCreditCard(Code string, optionalSetters ...AvatarsGetCreditCardOption)(*interface{}, error) {
	r := strings.NewReplacer("{code}", Code)
	path := r.Replace("/v1/avatars/credit-cards/{code}")
	options := AvatarsGetCreditCardOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["Width"] {
		params["width"] = options.Width
	}
	if options.enabledSetters["Height"] {
		params["height"] = options.Height
	}
	if options.enabledSetters["Quality"] {
		params["quality"] = options.Quality
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
	
// AvatarsGetFavicon use this endpoint to fetch the favorite icon (AKA
// favicon) of any remote website URL.
// 
// This endpoint does not follow HTTP redirects.
func (srv *Avatars) AvatarsGetFavicon(Url string)(*interface{}, error) {
	path := "/v1/avatars/favicon"
	params := map[string]interface{}{}
	params["url"] = Url
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
type AvatarsGetFlagOptions struct {
	Width int
	Height int
	Quality int
	enabledSetters map[string]bool
}
func (options AvatarsGetFlagOptions) New() *AvatarsGetFlagOptions {
	options.enabledSetters = map[string]bool{
		"Width": false,
		"Height": false,
		"Quality": false,
	}
	return &options
}
type AvatarsGetFlagOption func(*AvatarsGetFlagOptions)
func (srv *Avatars) WithAvatarsGetFlagWidth(v int) AvatarsGetFlagOption {
	return func(o *AvatarsGetFlagOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithAvatarsGetFlagHeight(v int) AvatarsGetFlagOption {
	return func(o *AvatarsGetFlagOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Avatars) WithAvatarsGetFlagQuality(v int) AvatarsGetFlagOption {
	return func(o *AvatarsGetFlagOptions) {
		o.Quality = v
		o.enabledSetters["Quality"] = true
	}
}
			
// AvatarsGetFlag you can use this endpoint to show different country flags
// icons to your users. The code argument receives the 2 letter country code.
// Use width, height and quality arguments to change the output settings.
// Country codes follow the [ISO
// 3166-1](https://en.wikipedia.org/wiki/ISO_3166-1) standard.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
func (srv *Avatars) AvatarsGetFlag(Code string, optionalSetters ...AvatarsGetFlagOption)(*interface{}, error) {
	r := strings.NewReplacer("{code}", Code)
	path := r.Replace("/v1/avatars/flags/{code}")
	options := AvatarsGetFlagOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["code"] = Code
	if options.enabledSetters["Width"] {
		params["width"] = options.Width
	}
	if options.enabledSetters["Height"] {
		params["height"] = options.Height
	}
	if options.enabledSetters["Quality"] {
		params["quality"] = options.Quality
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
type AvatarsGetImageOptions struct {
	Width int
	Height int
	enabledSetters map[string]bool
}
func (options AvatarsGetImageOptions) New() *AvatarsGetImageOptions {
	options.enabledSetters = map[string]bool{
		"Width": false,
		"Height": false,
	}
	return &options
}
type AvatarsGetImageOption func(*AvatarsGetImageOptions)
func (srv *Avatars) WithAvatarsGetImageWidth(v int) AvatarsGetImageOption {
	return func(o *AvatarsGetImageOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithAvatarsGetImageHeight(v int) AvatarsGetImageOption {
	return func(o *AvatarsGetImageOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
			
// AvatarsGetImage use this endpoint to fetch a remote image URL and crop it
// to any image size you want. This endpoint is very useful if you need to
// crop and display remote images in your app or in case you want to make sure
// a 3rd party image is properly served using a TLS protocol.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 400x400px.
// 
// This endpoint does not follow HTTP redirects.
func (srv *Avatars) AvatarsGetImage(Url string, optionalSetters ...AvatarsGetImageOption)(*interface{}, error) {
	path := "/v1/avatars/image"
	options := AvatarsGetImageOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["url"] = Url
	if options.enabledSetters["Width"] {
		params["width"] = options.Width
	}
	if options.enabledSetters["Height"] {
		params["height"] = options.Height
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
type AvatarsGetInitialsOptions struct {
	Name string
	Width int
	Height int
	Background string
	enabledSetters map[string]bool
}
func (options AvatarsGetInitialsOptions) New() *AvatarsGetInitialsOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"Width": false,
		"Height": false,
		"Background": false,
	}
	return &options
}
type AvatarsGetInitialsOption func(*AvatarsGetInitialsOptions)
func (srv *Avatars) WithAvatarsGetInitialsName(v string) AvatarsGetInitialsOption {
	return func(o *AvatarsGetInitialsOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Avatars) WithAvatarsGetInitialsWidth(v int) AvatarsGetInitialsOption {
	return func(o *AvatarsGetInitialsOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithAvatarsGetInitialsHeight(v int) AvatarsGetInitialsOption {
	return func(o *AvatarsGetInitialsOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Avatars) WithAvatarsGetInitialsBackground(v string) AvatarsGetInitialsOption {
	return func(o *AvatarsGetInitialsOptions) {
		o.Background = v
		o.enabledSetters["Background"] = true
	}
}
	
// AvatarsGetInitials use this endpoint to show your user initials avatar icon
// on your website or app. By default, this route will try to print your
// logged-in user name or email initials. You can also overwrite the user name
// if you pass the 'name' parameter. If no name is given and no user is
// logged, an empty avatar will be returned.
// 
// You can use the color and background params to change the avatar colors. By
// default, a random theme will be selected. The random theme will persist for
// the user's initials when reloading the same theme will always return for
// the same initials.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
func (srv *Avatars) AvatarsGetInitials(optionalSetters ...AvatarsGetInitialsOption)(*interface{}, error) {
	path := "/v1/avatars/initials"
	options := AvatarsGetInitialsOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Width"] {
		params["width"] = options.Width
	}
	if options.enabledSetters["Height"] {
		params["height"] = options.Height
	}
	if options.enabledSetters["Background"] {
		params["background"] = options.Background
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
type AvatarsGetQROptions struct {
	Size int
	Margin int
	Download bool
	enabledSetters map[string]bool
}
func (options AvatarsGetQROptions) New() *AvatarsGetQROptions {
	options.enabledSetters = map[string]bool{
		"Size": false,
		"Margin": false,
		"Download": false,
	}
	return &options
}
type AvatarsGetQROption func(*AvatarsGetQROptions)
func (srv *Avatars) WithAvatarsGetQRSize(v int) AvatarsGetQROption {
	return func(o *AvatarsGetQROptions) {
		o.Size = v
		o.enabledSetters["Size"] = true
	}
}
func (srv *Avatars) WithAvatarsGetQRMargin(v int) AvatarsGetQROption {
	return func(o *AvatarsGetQROptions) {
		o.Margin = v
		o.enabledSetters["Margin"] = true
	}
}
func (srv *Avatars) WithAvatarsGetQRDownload(v bool) AvatarsGetQROption {
	return func(o *AvatarsGetQROptions) {
		o.Download = v
		o.enabledSetters["Download"] = true
	}
}
			
// AvatarsGetQR converts a given plain text to a QR code image. You can use
// the query parameters to change the size and style of the resulting image.
func (srv *Avatars) AvatarsGetQR(Text string, optionalSetters ...AvatarsGetQROption)(*interface{}, error) {
	path := "/v1/avatars/qr"
	options := AvatarsGetQROptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["text"] = Text
	if options.enabledSetters["Size"] {
		params["size"] = options.Size
	}
	if options.enabledSetters["Margin"] {
		params["margin"] = options.Margin
	}
	if options.enabledSetters["Download"] {
		params["download"] = options.Download
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
type AvatarsGetScreenshotOptions struct {
	Headers interface{}
	ViewportWidth int
	ViewportHeight int
	Scale float64
	Theme string
	UserAgent string
	Fullpage bool
	Locale string
	Timezone string
	Latitude float64
	Longitude float64
	Accuracy float64
	Touch bool
	Permissions []string
	Sleep int
	Width int
	Height int
	Quality int
	Output string
	enabledSetters map[string]bool
}
func (options AvatarsGetScreenshotOptions) New() *AvatarsGetScreenshotOptions {
	options.enabledSetters = map[string]bool{
		"Headers": false,
		"ViewportWidth": false,
		"ViewportHeight": false,
		"Scale": false,
		"Theme": false,
		"UserAgent": false,
		"Fullpage": false,
		"Locale": false,
		"Timezone": false,
		"Latitude": false,
		"Longitude": false,
		"Accuracy": false,
		"Touch": false,
		"Permissions": false,
		"Sleep": false,
		"Width": false,
		"Height": false,
		"Quality": false,
		"Output": false,
	}
	return &options
}
type AvatarsGetScreenshotOption func(*AvatarsGetScreenshotOptions)
func (srv *Avatars) WithAvatarsGetScreenshotHeaders(v interface{}) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Headers = v
		o.enabledSetters["Headers"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotViewportWidth(v int) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.ViewportWidth = v
		o.enabledSetters["ViewportWidth"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotViewportHeight(v int) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.ViewportHeight = v
		o.enabledSetters["ViewportHeight"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotScale(v float64) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Scale = v
		o.enabledSetters["Scale"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotTheme(v string) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Theme = v
		o.enabledSetters["Theme"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotUserAgent(v string) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.UserAgent = v
		o.enabledSetters["UserAgent"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotFullpage(v bool) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Fullpage = v
		o.enabledSetters["Fullpage"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotLocale(v string) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Locale = v
		o.enabledSetters["Locale"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotTimezone(v string) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Timezone = v
		o.enabledSetters["Timezone"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotLatitude(v float64) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Latitude = v
		o.enabledSetters["Latitude"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotLongitude(v float64) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Longitude = v
		o.enabledSetters["Longitude"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotAccuracy(v float64) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Accuracy = v
		o.enabledSetters["Accuracy"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotTouch(v bool) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Touch = v
		o.enabledSetters["Touch"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotPermissions(v []string) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Permissions = v
		o.enabledSetters["Permissions"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotSleep(v int) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Sleep = v
		o.enabledSetters["Sleep"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotWidth(v int) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Width = v
		o.enabledSetters["Width"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotHeight(v int) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Height = v
		o.enabledSetters["Height"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotQuality(v int) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Quality = v
		o.enabledSetters["Quality"] = true
	}
}
func (srv *Avatars) WithAvatarsGetScreenshotOutput(v string) AvatarsGetScreenshotOption {
	return func(o *AvatarsGetScreenshotOptions) {
		o.Output = v
		o.enabledSetters["Output"] = true
	}
}
			
// AvatarsGetScreenshot use this endpoint to capture a screenshot of any
// website URL. This endpoint uses a headless browser to render the webpage
// and capture it as an image.
// 
// You can configure the browser viewport size, theme, user agent,
// geolocation, permissions, and more. Capture either just the viewport or the
// full page scroll.
// 
// When width and height are specified, the image is resized accordingly. If
// both dimensions are 0, the API provides an image at original size. If
// dimensions are not specified, the default viewport size is 1280x720px.
func (srv *Avatars) AvatarsGetScreenshot(Url string, optionalSetters ...AvatarsGetScreenshotOption)(*interface{}, error) {
	path := "/v1/avatars/screenshots"
	options := AvatarsGetScreenshotOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["url"] = Url
	if options.enabledSetters["Headers"] {
		params["headers"] = options.Headers
	}
	if options.enabledSetters["ViewportWidth"] {
		params["viewportWidth"] = options.ViewportWidth
	}
	if options.enabledSetters["ViewportHeight"] {
		params["viewportHeight"] = options.ViewportHeight
	}
	if options.enabledSetters["Scale"] {
		params["scale"] = options.Scale
	}
	if options.enabledSetters["Theme"] {
		params["theme"] = options.Theme
	}
	if options.enabledSetters["UserAgent"] {
		params["userAgent"] = options.UserAgent
	}
	if options.enabledSetters["Fullpage"] {
		params["fullpage"] = options.Fullpage
	}
	if options.enabledSetters["Locale"] {
		params["locale"] = options.Locale
	}
	if options.enabledSetters["Timezone"] {
		params["timezone"] = options.Timezone
	}
	if options.enabledSetters["Latitude"] {
		params["latitude"] = options.Latitude
	}
	if options.enabledSetters["Longitude"] {
		params["longitude"] = options.Longitude
	}
	if options.enabledSetters["Accuracy"] {
		params["accuracy"] = options.Accuracy
	}
	if options.enabledSetters["Touch"] {
		params["touch"] = options.Touch
	}
	if options.enabledSetters["Permissions"] {
		params["permissions"] = options.Permissions
	}
	if options.enabledSetters["Sleep"] {
		params["sleep"] = options.Sleep
	}
	if options.enabledSetters["Width"] {
		params["width"] = options.Width
	}
	if options.enabledSetters["Height"] {
		params["height"] = options.Height
	}
	if options.enabledSetters["Quality"] {
		params["quality"] = options.Quality
	}
	if options.enabledSetters["Output"] {
		params["output"] = options.Output
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
