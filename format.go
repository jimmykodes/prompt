package prompt

func formatPrompt() {
	Writer.Cursor.Clear().Bold().White()
}

func formatSelection() {
	Writer.Cursor.Clear().Magenta()
}
