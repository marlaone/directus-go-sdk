package directusgosdk

import (
	"encoding/json"
	"fmt"
)

type CollectionResponse struct {
	Data   []Collection
	Errors []*DirectusErrors
}

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

func (d *Directus) GetCollections() ([]Collection, error) {
	var collectionResponse CollectionResponse
	resp, err := NewDirectusRequest(d.client, "/collections", "GET", nil)

	if err != nil {
		return nil, fmt.Errorf("collections request error: %v", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&collectionResponse); err != nil {
		return nil, fmt.Errorf("directus response error: %v", err)
	}
	return collectionResponse.Data, nil
}

func (d *Directus) GetCollection(collectionName string) (Collection, error) {
	return Collection{}, nil
}
