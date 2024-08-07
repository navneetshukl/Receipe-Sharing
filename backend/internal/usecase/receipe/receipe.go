package receipe

import (
	"time"

	"github.com/navneetshukl/receipe-sharing/internal/adapter/persistence/ports"
	"github.com/navneetshukl/receipe-sharing/internal/core/receipe"
	"github.com/navneetshukl/receipe-sharing/pkg/logging"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReceipeUseCase struct {
	Receipe ports.ReceipeRepo
	Logs    logging.LogService
}

func NewReceipeUseCase(receipe ports.ReceipeRepo, logging logging.LogService) *ReceipeUseCase {
	return &ReceipeUseCase{Receipe: receipe,
		Logs: logging}
}

// AddReceipe function will add the receipe of user to db
func (ru *ReceipeUseCase) AddReceipe(userID string, data *receipe.Receipe) error {

	if data.Description == "" || len(data.Ingredients) == 0 || data.Name == "" {
		ru.Logs.ErrorLog("Some fields are missing ", nil)
		return receipe.ErrMissingField
	}

	userid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		ru.Logs.ErrorLog("Invalid user id", err)
		return receipe.ErrInvalidUserID
	}

	data.Created_At = time.Now()
	data.UserID = userid
	data.ID = primitive.NewObjectID()

	err = ru.Receipe.InsertReceipe(data)
	if err != nil {
		ru.Logs.ErrorLog("InsertReceipe", err)
		return receipe.ErrAddingReceipe
	}
	return nil

}
