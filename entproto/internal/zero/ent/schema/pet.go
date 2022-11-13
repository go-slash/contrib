package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return nil
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pet").
			Unique().
			Annotations(entproto.Field(2)),
	}
}

func (Pet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(entproto.PackageName("zero")),
		entproto.Service(entproto.BlockName("Zero")),
	}
}
