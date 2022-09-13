package tools

import (
	"os/exec"
	"strings"
)

// Trivy runs trivy against the image and returns the output
func Trivy(image string) (string, error) {
	im