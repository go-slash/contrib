package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (u User) Fields() []ent.Field {
	return u.Groups().ByName("common").Fields()
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pet", Pet.Type).
			Unique().
			Annotations(entproto.Field(7)),
	}
}

func (u User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Service(entproto.BlockName("Zero")),
		entproto.Message(entproto.PackageName("zero"),
			entproto.NamedMessages(
				entproto.NamedMessage("UpdateUserProfile").
					WithGroups(u.Groups().ByName("common")).
					WithExtraFields(
						entproto.ExtraField("extra_1", 10).WithType(entproto.TypeBool),
						entproto.ExtraField("extra_2", 11).WithType(entproto.TypeString),
					).
					WithSkipID(true),
			)),
	}
}

func (User) Groups() *entproto.FieldGroups {
	return entproto.Groups().
		Group("common", func(fg *entproto.FieldGroup) {
			fg.Fields = []ent.Field{
				field.Uint64("id").
					Unique().
					Annotations(entproto.Field(1)),

				field.String("user_name").
					Unique().
					Annotations(entproto.Field(2)),

				field.Time("joined").
					Immutable().
					Annotations(entproto.Field(3)),

				field.Uint("points").
					Annotations(entproto.Field(4)),

				field.Uint64("exp").
					Annotations(entproto.Field(5)),

				field.Enum("status").
					Values("pending", "active").
					Annotations(
						entproto.Field(6),
						entproto.Enum(map[string]int32{
							"pending": 1,
							"active":  2,
						}),
					),
			}
		})
}
