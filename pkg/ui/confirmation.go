package ui

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	log "github.com/sirupsen/logrus"
)

func ConfirmationRun(title string) bool {
	listSettings := ListSettings{Title: title, IsFilteringEnabled: false, IsShowStatusBar: true}
	res := ListRun(listSettings, []pkg.YesNo{pkg.YES, pkg.NO})
	log.Debugf(string(res))

	return res == pkg.YES
}
