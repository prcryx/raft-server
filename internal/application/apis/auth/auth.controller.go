package auth

import (
	"encoding/json"
	"net/http"

	"github.com/prcryx/datacheckr"
	e "github.com/prcryx/raft-server/internal/common/err"
	utils "github.com/prcryx/raft-server/internal/common/utils"
	"github.com/prcryx/raft-server/internal/domain/usecases"
)

//Auth DTO

type AuthDTO struct {
	Email    string
	Password string
}

type IAuthController interface {
	SignUpWithEmailAndPassword(http.ResponseWriter, *http.Request)
}

type AuthController struct {
	authUseCase *usecases.AuthUseCase
}

func NewAuthController(authUseCase *usecases.AuthUseCase) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
	}
}

func (ac *AuthController) SignUpWithEmailAndPassword(w http.ResponseWriter, request *http.Request) {

	//validating request
	authDTO := new(AuthDTO)
	err := json.NewDecoder(request.Body).Decode(&authDTO)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, e.InvalidBodyRequest)
		return
	}
	//checking the dto
	if ok := validate(authDTO); !ok {
		utils.ResponseWithError(w, http.StatusBadRequest, e.InvalidBodyRequest)
		return
	}

	userEntity, err := ac.authUseCase.SignUpWithEmailAndPassword(authDTO.Email, authDTO.Password)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, e.FailedToCreateUser)
		return
	}
	utils.ResponseWithJSONData(w, http.StatusOK, userEntity)

}

func validate(authDto *AuthDTO) bool {
	validator := datacheckr.NewValidatorInstance()

	//addvalidation rules
	validator.AddValidationRules(
		datacheckr.EmailValidation,
		datacheckr.MinStrLenValidation(5),
	)

	return validator.Validate(authDto.Email)
}
