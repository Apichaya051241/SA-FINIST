package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// Drug holds the schema definition for the Drug entity.
type Drug struct {
	ent.Schema
}

// Fields of the Drug.
func (Drug) Fields() []ent.Field {
	return []ent.Field{		
		field.String("name").NotEmpty(),		
}
}

// Edges of the Drug.
func (Drug) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Drug.Type).Ref("drugs").Unique(),
		edge.To("registerstores", Registerstore.Type).StorageKey(edge.Column("drug_drugs")),
		
	}
}
