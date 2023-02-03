package metion

import (
	"bizd/metion/global"
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

type MyClaim struct {
	UserId         string
	StandardClaims jwt.StandardClaims
}

func (m MyClaim) Valid() error {
	//TODO implement me
	return m.StandardClaims.Valid()
}

//创建token

func CreateToken(userId string) (s string, err error) {

	// Create the Claims
	claims := MyClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,                 //生效时间，这里是一分钟前生效
			ExpiresAt: time.Now().Unix() + global.ExpiresTime, //过期时间，这里是一小时过期
			Issuer:    global.Issuer,                          //签发人
		},
	}
	//SigningMethodHS256,HS256对称加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//通过自定义令牌加密
	ss, err := token.SignedString(global.MySigningKey)
	if err != nil {
		fmt.Println("生成token失败")
	}
	return ss, err
}

func ParseToken(token string) (*MyClaim, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return global.MySigningKey, nil
	})
	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*MyClaim); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func CheckToken(token string) bool {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return global.MySigningKey, nil
	})
	if tokenClaims != nil && err == nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if tokenClaims.Valid {
			return true
		}
	}
	return false
}
