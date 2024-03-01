package routine

import (
	"bufio"
	"final-project/models"
	"final-project/zincsearch"
	"log"
	"os"
	"sync"
)

func ReadFileWithBufio(filePath string, wg *sync.WaitGroup, emails *[]models.Email, batchSize int) {

	var m sync.Mutex
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error al abrir el archivo %s: %s\n", filePath, err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var emailContent string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {

			if err.Error() == "EOF" {
				break
			} else {
				log.Printf("Error al leer el archivo %s: %s\n", filePath, err)
				break
			}
		}

		emailContent += line

	}

	email := models.Email{
		Content: emailContent,
	}

	*emails = append(*emails, email)

	m.Lock()
	if len(*emails) >= batchSize {
		go zincsearch.PostZs(*emails)
		*emails = nil
	}
	m.Unlock()
}
