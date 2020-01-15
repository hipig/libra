package module

import (
	"github.com/lhlyu/libra/common"
	"github.com/lhlyu/yutil"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

type lg struct {
}

func (lg) seq() int {
	return 1 << 0
}

func (lg) SetUp() {
	common.L = NewEntry()
}

// 日志模块
var LgModule = lg{}

// 自定义一个logrus  hook
type defaultFieldHook struct {
}

func (hook *defaultFieldHook) Fire(entry *logrus.Entry) error {
	// todo
	return nil
}

// 只针对debug级别的日志
func (hook *defaultFieldHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.DebugLevel}
}

func NewEntry() *logrus.Entry {
	lr := logrus.New()
	out := common.Cfg.GetString("log.out")
	level := common.Cfg.GetString("log.level")
	if out != "" {
		dir := path.Dir(out)
		exists := yutil.FileIsExists(dir)
		if !exists {
			os.MkdirAll(dir, 0666)
		}
		f, err := os.OpenFile(out, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
			return nil
		}
		lr.SetOutput(f)
		lr.SetFormatter(new(logrus.TextFormatter))
	}
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		lv = logrus.InfoLevel
	}
	lr.SetLevel(lv)

	// 这里给日志加 hook
	lr.AddHook(new(defaultFieldHook))

	entry := logrus.NewEntry(lr)
	return entry
}
