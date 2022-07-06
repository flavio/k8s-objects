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

func easyjson80489276DecodeGithubComKubewardenK8sObjectsApiAppsV1(in *jlexer.Lexer, out *ReplicaSetStatus) {
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
		case "availableReplicas":
			out.AvailableReplicas = int32(in.Int32())
		case "conditions":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]*ReplicaSetCondition, 0, 8)
					} else {
						out.Conditions = []*ReplicaSetCondition{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *ReplicaSetCondition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(ReplicaSetCondition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Conditions = append(out.Conditions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "fullyLabeledReplicas":
			out.FullyLabeledReplicas = int32(in.Int32())
		case "observedGeneration":
			out.ObservedGeneration = int64(in.Int64())
		case "readyReplicas":
			out.ReadyReplicas = int32(in.Int32())
		case "replicas":
			if in.IsNull() {
				in.Skip()
				out.Replicas = nil
			} else {
				if out.Replicas == nil {
					out.Replicas = new(int32)
				}
				*out.Replicas = int32(in.Int32())
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
func easyjson80489276EncodeGithubComKubewardenK8sObjectsApiAppsV1(out *jwriter.Writer, in ReplicaSetStatus) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AvailableReplicas != 0 {
		const prefix string = ",\"availableReplicas\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.AvailableReplicas))
	}
	if len(in.Conditions) != 0 {
		const prefix string = ",\"conditions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Conditions {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.FullyLabeledReplicas != 0 {
		const prefix string = ",\"fullyLabeledReplicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FullyLabeledReplicas))
	}
	if in.ObservedGeneration != 0 {
		const prefix string = ",\"observedGeneration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ObservedGeneration))
	}
	if in.ReadyReplicas != 0 {
		const prefix string = ",\"readyReplicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ReadyReplicas))
	}
	{
		const prefix string = ",\"replicas\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Replicas == nil {
			out.RawString("null")
		} else {
			out.Int32(int32(*in.Replicas))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ReplicaSetStatus) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson80489276EncodeGithubComKubewardenK8sObjectsApiAppsV1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ReplicaSetStatus) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson80489276EncodeGithubComKubewardenK8sObjectsApiAppsV1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ReplicaSetStatus) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson80489276DecodeGithubComKubewardenK8sObjectsApiAppsV1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ReplicaSetStatus) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson80489276DecodeGithubComKubewardenK8sObjectsApiAppsV1(l, v)
}
