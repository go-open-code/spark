package version

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	goVersion string
	buildTime string
	gitCommit string
)

var (
	showVersion = flag.Bool("v", false, "show build version")
)

func Parse() {
	flag.Parse()
	if *showVersion {
		fmt.Println(String())
		os.Exit(0)
	}
}

func String() string {
	s := strings.Builder{}
	s.WriteString("Go Version: ")
	s.WriteString(goVersion + "\n")
	s.WriteString("Build Time: ")
	s.WriteString(buildTime + "\n")
	s.WriteString("Git Commit: ")
	s.WriteString(gitCommit + "\n")
	return s.String()
}
