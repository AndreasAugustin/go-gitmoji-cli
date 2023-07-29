package pkg

import "fmt"

type Gitmoji struct {
	Emoji  string `json:"emoji"`
	Entity string `json:"entity"`
	Code   string `json:"code"`
	Desc   string `json:"description"`
	Name   string `json:"name"`
	Semver string `json:"semver"`
}

func (i Gitmoji) FilterValue() string { return i.Name + i.Desc }
func (i Gitmoji) Title() string       { return fmt.Sprintf("%s %s", i.Emoji, i.Code) }
func (i Gitmoji) Description() string { return i.Desc }

type Gitmojis struct {
	Gitmojis []Gitmoji `json:"gitmojis"`
}
