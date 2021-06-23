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
	FinalRepr()
	HandleInput(input []byte) (bool, error)
}

var _ Question = &IntInputQuestion{}
var _ Question = &FloatInputQuestion{}
var _ Question = &StringInputQuestion{}
var _ Question = &StringMultiOptionQuestion{}
var _ Question = &IntMultiOptionQuestion{}
var _ Question = &FloatMultiOptionQuestion{}
var _ Question = &BoolQuestion{}
var _ Question = &StringOptionQuestion{}
var _ Question = &IntOptionQuestion{}
var _ Question = &FloatOptionQuestion{}

type Prompt struct {
	index     int
	Questions []Question
}

func (p *Prompt) next() Question {
	if p.index > len(p.Questions)-1 {
		return nil
	}
	q := p.Questions[p.index]
	p.index++
	return q
}

func (p *Prompt) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	im := newInputManager()
	Writer.Cursor.Hide()
	q := p.next()
RunLoop:
	for q != nil {
		q.Init()
		q.Repr()
		for {
			select {
			case <-ctx.Done():
				Writer.Cursor.Show()
				os.Exit(1)
			case msg := <-im.C:
				complete, err := q.HandleInput(msg)
				if err != nil {
					// todo: notify of invalid input and re-prompt
					return err
				}
				if complete {
					q.FinalRepr()
					q = p.next()
					continue RunLoop
				}
				q.Repr()
			}
		}
	}
	Writer.Cursor.Show()
	return nil
}
