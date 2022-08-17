package main

import (
	"testing"
)

func Test_BelowMax(t *testing.T) {
	config := &Config{}
	config.MaxVersions = 100

	if p, _ := calculateStartPage(config, 80); p != 0 {
		t.Errorf("expected page 0 but got %d", p)
	}
}

func Test_WithinPageSize(t *testing.T) {
	config := &Config{}
	config.MaxVersions = 20

	if p, _ := calculateStartPage(config, 20); p != 0 {

		t.Errorf("expected page 0 but got %d", p)
	}

	if p, _ := calculateStartPage(config, 21); p != 1 {

		t.Errorf("expected page 1 but got %d", p)
	}

	if p, _ := calculateStartPage(config, 80); p != 1 {

		t.Errorf("expected page 1 but got %d", p)
	}
}

func Test_OverPageSize(t *testing.T) {
	config := &Config{}
	config.MaxVersions = 120
	if p, _ := calculateStartPage(config, 140); p != 2 {
		t.Errorf("expected page 2 but got %d", p)
	}

	if p, _ := calculateStartPage(config, 220); p != 2 {

		t.Errorf("expected page 2 but got %d", p)
	}
}
