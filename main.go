package main

import (
	"github.com/spf13/cobra"
	"github.com/brndt/go_helloworld/beers"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.Execute()
}
