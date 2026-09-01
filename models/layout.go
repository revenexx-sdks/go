package models

import (
    "encoding/json"
    "errors"
)

// Layout model.
type Layout struct {
    // 
    ColorAccent string `json:"color_accent"`
    // 
    ColorBg string `json:"color_bg"`
    // 
    ColorText string `json:"color_text"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    Enabled bool `json:"enabled"`
    // 
    FontFamily string `json:"font_family"`
    // 
    FooterNote string `json:"footer_note"`
    // 
    Id string `json:"id"`
    // 
    IsDefault bool `json:"is_default"`
    // 
    LegalName string `json:"legal_name"`
    // 
    LifecycleState string `json:"lifecycle_state"`
    // 
    LogoUrl string `json:"logo_url"`
    // 
    Markets []interface{} `json:"markets"`
    // 
    MenuLinks []interface{} `json:"menu_links"`
    // 
    Name string `json:"name"`
    // 
    PostalAddress string `json:"postal_address"`
    // 
    SenderName string `json:"sender_name"`
    // 
    SocialLinks []interface{} `json:"social_links"`
    // 
    SupportEmail string `json:"support_email"`
    // 
    TenantId string `json:"tenant_id"`
    // 
    UpdatedAt string `json:"updated_at"`
    // 
    ValidFrom string `json:"valid_from"`
    // 
    ValidUntil string `json:"valid_until"`
    // 
    Width string `json:"width"`

    // Used by Decode() method
    data []byte
}

func (model Layout) New(data []byte) *Layout {
    model.data = data
    return &model
}

func (model *Layout) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}