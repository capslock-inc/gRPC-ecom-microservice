package model

import "github.com/gofrs/uuid"

type Product struct {
	id       uuid.UUID
	Name     string
	Price    int
	descript string
}
