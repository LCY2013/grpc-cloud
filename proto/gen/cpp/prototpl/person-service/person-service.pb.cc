// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: prototpl/person-service/person-service.proto

#include "prototpl/person-service/person-service.pb.h"

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
PROTOBUF_CONSTEXPR PersonMessage::PersonMessage(::_pbi::ConstantInitialized)
    : _impl_{
      /*decltype(_impl_.name_)*/ {
          &::_pbi::fixed_address_empty_string,
          ::_pbi::ConstantInitialized{},
      },
      /*decltype(_impl_.age_)*/ {
          &::_pbi::fixed_address_empty_string,
          ::_pbi::ConstantInitialized{},
      },
      /*decltype(_impl_._cached_size_)*/ {},
    } {}
struct PersonMessageDefaultTypeInternal {
  PROTOBUF_CONSTEXPR PersonMessageDefaultTypeInternal() : _instance(::_pbi::ConstantInitialized{}) {}
  ~PersonMessageDefaultTypeInternal() {}
  union {
    PersonMessage _instance;
  };
};

PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT
    PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 PersonMessageDefaultTypeInternal _PersonMessage_default_instance_;
}  // namespace v1
}  // namespace service
}  // namespace person
}  // namespace grpc
static ::_pb::Metadata file_level_metadata_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto[1];
static constexpr const ::_pb::EnumDescriptor**
    file_level_enum_descriptors_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto = nullptr;
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto = nullptr;
const ::uint32_t TableStruct_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(
    protodesc_cold) = {
    ~0u,  // no _has_bits_
    PROTOBUF_FIELD_OFFSET(::grpc::person::service::v1::PersonMessage, _internal_metadata_),
    ~0u,  // no _extensions_
    ~0u,  // no _oneof_case_
    ~0u,  // no _weak_field_map_
    ~0u,  // no _inlined_string_donated_
    ~0u,  // no _split_
    ~0u,  // no sizeof(Split)
    PROTOBUF_FIELD_OFFSET(::grpc::person::service::v1::PersonMessage, _impl_.name_),
    PROTOBUF_FIELD_OFFSET(::grpc::person::service::v1::PersonMessage, _impl_.age_),
};

static const ::_pbi::MigrationSchema
    schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
        {0, -1, -1, sizeof(::grpc::person::service::v1::PersonMessage)},
};

static const ::_pb::Message* const file_default_instances[] = {
    &::grpc::person::service::v1::_PersonMessage_default_instance_._instance,
};
const char descriptor_table_protodef_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n,prototpl/person-service/person-service"
    ".proto\022\026grpc.person.service.v1\032\034google/a"
    "pi/annotations.proto\032;grpc-gateway/proto"
    "c-gen-openapiv2/options/annotations.prot"
    "o\"5\n\rPersonMessage\022\022\n\004name\030\001 \001(\tR\004name\022\020"
    "\n\003age\030\002 \001(\tR\003age2\322\001\n\rPersonService\022\300\001\n\tG"
    "etPerson\022%.grpc.person.service.v1.Person"
    "Message\032%.grpc.person.service.v1.PersonM"
    "essage\"e\222AA\n\016person_service\022\021get a perso"
    "n info\032\004desc*\tgetPersonJ\013\n\003200\022\004\n\002OK\202\323\344\223"
    "\002\033\"\026/person/service/v1/get:\001*B\275\003\n\031io.grp"
    "c.person.service.v1B\022PersonServiceProtoH"
    "\002Z\026grpc/person/service/v1\242\002\003GPS\252\002\026Grpc.P"
    "erson.Service.V1\312\002\026Grpc\\Person\\Service\\V"
    "1\342\002\"Grpc\\Person\\Service\\V1\\GPBMetadata\352\002"
    "\031Grpc::Person::Service::V1\222A\367\001\022\273\001\n\nPerso"
    "n API\"Q\n\026person_service project\022%https:/"
    "/github.com/LCY2013/grpc-cloud\032\020none@exa"
    "mple.com*U\n\032Apache License Version 2.0\0227"
    "https://github.com/LCY2013/grpc-cloud/bl"
    "ob/main/LICENSE2\0031.0*\001\0022\020application/jso"
    "n:\020application/jsonj\020\n\016person_serviceb\006p"
    "roto3"
};
static const ::_pbi::DescriptorTable* const descriptor_table_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto_deps[2] =
    {
        &::descriptor_table_google_2fapi_2fannotations_2eproto,
        &::descriptor_table_grpc_2dgateway_2fprotoc_2dgen_2dopenapiv2_2foptions_2fannotations_2eproto,
};
static ::absl::once_flag descriptor_table_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto = {
    false,
    false,
    885,
    descriptor_table_protodef_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto,
    "prototpl/person-service/person-service.proto",
    &descriptor_table_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto_once,
    descriptor_table_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto_deps,
    2,
    1,
    schemas,
    file_default_instances,
    TableStruct_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto::offsets,
    file_level_metadata_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto,
    file_level_enum_descriptors_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto,
    file_level_service_descriptors_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto,
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
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto_getter() {
  return &descriptor_table_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto(&descriptor_table_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto);
namespace grpc {
namespace person {
namespace service {
namespace v1 {
// ===================================================================

class PersonMessage::_Internal {
 public:
};

PersonMessage::PersonMessage(::google::protobuf::Arena* arena)
    : ::google::protobuf::Message(arena) {
  SharedCtor(arena);
  // @@protoc_insertion_point(arena_constructor:grpc.person.service.v1.PersonMessage)
}
PersonMessage::PersonMessage(const PersonMessage& from) : ::google::protobuf::Message() {
  PersonMessage* const _this = this;
  (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.name_){},
      decltype(_impl_.age_){},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _internal_metadata_.MergeFrom<::google::protobuf::UnknownFieldSet>(
      from._internal_metadata_);
  _impl_.name_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.name_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_name().empty()) {
    _this->_impl_.name_.Set(from._internal_name(), _this->GetArenaForAllocation());
  }
  _impl_.age_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.age_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_age().empty()) {
    _this->_impl_.age_.Set(from._internal_age(), _this->GetArenaForAllocation());
  }

  // @@protoc_insertion_point(copy_constructor:grpc.person.service.v1.PersonMessage)
}
inline void PersonMessage::SharedCtor(::_pb::Arena* arena) {
  (void)arena;
  new (&_impl_) Impl_{
      decltype(_impl_.name_){},
      decltype(_impl_.age_){},
      /*decltype(_impl_._cached_size_)*/ {},
  };
  _impl_.name_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.name_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  _impl_.age_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        _impl_.age_.Set("", GetArenaForAllocation());
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}
PersonMessage::~PersonMessage() {
  // @@protoc_insertion_point(destructor:grpc.person.service.v1.PersonMessage)
  _internal_metadata_.Delete<::google::protobuf::UnknownFieldSet>();
  SharedDtor();
}
inline void PersonMessage::SharedDtor() {
  ABSL_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.name_.Destroy();
  _impl_.age_.Destroy();
}
void PersonMessage::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void PersonMessage::InternalSwap(PersonMessage* other) {
  using std::swap;
  GetReflection()->Swap(this, other);}

::google::protobuf::Metadata PersonMessage::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto_getter, &descriptor_table_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto_once,
      file_level_metadata_prototpl_2fperson_2dservice_2fperson_2dservice_2eproto[0]);
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