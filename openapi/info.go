package openapi

import (
	"errors"
	"net/url"
)

// codebeat:disable[TOO_MANY_IVARS]

// Info Object
type Info struct {
	Title          string   `yaml:"title,omitempty"`
	Description    string   `yaml:"description,omitempty"`
	TermsOfService string   `yaml:"termsOfService,omitempty"`
	Contact        *Contact `yaml:"contact,omitempty"`
	License        *License `yaml:"license,omitempty"`
	Version        string   `yaml:"version,omitempty"`
}

// Validate the values of Info object.
func (info Info) Validate() error {
	if info.Title == "" {
		return errors.New("info.title is required")
	}
	if info.TermsOfService != "" {
		if _, err := url.ParseRequestURI(info.TermsOfService); err != nil {
			return err
		}
	}
	validaters := []validater{}
	if info.Contact != nil {
		validaters = append(validaters, info.Contact)
	}
	if info.License != nil {
		validaters = append(validaters, info.License)
	}
	if err := validateAll(validaters); err != nil {
		return err
	}
	if info.Version == "" {
		return errors.New("info.version is required")
	}
	return nil
}
