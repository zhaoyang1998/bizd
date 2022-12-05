package metion

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type StandardClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

type ResponseMsg struct {
	Code int `json:"code"`
	Data struct {
		Token           string `json:"token"`
		AdministratorID int    `json:"administrator_id"`
		ExpiredAt       string `json:"expired_at"`
		CreatedAt       string `json:"created_at"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

var mySigningKey = []byte("Key of BIZD")

type MyClaim struct {
	Username       interface{}
	Id             int
	StandardClaims jwt.StandardClaims
}

func (m MyClaim) Valid() error {
	//TODO implement me
	panic("生产token的方法不对")
}

//创建token
func CreateToken(userid int, username interface{}) (s string, err error) {

	// Create the Claims
	claims := MyClaim{
		Username: username,
		Id:       userid,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,    //生效时间，这里是一分钟前生效
			ExpiresAt: time.Now().Unix() + 60*60, //过期时间，这里是一小时过期
			Issuer:    "BIZD",                    //签发人
		},
	}
	//SigningMethodHS256,HS256对称加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//通过自定义令牌加密
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("生成token失败")
	}
	return ss, err
}
