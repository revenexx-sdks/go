package models

import (
    "encoding/json"
    "errors"
)

// Template model.
type Template struct {
    // 
    BodyHtml string `json:"body_html"`
    // 
    BodyText string `json:"body_text"`
    // 
    Channel string `json:"channel"`
    // 
    ContentSid string `json:"content_sid"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Design []interface{} `json:"design"`
    // 
    Enabled bool `json:"enabled"`
    // 
    HasUnpublishedChanges string `json:"has_unpublished_changes"`
    // 
    Id string `json:"id"`
    // 
    IsPublished string `json:"is_published"`
    // 
    Key string `json:"key"`
    // 
    LayoutId string `json:"layout_id"`
    // 
    LifecycleState string `json:"lifecycle_state"`
    // 
    Locale string `json:"locale"`
    // 
    Markets []interface{} `json:"markets"`
    // 
    MessageClass string `json:"message_class"`
    // 
    PublishedVersionId string `json:"published_version_id"`
    // 
    SourceLibraryKey string `json:"source_library_key"`
    // 
    Subject string `json:"subject"`
    // 
    TenantId string `json:"tenant_id"`
    // 
    TestMode bool `json:"test_mode"`
    // 
    Title string `json:"title"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    UsesRawHtml string `json:"uses_raw_html"`
    // 
    ValidFrom string `json:"valid_from"`
    // 
    ValidUntil string `json:"valid_until"`
    // 
    VariableDefaults []interface{} `json:"variable_defaults"`
    // 
    Variables []interface{} `json:"variables"`
    // 
    WhatsappCategory string `json:"whatsapp_category"`

    // Used by Decode() method
    data []byte
}

func (model Template) New(data []byte) *Template {
    model.data = data
    return &model
}

func (model *Template) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}