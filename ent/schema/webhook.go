package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Webhook holds the schema definition for the Webhook entity.
type Webhook struct {
	ent.Schema
}

// Fields of the Webhook.
func (Webhook) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("id"),
		field.String("name").
			MaxLen(255),
		field.String("url").
			MaxLen(255),
		field.Bool("required_verification").
			Default(false),
		field.String("secret_key").
			Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.UUID("user_id", uuid.UUID{}),
	}
}

// Edges of the Webhook.
func (Webhook) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("webhooks").
			Field("user_id").
			Required().Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
