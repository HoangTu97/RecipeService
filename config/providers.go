package config

import (
  "p2/controller"
  "p2/pkg/service/Auth"
  "p2/pkg/service/Database"
  "p2/pkg/service/File"
  "p2/pkg/service/Cache"
  "p2/pkg/service/Mail"
  "p2/pkg/service/Log"
  "p2/pkg/service/Jwt"
  "p2/pkg/service/Hash"
  // "p2/pkg/service/Schedule"
  "p2/repository/impl"
  "p2/repository/proxy"
  "p2/service/impl"
  "p2/service/proxy"
  "p2/service/mapper/impl"
)

func Providers(
  dbManager Database.Manager, 
  jwtManager Jwt.Manager,
  cacheManager Cache.Manager,
  mailManager Mail.Manager,
  logManager Log.Manager,
) []controller.Base {
  db := dbManager.Connection("")

  // Mappers declare
  userMapper := mapper_impl.NewUser()
  linkMapper := mapper_impl.NewLink()
  clickMapper := mapper_impl.NewClick()
  // Mappers declare end : dont remove

  // Repositories declare
  userRepo := repository_impl.NewUser(db)
  linkRepo := repository_impl.NewLink(db)
  clickRepo := repository_impl.NewClick(db)
  // Repositories declare end : dont remove

  // Proxy Repositories declare
  userRepoProxy := repository_proxy.NewUser(userRepo)
  // Proxy Repositories declare end : dont remove

  // Services declare
  cacheService := cacheManager.Driver("")
  fileService := File.NewService()
  // mailService := mailManager.Mailer("smtp")
  hashService := Hash.NewService("")
  // scheduleService := Schedule.NewService()
  authService := Auth.NewService(jwtManager)
  userService := service_impl.NewUser(userRepoProxy, userMapper, hashService)
  linkService := service_impl.NewLink(linkRepo, linkMapper)
  clickService := service_impl.NewClick(clickRepo, clickMapper)
  // Services declare end : dont remove

  // Proxy Services declare
  userServiceProxy := service_proxy.NewUser(userService, cacheService)
  // Proxy Services declare end : dont remove

  // Controllers declare
  fileController := controller.NewFile(fileService)
  authController := controller.NewAuth(authService, userServiceProxy)
  userController := controller.NewUser(userServiceProxy)
  urlShortenerController := controller.NewUrlShortener(linkService, clickService, userService, authService, hashService)
  // Controllers declare end : dont remove

  return []controller.Base{
    // Register controller declare
    fileController,
    authController,
    userController,
    urlShortenerController,
    // Register controller declare end : dont remove
  }
}
