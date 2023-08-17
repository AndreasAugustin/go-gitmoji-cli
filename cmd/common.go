package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func bindAndAddBoolFlagP(cmd *cobra.Command, viperKey string, flagName string, shortHand string, usage string) {
	cmd.PersistentFlags().BoolP(flagName, shortHand, false, fmt.Sprintf("%s. If not set the configured value is used", usage))
	err := viper.BindPFlag(viperKey, CommitCmd.PersistentFlags().Lookup(flagName))

	if err != nil {
		log.Fatalf("issue with binding flags %s", err)
	}
}

func bindAndAddBoolFlag(cmd *cobra.Command, viperKey string, flagName string, usage string) {
	cmd.PersistentFlags().Bool(flagName, false, fmt.Sprintf("%s. If not set the configured value is used", usage))
	err := viper.BindPFlag(viperKey, CommitCmd.PersistentFlags().Lookup(flagName))

	if err != nil {
		log.Fatalf("issue with binding flags %s", err)
	}
}
