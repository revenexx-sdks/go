package models

import (
    "encoding/json"
    "errors"
)

// AttributeField One renderable field. A superset of the manifest's `Field`:
// the three additions (`localized`, `channel_scoped`, `storage`) carry what a
// static manifest never has to say, because a manifest's fields are columns
// and these are keys inside one.
type AttributeField struct {
    // One value per channel rather than one value.
    ChannelScoped bool `json:"channel_scoped"`
    // Dotted read paths, most specific first — the documented precedence
    // (channel+locale → locale → channel → common). `common` is always last
    // and always present, because early imports wrote there whatever the
    // attribute's flags say.
    From []string `json:"from"`
    // Attribute-group code — the section this field belongs in.
    Group string `json:"group"`
    // That section's heading, resolved for the requested locale — so a form can
    // be built without reading `attribute_groups` as well.
    GroupLabel string `json:"group_label"`
    // Resolved for the requested locale, falling back to English, then to the
    // code.
    Label string `json:"label"`
    // One value per locale rather than one value.
    Localized bool `json:"localized"`
    // The attribute code — the key the value is stored under.
    Name string `json:"name"`
    // Present on select / multi-select. Two sources, one shape: rows of
    // `attribute_options` for an enumeration the attribute owns, or the records
    // of a reference entity for an attribute that points at one. Empty is an
    // answer: the list has no members yet.
    Options []AttributeFieldOption `json:"options"`
    // The family's ordering of this attribute, falling back to the attribute's
    // own.
    Position int `json:"position"`
    // The field must not be edited in this context. Today the one cause is a
    // variant axis on a product model; `readonly_reason` says which.
    Readonly bool `json:"readonly"`
    // Why the field is locked — a variant axis on a product model is set on its
    // variants.
    ReadonlyReason string `json:"readonly_reason"`
    // Present when the options ARE a reference entity's records: the code of that
    // entity, so a client can offer to manage the values rather than only pick
    // from them.
    ReferenceEntity string `json:"reference_entity"`
    // The family's `is_required`, narrowed to the requested channel when
    // `required_channels` names any.
    Required bool `json:"required"`
    // Where the value lives. Absent on an app whose custom fields are plain
    // columns — then the field name IS the column.
    Storage AttributeFieldStorage `json:"storage"`
    // The control to draw. Mapped from `attributes.type`, which carries no CHECK
    // on purpose — an unknown type answers 'text' rather than nothing.
    Type string `json:"type"`
    // The attribute's `is_unique` — the value is meant to identify the product.
    // Advisory: no index enforces it, so a client that cares has to check.
    Unique bool `json:"unique"`
    // Offered units of a `measure` field, from the attribute's `config.units`.
    Units []string `json:"units"`
    // The limits the value has to satisfy, ready to hand to a form validator.
    // Only the seven keys below are republished; anything else the tenant stored
    // in `attributes.validation` stays there.
    Validation AttributeFieldValidation `json:"validation"`

    // Used by Decode() method
    data []byte
}

func (model AttributeField) New(data []byte) *AttributeField {
    model.data = data
    return &model
}

func (model *AttributeField) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}