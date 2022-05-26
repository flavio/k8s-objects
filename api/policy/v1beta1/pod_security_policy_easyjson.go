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

func easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(in *jlexer.Lexer, out *PodSecurityPolicy) {
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
			(out.Metadata).UnmarshalEasyJSON(in)
		case "spec":
			if in.IsNull() {
				in.Skip()
				out.Spec = nil
			} else {
				if out.Spec == nil {
					out.Spec = new(PodSecurityPolicySpec)
				}
				easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta11(in, out.Spec)
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
func easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(out *jwriter.Writer, in PodSecurityPolicy) {
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
	if true {
		const prefix string = ",\"metadata\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Metadata).MarshalEasyJSON(out)
	}
	if in.Spec != nil {
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta11(out, *in.Spec)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PodSecurityPolicy) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PodSecurityPolicy) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PodSecurityPolicy) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PodSecurityPolicy) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta1(l, v)
}
func easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta11(in *jlexer.Lexer, out *PodSecurityPolicySpec) {
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
		case "allowPrivilegeEscalation":
			out.AllowPrivilegeEscalation = bool(in.Bool())
		case "allowedCSIDrivers":
			if in.IsNull() {
				in.Skip()
				out.AllowedCSIDrivers = nil
			} else {
				in.Delim('[')
				if out.AllowedCSIDrivers == nil {
					if !in.IsDelim(']') {
						out.AllowedCSIDrivers = make([]*AllowedCSIDriver, 0, 8)
					} else {
						out.AllowedCSIDrivers = []*AllowedCSIDriver{}
					}
				} else {
					out.AllowedCSIDrivers = (out.AllowedCSIDrivers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *AllowedCSIDriver
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(AllowedCSIDriver)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.AllowedCSIDrivers = append(out.AllowedCSIDrivers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "allowedCapabilities":
			if in.IsNull() {
				in.Skip()
				out.AllowedCapabilities = nil
			} else {
				in.Delim('[')
				if out.AllowedCapabilities == nil {
					if !in.IsDelim(']') {
						out.AllowedCapabilities = make([]string, 0, 4)
					} else {
						out.AllowedCapabilities = []string{}
					}
				} else {
					out.AllowedCapabilities = (out.AllowedCapabilities)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.AllowedCapabilities = append(out.AllowedCapabilities, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "allowedFlexVolumes":
			if in.IsNull() {
				in.Skip()
				out.AllowedFlexVolumes = nil
			} else {
				in.Delim('[')
				if out.AllowedFlexVolumes == nil {
					if !in.IsDelim(']') {
						out.AllowedFlexVolumes = make([]*AllowedFlexVolume, 0, 8)
					} else {
						out.AllowedFlexVolumes = []*AllowedFlexVolume{}
					}
				} else {
					out.AllowedFlexVolumes = (out.AllowedFlexVolumes)[:0]
				}
				for !in.IsDelim(']') {
					var v3 *AllowedFlexVolume
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(AllowedFlexVolume)
						}
						(*v3).UnmarshalEasyJSON(in)
					}
					out.AllowedFlexVolumes = append(out.AllowedFlexVolumes, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "allowedHostPaths":
			if in.IsNull() {
				in.Skip()
				out.AllowedHostPaths = nil
			} else {
				in.Delim('[')
				if out.AllowedHostPaths == nil {
					if !in.IsDelim(']') {
						out.AllowedHostPaths = make([]*AllowedHostPath, 0, 8)
					} else {
						out.AllowedHostPaths = []*AllowedHostPath{}
					}
				} else {
					out.AllowedHostPaths = (out.AllowedHostPaths)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *AllowedHostPath
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(AllowedHostPath)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.AllowedHostPaths = append(out.AllowedHostPaths, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "allowedProcMountTypes":
			if in.IsNull() {
				in.Skip()
				out.AllowedProcMountTypes = nil
			} else {
				in.Delim('[')
				if out.AllowedProcMountTypes == nil {
					if !in.IsDelim(']') {
						out.AllowedProcMountTypes = make([]string, 0, 4)
					} else {
						out.AllowedProcMountTypes = []string{}
					}
				} else {
					out.AllowedProcMountTypes = (out.AllowedProcMountTypes)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.AllowedProcMountTypes = append(out.AllowedProcMountTypes, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "allowedUnsafeSysctls":
			if in.IsNull() {
				in.Skip()
				out.AllowedUnsafeSysctls = nil
			} else {
				in.Delim('[')
				if out.AllowedUnsafeSysctls == nil {
					if !in.IsDelim(']') {
						out.AllowedUnsafeSysctls = make([]string, 0, 4)
					} else {
						out.AllowedUnsafeSysctls = []string{}
					}
				} else {
					out.AllowedUnsafeSysctls = (out.AllowedUnsafeSysctls)[:0]
				}
				for !in.IsDelim(']') {
					var v6 string
					v6 = string(in.String())
					out.AllowedUnsafeSysctls = append(out.AllowedUnsafeSysctls, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "defaultAddCapabilities":
			if in.IsNull() {
				in.Skip()
				out.DefaultAddCapabilities = nil
			} else {
				in.Delim('[')
				if out.DefaultAddCapabilities == nil {
					if !in.IsDelim(']') {
						out.DefaultAddCapabilities = make([]string, 0, 4)
					} else {
						out.DefaultAddCapabilities = []string{}
					}
				} else {
					out.DefaultAddCapabilities = (out.DefaultAddCapabilities)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.DefaultAddCapabilities = append(out.DefaultAddCapabilities, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "defaultAllowPrivilegeEscalation":
			out.DefaultAllowPrivilegeEscalation = bool(in.Bool())
		case "forbiddenSysctls":
			if in.IsNull() {
				in.Skip()
				out.ForbiddenSysctls = nil
			} else {
				in.Delim('[')
				if out.ForbiddenSysctls == nil {
					if !in.IsDelim(']') {
						out.ForbiddenSysctls = make([]string, 0, 4)
					} else {
						out.ForbiddenSysctls = []string{}
					}
				} else {
					out.ForbiddenSysctls = (out.ForbiddenSysctls)[:0]
				}
				for !in.IsDelim(']') {
					var v8 string
					v8 = string(in.String())
					out.ForbiddenSysctls = append(out.ForbiddenSysctls, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "fsGroup":
			if in.IsNull() {
				in.Skip()
				out.FsGroup = nil
			} else {
				if out.FsGroup == nil {
					out.FsGroup = new(FSGroupStrategyOptions)
				}
				(*out.FsGroup).UnmarshalEasyJSON(in)
			}
		case "hostIPC":
			out.HostIPC = bool(in.Bool())
		case "hostNetwork":
			out.HostNetwork = bool(in.Bool())
		case "hostPID":
			out.HostPID = bool(in.Bool())
		case "hostPorts":
			if in.IsNull() {
				in.Skip()
				out.HostPorts = nil
			} else {
				in.Delim('[')
				if out.HostPorts == nil {
					if !in.IsDelim(']') {
						out.HostPorts = make([]*HostPortRange, 0, 8)
					} else {
						out.HostPorts = []*HostPortRange{}
					}
				} else {
					out.HostPorts = (out.HostPorts)[:0]
				}
				for !in.IsDelim(']') {
					var v9 *HostPortRange
					if in.IsNull() {
						in.Skip()
						v9 = nil
					} else {
						if v9 == nil {
							v9 = new(HostPortRange)
						}
						(*v9).UnmarshalEasyJSON(in)
					}
					out.HostPorts = append(out.HostPorts, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "privileged":
			out.Privileged = bool(in.Bool())
		case "readOnlyRootFilesystem":
			out.ReadOnlyRootFilesystem = bool(in.Bool())
		case "requiredDropCapabilities":
			if in.IsNull() {
				in.Skip()
				out.RequiredDropCapabilities = nil
			} else {
				in.Delim('[')
				if out.RequiredDropCapabilities == nil {
					if !in.IsDelim(']') {
						out.RequiredDropCapabilities = make([]string, 0, 4)
					} else {
						out.RequiredDropCapabilities = []string{}
					}
				} else {
					out.RequiredDropCapabilities = (out.RequiredDropCapabilities)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.RequiredDropCapabilities = append(out.RequiredDropCapabilities, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "runAsGroup":
			if in.IsNull() {
				in.Skip()
				out.RunAsGroup = nil
			} else {
				if out.RunAsGroup == nil {
					out.RunAsGroup = new(RunAsGroupStrategyOptions)
				}
				easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta12(in, out.RunAsGroup)
			}
		case "runAsUser":
			if in.IsNull() {
				in.Skip()
				out.RunAsUser = nil
			} else {
				if out.RunAsUser == nil {
					out.RunAsUser = new(RunAsUserStrategyOptions)
				}
				easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta13(in, out.RunAsUser)
			}
		case "runtimeClass":
			if in.IsNull() {
				in.Skip()
				out.RuntimeClass = nil
			} else {
				if out.RuntimeClass == nil {
					out.RuntimeClass = new(RuntimeClassStrategyOptions)
				}
				easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta14(in, out.RuntimeClass)
			}
		case "seLinux":
			if in.IsNull() {
				in.Skip()
				out.SeLinux = nil
			} else {
				if out.SeLinux == nil {
					out.SeLinux = new(SELinuxStrategyOptions)
				}
				easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta15(in, out.SeLinux)
			}
		case "supplementalGroups":
			if in.IsNull() {
				in.Skip()
				out.SupplementalGroups = nil
			} else {
				if out.SupplementalGroups == nil {
					out.SupplementalGroups = new(SupplementalGroupsStrategyOptions)
				}
				easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta16(in, out.SupplementalGroups)
			}
		case "volumes":
			if in.IsNull() {
				in.Skip()
				out.Volumes = nil
			} else {
				in.Delim('[')
				if out.Volumes == nil {
					if !in.IsDelim(']') {
						out.Volumes = make([]string, 0, 4)
					} else {
						out.Volumes = []string{}
					}
				} else {
					out.Volumes = (out.Volumes)[:0]
				}
				for !in.IsDelim(']') {
					var v11 string
					v11 = string(in.String())
					out.Volumes = append(out.Volumes, v11)
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
func easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta11(out *jwriter.Writer, in PodSecurityPolicySpec) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AllowPrivilegeEscalation {
		const prefix string = ",\"allowPrivilegeEscalation\":"
		first = false
		out.RawString(prefix[1:])
		out.Bool(bool(in.AllowPrivilegeEscalation))
	}
	{
		const prefix string = ",\"allowedCSIDrivers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.AllowedCSIDrivers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.AllowedCSIDrivers {
				if v12 > 0 {
					out.RawByte(',')
				}
				if v13 == nil {
					out.RawString("null")
				} else {
					(*v13).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"allowedCapabilities\":"
		out.RawString(prefix)
		if in.AllowedCapabilities == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.AllowedCapabilities {
				if v14 > 0 {
					out.RawByte(',')
				}
				out.String(string(v15))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"allowedFlexVolumes\":"
		out.RawString(prefix)
		if in.AllowedFlexVolumes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.AllowedFlexVolumes {
				if v16 > 0 {
					out.RawByte(',')
				}
				if v17 == nil {
					out.RawString("null")
				} else {
					(*v17).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"allowedHostPaths\":"
		out.RawString(prefix)
		if in.AllowedHostPaths == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v18, v19 := range in.AllowedHostPaths {
				if v18 > 0 {
					out.RawByte(',')
				}
				if v19 == nil {
					out.RawString("null")
				} else {
					(*v19).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"allowedProcMountTypes\":"
		out.RawString(prefix)
		if in.AllowedProcMountTypes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.AllowedProcMountTypes {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.String(string(v21))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"allowedUnsafeSysctls\":"
		out.RawString(prefix)
		if in.AllowedUnsafeSysctls == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v22, v23 := range in.AllowedUnsafeSysctls {
				if v22 > 0 {
					out.RawByte(',')
				}
				out.String(string(v23))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"defaultAddCapabilities\":"
		out.RawString(prefix)
		if in.DefaultAddCapabilities == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v24, v25 := range in.DefaultAddCapabilities {
				if v24 > 0 {
					out.RawByte(',')
				}
				out.String(string(v25))
			}
			out.RawByte(']')
		}
	}
	if in.DefaultAllowPrivilegeEscalation {
		const prefix string = ",\"defaultAllowPrivilegeEscalation\":"
		out.RawString(prefix)
		out.Bool(bool(in.DefaultAllowPrivilegeEscalation))
	}
	{
		const prefix string = ",\"forbiddenSysctls\":"
		out.RawString(prefix)
		if in.ForbiddenSysctls == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v26, v27 := range in.ForbiddenSysctls {
				if v26 > 0 {
					out.RawByte(',')
				}
				out.String(string(v27))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"fsGroup\":"
		out.RawString(prefix)
		if in.FsGroup == nil {
			out.RawString("null")
		} else {
			(*in.FsGroup).MarshalEasyJSON(out)
		}
	}
	if in.HostIPC {
		const prefix string = ",\"hostIPC\":"
		out.RawString(prefix)
		out.Bool(bool(in.HostIPC))
	}
	if in.HostNetwork {
		const prefix string = ",\"hostNetwork\":"
		out.RawString(prefix)
		out.Bool(bool(in.HostNetwork))
	}
	if in.HostPID {
		const prefix string = ",\"hostPID\":"
		out.RawString(prefix)
		out.Bool(bool(in.HostPID))
	}
	{
		const prefix string = ",\"hostPorts\":"
		out.RawString(prefix)
		if in.HostPorts == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v28, v29 := range in.HostPorts {
				if v28 > 0 {
					out.RawByte(',')
				}
				if v29 == nil {
					out.RawString("null")
				} else {
					(*v29).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Privileged {
		const prefix string = ",\"privileged\":"
		out.RawString(prefix)
		out.Bool(bool(in.Privileged))
	}
	if in.ReadOnlyRootFilesystem {
		const prefix string = ",\"readOnlyRootFilesystem\":"
		out.RawString(prefix)
		out.Bool(bool(in.ReadOnlyRootFilesystem))
	}
	{
		const prefix string = ",\"requiredDropCapabilities\":"
		out.RawString(prefix)
		if in.RequiredDropCapabilities == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v30, v31 := range in.RequiredDropCapabilities {
				if v30 > 0 {
					out.RawByte(',')
				}
				out.String(string(v31))
			}
			out.RawByte(']')
		}
	}
	if in.RunAsGroup != nil {
		const prefix string = ",\"runAsGroup\":"
		out.RawString(prefix)
		easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta12(out, *in.RunAsGroup)
	}
	{
		const prefix string = ",\"runAsUser\":"
		out.RawString(prefix)
		if in.RunAsUser == nil {
			out.RawString("null")
		} else {
			easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta13(out, *in.RunAsUser)
		}
	}
	if in.RuntimeClass != nil {
		const prefix string = ",\"runtimeClass\":"
		out.RawString(prefix)
		easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta14(out, *in.RuntimeClass)
	}
	{
		const prefix string = ",\"seLinux\":"
		out.RawString(prefix)
		if in.SeLinux == nil {
			out.RawString("null")
		} else {
			easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta15(out, *in.SeLinux)
		}
	}
	{
		const prefix string = ",\"supplementalGroups\":"
		out.RawString(prefix)
		if in.SupplementalGroups == nil {
			out.RawString("null")
		} else {
			easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta16(out, *in.SupplementalGroups)
		}
	}
	{
		const prefix string = ",\"volumes\":"
		out.RawString(prefix)
		if in.Volumes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v32, v33 := range in.Volumes {
				if v32 > 0 {
					out.RawByte(',')
				}
				out.String(string(v33))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta16(in *jlexer.Lexer, out *SupplementalGroupsStrategyOptions) {
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
		case "ranges":
			if in.IsNull() {
				in.Skip()
				out.Ranges = nil
			} else {
				in.Delim('[')
				if out.Ranges == nil {
					if !in.IsDelim(']') {
						out.Ranges = make([]*IDRange, 0, 8)
					} else {
						out.Ranges = []*IDRange{}
					}
				} else {
					out.Ranges = (out.Ranges)[:0]
				}
				for !in.IsDelim(']') {
					var v34 *IDRange
					if in.IsNull() {
						in.Skip()
						v34 = nil
					} else {
						if v34 == nil {
							v34 = new(IDRange)
						}
						(*v34).UnmarshalEasyJSON(in)
					}
					out.Ranges = append(out.Ranges, v34)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "rule":
			out.Rule = string(in.String())
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
func easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta16(out *jwriter.Writer, in SupplementalGroupsStrategyOptions) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ranges\":"
		out.RawString(prefix[1:])
		if in.Ranges == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v35, v36 := range in.Ranges {
				if v35 > 0 {
					out.RawByte(',')
				}
				if v36 == nil {
					out.RawString("null")
				} else {
					(*v36).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Rule != "" {
		const prefix string = ",\"rule\":"
		out.RawString(prefix)
		out.String(string(in.Rule))
	}
	out.RawByte('}')
}
func easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta15(in *jlexer.Lexer, out *SELinuxStrategyOptions) {
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
		case "rule":
			if in.IsNull() {
				in.Skip()
				out.Rule = nil
			} else {
				if out.Rule == nil {
					out.Rule = new(string)
				}
				*out.Rule = string(in.String())
			}
		case "seLinuxOptions":
			(out.SeLinuxOptions).UnmarshalEasyJSON(in)
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
func easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta15(out *jwriter.Writer, in SELinuxStrategyOptions) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rule\":"
		out.RawString(prefix[1:])
		if in.Rule == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Rule))
		}
	}
	if true {
		const prefix string = ",\"seLinuxOptions\":"
		out.RawString(prefix)
		(in.SeLinuxOptions).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}
func easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta14(in *jlexer.Lexer, out *RuntimeClassStrategyOptions) {
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
		case "allowedRuntimeClassNames":
			if in.IsNull() {
				in.Skip()
				out.AllowedRuntimeClassNames = nil
			} else {
				in.Delim('[')
				if out.AllowedRuntimeClassNames == nil {
					if !in.IsDelim(']') {
						out.AllowedRuntimeClassNames = make([]string, 0, 4)
					} else {
						out.AllowedRuntimeClassNames = []string{}
					}
				} else {
					out.AllowedRuntimeClassNames = (out.AllowedRuntimeClassNames)[:0]
				}
				for !in.IsDelim(']') {
					var v37 string
					v37 = string(in.String())
					out.AllowedRuntimeClassNames = append(out.AllowedRuntimeClassNames, v37)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "defaultRuntimeClassName":
			out.DefaultRuntimeClassName = string(in.String())
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
func easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta14(out *jwriter.Writer, in RuntimeClassStrategyOptions) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"allowedRuntimeClassNames\":"
		out.RawString(prefix[1:])
		if in.AllowedRuntimeClassNames == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v38, v39 := range in.AllowedRuntimeClassNames {
				if v38 > 0 {
					out.RawByte(',')
				}
				out.String(string(v39))
			}
			out.RawByte(']')
		}
	}
	if in.DefaultRuntimeClassName != "" {
		const prefix string = ",\"defaultRuntimeClassName\":"
		out.RawString(prefix)
		out.String(string(in.DefaultRuntimeClassName))
	}
	out.RawByte('}')
}
func easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta13(in *jlexer.Lexer, out *RunAsUserStrategyOptions) {
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
		case "ranges":
			if in.IsNull() {
				in.Skip()
				out.Ranges = nil
			} else {
				in.Delim('[')
				if out.Ranges == nil {
					if !in.IsDelim(']') {
						out.Ranges = make([]*IDRange, 0, 8)
					} else {
						out.Ranges = []*IDRange{}
					}
				} else {
					out.Ranges = (out.Ranges)[:0]
				}
				for !in.IsDelim(']') {
					var v40 *IDRange
					if in.IsNull() {
						in.Skip()
						v40 = nil
					} else {
						if v40 == nil {
							v40 = new(IDRange)
						}
						(*v40).UnmarshalEasyJSON(in)
					}
					out.Ranges = append(out.Ranges, v40)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "rule":
			if in.IsNull() {
				in.Skip()
				out.Rule = nil
			} else {
				if out.Rule == nil {
					out.Rule = new(string)
				}
				*out.Rule = string(in.String())
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
func easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta13(out *jwriter.Writer, in RunAsUserStrategyOptions) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ranges\":"
		out.RawString(prefix[1:])
		if in.Ranges == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v41, v42 := range in.Ranges {
				if v41 > 0 {
					out.RawByte(',')
				}
				if v42 == nil {
					out.RawString("null")
				} else {
					(*v42).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"rule\":"
		out.RawString(prefix)
		if in.Rule == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Rule))
		}
	}
	out.RawByte('}')
}
func easyjsonA70dd7d9DecodeGithubComKubewardenK8sObjectsApiPolicyV1beta12(in *jlexer.Lexer, out *RunAsGroupStrategyOptions) {
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
		case "ranges":
			if in.IsNull() {
				in.Skip()
				out.Ranges = nil
			} else {
				in.Delim('[')
				if out.Ranges == nil {
					if !in.IsDelim(']') {
						out.Ranges = make([]*IDRange, 0, 8)
					} else {
						out.Ranges = []*IDRange{}
					}
				} else {
					out.Ranges = (out.Ranges)[:0]
				}
				for !in.IsDelim(']') {
					var v43 *IDRange
					if in.IsNull() {
						in.Skip()
						v43 = nil
					} else {
						if v43 == nil {
							v43 = new(IDRange)
						}
						(*v43).UnmarshalEasyJSON(in)
					}
					out.Ranges = append(out.Ranges, v43)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "rule":
			if in.IsNull() {
				in.Skip()
				out.Rule = nil
			} else {
				if out.Rule == nil {
					out.Rule = new(string)
				}
				*out.Rule = string(in.String())
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
func easyjsonA70dd7d9EncodeGithubComKubewardenK8sObjectsApiPolicyV1beta12(out *jwriter.Writer, in RunAsGroupStrategyOptions) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ranges\":"
		out.RawString(prefix[1:])
		if in.Ranges == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v44, v45 := range in.Ranges {
				if v44 > 0 {
					out.RawByte(',')
				}
				if v45 == nil {
					out.RawString("null")
				} else {
					(*v45).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"rule\":"
		out.RawString(prefix)
		if in.Rule == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Rule))
		}
	}
	out.RawByte('}')
}
