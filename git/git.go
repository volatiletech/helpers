package git

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// GetAppPath executes the git cmd "git rev-parse --show-toplevel" to obtain
// the full path of the current app. The last folder in the path is the app name.
func GetAppPath() (string, error) {
	return getAppPath()
}

func getAppPath() (string, error) {
	gitCmd := exec.Command("git", "rev-parse", "--show-toplevel")

	b := &bytes.Buffer{}
	gitCmd.Stdout = b

	err := gitCmd.Run()
	if err != nil {
		return "", errors.Wrap(err, "cannot find app root dir git rev-parse failed")
	}

	output := b.String()

	if len(output) == 0 {
		return "", errors.New("cannot find app root dir git rev-parse had no output")
	}

	return strings.TrimSpace(output), nil
}

// GetAppName executes the git cmd "git rev-parse --show-toplevel" to obtain
// the app name of the current app. The last folder in the path is the app name.
func GetAppName() (string, error) {
	path, err := getAppPath()
	if err != nil {
		return "", err
	}

	return getAppName(path), err
}

// GetAppEnvName executes the git cmd "git rev-parse --show-toplevel" to obtain
// the app name of the current app in environment mode format.
// The last folder in the path is the app name.
// For example, "my-app" converts to "MY_APP".
func GetAppEnvName() (string, error) {
	path, err := getAppPath()
	if err != nil {
		return "", err
	}

	return strings.ToUpper(replaceNonAlpha(getAppName(path), '_')), nil
}

// getAppName gets the appname portion of a project path
func getAppName(appPath string) string {
	// Is "/" on both Windows and Linux
	split := strings.Split(appPath, "/")
	return split[len(split)-1]
}

// replaceNonAlpha replaces non alphabet characters with the replace byte.
func replaceNonAlpha(s string, replace byte) string {
	byts := []byte(s)
	newByts := make([]byte, len(byts))

	for i := 0; i < len(byts); i++ {
		if byts[i] < 'A' || (byts[i] > 'Z' && byts[i] < 'a') || byts[i] > 'z' {
			newByts[i] = replace
		} else {
			newByts[i] = byts[i]
		}
	}

	return string(newByts)
}
