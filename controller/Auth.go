package controller

import (
  "p2/dto"
  "p2/dto/request"
  AuthRequest "p2/dto/request/auth"
  "p2/dto/response"
  AuthResponse "p2/dto/response/auth"
  "p2/helpers/constants"
  AuthService "p2/pkg/service/Auth"
  "p2/service"

  "net/http"

  "github.com/gin-gonic/gin"
)

type Auth interface {
  GetRoutes() []RouteController
  Register(c *gin.Context)
  Login(c *gin.Context)
  Logout(c *gin.Context)
  UserInfo(c *gin.Context)
}

type auth struct {
  service AuthService.Service
  userService service.User
}

func NewAuth(service AuthService.Service, userService service.User) Auth {
  return &auth{service: service, userService: userService}
}

func (r *auth) GetRoutes() []RouteController {
  return []RouteController{
    {Method:http.MethodPost,Path:"/api/public/auth/register",Handler:r.Register},
    {Method:http.MethodPost,Path:"/api/public/auth/login",Handler:r.Login},
    {Method:http.MethodGet,Path:"/api/private/auth/userinfo",Handler:r.UserInfo},
  }
}

// Register register
// @Summary Register
// @Tags PublicAuth
// @Accept  json
// @Param RegisterDTO body AuthRequest.RegisterDTO true "RegisterDTO"
// @Success 200 {object} response.APIResponseDTO "desc"
// @Router /api/public/auth/register [post]
func (r *auth) Register(c *gin.Context) {
  var registerDTO AuthRequest.RegisterDTO
  err := request.BindAndValid(c, &registerDTO)
  if err != nil {
    response.CreateErrorResponse(c, err.Error())
    return
  }

  _, checkRegistered := r.userService.FindOneByUsername(registerDTO.Username)
  if checkRegistered {
    response.CreateErrorResponse(c, constants.ErrorStringApi.USER_EXISTED)
    return
  }

  userDTO := dto.UserDTO{
    Name: registerDTO.Username,
    Password: registerDTO.Password,
  }

  userDTO, isSuccess := r.userService.Create(userDTO)
  if !isSuccess {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INTERNAL_ERROR)
    return
  }

  tokenString, error := r.service.GenerateToken(userDTO)
  if error != nil {
    response.CreateErrorResponse(c, constants.ErrorStringApi.USER_TOKEN_GEN_FAILED)
    return
  }

  response.CreateSuccesResponse(c, AuthResponse.RegisterDTO{
    Token: tokenString,
  })
}

// Login login
// @Summary Login
// @Tags PublicAuth
// @Accept  json
// @Param LoginDTO body AuthRequest.LoginDTO true "LoginDTO"
// @Success 200 {object} response.APIResponseDTO "desc"
// @Router /api/public/auth/login [post]
func (r *auth) Login(c *gin.Context) {
  var loginDTO AuthRequest.LoginDTO
  err := request.BindAndValid(c, &loginDTO)
  if err != nil {
    response.CreateErrorResponse(c, err.Error())
    return
  }

  userDTO, isSuccess := r.userService.FindOneLogin(loginDTO.Username, loginDTO.Password)
  if !isSuccess {
    response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
    return
  }

  tokenString, err := r.service.GenerateToken(userDTO)
  if err != nil {
    response.CreateErrorResponse(c, constants.ErrorStringApi.USER_TOKEN_GEN_FAILED)
    return
  }

  response.CreateSuccesResponse(c, AuthResponse.LoginDTO{
    Token: tokenString,
  })
}

func (r *auth) Logout(c *gin.Context) {

}

func (r *auth) UserInfo(c *gin.Context) {

}
