package models

import (
    "encoding/json"
    "errors"
)

// Model
type SyncRuleResource struct {
    // 
    CreatedAt string `json:"created_at"`
    // 
    Enabled bool `json:"enabled"`
    // 
    Id string `json:"id"`
    // 
    LastRunAt string `json:"last_run_at"`
    // 
    Options []interface{} `json:"options"`
    // 
    Schedule string `json:"schedule"`
    // 
    SftpAccountId string `json:"sftp_account_id"`
    // 
    SourcePath string `json:"source_path"`
    // 
    TargetFolderId string `json:"target_folder_id"`
    // 
    TenantId string `json:"tenant_id"`

    // Used by Decode() method
    data []byte
}

func (model SyncRuleResource) New(data []byte) *SyncRuleResource {
    model.data = data
    return &model
}

func (model *SyncRuleResource) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}