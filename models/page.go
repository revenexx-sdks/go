package models

import (
    "encoding/json"
    "errors"
)

// Page One addressable page of the storefront: its metadata and publish
// pointer. Its CONTENT is not here — blocks live behind the editor and
// delivery routes.
type Page struct {
    // Identifiers of findings the blökkli analyze feature was told to stop
    // reporting for this page. Written by the `set_ignored_analyze` mutation and
    // carried through publish, so dismissing a finding survives the next edit.
    AnalyzeIgnored []string `json:"analyze_ignored"`
    // The page TYPE, e.g. `standard` or a landing-page type the theme defines. It
    // decides which fields the editor offers and which template the theme
    // renders; the value set belongs to the active theme, not to this app.
    Bundle string `json:"bundle"`
    // When the page was created.
    CreatedAt string `json:"created_at"`
    // The user id that created the page.
    CreatedBy string `json:"created_by"`
    // The tombstone. A soft-deleted page is never listed, never delivered and
    // answers 404 — and it drops out of the unique slug index at once, so
    // deleting a page frees its slug immediately.
    DeletedAt string `json:"deleted_at"`
    // Page-level blökkli display options, as a flat `option key → value` map
    // — the options that belong to the PAGE rather than to a block (background,
    // width, whether the header is shown). The keys are defined by the theme;
    // this app stores whatever the `update_host_options` mutation set.
    HostOptions interface{} `json:"host_options"`
    // The page id. Every editor and delivery route addresses a page by it, and it
    // never changes — publishing replaces a page's blocks, never the page.
    Id string `json:"id"`
    // The page's free-form metadata bag — SEO fields, social preview data,
    // whatever the theme asks the editor for. Nothing in this app reads a key of
    // it: it is stored, versioned into revisions and handed back to the renderer
    // untouched, so the theme owns its shape.
    Meta interface{} `json:"meta"`
    // The revision the storefront is currently serving. `null` means nothing has
    // ever been published, and delivery answers 404 for the page even when
    // `status` says `published`.
    PublishedRevisionId string `json:"published_revision_id"`
    // The path segment the storefront routes this page under, without a leading
    // slash. Unique per tenant among live pages, and `null` for a page that is
    // only ever reached by id. `GET /pages/delivery/page?slug=` matches it first
    // and the translations second.
    Slug string `json:"slug"`
    // The language the page was authored in. It is the fallback for every field a
    // translation leaves empty, so a page never renders as a hole.
    SourceLanguage string `json:"source_language"`
    // Where the page sits in the editorial lifecycle. Only `published` is ever
    // delivered, and only together with a `published_revision_id`.
    Status string `json:"status"`
    // The page title as an editor typed it, in the page's source language.
    // Publishing overwrites it with the title the edit state carries, so this is
    // always the last published (or last saved) wording.
    Title string `json:"title"`
    // When the page last changed. The default sort of `GET /pages/pages` is this
    // column descending, because "what did we touch last" is the question an
    // editorial list is opened with.
    UpdatedAt string `json:"updated_at"`
    // The user id that last changed the page — set by an update, a soft delete
    // and by publishing.
    UpdatedBy string `json:"updated_by"`

    // Used by Decode() method
    data []byte
}

func (model Page) New(data []byte) *Page {
    model.data = data
    return &model
}

func (model *Page) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}