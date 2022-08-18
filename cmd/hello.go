/*
Package cmd
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Name string

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "this command can be say `hello`",
	Long: `this command can be say "hello"

if 
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(Name) == 0 {
			fmt.Println("hello!")
		} else {
			fmt.Printf("hello %s!\n", Name)
		}
	},
}

func init() {
	helloCmd.Flags().StringVarP(&Name, "name", "n", "", "print your name")
	rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
