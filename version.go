package spark

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	GoVersion string
	BuildTime string
	GitCommit string
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
	s.WriteString(GoVersion + "\n")
	s.WriteString("Build Time: ")
	s.WriteString(BuildTime + "\n")
	s.WriteString("Git Commit: ")
	s.WriteString(GitCommit + "\n")
	return s.String()
}
