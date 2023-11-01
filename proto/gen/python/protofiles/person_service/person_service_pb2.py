# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protofiles/person-service/person-service.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from grpc_gateway.protoc_gen_openapiv2.options import annotations_pb2 as grpc__gateway_dot_protoc__gen__openapiv2_dot_options_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.protofiles/person-service/person-service.proto\x12\x16grpc.person.service.v1\x1a\x1cgoogle/api/annotations.proto\x1a;grpc-gateway/protoc-gen-openapiv2/options/annotations.proto\"5\n\rPersonMessage\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x10\n\x03\x61ge\x18\x02 \x01(\tR\x03\x61ge2\xd2\x01\n\rPersonService\x12\xc0\x01\n\tGetPerson\x12%.grpc.person.service.v1.PersonMessage\x1a%.grpc.person.service.v1.PersonMessage\"e\x92\x41\x41\n\x0eperson_service\x12\x11get a person info\x1a\x04\x64\x65sc*\tgetPersonJ\x0b\n\x03\x32\x30\x30\x12\x04\n\x02OK\x82\xd3\xe4\x93\x02\x1b\"\x16/person/service/v1/get:\x01*B\xbd\x03\n\x19io.grpc.person.service.v1B\x12PersonServiceProtoH\x02Z\x16grpc/person/service/v1\xa2\x02\x03GPS\xaa\x02\x16Grpc.Person.Service.V1\xca\x02\x16Grpc\\Person\\Service\\V1\xe2\x02\"Grpc\\Person\\Service\\V1\\GPBMetadata\xea\x02\x19Grpc::Person::Service::V1\x92\x41\xf7\x01\x12\xbb\x01\n\nPerson API\"Q\n\x16person_service project\x12%https://github.com/LCY2013/grpc-cloud\x1a\x10none@example.com*U\n\x1a\x41pache License Version 2.0\x12\x37https://github.com/LCY2013/grpc-cloud/blob/main/LICENSE2\x03\x31.0*\x01\x02\x32\x10\x61pplication/json:\x10\x61pplication/jsonj\x10\n\x0eperson_serviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'protofiles.person_service.person_service_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\031io.grpc.person.service.v1B\022PersonServiceProtoH\002Z\026grpc/person/service/v1\242\002\003GPS\252\002\026Grpc.Person.Service.V1\312\002\026Grpc\\Person\\Service\\V1\342\002\"Grpc\\Person\\Service\\V1\\GPBMetadata\352\002\031Grpc::Person::Service::V1\222A\367\001\022\273\001\n\nPerson API\"Q\n\026person_service project\022%https://github.com/LCY2013/grpc-cloud\032\020none@example.com*U\n\032Apache License Version 2.0\0227https://github.com/LCY2013/grpc-cloud/blob/main/LICENSE2\0031.0*\001\0022\020application/json:\020application/jsonj\020\n\016person_service'
  _globals['_PERSONSERVICE'].methods_by_name['GetPerson']._options = None
  _globals['_PERSONSERVICE'].methods_by_name['GetPerson']._serialized_options = b'\222AA\n\016person_service\022\021get a person info\032\004desc*\tgetPersonJ\013\n\003200\022\004\n\002OK\202\323\344\223\002\033\"\026/person/service/v1/get:\001*'
  _globals['_PERSONMESSAGE']._serialized_start=165
  _globals['_PERSONMESSAGE']._serialized_end=218
  _globals['_PERSONSERVICE']._serialized_start=221
  _globals['_PERSONSERVICE']._serialized_end=431
# @@protoc_insertion_point(module_scope)
