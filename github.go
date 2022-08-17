package gogithubpackageclean

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

// RateLimitError nice error for users
var RateLimitError = errors.New("Github rate limit hit")

func newClient(config Config) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}

func GetPackage(client *github.Client, config *Config) *github.Package {
	p, _, err := client.Organizations.GetPackage(context.Background(), config.Org, config.PackageType, config.PackageName)
	if _, ok := err.(*github.RateLimitError); ok {
		log.Fatal(RateLimitError)
	} else if err != nil {
		log.Fatal(err)
	}
	return p
}

func FetchPackageVersions(client *github.Client, config *Config, startingPage, versionsCount int) []int64 {
	opts := &github.PackageListOptions{}
	opts.ListOptions = github.ListOptions{
		Page:    startingPage,
		PerPage: 100,
	}
	packageVersionIds := make([]int64, 0, versionsCount)
	packageVersionIdsIndex := 0
	for {
		log.Printf("fetching page: %d\n", opts.ListOptions.Page)
		v, r, err := client.Organizations.PackageGetAllVersions(context.Background(), config.Org, config.PackageType, config.PackageName, opts)
		if _, ok := err.(*github.RateLimitError); ok {
			log.Fatal("hit rate limit")
		}
		if err != nil {
			log.Fatal(err)
		}
		index := (opts.ListOptions.Page - 1) * 100
		log.Printf("index: %d", index)
		for vi, vv := range v {
			if index+vi >= config.MaxVersions {
				if id := vv.GetID(); id != 0 {

					packageVersionIds = append(packageVersionIds, id)
					packageVersionIdsIndex += 1
				} else {
					log.Fatal("unexpected zero value for GetID")
				}
			} else {
				log.Printf("skipping index %d", index+vi)
			}
		}
		log.Printf("version Count: %d", len(packageVersionIds))
		if np := r.NextPage; np > 0 {
			opts.ListOptions.Page = np
		} else {
			break
		}
	}
	return packageVersionIds
}

func DeletePackageVersions(client *github.Client, config *Config, versions []int64) {
	for i, v := range versions {
		if !config.Commit {
			log.Printf("[%d] DRYRUN - Delete package version ID=%d", i, v)
			continue
		}
		log.Printf("[%d] Deleting package version ID=%d", i, versions[i])
		r, err := client.Organizations.PackageDeleteVersion(context.Background(), config.Org, config.PackageType, config.PackageName, v)
		if _, ok := err.(*github.RateLimitError); ok {
			log.Fatal("hit rate limit")
		}
		if err != nil {
			log.Printf("[%d] Failed to delete: %s", i, err)
			if config.ContinueOnFail {
				log.Printf("[%d] Continuing", i)
			} else {
				log.Fatalf("[%d] Stopping on fatal error to delete", i)
			}
		}
		if r.StatusCode < 300 {
			log.Printf("[%d] Success", i)
		} else {
			log.Printf("[%d] StatusCode: %d, continuing", i, r.StatusCode)
		}
	}
}

// Process what has been requested in the Config
func Process(config *Config) {
	client := newClient(*config)

	pkg := GetPackage(client, config)

	log.Printf("Version count: %d\n", *pkg.VersionCount)

	startingPage, versionsCount := CalculateStartPage(config, int(*pkg.VersionCount))

	if startingPage == 0 {
		log.Println("Nothing to do when under the max limit")
		os.Exit(0)
	}

	packageVersionIds := FetchPackageVersions(client, config, startingPage, versionsCount)

	if config.OldestToNewest {
		reverse(packageVersionIds)
	}
	DeletePackageVersions(client, config, packageVersionIds)

}

func CalculateStartPage(config *Config, versionCount int) (page int, versionsCount int) {
	if versionCount > config.MaxVersions {
		versionsCount := versionCount - config.MaxVersions
		log.Printf("%d package versions to remove\n", versionsCount)

		page = (config.MaxVersions / 100) + 1
	}
	return
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
