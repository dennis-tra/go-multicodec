// Code generated by go generate; DO NOT EDIT.
package multicodec

// Codec serves the purpose of type safety when passing around codecs
type Codec uint64

// These are multicodec-packed content types.
const (
{{ range .Codecs }}// {{ if .IsDeprecated }}Deprecated: {{ end }}{{ .Tag }}{{ if .Description }}: {{ .Description }}{{ end }}
   {{ .VarName }} Codec = {{ .Code }}
{{ end }}
)