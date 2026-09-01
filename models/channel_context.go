package models

import (
    "encoding/json"
    "errors"
)

// ChannelContext model.
type ChannelContext struct {
    // The channel that resolved, or null. Null on every answer where `resolved`
    // is false — including the everyday one on a tenant that has not created a
    // channel yet.
    Channel string `json:"channel"`
    // More than one channel claims is_default; the lowest position wins and this
    // says so.
    DefaultAmbiguous bool `json:"default_ambiguous"`
    // The visibility policy in force for the resolved channel.
    Policy ChannelPolicy `json:"policy"`
    // Why not, when resolved is false. Null when it resolved.
    Reason string `json:"reason"`
    // The channel code the request named, if any — lowercased and trimmed as it
    // was matched.
    Requested string `json:"requested"`
    // Whether a channel could be resolved for this request.
    Resolved bool `json:"resolved"`
    // Where the channel came from, in the order they are tried: 'body' (the
    // `channel` field, POST /channels/visibility only), 'query' (`?channel=`),
    // 'header' (x-revenexx-channel), 'jwt' (the scope_context.channel claim),
    // then 'default' (the channel flagged is_default). Null when nothing
    // resolved. Note that 'header' is not reachable through api.revenexx.com: the
    // gateway builds a fresh request to the app and copies a fixed set of headers
    // into it, and x-revenexx-channel is not among them — see `policy.header`.
    Source string `json:"source"`

    // Used by Decode() method
    data []byte
}

func (model ChannelContext) New(data []byte) *ChannelContext {
    model.data = data
    return &model
}

func (model *ChannelContext) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}