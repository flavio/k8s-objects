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

func easyjson3aa86ec1DecodeGithubComKubewardenK8sObjectsApiRbacV1(in *jlexer.Lexer, out *ClusterRole) {
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
		case "aggregationRule":
			if in.IsNull() {
				in.Skip()
				out.AggregationRule = nil
			} else {
				if out.AggregationRule == nil {
					out.AggregationRule = new(AggregationRule)
				}
				(*out.AggregationRule).UnmarshalEasyJSON(in)
			}
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "metadata":
			(out.Metadata).UnmarshalEasyJSON(in)
		case "rules":
			if in.IsNull() {
				in.Skip()
				out.Rules = nil
			} else {
				in.Delim('[')
				if out.Rules == nil {
					if !in.IsDelim(']') {
						out.Rules = make([]*PolicyRule, 0, 8)
					} else {
						out.Rules = []*PolicyRule{}
					}
				} else {
					out.Rules = (out.Rules)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *PolicyRule
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(PolicyRule)
						}
						easyjson3aa86ec1DecodeGithubComKubewardenK8sObjectsApiRbacV11(in, v1)
					}
					out.Rules = append(out.Rules, v1)
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
func easyjson3aa86ec1EncodeGithubComKubewardenK8sObjectsApiRbacV1(out *jwriter.Writer, in ClusterRole) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AggregationRule != nil {
		const prefix string = ",\"aggregationRule\":"
		first = false
		out.RawString(prefix[1:])
		(*in.AggregationRule).MarshalEasyJSON(out)
	}
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.APIVersion))
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Kind))
	}
	if true {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Metadata).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"rules\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Rules == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Rules {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson3aa86ec1EncodeGithubComKubewardenK8sObjectsApiRbacV11(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ClusterRole) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3aa86ec1EncodeGithubComKubewardenK8sObjectsApiRbacV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ClusterRole) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3aa86ec1EncodeGithubComKubewardenK8sObjectsApiRbacV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ClusterRole) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3aa86ec1DecodeGithubComKubewardenK8sObjectsApiRbacV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ClusterRole) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3aa86ec1DecodeGithubComKubewardenK8sObjectsApiRbacV1(l, v)
}
func easyjson3aa86ec1DecodeGithubComKubewardenK8sObjectsApiRbacV11(in *jlexer.Lexer, out *PolicyRule) {
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
					var v4 string
					v4 = string(in.String())
					out.APIGroups = append(out.APIGroups, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "nonResourceURLs":
			if in.IsNull() {
				in.Skip()
				out.NonResourceURLs = nil
			} else {
				in.Delim('[')
				if out.NonResourceURLs == nil {
					if !in.IsDelim(']') {
						out.NonResourceURLs = make([]string, 0, 4)
					} else {
						out.NonResourceURLs = []string{}
					}
				} else {
					out.NonResourceURLs = (out.NonResourceURLs)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.NonResourceURLs = append(out.NonResourceURLs, v5)
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
					var v6 string
					v6 = string(in.String())
					out.ResourceNames = append(out.ResourceNames, v6)
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
					var v7 string
					v7 = string(in.String())
					out.Resources = append(out.Resources, v7)
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
					var v8 string
					v8 = string(in.String())
					out.Verbs = append(out.Verbs, v8)
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
func easyjson3aa86ec1EncodeGithubComKubewardenK8sObjectsApiRbacV11(out *jwriter.Writer, in PolicyRule) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"apiGroups\":"
		out.RawString(prefix[1:])
		if in.APIGroups == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.APIGroups {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"nonResourceURLs\":"
		out.RawString(prefix)
		if in.NonResourceURLs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.NonResourceURLs {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"resourceNames\":"
		out.RawString(prefix)
		if in.ResourceNames == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v13, v14 := range in.ResourceNames {
				if v13 > 0 {
					out.RawByte(',')
				}
				out.String(string(v14))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"resources\":"
		out.RawString(prefix)
		if in.Resources == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v15, v16 := range in.Resources {
				if v15 > 0 {
					out.RawByte(',')
				}
				out.String(string(v16))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"verbs\":"
		out.RawString(prefix)
		if in.Verbs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Verbs {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.String(string(v18))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
