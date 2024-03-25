package controller

import (
	"Oracle/dao"
	"Oracle/middleware"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	var requestBody struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		logrus.Error(err)
	}

	err = dao.Register(requestBody.Name, requestBody.Email, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "注册失败,err:" + err.Error(),
		})
		return
	}

	token, err := middleware.GenerateToken(requestBody.Email)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": token,
		"msg":  "注册成功",
	})

}

func LoginHandler(c *gin.Context) {
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		logrus.Error(err)
	}

	err = dao.Login(requestBody.Email, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": err.Error(),
			"msg":  "登陆失败",
		})
		return
	}

	token, err := middleware.GenerateToken(requestBody.Email)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": token,
		"msg":  "登陆成功",
	})
}
