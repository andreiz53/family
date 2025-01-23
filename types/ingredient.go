package types

import "github.com/jackc/pgx/pgtype"

type Ingredient struct {
	ID        pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	DeletedAt pgtype.Timestamp
	Name      string
}
