package receipe

import (
	"log"

	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/ports"
	"github.com/navneetshukl/receipe-sharing/internal/core/receipe"
)

type ReceipeUseCase struct {
	Receipe ports.ReceipeRepo
}

func NewReceipeUseCase(receipe ports.ReceipeRepo) *ReceipeUseCase {
	return &ReceipeUseCase{Receipe: receipe}
}

func (ru *ReceipeUseCase) AddReceipe(data receipe.Receipe) error {

	err := ru.Receipe.InsertReceipe(data)
	if err != nil {
		log.Println("error in adding the receipe ", err)
	}
	return nil

}
