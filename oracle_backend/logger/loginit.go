package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

func LogInit() {
	log.Println("log init ...")
	logrus.SetReportCaller(true) //设置在输出日志中添加文件名和方法信息
	path, err := os.Getwd()
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile(fmt.Sprintf("%s/logger/log.txt", path), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))

	log.Println("log init success !")
}
