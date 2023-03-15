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

func easyjson39ff2ff3DecodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(in *jlexer.Lexer, out *ResourceRule) {
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
		case "apiGroups":
			if in.IsNull() {
				in.Skip()
				out.APIGroups = nil
			} else {
				in.Delim('[')
				if out.APIGroups == nil {
					if !in.IsDelim(']') {
						out.APIGroups = make([]string, 0, 4)
					} else {
						out.APIGroups = []string{}
					}
				} else {
					out.APIGroups = (out.APIGroups)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.APIGroups = append(out.APIGroups, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "resourceNames":
			if in.IsNull() {
				in.Skip()
				out.ResourceNames = nil
			} else {
				in.Delim('[')
				if out.ResourceNames == nil {
					if !in.IsDelim(']') {
						out.ResourceNames = make([]string, 0, 4)
					} else {
						out.ResourceNames = []string{}
					}
				} else {
					out.ResourceNames = (out.ResourceNames)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.ResourceNames = append(out.ResourceNames, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "resources":
			if in.IsNull() {
				in.Skip()
				out.Resources = nil
			} else {
				in.Delim('[')
				if out.Resources == nil {
					if !in.IsDelim(']') {
						out.Resources = make([]string, 0, 4)
					} else {
						out.Resources = []string{}
					}
				} else {
					out.Resources = (out.Resources)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.Resources = append(out.Resources, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "verbs":
			if in.IsNull() {
				in.Skip()
				out.Verbs = nil
			} else {
				in.Delim('[')
				if out.Verbs == nil {
					if !in.IsDelim(']') {
						out.Verbs = make([]string, 0, 4)
					} else {
						out.Verbs = []string{}
					}
				} else {
					out.Verbs = (out.Verbs)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Verbs = append(out.Verbs, v4)
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
func easyjson39ff2ff3EncodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(out *jwriter.Writer, in ResourceRule) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.APIGroups) != 0 {
		const prefix string = ",\"apiGroups\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v5, v6 := range in.APIGroups {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	if len(in.ResourceNames) != 0 {
		const prefix string = ",\"resourceNames\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.ResourceNames {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.String(string(v8))
			}
			out.RawByte(']')
		}
	}
	if len(in.Resources) != 0 {
		const prefix string = ",\"resources\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.Resources {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"verbs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Verbs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Verbs {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResourceRule) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson39ff2ff3EncodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResourceRule) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson39ff2ff3EncodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResourceRule) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson39ff2ff3DecodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResourceRule) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson39ff2ff3DecodeGithubComKubewardenK8sObjectsApiAuthorizationV1beta1(l, v)
}