// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"
	_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
	intstr "github.com/kubewarden/k8s-objects/apimachinery/pkg/util/intstr"
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

func easyjsonA77018b0DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(in *jlexer.Lexer, out *HTTPIngressPath) {
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
		case "backend":
			if in.IsNull() {
				in.Skip()
				out.Backend = nil
			} else {
				if out.Backend == nil {
					out.Backend = new(IngressBackend)
				}
				easyjsonA77018b0DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(in, out.Backend)
			}
		case "path":
			out.Path = string(in.String())
		case "pathType":
			out.PathType = string(in.String())
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
func easyjsonA77018b0EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(out *jwriter.Writer, in HTTPIngressPath) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"backend\":"
		out.RawString(prefix[1:])
		if in.Backend == nil {
			out.RawString("null")
		} else {
			easyjsonA77018b0EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(out, *in.Backend)
		}
	}
	if in.Path != "" {
		const prefix string = ",\"path\":"
		out.RawString(prefix)
		out.String(string(in.Path))
	}
	if in.PathType != "" {
		const prefix string = ",\"pathType\":"
		out.RawString(prefix)
		out.String(string(in.PathType))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HTTPIngressPath) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA77018b0EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HTTPIngressPath) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA77018b0EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HTTPIngressPath) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA77018b0DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HTTPIngressPath) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA77018b0DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(l, v)
}
func easyjsonA77018b0DecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(in *jlexer.Lexer, out *IngressBackend) {
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
		case "resource":
			if in.IsNull() {
				in.Skip()
				out.Resource = nil
			} else {
				if out.Resource == nil {
					out.Resource = new(_v1.TypedLocalObjectReference)
				}
				(*out.Resource).UnmarshalEasyJSON(in)
			}
		case "serviceName":
			out.ServiceName = string(in.String())
		case "servicePort":
			if in.IsNull() {
				in.Skip()
				out.ServicePort = nil
			} else {
				if out.ServicePort == nil {
					out.ServicePort = new(intstr.IntOrString)
				}
				*out.ServicePort = intstr.IntOrString(in.String())
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
func easyjsonA77018b0EncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(out *jwriter.Writer, in IngressBackend) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Resource != nil {
		const prefix string = ",\"resource\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Resource).MarshalEasyJSON(out)
	}
	if in.ServiceName != "" {
		const prefix string = ",\"serviceName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ServiceName))
	}
	if in.ServicePort != nil {
		const prefix string = ",\"servicePort\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.ServicePort))
	}
	out.RawByte('}')
}