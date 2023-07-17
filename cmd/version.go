package cmd

import (
	"fmt"
	"github.com/zeromicro/goctl-swagger/variable"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version info",
	Long:  `version info detail info`,
	Run: func(cmd *cobra.Command, args []string) {
		verSion := fmt.Sprintf("%s %s/%s", variable.BuildVersion, runtime.GOOS, runtime.GOARCH)
		fmt.Println(verSion)
	},
}
