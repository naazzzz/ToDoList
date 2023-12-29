package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io/ioutil"
	"learning-go/config"
	model "learning-go/internal/model"
	routes "learning-go/internal/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Init()

	reqBody := model.UserDTO{
		Username: "naaaddzzzzz",
		Password: "password1",
	}

	ts := httptest.NewServer(routes.Route())

	defer ts.Close()

	respBody, resp, err := makeUserRequest(ts, reqBody)

	if err != nil {
		panic(err)
		return
	}

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Equal(t, respBody, reqBody)

}

func makeUserRequest(ts *httptest.Server, reqBody interface{}) (model.UserDTO, *http.Response, error) {
	reqBodyJson, _ := json.Marshal(reqBody)

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/users", ts.URL),
		bytes.NewReader(reqBodyJson),
	)

	if err != nil {
		return model.UserDTO{}, nil, err
	}

	resp, err := config.GetHttpClient().Do(req)

	if err != nil {
		return model.UserDTO{}, resp, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.UserDTO{}, nil, err
	}

	var user model.UserDTO
	err = json.Unmarshal(respBody, &user)
	if err != nil {
		return model.UserDTO{}, nil, err
	}

	return user, resp, nil
}
