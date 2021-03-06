// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package types

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7f617826DecodeGithubComBytearenaBytearenaCommonTypes(in *jlexer.Lexer, out *PixelSurface) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Width":
			out.Width = PixelUnit(in.Float64())
		case "Height":
			out.Height = PixelUnit(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7f617826EncodeGithubComBytearenaBytearenaCommonTypes(out *jwriter.Writer, in PixelSurface) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Width\":")
	out.Float64(float64(in.Width))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Height\":")
	out.Float64(float64(in.Height))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PixelSurface) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7f617826EncodeGithubComBytearenaBytearenaCommonTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PixelSurface) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7f617826EncodeGithubComBytearenaBytearenaCommonTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PixelSurface) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7f617826DecodeGithubComBytearenaBytearenaCommonTypes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PixelSurface) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7f617826DecodeGithubComBytearenaBytearenaCommonTypes(l, v)
}
