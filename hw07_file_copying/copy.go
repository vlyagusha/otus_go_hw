package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var fileReader, fileWriter *os.File
	var err error

	fileReader, err = os.OpenFile(fromPath, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}
	defer fileReader.Close()

	stat, err := fileReader.Stat()
	if err != nil {
		return err
	}

	if (stat.Mode() & os.ModeType) != 0 {
		return ErrUnsupportedFile
	}
	if stat.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	bytesToRead := stat.Size() - offset
	if limit == 0 || limit > bytesToRead {
		limit = bytesToRead
	}

	_, err = fileReader.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	fileWriter, err = os.Create(toPath)
	if err != nil {
		return err
	}
	defer fileWriter.Close()

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(fileReader)
	defer bar.Finish()

	_, err = io.CopyN(fileWriter, barReader, limit)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	return nil
}
