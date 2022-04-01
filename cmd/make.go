/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/teppei22/fgen/gen"
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Automatic model generation",
	Long: `Automatic model generation.
if you enter the name of the model, it's automatically generated.
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("make called")

		model, err := cmd.Flags().GetString("model")
		if err != nil {
			log.Fatal(err)
		}

		conf := &gen.Config{
			OutputPath: "./output",
			Model:      model,
		}
		AutoGen := gen.NewAutoGen(*conf)

		if err := AutoGen.GenerateFileAll(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	makeCmd.Flags().String("model", "task", "Enter the name of the model by automatically generated")
}
