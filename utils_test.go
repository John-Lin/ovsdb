package ovsdb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestovsdbUnixPath(t *testing.T) {
	path := ovsdbUnixPath()
	assert.Equal(t, path, "punix:/usr/local/var/run/openvswitch/db.sock", "Those path should be same")
}
