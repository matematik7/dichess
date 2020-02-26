// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bluetoothpb.proto

package bluetoothpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Settings_Language int32

const (
	Settings_ENGLISH   Settings_Language = 0
	Settings_SLOVENIAN Settings_Language = 1
)

var Settings_Language_name = map[int32]string{
	0: "ENGLISH",
	1: "SLOVENIAN",
}

var Settings_Language_value = map[string]int32{
	"ENGLISH":   0,
	"SLOVENIAN": 1,
}

func (x Settings_Language) String() string {
	return proto.EnumName(Settings_Language_name, int32(x))
}

func (Settings_Language) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{0, 0}
}

type Settings_PlayerType int32

const (
	Settings_HUMAN    Settings_PlayerType = 0
	Settings_COMPUTER Settings_PlayerType = 1
)

var Settings_PlayerType_name = map[int32]string{
	0: "HUMAN",
	1: "COMPUTER",
}

var Settings_PlayerType_value = map[string]int32{
	"HUMAN":    0,
	"COMPUTER": 1,
}

func (x Settings_PlayerType) String() string {
	return proto.EnumName(Settings_PlayerType_name, int32(x))
}

func (Settings_PlayerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{0, 1}
}

type Response_Type int32

const (
	Response_NOOP         Response_Type = 0
	Response_GAME_UPDATE  Response_Type = 1
	Response_WIFI_UPDATE  Response_Type = 2
	Response_STATE_UPDATE Response_Type = 3
)

var Response_Type_name = map[int32]string{
	0: "NOOP",
	1: "GAME_UPDATE",
	2: "WIFI_UPDATE",
	3: "STATE_UPDATE",
}

var Response_Type_value = map[string]int32{
	"NOOP":         0,
	"GAME_UPDATE":  1,
	"WIFI_UPDATE":  2,
	"STATE_UPDATE": 3,
}

func (x Response_Type) String() string {
	return proto.EnumName(Response_Type_name, int32(x))
}

func (Response_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{1, 0}
}

type Request_Type int32

const (
	Request_NOOP            Request_Type = 0
	Request_START_WIFI_SCAN Request_Type = 1
	Request_STOP_WIFI_SCAN  Request_Type = 2
	Request_CONFIGURE_WIFI  Request_Type = 3
	Request_FORGET_WIFI     Request_Type = 4
	Request_CONNECT_WIFI    Request_Type = 5
	Request_UPDATE_SETTINGS Request_Type = 6
	Request_UNDO_MOVE       Request_Type = 7
	Request_MOVE            Request_Type = 8
)

var Request_Type_name = map[int32]string{
	0: "NOOP",
	1: "START_WIFI_SCAN",
	2: "STOP_WIFI_SCAN",
	3: "CONFIGURE_WIFI",
	4: "FORGET_WIFI",
	5: "CONNECT_WIFI",
	6: "UPDATE_SETTINGS",
	7: "UNDO_MOVE",
	8: "MOVE",
}

var Request_Type_value = map[string]int32{
	"NOOP":            0,
	"START_WIFI_SCAN": 1,
	"STOP_WIFI_SCAN":  2,
	"CONFIGURE_WIFI":  3,
	"FORGET_WIFI":     4,
	"CONNECT_WIFI":    5,
	"UPDATE_SETTINGS": 6,
	"UNDO_MOVE":       7,
	"MOVE":            8,
}

func (x Request_Type) String() string {
	return proto.EnumName(Request_Type_name, int32(x))
}

func (Request_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{2, 0}
}

type Settings struct {
	Sound                bool                       `protobuf:"varint,1,opt,name=sound,proto3" json:"sound,omitempty"`
	Language             Settings_Language          `protobuf:"varint,2,opt,name=language,proto3,enum=bluetoothpb.Settings_Language" json:"language,omitempty"`
	VoiceRecognition     bool                       `protobuf:"varint,3,opt,name=voice_recognition,json=voiceRecognition,proto3" json:"voice_recognition,omitempty"`
	AutoMove             bool                       `protobuf:"varint,4,opt,name=auto_move,json=autoMove,proto3" json:"auto_move,omitempty"`
	RandomBw             bool                       `protobuf:"varint,5,opt,name=random_bw,json=randomBw,proto3" json:"random_bw,omitempty"`
	Player1              Settings_PlayerType        `protobuf:"varint,6,opt,name=player1,proto3,enum=bluetoothpb.Settings_PlayerType" json:"player1,omitempty"`
	Player2              Settings_PlayerType        `protobuf:"varint,7,opt,name=player2,proto3,enum=bluetoothpb.Settings_PlayerType" json:"player2,omitempty"`
	ComputerSettings     *Settings_ComputerSettings `protobuf:"bytes,8,opt,name=computerSettings,proto3" json:"computerSettings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{0}
}

func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetSound() bool {
	if m != nil {
		return m.Sound
	}
	return false
}

func (m *Settings) GetLanguage() Settings_Language {
	if m != nil {
		return m.Language
	}
	return Settings_ENGLISH
}

func (m *Settings) GetVoiceRecognition() bool {
	if m != nil {
		return m.VoiceRecognition
	}
	return false
}

func (m *Settings) GetAutoMove() bool {
	if m != nil {
		return m.AutoMove
	}
	return false
}

func (m *Settings) GetRandomBw() bool {
	if m != nil {
		return m.RandomBw
	}
	return false
}

func (m *Settings) GetPlayer1() Settings_PlayerType {
	if m != nil {
		return m.Player1
	}
	return Settings_HUMAN
}

func (m *Settings) GetPlayer2() Settings_PlayerType {
	if m != nil {
		return m.Player2
	}
	return Settings_HUMAN
}

func (m *Settings) GetComputerSettings() *Settings_ComputerSettings {
	if m != nil {
		return m.ComputerSettings
	}
	return nil
}

type Settings_ComputerSettings struct {
	TimeLimitMs int32 `protobuf:"varint,1,opt,name=time_limit_ms,json=timeLimitMs,proto3" json:"time_limit_ms,omitempty"`
	// from 0 to 20
	SkillLevel int32 `protobuf:"varint,2,opt,name=skill_level,json=skillLevel,proto3" json:"skill_level,omitempty"`
	// enables elo over skill level
	LimitStrength bool `protobuf:"varint,3,opt,name=limit_strength,json=limitStrength,proto3" json:"limit_strength,omitempty"`
	// from 1350 to 2850
	Elo                  int32    `protobuf:"varint,4,opt,name=elo,proto3" json:"elo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_ComputerSettings) Reset()         { *m = Settings_ComputerSettings{} }
func (m *Settings_ComputerSettings) String() string { return proto.CompactTextString(m) }
func (*Settings_ComputerSettings) ProtoMessage()    {}
func (*Settings_ComputerSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{0, 0}
}

func (m *Settings_ComputerSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_ComputerSettings.Unmarshal(m, b)
}
func (m *Settings_ComputerSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_ComputerSettings.Marshal(b, m, deterministic)
}
func (m *Settings_ComputerSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_ComputerSettings.Merge(m, src)
}
func (m *Settings_ComputerSettings) XXX_Size() int {
	return xxx_messageInfo_Settings_ComputerSettings.Size(m)
}
func (m *Settings_ComputerSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_ComputerSettings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_ComputerSettings proto.InternalMessageInfo

func (m *Settings_ComputerSettings) GetTimeLimitMs() int32 {
	if m != nil {
		return m.TimeLimitMs
	}
	return 0
}

func (m *Settings_ComputerSettings) GetSkillLevel() int32 {
	if m != nil {
		return m.SkillLevel
	}
	return 0
}

func (m *Settings_ComputerSettings) GetLimitStrength() bool {
	if m != nil {
		return m.LimitStrength
	}
	return false
}

func (m *Settings_ComputerSettings) GetElo() int32 {
	if m != nil {
		return m.Elo
	}
	return 0
}

type Response struct {
	Type           Response_Type           `protobuf:"varint,1,opt,name=type,proto3,enum=bluetoothpb.Response_Type" json:"type,omitempty"`
	Networks       []*Response_WifiNetwork `protobuf:"bytes,2,rep,name=networks,proto3" json:"networks,omitempty"`
	Settings       *Settings               `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	GameInProgress bool                    `protobuf:"varint,4,opt,name=gameInProgress,proto3" json:"gameInProgress,omitempty"`
	Moves          []string                `protobuf:"bytes,5,rep,name=moves,proto3" json:"moves,omitempty"`
	// not used for now
	WhiteTurn            bool                 `protobuf:"varint,6,opt,name=whiteTurn,proto3" json:"whiteTurn,omitempty"`
	State                string               `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	ChessBoard           *Response_ChessBoard `protobuf:"bytes,8,opt,name=chess_board,json=chessBoard,proto3" json:"chess_board,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetType() Response_Type {
	if m != nil {
		return m.Type
	}
	return Response_NOOP
}

func (m *Response) GetNetworks() []*Response_WifiNetwork {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *Response) GetSettings() *Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *Response) GetGameInProgress() bool {
	if m != nil {
		return m.GameInProgress
	}
	return false
}

func (m *Response) GetMoves() []string {
	if m != nil {
		return m.Moves
	}
	return nil
}

func (m *Response) GetWhiteTurn() bool {
	if m != nil {
		return m.WhiteTurn
	}
	return false
}

func (m *Response) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Response) GetChessBoard() *Response_ChessBoard {
	if m != nil {
		return m.ChessBoard
	}
	return nil
}

type Response_WifiNetwork struct {
	Ssid                 string   `protobuf:"bytes,1,opt,name=ssid,proto3" json:"ssid,omitempty"`
	Connected            bool     `protobuf:"varint,2,opt,name=connected,proto3" json:"connected,omitempty"`
	Available            bool     `protobuf:"varint,3,opt,name=available,proto3" json:"available,omitempty"`
	Saved                bool     `protobuf:"varint,4,opt,name=saved,proto3" json:"saved,omitempty"`
	Connecting           bool     `protobuf:"varint,5,opt,name=connecting,proto3" json:"connecting,omitempty"`
	Failed               bool     `protobuf:"varint,6,opt,name=failed,proto3" json:"failed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_WifiNetwork) Reset()         { *m = Response_WifiNetwork{} }
func (m *Response_WifiNetwork) String() string { return proto.CompactTextString(m) }
func (*Response_WifiNetwork) ProtoMessage()    {}
func (*Response_WifiNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{1, 0}
}

func (m *Response_WifiNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_WifiNetwork.Unmarshal(m, b)
}
func (m *Response_WifiNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_WifiNetwork.Marshal(b, m, deterministic)
}
func (m *Response_WifiNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_WifiNetwork.Merge(m, src)
}
func (m *Response_WifiNetwork) XXX_Size() int {
	return xxx_messageInfo_Response_WifiNetwork.Size(m)
}
func (m *Response_WifiNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_WifiNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_Response_WifiNetwork proto.InternalMessageInfo

func (m *Response_WifiNetwork) GetSsid() string {
	if m != nil {
		return m.Ssid
	}
	return ""
}

func (m *Response_WifiNetwork) GetConnected() bool {
	if m != nil {
		return m.Connected
	}
	return false
}

func (m *Response_WifiNetwork) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Response_WifiNetwork) GetSaved() bool {
	if m != nil {
		return m.Saved
	}
	return false
}

func (m *Response_WifiNetwork) GetConnecting() bool {
	if m != nil {
		return m.Connecting
	}
	return false
}

func (m *Response_WifiNetwork) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

type Response_ChessBoard struct {
	Fen                  string   `protobuf:"bytes,1,opt,name=fen,proto3" json:"fen,omitempty"`
	Rotate               bool     `protobuf:"varint,2,opt,name=rotate,proto3" json:"rotate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_ChessBoard) Reset()         { *m = Response_ChessBoard{} }
func (m *Response_ChessBoard) String() string { return proto.CompactTextString(m) }
func (*Response_ChessBoard) ProtoMessage()    {}
func (*Response_ChessBoard) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{1, 1}
}

func (m *Response_ChessBoard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_ChessBoard.Unmarshal(m, b)
}
func (m *Response_ChessBoard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_ChessBoard.Marshal(b, m, deterministic)
}
func (m *Response_ChessBoard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_ChessBoard.Merge(m, src)
}
func (m *Response_ChessBoard) XXX_Size() int {
	return xxx_messageInfo_Response_ChessBoard.Size(m)
}
func (m *Response_ChessBoard) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_ChessBoard.DiscardUnknown(m)
}

var xxx_messageInfo_Response_ChessBoard proto.InternalMessageInfo

func (m *Response_ChessBoard) GetFen() string {
	if m != nil {
		return m.Fen
	}
	return ""
}

func (m *Response_ChessBoard) GetRotate() bool {
	if m != nil {
		return m.Rotate
	}
	return false
}

type Request struct {
	Type                 Request_Type `protobuf:"varint,1,opt,name=type,proto3,enum=bluetoothpb.Request_Type" json:"type,omitempty"`
	WifiSsid             string       `protobuf:"bytes,2,opt,name=wifi_ssid,json=wifiSsid,proto3" json:"wifi_ssid,omitempty"`
	WifiPsk              string       `protobuf:"bytes,3,opt,name=wifi_psk,json=wifiPsk,proto3" json:"wifi_psk,omitempty"`
	Settings             *Settings    `protobuf:"bytes,4,opt,name=settings,proto3" json:"settings,omitempty"`
	Move                 string       `protobuf:"bytes,5,opt,name=move,proto3" json:"move,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() Request_Type {
	if m != nil {
		return m.Type
	}
	return Request_NOOP
}

func (m *Request) GetWifiSsid() string {
	if m != nil {
		return m.WifiSsid
	}
	return ""
}

func (m *Request) GetWifiPsk() string {
	if m != nil {
		return m.WifiPsk
	}
	return ""
}

func (m *Request) GetSettings() *Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *Request) GetMove() string {
	if m != nil {
		return m.Move
	}
	return ""
}

func init() {
	proto.RegisterEnum("bluetoothpb.Settings_Language", Settings_Language_name, Settings_Language_value)
	proto.RegisterEnum("bluetoothpb.Settings_PlayerType", Settings_PlayerType_name, Settings_PlayerType_value)
	proto.RegisterEnum("bluetoothpb.Response_Type", Response_Type_name, Response_Type_value)
	proto.RegisterEnum("bluetoothpb.Request_Type", Request_Type_name, Request_Type_value)
	proto.RegisterType((*Settings)(nil), "bluetoothpb.Settings")
	proto.RegisterType((*Settings_ComputerSettings)(nil), "bluetoothpb.Settings.ComputerSettings")
	proto.RegisterType((*Response)(nil), "bluetoothpb.Response")
	proto.RegisterType((*Response_WifiNetwork)(nil), "bluetoothpb.Response.WifiNetwork")
	proto.RegisterType((*Response_ChessBoard)(nil), "bluetoothpb.Response.ChessBoard")
	proto.RegisterType((*Request)(nil), "bluetoothpb.Request")
}

func init() { proto.RegisterFile("bluetoothpb.proto", fileDescriptor_6671346f1aa9b484) }

var fileDescriptor_6671346f1aa9b484 = []byte{
	// 854 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0x8e, 0xc1, 0x0e, 0xf6, 0xf1, 0x86, 0xf5, 0x4e, 0x7f, 0xe4, 0xcd, 0x56, 0x5b, 0x8a, 0xb4,
	0x11, 0x52, 0x55, 0xa4, 0xa5, 0x52, 0x2f, 0x56, 0xea, 0x05, 0x4b, 0x1c, 0x16, 0x09, 0x6c, 0x34,
	0x36, 0xbb, 0x97, 0x96, 0x81, 0x09, 0x19, 0xc5, 0x78, 0xa8, 0x67, 0x00, 0xe5, 0x21, 0x2a, 0xf5,
	0x09, 0xfa, 0x02, 0x7d, 0xa1, 0xbe, 0x46, 0xdf, 0xa0, 0x9a, 0xb1, 0xf9, 0x49, 0xc2, 0x45, 0x7b,
	0x77, 0xce, 0x37, 0xdf, 0x19, 0x9d, 0xf9, 0xce, 0x77, 0x6c, 0x78, 0x35, 0x4d, 0xd7, 0x44, 0x30,
	0x26, 0xee, 0x56, 0xd3, 0xf6, 0x2a, 0x67, 0x82, 0x21, 0xfb, 0x08, 0x6a, 0xfe, 0xa3, 0x83, 0x19,
	0x12, 0x21, 0x68, 0xb6, 0xe0, 0xe8, 0x6b, 0x30, 0x38, 0x5b, 0x67, 0x73, 0x57, 0x6b, 0x68, 0x2d,
	0x13, 0x17, 0x09, 0xfa, 0x00, 0x66, 0x9a, 0x64, 0x8b, 0x75, 0xb2, 0x20, 0x6e, 0xa5, 0xa1, 0xb5,
	0xea, 0x9d, 0xb7, 0xed, 0xe3, 0x5b, 0x77, 0xe5, 0xed, 0x61, 0xc9, 0xc2, 0x7b, 0x3e, 0xfa, 0x11,
	0x5e, 0x6d, 0x18, 0x9d, 0x91, 0x38, 0x27, 0x33, 0xb6, 0xc8, 0xa8, 0xa0, 0x2c, 0x73, 0xab, 0xea,
	0x76, 0x47, 0x1d, 0xe0, 0x03, 0x8e, 0xde, 0x80, 0x95, 0xac, 0x05, 0x8b, 0x97, 0x6c, 0x43, 0x5c,
	0x5d, 0x91, 0x4c, 0x09, 0x8c, 0xd8, 0x86, 0xc8, 0xc3, 0x3c, 0xc9, 0xe6, 0x6c, 0x19, 0x4f, 0xb7,
	0xae, 0x51, 0x1c, 0x16, 0xc0, 0xc7, 0x2d, 0xfa, 0x00, 0xb5, 0x55, 0x9a, 0x3c, 0x90, 0xfc, 0xbd,
	0x7b, 0xae, 0x3a, 0x6c, 0x9c, 0xee, 0x70, 0xac, 0x48, 0xd1, 0xc3, 0x8a, 0xe0, 0x5d, 0xc1, 0xa1,
	0xb6, 0xe3, 0xd6, 0xfe, 0x5f, 0x6d, 0x07, 0x61, 0x70, 0x66, 0x6c, 0xb9, 0x5a, 0x0b, 0x92, 0xef,
	0x78, 0xae, 0xd9, 0xd0, 0x5a, 0x76, 0xe7, 0xea, 0xf4, 0x25, 0xbd, 0x27, 0x6c, 0xfc, 0xac, 0xfe,
	0xf2, 0x0f, 0x0d, 0x9c, 0xa7, 0x34, 0xd4, 0x84, 0x0b, 0x41, 0x97, 0x24, 0x4e, 0xe9, 0x92, 0x8a,
	0x78, 0xc9, 0xd5, 0x84, 0x0c, 0x6c, 0x4b, 0x70, 0x28, 0xb1, 0x11, 0x47, 0xdf, 0x83, 0xcd, 0xef,
	0x69, 0x9a, 0xc6, 0x29, 0xd9, 0x90, 0x54, 0x8d, 0xca, 0xc0, 0xa0, 0xa0, 0xa1, 0x44, 0xd0, 0x3b,
	0xa8, 0x17, 0xf5, 0x5c, 0xe4, 0x24, 0x5b, 0x88, 0xbb, 0x72, 0x12, 0x17, 0x0a, 0x0d, 0x4b, 0x10,
	0x39, 0x50, 0x25, 0x29, 0x53, 0x03, 0x30, 0xb0, 0x0c, 0x9b, 0x57, 0x60, 0xee, 0x66, 0x8b, 0x6c,
	0xa8, 0x79, 0x7e, 0x7f, 0x38, 0x08, 0x3f, 0x39, 0x67, 0xe8, 0x02, 0xac, 0x70, 0x18, 0x7c, 0xf6,
	0xfc, 0x41, 0xd7, 0x77, 0xb4, 0xe6, 0x3b, 0x80, 0x83, 0x4a, 0xc8, 0x02, 0xe3, 0xd3, 0x64, 0xd4,
	0xf5, 0x9d, 0x33, 0xf4, 0x02, 0xcc, 0x5e, 0x30, 0x1a, 0x4f, 0x22, 0x0f, 0x3b, 0x5a, 0xf3, 0x77,
	0x03, 0x4c, 0x4c, 0xf8, 0x8a, 0x65, 0x9c, 0xa0, 0x36, 0xe8, 0xe2, 0x61, 0x45, 0xd4, 0x83, 0xea,
	0x9d, 0xcb, 0x47, 0xb2, 0xed, 0x48, 0x6d, 0xa5, 0xba, 0xe2, 0xa1, 0x5f, 0xc1, 0xcc, 0x88, 0xd8,
	0xb2, 0xfc, 0x9e, 0xbb, 0x95, 0x46, 0xb5, 0x65, 0x77, 0x7e, 0x38, 0x5d, 0xf3, 0x85, 0xde, 0x52,
	0xbf, 0x60, 0xe2, 0x7d, 0x09, 0x7a, 0x0f, 0x26, 0xdf, 0x4d, 0xaa, 0xaa, 0x26, 0xf5, 0xcd, 0xc9,
	0x49, 0xe1, 0x3d, 0x0d, 0x5d, 0x41, 0x7d, 0x91, 0x2c, 0xc9, 0x20, 0x1b, 0xe7, 0x6c, 0x91, 0x13,
	0xce, 0x4b, 0x6f, 0x3e, 0x41, 0xe5, 0xf6, 0x48, 0xe7, 0x72, 0xd7, 0x68, 0x54, 0x5b, 0x16, 0x2e,
	0x12, 0xf4, 0x1d, 0x58, 0xdb, 0x3b, 0x2a, 0x48, 0xb4, 0xce, 0x33, 0x65, 0x4e, 0x13, 0x1f, 0x00,
	0xb5, 0x71, 0x22, 0x11, 0x44, 0x59, 0xcf, 0xc2, 0x45, 0x82, 0xba, 0x60, 0xcf, 0xee, 0x08, 0xe7,
	0xf1, 0x94, 0x25, 0xf9, 0xbc, 0x74, 0x54, 0xe3, 0xf4, 0x33, 0x7b, 0x92, 0xf8, 0x51, 0xf2, 0x30,
	0xcc, 0xf6, 0xf1, 0xe5, 0x5f, 0x1a, 0xd8, 0x47, 0x0a, 0x20, 0x04, 0x3a, 0xe7, 0xb4, 0xd8, 0x6c,
	0x0b, 0xab, 0x58, 0xb6, 0x36, 0x63, 0x59, 0x46, 0x66, 0x82, 0xcc, 0x95, 0x5d, 0x4c, 0x7c, 0x00,
	0xe4, 0x69, 0xb2, 0x49, 0x68, 0x9a, 0x4c, 0x53, 0x52, 0x1a, 0xe5, 0x00, 0xa8, 0xc6, 0x93, 0x0d,
	0x99, 0x97, 0x5a, 0x14, 0x09, 0x7a, 0x0b, 0x50, 0x5e, 0x40, 0xb3, 0x45, 0xb9, 0xa5, 0x47, 0x08,
	0xfa, 0x16, 0xce, 0x6f, 0x13, 0x9a, 0x92, 0x79, 0xa9, 0x44, 0x99, 0x5d, 0xfe, 0x02, 0x70, 0x78,
	0x87, 0x34, 0xe0, 0x2d, 0xc9, 0xca, 0x56, 0x65, 0x28, 0xeb, 0x72, 0xa6, 0x74, 0x2a, 0xda, 0x2c,
	0xb3, 0xe6, 0x35, 0xe8, 0xca, 0x6a, 0x26, 0xe8, 0x7e, 0x10, 0x8c, 0x9d, 0x33, 0xf4, 0x12, 0xec,
	0x7e, 0x77, 0xe4, 0xc5, 0x93, 0xf1, 0x75, 0x37, 0xf2, 0x1c, 0x4d, 0x02, 0x5f, 0x06, 0x37, 0x83,
	0x1d, 0x50, 0x41, 0x0e, 0xbc, 0x08, 0xa3, 0x6e, 0xb4, 0xa7, 0x54, 0x9b, 0x7f, 0x57, 0xa0, 0x86,
	0xc9, 0x6f, 0x6b, 0xc2, 0x05, 0xfa, 0xe9, 0x91, 0x1d, 0x5f, 0x3f, 0xd1, 0x5c, 0x71, 0x8e, 0xdd,
	0xf8, 0x06, 0xac, 0x2d, 0xbd, 0xa5, 0xb1, 0xd2, 0xb6, 0xa2, 0x1a, 0x36, 0x25, 0x10, 0x4a, 0x7d,
	0x5f, 0x83, 0x8a, 0xe3, 0x15, 0xbf, 0x57, 0x02, 0x5a, 0xb8, 0x26, 0xf3, 0x31, 0xbf, 0x7f, 0x64,
	0x43, 0xfd, 0xbf, 0xd9, 0x10, 0x81, 0xae, 0x3e, 0x8c, 0x46, 0x31, 0x41, 0x19, 0x37, 0xff, 0xd4,
	0x9e, 0x09, 0xf0, 0x15, 0xbc, 0x0c, 0xa3, 0x2e, 0x8e, 0x62, 0xf5, 0xea, 0xb0, 0x27, 0x17, 0x13,
	0x21, 0xa8, 0x87, 0x51, 0x30, 0x3e, 0xc2, 0x2a, 0x12, 0xeb, 0x05, 0xfe, 0xcd, 0xa0, 0x3f, 0xc1,
	0x9e, 0x3a, 0x70, 0xaa, 0x52, 0xac, 0x9b, 0x00, 0xf7, 0xbd, 0xa2, 0xda, 0xd1, 0xa5, 0x58, 0xbd,
	0xc0, 0xf7, 0xbd, 0x5e, 0x89, 0x18, 0xf2, 0xfe, 0x42, 0xb8, 0x38, 0xf4, 0xa2, 0x68, 0xe0, 0xf7,
	0x43, 0xe7, 0x5c, 0x7e, 0x07, 0x26, 0xfe, 0x75, 0x10, 0x8f, 0x82, 0xcf, 0x9e, 0x53, 0x93, 0xdd,
	0xa8, 0xc8, 0x9c, 0x9e, 0xab, 0x5f, 0xce, 0xcf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xb1,
	0x4d, 0xbe, 0x87, 0x06, 0x00, 0x00,
}
