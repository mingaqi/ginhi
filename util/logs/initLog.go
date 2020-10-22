package logs

import (
	"bytes"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var logOnce sync.Once

func InitLogs(path string) {

	//logFile, _ := os.OpenFile(path,os.O_CREATE|os.O_RDWR|os.O_APPEND,0644)
	logOnce.Do(func() {
		// 日志级别
		logrus.SetLevel(logrus.DebugLevel)
		// 使用自定义格式
		logrus.SetFormatter(new(LogFormatter))
		// 开启行号
		logrus.SetReportCaller(true)

		/*
			    日志轮转相关函数
			    `WithLinkName` 为最新的日志建立软连接
			    `WithRotationTime` 设置日志分割的时间，隔多久分割一次
			    `WithMaxAge` 和 `WithRotationCount` 二者只能设置一个
				`WithMaxAge` 设置文件清理前的最长保存时间
			    `WithRotationCount` 设置文件清理前最多保存的个数
				`WithRotationSize` 设置文件大小切分单位byte
		*/
		writer, _ := rotatelogs.New(
			fmt.Sprintf(path, "."+"%Y%m%d"),
			rotatelogs.WithLinkName(fmt.Sprintf(path, "")),
			rotatelogs.WithMaxAge(15*24*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
			rotatelogs.WithRotationSize(200*1000*1000),
		)
		logrus.SetOutput(io.MultiWriter(os.Stdout, writer)) // 控制台和文件打印

	})
}

//日志自定义格式
type LogFormatter struct{}

func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("[2006-01-02|15:04:05.000]")
	var file string
	var l int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		l = entry.Caller.Line
	}
	msg := fmt.Sprintf("[%s] [%s:%d][GOID:%d][%s] %s\n", timestamp, file, l, getGID(), strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil

}

// 获取当前协程ID
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	fmt.Println(string(b))
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n

}
