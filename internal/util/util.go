package util

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func IsExistsFile(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DownloadFile(URL, fileName string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}

	defer func() {
		err = response.Body.Close()
	}()

	if response.StatusCode != 200 {
		return errors.New("received none 200 response code")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
	}()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return err
}