// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package deathmatch

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

func easyjsonA8da870DecodeGithubComBytearenaBytearenaGameDeathmatch(in *jlexer.Lexer, out *mailboxMessagePerceptionWrapper) {
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
		case "subject":
			out.Subject = string(in.String())
		case "body":
			if m, ok := out.Body.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Body.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Body = in.Interface()
			}
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
func easyjsonA8da870EncodeGithubComBytearenaBytearenaGameDeathmatch(out *jwriter.Writer, in mailboxMessagePerceptionWrapper) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"subject\":")
	out.String(string(in.Subject))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"body\":")
	if m, ok := in.Body.(easyjson.Marshaler); ok {
		m.MarshalEasyJSON(out)
	} else if m, ok := in.Body.(json.Marshaler); ok {
		out.Raw(m.MarshalJSON())
	} else {
		out.Raw(json.Marshal(in.Body))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v mailboxMessagePerceptionWrapper) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA8da870EncodeGithubComBytearenaBytearenaGameDeathmatch(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v mailboxMessagePerceptionWrapper) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8da870EncodeGithubComBytearenaBytearenaGameDeathmatch(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *mailboxMessagePerceptionWrapper) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA8da870DecodeGithubComBytearenaBytearenaGameDeathmatch(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *mailboxMessagePerceptionWrapper) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8da870DecodeGithubComBytearenaBytearenaGameDeathmatch(l, v)
}
func easyjsonA8da870DecodeGithubComBytearenaBytearenaGameDeathmatch1(in *jlexer.Lexer, out *agentPerceptionVisionItem) {
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
		case "tag":
			out.Tag = string(in.String())
		case "nearedge":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v1 := 0
				for !in.IsDelim(']') {
					if v1 < 2 {
						out.NearEdge[v1] = float64(in.Float64())
						v1++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		case "center":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v2 := 0
				for !in.IsDelim(']') {
					if v2 < 2 {
						out.Center[v2] = float64(in.Float64())
						v2++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		case "faredge":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v3 := 0
				for !in.IsDelim(']') {
					if v3 < 2 {
						out.FarEdge[v3] = float64(in.Float64())
						v3++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		case "velocity":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v4 := 0
				for !in.IsDelim(']') {
					if v4 < 2 {
						out.Velocity[v4] = float64(in.Float64())
						v4++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonA8da870EncodeGithubComBytearenaBytearenaGameDeathmatch1(out *jwriter.Writer, in agentPerceptionVisionItem) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"tag\":")
	out.String(string(in.Tag))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"nearedge\":")
	out.RawByte('[')
	for v5 := range in.NearEdge {
		if v5 > 0 {
			out.RawByte(',')
		}
		out.Float64(float64(in.NearEdge[v5]))
	}
	out.RawByte(']')
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"center\":")
	out.RawByte('[')
	for v6 := range in.Center {
		if v6 > 0 {
			out.RawByte(',')
		}
		out.Float64(float64(in.Center[v6]))
	}
	out.RawByte(']')
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"faredge\":")
	out.RawByte('[')
	for v7 := range in.FarEdge {
		if v7 > 0 {
			out.RawByte(',')
		}
		out.Float64(float64(in.FarEdge[v7]))
	}
	out.RawByte(']')
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"velocity\":")
	out.RawByte('[')
	for v8 := range in.Velocity {
		if v8 > 0 {
			out.RawByte(',')
		}
		out.Float64(float64(in.Velocity[v8]))
	}
	out.RawByte(']')
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v agentPerceptionVisionItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA8da870EncodeGithubComBytearenaBytearenaGameDeathmatch1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v agentPerceptionVisionItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8da870EncodeGithubComBytearenaBytearenaGameDeathmatch1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *agentPerceptionVisionItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA8da870DecodeGithubComBytearenaBytearenaGameDeathmatch1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *agentPerceptionVisionItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8da870DecodeGithubComBytearenaBytearenaGameDeathmatch1(l, v)
}
func easyjsonA8da870DecodeGithubComBytearenaBytearenaGameDeathmatch2(in *jlexer.Lexer, out *agentPerception) {
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
		case "score":
			out.Score = int(in.Int())
		case "energy":
			out.Energy = float64(in.Float64())
		case "velocity":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v9 := 0
				for !in.IsDelim(']') {
					if v9 < 2 {
						out.Velocity[v9] = float64(in.Float64())
						v9++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		case "azimuth":
			out.Azimuth = float64(in.Float64())
		case "vision":
			if in.IsNull() {
				in.Skip()
				out.Vision = nil
			} else {
				in.Delim('[')
				if out.Vision == nil {
					if !in.IsDelim(']') {
						out.Vision = make([]agentPerceptionVisionItem, 0, 1)
					} else {
						out.Vision = []agentPerceptionVisionItem{}
					}
				} else {
					out.Vision = (out.Vision)[:0]
				}
				for !in.IsDelim(']') {
					var v10 agentPerceptionVisionItem
					(v10).UnmarshalEasyJSON(in)
					out.Vision = append(out.Vision, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "shootenergy":
			out.ShootEnergy = float64(in.Float64())
		case "shootcooldown":
			out.ShootCooldown = int(in.Int())
		case "messages":
			if in.IsNull() {
				in.Skip()
				out.Messages = nil
			} else {
				in.Delim('[')
				if out.Messages == nil {
					if !in.IsDelim(']') {
						out.Messages = make([]mailboxMessagePerceptionWrapper, 0, 2)
					} else {
						out.Messages = []mailboxMessagePerceptionWrapper{}
					}
				} else {
					out.Messages = (out.Messages)[:0]
				}
				for !in.IsDelim(']') {
					var v11 mailboxMessagePerceptionWrapper
					(v11).UnmarshalEasyJSON(in)
					out.Messages = append(out.Messages, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonA8da870EncodeGithubComBytearenaBytearenaGameDeathmatch2(out *jwriter.Writer, in agentPerception) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"score\":")
	out.Int(int(in.Score))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"energy\":")
	out.Float64(float64(in.Energy))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"velocity\":")
	out.RawByte('[')
	for v12 := range in.Velocity {
		if v12 > 0 {
			out.RawByte(',')
		}
		out.Float64(float64(in.Velocity[v12]))
	}
	out.RawByte(']')
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"azimuth\":")
	out.Float64(float64(in.Azimuth))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"vision\":")
	if in.Vision == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v13, v14 := range in.Vision {
			if v13 > 0 {
				out.RawByte(',')
			}
			(v14).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"shootenergy\":")
	out.Float64(float64(in.ShootEnergy))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"shootcooldown\":")
	out.Int(int(in.ShootCooldown))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"messages\":")
	if in.Messages == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v15, v16 := range in.Messages {
			if v15 > 0 {
				out.RawByte(',')
			}
			(v16).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v agentPerception) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA8da870EncodeGithubComBytearenaBytearenaGameDeathmatch2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v agentPerception) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8da870EncodeGithubComBytearenaBytearenaGameDeathmatch2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *agentPerception) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA8da870DecodeGithubComBytearenaBytearenaGameDeathmatch2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *agentPerception) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8da870DecodeGithubComBytearenaBytearenaGameDeathmatch2(l, v)
}
