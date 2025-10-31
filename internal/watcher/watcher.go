package watcher

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Start(dir, command string, delay int) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		var timer *time.Timer
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if isGoFile(event.Name) && (event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create) {
					if timer != nil {
						timer.Stop()
					}
					timer = time.AfterFunc(time.Duration(delay)*time.Millisecond, func() {
						clearScreen()
						fmt.Printf("🧱 Change detected: %s\n", filepath.Base(event.Name))
						runCommand(command)
					})
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Watcher error:", err)
			}
		}
	}()

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && !strings.Contains(path, "vendor") {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	runCommand(command) // Run once at start
	<-done
	return nil
}

func isGoFile(path string) bool {
	return strings.HasSuffix(path, ".go")
}

func runCommand(command string) {
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Command failed:", err)
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
