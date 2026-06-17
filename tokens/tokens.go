package tokens

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/models"
	"strings"
)

// Tokens service
type Tokens struct {
	client client.Client
}

func New(clt client.Client) *Tokens {
	return &Tokens{
		client: clt,
	}
}

type TokensListOptions struct {
	Queries []string
	Total bool
	enabledSetters map[string]bool
}
func (options TokensListOptions) New() *TokensListOptions {
	options.enabledSetters = map[string]bool{
		"Queries": false,
		"Total": false,
	}
	return &options
}
type TokensListOption func(*TokensListOptions)
func (srv *Tokens) WithTokensListQueries(v []string) TokensListOption {
	return func(o *TokensListOptions) {
		o.Queries = v
		o.enabledSetters["Queries"] = true
	}
}
func (srv *Tokens) WithTokensListTotal(v bool) TokensListOption {
	return func(o *TokensListOptions) {
		o.Total = v
		o.enabledSetters["Total"] = true
	}
}
					
// TokensList list all the tokens created for a specific file or bucket. You
// can use the query params to filter your results.
func (srv *Tokens) TokensList(BucketId string, FileId string, optionalSetters ...TokensListOption)(*models.ResourceTokenList, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/v1/tokens/buckets/{bucketId}/files/{fileId}")
	options := TokensListOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
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

		parsed := models.ResourceTokenList{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ResourceTokenList
	parsed, ok := resp.Result.(models.ResourceTokenList)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type TokensCreateFileTokenOptions struct {
	Expire string
	enabledSetters map[string]bool
}
func (options TokensCreateFileTokenOptions) New() *TokensCreateFileTokenOptions {
	options.enabledSetters = map[string]bool{
		"Expire": false,
	}
	return &options
}
type TokensCreateFileTokenOption func(*TokensCreateFileTokenOptions)
func (srv *Tokens) WithTokensCreateFileTokenExpire(v string) TokensCreateFileTokenOption {
	return func(o *TokensCreateFileTokenOptions) {
		o.Expire = v
		o.enabledSetters["Expire"] = true
	}
}
					
// TokensCreateFileToken create a new token. A token is linked to a file.
// Token can be passed as a request URL search parameter.
func (srv *Tokens) TokensCreateFileToken(BucketId string, FileId string, optionalSetters ...TokensCreateFileTokenOption)(*models.ResourceToken, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	path := r.Replace("/v1/tokens/buckets/{bucketId}/files/{fileId}")
	options := TokensCreateFileTokenOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["bucketId"] = BucketId
	params["fileId"] = FileId
	if options.enabledSetters["Expire"] {
		params["expire"] = options.Expire
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

		parsed := models.ResourceToken{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ResourceToken
	parsed, ok := resp.Result.(models.ResourceToken)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
	
// TokensDelete delete a token by its unique ID.
func (srv *Tokens) TokensDelete(TokenId string)(*interface{}, error) {
	r := strings.NewReplacer("{tokenId}", TokenId)
	path := r.Replace("/v1/tokens/{tokenId}")
	params := map[string]interface{}{}
	params["tokenId"] = TokenId
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
	
// TokensGet get a token by its unique ID.
func (srv *Tokens) TokensGet(TokenId string)(*models.ResourceToken, error) {
	r := strings.NewReplacer("{tokenId}", TokenId)
	path := r.Replace("/v1/tokens/{tokenId}")
	params := map[string]interface{}{}
	params["tokenId"] = TokenId
	headers := map[string]interface{}{
	}

	resp, err := srv.client.Call("GET", path, headers, params)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(resp.Type, "application/json") {
		bytes := []byte(resp.Result.(string))

		parsed := models.ResourceToken{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ResourceToken
	parsed, ok := resp.Result.(models.ResourceToken)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
type TokensUpdateOptions struct {
	Expire string
	enabledSetters map[string]bool
}
func (options TokensUpdateOptions) New() *TokensUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Expire": false,
	}
	return &options
}
type TokensUpdateOption func(*TokensUpdateOptions)
func (srv *Tokens) WithTokensUpdateExpire(v string) TokensUpdateOption {
	return func(o *TokensUpdateOptions) {
		o.Expire = v
		o.enabledSetters["Expire"] = true
	}
}
			
// TokensUpdate update a token by its unique ID. Use this endpoint to update a
// token's expiry date.
func (srv *Tokens) TokensUpdate(TokenId string, optionalSetters ...TokensUpdateOption)(*models.ResourceToken, error) {
	r := strings.NewReplacer("{tokenId}", TokenId)
	path := r.Replace("/v1/tokens/{tokenId}")
	options := TokensUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["tokenId"] = TokenId
	if options.enabledSetters["Expire"] {
		params["expire"] = options.Expire
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

		parsed := models.ResourceToken{}.New(bytes)

		err = json.Unmarshal(bytes, parsed)
		if err != nil {
			return nil, err
		}

		return parsed, nil
	}
	var parsed models.ResourceToken
	parsed, ok := resp.Result.(models.ResourceToken)
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil

}
