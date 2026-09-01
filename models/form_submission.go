package models

import (
    "encoding/json"
    "errors"
)

// FormSubmission model.
type FormSubmission struct {
    // When the submission arrived. This is the age the retention sweep measures
    // against `submission_retention_days`, and the column an inbox sorts by
    // (`order=created_at.desc`).
    CreatedAt string `json:"created_at"`
    // What the visitor typed — the substance of the submission, and the reason
    // this row is the payload of `form.submitted`.
    // 
    // It is an object keyed by the `name` of the definition node that collected
    // each value, so the keys of a submission are the named nodes of its form's
    // `definition` and nothing else. There is no fixed set of keys across forms:
    // a contact form yields `{name, email, message}`, a price request whatever
    // its operator built.
    // 
    // The VALUE type follows the input type, which is why this object is not
    // typed further: a `text`, `email` or `textarea` yields a string, a `number`
    // a number, a single `checkbox` a boolean, a `select`/`radio` the chosen
    // option value, a multi-select or a checkbox set an array of them, and a
    // `group` or `list` input nests an object or an array under its own name.
    // Nothing coerces them — a value arrives as the storefront sent it and is
    // stored as jsonb.
    // 
    // Two values are NOT here: the honeypot field, if the tenant configured one,
    // is stripped before the row is written (it is a trap, not an answer the
    // visitor gave), and the resolved notification recipient lives in `metadata`,
    // not in what somebody typed.
    Data interface{} `json:"data"`
    // The form this submission was made against. It is resolved at insert, so an
    // id no form in this tenant holds is a 404 and nothing is stored — a
    // submission with no form is a lead nobody can read.
    FormId string `json:"form_id"`
    // The form's slug as it stood when this submission arrived, copied onto the
    // row: the inbox filters by form without a join, and a submission still says
    // which form collected it after that form has been renamed. It does not
    // outlive a DELETED form — the foreign key cascades and takes the
    // submission with it. On a write the body's value WINS; omit it and the
    // form's own slug is copied in.
    FormSlug string `json:"form_slug"`
    // The submission's own id — what the inbox links to, and what a workflow
    // reading `form.submitted` gets handed.
    Id string `json:"id"`
    // Free-form metadata, plus what this app stamped on at insert. The recipient
    // is resolved ONCE, here, because this row is the payload of `form.submitted`
    // — a workflow reads the address off the event instead of re-resolving a
    // form's settings that may since have changed.
    Metadata FormSubmissionMetadata `json:"metadata"`
    // Where the submission came from. The storefront sends the
    // `window.location.pathname` of the page that carried the form, so this is
    // normally a path rather than an absolute URL; any other surface (an app, an
    // import) puts its own name here. Null when the caller sent none.
    Source string `json:"source"`
    // Inbox triage. `new` until somebody opens it, then `read`, and `archived`
    // once it is dealt with. `spam` is set by code in exactly one place — the
    // honeypot, and only while the tenant's spam_handling is 'flag'; under
    // 'reject' the submission is never stored at all. Default 'new'.
    Status string `json:"status"`
    // The tenant this row belongs to — the store slug, not an id. Set by the
    // platform from the authenticated context, never by a caller; a write that
    // carries it is ignored, and no request can read another tenant's rows by
    // sending a different one.
    TenantId string `json:"tenant_id"`
    // When the row was last written — a triage status change. It is not
    // evidence about the submitted data, which under the shipped policy cannot
    // change at all.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model FormSubmission) New(data []byte) *FormSubmission {
    model.data = data
    return &model
}

func (model *FormSubmission) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}