package controller

import (
  // "p2/dto"
  // "p2/dto/request"
  // UserRequest "p2/dto/request/user"
  // "p2/dto/response"
  // UserResponse "p2/dto/response/user"
  // "p2/helpers/constants"
  "p2/service"

  // "github.com/gin-gonic/gin"
)

type User interface {
  GetRoutes() []RouteController
}

type user struct {
  service service.User
}

func NewUser(service service.User) User {
  return &user{service: service}
}

func (r *user) GetRoutes() []RouteController {
  return []RouteController{}
}
