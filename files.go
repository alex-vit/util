package util

import "os"

func Touch(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		f, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		return f.Close()
	} else {
		return err
	}
}
