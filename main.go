package main

import (
	"fmt"
	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"os"
	configEnv "github.com/joho/godotenv"
	"strconv"
	"strings"
	"tennis_player/model"
)

func main() {

	const DefaultPort = 9000

	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	e := echo.New()
	e.Use(echoMid.Recover())
	e.Use(echoMid.Logger())

	e.PUT("/container-balls", AddBallToContainer)

	listenerPort := fmt.Sprintf(":%d", DefaultPort)
	e.Logger.Fatal(e.Start(listenerPort))

}

func AddBallToContainer(ctx echo.Context) error {
	plyId := ctx.QueryParam("player_id")
	ctrId := ctx.QueryParam("container_id")
	playerId, error := strconv.Atoi(plyId)
	if error != nil {
		log.Print("Cannot convert player id into integer")
		return ctx.JSON(http.StatusBadRequest, model.ResponseDetail{Message: "Player Id should be number"})
	}
	containerId, err := strconv.Atoi(ctrId)
	if err != nil {
		log.Print("Cannot convert container id into integer")
		return ctx.JSON(http.StatusBadRequest, model.ResponseDetail{Message:"Container Id should be number"})
	}

	playerData, stat := UpdatePlayerContainer(containerId, playerId)
	if stat == "full" {
		value := fmt.Sprintf("Players %s can play Tennis", strings.Join(playerData,", "))
		return ctx.JSON(http.StatusOK, model.ResponseDetail{Status: stat, Message: value})
	}
	if stat == "failed" {
		return ctx.JSON(http.StatusBadRequest, model.ResponseDetail{Message:"Cannot add ball into container"})
	}

	return ctx.JSON(http.StatusOK, model.ResponseDetail{Status: stat, Message:"Success add the ball into container"})

}