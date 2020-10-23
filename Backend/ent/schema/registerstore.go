package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Registerstore holds the schema definition for the Registerstore entity.
type Registerstore struct {
	ent.Schema
}

// Fields of the Registerstore.
func (Registerstore) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount").Positive(),
	}
}

// Edges of the Registerstore.
func (Registerstore) Edges() []ent.Edge {
	return []ent.Edge{
            		
			edge.From("stores", Store.Type).
            Ref("registerstores").
			Unique(),
			
			edge.From("users", User.Type).
            Ref("registerstores").
			Unique(),
			
			edge.From("drugs", Drug.Type).
            Ref("registerstores").
			Unique(),
		
			}
}