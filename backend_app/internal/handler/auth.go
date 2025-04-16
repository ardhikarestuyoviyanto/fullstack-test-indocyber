package handler

import (
	"auth_test/internal/domain"
	"auth_test/internal/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthEchoHandler struct {
	SvcAuth domain.AdapterAuthService
}

func(svc *AuthEchoHandler) LoginHandler(c echo.Context)error{
	dataRaw := echo.Map{}

	if err := c.Bind(&dataRaw); err != nil {
		return err
	}

	email, errEmail := dataRaw["email"].(string)
	password, errPassword := dataRaw["password"].(string)

	if !errEmail || !errPassword {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"message": "Masukkan email dan password",
		})
	}

	res, err := svc.SvcAuth.LoginUsers(email, password)

	if !err{
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"message": "Akun tidak ditemukan",
		})
	}else{

		token, _ := jwt.CreateToken(res.ID, res.Username)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"message": "Akun ditemukan",
			"token": token,
		})
	}

}