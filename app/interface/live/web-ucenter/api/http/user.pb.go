// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user.proto

package http

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import encoding_binary "encoding/binary"

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

// GetInfoReq get user info req, uid and platform get from header.metadata
type GetInfoReq struct {
	// platform in url
	Platform             string   `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty" form:"platform"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoReq) Reset()         { *m = GetInfoReq{} }
func (m *GetInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetInfoReq) ProtoMessage()    {}
func (*GetInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_7ff157935b3a3e82, []int{0}
}
func (m *GetInfoReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetInfoReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoReq.Merge(dst, src)
}
func (m *GetInfoReq) XXX_Size() int {
	return m.Size()
}
func (m *GetInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoReq proto.InternalMessageInfo

func (m *GetInfoReq) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

// GetInfoResp
type GetInfoResp struct {
	// 用户uid
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid"`
	// 用户名
	Uname string `protobuf:"bytes,2,opt,name=uname,proto3" json:"uname"`
	// 头像
	Face string `protobuf:"bytes,3,opt,name=face,proto3" json:"face"`
	// 主站硬币
	Coin float64 `protobuf:"fixed64,4,opt,name=coin,proto3" json:"billCoin"`
	// 用户银瓜子
	Silver int64 `protobuf:"varint,5,opt,name=silver,proto3" json:"silver"`
	// 用户金瓜子
	Gold int64 `protobuf:"varint,6,opt,name=gold,proto3" json:"gold"`
	// 用户成就点
	Achieve int64 `protobuf:"varint,7,opt,name=achieve,proto3" json:"achieve"`
	// 月费姥爷
	Vip int `protobuf:"varint,8,opt,name=vip,proto3,casttype=int" json:"vip"`
	// 年费姥爷
	Svip int `protobuf:"varint,9,opt,name=svip,proto3,casttype=int" json:"svip"`
	// 用户等级
	UserLevel int64 `protobuf:"varint,10,opt,name=user_level,json=userLevel,proto3" json:"user_level"`
	// 用户下一等级
	UserNextLevel int64 `protobuf:"varint,11,opt,name=user_next_level,json=userNextLevel,proto3" json:"user_next_level"`
	// 用户在当前等级已经获得的经验
	UserIntimacy int64 `protobuf:"varint,12,opt,name=user_intimacy,json=userIntimacy,proto3" json:"user_intimacy"`
	// 用户从当前等级升级到下一级所需总经验
	UserNextIntimacy int64 `protobuf:"varint,13,opt,name=user_next_intimacy,json=userNextIntimacy,proto3" json:"user_next_intimacy"`
	// 新增字段，判断用户是否达到满级 0:没有1:满级
	IsLevelTop int64 `protobuf:"varint,14,opt,name=is_level_top,json=isLevelTop,proto3" json:"is_level_top"`
	// 用户等级排名
	UserLevelRank string `protobuf:"bytes,15,opt,name=user_level_rank,json=userLevelRank,proto3" json:"user_level_rank"`
	// 年返逻辑，已失效
	UserCharged          int      `protobuf:"varint,16,opt,name=user_charged,json=userCharged,proto3,casttype=int" json:"user_charged"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoResp) Reset()         { *m = GetInfoResp{} }
func (m *GetInfoResp) String() string { return proto.CompactTextString(m) }
func (*GetInfoResp) ProtoMessage()    {}
func (*GetInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_7ff157935b3a3e82, []int{1}
}
func (m *GetInfoResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetInfoResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoResp.Merge(dst, src)
}
func (m *GetInfoResp) XXX_Size() int {
	return m.Size()
}
func (m *GetInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoResp proto.InternalMessageInfo

func (m *GetInfoResp) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *GetInfoResp) GetUname() string {
	if m != nil {
		return m.Uname
	}
	return ""
}

func (m *GetInfoResp) GetFace() string {
	if m != nil {
		return m.Face
	}
	return ""
}

func (m *GetInfoResp) GetCoin() float64 {
	if m != nil {
		return m.Coin
	}
	return 0
}

func (m *GetInfoResp) GetSilver() int64 {
	if m != nil {
		return m.Silver
	}
	return 0
}

func (m *GetInfoResp) GetGold() int64 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *GetInfoResp) GetAchieve() int64 {
	if m != nil {
		return m.Achieve
	}
	return 0
}

func (m *GetInfoResp) GetVip() int {
	if m != nil {
		return m.Vip
	}
	return 0
}

func (m *GetInfoResp) GetSvip() int {
	if m != nil {
		return m.Svip
	}
	return 0
}

func (m *GetInfoResp) GetUserLevel() int64 {
	if m != nil {
		return m.UserLevel
	}
	return 0
}

func (m *GetInfoResp) GetUserNextLevel() int64 {
	if m != nil {
		return m.UserNextLevel
	}
	return 0
}

func (m *GetInfoResp) GetUserIntimacy() int64 {
	if m != nil {
		return m.UserIntimacy
	}
	return 0
}

func (m *GetInfoResp) GetUserNextIntimacy() int64 {
	if m != nil {
		return m.UserNextIntimacy
	}
	return 0
}

func (m *GetInfoResp) GetIsLevelTop() int64 {
	if m != nil {
		return m.IsLevelTop
	}
	return 0
}

func (m *GetInfoResp) GetUserLevelRank() string {
	if m != nil {
		return m.UserLevelRank
	}
	return ""
}

func (m *GetInfoResp) GetUserCharged() int {
	if m != nil {
		return m.UserCharged
	}
	return 0
}

func init() {
	proto.RegisterType((*GetInfoReq)(nil), "live.webucenter.GetInfoReq")
	proto.RegisterType((*GetInfoResp)(nil), "live.webucenter.GetInfoResp")
}
func (m *GetInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetInfoReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Platform) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Platform)))
		i += copy(dAtA[i:], m.Platform)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GetInfoResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetInfoResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Uid))
	}
	if len(m.Uname) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Uname)))
		i += copy(dAtA[i:], m.Uname)
	}
	if len(m.Face) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Face)))
		i += copy(dAtA[i:], m.Face)
	}
	if m.Coin != 0 {
		dAtA[i] = 0x21
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Coin))))
		i += 8
	}
	if m.Silver != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Silver))
	}
	if m.Gold != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Gold))
	}
	if m.Achieve != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Achieve))
	}
	if m.Vip != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Vip))
	}
	if m.Svip != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Svip))
	}
	if m.UserLevel != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.UserLevel))
	}
	if m.UserNextLevel != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.UserNextLevel))
	}
	if m.UserIntimacy != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.UserIntimacy))
	}
	if m.UserNextIntimacy != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.UserNextIntimacy))
	}
	if m.IsLevelTop != 0 {
		dAtA[i] = 0x70
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.IsLevelTop))
	}
	if len(m.UserLevelRank) > 0 {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.UserLevelRank)))
		i += copy(dAtA[i:], m.UserLevelRank)
	}
	if m.UserCharged != 0 {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.UserCharged))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetInfoReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Platform)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetInfoResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovUser(uint64(m.Uid))
	}
	l = len(m.Uname)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Face)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.Coin != 0 {
		n += 9
	}
	if m.Silver != 0 {
		n += 1 + sovUser(uint64(m.Silver))
	}
	if m.Gold != 0 {
		n += 1 + sovUser(uint64(m.Gold))
	}
	if m.Achieve != 0 {
		n += 1 + sovUser(uint64(m.Achieve))
	}
	if m.Vip != 0 {
		n += 1 + sovUser(uint64(m.Vip))
	}
	if m.Svip != 0 {
		n += 1 + sovUser(uint64(m.Svip))
	}
	if m.UserLevel != 0 {
		n += 1 + sovUser(uint64(m.UserLevel))
	}
	if m.UserNextLevel != 0 {
		n += 1 + sovUser(uint64(m.UserNextLevel))
	}
	if m.UserIntimacy != 0 {
		n += 1 + sovUser(uint64(m.UserIntimacy))
	}
	if m.UserNextIntimacy != 0 {
		n += 1 + sovUser(uint64(m.UserNextIntimacy))
	}
	if m.IsLevelTop != 0 {
		n += 1 + sovUser(uint64(m.IsLevelTop))
	}
	l = len(m.UserLevelRank)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.UserCharged != 0 {
		n += 2 + sovUser(uint64(m.UserCharged))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUser(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: GetInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platform", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Platform = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetInfoResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: GetInfoResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetInfoResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Face", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Face = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Coin = float64(math.Float64frombits(v))
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Silver", wireType)
			}
			m.Silver = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Silver |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gold", wireType)
			}
			m.Gold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gold |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Achieve", wireType)
			}
			m.Achieve = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Achieve |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vip", wireType)
			}
			m.Vip = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vip |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Svip", wireType)
			}
			m.Svip = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Svip |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserLevel", wireType)
			}
			m.UserLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserLevel |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserNextLevel", wireType)
			}
			m.UserNextLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserNextLevel |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserIntimacy", wireType)
			}
			m.UserIntimacy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserIntimacy |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserNextIntimacy", wireType)
			}
			m.UserNextIntimacy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserNextIntimacy |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLevelTop", wireType)
			}
			m.IsLevelTop = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsLevelTop |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserLevelRank", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserLevelRank = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserCharged", wireType)
			}
			m.UserCharged = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserCharged |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUser
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
				next, err := skipUser(dAtA[start:])
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
	ErrInvalidLengthUser = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_7ff157935b3a3e82) }

var fileDescriptor_user_7ff157935b3a3e82 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0x68, 0xba, 0xb5, 0xaf, 0xdd, 0x5a, 0x3c, 0x09, 0xcc, 0x98, 0xea, 0xaa, 0x08, 0xa9,
	0x97, 0xb5, 0xd2, 0x90, 0x38, 0x0c, 0x71, 0xe9, 0x90, 0xd0, 0x10, 0xe2, 0x60, 0xc1, 0x85, 0x4b,
	0x95, 0xa6, 0x6e, 0x6a, 0x2d, 0x8d, 0x43, 0xe2, 0x84, 0xf1, 0x0f, 0x39, 0xee, 0x17, 0x58, 0xa8,
	0xc7, 0x1c, 0x39, 0x72, 0x42, 0x7e, 0x69, 0xd3, 0x6e, 0x82, 0x8b, 0xfd, 0xbe, 0xef, 0x7b, 0xdf,
	0x7b, 0x7e, 0x71, 0x0c, 0x90, 0xa5, 0x22, 0x19, 0xc5, 0x89, 0xd2, 0x8a, 0x74, 0x42, 0x99, 0x8b,
	0xd1, 0x77, 0x31, 0xcb, 0x7c, 0x11, 0x69, 0x91, 0x9c, 0x9e, 0x07, 0x52, 0x2f, 0xb3, 0xd9, 0xc8,
	0x57, 0xab, 0x71, 0xa0, 0x02, 0x35, 0xc6, 0xbc, 0x59, 0xb6, 0x40, 0x84, 0x00, 0xa3, 0xd2, 0x3f,
	0x78, 0x0b, 0xf0, 0x5e, 0xe8, 0xeb, 0x68, 0xa1, 0xb8, 0xf8, 0x46, 0xc6, 0xd0, 0x88, 0x43, 0x4f,
	0x2f, 0x54, 0xb2, 0xa2, 0x4e, 0xdf, 0x19, 0x36, 0x27, 0x27, 0xbf, 0x0d, 0xeb, 0x58, 0x7c, 0x39,
	0xd8, 0x2a, 0x03, 0x5e, 0x25, 0x0d, 0xee, 0xea, 0xd0, 0xaa, 0xfc, 0x69, 0x4c, 0x9e, 0x41, 0x2d,
	0x93, 0x73, 0xf4, 0xd6, 0x26, 0x87, 0x85, 0x61, 0x16, 0x72, 0xbb, 0x10, 0x06, 0xf5, 0x2c, 0xf2,
	0x56, 0x82, 0x3e, 0xc2, 0xc2, 0xcd, 0xc2, 0xb0, 0x92, 0xe0, 0xe5, 0x46, 0xce, 0xc0, 0x5d, 0x78,
	0xbe, 0xa0, 0x35, 0xd4, 0x1b, 0x85, 0x61, 0x88, 0x39, 0xae, 0xa4, 0x0f, 0xae, 0xaf, 0x64, 0x44,
	0xdd, 0xbe, 0x33, 0x74, 0x26, 0xed, 0xc2, 0xb0, 0xc6, 0x4c, 0x86, 0xe1, 0x95, 0x92, 0x11, 0x47,
	0x85, 0x0c, 0xe0, 0x20, 0x95, 0x61, 0x2e, 0x12, 0x5a, 0xc7, 0xf6, 0x50, 0x18, 0xb6, 0x61, 0xf8,
	0x66, 0xb7, 0x3d, 0x02, 0x15, 0xce, 0xe9, 0x01, 0x66, 0x60, 0x0f, 0x8b, 0x39, 0xae, 0xe4, 0x25,
	0x1c, 0x7a, 0xfe, 0x52, 0x8a, 0x5c, 0xd0, 0x43, 0x4c, 0x68, 0x15, 0x86, 0x6d, 0x29, 0xbe, 0x0d,
	0x48, 0x1f, 0x6a, 0xb9, 0x8c, 0x69, 0xa3, 0xef, 0x0c, 0xeb, 0x93, 0x63, 0x3b, 0x64, 0x2e, 0xe3,
	0x3f, 0x86, 0xd5, 0x64, 0xa4, 0xb9, 0x8d, 0xc9, 0x0b, 0x70, 0x53, 0x9b, 0xd2, 0xc4, 0x94, 0x8e,
	0x6d, 0x93, 0xee, 0xe5, 0x20, 0x20, 0xe7, 0xe5, 0x45, 0x4e, 0x43, 0x91, 0x8b, 0x90, 0x02, 0x36,
	0xb4, 0xd5, 0xf6, 0x58, 0xde, 0xb4, 0xf1, 0x47, 0x1b, 0x92, 0x37, 0xd0, 0x41, 0x21, 0x12, 0xb7,
	0x7a, 0xe3, 0x69, 0xa1, 0xe7, 0xa4, 0x30, 0xec, 0xa1, 0xc4, 0x8f, 0x2c, 0xf1, 0x49, 0xdc, 0xea,
	0xd2, 0xfc, 0x1a, 0x90, 0x98, 0xca, 0x48, 0xcb, 0x95, 0xe7, 0xff, 0xa0, 0x6d, 0xb4, 0x3e, 0x2e,
	0x0c, 0xbb, 0x2f, 0xf0, 0xb6, 0x85, 0xd7, 0x1b, 0x44, 0xde, 0x01, 0xd9, 0x55, 0xae, 0xcc, 0x47,
	0x68, 0x7e, 0x52, 0x18, 0xf6, 0x0f, 0x95, 0x77, 0xb7, 0xad, 0xab, 0x2a, 0x17, 0xd0, 0x96, 0x69,
	0x79, 0xb0, 0xa9, 0x56, 0x31, 0x3d, 0x46, 0x7f, 0xb7, 0x30, 0xec, 0x1e, 0xcf, 0x41, 0xa6, 0x78,
	0xdc, 0xcf, 0x2a, 0xae, 0xc6, 0x2d, 0xd5, 0xc4, 0x8b, 0x6e, 0x68, 0xa7, 0xfc, 0x23, 0xab, 0x71,
	0x77, 0x52, 0x39, 0x2e, 0x7a, 0xb9, 0x17, 0xdd, 0x90, 0x4b, 0xc0, 0x31, 0xa6, 0xfe, 0xd2, 0x4b,
	0x02, 0x31, 0xa7, 0x5d, 0xbc, 0x87, 0xa7, 0xb6, 0xe1, 0x3e, 0xbf, 0xbd, 0x8f, 0x96, 0x25, 0xaf,
	0x4a, 0xee, 0x82, 0x83, 0xfb, 0x25, 0x15, 0x09, 0xf9, 0x00, 0x47, 0x81, 0xd0, 0xd3, 0xcd, 0xd7,
	0x59, 0x28, 0xf2, 0x7c, 0xf4, 0xe0, 0xad, 0x8d, 0x76, 0x2f, 0xe7, 0xf4, 0xec, 0xff, 0x62, 0x1a,
	0x4f, 0xc8, 0xcf, 0x75, 0xcf, 0xb9, 0x5b, 0xf7, 0x9c, 0x5f, 0xeb, 0x9e, 0xf3, 0xd5, 0x5d, 0x6a,
	0x1d, 0xcf, 0x0e, 0xf0, 0x01, 0xbe, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xda, 0x13, 0xdc, 0xb3,
	0xce, 0x03, 0x00, 0x00,
}