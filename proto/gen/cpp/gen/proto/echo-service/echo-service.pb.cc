// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gen/proto/echo-service/echo-service.proto

#include "gen/proto/echo-service/echo-service.pb.h"

#include <algorithm>
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/extension_set.h"
#include "google/protobuf/wire_format_lite.h"
#include "google/protobuf/descriptor.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/reflection_ops.h"
#include "google/protobuf/wire_format.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"
PROTOBUF_PRAGMA_INIT_SEG
namespace _pb = ::google::protobuf;
namespace _pbi = ::google::protobuf::internal;
namespace grpc {
namespace person {
namespace service {
namespace v1 {
        template <typename>
PROTOBUF_CONSTEXPR StringMessage::StringMessage(::_pbi::ConstantInitialized)
    : _impl_{
      /*decltype(_impl_.value_)*/ {
          &::_pbi::fixed_address_empty_string,
          ::_pbi::ConstantInitialized{},
      },
      /*decltype(_impl_._cached_size_)*/ {},
    } {}
struct StringMessageDefaultTypeInternal {
  PROTOBUF_CONSTEXPR StringMessageDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~StringMessageDefaultTypeInternal() {}
  union {
    StringMessage _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 StringMessageDefaultTypeInternal _StringMessage_default_instance_;
}  // namespace v1
}  // namespace service
}  // namespace person
}  // namespace grpc
static ::_pb::Metadata file_level_metadata_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto[1];
static constexpr const ::_pb::EnumDescriptor**
    file_level_enum_descriptors_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto = nullptr;
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto = nullptr;
const ::uint32_t TableStruct_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(
    protodesc_cold) = {
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::grpc::person::service::v1::StringMessage, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::grpc::person::service::v1::StringMessage, _impl_.value_),
};

static const ::_pbi::MigrationSchema
    schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
        {0, -1, -1, sizeof(::grpc::person::service::v1::StringMessage)},
};

static const ::_pb::Message* const file_default_instances[] = {
    &::grpc::person::service::v1::_StringMessage_default_instance_._instance,
};
const char descriptor_table_protodef_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n)gen/proto/echo-service/echo-service.pr"
    "oto\022\026grpc.person.service.v1\032\034google/api/"
    "annotations.proto\032;grpc-gateway/protoc-g"
    "en-openapiv2/options/annotations.proto\"%"
    "\n\rStringMessage\022\024\n\005value\030\001 \001(\tR\005value2\236\003"
    "\n\013EchoService\022\276\001\n\004Echo\022%.grpc.person.ser"
    "vice.v1.StringMessage\032%.grpc.person.serv"
    "ice.v1.StringMessage\"h\222A=\n\014echo-service\022"
    "\016echo a message\032\004desc*\ngetMessageJ\013\n\003200"
    "\022\004\n\002OK\202\323\344\223\002\"\"\035/echo/service/v1/example/e"
    "cho:\001*\022\315\001\n\014EchoCustomer\022%.grpc.person.se"
    "rvice.v1.StringMessage\032%.grpc.person.ser"
    "vice.v1.StringMessage\"o\222A=\n\014echo-service"
    "\022\016echo a message\032\004desc*\ngetMessageJ\013\n\00320"
    "0\022\004\n\002OK\202\323\344\223\002):\001*B$\n\001*\022\037/echo/service/v1/"
    "example/echo/cB\263\003\n\031io.grpc.person.servic"
    "e.v1B\020EchoServiceProtoH\002Z\024grpc/echo/serv"
    "ice/v1\242\002\003GPS\252\002\026Grpc.Person.Service.V1\312\002\026"
    "Grpc\\Person\\Service\\V1\342\002\"Grpc\\Person\\Ser"
    "vice\\V1\\GPBMetadata\352\002\031Grpc::Person::Serv"
    "ice::V1\222A\361\001\022\267\001\n\010Echo API\"O\n\024echo_service"
    " project\022%https://github.com/LCY2013/grp"
    "c-cloud\032\020none@example.com*U\n\032Apache Lice"
    "nse Version 2.0\0227https://github.com/LCY2"
    "013/grpc-cloud/blob/main/LICENSE2\0031.0*\001\002"
    "2\020application/json:\020application/jsonj\016\n\014"
    "echo-serviceb\006proto3"
};
static const ::_pbi::DescriptorTable* const descriptor_table_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto_deps[2] =
    {
        &::descriptor_table_google_2fapi_2fannotations_2eproto,
        &::descriptor_table_grpc_2dgateway_2fprotoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto,
};
static ::absl::once_flag descriptor_table_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto = {
    false,
    false,
    1060,
    descriptor_table_protodef_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto,
    "gen/proto/echo-service/echo-service.proto",
    &descriptor_table_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto_once,
    descriptor_table_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto_deps,
    2,
    1,
    schemas,
    file_default_instances,
    TableStruct_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto::offsets,
    file_level_metadata_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto,
    file_level_enum_descriptors_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto,
    file_level_service_descriptors_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto,
};

// This function exists to be marked as weak.
// It can significantly speed up compilation by breaking up LLVM's SCC
// in the .pb.cc translation units. Large translation units see a
// reduction of more than 35% of walltime for optimized builds. Without
// the weak attribute all the messages in the file, including all the
// vtables and everything they use become part of the same SCC through
// a cycle like:
// GetMetadata -> descriptor table -> default instances ->
//   vtables -> GetMetadata
// By adding a weak function here we break the connection from the
// individual vtables back into the descriptor table.
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto_getter() {
  return &descriptor_table_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto(&descriptor_table_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto);
namespace grpc {
namespace person {
namespace service {
namespace v1 {
// ===================================================================

class StringMessage::_Internal {
 public:
};

StringMessage::StringMessage(::google::protobuf::Arena* arena)
    : ::google::protobuf::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:grpc.person.service.v1.StringMessage)
}
StringMessage::StringMessage(const StringMessage& from) : ::google::protobuf::Message() {
  StringMessage* const _this = this;
  (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.value_){},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);
  _impl_.value_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.value_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_value().empty()) {
    _this->_impl_.value_.Set(from._internal_value(), _this->GetArenaForAllocation());
  }

  // @@protoc_insertion_point(copy_constructor:grpc.person.service.v1.StringMessage)
}
inline void StringMessage::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.value_){},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _impl_.value_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.value_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}
StringMessage::~StringMessage() {
  // @@protoc_insertion_point(destructor:grpc.person.service.v1.StringMessage)
  _internal_metadata_.Delete<::google::protobuf::UnknownFieldSet>();
  SharedDtor();
}
inline void StringMessage::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.value_.Destroy();
}
void StringMessage::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void StringMessage::InternalSwap(StringMessage* other) {
  using std::swap;
  GetReflection()->Swap(this, other);}

::google::protobuf::Metadata StringMessage::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto_getter, &descriptor_table_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto_once,
      file_level_metadata_gen_2fproto_2fecho_2dservice_2fecho_2dservice_2eproto[0]);
}
// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace service
}  // namespace person
}  // namespace grpc
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google
// @@protoc_insertion_point(global_scope)
#include "google/protobuf/port_undef.inc"
