# README

Example implementation of a simple CLI tools that has sub-commands. The tool is implemented using the `cobra` package.

https://pkg.go.dev/github.com/spf13/cobra

Each sub-command performs a common setup by using `PersistentPreRun`.

## Reproduce

Prerequisites:

- Go 1.22
- cobra-cli has been installed

```bash
$ go mod init my-module

$ cobra-cli init # add ./main.go and ./cmd/root.go

$ cobra-cli add -p rootCmd subcmd1 # add ./cmd/subcmd1.go

$ cobra-cli add -p rootCmd subcmd2 # add ./cmd/subcmd2.go

$ cobra-cli add -p subcmd2Cmd nested # add ./cmd/nested.go
```
