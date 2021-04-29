package alfrednize

import "encoding/json"

type alfredFormat struct {
	Items []alfredFormatItem `json:"items"`
}

// AlfredFormatItem represents alfred script filter JSON format item.
type alfredFormatItem struct {
	UID          string `json:"uid"`
	Title        string `json:"title"`
	SubTitle     string `json:"subtitle"`
	Arg          string `json:"arg"`
	Match        string `json:"match"`
	AutoComplete string `json:"autocomplete"`
}

// Alfrednize convert from strings to Alfred format JSON
func Alfrednize(items []string) ([]byte, error) {
	if len(items) == 0 {
		return nil, nil
	}
	alfredItems := make([]alfredFormatItem, len(items))
	for i, v := range items {
		alfredItems[i] = alfredFormatItem{
			UID:          v,
			Title:        v,
			Arg:          v,
			Match:        v,
			AutoComplete: v,
		}
	}
	alfred := &alfredFormat{
		Items: alfredItems,
	}

	result, err := json.Marshal(alfred)
	if err != nil {
		return nil, err
	}

	return result, nil
}
