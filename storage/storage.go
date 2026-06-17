package storage

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"strings"
)

// Storage service
type Storage struct {
	client client.Client
}

func New(clt client.Client) *Storage {
	return &Storage{
		client: clt,
	}
}

type AssetIndexOptions struct {
	Search string
	enabledSetters map[string]bool
}
func (options AssetIndexOptions) New() *AssetIndexOptions {
	options.enabledSetters = map[string]bool{
		"Search": false,
	}
	return &options
}
type AssetIndexOption func(*AssetIndexOptions)
func (srv *Storage) WithAssetIndexSearch(v string) AssetIndexOption {
	return func(o *AssetIndexOptions) {
		o.Search = v
		o.enabledSetters["Search"] = true
	}
}
	
// AssetIndex
func (srv *Storage) AssetIndex(optionalSetters ...AssetIndexOption)(*interface{}, error) {
	path := "/v1/storage/assets"
	options := AssetIndexOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["Search"] {
		params["search"] = options.Search
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
type AssetStoreOptions struct {
	AltText string
	Description string
	DisplayName string
	FolderId string
	KeepArchive bool
	Tags []string
	Unpack bool
	Visibility string
	enabledSetters map[string]bool
}
func (options AssetStoreOptions) New() *AssetStoreOptions {
	options.enabledSetters = map[string]bool{
		"AltText": false,
		"Description": false,
		"DisplayName": false,
		"FolderId": false,
		"KeepArchive": false,
		"Tags": false,
		"Unpack": false,
		"Visibility": false,
	}
	return &options
}
type AssetStoreOption func(*AssetStoreOptions)
func (srv *Storage) WithAssetStoreAltText(v string) AssetStoreOption {
	return func(o *AssetStoreOptions) {
		o.AltText = v
		o.enabledSetters["AltText"] = true
	}
}
func (srv *Storage) WithAssetStoreDescription(v string) AssetStoreOption {
	return func(o *AssetStoreOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Storage) WithAssetStoreDisplayName(v string) AssetStoreOption {
	return func(o *AssetStoreOptions) {
		o.DisplayName = v
		o.enabledSetters["DisplayName"] = true
	}
}
func (srv *Storage) WithAssetStoreFolderId(v string) AssetStoreOption {
	return func(o *AssetStoreOptions) {
		o.FolderId = v
		o.enabledSetters["FolderId"] = true
	}
}
func (srv *Storage) WithAssetStoreKeepArchive(v bool) AssetStoreOption {
	return func(o *AssetStoreOptions) {
		o.KeepArchive = v
		o.enabledSetters["KeepArchive"] = true
	}
}
func (srv *Storage) WithAssetStoreTags(v []string) AssetStoreOption {
	return func(o *AssetStoreOptions) {
		o.Tags = v
		o.enabledSetters["Tags"] = true
	}
}
func (srv *Storage) WithAssetStoreUnpack(v bool) AssetStoreOption {
	return func(o *AssetStoreOptions) {
		o.Unpack = v
		o.enabledSetters["Unpack"] = true
	}
}
func (srv *Storage) WithAssetStoreVisibility(v string) AssetStoreOption {
	return func(o *AssetStoreOptions) {
		o.Visibility = v
		o.enabledSetters["Visibility"] = true
	}
}
			
// AssetStore
func (srv *Storage) AssetStore(File string, optionalSetters ...AssetStoreOption)(*interface{}, error) {
	path := "/v1/storage/assets"
	options := AssetStoreOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["file"] = File
	if options.enabledSetters["AltText"] {
		params["alt_text"] = options.AltText
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["DisplayName"] {
		params["display_name"] = options.DisplayName
	}
	if options.enabledSetters["FolderId"] {
		params["folder_id"] = options.FolderId
	}
	if options.enabledSetters["KeepArchive"] {
		params["keep_archive"] = options.KeepArchive
	}
	if options.enabledSetters["Tags"] {
		params["tags"] = options.Tags
	}
	if options.enabledSetters["Unpack"] {
		params["unpack"] = options.Unpack
	}
	if options.enabledSetters["Visibility"] {
		params["visibility"] = options.Visibility
	}
	headers := map[string]interface{}{
		"content-type": "multipart/form-data",
	}


    uploadId := ""

    resp, err := srv.client.FileUpload(path, headers, params, paramName, uploadId)
    if err != nil {
		return nil, err
	}
	var parsed interface{}
	if strings.HasPrefix(resp.Type, "application/json") {
		err = json.Unmarshal([]byte(resp.Result.(string)), &parsed)
		if err != nil {
			return nil, err
		}
		return &parsed, nil
	}
	parsed, ok := resp.Result.(interface{})
	if !ok {
		return nil, errors.New("unexpected response type")
	}
	return &parsed, nil
}
type AssetBulkOptions struct {
	FolderId string
	Visibility string
	enabledSetters map[string]bool
}
func (options AssetBulkOptions) New() *AssetBulkOptions {
	options.enabledSetters = map[string]bool{
		"FolderId": false,
		"Visibility": false,
	}
	return &options
}
type AssetBulkOption func(*AssetBulkOptions)
func (srv *Storage) WithAssetBulkFolderId(v string) AssetBulkOption {
	return func(o *AssetBulkOptions) {
		o.FolderId = v
		o.enabledSetters["FolderId"] = true
	}
}
func (srv *Storage) WithAssetBulkVisibility(v string) AssetBulkOption {
	return func(o *AssetBulkOptions) {
		o.Visibility = v
		o.enabledSetters["Visibility"] = true
	}
}
	
// AssetBulk
func (srv *Storage) AssetBulk(optionalSetters ...AssetBulkOption)(*interface{}, error) {
	path := "/v1/storage/assets/bulk"
	options := AssetBulkOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["FolderId"] {
		params["folder_id"] = options.FolderId
	}
	if options.enabledSetters["Visibility"] {
		params["visibility"] = options.Visibility
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
	
// AssetDestroy
func (srv *Storage) AssetDestroy(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/assets/{id}")
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
	
// AssetShow
func (srv *Storage) AssetShow(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/assets/{id}")
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
type AssetUpdateOptions struct {
	AltText string
	Description string
	DisplayName string
	FolderId string
	Name string
	Tags []string
	Visibility string
	enabledSetters map[string]bool
}
func (options AssetUpdateOptions) New() *AssetUpdateOptions {
	options.enabledSetters = map[string]bool{
		"AltText": false,
		"Description": false,
		"DisplayName": false,
		"FolderId": false,
		"Name": false,
		"Tags": false,
		"Visibility": false,
	}
	return &options
}
type AssetUpdateOption func(*AssetUpdateOptions)
func (srv *Storage) WithAssetUpdateAltText(v string) AssetUpdateOption {
	return func(o *AssetUpdateOptions) {
		o.AltText = v
		o.enabledSetters["AltText"] = true
	}
}
func (srv *Storage) WithAssetUpdateDescription(v string) AssetUpdateOption {
	return func(o *AssetUpdateOptions) {
		o.Description = v
		o.enabledSetters["Description"] = true
	}
}
func (srv *Storage) WithAssetUpdateDisplayName(v string) AssetUpdateOption {
	return func(o *AssetUpdateOptions) {
		o.DisplayName = v
		o.enabledSetters["DisplayName"] = true
	}
}
func (srv *Storage) WithAssetUpdateFolderId(v string) AssetUpdateOption {
	return func(o *AssetUpdateOptions) {
		o.FolderId = v
		o.enabledSetters["FolderId"] = true
	}
}
func (srv *Storage) WithAssetUpdateName(v string) AssetUpdateOption {
	return func(o *AssetUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Storage) WithAssetUpdateTags(v []string) AssetUpdateOption {
	return func(o *AssetUpdateOptions) {
		o.Tags = v
		o.enabledSetters["Tags"] = true
	}
}
func (srv *Storage) WithAssetUpdateVisibility(v string) AssetUpdateOption {
	return func(o *AssetUpdateOptions) {
		o.Visibility = v
		o.enabledSetters["Visibility"] = true
	}
}
			
// AssetUpdate
func (srv *Storage) AssetUpdate(Id string, optionalSetters ...AssetUpdateOption)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/assets/{id}")
	options := AssetUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["AltText"] {
		params["alt_text"] = options.AltText
	}
	if options.enabledSetters["Description"] {
		params["description"] = options.Description
	}
	if options.enabledSetters["DisplayName"] {
		params["display_name"] = options.DisplayName
	}
	if options.enabledSetters["FolderId"] {
		params["folder_id"] = options.FolderId
	}
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["Tags"] {
		params["tags"] = options.Tags
	}
	if options.enabledSetters["Visibility"] {
		params["visibility"] = options.Visibility
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
	
// AssetDownload
func (srv *Storage) AssetDownload(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/assets/{id}/download")
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
	
// AssetPermanent
func (srv *Storage) AssetPermanent(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/assets/{id}/permanent")
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
	
// AssetReprocess
func (srv *Storage) AssetReprocess(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/assets/{id}/reprocess")
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
	
// AssetRestore
func (srv *Storage) AssetRestore(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/assets/{id}/restore")
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
type AssetSignOptions struct {
	TtlSeconds int
	enabledSetters map[string]bool
}
func (options AssetSignOptions) New() *AssetSignOptions {
	options.enabledSetters = map[string]bool{
		"TtlSeconds": false,
	}
	return &options
}
type AssetSignOption func(*AssetSignOptions)
func (srv *Storage) WithAssetSignTtlSeconds(v int) AssetSignOption {
	return func(o *AssetSignOptions) {
		o.TtlSeconds = v
		o.enabledSetters["TtlSeconds"] = true
	}
}
			
// AssetSign
func (srv *Storage) AssetSign(Id string, optionalSetters ...AssetSignOption)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/assets/{id}/sign")
	options := AssetSignOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["TtlSeconds"] {
		params["ttl_seconds"] = options.TtlSeconds
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
type AssetUnpackOptions struct {
	KeepArchive bool
	TargetFolderId string
	enabledSetters map[string]bool
}
func (options AssetUnpackOptions) New() *AssetUnpackOptions {
	options.enabledSetters = map[string]bool{
		"KeepArchive": false,
		"TargetFolderId": false,
	}
	return &options
}
type AssetUnpackOption func(*AssetUnpackOptions)
func (srv *Storage) WithAssetUnpackKeepArchive(v bool) AssetUnpackOption {
	return func(o *AssetUnpackOptions) {
		o.KeepArchive = v
		o.enabledSetters["KeepArchive"] = true
	}
}
func (srv *Storage) WithAssetUnpackTargetFolderId(v string) AssetUnpackOption {
	return func(o *AssetUnpackOptions) {
		o.TargetFolderId = v
		o.enabledSetters["TargetFolderId"] = true
	}
}
			
// AssetUnpack
func (srv *Storage) AssetUnpack(Id string, optionalSetters ...AssetUnpackOption)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/assets/{id}/unpack")
	options := AssetUnpackOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["KeepArchive"] {
		params["keep_archive"] = options.KeepArchive
	}
	if options.enabledSetters["TargetFolderId"] {
		params["target_folder_id"] = options.TargetFolderId
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

// FolderIndex
func (srv *Storage) FolderIndex()(*interface{}, error) {
	path := "/v1/storage/folders"
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
type FolderStoreOptions struct {
	ParentId string
	enabledSetters map[string]bool
}
func (options FolderStoreOptions) New() *FolderStoreOptions {
	options.enabledSetters = map[string]bool{
		"ParentId": false,
	}
	return &options
}
type FolderStoreOption func(*FolderStoreOptions)
func (srv *Storage) WithFolderStoreParentId(v string) FolderStoreOption {
	return func(o *FolderStoreOptions) {
		o.ParentId = v
		o.enabledSetters["ParentId"] = true
	}
}
			
// FolderStore
func (srv *Storage) FolderStore(Name string, optionalSetters ...FolderStoreOption)(*interface{}, error) {
	path := "/v1/storage/folders"
	options := FolderStoreOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["name"] = Name
	if options.enabledSetters["ParentId"] {
		params["parent_id"] = options.ParentId
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
type FolderDestroyOptions struct {
	Recursive bool
	enabledSetters map[string]bool
}
func (options FolderDestroyOptions) New() *FolderDestroyOptions {
	options.enabledSetters = map[string]bool{
		"Recursive": false,
	}
	return &options
}
type FolderDestroyOption func(*FolderDestroyOptions)
func (srv *Storage) WithFolderDestroyRecursive(v bool) FolderDestroyOption {
	return func(o *FolderDestroyOptions) {
		o.Recursive = v
		o.enabledSetters["Recursive"] = true
	}
}
			
// FolderDestroy
func (srv *Storage) FolderDestroy(Id string, optionalSetters ...FolderDestroyOption)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/folders/{id}")
	options := FolderDestroyOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Recursive"] {
		params["recursive"] = options.Recursive
	}
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
	
// FolderShow
func (srv *Storage) FolderShow(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/folders/{id}")
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
type FolderUpdateOptions struct {
	Name string
	ParentId string
	enabledSetters map[string]bool
}
func (options FolderUpdateOptions) New() *FolderUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Name": false,
		"ParentId": false,
	}
	return &options
}
type FolderUpdateOption func(*FolderUpdateOptions)
func (srv *Storage) WithFolderUpdateName(v string) FolderUpdateOption {
	return func(o *FolderUpdateOptions) {
		o.Name = v
		o.enabledSetters["Name"] = true
	}
}
func (srv *Storage) WithFolderUpdateParentId(v string) FolderUpdateOption {
	return func(o *FolderUpdateOptions) {
		o.ParentId = v
		o.enabledSetters["ParentId"] = true
	}
}
			
// FolderUpdate
func (srv *Storage) FolderUpdate(Id string, optionalSetters ...FolderUpdateOption)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/folders/{id}")
	options := FolderUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Name"] {
		params["name"] = options.Name
	}
	if options.enabledSetters["ParentId"] {
		params["parent_id"] = options.ParentId
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

// SyncRuleIndex
func (srv *Storage) SyncRuleIndex()(*interface{}, error) {
	path := "/v1/storage/sftp/rules"
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

// SyncRuleStore
func (srv *Storage) SyncRuleStore()(*interface{}, error) {
	path := "/v1/storage/sftp/rules"
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
	
// SyncRuleDestroy
func (srv *Storage) SyncRuleDestroy(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/sftp/rules/{id}")
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
	
// SyncRuleShow
func (srv *Storage) SyncRuleShow(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/sftp/rules/{id}")
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
	
// SyncRuleUpdate
func (srv *Storage) SyncRuleUpdate(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/sftp/rules/{id}")
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
	
// SyncRuleRun
func (srv *Storage) SyncRuleRun(Id string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/sftp/rules/{id}/run")
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
			
// SyncRuleRunProtocol
func (srv *Storage) SyncRuleRunProtocol(Id string, RunId string)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id, "{runId}", RunId)
	path := r.Replace("/v1/storage/sftp/rules/{id}/runs/{runId}")
	params := map[string]interface{}{}
	params["id"] = Id
	params["runId"] = RunId
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
type SyncRuleHistoryOptions struct {
	RuleId string
	From string
	To string
	enabledSetters map[string]bool
}
func (options SyncRuleHistoryOptions) New() *SyncRuleHistoryOptions {
	options.enabledSetters = map[string]bool{
		"RuleId": false,
		"From": false,
		"To": false,
	}
	return &options
}
type SyncRuleHistoryOption func(*SyncRuleHistoryOptions)
func (srv *Storage) WithSyncRuleHistoryRuleId(v string) SyncRuleHistoryOption {
	return func(o *SyncRuleHistoryOptions) {
		o.RuleId = v
		o.enabledSetters["RuleId"] = true
	}
}
func (srv *Storage) WithSyncRuleHistoryFrom(v string) SyncRuleHistoryOption {
	return func(o *SyncRuleHistoryOptions) {
		o.From = v
		o.enabledSetters["From"] = true
	}
}
func (srv *Storage) WithSyncRuleHistoryTo(v string) SyncRuleHistoryOption {
	return func(o *SyncRuleHistoryOptions) {
		o.To = v
		o.enabledSetters["To"] = true
	}
}
	
// SyncRuleHistory
func (srv *Storage) SyncRuleHistory(optionalSetters ...SyncRuleHistoryOption)(*interface{}, error) {
	path := "/v1/storage/sftp/sync-history"
	options := SyncRuleHistoryOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	if options.enabledSetters["RuleId"] {
		params["rule_id"] = options.RuleId
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

// TenantStats
func (srv *Storage) TenantStats()(*interface{}, error) {
	path := "/v1/storage/tenant/stats"
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

// TenantUsage
func (srv *Storage) TenantUsage()(*interface{}, error) {
	path := "/v1/storage/tenant/usage"
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
