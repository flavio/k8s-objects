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

func easyjsonE440d6eaDecodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(in *jlexer.Lexer, out *IngressStatus) {
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
			(out.LoadBalancer).UnmarshalEasyJSON(in)
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
func easyjsonE440d6eaEncodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(out *jwriter.Writer, in IngressStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"loadBalancer\":"
		first = false
		out.RawString(prefix[1:])
		(in.LoadBalancer).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IngressStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE440d6eaEncodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IngressStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE440d6eaEncodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IngressStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE440d6eaDecodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IngressStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE440d6eaDecodeGithubComKubewardenK8sObjectsApiNetworkingV1beta1(l, v)
}
