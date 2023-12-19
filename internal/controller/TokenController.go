package internal

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	model "learning-go/internal/model"
	service "learning-go/internal/service"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	GrantType string `json:"grant_type"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Response struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expire_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId uint
}

const (
	tokenTTL        = 24 * time.Hour
	refreshTokenTTL = 24 * 7 * time.Hour
)

func TokenController(c *gin.Context) {
	var client model.Client
	var access model.AccessToken
	var refresh model.RefreshToken
	var request Request
	var user model.User

	db := service.CreateConnection()

	requestData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	requestString := string(requestData)
	arrayRequest, _ := url.ParseQuery(requestString)
	request.GrantType = arrayRequest.Get("grant_type")
	request.Username = arrayRequest.Get("username")
	request.Password = arrayRequest.Get("password")

	if request.GrantType != "password" {
		http.Error(c.Writer, "unsupported grant type", http.StatusBadRequest)
	}

	err = db.Model(&model.User{Username: request.Username, Password: request.Password}).First(&user).Error

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusForbidden)
		return
	}

	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		http.Error(c.Writer, err.Error(), 400)
	}

	authArray := strings.Split(auth, "Basic ")
	rawDecodedText, err := base64.StdEncoding.DecodeString(authArray[1])
	if err != nil {
		http.Error(c.Writer, err.Error(), 400)
	}
	clientCred := strings.Split(string(rawDecodedText), ":")
	clientId := clientCred[0]
	clientSecret := clientCred[1]

	err = db.Model(&model.Client{
		Identifier: clientId,
		Secret:     clientSecret,
	}).First(&client).Error

	if err != nil || !client.Active {
		http.Error(c.Writer, err.Error(), http.StatusForbidden)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, TokenClaims{
		jwt.StandardClaims{
			Id:        strconv.Itoa(int(user.ID)),
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	access.Identifier, err = token.SigningString()

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	access.Expiry = time.Now().Add(tokenTTL)
	access.ClientId = clientId
	access.UserIdentifier = user.Username

	db.Create(&access)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, TokenClaims{
		jwt.StandardClaims{
			Id:        strconv.Itoa(int(user.ID)),
			ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	refresh.AccessToken = access.Identifier
	refresh.Identifier, err = refreshToken.SigningString()

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	refresh.Revoked = false
	refresh.Expiry = time.Now().Add(refreshTokenTTL)

	db.Create(&refresh)

	var response Response

	response.TokenType = "Bearer"
	response.ExpiresIn = access.Expiry.Unix()
	response.AccessToken = access.Identifier
	response.RefreshToken = refresh.Identifier

	c.JSON(200, response)
}
