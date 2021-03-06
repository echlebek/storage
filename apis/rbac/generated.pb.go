/*
Copyright 2017 Greg Poirier

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the Software
is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
USE OR OTHER DEALINGS IN THE SOFTWARE.
*/ // Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/grepory/storage/apis/rbac/generated.proto

package rbac

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *ClusterRole) Reset()      { *m = ClusterRole{} }
func (*ClusterRole) ProtoMessage() {}
func (*ClusterRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_2686b6a78b197745, []int{0}
}
func (m *ClusterRole) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ClusterRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterRole.Merge(dst, src)
}
func (m *ClusterRole) XXX_Size() int {
	return m.Size()
}
func (m *ClusterRole) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterRole.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterRole proto.InternalMessageInfo

func (m *Group) Reset()      { *m = Group{} }
func (*Group) ProtoMessage() {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_2686b6a78b197745, []int{1}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(dst, src)
}
func (m *Group) XXX_Size() int {
	return m.Size()
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Role) Reset()      { *m = Role{} }
func (*Role) ProtoMessage() {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_2686b6a78b197745, []int{2}
}
func (m *Role) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(dst, src)
}
func (m *Role) XXX_Size() int {
	return m.Size()
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *RoleRef) Reset()      { *m = RoleRef{} }
func (*RoleRef) ProtoMessage() {}
func (*RoleRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_2686b6a78b197745, []int{3}
}
func (m *RoleRef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RoleRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *RoleRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleRef.Merge(dst, src)
}
func (m *RoleRef) XXX_Size() int {
	return m.Size()
}
func (m *RoleRef) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleRef.DiscardUnknown(m)
}

var xxx_messageInfo_RoleRef proto.InternalMessageInfo

func (m *Rule) Reset()      { *m = Rule{} }
func (*Rule) ProtoMessage() {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_2686b6a78b197745, []int{4}
}
func (m *Rule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(dst, src)
}
func (m *Rule) XXX_Size() int {
	return m.Size()
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClusterRole)(nil), "github.com.grepory.storage.apis.rbac.ClusterRole")
	proto.RegisterType((*Group)(nil), "github.com.grepory.storage.apis.rbac.Group")
	proto.RegisterType((*Role)(nil), "github.com.grepory.storage.apis.rbac.Role")
	proto.RegisterType((*RoleRef)(nil), "github.com.grepory.storage.apis.rbac.RoleRef")
	proto.RegisterType((*Rule)(nil), "github.com.grepory.storage.apis.rbac.Rule")
}
func (m *ClusterRole) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterRole) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Rules) > 0 {
		for _, msg := range m.Rules {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Group) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Group) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n2, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.RoleRefs) > 0 {
		for _, msg := range m.RoleRefs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Role) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Role) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ObjectMeta.Size()))
	n3, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.Rules) > 0 {
		for _, msg := range m.Rules {
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *RoleRef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RoleRef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Type)))
	i += copy(dAtA[i:], m.Type)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	return i, nil
}

func (m *Rule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rule) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Verbs) > 0 {
		for _, s := range m.Verbs {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Resources) > 0 {
		for _, s := range m.Resources {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Names) > 0 {
		for _, s := range m.Names {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ClusterRole) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *Group) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.RoleRefs) > 0 {
		for _, e := range m.RoleRefs {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *Role) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Rules) > 0 {
		for _, e := range m.Rules {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *RoleRef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *Rule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Verbs) > 0 {
		for _, s := range m.Verbs {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Resources) > 0 {
		for _, s := range m.Resources {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Names) > 0 {
		for _, s := range m.Names {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ClusterRole) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterRole{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "meta.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Rules:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Rules), "Rule", "Rule", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Group) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Group{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "meta.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`RoleRefs:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.RoleRefs), "RoleRef", "RoleRef", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Role) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Role{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "meta.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Rules:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Rules), "Rule", "Rule", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RoleRef) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RoleRef{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Rule) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Rule{`,
		`Verbs:` + fmt.Sprintf("%v", this.Verbs) + `,`,
		`Resources:` + fmt.Sprintf("%v", this.Resources) + `,`,
		`Names:` + fmt.Sprintf("%v", this.Names) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ClusterRole) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterRole: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterRole: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, Rule{})
			if err := m.Rules[len(m.Rules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Group) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Group: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Group: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoleRefs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoleRefs = append(m.RoleRefs, RoleRef{})
			if err := m.RoleRefs[len(m.RoleRefs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Role) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Role: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Role: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rules = append(m.Rules, Rule{})
			if err := m.Rules[len(m.Rules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RoleRef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RoleRef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RoleRef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Rule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Rule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verbs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Verbs = append(m.Verbs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resources = append(m.Resources, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Names", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Names = append(m.Names, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/grepory/storage/apis/rbac/generated.proto", fileDescriptor_generated_2686b6a78b197745)
}

var fileDescriptor_generated_2686b6a78b197745 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x91, 0x3f, 0x6f, 0xd4, 0x30,
	0x18, 0xc6, 0xe3, 0x5e, 0x22, 0x2e, 0x3e, 0x2a, 0xa1, 0x4c, 0x51, 0x07, 0x27, 0x8a, 0x18, 0x4e,
	0x45, 0x75, 0x50, 0xc5, 0x27, 0x08, 0x03, 0x53, 0xa9, 0x64, 0x21, 0x06, 0x58, 0x70, 0xd2, 0xb7,
	0xa1, 0x28, 0xa9, 0x23, 0xff, 0x41, 0xba, 0x8d, 0x8f, 0xc0, 0xd7, 0x61, 0x63, 0x42, 0x37, 0x76,
	0xec, 0x74, 0xe2, 0xc2, 0x17, 0x41, 0x4e, 0xdc, 0x06, 0x58, 0xee, 0x26, 0xa4, 0x6e, 0x79, 0x9f,
	0xf7, 0x79, 0x7e, 0x7e, 0xec, 0xe0, 0x17, 0xf5, 0x95, 0xfe, 0x68, 0x4a, 0x5a, 0x89, 0x36, 0xaf,
	0x25, 0x74, 0x42, 0xae, 0x72, 0xa5, 0x85, 0xe4, 0x35, 0xe4, 0xbc, 0xbb, 0x52, 0xb9, 0x2c, 0x79,
	0x95, 0xd7, 0x70, 0x0d, 0x92, 0x6b, 0xb8, 0xa0, 0x9d, 0x14, 0x5a, 0x44, 0x4f, 0xa7, 0x14, 0x75,
	0x29, 0xea, 0x52, 0xd4, 0xa6, 0xa8, 0x4d, 0x1d, 0x9d, 0xfc, 0xc9, 0x16, 0xb5, 0xc8, 0x87, 0x70,
	0x69, 0x2e, 0x87, 0x69, 0x18, 0x86, 0xaf, 0x11, 0x7a, 0xb4, 0xb3, 0x4a, 0x0b, 0x9a, 0xff, 0x5b,
	0x25, 0xfb, 0x8e, 0xf0, 0xe2, 0x65, 0x63, 0x94, 0x06, 0xc9, 0x44, 0x03, 0xd1, 0x07, 0x3c, 0xb7,
	0xbe, 0x0b, 0xae, 0x79, 0x8c, 0x52, 0xb4, 0x5c, 0x9c, 0x3e, 0xa7, 0xbb, 0xda, 0xda, 0x00, 0x3d,
	0x2f, 0x3f, 0x41, 0xa5, 0xcf, 0x40, 0xf3, 0x22, 0x5a, 0x6f, 0x12, 0xaf, 0xdf, 0x24, 0x78, 0xd2,
	0xd8, 0x3d, 0x35, 0x3a, 0xc7, 0x81, 0x34, 0x0d, 0xa8, 0xf8, 0x20, 0x9d, 0x2d, 0x17, 0xa7, 0xc7,
	0x74, 0x9f, 0xc7, 0xa0, 0xcc, 0x34, 0x50, 0x1c, 0x3a, 0x70, 0x60, 0x27, 0xc5, 0x46, 0x4e, 0xf6,
	0x03, 0xe1, 0xe0, 0x95, 0x14, 0xa6, 0xfb, 0x0f, 0xe5, 0xdf, 0xe3, 0xb9, 0x14, 0x0d, 0x30, 0xb8,
	0xbc, 0xeb, 0x7f, 0xb2, 0x67, 0xff, 0x31, 0x55, 0x3c, 0x71, 0xf8, 0xb9, 0x13, 0x14, 0xbb, 0x07,
	0x66, 0xdf, 0x10, 0xf6, 0x1f, 0xea, 0x4f, 0x38, 0xc3, 0x8f, 0xdc, 0x8d, 0xa2, 0x14, 0xfb, 0x7a,
	0xd5, 0xc1, 0xd0, 0x3c, 0x2c, 0x1e, 0x3b, 0xbb, 0xff, 0x66, 0xd5, 0x01, 0x1b, 0x36, 0xd6, 0x71,
	0xcd, 0x5b, 0x88, 0x0f, 0xfe, 0x76, 0xbc, 0xe6, 0x2d, 0xb0, 0x61, 0x93, 0x19, 0xec, 0x5b, 0x7c,
	0x94, 0xe0, 0xe0, 0x33, 0xc8, 0x52, 0xc5, 0x28, 0x9d, 0x2d, 0xc3, 0x22, 0xb4, 0xe7, 0xbe, 0xb5,
	0x02, 0x1b, 0xf5, 0xe8, 0x19, 0x0e, 0x25, 0x28, 0x61, 0x64, 0xe5, 0x2e, 0x13, 0x16, 0x87, 0xfd,
	0x26, 0x09, 0xd9, 0x9d, 0xc8, 0xa6, 0xbd, 0xa5, 0x59, 0xba, 0x8a, 0x67, 0x13, 0xcd, 0x1e, 0xaa,
	0xd8, 0xa8, 0x17, 0xc7, 0xeb, 0x2d, 0xf1, 0x6e, 0xb6, 0xc4, 0xbb, 0xdd, 0x12, 0xef, 0x4b, 0x4f,
	0xd0, 0xba, 0x27, 0xe8, 0xa6, 0x27, 0xe8, 0xb6, 0x27, 0xe8, 0x67, 0x4f, 0xd0, 0xd7, 0x5f, 0xc4,
	0x7b, 0xe7, 0xdb, 0xc7, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x26, 0x06, 0xd6, 0x9b, 0xfb, 0x03,
	0x00, 0x00,
}
