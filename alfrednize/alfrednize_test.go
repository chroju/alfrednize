package alfrednize

import (
	"reflect"
	"testing"
)

var (
	jsonSingleItem    = []byte("{\"items\":[{\"uid\":\"foo\",\"title\":\"foo\",\"subtitle\":\"\",\"arg\":\"foo\",\"match\":\"foo\",\"autocomplete\":\"foo\"}]}")
	jsonMultipleItems = []byte("{\"items\":[{\"uid\":\"foo\",\"title\":\"foo\",\"subtitle\":\"\",\"arg\":\"foo\",\"match\":\"foo\",\"autocomplete\":\"foo\"},{\"uid\":\"bar\",\"title\":\"bar\",\"subtitle\":\"\",\"arg\":\"bar\",\"match\":\"bar\",\"autocomplete\":\"bar\"}]}")
)

func TestAlfrednize(t *testing.T) {
	tests := []struct {
		name    string
		items   []string
		want    []byte
		wantErr bool
	}{
		{
			name:    "no items",
			items:   []string{},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "single item",
			items:   []string{"foo"},
			want:    jsonSingleItem,
			wantErr: false,
		},
		{
			name:    "multiple items",
			items:   []string{"foo", "bar"},
			want:    jsonMultipleItems,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Alfrednize(tt.items)
			if (err != nil) != tt.wantErr {
				t.Errorf("Alfrednize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Alfrednize() = %v, want %v", got, tt.want)
			}
		})
	}
}
