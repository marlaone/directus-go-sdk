package directusgosdk

import "fmt"

type CollectionMeta struct {
	Collection       string
	Icon             string
	Note             string
	DisplayTemplate  string `json:"display_template"`
	Hidden           bool
	Singleton        bool
	Translations     []Translation
	ArchiveField     string `json:"archive_field"`
	ArchiveValue     string `json:"archive_value"`
	UnarchiveValue   string `json:"unarchive_value"`
	ArchiveAppFilter bool   `json:"archive_app_filter"`
	SortField        bool   `json:"sort_field"`
}

type CollectionSchema struct {
	Name    string
	Comment *string
}

type Collection struct {
	Collection string
	Meta       CollectionMeta
	Schema     CollectionSchema
}

func (d *Directus) GetCollections() ([]*Collection, error) {

	authResponse, err := NewDirectusRequest(d.client, "/collections", "GET", nil)

	if err != nil {
		return nil, fmt.Errorf("collections request error: %v", err)
	}

	fmt.Println(authResponse.Data)

	return []*Collection{}, nil
}
