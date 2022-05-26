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

func easyjson391164eaDecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1(in *jlexer.Lexer, out *APIServiceSpec) {
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
		case "caBundle":
			if in.IsNull() {
				in.Skip()
				out.CaBundle = nil
			} else {
				out.CaBundle = in.Bytes()
			}
		case "group":
			out.Group = string(in.String())
		case "groupPriorityMinimum":
			if in.IsNull() {
				in.Skip()
				out.GroupPriorityMinimum = nil
			} else {
				if out.GroupPriorityMinimum == nil {
					out.GroupPriorityMinimum = new(int32)
				}
				*out.GroupPriorityMinimum = int32(in.Int32())
			}
		case "insecureSkipTLSVerify":
			out.InsecureSkipTLSVerify = bool(in.Bool())
		case "service":
			if in.IsNull() {
				in.Skip()
				out.Service = nil
			} else {
				if out.Service == nil {
					out.Service = new(ServiceReference)
				}
				easyjson391164eaDecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV11(in, out.Service)
			}
		case "version":
			out.Version = string(in.String())
		case "versionPriority":
			if in.IsNull() {
				in.Skip()
				out.VersionPriority = nil
			} else {
				if out.VersionPriority == nil {
					out.VersionPriority = new(int32)
				}
				*out.VersionPriority = int32(in.Int32())
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
func easyjson391164eaEncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1(out *jwriter.Writer, in APIServiceSpec) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.CaBundle) != 0 {
		const prefix string = ",\"caBundle\":"
		first = false
		out.RawString(prefix[1:])
		out.Base64Bytes(in.CaBundle)
	}
	if in.Group != "" {
		const prefix string = ",\"group\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Group))
	}
	{
		const prefix string = ",\"groupPriorityMinimum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.GroupPriorityMinimum == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.GroupPriorityMinimum))
		}
	}
	if in.InsecureSkipTLSVerify {
		const prefix string = ",\"insecureSkipTLSVerify\":"
		out.RawString(prefix)
		out.Bool(bool(in.InsecureSkipTLSVerify))
	}
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix)
		if in.Service == nil {
			out.RawString("null")
		} else {
			easyjson391164eaEncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV11(out, *in.Service)
		}
	}
	if in.Version != "" {
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"versionPriority\":"
		out.RawString(prefix)
		if in.VersionPriority == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.VersionPriority))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v APIServiceSpec) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson391164eaEncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v APIServiceSpec) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson391164eaEncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *APIServiceSpec) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson391164eaDecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *APIServiceSpec) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson391164eaDecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV1(l, v)
}
func easyjson391164eaDecodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV11(in *jlexer.Lexer, out *ServiceReference) {
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
		case "namespace":
			out.Namespace = string(in.String())
		case "port":
			out.Port = int32(in.Int32())
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
func easyjson391164eaEncodeGithubComKubewardenK8sObjectsKubeAggregatorPkgApisApiregistrationV11(out *jwriter.Writer, in ServiceReference) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Namespace != "" {
		const prefix string = ",\"namespace\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Namespace))
	}
	if in.Port != 0 {
		const prefix string = ",\"port\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Port))
	}
	out.RawByte('}')
}
