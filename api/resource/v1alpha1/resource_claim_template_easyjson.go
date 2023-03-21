// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
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

func easyjsonDd56f20cDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha1(in *jlexer.Lexer, out *ResourceClaimTemplate) {
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
		case "apiVersion":
			out.APIVersion = string(in.String())
		case "kind":
			out.Kind = string(in.String())
		case "metadata":
			if in.IsNull() {
				in.Skip()
				out.Metadata = nil
			} else {
				if out.Metadata == nil {
					out.Metadata = new(_v1.ObjectMeta)
				}
				(*out.Metadata).UnmarshalEasyJSON(in)
			}
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(ResourceClaimTemplateSpec)
				}
				easyjsonDd56f20cDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha11(in, out.Spec)
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
func easyjsonDd56f20cEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha1(out *jwriter.Writer, in ResourceClaimTemplate) {
	out.RawByte('{')
	first := true
	_ = first
	if in.APIVersion != "" {
		const prefix string = ",\"apiVersion\":"
		first = false
		out.RawString(prefix[1:])
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
	if in.Metadata != nil {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Metadata).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Spec == nil {
			out.RawString("null")
		} else {
			easyjsonDd56f20cEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha11(out, *in.Spec)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ResourceClaimTemplate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDd56f20cEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ResourceClaimTemplate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDd56f20cEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ResourceClaimTemplate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDd56f20cDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ResourceClaimTemplate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDd56f20cDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha1(l, v)
}
func easyjsonDd56f20cDecodeGithubComKubewardenK8sObjectsApiResourceV1alpha11(in *jlexer.Lexer, out *ResourceClaimTemplateSpec) {
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
		case "metadata":
			if in.IsNull() {
				in.Skip()
				out.Metadata = nil
			} else {
				if out.Metadata == nil {
					out.Metadata = new(_v1.ObjectMeta)
				}
				(*out.Metadata).UnmarshalEasyJSON(in)
			}
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(ResourceClaimSpec)
				}
				(*out.Spec).UnmarshalEasyJSON(in)
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
func easyjsonDd56f20cEncodeGithubComKubewardenK8sObjectsApiResourceV1alpha11(out *jwriter.Writer, in ResourceClaimTemplateSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Metadata != nil {
		const prefix string = ",\"metadata\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Metadata).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Spec == nil {
			out.RawString("null")
		} else {
			(*in.Spec).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}
