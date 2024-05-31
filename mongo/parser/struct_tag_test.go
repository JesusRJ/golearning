package parser

import (
	"reflect"
	"testing"
)

func TestDefaultStructTagParser(t *testing.T) {
	tb := []struct {
		name string
		tag  string
		want StructTag
		err  bool
	}{
		{
			"belongsTo",
			`ref:"belongsTo,company,company_id,_id"`,
			StructTag{
				BelongsTo:    true,
				From:         "company",
				LocalField:   "company_id",
				ForeignField: "_id",
				As:           "belongsto",
			},
			false,
		},
		{
			"hasMany",
			`ref:"hasMany,company,company_id,_id"`,
			StructTag{
				HasMany:      true,
				From:         "company",
				LocalField:   "company_id",
				ForeignField: "_id",
				As:           "hasmany",
			},
			false,
		},
		{
			"with_as_field",
			`ref:"belongsTo,company,company_id,_id,asField"`,
			StructTag{
				BelongsTo:    true,
				From:         "company",
				LocalField:   "company_id",
				ForeignField: "_id",
				As:           "asField",
			},
			false,
		},
		{
			"empty_fields",
			`ref:",,,,"`,
			StructTag{
				From:         "",
				LocalField:   "",
				ForeignField: "",
				As:           "",
			},
			true,
		},
		{
			"error_invalid_relation",
			`ref:"oneToOne,,,,"`,
			StructTag{
				From:         "",
				LocalField:   "",
				ForeignField: "",
				As:           "",
			},
			true,
		},
		{
			"no_tag",
			"",
			StructTag{},
			false,
		},
	}

	for _, tt := range tb {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DefaultStructTagParser.ParseStructTag(reflect.StructField{
				Name: tt.name,
				Tag:  reflect.StructTag(tt.tag),
			})
			if err != nil && !tt.err {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
