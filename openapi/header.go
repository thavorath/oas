package openapi

import "errors"

// codebeat:disable[TOO_MANY_IVARS]

// Header Object
type Header struct {
	Description     string
	Required        bool
	Deprecated      string
	AllowEmptyValue bool `yaml:"allowEmptyValue"`

	Style         string
	Explode       bool
	AllowReserved bool `yaml:"allowReserved"`
	Schema        *Schema
	Example       interface{}
	Examples      map[string]*Example

	Content map[string]*MediaType

	Ref string `yaml:"$ref"`
}

// Validate the values of Header object.
func (header Header) Validate() error {
	validaters := []validater{header.Schema}
	if v, ok := header.Example.(validater); ok {
		validaters = append(validaters, v)
	}

	// example has no validation

	if len(header.Content) > 1 {
		return errors.New("header.content must only contain one entry")
	}
	for _, mediaType := range header.Content {
		validaters = append(validaters, mediaType)
	}
	return validateAll(validaters)
}
