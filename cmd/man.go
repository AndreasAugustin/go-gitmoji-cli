package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/charmbracelet/glamour/ansi"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/mattn/go-isatty"
	mcobra "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"
)

// We use this special character to represent the character that delimits the
// important symbols in the text. We replace these to code in markdown and bold
// in roff.
const specialChar = "%"
const margin = 2
const brightBlack = "#4d4d4d"

func boolPtr(b bool) *bool       { return &b }
func stringPtr(s string) *string { return &s }
func uintPtr(u uint) *uint       { return &u }

var GlamourTheme = ansi.StyleConfig{
	Document: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			BlockPrefix: "\n",
			BlockSuffix: "\n",
		},
		Margin: uintPtr(margin),
	},
	Heading: ansi.StyleBlock{
		StylePrimitive: ansi.StylePrimitive{
			BlockSuffix: "\n",
			Color:       stringPtr("99"),
			Bold:        boolPtr(true),
		},
	},
	Item:     ansi.StylePrimitive{Prefix: "· "},
	Emph:     ansi.StylePrimitive{Color: stringPtr(brightBlack)},
	Strong:   ansi.StylePrimitive{Bold: boolPtr(true)},
	Link:     ansi.StylePrimitive{Color: stringPtr("42"), Underline: boolPtr(true)},
	LinkText: ansi.StylePrimitive{Color: stringPtr("207")},
	Code:     ansi.StyleBlock{StylePrimitive: ansi.StylePrimitive{Color: stringPtr("204")}},
}

var (
	manDescription = fmt.Sprintf(`%s is a cli which lets you do git commits with conventional commits together with gitmoji format.`, pkg.ProgramName)

	manBugs = "See GitHub Issues: <https://github.com/AndreasAugustin/go-gitmoji-cli/issues>"

	manAuthor = "Andreas Augustin"
)

var manCmd = &cobra.Command{
	Use:     "manual",
	Aliases: []string{"man"},
	Short:   "Generate man pages",
	Args:    cobra.NoArgs,
	Hidden:  true,
	RunE: func(_ *cobra.Command, _ []string) error {
		if isatty.IsTerminal(os.Stdout.Fd()) {
			renderer, err := glamour.NewTermRenderer(
				glamour.WithStyles(GlamourTheme),
			)
			if err != nil {
				return err
			}
			v, err := renderer.Render(markdownManual())
			if err != nil {
				return err
			}
			fmt.Println(v)
			return nil
		}

		manPage, err := mcobra.NewManPage(1, RootCmd)
		if err != nil {
			return err
		}

		manPage = manPage.
			WithLongDescription(sanitizeSpecial(manDescription)).
			WithSection("Bugs", sanitizeSpecial(manBugs)).
			WithSection("Author", sanitizeSpecial(manAuthor)).
			WithSection("Copyright", "(C) 2021-2022 Charmbracelet, Inc.\n"+
				"Released under MIT license.")

		fmt.Println(manPage.Build(roff.NewDocument()))
		return nil
	},
}

func markdownManual() string {
	return fmt.Sprint(
		"# MANUAL\n", sanitizeMarkdown(manDescription),
		"# BUGS\n", manBugs,
		"\n# AUTHOR\n", manAuthor,
	)
}

func sanitizeMarkdown(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(
		s, "<", "&lt;"), ">", "&gt;"), specialChar, "`")
}

func sanitizeSpecial(s string) string {
	return strings.ReplaceAll(s, specialChar, "")
}

func init() {
	RootCmd.AddCommand(manCmd)
}
