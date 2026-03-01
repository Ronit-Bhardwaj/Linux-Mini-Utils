# Linux Mini-Utils (Go)

---

## mini-grep

Search for a regex pattern across files or stdin, with optional line numbers.

Usage:
./mini-grep-bin [-q] -e PATTERN [FILE...]

Flags:
  -e PATTERN   Regex pattern to search (required)
  -q           Quiet mode — omit line numbers

Examples:
  ./mini-grep-bin -e "error" app.log
  cat file.txt | ./mini-grep-bin -e "foo"

Build:
go build -o mini-grep-bin ./mini-grep

Run:
./mini-grep-bin -e "hello" somefile.txt

---

## mini-df

Displays total and free space for the file system containing each provided path.

If no PATH is given, all mounted file systems are displayed.

Usage:
./mini-df-bin [-h] [PATH...]

Flags:
  -h   Human-readable output (powers of 1024, e.g. 1.5G)

Examples:
  ./mini-df-bin
  ./mini-df-bin -h /home /tmp
  ./mini-df-bin /dev/shm

Build:
go build -o mini-df-bin ./mini-df

Run:
./mini-df-bin
./mini-df-bin -h

---

## Cross-Platform Compatibility

`mini-df` automatically detects mounted filesystems when no PATH is provided.

On Linux:
- Filesystem information is retrieved using system calls.

On macOS:
- Mounted filesystems are detected using the `mount` command.
- Disk statistics are retrieved using `syscall.Statfs`.

---

## Testing in Linux using Docker

These utilities can be tested in a real Linux environment using Docker.

Run:

docker run -it --rm -v "$PWD":/app -w /app golang:1.25 bash

Inside the container:

Build:

go build -o mini-grep-bin ./mini-grep
go build -o mini-df-bin ./mini-df

Run:

./mini-grep-bin -e "test" somefile.txt
./mini-df-bin -h

This ensures correct Linux behavior.

---

## Project Structure

```
Linux-mini-utils
├── mini-grep
│   └── main.go
├── mini-df
│   ├── main.go
│   ├── fs_linux.go
│   └── fs_darwin.go
├── go.mod
└── README.md
```

---

## Note

Please build locally before running:

go build -o mini-grep-bin ./mini-grep
go build -o mini-df-bin ./mini-df