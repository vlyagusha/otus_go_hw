package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var fileRead, fileWrite *os.File
	var err error

	fileRead, err = os.OpenFile(fromPath, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}
	defer fileRead.Close()

	fileInfo, err := fileRead.Stat()
	if err != nil {
		return err
	}

	if (fileInfo.Mode() & os.ModeType) != 0 {
		return ErrUnsupportedFile
	}
	if fileInfo.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	fileWrite, err = os.Create(toPath)
	if err != nil {
		return err
	}
	defer fileWrite.Close()

	buf := make([]byte, 1024)
	_, err = fileRead.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	for {
		read, err := fileRead.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if limit > 0 && int64(read) >= limit {
			_, err := fileWrite.Write(buf[:limit])
			if err != nil {
				return err
			}
			break
		}
		limit -= int64(read)

		_, err = fileWrite.Write(buf[:read])
		if err != nil {
			return err
		}
	}

	return nil
}
