package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/gland/pkg/log"
)

var (
	rootCmd = &cobra.Command{
		Use:          "coordinator",
		Short:        "Gland saga execution coordinator",
		Long:         "Gland saga execution coordinator",
		SilenceUsage: true,
	}
)

func main()  {
	if err := rootCmd.Execute(); err != nil {
		log.Error("", err)
		os.Exit(-1)
	}
}
