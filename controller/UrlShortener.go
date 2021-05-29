package controller

import (
  "p2/dto"
  "p2/dto/request"
  LinkRequest "p2/dto/request/link"
  "p2/dto/response"
  LinkResponse "p2/dto/response/link"
  "p2/helpers/constants"
  // "p2/pkg/converter"
  AuthService "p2/pkg/service/Auth"
  "p2/pkg/service/Hash"
  "p2/service"

  "net/http"

  "github.com/gin-gonic/gin"
)

type UrlShortener interface {
  GetRoutes() []RouteController
  Create(c *gin.Context)
  Click(c *gin.Context)
}

type urlShortener struct {
  linkService service.Link
  clickService service.Click
  userService service.User
  authService AuthService.Service
  hashService Hash.Service
}

func NewUrlShortener(linkService service.Link, clickService service.Click, userService service.User, authService AuthService.Service, hashService Hash.Service) UrlShortener {
  return &urlShortener{
    linkService: linkService, clickService: clickService, userService: userService,
    authService: authService, hashService: hashService,
  }
}

func (r *urlShortener) GetRoutes() []RouteController {
  return []RouteController{
    {Method:http.MethodPost,Path:"/api/private/url",Handler:r.Create},
    {Method:http.MethodGet,Path:"/api/public/url/:id",Handler:r.Click},
  }
}

// Url create
// @Summary Create
// @Tags PrivateCreate
// @Accept json
// @Security ApiKeyAuth
// @Param body body LinkRequest.CreateRequestDTO true "body"
// @Success 200 {object} response.APIResponseDTO{data=LinkResponse.CreateResponseDTO} "desc"
// @Router /api/private/url [post]
func (r *urlShortener) Create(c *gin.Context) {
  var success bool

  var requestDTO LinkRequest.CreateRequestDTO
  err := request.BindAndValid(c, &requestDTO)
  if err != nil {
    response.CreateErrorResponse(c, err.Error())
    return
  }

  userID := r.authService.GetUserID(c)
  userDTO, success := r.userService.FindOneByUserID(userID)
  if !success {
    response.CreateErrorResponse(c, constants.ErrorStringApi.UNAUTHORIZED_ACCESS)
    return
  }

  linkDTO := dto.LinkDTO{
    LongUrl: requestDTO.Url,
    LongUrlHash: r.hashService.Make(requestDTO.Url),
    ClickNum: 0,
    Ip: c.ClientIP(),
    IsDisabled: false,
    IsCustom: false,
    IsApi: true,
    CreatorID: userDTO.ID,
  }

  linkDTO, success = r.linkService.Save(linkDTO)
  if !success {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INTERNAL_ERROR)
    return
  }

  linkDTO.ShortUrl = r.linkService.IdToShortUrl(linkDTO.ID)
  linkDTO, success = r.linkService.Save(linkDTO)
  if !success {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INTERNAL_ERROR)
    return
  }

  response.CreateSuccesResponse(c, LinkResponse.CreateResponseDTO{Url:linkDTO.ShortUrl})
}



// Url click
// @Summary Click
// @Tags PublicClick
// @Success 200 {object} response.APIResponseDTO{data=LinkResponse.CreateResponseDTO} "desc"
// @Router /api/public/url/:shorturl [post]
func (r *urlShortener) Click(c *gin.Context) {
  var success bool

  short_url := c.Param("id")

  linkID := r.linkService.ShortUrlToId(short_url)
  var linkDTO dto.LinkDTO
  linkDTO, success = r.linkService.FindOne(linkID)
  if !success {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INTERNAL_ERROR)
    return
  }

  linkDTO.ClickNum = linkDTO.ClickNum + 1

  clickDTO := dto.ClickDTO{
    LinkID: linkDTO.ID,
    Ip: c.ClientIP(),
    Country: "",
    Referer: c.Request.Referer(),
    RefererHost: c.Request.Host,
    UserAgent: c.Request.UserAgent(),
  }

  _, success = r.clickService.Save(clickDTO)
  if !success {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INTERNAL_ERROR)
    return
  }

  linkDTO, success = r.linkService.Save(linkDTO)
  if !success {
    response.CreateErrorResponse(c, constants.ErrorStringApi.INTERNAL_ERROR)
    return
  }

  response.CreateSuccesResponse(c, LinkResponse.CreateResponseDTO{Url:linkDTO.LongUrl})
}
