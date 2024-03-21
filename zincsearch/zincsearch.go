package zincsearch

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"final-project/models"
)

func PostZs(records []models.Email) {
	// URL de la API externa
	apiURL := "http://localhost:4080/api/_bulkv2"

	body := models.CreateIndexer{
		Index:   "Enron",
		Records: records,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	log.Printf("JSON Data:\n%s\n", jsonData)

	// Crear una solicitud POST
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		return
	}

	user := "admin"
	password := "Complexpass#123"

	// Codificar las credenciales en base64
	auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))

	// Configurar encabezados de la solicitud según sea necesario
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+auth)

	// Cliente HTTP personalizado
	client := &http.Client{}

	// Realizar la solicitud POST
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta de la API
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta de la API:", err)
		return
	}

	// Imprimir la respuesta en la consola
	fmt.Println("Respuesta de la API:", string(responseBody))
}

// SearchZs realiza una búsqueda en la API de Zinc y devuelve los resultados
func SearchZs(searchTerm string) (models.SearchResponse, error) {

	// Realizar búsqueda en la API de Zinc después de procesar los correos electrónicos
	searchQuery := models.SearchQuery{
		SearchType: "match_phrase",
		Query: models.Query{
			Term:      searchTerm,
			StartTime: "2021-06-02T14:28:31.894Z",
			EndTime:   "2025-12-02T15:28:31.894Z",
		},

		From:       0,
		MaxResults: 20,
	}

	// URL de la API de Zinc
	apiURL := "http://localhost:4080/api/Enron/_search"

	// Convertir la consulta a formato JSON
	jsonData, err := json.Marshal(searchQuery)
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error al convertir la consulta a JSON: %v", err)
	}

	// Crear una solicitud POST con el cuerpo de la consulta JSON
	req, err := http.NewRequest("POST", apiURL, bytes.NewReader(jsonData))
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error al crear la solicitud: %v", err)
	}

	user := "admin"
	password := "Complexpass#123"

	// Codificar las credenciales en base64
	auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))

	// Configurar encabezados de la solicitud según sea necesario
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+auth)

	// Cliente HTTP personalizado (ajusta el timeout según tus necesidades)
	client := &http.Client{}

	// Realizar la solicitud POST
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("Error al realizar la solicitud:", err)
		return models.SearchResponse{}, fmt.Errorf("error al realizar la solicitud: %v", err)
	}
	defer resp.Body.Close()

	// Leer la respuesta de la API
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error al leer la respuesta de la API: %v", err)
	}

	// Crear una estructura de respuesta y deserializar la respuesta JSON
	var response models.SearchResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return models.SearchResponse{}, fmt.Errorf("error al deserializar la respuesta JSON: %v", err)
	}

	return response, nil
}
