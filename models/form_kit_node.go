package models

import (
    "encoding/json"
    "errors"
)

// FormKitNode One node of a form definition.
// 
// A definition is a FLAT ARRAY of these, and the storefront hands each one to
// `<FormKitSchema>` verbatim — it maps nothing, so every key FormKit
// understands works here whether or not it is named below (`options`, `if`,
// `rows`, `autocomplete`, `min`, `max`, `$cmp`, …). Three kinds of node
// occur:
// 
// • an INPUT node (`$formkit`) collects a value and, if it carries a
// `name`, contributes exactly one key to a submission's `data`;
// • a CONTENT node (`$el`) renders markup — a paragraph of legal text, a
// heading — and collects nothing;
// • a STEP MARKER (`$rxStep`) is a Revenexx extension the storefront
// consumes and strips before FormKit sees the node; it splits the flat array
// into wizard steps.
// 
// Only the four keys `name`, `label`, `placeholder` and `help` are read by
// Revenexx code at all (the last three are what the per-form i18n overlay
// translates). Everything else is FormKit's business.
type FormKitNode struct {
    // A CONTENT node instead of an input: a raw element name ('p', 'h2', 'div').
    // It collects no value and contributes no key to `data`.
    El string `json:"$el"`
    // An INPUT node: the FormKit input type — 'text', 'email', 'textarea',
    // 'number', 'select', 'checkbox', 'radio', 'date', 'group', 'list', … . The
    // set is FormKit's, not this app's, which is why nothing here enforces it and
    // no vocabulary is published for it; the storefront adds one input of its
    // own, `datepicker`, and three validation rules (`zip`, `companyName`,
    // `phoneNumber`).
    Formkit string `json:"$formkit"`
    // A Revenexx step marker. The storefront cuts the flat array at each marker
    // and renders the nodes that follow it as one wizard step, then removes the
    // marker before FormKit renders anything. A definition with no marker is a
    // single-step form.
    RxStep FormKitStepMarker `json:"$rxStep"`
    // The content of an `$el` node: a string of text, or nested nodes.
    Children string `json:"children"`
    // The hint under the input. Translatable.
    Help string `json:"help"`
    // What the visitor reads above the input. Translatable: the per-form i18n
    // overlay replaces it per locale.
    Label string `json:"label"`
    // The key this input writes into a submission's `data` — `{ "$formkit":
    // "email", "name": "email" }` here is the `"email"` key there, and that
    // correspondence is the whole contract between a form and its inbox. A node
    // with a non-empty `name` is a FIELD: only fields count against the tenant's
    // `max_form_fields`, so a form with twenty paragraphs of legal text and three
    // inputs is a three-field form. A `group` or `list` input nests, and its
    // `name` keys the nested object or array.
    Name string `json:"name"`
    // Placeholder text inside the input. Translatable.
    Placeholder string `json:"placeholder"`
    // A Revenexx hint about where the value comes from rather than what it looks
    // like. 'product' means the storefront prefills this input from the page
    // context or the query string (`?sku=…`) and renders it read-only — how a
    // price request knows which article it is about. Stripped before FormKit
    // renders the node.
    RxKind string `json:"rxKind"`
    // FormKit validation, in either notation FormKit accepts: the pipe string
    // 'required|email', or the array form. It is enforced in the browser by
    // FormKit — this API stores whatever `data` it is sent, so a server-side
    // integration must not treat it as a guarantee.
    Validation string `json:"validation"`

    // Used by Decode() method
    data []byte
}

func (model FormKitNode) New(data []byte) *FormKitNode {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *FormKitNode) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}