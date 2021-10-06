package directusgosdk

import (
	"encoding/json"
	"fmt"
)

type ItemsResponse struct {
	Data   []Item
	Errors []DirectusErrors
}

type Item interface{}

func (d *Directus) GetItems(collectionName string) ([]Item, error) {
	var itemsResponse ItemsResponse
	resp, err := NewDirectusRequest(d.client, fmt.Sprintf("/items/%s", collectionName), "GET", nil)

	if err != nil {
		return nil, fmt.Errorf("items request error: %v", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&itemsResponse); err != nil {
		return nil, fmt.Errorf("items response error: %v", err)
	}

	if err := CheckDirectusErrors(itemsResponse.Errors); err != nil {
		return nil, fmt.Errorf("items response errors: %v", err)

	}

	return itemsResponse.Data, nil
}
