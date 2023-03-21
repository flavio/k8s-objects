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

func easyjson44f85682DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *NodeSelectorRequirement) {
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
		case "key":
			if in.IsNull() {
				in.Skip()
				out.Key = nil
			} else {
				if out.Key == nil {
					out.Key = new(string)
				}
				*out.Key = string(in.String())
			}
		case "operator":
			if in.IsNull() {
				in.Skip()
				out.Operator = nil
			} else {
				if out.Operator == nil {
					out.Operator = new(string)
				}
				*out.Operator = string(in.String())
			}
		case "values":
			if in.IsNull() {
				in.Skip()
				out.Values = nil
			} else {
				in.Delim('[')
				if out.Values == nil {
					if !in.IsDelim(']') {
						out.Values = make([]string, 0, 4)
					} else {
						out.Values = []string{}
					}
				} else {
					out.Values = (out.Values)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Values = append(out.Values, v1)
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
func easyjson44f85682EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in NodeSelectorRequirement) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"key\":"
		out.RawString(prefix[1:])
		if in.Key == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Key))
		}
	}
	{
		const prefix string = ",\"operator\":"
		out.RawString(prefix)
		if in.Operator == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Operator))
		}
	}
	if len(in.Values) != 0 {
		const prefix string = ",\"values\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.Values {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NodeSelectorRequirement) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson44f85682EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NodeSelectorRequirement) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson44f85682EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NodeSelectorRequirement) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson44f85682DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NodeSelectorRequirement) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson44f85682DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
