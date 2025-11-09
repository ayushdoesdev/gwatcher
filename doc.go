// Package gwatcher is a lightweight file watcher for Go projects.
//
// gwatcher watches .go files and runs a command when files change.
// It's minimal and intended to be used during development to rebuild,
// re-run tests, or re-run your app automatically.
//
// Usage
//
//   // run default action (go test ./...)
//   gwatcher
//
//   // run a custom command
//   gwatcher --cmd "go run main.go"
//
// The command, directory, and debounce delay are configurable via flags.
//
// See the README for full examples and installation.
package gwatcher
