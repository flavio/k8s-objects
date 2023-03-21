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

func easyjsonEd5203d4DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(in *jlexer.Lexer, out *ValidationRule) {
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
		case "message":
			out.Message = string(in.String())
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
func easyjsonEd5203d4EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(out *jwriter.Writer, in ValidationRule) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"message\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"rule\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Rule == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Rule))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ValidationRule) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEd5203d4EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ValidationRule) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEd5203d4EncodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ValidationRule) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEd5203d4DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ValidationRule) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEd5203d4DecodeGithubComKubewardenK8sObjectsApiextensionsApiserverPkgApisApiextensionsV1(l, v)
}
