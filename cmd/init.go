/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/teppei22/fji-codegen/gen"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "generate the setup code for golang application",
	Long: `
init of Fgen is a command to init the golang application composed of layered architecture.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")

		conf := &gen.Config{
			OutputPath: "./output",
			Model:      "task",
		}
		AutoGen := gen.NewAutoGen(*conf)

		if err := AutoGen.Init(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
