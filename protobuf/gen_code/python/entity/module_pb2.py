# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: module.proto
# Protobuf Python Version: 4.25.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0cmodule.proto\x12\x06module\"\xcd\x01\n\x06Module\x12\x1c\n\tmodule_id\x18\x01 \x01(\tR\tmodule_id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x1c\n\tcover_url\x18\x03 \x01(\tR\tcover_url\x12\x1a\n\x08\x63\x61tegory\x18\x04 \x01(\tR\x08\x63\x61tegory\x12\x18\n\x07\x63reator\x18\x05 \x01(\tR\x07\x63reator\x12\x14\n\x05intro\x18\x06 \x01(\tR\x05intro\x12\'\n\x07\x65ntries\x18\x07 \x03(\x0b\x32\r.module.EntryR\x07\x65ntries\"Q\n\x05\x45ntry\x12\x1a\n\x08\x65ntry_id\x18\x01 \x01(\tR\x08\x65ntry_id\x12\x12\n\x04keys\x18\x02 \x03(\tR\x04keys\x12\x18\n\x07\x63ontent\x18\x03 \x01(\tR\x07\x63ontentb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'module_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_MODULE']._serialized_start=25
  _globals['_MODULE']._serialized_end=230
  _globals['_ENTRY']._serialized_start=232
  _globals['_ENTRY']._serialized_end=313
# @@protoc_insertion_point(module_scope)
