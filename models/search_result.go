package models

import (
    "encoding/json"
    "errors"
)

// SearchResult A Typesense search response, passed through verbatim.
type SearchResult struct {
    // 
    FacetCounts []FacetCount `json:"facet_counts"`
    // Total matching documents.
    Found int `json:"found"`
    // 
    Hits []SearchHit `json:"hits"`
    // Documents searched.
    OutOf int `json:"out_of"`
    // 1-based page this result is for.
    Page int `json:"page"`
    // 
    SearchTimeMs int `json:"search_time_ms"`

    // Used by Decode() method
    data []byte
}

func (model SearchResult) New(data []byte) *SearchResult {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *SearchResult) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}