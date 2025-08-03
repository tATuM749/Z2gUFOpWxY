// 代码生成时间: 2025-08-03 14:47:15
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// FolderStructureOrganizer is the main struct that holds the configuration for the organizer.
type FolderStructureOrganizer struct {
    RootPath string
}

// NewFolderStructureOrganizer creates a new instance of FolderStructureOrganizer.
func NewFolderStructureOrganizer(rootPath string) *FolderStructureOrganizer {
    return &FolderStructureOrganizer{
        RootPath: rootPath,
    }
}

// Organize scans the root path and organizes the folders based on a predefined structure.
func (fso *FolderStructureOrganizer) Organize() error {
    // Check if the root path exists
    if _, err := os.Stat(fso.RootPath); os.IsNotExist(err) {
        return fmt.Errorf("root path does not exist: %w", err)
    }

    // Create a map to hold the desired folder structure
    desiredStructure := map[string][]string{
        "Documents": []string{
            "Financial", "Legal", "Personal",\        },
        "Pictures": []string{
            "Holidays", "Family", "Events",
        },
        "Projects": []string{}, // This can be filled dynamically based on some criteria
    }

    // Iterate over the desired structure and create folders if they don't exist
    for parent, children := range desiredStructure {
        // Create parent folder if it doesn't exist
        if err := os.MkdirAll(filepath.Join(fso.RootPath, parent), 0755); err != nil {
            return fmt.Errorf("failed to create parent folder %s: %w", parent, err)
        }

        // Create child folders if they don't exist
        for _, child := range children {
            if err := os.MkdirAll(filepath.Join(fso.RootPath, parent, child), 0755); err != nil {
                return fmt.Errorf("failed to create child folder %s: %w", child, err)
            }
        }
    }

    return nil
}

// MoveFiles attempts to move files from the root path to the newly created folders based on file extensions.
func (fso *FolderStructureOrganizer) MoveFiles() error {
    // Read all files from the root path
    files, err := os.ReadDir(fso.RootPath)
    if err != nil {
        return fmt.Errorf("failed to read files from root path: %w", err)
    }

    // Define a map of file extensions to folder names
    extensionToFolder := map[string]string{
        ".pdf": "Documents/Financial",
        ".txt": "Documents/Legal",
        ".jpg": "Pictures/Holidays",
        // Add more mappings as needed
    }

    // Iterate over the files and move them to the appropriate folders
    for _, file := range files {
        if file.IsDir() {
            continue
        }

        fileExtension := strings.TrimPrefix(strings.ToLower(filepath.Ext(file.Name())), ".")
        targetFolder, ok := extensionToFolder[fileExtension]
        if !ok {
            continue // Skip files that don't have a specified target folder
        }

        // Construct the source and target file paths
        srcPath := filepath.Join(fso.RootPath, file.Name())
        dstPath := filepath.Join(fso.RootPath, targetFolder, file.Name())

        // Create the target folder if it doesn't exist
        if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
            return fmt.Errorf("failed to create target folder: %w", err)
        }

        // Move the file to the target folder
        if err := os.Rename(srcPath, dstPath); err != nil {
            return fmt.Errorf("failed to move file %s: %w", file.Name(), err)
        }
    }

    return nil
}

func main() {
    // Define the root path to organize
    rootPath := "/path/to/your/folder"

    // Create a new folder structure organizer
    organizer := NewFolderStructureOrganizer(rootPath)

    // Organize the folders
    if err := organizer.Organize(); err != nil {
        log.Fatalf("failed to organize folders: %v", err)
    }

    // Move files to the newly created folders
    if err := organizer.MoveFiles(); err != nil {
        log.Fatalf("failed to move files: %v", err)
    }

    fmt.Println("Folder structure organized successfully!")
}