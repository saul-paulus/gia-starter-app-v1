package cli

import "os"

func createFile(path string, content string) {
	if _, err := os.Stat(path); err == nil {
		return // skip kalau sudah ada
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(content)
}
