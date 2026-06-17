package models

import (
    "encoding/json"
    "errors"
)

// File Model
type File struct {
    // File creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // File ID.
    Id string `json:"$id"`
    // File permissions. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    Permissions []string `json:"$permissions"`
    // File update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Bucket ID.
    BucketId string `json:"bucketId"`
    // Total number of chunks available
    ChunksTotal int `json:"chunksTotal"`
    // Total number of chunks uploaded
    ChunksUploaded int `json:"chunksUploaded"`
    // Compression algorithm used for the file. Will be one of none,
    // [gzip](https://en.wikipedia.org/wiki/Gzip), or
    // [zstd](https://en.wikipedia.org/wiki/Zstd).
    Compression string `json:"compression"`
    // Whether file contents are encrypted at rest.
    Encryption bool `json:"encryption"`
    // File mime type.
    MimeType string `json:"mimeType"`
    // File name.
    Name string `json:"name"`
    // File MD5 signature.
    Signature string `json:"signature"`
    // File original size in bytes.
    SizeOriginal int `json:"sizeOriginal"`

    // Used by Decode() method
    data []byte
}

func (model File) New(data []byte) *File {
    model.data = data
    return &model
}

func (model *File) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}