package git

import (
	"strings"
	"testing"
)

func TestGetAppPath(t *testing.T) {
	path, err := getAppPath()
	if err != nil {
		t.Error(err)
	}
	if !strings.HasSuffix(path, "volatiletech/helpers") {
		t.Error("path not correct value")
	}
}

func TestGetAppName(t *testing.T) {
	name := GetAppName()
	if name != "helpers" {
		t.Errorf("expected helpers, got %s", name)
	}
}

func TestGetAppEnvName(t *testing.T) {
	name := GetAppEnvName()
	if name != "HELPERS" {
		t.Errorf("expected HELPERS, got %s", name)
	}
}
