package internal

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/magefile/mage/sh"
)

// CleanDownloads removes files and dirs modified more than 30 days ago from ~/downloads
// TODO test this function
func CleanDownloads() {
	retentionTime := 30 * 24 * time.Hour
	log.Println("Retention time:", retentionTime.Hours(), "hours")

	hasTrashCLI := checkCLI("trash-put")
	count := 0

	err := os.Chdir(filepath.Join(getHome(), "downloads"))
	must(err)

	log.Println("Cleaning downloads folder...")
	files, err := os.ReadDir(".")
	must(err)

	for _, file := range files {
		info, err := file.Info()
		must(err)

		// log.Println("File", file.Name(), "modified", time.Since(info.ModTime()).Hours(), "hours ago")

		if time.Since(info.ModTime()).Hours() > retentionTime.Hours() {
			log.Println("Removing: ", file.Name())
			if hasTrashCLI {
				err := sh.Run("trash-put", file.Name())
				must(err)
			} else {
				err := os.Rename(file.Name(), filepath.Join(getTrashDir(), file.Name()))
				must(err)
			}
			count++
		}
	}
	log.Println("Cleaning downloads folder... done")
	log.Println("Removed", count, "files/dirs")
}
