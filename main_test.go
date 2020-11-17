package main

import (
	"fmt"
	configEnv "github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)


func TestAddBall(t *testing.T) {

	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	e := echo.New()
	q := make(url.Values)
	q.Set("player_id", "2")
	q.Set("container_id", "5")
	req := httptest.NewRequest(http.MethodPut, "/container-balls?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, AddBallToContainer(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestAddBallIfPlayerIdIsChar(t *testing.T) {

	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	e := echo.New()
	q := make(url.Values)
	q.Set("player_id", "abc")
	q.Set("container_id", "5")
	req := httptest.NewRequest(http.MethodPut, "/container-balls?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, AddBallToContainer(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestAddBallIfContainerIdIsChar(t *testing.T) {

	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	e := echo.New()
	q := make(url.Values)
	q.Set("player_id", "1")
	q.Set("container_id", "abc")
	req := httptest.NewRequest(http.MethodPut, "/container-balls?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, AddBallToContainer(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}