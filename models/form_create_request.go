package models

import (
    "encoding/json"
    "errors"
)

// FormCreateRequest model.
type FormCreateRequest struct {
    // The form itself: a FormKit schema, held as a flat ARRAY of nodes (it
    // defaults to `[]`, never to an object) and rendered verbatim by the
    // storefront.
    // 
    // Read it as the field list. Every node carrying a non-empty `name` collects
    // one value and writes it into a submission's `data` under exactly that name
    // — the example below produces `{"company": …, "email": …, "message":
    // …}` — while `$el` content nodes and `$rxStep` step markers collect
    // nothing. Order is render order, and a `$rxStep` marker starts a new wizard
    // step.
    // 
    // See the `FormKitNode` schema for what a node may carry.
    // 
    // On the way IN a node is any object: this is unconstrained jsonb, FormKit
    // owns the grammar, and the one rule this app applies is the tenant's
    // `max_form_fields` ceiling counted over the nodes with a non-empty `name`.
    // Anything that is not an array at all is a 400.
    Definition []interface{} `json:"definition"`
    // Free-form metadata on the FORM, which this app neither reads nor writes:
    // yours to key however an integration needs, stored and returned verbatim.
    // (The metadata this app does write is on a SUBMISSION — see
    // `FormSubmissionMetadata`.)
    Metadata interface{} `json:"metadata"`
    // What this form is called in the Cockpit's form list. Operator-facing only
    // — the storefront never renders it, so renaming a form breaks no page.
    Name string `json:"name"`
    // Submit label, success message, per-form notify email, post-submit actions,
    // translations — see the `FormSettings` schema for every key that is read.
    // Unconstrained jsonb on the way in: nothing here is required and no key is
    // refused.
    Settings interface{} `json:"settings"`
    // URL-safe identifier, unique per tenant. This is the name a storefront
    // resolves a form by (`GET /v1/forms?slug=contact&status=live&limit=1`), so
    // it is part of the page's contract: changing it changes which form a page
    // renders. Lower-case letters, digits and inner hyphens. Taken already? That
    // is the 409 — one slug answers for one form.
    Slug string `json:"slug"`
    // Lifecycle. `draft` while it is being built; `live` once the storefront may
    // render it — the cover BFF resolves live forms only, so a draft is a 404
    // on the storefront and never a broken page; `archived` for a form that is
    // kept for its submissions but no longer offered. Default 'draft'.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model FormCreateRequest) New(data []byte) *FormCreateRequest {
    model.data = data
    return &model
}

func (model *FormCreateRequest) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}