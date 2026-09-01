package models

import (
    "encoding/json"
    "errors"
)

// Bucket Bucket
type Bucket struct {
    // Bucket creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Bucket ID.
    Id string `json:"$id"`
    // Bucket permissions. Each entry is a permission string: an action wrapping a
    // role, e.g. `read("any")`, `update("user:abc")`, `delete("team:abc/owner")`.
    // Actions are `read`, `create`, `update`, `delete` and the aggregate `write`
    // (= create + update + delete); the role inside the quotes takes the form
    // described under “Role strings” in this document's introduction.
    Permissions []string `json:"$permissions"`
    // Bucket update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Allowed file extensions.
    AllowedFileExtensions []string `json:"allowedFileExtensions"`
    // Virus scanning is enabled.
    Antivirus bool `json:"antivirus"`
    // Compression algorithm chosen for compression. Will be one of none,
    // [gzip](https://en.wikipedia.org/wiki/Gzip), or
    // [zstd](https://en.wikipedia.org/wiki/Zstd).
    Compression string `json:"compression"`
    // Bucket enabled.
    Enabled bool `json:"enabled"`
    // Bucket is encrypted.
    Encryption bool `json:"encryption"`
    // Whether file-level security is enabled. When it is, each record's own
    // `$permissions` are enforced on top of the container's.
    FileSecurity bool `json:"fileSecurity"`
    // Maximum file size supported.
    MaximumFileSize int `json:"maximumFileSize"`
    // Bucket name.
    Name string `json:"name"`
    // Total size of this bucket in bytes.
    TotalSize int `json:"totalSize"`
    // Image transformations are enabled.
    Transformations bool `json:"transformations"`

    // Used by Decode() method
    data []byte
}

func (model Bucket) New(data []byte) *Bucket {
    model.data = data
    return &model
}

func (model *Bucket) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}