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

func easyjsonB800a7a5DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(in *jlexer.Lexer, out *ServiceReference) {
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
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "namespace":
			if in.IsNull() {
				in.Skip()
				out.Namespace = nil
			} else {
				if out.Namespace == nil {
					out.Namespace = new(string)
				}
				*out.Namespace = string(in.String())
			}
		case "path":
			out.Path = string(in.String())
		case "port":
			out.Port = int32(in.Int32())
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
func easyjsonB800a7a5EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(out *jwriter.Writer, in ServiceReference) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	{
		const prefix string = ",\"namespace\":"
		out.RawString(prefix)
		if in.Namespace == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Namespace))
		}
	}
	if in.Path != "" {
		const prefix string = ",\"path\":"
		out.RawString(prefix)
		out.String(string(in.Path))
	}
	if in.Port != 0 {
		const prefix string = ",\"port\":"
		out.RawString(prefix)
		out.Int32(int32(in.Port))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceReference) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB800a7a5EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceReference) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB800a7a5EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceReference) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB800a7a5DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceReference) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB800a7a5DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1beta1(l, v)
}