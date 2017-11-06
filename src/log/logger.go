package logger

import (
	"fmt"
	"log"
	"sync"

	"github.com/fatih/color"
	shortid "github.com/ventu-io/go-shortid"
)

//Logger ...
type Logger struct {
	UUID      string
	mpColor   func(a ...interface{}) string
	infoColor func(a ...interface{}) string
}

func (l *Logger) prefix() string {
	return fmt.Sprintf("[%s]", l.mpColor(l.UUID))
}

//Info for program alerts more than anything else
func (l *Logger) Info(args ...string) {
	log.Printf(fmt.Sprintf("%s:%s", l.prefix(), l.infoColor(args)))
}

//Log ...
func (l *Logger) Log(args ...string) {
	log.Printf(fmt.Sprintf("%s:%s", l.prefix(), args))
}

//Warn ...
func (l *Logger) Warn(args ...string) {
	red := color.New(color.FgRed).SprintFunc()
	log.Printf(fmt.Sprintf("%s:%s", l.prefix(), red(args)))
}

//Fatal ...
func (l *Logger) Fatal(args ...string) {
	red := color.New(color.FgRed).SprintFunc()
	log.Fatalf("%s:%s", l.prefix(), red(args))
}

var instance *Logger
var once sync.Once

//GetInstance of logger singleton
func GetInstance() *Logger {
	once.Do(func() {
		green := color.New(color.FgGreen).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()
		sid, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			panic(err)
		}
		instance = &Logger{UUID: sid.MustGenerate(), mpColor: green, infoColor: blue}
	})
	return instance
}
