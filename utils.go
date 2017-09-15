package ovsdb

import (
	"os/exec"
	"strings"
)

func ovsdbUnixPath() string {
	var path string
	cmd := "ps aux | awk '{print $12}' | grep -E \"unix:.*sock\""
	out, _ := exec.Command("sh", "-c", cmd).Output()
	if len(out) != 0 {
		path = strings.Split(string(out[:]), "unix:")[1]
	}
	return path
}
