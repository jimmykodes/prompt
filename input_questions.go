package prompt

import (
	"fmt"
	"strconv"
	"strings"
)

func inputPrompt(called bool, prompt string, val strings.Builder, def interface{}, showDefault bool) {
	if called {
		Writer.Cursor.ToCol(0).ClearToEndOfScreen()
	}
	formatPrompt()
	fmt.Fprint(Writer, prompt, " ")
	Writer.Cursor.Clear()
	if showDefault && val.Len() == 0 {
		fmt.Fprint(Writer, "(default:", def, ")")
	}
	fmt.Fprint(Writer, val.String())
}
func finalInput(prompt string, val strings.Builder, def interface{}) {
	Writer.Cursor.Up(1).ToCol(0).ClearToEndOfScreen()
	var v interface{}
	if val.Len() == 0 {
		v = def
	} else {
		v = val.String()
	}
	formatPrompt()
	fmt.Fprint(Writer, prompt, " ")
	formatSelection()
	fmt.Fprintln(Writer, v)
	Writer.Cursor.Clear()
}

type IntInputQuestion struct {
	called      bool
	val         strings.Builder
	Prompt      string
	Destination *int
	Default     int
	OnComplete  func()
}

func (i *IntInputQuestion) FinalRepr() {
	finalInput(i.Prompt, i.val, i.Default)
}

func (i *IntInputQuestion) Init() {
	Writer.Cursor.Show()
}

func (i *IntInputQuestion) Repr() {
	inputPrompt(i.called, i.Prompt, i.val, i.Default, i.Default > 0)
	if !i.called {
		i.called = true
	}
}

func (i *IntInputQuestion) HandleInput(input []byte) (bool, error) {
	switch {
	case isEnter(input):
		Writer.Cursor.Hide()
		// make sure to write a newline so since this is inline input
		fmt.Fprintln(Writer)
		if i.val.Len() > 0 {
			v, err := strconv.ParseInt(i.val.String(), 10, 64)
			if err != nil {
				return true, err
			}
			*i.Destination = int(v)
		} else {
			*i.Destination = i.Default
		}
		if f := i.OnComplete; f != nil {
			f()
		}
		return true, nil
	case isBackspace(input):
		{
			v := i.val.String()[:i.val.Len()-1]
			i.val.Reset()
			i.val.WriteString(v)
		}
	case !isEscapeSequence(input[:2]):
		i.val.Write(reduceInput(input))
	}
	return false, nil
}

type FloatInputQuestion struct {
	called      bool
	val         strings.Builder
	Prompt      string
	Destination *float64
	Default     float64
	OnComplete  func()
}

func (f *FloatInputQuestion) FinalRepr() {
	finalInput(f.Prompt, f.val, f.Default)
}

func (f *FloatInputQuestion) Init() {
	Writer.Cursor.Show()
}

func (f *FloatInputQuestion) Repr() {
	inputPrompt(f.called, f.Prompt, f.val, f.Default, f.Default > 0)
	if !f.called {
		f.called = true
	}
}

func (f *FloatInputQuestion) HandleInput(input []byte) (bool, error) {
	switch {
	case isEnter(input):
		Writer.Cursor.Hide()
		// make sure to write a newline so since this is inline input
		fmt.Fprintln(Writer)
		if f.val.Len() > 0 {
			v, err := strconv.ParseFloat(f.val.String(), 64)
			if err != nil {
				return true, err
			}
			*f.Destination = v
		} else {
			*f.Destination = f.Default
		}
		if f := f.OnComplete; f != nil {
			f()
		}
		return true, nil
	case isBackspace(input):
		{
			v := f.val.String()[:f.val.Len()-1]
			f.val.Reset()
			f.val.WriteString(v)
		}
	case !isEscapeSequence(input[:2]):
		f.val.Write(reduceInput(input))
	}
	return false, nil
}

type StringInputQuestion struct {
	called      bool
	val         strings.Builder
	Prompt      string
	Destination *string
	Default     string
	OnComplete  func()
}

func (s *StringInputQuestion) FinalRepr() {
	finalInput(s.Prompt, s.val, s.Default)
}

func (s *StringInputQuestion) Init() {
	Writer.Cursor.Show()
}

func (s *StringInputQuestion) Repr() {
	inputPrompt(s.called, s.Prompt, s.val, s.Default, s.Default != "")
	if !s.called {
		s.called = true
	}
}

func (s *StringInputQuestion) HandleInput(input []byte) (bool, error) {
	switch {
	case isEnter(input):
		Writer.Cursor.Hide()
		// make sure to write a newline so since this is inline input
		fmt.Fprintln(Writer)
		if s.val.Len() > 0 {
			*s.Destination = s.val.String()
		} else {
			*s.Destination = s.Default
		}
		if f := s.OnComplete; f != nil {
			f()
		}
		return true, nil
	case isBackspace(input):
		{
			v := s.val.String()[:s.val.Len()-1]
			s.val.Reset()
			s.val.WriteString(v)
		}
	case !isEscapeSequence(input[:2]):
		s.val.Write(reduceInput(input))
	}
	return false, nil
}
