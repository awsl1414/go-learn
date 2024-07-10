/*
 * @Author: awsl1414 3030994569@qq.com
 * @Date: 2024-07-09 16:40:39
 * @LastEditors: awsl1414 3030994569@qq.com
 * @LastEditTime: 2024-07-10 16:14:17
 * @FilePath: /go-learn/gin/project/ranking/pkg/logger/logger.go
 * @Description:
 *
 */
package logger

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	LOG_DIR string = "./runtime/log"
	mu      sync.Mutex
)

func init() {
	// 设置日志格式为json
	logrus.SetFormatter(&logrus.JSONFormatter{

		// 作为时间格式的参考值，因此这个格式字符串表示日期和时间的格式。
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}

func Write(msg string, filename string) {
	setOutputFile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}
func Debug(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args)

}
func Info(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.InfoLevel, "info")
	logrus.WithFields(fields).Info(args)
}
func Warn(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.WarnLevel, "warn")
	logrus.WithFields(fields).Warn(args)
}
func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.FatalLevel, "fatal")
	logrus.WithFields(fields).Fatal(args)
}
func Error(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.ErrorLevel, "error")
	logrus.WithFields(fields).Error(args)
}
func Panic(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.PanicLevel, "panic")
	logrus.WithFields(fields).Panic(args)
}
func Trace(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.TraceLevel, "trace")
	logrus.WithFields(fields).Trace(args)
}

func checkLogDir() {
	if _, err := os.Stat(LOG_DIR); os.IsNotExist(err) {
		err = os.MkdirAll(LOG_DIR, 0777)
		if err != nil {
			panic(fmt.Errorf("creat log dir '%s' err: %s", LOG_DIR, err))
		}
	}
}

func createLogFile(link string, logName string) *os.File {
	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join(LOG_DIR, logName+link+timeStr+".log")

	out, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("open log file err", err)

	}

	return out

}
func setOutputFile(level logrus.Level, logName string) {
	// 加互斥锁保证线程安全 setOutputFile和logrus.SetOutput线程不安全
	mu.Lock()
	defer mu.Unlock()

	checkLogDir()

	out := createLogFile("-", logName)

	logrus.SetOutput(out)
	logrus.SetLevel(level)

}

func LoggerToFile() gin.LoggerConfig {

	checkLogDir()

	out := createLogFile("success_", "")

	var conf = gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - %s \"%s %s %s %d %s \"%s\" %s\"\n",
				param.TimeStamp.Format("2006-01-02 15:04:05"),
				param.ClientIP,
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
		Output: io.MultiWriter(os.Stdout, out), // 同时输出日志到文件和控制台
	}
	return conf

}

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			checkLogDir()
			f := createLogFile("error_", "")
			timeFileStr := time.Now().Format("2006-01-02 15:04:05")
			f.WriteString("panic err time:" + timeFileStr + "\n")
			f.WriteString(fmt.Sprintf("%v", err) + "\n")
			f.WriteString("stacktrace from panic:" + string(debug.Stack()) + "\n")
			f.Close()
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("%v", err),
			})

			// 终止后续接口调用，不加的话，recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}

	}()

	c.Next()
}
