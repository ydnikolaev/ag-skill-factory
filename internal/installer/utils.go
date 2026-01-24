package installer

import (
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

// copyDir copies a directory recursively using the Installer's Fs.
func (i *Installer) copyDir(src, dst string) error {
	_ = i.Fs.RemoveAll(dst)

	return afero.Walk(i.Fs, src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return i.Fs.MkdirAll(dstPath, info.Mode())
		}

		return i.copyFile(path, dstPath)
	})
}

// copyFile copies a single file using the Installer's Fs.
func (i *Installer) copyFile(src, dst string) error {
	srcFile, err := i.Fs.Open(src)
	if err != nil {
		return err
	}
	defer func() { _ = srcFile.Close() }()

	dstFile, err := i.Fs.Create(dst)
	if err != nil {
		return err
	}
	defer func() { _ = dstFile.Close() }()

	_, err = io.Copy(dstFile, srcFile)

	return err
}
