package main

import (
	"github.com/spf13/cobra"

	"github.com/merico-dev/stream/internal/pkg/pluginengine"
	"github.com/merico-dev/stream/pkg/util/log"
)

var verifyCMD = &cobra.Command{
	Use:   "verify",
	Short: "Verify DevOps tools according to DevStream configuration file",
	Long:  `Verify DevOps tools according to DevStream configuration file.`,
	Run:   verifyCMDFunc,
}

func verifyCMDFunc(cmd *cobra.Command, args []string) {
	log.Info("Verify started.")
	healthy, err := pluginengine.Verify(configFile)
	if err != nil {
		log.Fatalf("Verify error: %s.", err)
	}

	if healthy {
		log.Success("All tools are healthy.")
	} else {
		log.Error("Some tools are NOT healthy!!!")
	}
}
