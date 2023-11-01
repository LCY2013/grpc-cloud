// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gen/proto/person-service/person-service.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_gen_2fproto_2fperson_2dservice_2fperson_2dservice_2eproto_2epb_2eh
#define GOOGLE_PROTOBUF_INCLUDED_gen_2fproto_2fperson_2dservice_2fperson_2dservice_2eproto_2epb_2eh

#include <limits>
#include <string>
#include <type_traits>

#include "google/protobuf/port_def.inc"
#if PROTOBUF_VERSION < 4024000
#error "This file was generated by a newer version of protoc which is"
#error "incompatible with your Protocol Buffer headers. Please update"
#error "your headers."
#endif  // PROTOBUF_VERSION

#if 4024004 < PROTOBUF_MIN_PROTOC_VERSION
#error "This file was generated by an older version of protoc which is"
#error "incompatible with your Protocol Buffer headers. Please"
#error "regenerate this file with a newer version of protoc."
#endif  // PROTOBUF_MIN_PROTOC_VERSION
#include "google/protobuf/port_undef.inc"
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/arena.h"
#include "google/protobuf/arenastring.h"
#include "google/protobuf/generated_message_util.h"
#include "google/protobuf/metadata_lite.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/message.h"
#include "google/protobuf/repeated_field.h"  // IWYU pragma: export
#include "google/protobuf/extension_set.h"  // IWYU pragma: export
#include "google/protobuf/unknown_field_set.h"
#include "google/api/annotations.pb.h"
#include "grpc-gateway/protoc-gen-openapiv2/options/annotations.pb.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"

#define PROTOBUF_INTERNAL_EXPORT_gen_2fproto_2fperson_2dservice_2fperson_2dservice_2eproto

namespace google {
namespace protobuf {
namespace internal {
class AnyMetadata;
}  // namespace internal
}  // namespace protobuf
}  // namespace google

// Internal implementation detail -- do not use these members.
struct TableStruct_gen_2fproto_2fperson_2dservice_2fperson_2dservice_2eproto {
  static const ::uint32_t offsets[];
};
extern const ::google::protobuf::internal::DescriptorTable
    descriptor_table_gen_2fproto_2fperson_2dservice_2fperson_2dservice_2eproto;
namespace grpc {
namespace person {
namespace service {
namespace v1 {
class PersonMessage;
struct PersonMessageDefaultTypeInternal;
extern PersonMessageDefaultTypeInternal _PersonMessage_default_instance_;
}  // namespace v1
}  // namespace service
}  // namespace person
}  // namespace grpc
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google

namespace grpc {
namespace person {
namespace service {
namespace v1 {

// ===================================================================


// -------------------------------------------------------------------

class PersonMessage final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:grpc.person.service.v1.PersonMessage) */ {
 public:
  inline PersonMessage() : PersonMessage(nullptr) {}
  ~PersonMessage() override;
  template<typename = void>
  explicit PROTOBUF_CONSTEXPR PersonMessage(::google::protobuf::internal::ConstantInitialized);

  PersonMessage(const PersonMessage& from);
  PersonMessage(PersonMessage&& from) noexcept
    : PersonMessage() {
    *this = ::std::move(from);
  }

  inline PersonMessage& operator=(const PersonMessage& from) {
    CopyFrom(from);
    return *this;
  }
  inline PersonMessage& operator=(PersonMessage&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const {
    return _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance);
  }
  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields() {
    return _internal_metadata_.mutable_unknown_fields<::google::protobuf::UnknownFieldSet>();
  }

  static const ::google::protobuf::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::google::protobuf::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::google::protobuf::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const PersonMessage& default_instance() {
    return *internal_default_instance();
  }
  static inline const PersonMessage* internal_default_instance() {
    return reinterpret_cast<const PersonMessage*>(
               &_PersonMessage_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(PersonMessage& a, PersonMessage& b) {
    a.Swap(&b);
  }
  inline void Swap(PersonMessage* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::google::protobuf::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(PersonMessage* other) {
    if (other == this) return;
    ABSL_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  PersonMessage* New(::google::protobuf::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<PersonMessage>(arena);
  }
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::google::protobuf::Arena* arena);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(PersonMessage* other);

  private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() {
    return "grpc.person.service.v1.PersonMessage";
  }
  protected:
  explicit PersonMessage(::google::protobuf::Arena* arena);
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kNameFieldNumber = 1,
    kAgeFieldNumber = 2,
  };
  // string name = 1 [json_name = "name"];
  void clear_name() ;
  const std::string& name() const;
  template <typename Arg_ = const std::string&, typename... Args_>
  void set_name(Arg_&& arg, Args_... args);
  std::string* mutable_name();
  PROTOBUF_NODISCARD std::string* release_name();
  void set_allocated_name(std::string* ptr);

  private:
  const std::string& _internal_name() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_name(
      const std::string& value);
  std::string* _internal_mutable_name();

  public:
  // string age = 2 [json_name = "age"];
  void clear_age() ;
  const std::string& age() const;
  template <typename Arg_ = const std::string&, typename... Args_>
  void set_age(Arg_&& arg, Args_... args);
  std::string* mutable_age();
  PROTOBUF_NODISCARD std::string* release_age();
  void set_allocated_age(std::string* ptr);

  private:
  const std::string& _internal_age() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_age(
      const std::string& value);
  std::string* _internal_mutable_age();

  public:
  // @@protoc_insertion_point(class_scope:grpc.person.service.v1.PersonMessage)
 private:
  class _Internal;

  template <typename T> friend class ::google::protobuf::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::google::protobuf::internal::ArenaStringPtr name_;
    ::google::protobuf::internal::ArenaStringPtr age_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    PROTOBUF_TSAN_DECLARE_MEMBER
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_gen_2fproto_2fperson_2dservice_2fperson_2dservice_2eproto;
};

// ===================================================================




// ===================================================================


#ifdef __GNUC__
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// -------------------------------------------------------------------

// PersonMessage

// string name = 1 [json_name = "name"];
inline void PersonMessage::clear_name() {
  _impl_.name_.ClearToEmpty();
}
inline const std::string& PersonMessage::name() const {
  // @@protoc_insertion_point(field_get:grpc.person.service.v1.PersonMessage.name)
  return _internal_name();
}
template <typename Arg_, typename... Args_>
inline PROTOBUF_ALWAYS_INLINE void PersonMessage::set_name(Arg_&& arg,
                                                     Args_... args) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.name_.Set(static_cast<Arg_&&>(arg), args..., GetArenaForAllocation());
  // @@protoc_insertion_point(field_set:grpc.person.service.v1.PersonMessage.name)
}
inline std::string* PersonMessage::mutable_name() {
  std::string* _s = _internal_mutable_name();
  // @@protoc_insertion_point(field_mutable:grpc.person.service.v1.PersonMessage.name)
  return _s;
}
inline const std::string& PersonMessage::_internal_name() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.name_.Get();
}
inline void PersonMessage::_internal_set_name(const std::string& value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.name_.Set(value, GetArenaForAllocation());
}
inline std::string* PersonMessage::_internal_mutable_name() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  return _impl_.name_.Mutable( GetArenaForAllocation());
}
inline std::string* PersonMessage::release_name() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  // @@protoc_insertion_point(field_release:grpc.person.service.v1.PersonMessage.name)
  return _impl_.name_.Release();
}
inline void PersonMessage::set_allocated_name(std::string* value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  _impl_.name_.SetAllocated(value, GetArenaForAllocation());
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        if (_impl_.name_.IsDefault()) {
          _impl_.name_.Set("", GetArenaForAllocation());
        }
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:grpc.person.service.v1.PersonMessage.name)
}

// string age = 2 [json_name = "age"];
inline void PersonMessage::clear_age() {
  _impl_.age_.ClearToEmpty();
}
inline const std::string& PersonMessage::age() const {
  // @@protoc_insertion_point(field_get:grpc.person.service.v1.PersonMessage.age)
  return _internal_age();
}
template <typename Arg_, typename... Args_>
inline PROTOBUF_ALWAYS_INLINE void PersonMessage::set_age(Arg_&& arg,
                                                     Args_... args) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.age_.Set(static_cast<Arg_&&>(arg), args..., GetArenaForAllocation());
  // @@protoc_insertion_point(field_set:grpc.person.service.v1.PersonMessage.age)
}
inline std::string* PersonMessage::mutable_age() {
  std::string* _s = _internal_mutable_age();
  // @@protoc_insertion_point(field_mutable:grpc.person.service.v1.PersonMessage.age)
  return _s;
}
inline const std::string& PersonMessage::_internal_age() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.age_.Get();
}
inline void PersonMessage::_internal_set_age(const std::string& value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.age_.Set(value, GetArenaForAllocation());
}
inline std::string* PersonMessage::_internal_mutable_age() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  return _impl_.age_.Mutable( GetArenaForAllocation());
}
inline std::string* PersonMessage::release_age() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  // @@protoc_insertion_point(field_release:grpc.person.service.v1.PersonMessage.age)
  return _impl_.age_.Release();
}
inline void PersonMessage::set_allocated_age(std::string* value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  _impl_.age_.SetAllocated(value, GetArenaForAllocation());
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        if (_impl_.age_.IsDefault()) {
          _impl_.age_.Set("", GetArenaForAllocation());
        }
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:grpc.person.service.v1.PersonMessage.age)
}

#ifdef __GNUC__
#pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace service
}  // namespace person
}  // namespace grpc


// @@protoc_insertion_point(global_scope)

#include "google/protobuf/port_undef.inc"

#endif  // GOOGLE_PROTOBUF_INCLUDED_gen_2fproto_2fperson_2dservice_2fperson_2dservice_2eproto_2epb_2eh
