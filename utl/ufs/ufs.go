package ufs

import (
	"embed"
	"os"
	"path"
	"strings"
)

// return list of fs files under a directory recursively
func GetFSFilesRecursively(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}

	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() {
			res, err := GetFSFilesRecursively(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		out = append(out, fp)
	}
	return
}

// get a list of files of some extension, in a directory (not recursive)
// the "x" paramater must have the "." example: .json.
func GetFilesOfExtension(dir, x string) ([]string, error) {

	// Open the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var f []string

	// Iterate through the files in the directory
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), x) {
			f = append(f, file.Name())
		}
	}
	return f, nil
}
