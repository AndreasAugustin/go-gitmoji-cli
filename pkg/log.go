package pkg

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	log "github.com/sirupsen/logrus"
)

var levelColor = lipgloss.Color("205")

type PlainFormatter struct {
}

func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s: %s\n", lipgloss.NewStyle().SetString(entry.Level.String()).Foreground(levelColor), entry.Message)), nil
}
func ToggleDebug(debug bool) {
	if debug {
		log.Info("Debug logs enabled")
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{})
	} else {
		plainFormatter := new(PlainFormatter)
		log.SetFormatter(plainFormatter)
	}
}
