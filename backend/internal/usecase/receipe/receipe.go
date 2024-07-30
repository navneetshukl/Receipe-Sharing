package receipe

import (
	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/ports"
	"github.com/navneetshukl/receipe-sharing/internal/core/receipe"
	"github.com/navneetshukl/receipe-sharing/pkg/logging"
)

type ReceipeUseCase struct {
	Receipe ports.ReceipeRepo
	Logs    logging.LogService
}

func NewReceipeUseCase(receipe ports.ReceipeRepo, logging logging.LogService) *ReceipeUseCase {
	return &ReceipeUseCase{Receipe: receipe,
		Logs: logging}
}

func (ru *ReceipeUseCase) AddReceipe(data receipe.Receipe) error {

	err := ru.Receipe.InsertReceipe(data)
	if err != nil {
		ru.Logs.ErrorLog("InsertReceipe ", err)
		return receipe.ErrAddingReceipe
	}
	return nil

}
