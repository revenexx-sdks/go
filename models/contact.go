package models

import (
    "encoding/json"
    "errors"
)

// Contact A PERSON, and the unit that logs in: one platform user, one email,
// one role inside its organization. A contact without an organization is a
// standalone buyer, not an error.
type Contact struct {
    // When this person record was created in this app.
    CreatedAt string `json:"created_at"`
    // Login identity and the unique key of a person within the tenant. Changing
    // it changes the platform login with it. Two people at the same company
    // therefore need two addresses — a shared purchasing mailbox is one
    // contact, not several.
    Email string `json:"email"`
    // Id of the platform USER this contact is mirrored as — the account that
    // actually holds the password and the sessions. Written by the mirror and
    // ignored on every write a caller sends.
    ExternalUserId string `json:"external_user_id"`
    // Given name. Optional: an ERP import often has only a mailbox.
    FirstName string `json:"first_name"`
    // Primary key of the person record. What the timeline, the permission routes
    // and the principal resolution all name.
    Id string `json:"id"`
    // The main contact of its organization — who a merchant calls first. At
    // most one per company is the intent; the tenant's `primary_contact_required`
    // setting decides whether the last one may be demoted or deleted.
    IsPrimary bool `json:"is_primary"`
    // What this person does at the company — free text on purpose, because it
    // is a title and not a grant. The permission ladder is `role`; overloading a
    // job title with authority silently un-grants everyone the day the ledger is
    // enforced.
    JobTitle string `json:"job_title"`
    // Family name. Optional for the same reason.
    LastName string `json:"last_name"`
    // The language this person is written to in — BCP 47, and one of the
    // store's configured locales. Null falls back to the store default.
    Locale string `json:"locale"`
    // Amount ceiling for this person, in the market's currency: with the
    // `orders.approve` permission it is the most they may sign off. Null means no
    // ceiling. An amount, never a grant — the grant comes from the role.
    OrderApprovalLimit float64 `json:"order_approval_limit"`
    // The company this person belongs to. NULL is a legitimate state, not a
    // defect: a standalone buyer with no company behind them. Deleting the
    // organization sets this null and keeps the person.
    OrganizationId string `json:"organization_id"`
    // Direct number of this person, as somebody typed it — free text, no format
    // is enforced or normalized. E.164 is what an integration should send.
    Phone string `json:"phone"`
    // When a merchant approved or rejected the application. Null while nobody has
    // decided.
    RegistrationDecidedAt string `json:"registration_decided_at"`
    // Who decided — free text as the deciding client supplied it (an operator
    // id or an email address), not a resolvable user reference.
    RegistrationDecidedBy string `json:"registration_decided_by"`
    // Why the application was declined. Always recorded here; whether the
    // APPLICANT is ever told it is the tenant's `registration_reason_disclosed`
    // setting, because that is a legal decision and not a template one.
    RegistrationReason string `json:"registration_reason"`
    // Where this person's own application stands: 'approved' (the default, and
    // what an open store creates), 'pending' while a merchant has yet to decide,
    // 'rejected' once they declined. Only the approve/reject routes move it; it
    // is ignored on an ordinary update.
    RegistrationStatus string `json:"registration_status"`
    // The person's role INSIDE its organization, and the only thing permissions
    // are derived from. One of the tenant's own roles (GET /customers/roles); a
    // tenant that never edited the ledger has viewer, requester, buyer, approver,
    // admin. Also the team role on the platform mirror. There is no global role
    // — the same person in two companies is two contacts.
    Role string `json:"role"`
    // Whether this person may act: 'invited' has been created but has not
    // accepted, 'active' works, 'blocked' cannot log in. A create through the API
    // defaults to 'invited'; a self-registration in an open store lands 'active'.
    Status string `json:"status"`
    // The tenant this row belongs to — the store slug, not an id. Set by the
    // platform from the authenticated context, never by a caller; a write that
    // carries it is ignored, and no request can read another tenant's rows by
    // sending a different one.
    TenantId string `json:"tenant_id"`
    // When any column of this row last changed.
    UpdatedAt string `json:"updated_at"`

    // Used by Decode() method
    data []byte
}

func (model Contact) New(data []byte) *Contact {
    model.data = data
    return &model
}

func (model *Contact) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}