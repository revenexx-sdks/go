package models

import (
    "encoding/json"
    "errors"
)

// PageCreateRequest A new page. Only the title is yours to supply —
// everything else has a tenant default behind it.
type PageCreateRequest struct {
    // The page type. Omit to take the default_page_bundle setting.
    Bundle string `json:"bundle"`
    // Page-level blökkli display options as a flat `option key → value` map.
    // Theme-defined; usually left out and set later from the editor.
    HostOptions interface{} `json:"hostOptions"`
    // The page's metadata bag (SEO and social fields). Stored and handed back
    // untouched — this app reads no key of it, so the theme decides what goes
    // in.
    Meta interface{} `json:"meta"`
    // The path segment the storefront routes it under, without a leading slash.
    // Unique per tenant among live pages; omit or send null for a page reached
    // only by id. Nothing here derives one from the title.
    Slug string `json:"slug"`
    // The language you are authoring in, and the fallback for every later
    // translation. Omit to take the default_source_language setting for the
    // request market.
    SourceLanguage string `json:"sourceLanguage"`
    // What the page is called, in its source language. Shown in the editorial
    // list and searched by `?q=`.
    Title string `json:"title"`

    // Used by Decode() method
    data []byte
}

func (model PageCreateRequest) New(data []byte) *PageCreateRequest {
    model.data = data
    return &model
}

func (model *PageCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}