package models

import (
    "encoding/json"
    "errors"
)

// OrderListItemInput model.
type OrderListItemInput struct {
    // The catalogue category the article sat in when the position was saved, as a
    // slug. Kept so a long list can be grouped the way the shop groups it without
    // a call to the catalogue.
    CategorySlug string `json:"category_slug"`
    // The cost centre this position books to, as the tenant's ERP names it. Free
    // text and not our enum. It survives into the ORDER position, which has a
    // `cost_center` column; a CART line has none, so the cart conversion carries
    // it in the line snapshot instead.
    CostCenterId string `json:"cost_center_id"`
    // The buyer's OWN article number for this article — what their purchasing
    // system calls it, which is rarely what the shop calls it. Free text, and the
    // field a B2B buyer searches their own lists by.
    CustomSku string `json:"custom_sku"`
    // The article image at the time the position was saved, as a URL or a path
    // — a snapshot like `name`, and nothing here refreshes it. It rides into
    // the cart line and the order position in their snapshot, because neither has
    // a column for it.
    Image string `json:"image"`
    // Free-form data the tenant keeps on the position. Never read by this app; it
    // travels into the cart line / order position snapshot untouched. A write
    // replaces the whole document rather than merging into it.
    Metadata interface{} `json:"metadata"`
    // The article name AS IT WAS when the position was saved. A snapshot on
    // purpose: the list is the buyer's own record, so a renamed or withdrawn
    // article still reads the way they wrote it down.
    Name string `json:"name"`
    // Sort order within the list, ascending — the order the positions
    // collection returns by default and the order the conversions hand the lines
    // over in. Neither dense nor unique: an add with no `position` of its own
    // takes the list's current position COUNT, so removing a position from the
    // middle and adding another leaves two rows sharing a number. A bulk replace
    // assigns the array index the same way, so it renumbers only the positions it
    // is not given explicitly.
    Position int `json:"position"`
    // Per-position notes the buyer wrote — an engraving, a delivery
    // instruction, a reference for the picker. An ARRAY OF STRINGS, one entry per
    // line; the order conversion joins them with newlines into the order
    // position's single `position_text`, and the cart conversion carries the
    // array in the line snapshot.
    PositionTexts []string `json:"position_texts"`
    // Unit price snapshot — what the buyer saw when they saved the position, in
    // whatever way the catalogue quoted it. It is a record, not a live price: the
    // cart and the order reprice on their own terms, so this never becomes what
    // somebody is charged.
    Price float64 `json:"price"`
    // The catalogue product this position stands for. One of `product_id` / `sku`
    // must be set (the database enforces it); this is the identity the products
    // app answers to, and the one `reject_unknown_articles` and the conversions
    // check against.
    ProductId string `json:"product_id"`
    // How much of the article the list holds. Greater than zero — the database
    // refuses the rest — and fractional to three decimals, because a B2B
    // position may be 2.5 metres or 0.75 kilos.
    Quantity float64 `json:"quantity"`
    // The article number as the catalogue knows it — the alternative identity
    // to `product_id`, and the one an ERP integration usually joins on.
    Sku string `json:"sku"`
    // The catalogue subcategory, as a slug. Same purpose as `category_slug`, one
    // level down.
    SubcategorySlug string `json:"subcategory_slug"`
    // The VAT rate that applied when the position was saved, as a PERCENT (19 =
    // 19 %). Four decimals so a rate like 8.25 % survives; carts and orders
    // document the same field the same way, and the conversion forwards the
    // number unchanged.
    TaxRate float64 `json:"tax_rate"`
    // The unit `quantity` counts in, in the tenant's own words. Deliberately open
    // text and deliberately NOT a vocabulary: a B2B catalogue units in pieces,
    // metres, kilos, rolls and pallets, and any closed list published here would
    // be a guess.
    Unit string `json:"unit"`

    // Used by Decode() method
    data []byte
}

func (model OrderListItemInput) New(data []byte) *OrderListItemInput {
    model.data = data
    return &model
}

func (model *OrderListItemInput) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}