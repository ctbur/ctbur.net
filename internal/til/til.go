package til

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type TilsFile struct {
	Tils []Til `toml:"tils"`
}

type Til struct {
	ID      uint      `toml:"id"`
	Date    time.Time `toml:"date"`
	Title   string    `toml:"title"`
	Content string    `toml:"content"`
	Tags    []string  `toml:"tags"`
}

func LoadTils(tilsFilePath string) ([]Til, error) {
	var tf TilsFile
	toml.DecodeFile(tilsFilePath, &tf)

	// Ensure there is at least one TIL
	if len(tf.Tils) == 0 {
		return nil, fmt.Errorf("No TILs found in file '%s'", tilsFilePath)
	}

	// Ensure that IDs are unique
	ids := make(map[uint]struct{}, len(tf.Tils))
	for _, til := range tf.Tils {
		if _, idExists := ids[til.ID]; idExists {
			return nil, fmt.Errorf("TIL ID %d occurs multiple times", til.ID)
		}
	}

	// Ensure that each til all properties set
	for _, til := range tf.Tils {
		if til.Date.IsZero() {
			return nil, fmt.Errorf("TIL %d has no date", til.ID)
		}
		if til.Title == "" {
			return nil, fmt.Errorf("TIL %d has no title", til.ID)
		}
		if til.Content == "" {
			return nil, fmt.Errorf("TIL %d has no content", til.ID)
		}
		if len(til.Tags) == 0 {
			return nil, fmt.Errorf("TIL %d has no tags", til.ID)
		}
	}

	return tf.Tils, nil
}
