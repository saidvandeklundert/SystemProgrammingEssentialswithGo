package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"io"
)

func main() {
	// joining file paths:
	dir := "home/klundert"
	file := "document.txt"

	fullPath := filepath.Join(dir, file)
	fmt.Println("Full path:", fullPath)

	// cleaning file paths:
	uncleanPath := "/home/user/../documents/file.txt"
	cleanPath := filepath.Clean(uncleanPath)
	fmt.Println("Cleaned path:", cleanPath)
	// splitting file paths:
	path := "/home/user/documents/myfile.txt"
	dir, file = filepath.Split(path)
	fmt.Println("Directory:", dir)
	fmt.Println("File:", file)
	// traverse directories:
	err := filepath.WalkDir("/home/klundert/SystemProgrammingEssentialswithGo", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		if d.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
	/*
		// Define the source file path.
		sourcePath := "/home/user/Documents/important_document.txt"
		// Define the symlink path.
		symlinkPath := "/home/user/Desktop/shortcut_to_document.txt"
		// Create the symlink.
		err := os.Symlink(sourcePath, symlinkPath)
		if err != nil {
		fmt.Printf("Error creating symlink: %v\n", err)
		return
		}
		fmt.Printf("Symlink created: %s -> %s\n", symlinkPath, sourcePath)
	*/
	createFile()
	checkFiles()
	// Define the path to the file or symlink you want to remove.
	filePath := "hello.txt"
	// Attempt to remove the file.
	err = os.Remove(filePath)
	if err != nil {
		fmt.Printf("Error removing the file: %v\n", err)
		return
	}
	fmt.Printf("File removed: %s\n", filePath)
	checkFiles()
	// find duplicates:
	duplicates, err := findDuplicateFiles("/home/klundert/SystemProgrammingEssentialswithGo")
	if err != nil {
		fmt.Println("Error finding duplicate files:", err)
		return
	}

	for hash, files := range duplicates {
		if len(files) > 1 {
			fmt.Printf("Duplicate Hash: %s\n", hash)
			for _, file := range files {
				fmt.Println("  -", file)
			}
		}
	}

	fmt.Println("fin")

}

func createFile() {
	// Create a file
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write "Hello, World!" to the file
	_, err = file.WriteString("Hello, World!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and written successfully.")
	// check dir size:
	dirSize, err := calculateDirSize("/home/klundert/SystemProgrammingEssentialswithGo")
	if err != nil {
		fmt.Println("Error checking dir size:", err)
		return
	}
	fmt.Println("Dir size in bytes:", dirSize)
}

func checkFiles() {
	cmd := exec.Command("ls", "-ltr")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func calculateDirSize(path string) (int64, error) {
	var size int64

	err := filepath.Walk(path, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fileInfo.IsDir() {
			size += fileInfo.Size()
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return size, nil
}

func computeFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func findDuplicateFiles(rootDir string) (map[string][]string, error) {
	duplicates := make(map[string][]string)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			hash, err := computeFileHash(path)
			if err != nil {
				return err
			}

			duplicates[hash] = append(duplicates[hash], path)
		}

		return nil
	})

	return duplicates, err
}
