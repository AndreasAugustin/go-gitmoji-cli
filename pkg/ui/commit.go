package ui

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	log "github.com/sirupsen/logrus"
)

func CommitPrompt(config pkg.Config, gitmojis []pkg.Gitmoji, initialValues pkg.InitialCommitValues, isBreaking bool) pkg.CommitValues {
	defaultTypes := pkg.DefaultCommitTypes()
	listSettingsGitmojis := ListSettings{IsShowStatusBar: true, IsFilteringEnabled: true, Title: "Gitmojis"}
	listSettingsCommitTypes := ListSettings{Title: "Commit types", IsShowStatusBar: true, IsFilteringEnabled: true}

	selectedGitmoji := ListRun(listSettingsGitmojis, gitmojis)
	log.Debugf("selected gitmoji %v", selectedGitmoji)
	if initialValues.Type == "" {
		log.Debug("no type provided with `-m` flag. Starting type prompt")
		selectedPromptType := ListRun(listSettingsCommitTypes, defaultTypes)
		log.Debugf("selected %s", selectedPromptType)
		initialValues.Type = selectedPromptType.Type
	}

	textInputsData := initialValues.BuildTextInputsData(config)
	inputsRes := TextInputsRun("please add the related information", textInputsData)

	commitValues := pkg.CreateMessage(inputsRes, selectedGitmoji, initialValues, config, isBreaking)

	log.Debugf("complete title: %v", commitValues.Title)
	return commitValues
}
