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

func easyjson97a47480DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(in *jlexer.Lexer, out *HorizontalPodAutoscalerSpec) {
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
		case "maxReplicas":
			if in.IsNull() {
				in.Skip()
				out.MaxReplicas = nil
			} else {
				if out.MaxReplicas == nil {
					out.MaxReplicas = new(int32)
				}
				*out.MaxReplicas = int32(in.Int32())
			}
		case "metrics":
			if in.IsNull() {
				in.Skip()
				out.Metrics = nil
			} else {
				in.Delim('[')
				if out.Metrics == nil {
					if !in.IsDelim(']') {
						out.Metrics = make([]*MetricSpec, 0, 8)
					} else {
						out.Metrics = []*MetricSpec{}
					}
				} else {
					out.Metrics = (out.Metrics)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *MetricSpec
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(MetricSpec)
						}
						easyjson97a47480DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(in, v1)
					}
					out.Metrics = append(out.Metrics, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "minReplicas":
			out.MinReplicas = int32(in.Int32())
		case "scaleTargetRef":
			if in.IsNull() {
				in.Skip()
				out.ScaleTargetRef = nil
			} else {
				if out.ScaleTargetRef == nil {
					out.ScaleTargetRef = new(CrossVersionObjectReference)
				}
				(*out.ScaleTargetRef).UnmarshalEasyJSON(in)
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
func easyjson97a47480EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(out *jwriter.Writer, in HorizontalPodAutoscalerSpec) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"maxReplicas\":"
		out.RawString(prefix[1:])
		if in.MaxReplicas == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.MaxReplicas))
		}
	}
	{
		const prefix string = ",\"metrics\":"
		out.RawString(prefix)
		if in.Metrics == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Metrics {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson97a47480EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	if in.MinReplicas != 0 {
		const prefix string = ",\"minReplicas\":"
		out.RawString(prefix)
		out.Int32(int32(in.MinReplicas))
	}
	{
		const prefix string = ",\"scaleTargetRef\":"
		out.RawString(prefix)
		if in.ScaleTargetRef == nil {
			out.RawString("null")
		} else {
			(*in.ScaleTargetRef).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HorizontalPodAutoscalerSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson97a47480EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HorizontalPodAutoscalerSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson97a47480EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HorizontalPodAutoscalerSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson97a47480DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HorizontalPodAutoscalerSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson97a47480DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta1(l, v)
}
func easyjson97a47480DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(in *jlexer.Lexer, out *MetricSpec) {
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
				easyjson97a47480DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(in, out.Object)
			}
		case "pods":
			if in.IsNull() {
				in.Skip()
				out.Pods = nil
			} else {
				if out.Pods == nil {
					out.Pods = new(PodsMetricSource)
				}
				easyjson97a47480DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(in, out.Pods)
			}
		case "resource":
			if in.IsNull() {
				in.Skip()
				out.Resource = nil
			} else {
				if out.Resource == nil {
					out.Resource = new(ResourceMetricSource)
				}
				easyjson97a47480DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta14(in, out.Resource)
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
func easyjson97a47480EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta11(out *jwriter.Writer, in MetricSpec) {
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
		easyjson97a47480EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(out, *in.Object)
	}
	if in.Pods != nil {
		const prefix string = ",\"pods\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson97a47480EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(out, *in.Pods)
	}
	if in.Resource != nil {
		const prefix string = ",\"resource\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson97a47480EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta14(out, *in.Resource)
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
func easyjson97a47480DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta14(in *jlexer.Lexer, out *ResourceMetricSource) {
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
func easyjson97a47480EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta14(out *jwriter.Writer, in ResourceMetricSource) {
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
func easyjson97a47480DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(in *jlexer.Lexer, out *PodsMetricSource) {
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
func easyjson97a47480EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta13(out *jwriter.Writer, in PodsMetricSource) {
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
func easyjson97a47480DecodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(in *jlexer.Lexer, out *ObjectMetricSource) {
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
func easyjson97a47480EncodeGithubComKubewardenK8sObjectsApiAutoscalingV2beta12(out *jwriter.Writer, in ObjectMetricSource) {
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
