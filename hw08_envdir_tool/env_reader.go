package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrIllegalCharInFileName = errors.New("illegal char in file name")
)

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	environment := make(Environment)
	for _, dirEntry := range dirEntries {
		if (dirEntry.Type() & os.ModeType) != 0 {
			return nil, ErrUnsupportedFile
		}

		if strings.Contains(dirEntry.Name(), "=") {
			return nil, ErrIllegalCharInFileName
		}

		value, err := getValue(filepath.Join(dir, dirEntry.Name()))
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}

		environment[dirEntry.Name()] = EnvValue{
			Value:      value,
			NeedRemove: value == "",
		}
	}

	return environment, nil
}

func getValue(filename string) (string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0o644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	line = bytes.ReplaceAll(line, []byte{0}, []byte("\n"))

	return strings.TrimRight(string(line), " \t\n"), nil
}
