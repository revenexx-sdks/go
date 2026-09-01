package models

import (
    "encoding/json"
    "errors"
)

// PriceEntry One rung of one item’s quantity ladder inside one price list.
// The ladder IS the set of entries sharing an identity (product_id or sku);
// the amount is in the LIST’s currency and on the LIST’s tax basis.
type PriceEntry struct {
    // When the entry was created.
    CreatedAt string `json:"created_at"`
    // The entry itself — one rung of one item’s quantity ladder.
    Id string `json:"id"`
    // Free-form bag, unvalidated and never read by this app: whatever JSON object
    // you write round-trips exactly. Its keys are the integration’s own, e.g.
    // {"source_system": "erp", "imported_batch": "2026-02-14"}.
    Metadata interface{} `json:"metadata"`
    // The price list this entry belongs to, and therefore the currency and tax
    // basis its amount is on. Set from the path on write.
    PriceListId string `json:"price_list_id"`
    // `standard` is a number. `on_request` is the explicit no-price marker: it
    // STOPS resolution for this item on this list and answers price-on-request,
    // even where a cheaper list exists — the list is authoritative for this
    // buyer and it says "ask us".
    PriceType string `json:"price_type"`
    // The product this rung prices. An entry needs `product_id` or `sku` (a row
    // CHECK enforces it); an entry that carries both prices whichever of the two
    // the resolve item names.
    ProductId string `json:"product_id"`
    // Lowest quantity this price applies from (Staffelpreis). The ladder for one
    // item is the set of entries sharing its identity: the rung with the HIGHEST
    // quantity_min at or below the requested quantity wins, and below the first
    // rung the first rung’s price applies — a minimum order quantity belongs
    // to the catalog, not to the ladder.
    QuantityMin float64 `json:"quantity_min"`
    // The article number this rung prices, for a price book keyed by SKU rather
    // than by product id — matched exactly, never normalised or case-folded.
    Sku string `json:"sku"`
    // The unit of measure the price is per — ‘pcs’, ‘m’, ‘kg’, a
    // packaging size. Free text: this app neither validates nor converts it, and
    // the `quantity` of a resolve call is counted in it.
    Unit string `json:"unit"`
    // Price for ONE unit of `unit`, expressed in the list’s `currency` and on
    // the list’s `tax_basis` — a decimal amount in major units (19.90 EUR),
    // never minor units/cents. Stored at 4 decimals so a per-1000-piece price
    // survives, and echoed back exactly as it was written; only DERIVED amounts
    // (net, gross, line totals) are rounded to the tenant’s `price_precision`.
    UnitPrice float64 `json:"unit_price"`
    // When the entry last changed. A bulk adjust only writes the rows whose price
    // actually moved, so this is a real "the price changed here" marker.
    UpdatedAt string `json:"updated_at"`
    // Start of this entry’s own validity; null = open-ended. This is how a
    // promo price is expressed — a second rung for the same item and quantity,
    // live only for its window.
    ValidFrom string `json:"valid_from"`
    // End of this entry’s own validity; null = open-ended. Outside the window
    // the rung is skipped and the ladder resolves as if it were not there.
    ValidUntil string `json:"valid_until"`

    // Used by Decode() method
    data []byte
}

func (model PriceEntry) New(data []byte) *PriceEntry {
    model.data = data
    return &model
}

func (model *PriceEntry) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}