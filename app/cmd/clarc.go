package cmd

import (
	"fmt"
	"os"

	"github.com/mtoku/clarc/app/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clarc",
	Short: "clarc",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		_ = err
		err = server.RunServer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
