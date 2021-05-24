package prompt

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

type Question interface {
	Init()
	Repr()
	HandleInput(input []byte) (bool, error)
}

type Questions struct {
	Q []Question
}

func (q Questions) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	im := newInputManager()
	Writer.Cursor.Hide()
QUESTIONS:
	for _, i := range q.Q {
		i.Init()
		i.Repr()
		for {
			select {
			case <-ctx.Done():
				Writer.Cursor.Show()
				os.Exit(1)
			case msg := <-im.C:
				complete, err := i.HandleInput(msg)
				if err != nil {
					// todo: notify of invalid input and re-prompt
					return err
				}
				if complete {
					continue QUESTIONS
				}
				i.Repr()
			}
		}
	}
	Writer.Cursor.Show()
	return nil
}
