package schema

import (
	"entgo.io/ent"
	. "github.com/tarqeem/template/utl/uent"
)

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		String("name"),
		Passowrd(),
		Email(),
		Phone(),
		Created_at(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
