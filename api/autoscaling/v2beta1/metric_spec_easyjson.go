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

func easyjson99ef38b0DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(in *jlexer.Lexer, out *MetricSpec) {
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
		case "external":
			if in.IsNull() {
				in.Skip()
				out.External = nil
			} else {
				if out.External == nil {
					out.External = new(ExternalMetricSource)
				}
				(*out.External).UnmarshalEasyJSON(in)
			}
		case "object":
			if in.IsNull() {
				in.Skip()
				out.Object = nil
			} else {
				if out.Object == nil {
					out.Object = new(ObjectMetricSource)
				}
				easyjson99ef38b0DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(in, out.Object)
			}
		case "pods":
			if in.IsNull() {
				in.Skip()
				out.Pods = nil
			} else {
				if out.Pods == nil {
					out.Pods = new(PodsMetricSource)
				}
				easyjson99ef38b0DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(in, out.Pods)
			}
		case "resource":
			if in.IsNull() {
				in.Skip()
				out.Resource = nil
			} else {
				if out.Resource == nil {
					out.Resource = new(ResourceMetricSource)
				}
				easyjson99ef38b0DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(in, out.Resource)
			}
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(string)
				}
				*out.Type = string(in.String())
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
func easyjson99ef38b0EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(out *jwriter.Writer, in MetricSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.External != nil {
		const prefix string = ",\"external\":"
		first = false
		out.RawString(prefix[1:])
		(*in.External).MarshalEasyJSON(out)
	}
	if in.Object != nil {
		const prefix string = ",\"object\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson99ef38b0EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(out, *in.Object)
	}
	if in.Pods != nil {
		const prefix string = ",\"pods\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson99ef38b0EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(out, *in.Pods)
	}
	if in.Resource != nil {
		const prefix string = ",\"resource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson99ef38b0EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(out, *in.Resource)
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Type == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Type))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MetricSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson99ef38b0EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MetricSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson99ef38b0EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MetricSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson99ef38b0DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MetricSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson99ef38b0DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(l, v)
}
func easyjson99ef38b0DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(in *jlexer.Lexer, out *ResourceMetricSource) {
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
		case "targetAverageUtilization":
			out.TargetAverageUtilization = int32(in.Int32())
		case "targetAverageValue":
			out.TargetAverageValue = resource.Quantity(in.String())
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
func easyjson99ef38b0EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(out *jwriter.Writer, in ResourceMetricSource) {
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
	if in.TargetAverageUtilization != 0 {
		const prefix string = ",\"targetAverageUtilization\":"
		out.RawString(prefix)
		out.Int32(int32(in.TargetAverageUtilization))
	}
	if in.TargetAverageValue != "" {
		const prefix string = ",\"targetAverageValue\":"
		out.RawString(prefix)
		out.String(string(in.TargetAverageValue))
	}
	out.RawByte('}')
}
func easyjson99ef38b0DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(in *jlexer.Lexer, out *PodsMetricSource) {
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
		case "metricName":
			if in.IsNull() {
				in.Skip()
				out.MetricName = nil
			} else {
				if out.MetricName == nil {
					out.MetricName = new(string)
				}
				*out.MetricName = string(in.String())
			}
		case "selector":
			(out.Selector).UnmarshalEasyJSON(in)
		case "targetAverageValue":
			if in.IsNull() {
				in.Skip()
				out.TargetAverageValue = nil
			} else {
				if out.TargetAverageValue == nil {
					out.TargetAverageValue = new(resource.Quantity)
				}
				*out.TargetAverageValue = resource.Quantity(in.String())
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
func easyjson99ef38b0EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(out *jwriter.Writer, in PodsMetricSource) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"metricName\":"
		out.RawString(prefix[1:])
		if in.MetricName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.MetricName))
		}
	}
	if true {
		const prefix string = ",\"selector\":"
		out.RawString(prefix)
		(in.Selector).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"targetAverageValue\":"
		out.RawString(prefix)
		if in.TargetAverageValue == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.TargetAverageValue))
		}
	}
	out.RawByte('}')
}
func easyjson99ef38b0DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(in *jlexer.Lexer, out *ObjectMetricSource) {
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
		case "averageValue":
			out.AverageValue = resource.Quantity(in.String())
		case "metricName":
			if in.IsNull() {
				in.Skip()
				out.MetricName = nil
			} else {
				if out.MetricName == nil {
					out.MetricName = new(string)
				}
				*out.MetricName = string(in.String())
			}
		case "selector":
			(out.Selector).UnmarshalEasyJSON(in)
		case "target":
			if in.IsNull() {
				in.Skip()
				out.Target = nil
			} else {
				if out.Target == nil {
					out.Target = new(CrossVersionObjectReference)
				}
				(*out.Target).UnmarshalEasyJSON(in)
			}
		case "targetValue":
			if in.IsNull() {
				in.Skip()
				out.TargetValue = nil
			} else {
				if out.TargetValue == nil {
					out.TargetValue = new(resource.Quantity)
				}
				*out.TargetValue = resource.Quantity(in.String())
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
func easyjson99ef38b0EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(out *jwriter.Writer, in ObjectMetricSource) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AverageValue != "" {
		const prefix string = ",\"averageValue\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.AverageValue))
	}
	{
		const prefix string = ",\"metricName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.MetricName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.MetricName))
		}
	}
	if true {
		const prefix string = ",\"selector\":"
		out.RawString(prefix)
		(in.Selector).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"target\":"
		out.RawString(prefix)
		if in.Target == nil {
			out.RawString("null")
		} else {
			(*in.Target).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"targetValue\":"
		out.RawString(prefix)
		if in.TargetValue == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.TargetValue))
		}
	}
	out.RawByte('}')
}
