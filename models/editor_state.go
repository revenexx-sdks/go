package models

import (
    "encoding/json"
    "errors"
)

// EditorState Everything the blökkli editor runs on, for one page in one
// language, materialized at the current point of the undo history. The theme
// adapter maps it 1:1 onto blökkli's MappedState.
type EditorState struct {
    // Whether the caller may write. False means every write answers 409 until
    // `POST …/take-ownership` — so the editor should go read-only rather than
    // let someone type into a refusal.
    CurrentUserIsOwner bool `json:"currentUserIsOwner"`
    // Every entity-reference field of every block — the fields an editor drags
    // a product or a media item into.
    DroppableFieldValues []interface{} `json:"droppableFieldValues"`
    // The open working copy, or `null` when nobody has started editing — in
    // which case the state shown is simply the published one.
    EditState interface{} `json:"editState"`
    // What the tenant's settings allow, so a client hides a control instead of
    // discovering the refusal.
    Features interface{} `json:"features"`
    // The block tree, flattened into one entry per (host, field) pair. This is
    // the list the editor renders and drops into.
    Fields []interface{} `json:"fields"`
    // Analyze findings that were dismissed for this page, so the editor stops
    // reporting them.
    IgnoredAnalyzeIdentifiers []string `json:"ignoredAnalyzeIdentifiers"`
    // The language this whole state was resolved for — the `?langcode` that was
    // applied, or the page's source language.
    Langcode string `json:"langcode"`
    // The page-level field values the edit state changed, merged
    // source-then-language — `{ "title": …, "slug": …, "meta": … }`.
    // Empty when nobody edited the page itself, only its blocks.
    MutatedEntity interface{} `json:"mutatedEntity"`
    // The PAGE-level display options after the unpublished changes, as a flat
    // `option key → value` map. Theme-defined.
    MutatedHostOptions interface{} `json:"mutatedHostOptions"`
    // Every block's display options after the unpublished changes, keyed by block
    // uuid: `{ "<uuid>": { "background": "grey" } }`.
    MutatedOptions interface{} `json:"mutatedOptions"`
    // The undo/redo history, oldest first. Its length and
    // `editState.currentIndex` are what an undo button and a history sidebar are
    // drawn from.
    Mutations []interface{} `json:"mutations"`
    // The page itself, with the unpublished edits already applied — so the
    // title here is what publishing would store, not what is stored now.
    Page interface{} `json:"page"`
    // Every string field of every block, flattened. It is what the translation
    // view and the CSV export are built on — one row per translatable string.
    TextFieldValues []interface{} `json:"textFieldValues"`
    // Every language this page exists in, so the editor can offer a language
    // switcher that shows what is missing.
    Translations []interface{} `json:"translations"`
    // Why publishing would be refused right now. Empty means `POST …/publish`
    // succeeds without `force`.
    Violations []interface{} `json:"violations"`

    // Used by Decode() method
    data []byte
}

func (model EditorState) New(data []byte) *EditorState {
    model.data = data
    return &model
}

func (model *EditorState) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}