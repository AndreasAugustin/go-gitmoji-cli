package pkg

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

func ExecuteCommit(title string, body string, config Config) {
	commitCmd, err := utils.BuildGitCommitCommandStr(config.AutoAdd, config.AutoSign, title, body)
	if err != nil {
		log.Fatalf("error building commit message: %s", err)
	}
	log.Debugf("commit cmd: %s", commitCmd)
	utils.Commit(commitCmd)
}

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
