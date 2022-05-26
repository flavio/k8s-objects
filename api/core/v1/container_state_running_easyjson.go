// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

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

func easyjsonB17754e5DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *ContainerStateRunning) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "startedAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.StartedAt).UnmarshalJSON(data))
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
func easyjsonB17754e5EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in ContainerStateRunning) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"startedAt\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((in.StartedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ContainerStateRunning) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB17754e5EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ContainerStateRunning) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB17754e5EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ContainerStateRunning) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB17754e5DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ContainerStateRunning) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB17754e5DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
