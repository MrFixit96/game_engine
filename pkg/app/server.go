package app

import (
    "github.com/gin-gonic/gin"
    "log"
    "game_engine/pkg/api"
)

type Server struct {
    router        *gin.Engine
    userService   api.UserService
    gameService api.GameService
}

func NewServer(router *gin.Engine, userService api.UserService, gameService api.GameService) *Server {
    return &Server{
        router:        router,
        userService:   userService,
        gameService: gameService,
    }
}

func (s *Server) Run() error {
    // run function that initializes the routes
    r := s.Routes()

    // run the server through the router
    err := r.Run()

    if err != nil {
        log.Printf("Server - there was an error calling Run on router: %v", err)
        return err
    }

    return nil
}