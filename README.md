# 🦫 gwatcher

A lightweight Go file watcher that automatically rebuilds or reruns your Go project (like nodemon, but for Go).

---

## 🚀 Features

- Automatically watches `.go` files in your project  
- Runs any command (default: `go test ./...`) on save  
- Simple, fast, and dependency-light  
- Supports debounce delay and custom watch directories  
- Works on macOS, Linux, and Windows

---

## 🧩 Installation

### Requirements
- Go 1.18+

### Install
```bash
go install github.com/ayushdoesdev/gwatcher/cmd/gwatcher@latest
```

Make sure your Go bin directory is in your PATH:
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Verify installation:
```bash
gwatcher --help
```

## ⚙️ Usage
```bash
# Watch current directory and run tests on change
gwatcher

# Watch and run main.go automatically
gwatcher --cmd "go run main.go"

# Watch a specific directory
gwatcher --dir ./internal --cmd "go test ./internal/..."

# Set debounce delay (milliseconds)
gwatcher --delay 1000 --cmd "go run main.go"

```

Example output

```bash
👀 Watching ./ | Running: go test ./...
🧱 Change detected: main.go
ok  	example.com/app	0.45s

```


---

## 🛠 Flags

| Flag | Default | Description |
|------|----------|-------------|
| `--cmd`, `-c` | `go test ./...` | Command to run when a `.go` file changes |
| `--dir`, `-d` | `.` | Directory to watch recursively |
| `--delay`, `-t` | `500` | Delay in milliseconds before re-running command |

---

## 💻 Example Workflow

```bash
# Start watcher
gwatcher --cmd "go run main.go"

# Edit your Go code
# gwatcher automatically rebuilds and re-runs your app!
```

## 🤝 Contributing

Contributions are welcome!

1. **Fork** the repository  
2. **Create a feature branch:**
```bash
   git checkout -b feature/your-feature
```
3. Make your changes
4. Run formatting and vet:
```bash
  go fmt ./...
  go vet ./...
```
5. Commit and open a Pull Request

