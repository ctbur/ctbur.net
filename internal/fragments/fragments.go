package fragments

import (
	"bytes"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/yuin/goldmark"
)

type FragmentsFile struct {
	Fragments []Fragment `toml:"fragments"`
}

type Fragment struct {
	ID      uint      `toml:"id"`
	Date    time.Time `toml:"date"`
	Title   string    `toml:"title"`
	Content string    `toml:"content"`
	Tags    []string  `toml:"tags"`
}

func LoadFragments(fragmentsFilePath string, markdown goldmark.Markdown) ([]Fragment, error) {
	var ff FragmentsFile
	toml.DecodeFile(fragmentsFilePath, &ff)

	// Ensure there is at least one entry
	if len(ff.Fragments) == 0 {
		return nil, fmt.Errorf("No fragments found in file '%s'", fragmentsFilePath)
	}

	// Ensure that IDs are unique
	ids := make(map[uint]struct{}, len(ff.Fragments))
	for _, fragment := range ff.Fragments {
		if _, idExists := ids[fragment.ID]; idExists {
			return nil, fmt.Errorf("Fragment ID %d occurs multiple times", fragment.ID)
		}
	}

	for idx, _ := range ff.Fragments {
		fragment := &ff.Fragments[idx]

		// Ensure that each entry has all properties set
		if fragment.Date.IsZero() {
			return nil, fmt.Errorf("Fragment %d has no date", fragment.ID)
		}
		if fragment.Title == "" {
			return nil, fmt.Errorf("Fragment %d has no title", fragment.ID)
		}
		if fragment.Content == "" {
			return nil, fmt.Errorf("Fragment %d has no content", fragment.ID)
		}
		if len(fragment.Tags) == 0 {
			return nil, fmt.Errorf("Fragment %d has no tags", fragment.ID)
		}

		// Convert markdown content
		var buf bytes.Buffer
		if err := markdown.Convert([]byte(fragment.Content), &buf); err != nil {
			return nil, fmt.Errorf("Fragment markdown parse failed: %w", err)
		}
		fragment.Content = buf.String()
	}

	return ff.Fragments, nil
}
