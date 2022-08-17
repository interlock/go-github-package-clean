package gogithubpackageclean

import (
	"flag"
	"os"
)

type Config struct {
	GithubToken    string
	Org            string
	PackageType    string
	PackageName    string
	Commit         bool
	MaxVersions    int
	MinAge         int
	OldestToNewest bool
	ContinueOnFail bool
}

func NewConfig() *Config {
	config := &Config{}
	if v, ok := os.LookupEnv("GITHUB_TOKEN"); ok {
		config.GithubToken = v
	}
	if v, ok := os.LookupEnv("GITHUB_ORG"); ok {
		config.Org = v
	}
	if v, ok := os.LookupEnv("GITHUB_PACKAGE_TYPE"); ok {
		config.PackageType = v
	}
	if v, ok := os.LookupEnv("GITHUB_PACKAGE_NAME"); ok {
		config.PackageName = v
	}
	return config
}

// flags

func ParseFlags() *Config {
	config := NewConfig()
	flag.StringVar(&config.GithubToken, "github-token", "", "github personal access token [GITHUB_TOKEN]")
	flag.StringVar(&config.Org, "org", "", "github org or username [GITHUB_ORG]")
	flag.StringVar(&config.PackageType, "type", "", "package type: npm, maven, rubygems, docker, nuget, container [GITHUB_PACKAGE_TYPE]")
	flag.StringVar(&config.PackageName, "name", "", "package name [GITHUB_PACKAGE_NAME]")
	flag.BoolVar(&config.Commit, "commit", false, "")
	flag.BoolVar(&config.OldestToNewest, "oldest-to-newest", false, "delete from oldest to newest")
	flag.BoolVar(&config.ContinueOnFail, "continue-on-fail", false, "continue on failure to delete")
	flag.IntVar(&config.MaxVersions, "max", -1, "max number of versions that should be allowed to exist")
	flag.IntVar(&config.MinAge, "min_age", -1, "minimum age in days for version before deleting it")
	flag.Parse()
	return config
}
