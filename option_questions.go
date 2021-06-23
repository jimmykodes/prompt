package prompt

import (
	"fmt"
)

func optionsRepr(prompt string, options []interface{}, index int, clear bool) {
	if clear {
		Writer.Cursor.Up(len(options) + 1).ClearToEndOfScreen()
	}
	formatPrompt()
	fmt.Fprintln(Writer, prompt)
	for i, option := range options {
		prefix := " "
		if i == index {
			prefix = "‣"
			formatSelection()
		} else {
			Writer.Cursor.Clear()
		}
		fmt.Fprintln(Writer, prefix, option)
		Writer.Cursor.Clear()
	}
}
func finalOption(prompt string, options []interface{}, choice interface{}) {
	Writer.Cursor.Up(len(options) + 1).ClearToEndOfScreen()
	formatPrompt()
	fmt.Fprintf(Writer, "%s ", prompt)
	formatSelection()
	fmt.Fprintln(Writer, choice)
	Writer.Cursor.Clear()
}

type BoolQuestion struct {
	called      bool
	index       int
	Prompt      string
	Destination *bool
	OnComplete  func()
}

func (b *BoolQuestion) Init() {}

func (b *BoolQuestion) FinalRepr() {
	choices := []interface{}{"True", "False"}
	finalOption(b.Prompt, []interface{}{"True", "False"}, choices[b.index])
}

func (b *BoolQuestion) Repr() {
	optionsRepr(b.Prompt, []interface{}{"True", "False"}, b.index, b.called)
	if !b.called {
		b.called = true
	}
}

func (b *BoolQuestion) HandleInput(input []byte) (bool, error) {
	switch {
	case isEnter(input):
		*b.Destination = b.index == 0
		if f := b.OnComplete; f != nil {
			f()
		}
		return true, nil
	case isDownArrow(input), isUpArrow(input):
		b.index = (b.index + 1) % 2
	}
	return false, nil
}

type StringOptionQuestion struct {
	index       int
	called      bool
	Prompt      string
	Options     []string
	Destination *string
	OnComplete  func()
}

func (o *StringOptionQuestion) FinalRepr() {
	opt := make([]interface{}, len(o.Options))
	for i, option := range o.Options {
		opt[i] = option
	}
	finalOption(o.Prompt, opt, o.Options[o.index])
}

func (o *StringOptionQuestion) Init() {}

func (o *StringOptionQuestion) Repr() {
	opt := make([]interface{}, len(o.Options))
	for i, option := range o.Options {
		opt[i] = option
	}
	optionsRepr(o.Prompt, opt, o.index, o.called)
	if !o.called {
		o.called = true
	}
}

func (o *StringOptionQuestion) HandleInput(input []byte) (bool, error) {
	switch {
	case isEnter(input):
		*o.Destination = o.Options[o.index]
		if f := o.OnComplete; f != nil {
			f()
		}
		return true, nil
	case isUpArrow(input):
		o.index = (o.index - 1 + len(o.Options)) % len(o.Options)
	case isDownArrow(input):
		o.index = (o.index + 1) % len(o.Options)
	}
	return false, nil
}

type IntOptionQuestion struct {
	index       int
	called      bool
	Prompt      string
	Options     []int
	Destination *int
	OnComplete  func()
}

func (o *IntOptionQuestion) FinalRepr() {
	opt := make([]interface{}, len(o.Options))
	for i, option := range o.Options {
		opt[i] = option
	}
	finalOption(o.Prompt, opt, o.Options[o.index])
}

func (o *IntOptionQuestion) Init() {}

func (o *IntOptionQuestion) Repr() {
	opt := make([]interface{}, len(o.Options))
	for i, option := range o.Options {
		opt[i] = option
	}
	optionsRepr(o.Prompt, opt, o.index, o.called)
	if !o.called {
		o.called = true
	}
}

func (o *IntOptionQuestion) HandleInput(input []byte) (bool, error) {
	switch {
	case isEnter(input):
		*o.Destination = o.Options[o.index]
		if f := o.OnComplete; f != nil {
			f()
		}
		return true, nil
	case isUpArrow(input):
		o.index = (o.index - 1 + len(o.Options)) % len(o.Options)
	case isDownArrow(input):
		o.index = (o.index + 1) % len(o.Options)
	}
	return false, nil
}

type FloatOptionQuestion struct {
	index       int
	called      bool
	Prompt      string
	Options     []float64
	Destination *float64
	OnComplete  func()
}

func (o *FloatOptionQuestion) FinalRepr() {
	opt := make([]interface{}, len(o.Options))
	for i, option := range o.Options {
		opt[i] = option
	}
	finalOption(o.Prompt, opt, o.Options[o.index])
}

func (o *FloatOptionQuestion) Init() {}

func (o *FloatOptionQuestion) Repr() {
	opt := make([]interface{}, len(o.Options))
	for i, option := range o.Options {
		opt[i] = option
	}
	optionsRepr(o.Prompt, opt, o.index, o.called)
	if !o.called {
		o.called = true
	}
}

func (o *FloatOptionQuestion) HandleInput(input []byte) (bool, error) {
	switch {
	case isEnter(input):
		*o.Destination = o.Options[o.index]
		if f := o.OnComplete; f != nil {
			f()
		}
		return true, nil
	case isUpArrow(input):
		o.index = (o.index - 1 + len(o.Options)) % len(o.Options)
	case isDownArrow(input):
		o.index = (o.index + 1) % len(o.Options)
	}
	return false, nil
}
