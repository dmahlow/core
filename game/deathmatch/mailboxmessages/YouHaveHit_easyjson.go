// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package mailboxmessages

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

func easyjsonA5513ec6DecodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(in *jlexer.Lexer, out *YouHaveHit) {
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
		case "Who":
			out.Who = string(in.String())
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
func easyjsonA5513ec6EncodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(out *jwriter.Writer, in YouHaveHit) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Who\":")
	out.String(string(in.Who))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v YouHaveHit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA5513ec6EncodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v YouHaveHit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA5513ec6EncodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *YouHaveHit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA5513ec6DecodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *YouHaveHit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA5513ec6DecodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(l, v)
}
