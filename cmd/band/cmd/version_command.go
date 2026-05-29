package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

const (
	appVersion = "0.5.4"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  `Show the version info of band tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("band v%s\n", appVersion)
		fmt.Printf("Go Runtime: %s\n", runtime.Version())
		fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}
