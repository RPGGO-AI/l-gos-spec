# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: character.proto
# Protobuf Python Version: 4.25.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import module_pb2 as module__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0f\x63haracter.proto\x12\tcharacter\x1a\x0cmodule.proto\"\xc3\x03\n\tCharacter\x12\"\n\x0c\x63haracter_id\x18\x01 \x01(\tR\x0c\x63haracter_id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x1e\n\navatar_url\x18\x03 \x01(\tR\navatar_url\x12\x1e\n\nbackground\x18\x04 \x01(\tR\nbackground\x12\x1e\n\nappearance\x18\x05 \x01(\tR\nappearance\x12\x16\n\x06traits\x18\x06 \x03(\tR\x06traits\x12\x12\n\x04tone\x18\x07 \x03(\tR\x04tone\x12\x14\n\x05intro\x18\x08 \x01(\tR\x05intro\x12&\n\x0e\x63haracter_tags\x18\t \x03(\tR\x0e\x63haracter_tags\x12\"\n\x0copening_line\x18\n \x01(\tR\x0copening_line\x12\x18\n\x07\x63reator\x18\x0b \x01(\tR\x07\x63reator\x12$\n\rcreator_notes\x18\x0c \x01(\tR\rcreator_notes\x12&\n\x0epublish_status\x18\r \x01(\tR\x0epublish_status\x12(\n\x07modules\x18\x0e \x03(\x0b\x32\x0e.module.ModuleR\x07modulesb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'character_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_CHARACTER']._serialized_start=45
  _globals['_CHARACTER']._serialized_end=496
# @@protoc_insertion_point(module_scope)
