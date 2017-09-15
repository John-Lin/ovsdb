package ovsdb

import (
	"os/exec"
	"strings"
)

func ovsdbUnixPath() string {
	var path string
	cmd := "ps aux | awk '{print $12}' | grep -Eo unix:.*openvswitch.*sock"
	out, _ := exec.Command("sh", "-c", cmd).Output()
	if len(out) != 0 {
		m := strings.Split(string(out[:]), "unix:")
		path = m[len(m)-1]
	}
	return path
}
