package files

import (
	"bufio"
	"log"
	"os"
)

func ReadLinesFromFile(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("error: ", err.Error())
		return nil
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
