package gogithubpackageclean

import (
	"testing"
)

var calcTests [][]int = [][]int{
	{100, 80, 0},
	{20, 20, 0},
	{20, 21, 1},
	{20, 80, 1},
	{120, 140, 2},
	{120, 220, 2},
}

func Test_CalculateStartPage(t *testing.T) {
	config := &Config{}
	for _, v := range calcTests {
		config.MaxVersions = v[0]
		if p, _ := CalculateStartPage(config, v[1]); p != v[2] {
			t.Errorf("maxVersion=%d versionCount=%d expected page %d but got %d", v[0], v[1], v[2], p)
		}
	}
}

func Test_GetPackage(t *testing.T) {

}
