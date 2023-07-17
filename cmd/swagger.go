package cmd

import (
	"github.com/zeromicro/goctl-swagger/action"

	"github.com/spf13/cobra"
)

var swaggerCmd = &cobra.Command{
	Use:   "swagger",
	Short: "generates swagger.json",
	Long:  `generates swagger.json`,
	RunE:  action.Generator,
}

func init() {
	rootCmd.AddCommand(swaggerCmd)

	swaggerCmd.Flags().StringVar(&action.VarStringHost, "host", "", `api request address`)
	swaggerCmd.Flags().StringVar(&action.VarStringBasePath, "basepath", "", `url request prefix`)
	swaggerCmd.Flags().StringVar(&action.VarStringFileName, "filename", "", `swagger save file name`)
}
