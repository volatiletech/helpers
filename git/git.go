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
