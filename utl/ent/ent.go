package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

const (
	MaxNormalNameLength = 300
	MinNameLen          = 2
)

func String(n string) ent.Field      { return field.String(n).MinLen(MinNameLen) }
func IntPositive(n string) ent.Field { return field.Int(n).Min(0) }
