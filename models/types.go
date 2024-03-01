package models

type Email struct {
	// Subject   string `json:"subject"`
	// From      string `json:"from"`
	// To        string `json:"to"`
	Content string `json:"content"`
	// MessageID string `json:"message_id"`
	// Date      string `json:"date"`
}

type CreateIndexer struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}

// Solicitud de busqueda
type SearchQuery struct {
	SearchType   string   `json:"search_type"`
	Query        Query    `json:"query"`
	SortFields   []string `json:"sort_fields"`
	From         int      `json:"from"`
	MaxResults   int      `json:"max_results"`
	SourceFields []string `json:"_source"`
}

// Parámetros de la consulta de búsqueda
type Query struct {
	Term      string `json:"term"`
	Field     string `json:"field"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// Búsqueda
type SearchResponse struct {
	ResultCount       int    `json:"result_count"`
	Success           bool   `json:"success"`
	TotalResultsCount int    `json:"total_results_count"`
	Time              string `json:"time"`
	Hits              Hits   `json:"hits"`
}

// Resultados de la búsqueda
type Hits struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`
	Hits []Hit `json:"hits"`
}

// Resultado individual de la búsqueda
type Hit struct {
	Index     string      `json:"_index"`
	Type      string      `json:"_type"`
	ID        string      `json:"_id"`
	Score     float64     `json:"_score"`
	Timestamp string      `json:"@timestamp"`
	Source    interface{} `json:"_source"`
}
