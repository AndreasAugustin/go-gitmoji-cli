package pkg

import (
	"fmt"
	"github.com/fatih/color"
)

func ColoredRepr(gitmoji Gitmoji) string {
	blue := color.New(color.FgBlue).SprintFunc()
	return fmt.Sprintf("%s %s %s", gitmoji.Emoji, blue(gitmoji.Code), gitmoji.Description)
}

func PrintEmojis(gitmojis Gitmojis) {
	for _, gitmoji := range gitmojis.Gitmojis {
		fmt.Println(ColoredRepr(gitmoji))
	}
}
