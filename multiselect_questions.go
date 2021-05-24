package prompt

import "fmt"

func multiOptionsRepr(prompt string, options []interface{}, index int, selected map[int]bool, clear bool) {
	if clear {
		Writer.Cursor.Up(len(options) + 1).ClearToEndOfScreen()
	}
	formatPrompt()
	fmt.Fprintln(Writer, prompt)
	for i, option := range options {
		Writer.Cursor.Clear()
		if selected[i] {
			Writer.Cursor.Green()
			fmt.Fprint(Writer, "⦿ ")
			Writer.Cursor.Clear()
		} else {
			fmt.Fprint(Writer, "○ ")
		}
		if i == index {
			formatSelection()
		}
		fmt.Fprintln(Writer, option)
		Writer.Cursor.Clear()
	}
}

type StringMultiOptionQuestion struct {
	called      bool
	index       int
	selected    map[int]bool
	Prompt      string
	Options     []string
	Destination *[]string
	OnComplete  func()
}

func (o *StringMultiOptionQuestion) Init() {
	o.selected = make(map[int]bool)
	for i := range o.Options {
		o.selected[i] = false
	}
}

func (o *StringMultiOptionQuestion) Repr() {
	opt := make([]interface{}, len(o.Options))
	for i, option := range o.Options {
		opt[i] = option
	}
	multiOptionsRepr(o.Prompt, opt, o.index, o.selected, o.called)
	if !o.called {
		o.called = true
	}
}

func (o *StringMultiOptionQuestion) HandleInput(input []byte) (bool, error) {
	switch {
	case isEnter(input):
		v := make([]string, 0)
		for index, selected := range o.selected {
			if selected {
				v = append(v, o.Options[index])
			}
		}
		*o.Destination = v
		if f := o.OnComplete; f != nil {
			f()
		}
		return true, nil
	case isUpArrow(input):
		o.index = (o.index - 1 + len(o.Options)) % len(o.Options)
	case isDownArrow(input):
		o.index = (o.index + 1) % len(o.Options)
	case isSpace(input):
		o.selected[o.index] = !o.selected[o.index]
	}
	return false, nil
}

type IntMultiOptionQuestion struct {
	called      bool
	index       int
	selected    map[int]bool
	Prompt      string
	Options     []int
	Destination *[]int
	OnComplete  func()
}

func (o *IntMultiOptionQuestion) Init() {
	o.selected = make(map[int]bool)
	for i := range o.Options {
		o.selected[i] = false
	}
}

func (o *IntMultiOptionQuestion) Repr() {
	opt := make([]interface{}, len(o.Options))
	for i, option := range o.Options {
		opt[i] = option
	}
	multiOptionsRepr(o.Prompt, opt, o.index, o.selected, o.called)
	if !o.called {
		o.called = true
	}
}

func (o *IntMultiOptionQuestion) HandleInput(input []byte) (bool, error) {
	switch {
	case isEnter(input):
		v := make([]int, 0)
		for index, selected := range o.selected {
			if selected {
				v = append(v, o.Options[index])
			}
		}
		*o.Destination = v
		if f := o.OnComplete; f != nil {
			f()
		}
		return true, nil
	case isUpArrow(input):
		o.index = (o.index - 1 + len(o.Options)) % len(o.Options)
	case isDownArrow(input):
		o.index = (o.index + 1) % len(o.Options)
	case isSpace(input):
		o.selected[o.index] = !o.selected[o.index]
	}
	return false, nil
}

type FloatMultiOptionQuestion struct {
	called      bool
	index       int
	selected    map[int]bool
	Prompt      string
	Options     []float64
	Destination *[]float64
	OnComplete  func()
}

func (o *FloatMultiOptionQuestion) Init() {
	o.selected = make(map[int]bool)
	for i := range o.Options {
		o.selected[i] = false
	}
}

func (o *FloatMultiOptionQuestion) Repr() {
	opt := make([]interface{}, len(o.Options))
	for i, option := range o.Options {
		opt[i] = option
	}
	multiOptionsRepr(o.Prompt, opt, o.index, o.selected, o.called)
	if !o.called {
		o.called = true
	}
}

func (o *FloatMultiOptionQuestion) HandleInput(input []byte) (bool, error) {
	switch {
	case isEnter(input):
		v := make([]float64, 0)
		for index, selected := range o.selected {
			if selected {
				v = append(v, o.Options[index])
			}
		}
		*o.Destination = v
		if f := o.OnComplete; f != nil {
			f()
		}
		return true, nil
	case isUpArrow(input):
		o.index = (o.index - 1 + len(o.Options)) % len(o.Options)
	case isDownArrow(input):
		o.index = (o.index + 1) % len(o.Options)
	case isSpace(input):
		o.selected[o.index] = !o.selected[o.index]
	}
	return false, nil
}
