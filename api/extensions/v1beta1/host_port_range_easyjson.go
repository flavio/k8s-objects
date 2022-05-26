// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

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

func easyjson10e431ecDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(in *jlexer.Lexer, out *HostPortRange) {
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
		case "max":
			if in.IsNull() {
				in.Skip()
				out.Max = nil
			} else {
				if out.Max == nil {
					out.Max = new(int32)
				}
				*out.Max = int32(in.Int32())
			}
		case "min":
			if in.IsNull() {
				in.Skip()
				out.Min = nil
			} else {
				if out.Min == nil {
					out.Min = new(int32)
				}
				*out.Min = int32(in.Int32())
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
func easyjson10e431ecEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(out *jwriter.Writer, in HostPortRange) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"max\":"
		out.RawString(prefix[1:])
		if in.Max == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Max))
		}
	}
	{
		const prefix string = ",\"min\":"
		out.RawString(prefix)
		if in.Min == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Min))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HostPortRange) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson10e431ecEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HostPortRange) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson10e431ecEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HostPortRange) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson10e431ecDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HostPortRange) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson10e431ecDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(l, v)
}
