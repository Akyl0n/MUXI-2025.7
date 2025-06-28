package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "123456"

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// 生成JWT token
func GenerateToken(id int) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := Claims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretkey))
}

// 登录接口
func Login(c *gin.Context) {
	// 定义接收登录参数的结构体
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误", "error": err.Error()})
		return
	}

	// 假设用户名/密码为 admin / 123456
	if loginData.Username != "admin" || loginData.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "用户名或密码错误"})
		return
	}

	// 假设用户ID是 1，生成 token
	token, err := GenerateToken(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "生成Token失败", "error": err.Error()})
		return
	}

	// 返回 token
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
	})
}

// JWT中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未认证"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretkey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token", "error": err.Error()})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未认证"})
			c.Abort()
			return
		}
		userId, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "无法获取id"})
			c.Abort()
			return
		}
		fmt.Println(userId)
		c.Set("user_id", int(userId))
		c.Next()
	}
}

// 受保护接口
func Protected(c *gin.Context) {
	userId := c.GetInt("user_id")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("欢迎，用户ID：%d", userId)})
}

func main() {
	r := gin.Default()
	r.POST("/login", Login)
	r.GET("/protected", AuthMiddleware(), Protected)
	r.Run(":8080")
}
