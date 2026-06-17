package models

import (
    "encoding/json"
    "errors"
)

// Model
type SyncHistory struct {
    // 
    BytesSynced int `json:"bytes_synced"`
    // 
    CreatedAt string `json:"created_at"`
    // 
    DurationMs int `json:"duration_ms"`
    // 
    Error string `json:"error"`
    // 
    Id int `json:"id"`
    // 
    RuleId string `json:"rule_id"`
    // 
    RunId string `json:"run_id"`
    // 
    SourcePath string `json:"source_path"`
    // 
    Status string `json:"status"`
    // 
    TargetAssetId string `json:"target_asset_id"`
    // 
    TenantId string `json:"tenant_id"`

    // Used by Decode() method
    data []byte
}

func (model SyncHistory) New(data []byte) *SyncHistory {
    model.data = data
    return &model
}

func (model *SyncHistory) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}