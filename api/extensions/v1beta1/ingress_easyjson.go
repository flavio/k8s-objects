// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"
	_v11 "github.com/kubewarden/k8s-objects/api/core/v1"
	_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
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

func easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(in *jlexer.Lexer, out *Ingress) {
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
					out.Spec = new(IngressSpec)
				}
				easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(IngressStatus)
				}
				easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(in, out.Status)
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
func easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(out *jwriter.Writer, in Ingress) {
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
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Ingress) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Ingress) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Ingress) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Ingress) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta1(l, v)
}
func easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(in *jlexer.Lexer, out *IngressStatus) {
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
		case "loadBalancer":
			if in.IsNull() {
				in.Skip()
				out.LoadBalancer = nil
			} else {
				if out.LoadBalancer == nil {
					out.LoadBalancer = new(_v11.LoadBalancerStatus)
				}
				(*out.LoadBalancer).UnmarshalEasyJSON(in)
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
func easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta12(out *jwriter.Writer, in IngressStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.LoadBalancer != nil {
		const prefix string = ",\"loadBalancer\":"
		first = false
		out.RawString(prefix[1:])
		(*in.LoadBalancer).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(in *jlexer.Lexer, out *IngressSpec) {
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
				easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(in, out.Backend)
			}
		case "rules":
			if in.IsNull() {
				in.Skip()
				out.Rules = nil
			} else {
				in.Delim('[')
				if out.Rules == nil {
					if !in.IsDelim(']') {
						out.Rules = make([]*IngressRule, 0, 8)
					} else {
						out.Rules = []*IngressRule{}
					}
				} else {
					out.Rules = (out.Rules)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *IngressRule
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(IngressRule)
						}
						easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta14(in, v1)
					}
					out.Rules = append(out.Rules, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tls":
			if in.IsNull() {
				in.Skip()
				out.TLS = nil
			} else {
				in.Delim('[')
				if out.TLS == nil {
					if !in.IsDelim(']') {
						out.TLS = make([]*IngressTLS, 0, 8)
					} else {
						out.TLS = []*IngressTLS{}
					}
				} else {
					out.TLS = (out.TLS)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *IngressTLS
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(IngressTLS)
						}
						easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta15(in, v2)
					}
					out.TLS = append(out.TLS, v2)
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
func easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta11(out *jwriter.Writer, in IngressSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Backend != nil {
		const prefix string = ",\"backend\":"
		first = false
		out.RawString(prefix[1:])
		easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(out, *in.Backend)
	}
	if len(in.Rules) != 0 {
		const prefix string = ",\"rules\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v3, v4 := range in.Rules {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta14(out, *v4)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.TLS) != 0 {
		const prefix string = ",\"tls\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.TLS {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta15(out, *v6)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta15(in *jlexer.Lexer, out *IngressTLS) {
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
		case "hosts":
			if in.IsNull() {
				in.Skip()
				out.Hosts = nil
			} else {
				in.Delim('[')
				if out.Hosts == nil {
					if !in.IsDelim(']') {
						out.Hosts = make([]string, 0, 4)
					} else {
						out.Hosts = []string{}
					}
				} else {
					out.Hosts = (out.Hosts)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.Hosts = append(out.Hosts, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "secretName":
			out.SecretName = string(in.String())
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
func easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta15(out *jwriter.Writer, in IngressTLS) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Hosts) != 0 {
		const prefix string = ",\"hosts\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v8, v9 := range in.Hosts {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if in.SecretName != "" {
		const prefix string = ",\"secretName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SecretName))
	}
	out.RawByte('}')
}
func easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta14(in *jlexer.Lexer, out *IngressRule) {
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
		case "host":
			out.Host = string(in.String())
		case "http":
			if in.IsNull() {
				in.Skip()
				out.HTTP = nil
			} else {
				if out.HTTP == nil {
					out.HTTP = new(HTTPIngressRuleValue)
				}
				(*out.HTTP).UnmarshalEasyJSON(in)
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
func easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta14(out *jwriter.Writer, in IngressRule) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Host != "" {
		const prefix string = ",\"host\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Host))
	}
	if in.HTTP != nil {
		const prefix string = ",\"http\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.HTTP).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjson3281b84bDecodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(in *jlexer.Lexer, out *IngressBackend) {
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
		case "serviceName":
			if in.IsNull() {
				in.Skip()
				out.ServiceName = nil
			} else {
				if out.ServiceName == nil {
					out.ServiceName = new(string)
				}
				*out.ServiceName = string(in.String())
			}
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
func easyjson3281b84bEncodeGithubComKubewardenK8sObjectsApiExtensionsV1beta13(out *jwriter.Writer, in IngressBackend) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"serviceName\":"
		out.RawString(prefix[1:])
		if in.ServiceName == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ServiceName))
		}
	}
	{
		const prefix string = ",\"servicePort\":"
		out.RawString(prefix)
		if in.ServicePort == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.ServicePort))
		}
	}
	out.RawByte('}')
}
