package models

import (
    "encoding/json"
    "errors"
)

// ChannelPolicy The visibility policy in force for the resolved channel.
type ChannelPolicy struct {
    // Always 'channel' — the scope dimension this app provides.
    Dimension string `json:"dimension"`
    // The header name Baseline uses for this dimension. Through api.revenexx.com
    // it does NOT reach the app — the gateway builds a fresh request downstream
    // and forwards only its own headers — so use `?channel=` (or `channel` in
    // the body of POST /channels/visibility) instead. The header path applies to
    // a direct in-cluster call to the app.
    Header string `json:"header"`
    // The tenant setting, echoed: what `status = 'inactive'` DOES. 'serve' makes
    // it a label and the channel still resolves; 'block' makes resolution fail
    // with reason 'channel_inactive', and the policy then falls back to the
    // tenant answer.
    InactiveChannelBehavior string `json:"inactive_channel_behavior"`
    // The claim path in the forwarded identity token that names the active
    // channel, tried after the query and the header and before the default
    // channel.
    JwtPath string `json:"jwt_path"`
    // How Baseline matches the dimension — 'single': a request is in exactly
    // one channel at a time, never a set.
    MatchMode string `json:"match_mode"`
    // The tenant setting, echoed: whether a request naming no channel is refused
    // rather than falling back to the default channel. On POST
    // /channels/visibility that refusal is the single 400 this app makes of its
    // own accord.
    RequireChannelContext bool `json:"require_channel_context"`
    // Whether the answer came from the tenant setting or this channel's own
    // override. Only a channel that actually resolved gets a say — a blocked or
    // unknown channel falls back to 'tenant'.
    Source string `json:"source"`
    // The tenant-wide baseline, so a caller can see what this channel overrode.
    // Equal to `unassigned_visibility` whenever `source` is 'tenant'.
    TenantDefault string `json:"tenant_default"`
    // What a row with NO channel assignment means. 'all' is Baseline's
    // open-by-default semantic, reproduced exactly; 'assigned_only' is the closed
    // assortment the _scoped view cannot express.
    UnassignedVisibility string `json:"unassigned_visibility"`

    // Used by Decode() method
    data []byte
}

func (model ChannelPolicy) New(data []byte) *ChannelPolicy {
    model.data = data
    return &model
}

func (model *ChannelPolicy) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}