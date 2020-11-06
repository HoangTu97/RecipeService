package controller

import (
	"Food/dto"
	"Food/dto/request"
	UserRequest "Food/dto/request/user"
	"Food/dto/response"
	UserResponse "Food/dto/response/user"
	"Food/helpers/e"
	"Food/helpers/jwt"
	"Food/service"

	"github.com/gin-gonic/gin"
)

type User interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type user struct {
	service service.User
}

func NewUser(service service.User) User {
	return &user{service: service}
}

// Register register
// @Summary Register
// @Tags PublicUser
// @Accept  json
// @Param RegisterDTO body requestuser.RegisterDTO true "RegisterDTO"
// @Success 200 {object} response.APIResponseDTO "desc"
// @Router /api/public/user/register [post]
func (r *user) Register(c *gin.Context) {
	var registerDTO UserRequest.RegisterDTO
	errCode := request.BindAndValid(c, &registerDTO)
	if errCode != e.SUCCESS {
		response.CreateErrorResponse(c, e.GetMsg(errCode))
		return
	}

	userDTO := dto.UserDTO{Name: registerDTO.Username, Password: registerDTO.Password}

	_, isSuccess := r.service.Create(userDTO)
	if !isSuccess {
		response.CreateErrorResponse(c, "Register failed!!!")
		return
	}

	response.CreateSuccesResponse(c, nil)
}

// Login login
// @Summary Login
// @Tags PublicUser
// @Accept  json
// @Param LoginDTO body requestuser.LoginDTO true "LoginDTO"
// @Success 200 {object} response.APIResponseDTO "desc"
// @Router /api/public/user/login [post]
func (r *user) Login(c *gin.Context) {
	var loginDTO UserRequest.LoginDTO
	errCode := request.BindAndValid(c, &loginDTO)
	if errCode != e.SUCCESS {
		response.CreateErrorResponse(c, e.GetMsg(errCode))
		return
	}

	userDTO, isSuccess := r.service.FindOneLogin(loginDTO.Username, loginDTO.Password)
	if !isSuccess {
		response.CreateErrorResponse(c, "UNAUTHORIZED")
		return
	}

	tokenString, error := jwt.GenerateToken(userDTO.UserID, userDTO.Name, userDTO.GetRolesStr())
	if error != nil {
		response.CreateErrorResponse(c, "UNAUTHORIZED")
		return
	}

	response.CreateSuccesResponse(c, UserResponse.LoginResponseDTO{
		Token: tokenString,
	})
}
