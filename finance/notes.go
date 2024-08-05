package finance

import (
	"context"
	"os"
	"strings"
)

type NotesResolver struct{}

func (r *NotesResolver) Notes(ctx context.Context) ([]string, error) {
	data, err := os.ReadFile("path/to/your/markdown/file.md")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	var notes []string
	for _, line := range lines {
		if strings.HasPrefix(line, "- ") {
			notes = append(notes, strings.TrimPrefix(line, "- "))
		}
	}

	return notes, nil
}
