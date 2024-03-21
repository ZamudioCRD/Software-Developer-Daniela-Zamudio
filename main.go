package main

import (
	"final-project/server"
)

func main() {

	// var wg sync.WaitGroup
	// var batchSize = 50000
	// var emails []models.Email

	// //mailPath := "./enron_mail_20110402"
	// mailPath := "./enron_mail"

	// err := filepath.Walk(mailPath,

	// 	func(path string, info os.FileInfo, err error) error {
	// 		if err != nil {
	// 			return err
	// 		}

	// 		if info.IsDir() {
	// 			//fmt.Printf("Directorio: %s\n", path)
	// 		} else {
	// 			//hola
	// 			wg.Add(1)
	// 			go routine.ReadFileWithBufio(path, &wg, &emails, batchSize)
	// 		}
	// 		return nil
	// 	})

	// if err != nil {
	// 	log.Println(err)
	// }

	// wg.Wait() // Espera a que ambas goroutines terminen antes de salir

	// if len(emails) > 0 {
	// 	zincsearch.PostZs(emails)
	// }

	server.ConfigureRoutes()
}
