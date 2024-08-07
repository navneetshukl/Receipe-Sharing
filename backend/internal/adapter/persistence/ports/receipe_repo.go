package ports

import (
	"github.com/navneetshukl/receipe-sharing/internal/core/receipe"
)

type ReceipeRepo interface {
	InsertReceipe(data *receipe.Receipe) error
}
