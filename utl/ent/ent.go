package ent

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

var (
	MaxNormalNameLength = 300
	MinNameLen          = 2
	PassRegex           = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).+$`)
	EmailRegex          = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	PhoneRegex          = regexp.MustCompile(`^\+?[0-9]{1,3}-?[0-9]{3}-?[0-9]{3}-?[0-9]{4}$`)
	CreatedTableName    = "created_at"
	PasswordTableName   = "passowrd"
	PhoneTableName      = "phone"
)

// return a field `created_at`. Configure name with CreatedTableName
func Created() ent.Field { return field.Time(CreatedTableName).Default(time.Now) }

// return a string field with `n` name and `MinNameLen` minimum length
func String(n string) ent.Field { return field.String(n).MinLen(MinNameLen) }

// return a field `password`. Configure name with `CreatedTableName` and Regex
// with PassRegex
func Passowrd() ent.Field {
	return field.String(PasswordTableName).Sensitive().Match(PassRegex).NotEmpty()
}

// return a non negative integer field
func NonNegative(n string) ent.Field { return field.Int(n).Min(0) }

// return a field named `phone`. Configure name with `PhoneTableName` and Regex
// with `PhoneRegex`
func Phone() ent.Field {
	return field.String(PhoneTableName).MaxLen(MinNameLen).NotEmpty().Match(PhoneRegex).Unique()
}

// return a field named `email`. Configure name with `EmailTableName` and Regex
// with `EmailRegex`
func Email() ent.Field {
	return field.String("email").MaxLen(MaxNormalNameLength).
		NotEmpty().Match(EmailRegex)
}