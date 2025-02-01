package terminal

import (
	"os"
	"os/exec"
	"runtime"
)

// Clear löscht das Terminal unabhängig vom Betriebssystem.
func Clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
