// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v2

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

func easyjson7d7eb819DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(in *jlexer.Lexer, out *ObjectMetricStatus) {
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
		case "current":
			if in.IsNull() {
				in.Skip()
				out.Current = nil
			} else {
				if out.Current == nil {
					out.Current = new(MetricValueStatus)
				}
				(*out.Current).UnmarshalEasyJSON(in)
			}
		case "describedObject":
			if in.IsNull() {
				in.Skip()
				out.DescribedObject = nil
			} else {
				if out.DescribedObject == nil {
					out.DescribedObject = new(CrossVersionObjectReference)
				}
				(*out.DescribedObject).UnmarshalEasyJSON(in)
			}
		case "metric":
			if in.IsNull() {
				in.Skip()
				out.Metric = nil
			} else {
				if out.Metric == nil {
					out.Metric = new(MetricIdentifier)
				}
				(*out.Metric).UnmarshalEasyJSON(in)
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
func easyjson7d7eb819EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(out *jwriter.Writer, in ObjectMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"current\":"
		out.RawString(prefix[1:])
		if in.Current == nil {
			out.RawString("null")
		} else {
			(*in.Current).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"describedObject\":"
		out.RawString(prefix)
		if in.DescribedObject == nil {
			out.RawString("null")
		} else {
			(*in.DescribedObject).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"metric\":"
		out.RawString(prefix)
		if in.Metric == nil {
			out.RawString("null")
		} else {
			(*in.Metric).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ObjectMetricStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7d7eb819EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ObjectMetricStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7d7eb819EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ObjectMetricStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7d7eb819DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ObjectMetricStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7d7eb819DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2(l, v)
}