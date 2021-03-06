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
// source: github.com/grepory/storage/apis/meta/generated.proto

package meta

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

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

func (m *ObjectMeta) Reset()      { *m = ObjectMeta{} }
func (*ObjectMeta) ProtoMessage() {}
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_139764bbd714ad10, []int{0}
}
func (m *ObjectMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMeta.Merge(dst, src)
}
func (m *ObjectMeta) XXX_Size() int {
	return m.Size()
}
func (m *ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMeta proto.InternalMessageInfo

func (m *Time) Reset()      { *m = Time{} }
func (*Time) ProtoMessage() {}
func (*Time) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_139764bbd714ad10, []int{1}
}
func (m *Time) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Time.Unmarshal(m, b)
}
func (m *Time) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Time.Marshal(b, m, deterministic)
}
func (dst *Time) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Time.Merge(dst, src)
}
func (m *Time) XXX_Size() int {
	return xxx_messageInfo_Time.Size(m)
}
func (m *Time) XXX_DiscardUnknown() {
	xxx_messageInfo_Time.DiscardUnknown(m)
}

var xxx_messageInfo_Time proto.InternalMessageInfo

func (m *Timestamp) Reset()      { *m = Timestamp{} }
func (*Timestamp) ProtoMessage() {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_139764bbd714ad10, []int{2}
}
func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(dst, src)
}
func (m *Timestamp) XXX_Size() int {
	return m.Size()
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *TypeMeta) Reset()      { *m = TypeMeta{} }
func (*TypeMeta) ProtoMessage() {}
func (*TypeMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_generated_139764bbd714ad10, []int{3}
}
func (m *TypeMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TypeMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *TypeMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeMeta.Merge(dst, src)
}
func (m *TypeMeta) XXX_Size() int {
	return m.Size()
}
func (m *TypeMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeMeta.DiscardUnknown(m)
}

var xxx_messageInfo_TypeMeta proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ObjectMeta)(nil), "github.com.grepory.storage.apis.meta.ObjectMeta")
	proto.RegisterMapType((map[string]string)(nil), "github.com.grepory.storage.apis.meta.ObjectMeta.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "github.com.grepory.storage.apis.meta.ObjectMeta.LabelsEntry")
	proto.RegisterType((*Time)(nil), "github.com.grepory.storage.apis.meta.Time")
	proto.RegisterType((*Timestamp)(nil), "github.com.grepory.storage.apis.meta.Timestamp")
	proto.RegisterType((*TypeMeta)(nil), "github.com.grepory.storage.apis.meta.TypeMeta")
}
func (m *ObjectMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObjectMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Namespace)))
	i += copy(dAtA[i:], m.Namespace)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.UUID)))
	i += copy(dAtA[i:], m.UUID)
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ResourceVersion)))
	i += copy(dAtA[i:], m.ResourceVersion)
	if m.CreationTimestamp != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.CreationTimestamp.Size()))
		n1, err := m.CreationTimestamp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.DeletionTimestamp != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintGenerated(dAtA, i, uint64(m.DeletionTimestamp.Size()))
		n2, err := m.DeletionTimestamp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Labels) > 0 {
		keysForLabels := make([]string, 0, len(m.Labels))
		for k := range m.Labels {
			keysForLabels = append(keysForLabels, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
		for _, k := range keysForLabels {
			dAtA[i] = 0x42
			i++
			v := m.Labels[string(k)]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Annotations) > 0 {
		keysForAnnotations := make([]string, 0, len(m.Annotations))
		for k := range m.Annotations {
			keysForAnnotations = append(keysForAnnotations, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForAnnotations)
		for _, k := range keysForAnnotations {
			dAtA[i] = 0x4a
			i++
			v := m.Annotations[string(k)]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	dAtA[i] = 0x52
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ClusterName)))
	i += copy(dAtA[i:], m.ClusterName)
	return i, nil
}

func (m *Timestamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Timestamp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Seconds))
	dAtA[i] = 0x10
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Nanos))
	return i, nil
}

func (m *TypeMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TypeMeta) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Kind)))
	i += copy(dAtA[i:], m.Kind)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.APIVersion)))
	i += copy(dAtA[i:], m.APIVersion)
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
func (m *ObjectMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Namespace)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.UUID)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.ResourceVersion)
	n += 1 + l + sovGenerated(uint64(l))
	if m.CreationTimestamp != nil {
		l = m.CreationTimestamp.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.DeletionTimestamp != nil {
		l = m.DeletionTimestamp.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if len(m.Annotations) > 0 {
		for k, v := range m.Annotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	l = len(m.ClusterName)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *Timestamp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Seconds))
	n += 1 + sovGenerated(uint64(m.Nanos))
	return n
}

func (m *TypeMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Kind)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.APIVersion)
	n += 1 + l + sovGenerated(uint64(l))
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
func (this *ObjectMeta) String() string {
	if this == nil {
		return "nil"
	}
	keysForLabels := make([]string, 0, len(this.Labels))
	for k := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]string{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%v: %v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	keysForAnnotations := make([]string, 0, len(this.Annotations))
	for k := range this.Annotations {
		keysForAnnotations = append(keysForAnnotations, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForAnnotations)
	mapStringForAnnotations := "map[string]string{"
	for _, k := range keysForAnnotations {
		mapStringForAnnotations += fmt.Sprintf("%v: %v,", k, this.Annotations[k])
	}
	mapStringForAnnotations += "}"
	s := strings.Join([]string{`&ObjectMeta{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`UUID:` + fmt.Sprintf("%v", this.UUID) + `,`,
		`ResourceVersion:` + fmt.Sprintf("%v", this.ResourceVersion) + `,`,
		`CreationTimestamp:` + strings.Replace(fmt.Sprintf("%v", this.CreationTimestamp), "Time", "Time", 1) + `,`,
		`DeletionTimestamp:` + strings.Replace(fmt.Sprintf("%v", this.DeletionTimestamp), "Time", "Time", 1) + `,`,
		`Labels:` + mapStringForLabels + `,`,
		`Annotations:` + mapStringForAnnotations + `,`,
		`ClusterName:` + fmt.Sprintf("%v", this.ClusterName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Timestamp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Timestamp{`,
		`Seconds:` + fmt.Sprintf("%v", this.Seconds) + `,`,
		`Nanos:` + fmt.Sprintf("%v", this.Nanos) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TypeMeta) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TypeMeta{`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`APIVersion:` + fmt.Sprintf("%v", this.APIVersion) + `,`,
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
func (m *ObjectMeta) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ObjectMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObjectMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
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
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
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
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceVersion", wireType)
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
			m.ResourceVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTimestamp", wireType)
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
			if m.CreationTimestamp == nil {
				m.CreationTimestamp = &Time{}
			}
			if err := m.CreationTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletionTimestamp", wireType)
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
			if m.DeletionTimestamp == nil {
				m.DeletionTimestamp = &Time{}
			}
			if err := m.DeletionTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
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
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
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
			if m.Annotations == nil {
				m.Annotations = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Annotations[mapkey] = mapvalue
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
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
			m.ClusterName = string(dAtA[iNdEx:postIndex])
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
func (m *Timestamp) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seconds |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nanos", wireType)
			}
			m.Nanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nanos |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *TypeMeta) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TypeMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TypeMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
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
			m.Kind = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIVersion", wireType)
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
			m.APIVersion = string(dAtA[iNdEx:postIndex])
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
	proto.RegisterFile("github.com/grepory/storage/apis/meta/generated.proto", fileDescriptor_generated_139764bbd714ad10)
}

var fileDescriptor_generated_139764bbd714ad10 = []byte{
	// 594 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x4f, 0xd4, 0x40,
	0x18, 0x6e, 0xd9, 0x0f, 0xd8, 0x77, 0x55, 0x60, 0xc4, 0xd8, 0x70, 0x68, 0x37, 0xab, 0x07, 0x24,
	0xb1, 0x4d, 0x88, 0x26, 0x4a, 0x8c, 0xc9, 0x2e, 0x78, 0x20, 0x2a, 0x9a, 0x11, 0x3c, 0xe8, 0xc5,
	0xd9, 0x76, 0x5c, 0x2b, 0xed, 0x4c, 0xd3, 0x99, 0x1a, 0xf7, 0xc6, 0xd1, 0x23, 0x47, 0x8f, 0xf0,
	0x27, 0xfc, 0x0d, 0x1c, 0x39, 0x72, 0x30, 0x44, 0xea, 0x1f, 0x31, 0x9d, 0x76, 0x6d, 0x77, 0xb9,
	0x80, 0xf1, 0xd4, 0xce, 0x3c, 0x5f, 0xf3, 0xf1, 0xce, 0x0b, 0x0f, 0x86, 0xbe, 0xfc, 0x94, 0x0c,
	0x6c, 0x97, 0x87, 0xce, 0x30, 0xa6, 0x11, 0x8f, 0x47, 0x8e, 0x90, 0x3c, 0x26, 0x43, 0xea, 0x90,
	0xc8, 0x17, 0x4e, 0x48, 0x25, 0x71, 0x86, 0x94, 0xd1, 0x98, 0x48, 0xea, 0xd9, 0x51, 0xcc, 0x25,
	0x47, 0x77, 0x4b, 0x95, 0x5d, 0xa8, 0xec, 0x42, 0x65, 0x67, 0x2a, 0x3b, 0x53, 0x2d, 0xdf, 0xaf,
	0x7a, 0xf3, 0x21, 0x77, 0x94, 0x78, 0x90, 0x7c, 0x54, 0x23, 0x35, 0x50, 0x7f, 0xb9, 0x69, 0xf7,
	0x47, 0x13, 0xe0, 0xd5, 0xe0, 0x33, 0x75, 0xe5, 0x4b, 0x2a, 0x09, 0xea, 0x40, 0x9d, 0x91, 0x90,
	0x1a, 0x7a, 0x47, 0x5f, 0x69, 0xf5, 0xaf, 0x1d, 0x9f, 0x59, 0x5a, 0x7a, 0x66, 0xd5, 0xb7, 0x49,
	0x48, 0xb1, 0x42, 0x90, 0x03, 0xad, 0xec, 0x2b, 0x22, 0xe2, 0x52, 0x63, 0x46, 0xd1, 0x16, 0x0b,
	0x5a, 0x6b, 0x7b, 0x0c, 0xe0, 0x92, 0x83, 0x4c, 0xa8, 0x25, 0xbe, 0x67, 0xd4, 0x26, 0x1d, 0x77,
	0x77, 0xb7, 0x36, 0x71, 0x06, 0xa0, 0x1e, 0xcc, 0xc7, 0x54, 0xf0, 0x24, 0x76, 0xe9, 0x5b, 0x1a,
	0x0b, 0x9f, 0x33, 0xa3, 0xae, 0xb8, 0xb7, 0x0b, 0xee, 0x3c, 0x9e, 0x84, 0xf1, 0x34, 0x1f, 0x71,
	0x58, 0x74, 0x63, 0x4a, 0xa4, 0xcf, 0xd9, 0x8e, 0x1f, 0x52, 0x21, 0x49, 0x18, 0x19, 0x8d, 0x8e,
	0xbe, 0xd2, 0x5e, 0x5b, 0xb5, 0x2f, 0x73, 0x6a, 0x76, 0x26, 0xeb, 0xdf, 0x4a, 0xcf, 0xac, 0xc5,
	0x8d, 0x69, 0x23, 0x7c, 0xd1, 0x3b, 0x0b, 0xf4, 0x68, 0x40, 0x27, 0x03, 0x9b, 0xff, 0x16, 0xb8,
	0x39, 0x6d, 0x84, 0x2f, 0x7a, 0x23, 0x0f, 0x9a, 0x01, 0x19, 0xd0, 0x40, 0x18, 0x73, 0x9d, 0xda,
	0x4a, 0x7b, 0xed, 0xc9, 0xe5, 0x52, 0xca, 0x9b, 0xb5, 0x5f, 0x28, 0xf9, 0x33, 0x26, 0xe3, 0x51,
	0xff, 0x46, 0x71, 0xb2, 0xcd, 0x7c, 0x12, 0x17, 0xde, 0xe8, 0x2b, 0xb4, 0x09, 0x63, 0x5c, 0xaa,
	0xdd, 0x0a, 0xa3, 0xa5, 0xa2, 0x7a, 0x57, 0x8e, 0xea, 0x95, 0x1e, 0x79, 0xde, 0xcd, 0x22, 0xaf,
	0x5d, 0x41, 0x70, 0x35, 0x0a, 0x3d, 0x84, 0xb6, 0x1b, 0x24, 0x42, 0xd2, 0x38, 0xab, 0x21, 0x03,
	0x54, 0x01, 0xfc, 0x95, 0x6d, 0x94, 0x10, 0xae, 0xf2, 0x96, 0x1f, 0x43, 0xbb, 0xb2, 0x2f, 0xb4,
	0x00, 0xb5, 0x3d, 0x3a, 0xca, 0x8b, 0x17, 0x67, 0xbf, 0x68, 0x09, 0x1a, 0x5f, 0x48, 0x90, 0x14,
	0x95, 0x8a, 0xf3, 0xc1, 0xfa, 0xcc, 0x23, 0x7d, 0xf9, 0x29, 0x2c, 0x4c, 0xaf, 0xf3, 0x2a, 0xfa,
	0x6e, 0x00, 0xf5, 0xec, 0x7a, 0xd0, 0x3d, 0x98, 0x15, 0xd4, 0xe5, 0xcc, 0x13, 0x4a, 0x57, 0xeb,
	0xcf, 0x17, 0xab, 0x9e, 0x7d, 0x93, 0x4f, 0xe3, 0x31, 0x8e, 0xee, 0x40, 0x83, 0x11, 0xc6, 0x85,
	0x32, 0x6b, 0xf4, 0xaf, 0x17, 0xc4, 0xc6, 0x76, 0x36, 0x89, 0x73, 0x6c, 0x7d, 0xe9, 0xfb, 0xa1,
	0xa5, 0x7d, 0x3b, 0xb2, 0xb4, 0x83, 0x23, 0x4b, 0x3b, 0x3c, 0xb2, 0xb4, 0xfd, 0x9f, 0x1d, 0xad,
	0xfb, 0x1e, 0x5a, 0x65, 0x31, 0xfc, 0xe7, 0xc8, 0xee, 0x07, 0x98, 0xdb, 0x19, 0x45, 0x74, 0xdc,
	0x00, 0xf6, 0x7c, 0xe6, 0x4d, 0x37, 0x80, 0xe7, 0x3e, 0xf3, 0xb0, 0x42, 0xd0, 0x1a, 0x00, 0x89,
	0xfc, 0xf1, 0x53, 0xcd, 0x3b, 0x00, 0x2a, 0x78, 0xd0, 0x7b, 0xbd, 0x35, 0x7e, 0xa5, 0x15, 0x56,
	0x7f, 0xf5, 0xf8, 0xdc, 0xd4, 0x4e, 0xce, 0x4d, 0xed, 0xf4, 0xdc, 0xd4, 0xf6, 0x53, 0x53, 0x3f,
	0x4e, 0x4d, 0xfd, 0x24, 0x35, 0xf5, 0xd3, 0xd4, 0xd4, 0x7f, 0xa5, 0xa6, 0x7e, 0xf0, 0xdb, 0xd4,
	0xde, 0xd5, 0xb3, 0x42, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x88, 0x22, 0x31, 0x1d, 0x05,
	0x00, 0x00,
}
