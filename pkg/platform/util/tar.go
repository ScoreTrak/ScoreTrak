package util

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"strings"
)

func TarRecurse(path string, tw *tar.Writer, skipMatch string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if strings.Contains(file.Name(), skipMatch) {
		return nil
	}
	s, err := file.Stat()
	if err != nil {
		return err
	}
	if s.IsDir() {
		rdr, err := file.Readdir(0)
		if err != nil {
			return err
		}
		for _, fi := range rdr {
			dir := fmt.Sprintf("%s/%s", path, fi.Name())
			err := TarRecurse(dir, tw, skipMatch)
			if err != nil {
				return err
			}
		}
	} else {
		err := TarWrite(path, tw, s)
		if err != nil {
			return err
		}
	}
	return nil
}

func TarWrite(_path string, tw *tar.Writer, fi os.FileInfo) error {
	fr, err := os.Open(_path)
	if err != nil {
		return err
	}
	defer fr.Close()
	tarHeader := &tar.Header{
		Name: _path,
		Size: fi.Size(),
		Mode: int64(fi.Mode()),
	}
	err = tw.WriteHeader(tarHeader)
	if err != nil {
		return err
	}
	_, err = io.Copy(tw, fr)
	if err != nil {
		return err
	}
	return nil
}
