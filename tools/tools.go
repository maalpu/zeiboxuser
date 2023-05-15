package tools

import (
	"fmt"
	"time"
)

// Em MySQL la fecha se graba en el siguiente formato:
// AAAA-MM-DD THH:MM:SS

func FechaMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
