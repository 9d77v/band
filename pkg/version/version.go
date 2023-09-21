package version

import (
	"flag"
	"fmt"
	"os"
)

// 版本参数
var (
	git_tag    string
	git_commit string
	git_branch string
	build_time string
	build_name string
	go_version string
)

// FullVersion show the version info
func FullVersion() string {
	version := fmt.Sprintf(`%s version %s 
Build Name: %s
Build Time: %s
Git Branch: %s
Git Commit: %s
Go Version: %s`, build_name, git_tag,
		build_name, build_time,
		git_branch, git_commit, go_version)
	return version
}

func Short() string {
	return fmt.Sprintf("%s version %s", build_name, git_tag)
}

func Init() {
	var showVer bool

	var showShortVersion bool

	flag.BoolVar(&showVer, "v", false, "show build version")
	flag.BoolVar(&showShortVersion, "s", false, "show short build version")

	flag.Parse()
	if showVer {
		fmt.Println(FullVersion())
		os.Exit(0)
	}
	if showShortVersion {
		fmt.Println(Short())
		os.Exit(0)
	}
}
