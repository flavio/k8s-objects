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

func easyjson17697e04DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *LocalVolumeSource) {
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
		case "fsType":
			out.FsType = string(in.String())
		case "path":
			if in.IsNull() {
				in.Skip()
				out.Path = nil
			} else {
				if out.Path == nil {
					out.Path = new(string)
				}
				*out.Path = string(in.String())
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
func easyjson17697e04EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in LocalVolumeSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FsType != "" {
		const prefix string = ",\"fsType\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.FsType))
	}
	{
		const prefix string = ",\"path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Path == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Path))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LocalVolumeSource) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson17697e04EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LocalVolumeSource) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson17697e04EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LocalVolumeSource) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson17697e04DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LocalVolumeSource) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson17697e04DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
