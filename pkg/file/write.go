package file

import (
	"bufio"
	"log"
	"os"
)

//WriteFile is to write file
func WriteFile(filename string, fileContext string) {
	sumfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer sumfile.Close()

	writer := bufio.NewWriter(sumfile)
	writer.WriteString(fileContext)
	writer.Flush()
}
