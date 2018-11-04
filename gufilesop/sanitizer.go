package gufilesop

import "path/filepath"

func NormalizePath(path string) string {
	pointer, err := filepath.EvalSymlinks(path)
	if err != nil {
		return path
	}
	abs, err := filepath.Abs(pointer)
	if err != nil {
		return pointer
	}
	return abs
}
