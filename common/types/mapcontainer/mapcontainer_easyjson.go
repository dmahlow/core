// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package mapcontainer

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

func easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer(in *jlexer.Lexer, out *MapStart) {
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
		case "id":
			out.Id = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "point":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v1 := 0
				for !in.IsDelim(']') {
					if v1 < 2 {
						out.Point[v1] = float64(in.Float64())
						v1++
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
func easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer(out *jwriter.Writer, in MapStart) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.String(string(in.Id))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"point\":")
	out.RawByte('[')
	for v2 := range in.Point {
		if v2 > 0 {
			out.RawByte(',')
		}
		out.Float64(float64(in.Point[v2]))
	}
	out.RawByte(']')
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MapStart) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MapStart) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MapStart) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MapStart) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer(l, v)
}
func easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer1(in *jlexer.Lexer, out *MapPolygon) {
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
		case "points":
			if in.IsNull() {
				in.Skip()
				out.Points = nil
			} else {
				in.Delim('[')
				if out.Points == nil {
					if !in.IsDelim(']') {
						out.Points = make([]MapPoint, 0, 4)
					} else {
						out.Points = []MapPoint{}
					}
				} else {
					out.Points = (out.Points)[:0]
				}
				for !in.IsDelim(']') {
					var v3 MapPoint
					if in.IsNull() {
						in.Skip()
					} else {
						in.Delim('[')
						v4 := 0
						for !in.IsDelim(']') {
							if v4 < 2 {
								v3[v4] = float64(in.Float64())
								v4++
							} else {
								in.SkipRecursive()
							}
							in.WantComma()
						}
						in.Delim(']')
					}
					out.Points = append(out.Points, v3)
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
func easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer1(out *jwriter.Writer, in MapPolygon) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"points\":")
	if in.Points == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in.Points {
			if v5 > 0 {
				out.RawByte(',')
			}
			out.RawByte('[')
			for v7 := range v6 {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.Float64(float64(v6[v7]))
			}
			out.RawByte(']')
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MapPolygon) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MapPolygon) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MapPolygon) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MapPolygon) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer1(l, v)
}
func easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer2(in *jlexer.Lexer, out *MapObstacleObject) {
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
		case "id":
			out.Id = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "polygon":
			(out.Polygon).UnmarshalEasyJSON(in)
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
func easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer2(out *jwriter.Writer, in MapObstacleObject) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.String(string(in.Id))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"polygon\":")
	(in.Polygon).MarshalEasyJSON(out)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MapObstacleObject) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MapObstacleObject) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MapObstacleObject) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MapObstacleObject) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer2(l, v)
}
func easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer3(in *jlexer.Lexer, out *MapGround) {
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
		case "id":
			out.Id = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "polygon":
			(out.Polygon).UnmarshalEasyJSON(in)
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
func easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer3(out *jwriter.Writer, in MapGround) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"id\":")
	out.String(string(in.Id))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"polygon\":")
	(in.Polygon).MarshalEasyJSON(out)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MapGround) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MapGround) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MapGround) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MapGround) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer3(l, v)
}
func easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer4(in *jlexer.Lexer, out *MapContainer) {
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
		case "meta":
			easyjsonFa548efbDecode(in, &out.Meta)
		case "data":
			easyjsonFa548efbDecode1(in, &out.Data)
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
func easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer4(out *jwriter.Writer, in MapContainer) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"meta\":")
	easyjsonFa548efbEncode(out, in.Meta)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"data\":")
	easyjsonFa548efbEncode1(out, in.Data)
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MapContainer) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MapContainer) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonFa548efbEncodeGithubComBytearenaBytearenaCommonTypesMapcontainer4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MapContainer) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MapContainer) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonFa548efbDecodeGithubComBytearenaBytearenaCommonTypesMapcontainer4(l, v)
}
func easyjsonFa548efbDecode1(in *jlexer.Lexer, out *struct {
	Grounds   []MapGround         `json:"grounds"`
	Starts    []MapStart          `json:"starts"`
	Obstacles []MapObstacleObject `json:"obstacles"`
}) {
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
		case "grounds":
			if in.IsNull() {
				in.Skip()
				out.Grounds = nil
			} else {
				in.Delim('[')
				if out.Grounds == nil {
					if !in.IsDelim(']') {
						out.Grounds = make([]MapGround, 0, 1)
					} else {
						out.Grounds = []MapGround{}
					}
				} else {
					out.Grounds = (out.Grounds)[:0]
				}
				for !in.IsDelim(']') {
					var v8 MapGround
					(v8).UnmarshalEasyJSON(in)
					out.Grounds = append(out.Grounds, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "starts":
			if in.IsNull() {
				in.Skip()
				out.Starts = nil
			} else {
				in.Delim('[')
				if out.Starts == nil {
					if !in.IsDelim(']') {
						out.Starts = make([]MapStart, 0, 1)
					} else {
						out.Starts = []MapStart{}
					}
				} else {
					out.Starts = (out.Starts)[:0]
				}
				for !in.IsDelim(']') {
					var v9 MapStart
					(v9).UnmarshalEasyJSON(in)
					out.Starts = append(out.Starts, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "obstacles":
			if in.IsNull() {
				in.Skip()
				out.Obstacles = nil
			} else {
				in.Delim('[')
				if out.Obstacles == nil {
					if !in.IsDelim(']') {
						out.Obstacles = make([]MapObstacleObject, 0, 1)
					} else {
						out.Obstacles = []MapObstacleObject{}
					}
				} else {
					out.Obstacles = (out.Obstacles)[:0]
				}
				for !in.IsDelim(']') {
					var v10 MapObstacleObject
					(v10).UnmarshalEasyJSON(in)
					out.Obstacles = append(out.Obstacles, v10)
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
func easyjsonFa548efbEncode1(out *jwriter.Writer, in struct {
	Grounds   []MapGround         `json:"grounds"`
	Starts    []MapStart          `json:"starts"`
	Obstacles []MapObstacleObject `json:"obstacles"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"grounds\":")
	if in.Grounds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v11, v12 := range in.Grounds {
			if v11 > 0 {
				out.RawByte(',')
			}
			(v12).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"starts\":")
	if in.Starts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v13, v14 := range in.Starts {
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
	out.RawString("\"obstacles\":")
	if in.Obstacles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v15, v16 := range in.Obstacles {
			if v15 > 0 {
				out.RawByte(',')
			}
			(v16).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}
func easyjsonFa548efbDecode(in *jlexer.Lexer, out *struct {
	Readme         string `json:"readme"`
	Kind           string `json:"kind"`
	MaxContestants int    `json:"maxcontestants"`
	Date           string `json:"date"`
}) {
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
		case "readme":
			out.Readme = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "maxcontestants":
			out.MaxContestants = int(in.Int())
		case "date":
			out.Date = string(in.String())
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
func easyjsonFa548efbEncode(out *jwriter.Writer, in struct {
	Readme         string `json:"readme"`
	Kind           string `json:"kind"`
	MaxContestants int    `json:"maxcontestants"`
	Date           string `json:"date"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"readme\":")
	out.String(string(in.Readme))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"kind\":")
	out.String(string(in.Kind))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"maxcontestants\":")
	out.Int(int(in.MaxContestants))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"date\":")
	out.String(string(in.Date))
	out.RawByte('}')
}
