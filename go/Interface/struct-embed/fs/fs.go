package main

import "fmt"

// 定义一个文件系统条目接口
type FileSystemItem interface {
	Info() string
}

// File struct
type File struct {
	filename string
	size     int64
}

// impl FileSystemItem Interface in File
func (f File) Info() string {
	return fmt.Sprintf("File: %s, Size: %d bytes", f.filename, f.size)
}

func (f File) IsLarge() bool {
	return f.size > 1024
}

// Folder struct
type Folder struct {
	contents []FileSystemItem // 包含的文件或文件夹
	name     string
}

// impl FileSystemItem interface in Folder
func (f Folder) Info() string {
	return fmt.Sprintf("Folder: %s", f.name)
}

// append file or folder into folder
func (f *Folder) Add(item FileSystemItem) {
	f.contents = append(f.contents, item)
}

func main() {
	// create file
	file1 := File{filename: "file1.txt", size: 512}
	file2 := File{filename: "file2.txt", size: 2028}

	// create folder, append to file
	folder1 := Folder{name: "Folder1"}
	folder1.Add(file1)
	folder1.Add(file2)

	// create another folder
	folder2 := Folder{name: "Folder2"}
	folder2.Add(File{filename: "file3.txt", size: 3072})
	folder2.Add(Folder{name: "Subfolder"})

	// create root folder
	rootFolder := Folder{name: "Root"}
	rootFolder.Add(folder1)
	rootFolder.Add(folder2)

	// print fs structure
	printFS(rootFolder, 0)
}

func printFS(item FileSystemItem, indentLevel int) {
	for i := 0; i < indentLevel; i++ {
		fmt.Print("  ")
	}
	fmt.Println(item.Info())

	if folder, ok := item.(Folder); ok {
		for _, subItem := range folder.contents {
			printFS(subItem, indentLevel+1)
		}
	}
}
