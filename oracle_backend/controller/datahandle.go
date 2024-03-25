package controller

import (
	"Oracle/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GradeDataHandle(c *gin.Context) {
	_, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "未登录",
		})
		return
	}

	var requestBody struct {
		NodeId          int64   `json:"node_id"`
		DataImportance  float32 `json:"data_importance"`
		NodeFailureRate float32 `json:"node_failure_rate"`
		Source          float32 `json:"source"`
		Integrity       float32 `json:"integrity"`
		Utilization     float32 `json:"utilization"`
		Result          int64   `json:"result"`
	}

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		logrus.Error(err)
	}

	res, err := dao.GradeData(requestBody.NodeId, requestBody.DataImportance, requestBody.NodeFailureRate, requestBody.Source, requestBody.Integrity, requestBody.Utilization, requestBody.Result)
	if res == true {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
			"msg":  "评估节点成功",
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"code": 1,
		"data": err.Error(),
		"msg":  "评估节点失败",
	})

}

func FindDataByIdHandle(c *gin.Context) {
	_, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "未登录",
		})
		return
	}

	var requestBody struct {
		DataSourceId int `json:"data_source_id"`
	}

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		logrus.Error(err)
	}

	data := dao.FindDataById(requestBody.DataSourceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "获取数据,err:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
		"msg":  "获取数据成功",
	})

}

func GetStartTimeHandle(c *gin.Context) {
	_, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "未登录",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "初始时间成功",
	})

}

func UseDataHandle(c *gin.Context) {
	_, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "未登录",
		})
		return
	}

	var requestBody struct {
		DataSourceId int    `json:"data_source_id"`
		DurationTime string `json:"duration_time"`
	}

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		logrus.Error(err)
	}

	err = dao.UseData(requestBody.DataSourceId, requestBody.DurationTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "使用预言机失败，err：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "使用预言机成功",
	})

}

func Feedback(c *gin.Context) {
	_, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "未登录",
		})
		return
	}

	var requestBody struct {
		DataSourceId int `json:"data_source_id"`
	}

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		logrus.Error(err)
	}

	err = dao.FeedBack(requestBody.DataSourceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"data": nil,
			"msg":  "反馈失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": nil,
		"msg":  "反馈成功",
	})

}

func CorsMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
