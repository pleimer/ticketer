package ent

//go:generate go run entc.go
//go:generate go run entgo.io/ent/cmd/ent generate  --feature sql/upsert ./schema
