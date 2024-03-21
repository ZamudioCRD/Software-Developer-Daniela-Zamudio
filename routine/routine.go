package routine

import (
	"bufio"
	"final-project/models"
	"final-project/zincsearch"
	"log"
	"os"
	"strings"
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

	var (
		email        models.Email
		emailContent string
		subjectFound bool
		fromFound    bool
		toFound      bool
		messageFound bool
		dateFound    bool
	)

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

		switch {

		case strings.HasPrefix(line, "Message-ID:") && !messageFound:
			email.MessageID = strings.TrimSpace(line[12:])
			messageFound = true

		case strings.HasPrefix(line, "Date:") && !dateFound:
			email.Date = strings.TrimSpace(line[5:])
			dateFound = true

		case strings.HasPrefix(line, "From:") && !fromFound:
			email.From = strings.TrimSpace(line[5:])
			fromFound = true

		case strings.HasPrefix(line, "To:") && !toFound:
			email.To = strings.TrimSpace(line[3:])
			toFound = true

		case strings.HasPrefix(line, "Subject:") && !subjectFound:
			email.Subject = strings.TrimSpace(line[8:])
			subjectFound = true

		default:
			emailContent += line
		}
	}

	//log.Printf("Email Content:\n%s\n", emailContent)

	email.Content = emailContent
	*emails = append(*emails, email)

	m.Lock()
	if len(*emails) >= batchSize {
		go zincsearch.PostZs(*emails)
		*emails = nil
	}
	m.Unlock()
}
