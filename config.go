package gogithubpackageclean

import (
	"os"

	"github.com/isobit/cli"
)

type Config struct {
	GithubToken    string `cli:"help='github personal access token',env=GITHUB_TOKEN,required"`
	Org            string `cli:"help='github org or username',env=GITHUB_ORG,required"`
	PackageType    string `cli:"help='package type: npm, maven, rubygems, docker, nuget, container',env=GITHUB_PACKAGE_TYPE,required"`
	PackageName    string `cli:"help='package name',env=GITHUB_PACKAGE_NAME,required"`
	Commit         bool   `cli:"help='make actual changes'"`
	MaxVersions    int    `cli:"help='max number of versions that should be allowed to exist',name=max"`
	MinAge         int    `cli:"help='minimum age in days for version before deleting it',name=min-age"`
	OldestToNewest bool   `cli:"help='delete from oldest to newest'"`
	ContinueOnFail bool   `cli:"help='continue on failure to delete'"`
}

func (c *Config) Run() error {
	Process(c)
	return nil
}

func NewConfig() cli.ParseResult {
	return cli.New(os.Args[0], &Config{}).Parse()
}
