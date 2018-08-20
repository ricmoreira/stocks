package mresponse

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type StockMovCreate struct {
	ID             string            `json:"id,omitempty"`
	IDdb           objectid.ObjectID `json:"-" bson:"_id"`
	MovemementType string            `json:"MovemementType" bson:"MovemementType"`
	DocumentID     string            `json:"DocumentID" bson:"DocumentID"`
	Line           int32             `json:"Line" bson:"Line"`
	Quantity       float32           `json:"Quantity" bson:"Quantity"`
	ProductCode    string            `json:"ProductCode" bson:"ProductCode"`
	UnitOfMeasure  string            `json:"UnitOfMeasure" bson:"UnitOfMeasure"`
	Time           time.Time         `json:"Time" bson:"Time"`
}

type StockMovRead struct {
	ID             string            `json:"id,omitempty"`
	IDdb           objectid.ObjectID `json:"-" bson:"_id"`
	MovemementType string            `json:"MovemementType" bson:"MovemementType"`
	DocumentID     string            `json:"DocumentID" bson:"DocumentID"`
	Line           int32             `json:"Line" bson:"Line"`
	Quantity       float32           `json:"Quantity" bson:"Quantity"`
	ProductCode    string            `json:"ProductCode" bson:"ProductCode"`
	UnitOfMeasure  string            `json:"UnitOfMeasure" bson:"UnitOfMeasure"`
	Time           time.Time         `json:"Time" bson:"Time"`
}

type StockMovUpdate struct {
	ID             string            `json:"id,omitempty"`
	IDdb           objectid.ObjectID `json:"-" bson:"_id"`
	MovemementType string            `json:"MovemementType" bson:"MovemementType"`
	DocumentID     string            `json:"DocumentID" bson:"DocumentID"`
	Line           int32             `json:"Line" bson:"Line"`
	Quantity       float32           `json:"Quantity" bson:"Quantity"`
	ProductCode    string            `json:"ProductCode" bson:"ProductCode"`
	UnitOfMeasure  string            `json:"UnitOfMeasure" bson:"UnitOfMeasure"`
	Time           time.Time         `json:"Time" bson:"Time"`
}

type StockMovDelete struct {
	ID   string            `json:"id,omitempty"`
	IDdb objectid.ObjectID `json:"-" bson:"_id"`
}

type StockMovList struct {
	Total   int64            `json:"total"`
	PerPage int64            `json:"per_page"`
	Page    int64            `json:"page"`
	Items   *[]*StockMovRead `json:"items"`
}