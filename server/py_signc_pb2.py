# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: py_signc.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0epy_signc.proto\x12\x02pb\"\x9f\x01\n\x0bSignRequest\x12\x17\n\nsignmethod\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x12\n\x05keyid\x18\x02 \x01(\tH\x01\x88\x01\x01\x12\x15\n\x08inputUrl\x18\x03 \x01(\tH\x02\x88\x01\x01\x12\x17\n\noutputPath\x18\x04 \x01(\tH\x03\x88\x01\x01\x42\r\n\x0b_signmethodB\x08\n\x06_keyidB\x0b\n\t_inputUrlB\r\n\x0b_outputPath\"T\n\x0cSignResponse\x12\x13\n\x06status\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x16\n\toutputUrl\x18\x02 \x01(\tH\x01\x88\x01\x01\x42\t\n\x07_statusB\x0c\n\n_outputUrl\",\n\rUploadRequest\x12\x0c\n\x04mime\x18\x01 \x01(\t\x12\r\n\x05\x63hunk\x18\x02 \x01(\x0c\"\x1e\n\x0eUploadResponse\x12\x0c\n\x04name\x18\x01 \x01(\t2\x81\x01\n\x0eSigningRequest\x12\x32\n\x0bSendRequest\x12\x0f.pb.SignRequest\x1a\x10.pb.SignResponse\"\x00\x12;\n\x0eUploadToServer\x12\x11.pb.UploadRequest\x1a\x12.pb.UploadResponse\"\x00(\x01\x62\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'py_signc_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _SIGNREQUEST._serialized_start=23
  _SIGNREQUEST._serialized_end=182
  _SIGNRESPONSE._serialized_start=184
  _SIGNRESPONSE._serialized_end=268
  _UPLOADREQUEST._serialized_start=270
  _UPLOADREQUEST._serialized_end=314
  _UPLOADRESPONSE._serialized_start=316
  _UPLOADRESPONSE._serialized_end=346
  _SIGNINGREQUEST._serialized_start=349
  _SIGNINGREQUEST._serialized_end=478
# @@protoc_insertion_point(module_scope)
