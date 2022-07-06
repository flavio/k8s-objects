// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
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

func easyjson7f4cd7fbDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(in *jlexer.Lexer, out *SELinuxStrategyOptions) {
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
		case "rule":
			if in.IsNull() {
				in.Skip()
				out.Rule = nil
			} else {
				if out.Rule == nil {
					out.Rule = new(string)
				}
				*out.Rule = string(in.String())
			}
		case "seLinuxOptions":
			if in.IsNull() {
				in.Skip()
				out.SeLinuxOptions = nil
			} else {
				if out.SeLinuxOptions == nil {
					out.SeLinuxOptions = new(_v1.SELinuxOptions)
				}
				(*out.SeLinuxOptions).UnmarshalEasyJSON(in)
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
func easyjson7f4cd7fbEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(out *jwriter.Writer, in SELinuxStrategyOptions) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rule\":"
		out.RawString(prefix[1:])
		if in.Rule == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Rule))
		}
	}
	if in.SeLinuxOptions != nil {
		const prefix string = ",\"seLinuxOptions\":"
		out.RawString(prefix)
		(*in.SeLinuxOptions).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SELinuxStrategyOptions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7f4cd7fbEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SELinuxStrategyOptions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7f4cd7fbEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SELinuxStrategyOptions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7f4cd7fbDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SELinuxStrategyOptions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7f4cd7fbDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(l, v)
}
