// Use proto3 syntax, use proto2 if not specified

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.2
// source: game.proto

// Generate the package path where the proto file is located

package entity

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Define the response to user information
type Game struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	GameId           string                 `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Name             string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CoverUrl         string                 `protobuf:"bytes,3,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`
	Lang             string                 `protobuf:"bytes,4,opt,name=lang,proto3" json:"lang,omitempty"`
	Intro            string                 `protobuf:"bytes,5,opt,name=intro,proto3" json:"intro,omitempty"`
	Background       string                 `protobuf:"bytes,6,opt,name=background,proto3" json:"background,omitempty"`
	GameTags         []string               `protobuf:"bytes,7,rep,name=game_tags,json=gameTags,proto3" json:"game_tags,omitempty"`
	Mechanics        string                 `protobuf:"bytes,8,opt,name=mechanics,proto3" json:"mechanics,omitempty"`
	BackgroundMusics []string               `protobuf:"bytes,9,rep,name=background_musics,json=backgroundMusics,proto3" json:"background_musics,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Game) Reset() {
	*x = Game{}
	mi := &file_game_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *Game) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Game) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Game) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Game) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *Game) GetBackground() string {
	if x != nil {
		return x.Background
	}
	return ""
}

func (x *Game) GetGameTags() []string {
	if x != nil {
		return x.GameTags
	}
	return nil
}

func (x *Game) GetMechanics() string {
	if x != nil {
		return x.Mechanics
	}
	return ""
}

func (x *Game) GetBackgroundMusics() []string {
	if x != nil {
		return x.BackgroundMusics
	}
	return nil
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x22, 0x82, 0x02, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74,
	0x72, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x62, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData []byte
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_game_proto_rawDesc), len(file_game_proto_rawDesc)))
	})
	return file_game_proto_rawDescData
}

var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_game_proto_goTypes = []any{
	(*Game)(nil), // 0: game.Game
}
var file_game_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_game_proto_rawDesc), len(file_game_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
