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

func easyjson230069fcDecodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(in *jlexer.Lexer, out *Score) {
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
		case "value":
			out.Value = int(in.Int())
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
func easyjson230069fcEncodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(out *jwriter.Writer, in Score) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"value\":")
	out.Int(int(in.Value))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Score) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson230069fcEncodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Score) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson230069fcEncodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Score) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson230069fcDecodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Score) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson230069fcDecodeGithubComBytearenaBytearenaGameDeathmatchMailboxmessages(l, v)
}
