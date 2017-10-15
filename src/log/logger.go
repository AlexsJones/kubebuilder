package logger

import (
	"fmt"
	"log"
	"sync"

	shortid "github.com/ventu-io/go-shortid"
)

//Logger ...
type Logger struct {
	UUID string
}

func (l *Logger) prefix() string {
	return fmt.Sprintf("[%s]", l.UUID)
}

//Log ...
func (l *Logger) Log(args ...string) {
	log.Printf(fmt.Sprintf("%s:%s", l.prefix(), args))
}

//Fatal ...
func (l *Logger) Fatal(args ...string) {
	log.Fatalf("%s:%s", l.prefix(), args)
}

var instance *Logger
var once sync.Once

//GetInstance of logger singleton
func GetInstance() *Logger {
	once.Do(func() {
		sid, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			panic(err)
		}
		instance = &Logger{UUID: sid.MustGenerate()}
	})
	return instance
}
