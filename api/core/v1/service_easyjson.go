// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package v1

import (
	json "encoding/json"
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

func easyjsonCd93bc43DecodeGithubComKubewardenK8sObjectsApiCoreV1(in *jlexer.Lexer, out *Service) {
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
					out.Spec = new(ServiceSpec)
				}
				easyjsonCd93bc43DecodeGithubComKubewardenK8sObjectsApiCoreV11(in, out.Spec)
			}
		case "status":
			if in.IsNull() {
				in.Skip()
				out.Status = nil
			} else {
				if out.Status == nil {
					out.Status = new(ServiceStatus)
				}
				easyjsonCd93bc43DecodeGithubComKubewardenK8sObjectsApiCoreV12(in, out.Status)
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
func easyjsonCd93bc43EncodeGithubComKubewardenK8sObjectsApiCoreV1(out *jwriter.Writer, in Service) {
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
		easyjsonCd93bc43EncodeGithubComKubewardenK8sObjectsApiCoreV11(out, *in.Spec)
	}
	if in.Status != nil {
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonCd93bc43EncodeGithubComKubewardenK8sObjectsApiCoreV12(out, *in.Status)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Service) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCd93bc43EncodeGithubComKubewardenK8sObjectsApiCoreV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Service) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCd93bc43EncodeGithubComKubewardenK8sObjectsApiCoreV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Service) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCd93bc43DecodeGithubComKubewardenK8sObjectsApiCoreV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Service) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCd93bc43DecodeGithubComKubewardenK8sObjectsApiCoreV1(l, v)
}
func easyjsonCd93bc43DecodeGithubComKubewardenK8sObjectsApiCoreV12(in *jlexer.Lexer, out *ServiceStatus) {
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
					out.LoadBalancer = new(LoadBalancerStatus)
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
func easyjsonCd93bc43EncodeGithubComKubewardenK8sObjectsApiCoreV12(out *jwriter.Writer, in ServiceStatus) {
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
func easyjsonCd93bc43DecodeGithubComKubewardenK8sObjectsApiCoreV11(in *jlexer.Lexer, out *ServiceSpec) {
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
		case "clusterIP":
			out.ClusterIP = string(in.String())
		case "externalIPs":
			if in.IsNull() {
				in.Skip()
				out.ExternalIPs = nil
			} else {
				in.Delim('[')
				if out.ExternalIPs == nil {
					if !in.IsDelim(']') {
						out.ExternalIPs = make([]string, 0, 4)
					} else {
						out.ExternalIPs = []string{}
					}
				} else {
					out.ExternalIPs = (out.ExternalIPs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.ExternalIPs = append(out.ExternalIPs, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "externalName":
			out.ExternalName = string(in.String())
		case "externalTrafficPolicy":
			out.ExternalTrafficPolicy = string(in.String())
		case "healthCheckNodePort":
			out.HealthCheckNodePort = int32(in.Int32())
		case "ipFamily":
			out.IPFamily = string(in.String())
		case "loadBalancerIP":
			out.LoadBalancerIP = string(in.String())
		case "loadBalancerSourceRanges":
			if in.IsNull() {
				in.Skip()
				out.LoadBalancerSourceRanges = nil
			} else {
				in.Delim('[')
				if out.LoadBalancerSourceRanges == nil {
					if !in.IsDelim(']') {
						out.LoadBalancerSourceRanges = make([]string, 0, 4)
					} else {
						out.LoadBalancerSourceRanges = []string{}
					}
				} else {
					out.LoadBalancerSourceRanges = (out.LoadBalancerSourceRanges)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.LoadBalancerSourceRanges = append(out.LoadBalancerSourceRanges, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ports":
			if in.IsNull() {
				in.Skip()
				out.Ports = nil
			} else {
				in.Delim('[')
				if out.Ports == nil {
					if !in.IsDelim(']') {
						out.Ports = make([]*ServicePort, 0, 8)
					} else {
						out.Ports = []*ServicePort{}
					}
				} else {
					out.Ports = (out.Ports)[:0]
				}
				for !in.IsDelim(']') {
					var v3 *ServicePort
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(ServicePort)
						}
						easyjsonCd93bc43DecodeGithubComKubewardenK8sObjectsApiCoreV13(in, v3)
					}
					out.Ports = append(out.Ports, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "publishNotReadyAddresses":
			out.PublishNotReadyAddresses = bool(in.Bool())
		case "selector":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Selector = make(map[string]string)
				} else {
					out.Selector = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v4 string
					v4 = string(in.String())
					(out.Selector)[key] = v4
					in.WantComma()
				}
				in.Delim('}')
			}
		case "sessionAffinity":
			out.SessionAffinity = string(in.String())
		case "sessionAffinityConfig":
			if in.IsNull() {
				in.Skip()
				out.SessionAffinityConfig = nil
			} else {
				if out.SessionAffinityConfig == nil {
					out.SessionAffinityConfig = new(SessionAffinityConfig)
				}
				easyjsonCd93bc43DecodeGithubComKubewardenK8sObjectsApiCoreV14(in, out.SessionAffinityConfig)
			}
		case "type":
			out.Type = string(in.String())
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
func easyjsonCd93bc43EncodeGithubComKubewardenK8sObjectsApiCoreV11(out *jwriter.Writer, in ServiceSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ClusterIP != "" {
		const prefix string = ",\"clusterIP\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ClusterIP))
	}
	if len(in.ExternalIPs) != 0 {
		const prefix string = ",\"externalIPs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.ExternalIPs {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	if in.ExternalName != "" {
		const prefix string = ",\"externalName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExternalName))
	}
	if in.ExternalTrafficPolicy != "" {
		const prefix string = ",\"externalTrafficPolicy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExternalTrafficPolicy))
	}
	if in.HealthCheckNodePort != 0 {
		const prefix string = ",\"healthCheckNodePort\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.HealthCheckNodePort))
	}
	if in.IPFamily != "" {
		const prefix string = ",\"ipFamily\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IPFamily))
	}
	if in.LoadBalancerIP != "" {
		const prefix string = ",\"loadBalancerIP\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LoadBalancerIP))
	}
	if len(in.LoadBalancerSourceRanges) != 0 {
		const prefix string = ",\"loadBalancerSourceRanges\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.LoadBalancerSourceRanges {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.String(string(v8))
			}
			out.RawByte(']')
		}
	}
	if len(in.Ports) != 0 {
		const prefix string = ",\"ports\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.Ports {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					easyjsonCd93bc43EncodeGithubComKubewardenK8sObjectsApiCoreV13(out, *v10)
				}
			}
			out.RawByte(']')
		}
	}
	if in.PublishNotReadyAddresses {
		const prefix string = ",\"publishNotReadyAddresses\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.PublishNotReadyAddresses))
	}
	if len(in.Selector) != 0 {
		const prefix string = ",\"selector\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v11First := true
			for v11Name, v11Value := range in.Selector {
				if v11First {
					v11First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v11Name))
				out.RawByte(':')
				out.String(string(v11Value))
			}
			out.RawByte('}')
		}
	}
	if in.SessionAffinity != "" {
		const prefix string = ",\"sessionAffinity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SessionAffinity))
	}
	if in.SessionAffinityConfig != nil {
		const prefix string = ",\"sessionAffinityConfig\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonCd93bc43EncodeGithubComKubewardenK8sObjectsApiCoreV14(out, *in.SessionAffinityConfig)
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	out.RawByte('}')
}
func easyjsonCd93bc43DecodeGithubComKubewardenK8sObjectsApiCoreV14(in *jlexer.Lexer, out *SessionAffinityConfig) {
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
		case "clientIP":
			if in.IsNull() {
				in.Skip()
				out.ClientIP = nil
			} else {
				if out.ClientIP == nil {
					out.ClientIP = new(ClientIPConfig)
				}
				(*out.ClientIP).UnmarshalEasyJSON(in)
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
func easyjsonCd93bc43EncodeGithubComKubewardenK8sObjectsApiCoreV14(out *jwriter.Writer, in SessionAffinityConfig) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ClientIP != nil {
		const prefix string = ",\"clientIP\":"
		first = false
		out.RawString(prefix[1:])
		(*in.ClientIP).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjsonCd93bc43DecodeGithubComKubewardenK8sObjectsApiCoreV13(in *jlexer.Lexer, out *ServicePort) {
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
			out.Name = string(in.String())
		case "nodePort":
			out.NodePort = int32(in.Int32())
		case "port":
			if in.IsNull() {
				in.Skip()
				out.Port = nil
			} else {
				if out.Port == nil {
					out.Port = new(int32)
				}
				*out.Port = int32(in.Int32())
			}
		case "protocol":
			out.Protocol = string(in.String())
		case "targetPort":
			if in.IsNull() {
				in.Skip()
				out.TargetPort = nil
			} else {
				if out.TargetPort == nil {
					out.TargetPort = new(intstr.IntOrString)
				}
				*out.TargetPort = intstr.IntOrString(in.String())
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
func easyjsonCd93bc43EncodeGithubComKubewardenK8sObjectsApiCoreV13(out *jwriter.Writer, in ServicePort) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.NodePort != 0 {
		const prefix string = ",\"nodePort\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.NodePort))
	}
	{
		const prefix string = ",\"port\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Port == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Port))
		}
	}
	if in.Protocol != "" {
		const prefix string = ",\"protocol\":"
		out.RawString(prefix)
		out.String(string(in.Protocol))
	}
	if in.TargetPort != nil {
		const prefix string = ",\"targetPort\":"
		out.RawString(prefix)
		out.String(string(*in.TargetPort))
	}
	out.RawByte('}')
}
