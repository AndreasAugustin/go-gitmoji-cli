package pkg

import (
	"fmt"
	"strings"
)

func BuildCommitTitle(_type string, scope string, isBreaking bool, desc string, gitmoji Gitmoji, config Config) string {

	var s strings.Builder

	s.WriteString(_type)

	if scope != "" {
		s.WriteString(fmt.Sprintf("(%s)", scope))
	}

	if isBreaking {
		s.WriteString("!")
	}

	s.WriteString(fmt.Sprintf(": %s ", gitmojiToString(gitmoji, config)))
	s.WriteString(eventualCapitalizeTitle(desc, config))

	return s.String()
}

func gitmojiToString(gitmoji Gitmoji, config Config) string {
	if config.EmojiFormat == CODE {
		return gitmoji.Code
	} else {
		return gitmoji.Emoji
	}
}

func eventualCapitalizeTitle(title string, config Config) string {
	if config.CapitalizeTitle {
		return strings.ToUpper(title)
	}
	return title
}
