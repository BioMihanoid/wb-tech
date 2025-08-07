package files

import (
	"bufio"
	"log"
	"os"
)

func WriteLinesToFile(lines []string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Println("error: ", err.Error())
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	writer := bufio.NewWriter(f)

	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			log.Println("error: ", err.Error())
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		log.Println("error: ", err.Error())
	}
}
