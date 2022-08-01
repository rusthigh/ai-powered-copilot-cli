
package tools

import (
	"strings"
	"testing"
)

func TestPythonREPL(t *testing.T) {
	type args struct {
		script string
	}
	tests := []struct {
		name    string