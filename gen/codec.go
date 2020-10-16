package main

import "strings"

// Codec represents a row in the canonical multicodec table
type Codec struct {
	Name        string
	Tag         string
	Code        string
	Description string
}

// IsDeprecated returns true if the description contains the word deprecated
func (c Codec) IsDeprecated() bool {
	return strings.Contains(strings.ToLower(c.Description), "deprecated")
}

// VarName returns the variable name to be used in Go for this codec
func (c Codec) VarName() string {
	return toCamel(c.Name)
}
