package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

func must(err error) {
	if err != nil {
		log.Panic(err)
		return
	}
}

func getHome() string {
	home, err := os.UserHomeDir()
	must(err)
	return home
}

func checkCLI(cmd string) bool {
	_, err := exec.Command("which", cmd).Output()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			log.Println("Command not found:", cmd)
			return false
		}
		log.Println("Error while finding cmd:", err)
		return false
	}
	log.Println("Command found:", cmd)
	return true
}

func getTrashDir() string {
	home := getHome()
	return filepath.Join(home, ".local/share/Trash/files")
}

func getStats() {
	// Gets stats of the file
	stats, err := os.Stat("bb")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %s\n", stats.Name())
	fmt.Printf("Modification Time: %s\n", stats.ModTime())
	atime := stats.Sys().(*syscall.Stat_t).Atim
	fmt.Println(time.Unix(atime.Sec, atime.Nsec))
}
