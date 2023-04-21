package gotine

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
	"time"
)

func TestTaskSubmit(t *testing.T) {
	Convey("test for TaskSubmit", t, func() {
		Convey("submit same func", func() {
			fs := []func(){testFunc1, testFunc1, testFunc1, testFunc1, testFunc1, testFunc1}
			g := new(Gotine)
			g.TaskSubmit(fs)
		})
	})
}

func testFunc1() {
	t := time.Now()
	log.Printf("routine ID is %v \n", GetID())
	time.Sleep(time.Second * 2)
	log.Println(time.Since(t).String())
	log.Printf("routine ID %v done \n", GetID())
}
