package proxy

import (
	"github.com/xidongc-wish/mgo/bson"
)

type Empty struct {}

type FindAndModifyMode int

// Query param for upper services
type QueryParam struct {
Filter 			bson.M
Fields 			bson.M
Limit			int64
Skip 			int64
Sort 			[]string
Distinctkey 	string
FindOne			bool
UsingIndex		[]string
Amp 			Amplifier
}

// Insert param for upper services
type InsertParam struct {
Docs 			[]interface{}
Amp 			Amplifier
}

// Remove param for upper services
type RemoveParam struct {
Filter			bson.M
Amp 			Amplifier
}

// Update param for upper services
type UpdateParam struct {
Filter			bson.M
Update			bson.M
Upsert          bool
Multi           bool
Amp 			Amplifier
}

// FindAndModify param for upper services
type FindModifyParam struct {
Filter			bson.M
Desired			bson.M
Mode			FindAndModifyMode
SortRule		[]string
Fields 			bson.M
Amp 			Amplifier
}

// Aggregate param for upper services
type AggregateParam struct {
Pipeline 		bson.M
}