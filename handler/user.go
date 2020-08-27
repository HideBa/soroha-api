package handler

import (
	"net/http"
	"time"

	"github.com/HideBa/soroha-api/config"
	"github.com/HideBa/soroha-api/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) SignUp(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err.Error}
	}
	if user.Username == "" || user.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}
	hash, hashError := h.hashPassword(user.Password)
	if hashError != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: hashError}
	}
	user.Password = hash
	if err := h.DB.Create(&user).Error; err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err}
	}

	return c.JSON(http.StatusCreated, map[string]model.UserResponse{"user": user.UserTransformer()})
}

func (h *Handler) Login(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err}
	}
	record, err := h.findUserByName(user.Username)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	if err := h.checkPassword(record.Password, user.Password); err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	token := jwt.New(jwt.SigningMethodES256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	user.Token, err = token.SignedString([]byte(config.GetConfig().Server.KEY))
	if err != nil {
		return &echo.HTTPError{Code: http.StatusInternalServerError, Message: err}
	}

	cookies := &http.Cookie{}
	cookies.Name = "JWTcookie"
	cookies.Value = user.Token
	cookies.Expires = time.Now().Add(time.Hour * 48)
	c.SetCookie(cookies)
	return c.JSON(http.StatusOK, record.UserTransformer())

}

func (h *Handler) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (h *Handler) validatePassword(hash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
