package models

import (
    "encoding/json"
    "errors"
)

// Model
type AssetResource struct {
    // 
    AltText string `json:"alt_text"`
    // 
    ContentHash string `json:"content_hash"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    DeletedAt string `json:"deleted_at"`
    // 
    Description string `json:"description"`
    // 
    DisplayName string `json:"display_name"`
    // 
    DominantColor string `json:"dominant_color"`
    // 
    DurationMs int `json:"duration_ms"`
    // 
    FolderId string `json:"folder_id"`
    // 
    Height int `json:"height"`
    // 
    Id string `json:"id"`
    // 
    Kind string `json:"kind"`
    // 
    Metadata []interface{} `json:"metadata"`
    // 
    MimeType string `json:"mime_type"`
    // 
    OriginalName string `json:"original_name"`
    // 
    PageCount int `json:"page_count"`
    // 
    PathName string `json:"path_name"`
    // 
    ProcessedAt string `json:"processed_at"`
    // 
    SizeBytes int `json:"size_bytes"`
    // 
    Status string `json:"status"`
    // 
    Tags []interface{} `json:"tags"`
    // 
    TenantId string `json:"tenant_id"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    Url string `json:"url"`
    // 
    Visibility string `json:"visibility"`
    // 
    Width int `json:"width"`

    // Used by Decode() method
    data []byte
}

func (model AssetResource) New(data []byte) *AssetResource {
    model.data = data
    return &model
}

func (model *AssetResource) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}