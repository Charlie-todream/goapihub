package cmd

import (
	"github.com/spf13/cobra"
	"goapihub/pkg/helpers"
	console "goapihub/pkg/pkg"
)

var CmdKey = &cobra.Command{
	Use:"Key",
	Short:"Generate App Key, will print the generated Key",
	Run: runKeyGenerate,
	Args:cobra.NoArgs, // 不允许传参
}

func runKeyGenerate(cmd *cobra.Command,args []string)  {
	console.Success("App Key:")
	console.Success(helpers.RandomString(32))
	console.Success("---")
	console.Warning("please go to .env file to change the APP_KEY option")
}
