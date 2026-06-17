package models

import (
    "encoding/json"
    "errors"
)

// AttributeRelationship Model
type AttributeRelationship struct {
    // Attribute creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Attribute update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Is attribute an array?
    Array bool `json:"array"`
    // Error message. Displays error generated on failure of creating or deleting
    // an attribute.
    Error string `json:"error"`
    // Attribute Key.
    Key string `json:"key"`
    // How deleting the parent document will propagate to child documents.
    OnDelete string `json:"onDelete"`
    // The ID of the related collection.
    RelatedCollection string `json:"relatedCollection"`
    // The type of the relationship.
    RelationType string `json:"relationType"`
    // Is attribute required?
    Required bool `json:"required"`
    // Whether this is the parent or child side of the relationship
    Side string `json:"side"`
    // Attribute status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Is the relationship two-way?
    TwoWay bool `json:"twoWay"`
    // The key of the two-way relationship.
    TwoWayKey string `json:"twoWayKey"`
    // Attribute type.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model AttributeRelationship) New(data []byte) *AttributeRelationship {
    model.data = data
    return &model
}

func (model *AttributeRelationship) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}