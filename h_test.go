package gone_test

import (
	"github.com/gone-io/gone"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Triangle struct {
	gone.Flag
	a gone.XPoint `gone:"point-a"`
	b gone.Point  `gone:"point-b"`
	c *gone.Point `gone:"point-c"`
}

func TestRun(t *testing.T) {
	gone.Run(func(cemetery gone.Cemetery) error {
		cemetery.
			Bury(&Triangle{}).
			Bury(&gone.Point{Index: 1}, gone.GonerId("point-a")).
			Bury(&gone.Point{Index: 2}, gone.GonerId("point-b")).
			Bury(&gone.Point{Index: 3}, gone.GonerId("point-c"))

		return nil
	})
}

func TestNew(t *testing.T) {
	var sort []int

	const gonerId gone.GonerId = "Triangle"

	heaven := gone.
		New(func(cemetery gone.Cemetery) error {
			cemetery.
				Bury(&Triangle{}, gonerId).
				Bury(&gone.Point{Index: 1}, gone.GonerId("point-a")).
				Bury(&gone.Point{Index: 2}, gone.GonerId("point-b")).
				Bury(&gone.Point{Index: 3}, gone.GonerId("point-c"))

			return nil
		})

	heaven.
		BeforeStart(func(cemetery gone.Cemetery) error {
			tomb := cemetery.GetTomById(gonerId)
			triangle, ok := tomb.GetGoner().(*Triangle)
			assert.True(t, ok)
			assert.Equal(t, triangle.a.GetIndex(), 1)
			assert.Equal(t, triangle.b.GetIndex(), 2)
			assert.Equal(t, triangle.c.Index, 3)

			sort = append(sort, 0)
			return nil
		}).
		AfterStart(func(cemetery gone.Cemetery) error {
			sort = append(sort, 1)
			return nil
		}).
		BeforeStop(func(cemetery gone.Cemetery) error {
			sort = append(sort, 2)
			return nil
		}).
		AfterStop(func(cemetery gone.Cemetery) error {
			sort = append(sort, 3)
			return nil
		})

	heaven.Install()
	heaven.Start()
	heaven.Stop()

	for i := range sort {
		assert.Equal(t, i, sort[i])
	}
}

func TestServe(t *testing.T) {
	gone.AfterStopSignalWaitSecond = 1
	gone.Serve(func(cemetery gone.Cemetery) error {
		go func() {
			tom := cemetery.GetTomById(gone.IdGoneHeaven)
			heaven := tom.GetGoner().(gone.Heaven)
			heaven.End()
		}()
		return nil
	})
}
