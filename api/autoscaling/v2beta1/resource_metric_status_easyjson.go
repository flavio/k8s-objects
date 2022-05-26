// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v2beta1

import (
	json "encoding/json"
	resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
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

func easyjsonF2b5af68DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(in *jlexer.Lexer, out *ResourceMetricStatus) {
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
		case "currentAverageUtilization":
			out.CurrentAverageUtilization = int32(in.Int32())
		case "currentAverageValue":
			if in.IsNull() {
				in.Skip()
				out.CurrentAverageValue = nil
			} else {
				if out.CurrentAverageValue == nil {
					out.CurrentAverageValue = new(resource.Quantity)
				}
				*out.CurrentAverageValue = resource.Quantity(in.String())
			}
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
func easyjsonF2b5af68EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(out *jwriter.Writer, in ResourceMetricStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CurrentAverageUtilization != 0 {
		const prefix string = ",\"currentAverageUtilization\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.CurrentAverageUtilization))
	}
	{
		const prefix string = ",\"currentAverageValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.CurrentAverageValue == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.CurrentAverageValue))
		}
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		if in.Name == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Name))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResourceMetricStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF2b5af68EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResourceMetricStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF2b5af68EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResourceMetricStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF2b5af68DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResourceMetricStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF2b5af68DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(l, v)
}
