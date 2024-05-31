package parser

import (
	"errors"
	"reflect"
	"strings"
)

const minTags = 4
const maxTags = 5

// StructTagParser returns the struct tags for a given struct field.
type StructTagParser interface {
	ParseStructTag(reflect.StructField) (StructTag, error)
}

// StructTagParserFunc is an adapter that allows a generic function to be used
// as a StructTagParser.
type StructTagParserFunc func(sf reflect.StructField) (StructTag, error)

func (f StructTagParserFunc) ParseStructTag(sf reflect.StructField) (StructTag, error) {
	return f(sf)
}

// StructTag represents the struct tag fields
type StructTag struct {
	BelongsTo    bool
	HasMany      bool
	From         string
	LocalField   string
	ForeignField string
	As           string
}

// DefaultStructTagParser is the StructTagParser used by default.
// It will handle the ref struct tag.
//
// The tag formats accepted are:
//
//	"<relation>,<from>,<localField>,<foreignField>[,<alias>]"
//
//	`(...) ref:"<relation>,<from>,<localField>,<foreignField>[,<alias>]" (...)`
//
// If there is no "as" name in the struct tag fields, the struct field name is used in lowercase.
//
// An example:
//
//	type T struct {
//	    A struct{} `ref:"belongsTo,collectionName,localFieldName,foreignFieldName"`
//	}
var DefaultStructTagParser StructTagParserFunc = func(sf reflect.StructField) (StructTag, error) {
	key := strings.ToLower(sf.Name)
	tag, ok := sf.Tag.Lookup("ref")
	if !ok {
		return StructTag{}, nil
	}
	return parseTags(key, tag)
}

func parseTags(key string, tag string) (StructTag, error) {
	tags := strings.Split(tag, ",")

	if len(tags) < minTags || len(tags) > maxTags {
		return StructTag{}, errors.New("invalid tag format")
	}

	for i := range tags {
		tags[i] = strings.TrimSpace(tags[i])
	}

	st := StructTag{
		From:         tags[1],
		LocalField:   tags[2],
		ForeignField: tags[3],
		As:           key,
	}

	switch tags[0] {
	case "belongsTo":
		st.BelongsTo = true
	case "hasMany":
		st.HasMany = true
	default:
		return StructTag{}, errors.New("invalid relation type, must be 'belongsTo' or 'hasMany'")
	}

	if len(tags) == maxTags {
		st.As = tags[4]
	}

	return st, nil
}
