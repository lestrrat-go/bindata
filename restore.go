// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package bindata

import (
	"fmt"
	"io"
)

func writeRestore(w io.Writer) error {
	_, err := fmt.Fprintf(w, `
// _RestoreAsset restores an asset under the given directory
func _RestoreAsset(dir, name string) error {
	data, err := _Asset(name)
	if err != nil {
		return err
	}
	info, err := _AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// _RestoreAssets restores an asset under the given directory recursively
func _RestoreAssets(dir, name string) error {
	children, err := _AssetDir(name)
	// File
	if err != nil {
		return _RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = _RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
`)
	return err
}
