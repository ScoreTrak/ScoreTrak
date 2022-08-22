package util

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"strings"
)

func TarRecurse(path string, tarWriter *tar.Writer, skipMatch string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if strings.Contains(file.Name(), skipMatch) {
		return nil
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.IsDir() {
		rdr, err := file.Readdir(0)
		if err != nil {
			return err
		}
		for _, fi := range rdr {
			dir := fmt.Sprintf("%s/%s", path, fi.Name())
			err := TarRecurse(dir, tarWriter, skipMatch)
			if err != nil {
				return err
			}
		}
	} else {
		err := TarWrite(path, tarWriter, fileInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func TarWrite(_path string, tw *tar.Writer, fi os.FileInfo) error {
	file, err := os.Open(_path)
	if err != nil {
		return err
	}
	defer file.Close()
	tarHeader := &tar.Header{
		Name: _path,
		Size: fi.Size(),
		Mode: int64(fi.Mode()),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		return err
	}
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}
	return nil
}
