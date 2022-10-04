package cmd

import (
	"fmt"
	"os"

	"github.com/grammyjs/gmy/internal/files"
	"github.com/grammyjs/gmy/internal/prompts/additional"
	"github.com/grammyjs/gmy/internal/prompts/name"
	"github.com/grammyjs/gmy/internal/prompts/platform"
	template_ "github.com/grammyjs/gmy/internal/prompts/template"

	"github.com/spf13/cobra"
)

const footer = `
 + Project '%s' created successfully!
 
 Thank you for using grammY <3
 . Documentation  https://grammy.dev
 . GitHub         https://github.com/grammyjs/grammY
 . Community      https://telegram.me/grammyjs
`

func Run(cmd *cobra.Command, args []string) {
	var projectName string
	if len(args) > 0 {
		err := name.ValidateProjectName(args[0])
		if err == nil {
			projectName = args[0]
		}
	}
	if projectName == "" {
		projectName = name.Prompt()
	}

	platform := platform.Prompt()
	template := template_.Prompt(platform)

	// Download, extract files from template source
	files.SetupFiles(projectName, template)

	fmt.Println(`
 + Project folder set up successfully!
   Additional configuration options for you:\n`)

	// Additional prompts and configuration
	additional.Prompt(projectName, platform, template)
	fmt.Printf(footer, projectName)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gmy [project name]",
	Short: "Create grammY projects from the command line!",
	Long: `grammY is a framework for creating Telegram bots. It can be used from
TypeScript and JavaScript and runs on Node.js, Deno, and in the browser.
This application is a tool to generate the needed files to quickly
create a Telegram bot from several templates powered by grammY.

Feel free to add more templates to our template registry:
https://github.com/grammyjs/gmy#templates

https://grammy.dev`,
	Args: cobra.ArbitraryArgs,
	Run:  Run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
