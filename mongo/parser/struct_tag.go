package parser

import (
	"errors"
	"reflect"
	"strings"
)

const minTags = 4
const maxTags = 5
const tagRef = "ref"

const (
	None RelationType = iota
	BelongsTo
	HasMany
)

// Relation represents the relation type between two models.
type RelationType int

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
	Relation     RelationType
	From         string
	LocalField   string
	ForeignField string
	As           string
}

// BelongsTo returns true if the relation is belongs to.
func (st StructTag) BelongsTo() bool {
	return st.Relation == BelongsTo
}

// HasMany returns true if the relation is has many.
func (st StructTag) HasMany() bool {
	return st.Relation == HasMany
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
	tag, ok := sf.Tag.Lookup(tagRef)
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

	st := StructTag{
		From:         strings.TrimSpace(tags[1]),
		LocalField:   strings.TrimSpace(tags[2]),
		ForeignField: strings.TrimSpace(tags[3]),
		As:           strings.TrimSpace(key),
	}

	switch tags[0] {
	case "belongsTo":
		st.Relation = BelongsTo
	case "hasMany":
		st.Relation = HasMany
	default:
		st.Relation = None
	}

	// Replace default "As" name with the one provided in the tag
	if len(tags) == maxTags {
		st.As = strings.TrimSpace(tags[4])
	}

	return st, nil
}
