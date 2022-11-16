package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const configDir = "configs"

var rootCmd = &cobra.Command{
	Use:   "vcli",
	Short: "Vending CLI in Go",
	Long:  "Vending CLI application written in Go.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
