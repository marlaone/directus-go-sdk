package directusgosdk

import (
	"encoding/json"
	"fmt"
	"io"
)

type ItemsResponse struct {
	Data   []Item
	Errors []DirectusErrors
}

type Item interface{}

func (d *Directus) GetItems(collectionName string, query *Query) ([]Item, error) {
	var itemsResponse ItemsResponse

	decoder, body, err := d.GetItemsWithDecoder(collectionName, query)

	if err != nil {
		return nil, err
	}

	defer body.Close()

	if err := decoder.Decode(&itemsResponse); err != nil {
		return nil, fmt.Errorf("items response error: %v", err)
	}

	if err := CheckDirectusErrors(itemsResponse.Errors); err != nil {
		return nil, fmt.Errorf("items response errors: %v", err)
	}

	return itemsResponse.Data, nil
}

func (d *Directus) GetItemsWithDecoder(collectionName string, query *Query) (*json.Decoder, io.ReadCloser, error) {
	resp, err := NewDirectusRequest(d.client, fmt.Sprintf("/items/%s", collectionName), "GET", nil)

	if err != nil {
		return nil, nil, fmt.Errorf("items request error: %v", err)
	}

	return json.NewDecoder(resp.Body), resp.Body, nil
}
