// Package directory_tree provides a way to generate a directory tree.
//
// Example usage:
//
//	tree, err := directory_tree.NewTree("/home/me")
//
// I did my best to keep it OS-independent but truth be told I only tested it
// on OS X and Debian Linux so YMMV. You've been warned.
package dt

import (
	"archive/zip"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// FileInfo is a struct created from os.FileInfo interface for serialization.
type FileInfo struct {
	Name    string      `json:"name"`
	Size    int64       `json:"size"`
	Mode    os.FileMode `json:"mode"`
	ModTime time.Time   `json:"mod_time"`
	IsDir   bool        `json:"is_dir"`
}

// Helper function to create a local FileInfo struct from os.FileInfo interface.
func fileInfoFromInterface(v os.FileInfo) *FileInfo {
	return &FileInfo{v.Name(), v.Size(), v.Mode(), v.ModTime(), v.IsDir()}
}

// Node represents a node in a directory tree.
type Node struct {
	FullPath string    `json:"path"`
	Info     *FileInfo `json:"info"`
	Children []*Node   `json:"children"`
	Parent   *Node     `json:"-"`
}

// func walk(parent string, walkFunc func(path string, info os.FileInfo, err error) error) error {

// 	return nil
// }

type WalkFunc func(path string, info fs.FileInfo, err error) error

var lstat = os.Lstat // for testing
var SkipDir error = fs.SkipDir

func walk(path string, info fs.FileInfo, walkFn WalkFunc) error {
	if !info.IsDir() {
		return walkFn(path, info, nil)
	}

	names, err := readDirNames(path)
	err1 := walkFn(path, info, err)
	// If err != nil, walk can't walk into this directory.
	// err1 != nil means walkFn want walk to skip this directory or stop walking.
	// Therefore, if one of err and err1 isn't nil, walk will return.
	if err != nil || err1 != nil {
		// The caller's behavior is controlled by the return value, which is decided
		// by walkFn. walkFn may ignore err and return nil.
		// If walkFn returns SkipDir, it will be handled by the caller.
		// So walk should return whatever walkFn returns.
		return err1
	}

	for _, name := range names {
		filename := filepath.Join(path, name)
		fileInfo, err := lstat(filename)
		if err != nil {
			if err := walkFn(filename, fileInfo, err); err != nil && err != SkipDir {
				return err
			}
		} else {
			err = walk(filename, fileInfo, walkFn)
			if err != nil {
				if !fileInfo.IsDir() || err != SkipDir {
					return err
				}
			}
		}
	}
	return nil
}

func Walk(root string, walkFunc WalkFunc) error {
	zipReader, err := zip.OpenReader(root)
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	for _, file := range zipReader.Reader.File {
		walkFunc(file.Name, file.FileInfo(), err)
	}

	rootInfo, err := os.Stat(root)
	walkFunc(".", rootInfo, err)

	// TODO the root itself
	// walkFunc("", nil, err)

	// info, err := os.Lstat(root)
	// if err != nil {
	// 	err = fn(root, nil, err)
	// } else {
	// 	err = walk(root, info, fn)
	// }
	// if err == SkipDir {
	// 	return nil
	// }
	return err
}

// readDirNames reads the directory named by dirname and returns
// a sorted list of directory entry names.
func readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}

	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Strings(names)
	return names, nil
}

func NewZipTree(root string) (result *Node, err error) {
	parents := make(map[string]*Node)
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		path = strings.TrimSuffix(path, "/")
		parents[path] = &Node{
			FullPath: path,
			Info:     fileInfoFromInterface(info),
			Children: make([]*Node, 0),
		}
		return nil
	}

	if err = Walk(root, walkFunc); err != nil {
		return
	}

	for path, node := range parents {
		if path == "." {
			result = node
			continue
		}
		parentPath := filepath.Dir(path)
		parentPath = filepath.ToSlash(parentPath)
		parent, exists := parents[parentPath]
		if !exists {
			err = fmt.Errorf("Directory doesn't exist: %s", parentPath)
			return
		}
		node.Parent = parent
		parent.Children = append(parent.Children, node)

	}

	return
}

// Create directory hierarchy.
func NewTree(root string) (result *Node, err error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return
	}
	parents := make(map[string]*Node)
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		parents[path] = &Node{
			FullPath: path,
			Info:     fileInfoFromInterface(info),
			Children: make([]*Node, 0),
		}
		return nil
	}
	if err = filepath.Walk(absRoot, walkFunc); err != nil {
		return
	}
	for path, node := range parents {
		parentPath := filepath.Dir(path)
		parent, exists := parents[parentPath]
		if !exists { // If a parent does not exist, this is the root.
			result = node
		} else {
			node.Parent = parent
			parent.Children = append(parent.Children, node)
		}
	}
	return
}
