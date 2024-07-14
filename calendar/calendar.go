package calendar

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/apognu/gocal"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

// Event represents a single event from an iCal file
type Event struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`
}

// processICalFile parses an iCal file and returns a slice of Event structs
func processICalFile(r io.Reader) ([]Event, error) {
	parser := gocal.NewParser(r)
	err := parser.Parse()
	if err != nil {
		return nil, err
	}

	var events []Event
	for _, e := range parser.Events {
		event := Event{
			Summary:     e.Summary,
			Description: e.Description,
			Start:       e.Start.String(),
			End:         e.End.String(),
		}
		events = append(events, event)
	}

	return events, nil
}

// saveToCSV saves a slice of Events to a CSV file
func saveToCSV(events []Event, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Summary", "Description", "Start", "End"}
	if err := writer.Write(headers); err != nil {
		return err
	}

	for _, event := range events {
		row := []string{event.Summary, event.Description, event.Start, event.End}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func calendar() {
	// Define the GraphQL schema
	eventType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Event",
		Fields: graphql.Fields{
			"summary":     &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"start":       &graphql.Field{Type: graphql.String},
			"end":         &graphql.Field{Type: graphql.String},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"processICalFile": &graphql.Field{
				Type: graphql.NewList(eventType),
				Args: graphql.FieldConfigArgument{
					"file": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					fileContent, ok := p.Args["file"].(string)
					if !ok {
						return nil, fmt.Errorf("invalid file content")
					}

					events, err := processICalFile(stringReader(fileContent))
					if err != nil {
						return nil, err
					}

					if err := saveToCSV(events, "events.csv"); err != nil {
						return nil, err
					}

					return events, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Set up Gin
	r := gin.Default()

	// Define the GraphQL endpoint
	r.POST("/graphql", func(c *gin.Context) {
		var request struct {
			Query     string                 `json:"query"`
			Variables map[string]interface{} `json:"variables"`
		}

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  request.Query,
			VariableValues: request.Variables,
		})

		c.JSON(http.StatusOK, result)
	})

	// Start the server
	log.Println("Server is running on http://localhost:8080/graphql")
	log.Fatal(r.Run(":8080"))
}

// stringReader is a custom type that implements the io.Reader interface
type stringReader string

// Read implements the io.Reader interface for stringReader
func (s stringReader) Read(p []byte) (n int, err error) {
	return copy(p, s), io.EOF
}
