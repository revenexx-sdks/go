package storage

import (
	"encoding/json"
	"errors"
	"github.com/revenexx-sdks/go/client"
	"github.com/revenexx-sdks/go/file"
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
	
// AssetIndex list the media assets in this tenant, newest first. Narrow the
// list with
// `filter[folder_id]`, `filter[kind]`, `filter[status]` and a
// `filter[created_at][gte]`/`[lte]` range; search original names, display
// names, alt text and descriptions with `search`; order by `created_at`,
// `size_bytes` or `original_name` (prefix with `-` to reverse). One page is
// returned, 50 records by default and 200 at most.
// 
// Records only: no file content is returned — fetch bytes with
// `GET /assets/{id}/download` or hand out a link with
// `POST /assets/{id}/sign`. Deleted assets are not listed.
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
			
// AssetStore upload one file into this tenant's media library. The file is
// checked
// against the tenant's single-file limit and its remaining storage quota,
// its media type is sniffed from the content rather than trusted from the
// request, and it is virus-scanned before anything is written. The stored
// asset comes back with status `pending_processing`; metadata extraction
// finishes asynchronously and moves it to `available`. `folder_id`,
// `visibility`, `alt_text`, `description`, `display_name` and `tags` are
// applied on the way in; set `unpack` to also queue an uploaded archive's
// members for ingestion.
// 
// Every call creates a new asset — this never replaces the content of an
// existing one — and it takes exactly one file. Use `POST /assets/bulk` for
// several.
func (srv *Storage) AssetStore(File file.InputFile, optionalSetters ...AssetStoreOption)(*interface{}, error) {
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

    paramName := "file"


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
	
// AssetBulk upload a batch of files in one request under `files`, each
// ingested
// exactly as `POST /assets` ingests a single file. The batch is rejected as
// a whole when it carries no files, more files than one request may carry,
// or too many bytes in total. Past that point every file is attempted
// independently and the call answers 207 with a `results` entry per file:
// either the created asset or the error that rejected it. A partial failure
// is therefore a successful call, not an error status — read `results`.
// 
// Only `folder_id` and `visibility` apply, and they apply to the whole
// batch; per-file metadata is not accepted here. Set it afterwards with
// `PATCH /assets/{id}`.
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
	
// AssetDestroy soft-delete an asset: it stops being listed and served, its
// status
// becomes `soft_deleted`, and it is scheduled for permanent deletion once
// the retention window has passed. Until then `POST /assets/{id}/restore`
// brings it back.
// 
// The stored file is not erased at this point and its bytes still count
// against the tenant's storage quota — use `DELETE /assets/{id}/permanent`
// to erase it and free the quota immediately.
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
	
// AssetShow fetch one asset's record by id: name, folder, media type, size,
// status,
// tags, the extracted metadata and the delivery URL (null for a private
// asset, which is reachable only through a signed URL). Metadata only — the
// bytes are served by `GET /assets/{id}/download`. A deleted asset is not
// visible here until `POST /assets/{id}/restore` brings it back.
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
			
// AssetUpdate change an asset's metadata: `display_name`, `alt_text`,
// `description`,
// `visibility` and `tags`. Sending `folder_id` moves it and sending `name`
// renames it; either re-derives the asset's public delivery path, so links
// built from the old path stop resolving. Only the fields present in the
// request are touched.
// 
// The stored file itself is never modified here — to change the content,
// upload a new asset.
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
	
// AssetDownload stream the asset's original file back as an attachment, named
// after the
// asset. This is the authenticated read path — every call carries the
// caller's credentials — and the bytes are the ones that were uploaded: no
// resizing, re-encoding or other transformation is applied.
// 
// To let a browser, an email or a third party fetch the file without an API
// credential, mint a link with `POST /assets/{id}/sign` instead.
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
	
// AssetPermanent erase an asset and its stored file for good and credit its
// bytes back to
// the tenant's used storage. Works on live and soft-deleted assets alike.
// 
// This cannot be undone: there is no restore afterwards, and links to the
// asset stop resolving at once. Use `DELETE /assets/{id}` for the
// reversible variant. Requires the elevated (admin) tier.
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
	
// AssetReprocess re-run post-upload processing for one asset. It returns to
// `pending_processing` and the job re-extracts its metadata — and, for a 3D
// model, re-renders the preview and mesh derivatives — before marking it
// `available` again. The usual reason is an asset stuck in
// `processing_failed`.
// 
// The stored file is neither re-uploaded nor altered, and no thumbnails are
// produced: delivery transforms are applied on the fly when the asset is
// served, not here.
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
	
// AssetRestore bring a soft-deleted asset back: the scheduled permanent
// deletion is
// cleared and the asset returns to `available`, listed and served again
// under its original path. Only works while the asset is still inside its
// retention window — once it has been erased, by
// `DELETE /assets/{id}/permanent` or by the retention sweep, there is
// nothing left to restore.
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
			
// AssetSign mint a time-limited URL that serves this asset without an API
// credential
// — the way to hand a private asset to a browser, an email or a third
// party. `ttl_seconds` sets the lifetime: one hour by default, seven days
// at most. The response carries the URL and the lifetime it was issued
// with.
// 
// The signature is checked at the delivery edge. A link cannot be revoked
// before it expires, so keep the lifetime short. A public asset already
// carries an unsigned delivery URL on its record and does not need this.
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
			
// AssetUnpack ingest the members of an already-uploaded archive as individual
// assets.
// They land in a folder named after the archive, created under
// `target_folder_id` or, when that is omitted, under the archive's own
// folder, and the archive's internal directory structure is mirrored
// beneath it. Each member goes through the same pipeline as an upload —
// media-type sniff, virus scan, quota — and a member that fails is skipped
// rather than failing the run. `keep_archive` (true by default) decides
// whether the archive asset itself survives.
// 
// Asynchronous: this answers 202 as soon as the work is queued, so poll the
// folder or asset list for the results. Only an asset that is an archive of
// a supported type can be unpacked; an upload can ask for the same thing
// inline with `unpack`.
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

// FolderIndex return every folder in this tenant as one flat list ordered by
// path, each
// record carrying its `parent_id` and its materialized `path`, so a client
// can rebuild the tree without walking it. Not paginated and not filtered.
// 
// Folders hold no file content of their own — list a folder's assets with
// `GET /assets` and `filter[folder_id]`.
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
			
// FolderStore create a folder under `parent_id`, or at the library root when
// it is
// omitted. The `name` is slugged into a path segment and appended to the
// parent's path; that path is what the public delivery URL of every asset
// inside it is built from, so two siblings may not slug to the same
// segment.
// 
// Creating a folder moves nothing into it — assign assets with
// `folder_id` on upload or with `PATCH /assets/{id}`.
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
			
// FolderDestroy delete a folder. By default it has to be empty: a folder that
// still holds
// folders or assets is refused, so pass `recursive=true` to delete it
// together with everything beneath it.
// 
// A recursive delete soft-deletes the assets it takes with it — their files
// are not erased and their bytes still count against the tenant's storage
// quota, and each remains restorable through `POST /assets/{id}/restore`.
// System folders cannot be deleted.
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
	
// FolderShow fetch one folder's record by id: its name, its parent, the
// materialized
// path assets inside it are delivered under, and whether it is a system
// folder (system folders cannot be renamed, moved or deleted).
// 
// Its contents are not included — list them with `GET /assets` and
// `filter[folder_id]`, and its child folders with `GET /folders`.
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
			
// FolderUpdate rename a folder with `name`, move it under a different parent
// with
// `parent_id` (null for the root), or both at once. Either rewrites the
// folder's materialized path and the path of every folder beneath it, which
// changes the public delivery URL of every asset they hold — existing links
// built from the old path stop resolving.
// 
// Nothing else about the assets changes; they are not moved, re-uploaded or
// reprocessed. A system folder cannot be changed, a folder cannot be moved
// inside its own subtree, and the new name has to slug to a segment free
// among its new siblings.
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

// SyncRuleIndex return this tenant's SFTP sync rules, newest first, each with
// the account
// and remote path it pulls from, the folder it imports into, its cron
// schedule, whether it is enabled and when it last ran. Not paginated and
// not filtered.
// 
// These are the rules themselves, not what they moved: for the files a rule
// has actually transferred, see `GET /sftp/sync-history`.
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
type SyncRuleStoreOptions struct {
	Enabled bool
	Options []string
	Schedule string
	TargetFolderId string
	enabledSetters map[string]bool
}
func (options SyncRuleStoreOptions) New() *SyncRuleStoreOptions {
	options.enabledSetters = map[string]bool{
		"Enabled": false,
		"Options": false,
		"Schedule": false,
		"TargetFolderId": false,
	}
	return &options
}
type SyncRuleStoreOption func(*SyncRuleStoreOptions)
func (srv *Storage) WithSyncRuleStoreEnabled(v bool) SyncRuleStoreOption {
	return func(o *SyncRuleStoreOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Storage) WithSyncRuleStoreOptions(v []string) SyncRuleStoreOption {
	return func(o *SyncRuleStoreOptions) {
		o.Options = v
		o.enabledSetters["Options"] = true
	}
}
func (srv *Storage) WithSyncRuleStoreSchedule(v string) SyncRuleStoreOption {
	return func(o *SyncRuleStoreOptions) {
		o.Schedule = v
		o.enabledSetters["Schedule"] = true
	}
}
func (srv *Storage) WithSyncRuleStoreTargetFolderId(v string) SyncRuleStoreOption {
	return func(o *SyncRuleStoreOptions) {
		o.TargetFolderId = v
		o.enabledSetters["TargetFolderId"] = true
	}
}
					
// SyncRuleStore schedule a recurring one-way pull from a directory on the
// tenant's SFTP
// storage box into this media library. `sftp_account_id` selects the
// account, `source_path` the remote directory, `target_folder_id` the
// folder imported assets land in, and `schedule` a cron expression (every
// five minutes when omitted) at which the rule falls due. `options` carries
// the per-rule knobs: recursion, include/exclude and size filters, how long
// a remote file has to have stopped changing before it is taken, and
// whether it is deleted from the remote after a successful transfer.
// 
// Each run ingests every matching remote file exactly as an upload would,
// quota, media-type and virus checks included, and records one history
// entry per file. Creating the rule transfers nothing: the first run
// happens when the schedule next falls due, or immediately if you call
// `POST /sftp/rules/{id}/run`. Nothing is ever pushed back to the remote,
// beyond the optional delete after a successful transfer. Requires the
// elevated (admin) tier.
func (srv *Storage) SyncRuleStore(SftpAccountId string, SourcePath string, optionalSetters ...SyncRuleStoreOption)(*interface{}, error) {
	path := "/v1/storage/sftp/rules"
	options := SyncRuleStoreOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["sftp_account_id"] = SftpAccountId
	params["source_path"] = SourcePath
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Options"] {
		params["options"] = options.Options
	}
	if options.enabledSetters["Schedule"] {
		params["schedule"] = options.Schedule
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
	
// SyncRuleDestroy delete a sync rule so it is never scheduled again. The
// assets it already
// imported stay exactly where they are, its recorded run history is kept,
// and nothing on the remote is touched.
// 
// To stop a rule only for a while, set `enabled` to false with
// `PATCH /sftp/rules/{id}` instead — a deleted rule cannot be restored.
// Requires the elevated (admin) tier.
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
	
// SyncRuleShow fetch one sync rule's configuration by id: the account and
// remote path it
// pulls from, its target folder, its cron schedule, its `options` and
// `last_run_at`.
// 
// Configuration only, and `last_run_at` says when a run was last attempted,
// not whether it succeeded. What a run did is in
// `GET /sftp/rules/{id}/runs/{runId}` and `GET /sftp/sync-history`.
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
type SyncRuleUpdateOptions struct {
	Enabled bool
	Options []string
	Schedule string
	SftpAccountId string
	SourcePath string
	TargetFolderId string
	enabledSetters map[string]bool
}
func (options SyncRuleUpdateOptions) New() *SyncRuleUpdateOptions {
	options.enabledSetters = map[string]bool{
		"Enabled": false,
		"Options": false,
		"Schedule": false,
		"SftpAccountId": false,
		"SourcePath": false,
		"TargetFolderId": false,
	}
	return &options
}
type SyncRuleUpdateOption func(*SyncRuleUpdateOptions)
func (srv *Storage) WithSyncRuleUpdateEnabled(v bool) SyncRuleUpdateOption {
	return func(o *SyncRuleUpdateOptions) {
		o.Enabled = v
		o.enabledSetters["Enabled"] = true
	}
}
func (srv *Storage) WithSyncRuleUpdateOptions(v []string) SyncRuleUpdateOption {
	return func(o *SyncRuleUpdateOptions) {
		o.Options = v
		o.enabledSetters["Options"] = true
	}
}
func (srv *Storage) WithSyncRuleUpdateSchedule(v string) SyncRuleUpdateOption {
	return func(o *SyncRuleUpdateOptions) {
		o.Schedule = v
		o.enabledSetters["Schedule"] = true
	}
}
func (srv *Storage) WithSyncRuleUpdateSftpAccountId(v string) SyncRuleUpdateOption {
	return func(o *SyncRuleUpdateOptions) {
		o.SftpAccountId = v
		o.enabledSetters["SftpAccountId"] = true
	}
}
func (srv *Storage) WithSyncRuleUpdateSourcePath(v string) SyncRuleUpdateOption {
	return func(o *SyncRuleUpdateOptions) {
		o.SourcePath = v
		o.enabledSetters["SourcePath"] = true
	}
}
func (srv *Storage) WithSyncRuleUpdateTargetFolderId(v string) SyncRuleUpdateOption {
	return func(o *SyncRuleUpdateOptions) {
		o.TargetFolderId = v
		o.enabledSetters["TargetFolderId"] = true
	}
}
			
// SyncRuleUpdate change a sync rule in place: its account, remote path,
// target folder,
// schedule or options, or `enabled` to pause and resume it without deleting
// it. Only the fields present in the request are touched, but `options` is
// replaced wholesale rather than merged — send the whole object.
// 
// A change takes effect from the next run; a run already in flight is not
// affected, and nothing a previous run imported is revisited or undone.
// Requires the elevated (admin) tier.
func (srv *Storage) SyncRuleUpdate(Id string, optionalSetters ...SyncRuleUpdateOption)(*interface{}, error) {
	r := strings.NewReplacer("{id}", Id)
	path := r.Replace("/v1/storage/sftp/rules/{id}")
	options := SyncRuleUpdateOptions{}.New()
	for _, opt := range optionalSetters {
		opt(options)
	}
	params := map[string]interface{}{}
	params["id"] = Id
	if options.enabledSetters["Enabled"] {
		params["enabled"] = options.Enabled
	}
	if options.enabledSetters["Options"] {
		params["options"] = options.Options
	}
	if options.enabledSetters["Schedule"] {
		params["schedule"] = options.Schedule
	}
	if options.enabledSetters["SftpAccountId"] {
		params["sftp_account_id"] = options.SftpAccountId
	}
	if options.enabledSetters["SourcePath"] {
		params["source_path"] = options.SourcePath
	}
	if options.enabledSetters["TargetFolderId"] {
		params["target_folder_id"] = options.TargetFolderId
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
	
// SyncRuleRun queue a run of this rule straight away, outside its schedule.
// Answers 202
// with the rule id as soon as the job is queued — it does not wait for the
// transfer and it does not hand back a run id, so follow the outcome in
// `GET /sftp/sync-history`.
// 
// The rule's own schedule is untouched, and this does not enable a disabled
// rule: the job is queued but does nothing when it picks a disabled rule
// up. Requires the elevated (admin) tier.
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
			
// SyncRuleRunProtocol return the per-file protocol of one run of one sync
// rule: every entry the
// run recorded, oldest first, with the remote source path, the asset it
// produced, the bytes transferred, the duration and the error where one
// applies — plus a `summary` counting those entries by status (`success`,
// `skipped`, `failed`, `quarantined`).
// 
// Use it to find out what one run actually did. It is not paginated, and it
// does not list a rule's runs: take the `run_id` from
// `GET /sftp/sync-history`. An unknown `runId` under a rule that does exist
// is an empty protocol, not a 404.
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
	
// SyncRuleHistory page through this tenant's per-file sync records across
// every rule,
// newest first. Each entry names the run it belongs to, the rule, the
// remote source path, the asset it produced where there is one, the
// outcome — `success`, `skipped`, `failed` or `quarantined` — the bytes
// transferred and how long it took. Narrow it with `rule_id` and a
// `from`/`to` range on when the entry was recorded; one page is returned,
// 50 entries by default and 200 at most.
// 
// This is the audit trail of what SFTP sync has brought in: every file
// taken, skipped and rejected leaves an entry, and a run that matched
// nothing leaves one too. To read a single run whole instead, group by
// `run_id` and call `GET /sftp/rules/{id}/runs/{runId}`.
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

// TenantStats break this tenant's library down by asset kind — `image`,
// `video`,
// `audio`, `pdf`, `document`, `archive`, `model3d`, `other` — with a count
// and a byte total for each kind that has at least one asset, alongside the
// tenant-wide totals.
// 
// A dashboard figure, not a listing: no asset is named, and nothing here
// can be filtered. The tenant-wide byte total is the same running figure
// `GET /tenant/usage` reports, so soft-deleted assets are counted in it.
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

// TenantUsage report this tenant's storage consumption: the bytes in use, the
// byte
// quota in force (null when the tenant is uncapped) and how many assets it
// holds. This is the figure the quota check on upload compares against — it
// is maintained as a running total on every upload and permanent delete
// rather than summed on read.
// 
// Soft-deleted assets are still counted, because their files are still
// stored; their bytes come back only once they are permanently deleted. For
// the breakdown by asset kind, see `GET /tenant/stats`.
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
