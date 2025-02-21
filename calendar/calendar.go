package calendar

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/apognu/gocal"
)

// Event represents a single event from an iCal file
type Event struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`
}

// processICalFile parses an iCal file and returns a slice of Event structs
func ProcessICalFile(r io.Reader) ([]Event, error) {
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

func SaveToCSV(events []Event, filename string) error {
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
