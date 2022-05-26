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

func easyjson227342c3DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(in *jlexer.Lexer, out *FSGroupStrategyOptions) {
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
		case "ranges":
			if in.IsNull() {
				in.Skip()
				out.Ranges = nil
			} else {
				in.Delim('[')
				if out.Ranges == nil {
					if !in.IsDelim(']') {
						out.Ranges = make([]*IDRange, 0, 8)
					} else {
						out.Ranges = []*IDRange{}
					}
				} else {
					out.Ranges = (out.Ranges)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *IDRange
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(IDRange)
						}
						easyjson227342c3DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(in, v1)
					}
					out.Ranges = append(out.Ranges, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "rule":
			out.Rule = string(in.String())
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
func easyjson227342c3EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(out *jwriter.Writer, in FSGroupStrategyOptions) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ranges\":"
		out.RawString(prefix[1:])
		if in.Ranges == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Ranges {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson227342c3EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Rule != "" {
		const prefix string = ",\"rule\":"
		out.RawString(prefix)
		out.String(string(in.Rule))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FSGroupStrategyOptions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson227342c3EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FSGroupStrategyOptions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson227342c3EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FSGroupStrategyOptions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson227342c3DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FSGroupStrategyOptions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson227342c3DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(l, v)
}
func easyjson227342c3DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(in *jlexer.Lexer, out *IDRange) {
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
					out.Max = new(int64)
				}
				*out.Max = int64(in.Int64())
			}
		case "min":
			if in.IsNull() {
				in.Skip()
				out.Min = nil
			} else {
				if out.Min == nil {
					out.Min = new(int64)
				}
				*out.Min = int64(in.Int64())
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
func easyjson227342c3EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(out *jwriter.Writer, in IDRange) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"max\":"
		out.RawString(prefix[1:])
		if in.Max == nil {
			out.RawString("null")
		} else {
			out.Int64(int64(*in.Max))
		}
	}
	{
		const prefix string = ",\"min\":"
		out.RawString(prefix)
		if in.Min == nil {
			out.RawString("null")
		} else {
			out.Int64(int64(*in.Min))
		}
	}
	out.RawByte('}')
}
