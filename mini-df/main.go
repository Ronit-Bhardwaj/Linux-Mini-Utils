package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type diskInfo struct {
	Path  string
	Total uint64
	Free  uint64
}

func humanReadable(bytes uint64) string {
	units := []string{"B", "K", "M", "G", "T", "P"}
	value := float64(bytes)
	unit := units[0]
	for i := 1; i < len(units); i++ {
		if value < 1024 {
			break
		}
		value /= 1024
		unit = units[i]
	}
	if unit == "B" {
		return fmt.Sprintf("%dB", bytes)
	}
	return fmt.Sprintf("%.4g%s", value, unit)
}

func main() {
	human := flag.Bool("h", false, "Human-readable output (powers of 1024)")
	flag.Parse()

	paths := flag.Args()

	if len(paths) == 0 {
		var err error
		paths, err = getMountedFilesystems()
		if err != nil {
			fmt.Fprintf(os.Stderr, "mini-df: could not read mounted filesystems: %v\n", err)
			os.Exit(1)
		}
	}

	if *human {
		fmt.Printf("%-40s  %10s  %10s\n", "Path", "Total", "Free")
	} else {
		fmt.Printf("%-40s  %20s  %20s\n", "Path", "Total (bytes)", "Free (bytes)")
	}
	fmt.Println(strings.Repeat("-", 76))

	for _, path := range paths {
		info, err := getDiskInfo(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mini-df: %s: %v\n", path, err)
			continue
		}
		if *human {
			fmt.Printf("%-40s  %10s  %10s\n", info.Path, humanReadable(info.Total), humanReadable(info.Free))
		} else {
			fmt.Printf("%-40s  %20d  %20d\n", info.Path, info.Total, info.Free)
		}
	}
}
