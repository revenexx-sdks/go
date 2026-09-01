package models

import (
    "encoding/json"
    "errors"
)

// ReferenceEntityRecordsCreateRequest model.
type ReferenceEntityRecordsCreateRequest struct {
    // Every attribute value the record carries, in ONE jsonb document — the
    // core of an attribute-driven PIM. A record's properties are not columns
    // here: they are rows in `attributes`, selected per family by
    // `family_attributes`, and their values live under their attribute CODE
    // inside this object.
    // 
    // Four buckets, and an attribute's own flags decide which one it writes to:
    // 
    // `common`                    the attribute is neither localizable nor
    // scopable — one value, full stop.
    // `{"common": {"net_weight": 2.4, "colour": "black"}}`
    // `locale_specific`           `localizable`: one value per language tag.
    // `{"locale_specific": {"de_DE": {"name": "Akku-Bohrschrauber"}}}`
    // `channel_specific`          `scopable`: one value per channel.
    // `{"channel_specific": {"b2b": {"minimum_order_quantity": 6}}}`
    // `channel_locale_specific`   both: one value per channel AND language tag.
    // `{"channel_locale_specific": {"b2b": {"de_DE": {"description": "…"}}}}`
    // 
    // A reader takes the most specific bucket that carries the code and falls
    // back through locale, then channel, then `common`. `common` is always last
    // and always consulted, because early imports wrote everything there whatever
    // an attribute's flags said — a reader that skipped it reports an imported
    // catalog as empty. `GET /products/attribute-schema` answers, per field, the
    // exact path a value belongs at (`storage.path`) and that full fallback order
    // (`from`), so no client has to re-derive any of this.
    // 
    // The value itself is whatever the attribute's `type` implies: a string, a
    // number, a boolean, an option CODE for a select (never its label), a list of
    // codes for a multi-select, `{"amount": …, "unit": …}` for a measure, a
    // list of `{"amount": …, "currency": …}` for a price, an asset code for
    // media.
    // 
    // Defaults to `{}`, and an empty object is a normal state — a record nobody
    // has enriched yet. The declared type also admits an array only because every
    // jsonb column of this app shares one mapping; an array is not meaningful
    // here and every reader in this app treats a non-object as empty.
    // 
    // Which attributes a record of this entity has comes from `attributes` rows
    // with `entity_type: "reference_entity"` and `entity_ref` equal to the
    // entity's code — `GET
    // /products/attribute-schema?entity_type=reference_entity&entity_ref=brand`
    // answers it in one call.
    AttributeValues interface{} `json:"attribute_values"`
    // The record's stable identifier — the value a product stores when it
    // points at this record, the same way a select stores an option code. Unique
    // within the entity.
    Code string `json:"code"`
    // What the record is called, per language tag — the text a picker shows
    // while the code is what gets written.
    Labels interface{} `json:"labels"`
    // Which reference entity this record belongs to.
    ReferenceEntityId string `json:"reference_entity_id"`

    // Used by Decode() method
    data []byte
}

func (model ReferenceEntityRecordsCreateRequest) New(data []byte) *ReferenceEntityRecordsCreateRequest {
    model.data = data
    return &model
}

func (model *ReferenceEntityRecordsCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}