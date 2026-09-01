package models

import (
    "encoding/json"
    "errors"
)

// AttributesUpdateRequest Partial update — omitted fields keep their
// current value.
type AttributesUpdateRequest struct {
    // The attribute's stable identifier — the KEY its value is stored under
    // inside `attribute_values`, and the name a category rule addresses as
    // `attribute:<code>`. Unique per (`entity_type`, `entity_ref`) in this
    // tenant.
    Code string `json:"code"`
    // Type-specific settings; which keys apply depends on `type`. The ones this
    // app reads: `units` (the unit list a measure attribute offers) and
    // `reference_entity` (which entity a reference attribute draws its options
    // from). The ones the cockpit edits alongside them: `unit`, `metric_family`,
    // `decimals_allowed`, `asset_family`, `max_file_size`, `allowed_extensions`.
    Config interface{} `json:"config"`
    // Narrows `entity_type` to ONE reference entity or asset family, by its code
    // — the attributes of `brand` rather than of every reference entity. Null
    // for a plain product attribute.
    EntityRef string `json:"entity_ref"`
    // Which kind of record carries this attribute: 'product' for the catalog
    // itself, 'reference_entity', 'asset' or 'category' for the other things in
    // this app that have attributes. Deliberately carries no CHECK — a tenant
    // that models a fifth kind is served on it too.
    EntityType string `json:"entity_type"`
    // The `attribute_groups` row this attribute is filed under — the form
    // section it appears in. Null is ungrouped, and an ungrouped field is
    // rendered after every section that has a name.
    GroupId string `json:"group_id"`
    // Offer this attribute as a filter in a product list. `GET /products/grid`
    // reports exactly these attributes in its `filters` array, and nothing else
    // reads the flag.
    IsFilterable bool `json:"is_filterable"`
    // Declares that the value identifies the product — an EAN, a manufacturer
    // part number. It is metadata a form and an importer read: no database index
    // enforces it, because the value lives inside jsonb rather than in a column.
    IsUnique bool `json:"is_unique"`
    // The field label a person sees, keyed by language tag. Resolution falls back
    // to English and then to the code, so an untranslated attribute is still
    // renderable.
    Labels interface{} `json:"labels"`
    // True → the record holds ONE VALUE PER LOCALE, under
    // `attribute_values.locale_specific.<locale>.<code>`. False → one value,
    // under `attribute_values.common.<code>`. This flag is what decides where a
    // write goes.
    Localizable bool `json:"localizable"`
    // Where the field sits inside its group. A family may override it for its own
    // form through `family_attributes.position`; this is the attribute's default.
    Position int `json:"position"`
    // True → one value PER CHANNEL, under
    // `attribute_values.channel_specific.<channel>.<code>`. Set together with
    // `localizable` it means one value per channel AND locale, in
    // `channel_locale_specific`.
    Scopable bool `json:"scopable"`
    // Which editor the value asks for — 'text', 'select', 'metric', 'price',
    // 'asset_collection', 'reference_entity'. Carries no CHECK on purpose: an
    // integrator adds a type, and `GET /products/attribute-schema` maps an
    // unknown one onto a text field rather than refusing to answer.
    Type string `json:"type"`
    // Show this attribute as a COLUMN in the product grid. `GET /products/grid`
    // returns a column definition and a per-row value for exactly these.
    UsableInGrid bool `json:"usable_in_grid"`
    // Limits a value has to satisfy, as a flat object. The seven keys a client
    // can act on are `min`, `max`, `min_length`, `max_length`, `pattern`,
    // `min_items`, `max_items` — `GET /products/attribute-schema` republishes
    // those and leaves anything else the tenant stored untouched.
    Validation interface{} `json:"validation"`

    // Used by Decode() method
    data []byte
}

func (model AttributesUpdateRequest) New(data []byte) *AttributesUpdateRequest {
    model.data = data
    return &model
}

func (model *AttributesUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}