package user

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mtoku/di/app/com/domain/app/user"
	"github.com/mtoku/di/app/com/usecase/input_data"
	"github.com/mtoku/di/app/com/usecase/output_data"
	log "github.com/sirupsen/logrus"
)

func NewUserController(createService user.IUserCreateService) UserController {
	return UserController{
		UserCreateService: createService,
	}
}

type UserController struct {
	UserCreateService user.IUserCreateService
}

func (controller UserController) Create(w http.ResponseWriter, r *http.Request) {

	userCreateResult := output_data.UserCreateResult{
		Status:  "0",
		Message: "Success",
	}
	defer func() {
		json, _ := userCreateResult.MarshalJSON()
		fmt.Fprint(w, string(json))
	}()

	req, err := controller.newUserCreateRequest(r)
	if err != nil {
		log.Error(err)
		userCreateResult.Status = "1"
		userCreateResult.Message = "Json Marshall Error"
		return
	}

	res := controller.UserCreateService.Handle(req)
	if res.Error != nil {
		log.Error(res.Error)
		userCreateResult.Status = "2"
		userCreateResult.Message = "Error"
	}

	return
}

func (controller UserController) newUserCreateRequest(r *http.Request) (req user.UserCreateRequest, err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	apiReq := &input_data.APIUserCreateRequest{}
	err = apiReq.UnmarshalJSON(body)
	if err != nil {
		return
	}

	req.UserID = apiReq.UserID
	req.Password = apiReq.Password
	req.Nickname = apiReq.Nickname

	return
}
