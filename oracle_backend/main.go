package main

import (
	"Oracle/dao"
	"Oracle/logger"
	"Oracle/route"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.LogInit()
	dao.Init()
	r := route.GetRoutes()
	err := r.Run(":8080")
	if err != nil {
		logrus.Error(err)
	}

}
