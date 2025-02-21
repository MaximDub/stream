package pluginengine

import (
	"fmt"
	"os"

	"github.com/tcnksm/go-input"

	"github.com/merico-dev/stream/internal/pkg/configloader"
	"github.com/merico-dev/stream/pkg/util/log"
)

func getStateKeyFromTool(t *configloader.Tool) string {
	return fmt.Sprintf("%s_%s", t.Name, t.Plugin.Kind)
}

func readUserInput() string {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	query := "Continue? [y/n]"
	userInput, err := ui.Ask(query, &input.Options{
		Required: true,
		Default:  "n",
		Loop:     true,
		ValidateFunc: func(s string) error {
			if s != "y" && s != "n" {
				return fmt.Errorf("input must be y or n")
			}
			return nil
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return userInput
}
