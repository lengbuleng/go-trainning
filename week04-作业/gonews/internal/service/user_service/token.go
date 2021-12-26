package user_service

import (
	"app/internal/models"
	"app/pkg/config"
	"app/pkg/e"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenSvc struct {
}

func (*TokenSvc) MakeToken(user *models.User) (string, error) {

	tk := jwt.New(jwt.SigningMethodHS256)
	claims := tk.Claims.(jwt.MapClaims)
	claims["userId"] = user.Id
	claims["role"] = user.Role
	claims["expDate"] = time.Now().Add(time.Hour * 500).Format("2006-01-02 15:04:05")

	token, err := tk.SignedString([]byte(config.GetConfig().Token_KEY))
	if err != nil {
		return "", err
	}
	return token, nil

}

func (*TokenSvc) ParseToken(token string) (*models.User, error) {

	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().Token_KEY), nil
	})
	if err != nil {
		return nil, e.NewError("token解析失败")
	}

	if claims, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		userId := claims["userId"]
		role := claims["role"]
		expDate := claims["expDate"]
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		if timeNow > expDate.(string) {
			return nil, e.NewError("token已过期")
		}
		if userId == "" {
			return nil, e.NewError("token解密，id为空")
		}
		user := &models.User{}
		user.Role = int(role.(float64))
		user.Id, _ = primitive.ObjectIDFromHex(userId.(string))
		return user, nil
	} else {
		return nil, e.NewError("token claims 解析失败")
	}

}
