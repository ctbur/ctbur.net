package fragments

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
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

func LoadFragments(fragmentsFilePath string) ([]Fragment, error) {
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

	// Ensure that each entry has all properties set
	for _, fragment := range ff.Fragments {
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
	}

	return ff.Fragments, nil
}
