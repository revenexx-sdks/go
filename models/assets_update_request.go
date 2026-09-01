package models

import (
    "encoding/json"
    "errors"
)

// AssetsUpdateRequest Partial update — omitted fields keep their current
// value.
type AssetsUpdateRequest struct {
    // The asset family this asset belongs to — which attributes it carries and
    // how its file is named. A create falls back to the `default_asset_family`
    // tenant setting when the body names none.
    AssetFamilyId string `json:"asset_family_id"`
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
    // Which attributes an asset of this family has comes from `attributes` rows
    // with `entity_type: "asset"` and `entity_ref` equal to the family's code —
    // alt text, copyright, an expiry date.
    AttributeValues interface{} `json:"attribute_values"`
    // The asset's stable identifier within its family — the value a product's
    // media attribute stores. Unique per family.
    Code string `json:"code"`
    // The path the CDN serves this asset under — the convenient value for
    // rendering. It changes when the file is moved, so never join on it.
    DeliveryPath string `json:"delivery_path"`
    // Absolute URL of an externally hosted file. Required when `source` is
    // `external`, and accepted only when the tenant has `allow_external_media` on
    // and the host is on its `external_media_allowed_hosts` list — `POST
    // /products/assets` is the only place an external URL can enter the catalog,
    // so it is the only place those are enforced.
    ExternalUrl string `json:"external_url"`
    // Where the bytes live: 'storage' is this platform's object store and needs
    // `storage_asset_id`, 'external' is somebody else's host and needs
    // `external_url`. The database enforces the pair, so neither half can be
    // stored on its own.
    Source string `json:"source"`
    // The stable `ast_…` id of the storage object. It survives a rename or a
    // folder move, which is exactly why it and not the delivery path is the
    // identifier. Required when `source` is `storage`.
    StorageAssetId string `json:"storage_asset_id"`

    // Used by Decode() method
    data []byte
}

func (model AssetsUpdateRequest) New(data []byte) *AssetsUpdateRequest {
    model.data = data
    return &model
}

func (model *AssetsUpdateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}