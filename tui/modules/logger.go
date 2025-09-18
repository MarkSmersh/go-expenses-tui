package modules

import (
	"fmt"
	"os"
)

type Logger struct {
	File string
}

func (l Logger) Logf(s string, args ...any) {
	formatted := fmt.Sprintf(s, args...)

	f, _ := os.OpenFile(l.File, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	f.WriteString(formatted + "\n")

	defer f.Close()
}
