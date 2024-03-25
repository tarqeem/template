package schema

import (
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	. "github.com/tarqeem/template/utl"
)

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	sp := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).+$`)
	ve := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	vp := regexp.MustCompile(`^\+?[0-9]{1,3}-?[0-9]{3}-?[0-9]{3}-?[0-9]{4}$`)

	return []ent.Field{
		EntString("name"),
		field.String("password").Sensitive().Match(sp).NotEmpty(),
		field.String("email").MaxLen(MaxNormalNameLength).NotEmpty().Match(ve).Unique(),
		field.String("phone").MaxLen(MinNameLen).NotEmpty().Match(vp).Unique(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
