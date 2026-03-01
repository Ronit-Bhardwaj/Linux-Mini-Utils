package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

// getDiskInfo uses syscall.Statfs, available in Go's standard library on macOS.
func getDiskInfo(path string) (diskInfo, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return diskInfo{}, err
	}
	// stat.Bsize is int32 on macOS, cast carefully
	blockSize := uint64(stat.Bsize)
	return diskInfo{
		Path:  path,
		Total: stat.Blocks * blockSize,
		Free:  stat.Bfree * blockSize,
	}, nil
}

// getMountedFilesystems shells out to `mount` and parses mount points.
func getMountedFilesystems() ([]string, error) {
	var out bytes.Buffer
	cmd := exec.Command("mount")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to run mount: %w", err)
	}

	var paths []string
	for _, line := range strings.Split(out.String(), "\n") {
		parts := strings.Fields(line)
		if len(parts) >= 3 && parts[1] == "on" {
			paths = append(paths, parts[2])
		}
	}
	return paths, nil
}