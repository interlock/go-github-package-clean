# go-github-package-clean

Small static binary utility to clean up packages on github.

## Params

|Flag|Environment|Description|
|----|-----------|-----------|
|github-token|GITHUB_TOKEN|github personal access token|
|org|GITHUB_ORG|github org or username|
|package-type|GITHUB_PACKAGE_TYPE|package type: npm, maven, rubygems, docker, nuget, container|
|package-name|GITHUB_PACKAGE_NAME|name of package|
|commit||make all changes, otherwise runs are dry|
|max||max number of versions that should be allowed to exist|
|min-age||minimum age in days for a version before deleting it|
|oldest-to-newest||delete from oldest to newest|
|continue-on-fail||continue on failure to delete|

## Examples

### CircleCI

### Github
