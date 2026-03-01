package main

import (
	"bufio"
	"os"
	"strings"
	"syscall"
)

// getDiskInfo works on Linux using Statfs
func getDiskInfo(path string) (diskInfo, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return diskInfo{}, err
	}

	blockSize := uint64(stat.Bsize)

	return diskInfo{
		Path:  path,
		Total: stat.Blocks * blockSize,
		Free:  stat.Bfree * blockSize,
	}, nil
}

// getMountedFilesystems reads /proc/mounts
func getMountedFilesystems() ([]string, error) {
	file, err := os.Open("/proc/mounts")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var paths []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) >= 2 {
			paths = append(paths, parts[1])
		}
	}

	return paths, scanner.Err()
}