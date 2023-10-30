// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gen/go/echo-service/echo-service.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_gen_2fgo_2fecho_2dservice_2fecho_2dservice_2eproto_2epb_2eh
#define GOOGLE_PROTOBUF_INCLUDED_gen_2fgo_2fecho_2dservice_2fecho_2dservice_2eproto_2epb_2eh

#include <limits>
#include <string>
#include <type_traits>

#include "google/protobuf/port_def.inc"
#if PROTOBUF_VERSION < 4024000
#error "This file was generated by a newer version of protoc which is"
#error "incompatible with your Protocol Buffer headers. Please update"
#error "your headers."
#endif  // PROTOBUF_VERSION

#if 4024003 < PROTOBUF_MIN_PROTOC_VERSION
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

#define PROTOBUF_INTERNAL_EXPORT_gen_2fgo_2fecho_2dservice_2fecho_2dservice_2eproto

namespace google {
namespace protobuf {
namespace internal {
class AnyMetadata;
}  // namespace internal
}  // namespace protobuf
}  // namespace google

// Internal implementation detail -- do not use these members.
struct TableStruct_gen_2fgo_2fecho_2dservice_2fecho_2dservice_2eproto {
  static const ::uint32_t offsets[];
};
extern const ::google::protobuf::internal::DescriptorTable
    descriptor_table_gen_2fgo_2fecho_2dservice_2fecho_2dservice_2eproto;
namespace grpc {
namespace gateway {
namespace service {
namespace v1 {
class StringMessage;
struct StringMessageDefaultTypeInternal;
extern StringMessageDefaultTypeInternal _StringMessage_default_instance_;
}  // namespace v1
}  // namespace service
}  // namespace gateway
}  // namespace grpc
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google

namespace grpc {
namespace gateway {
namespace service {
namespace v1 {

// ===================================================================


// -------------------------------------------------------------------

class StringMessage final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:grpc.gateway.service.v1.StringMessage) */ {
 public:
  inline StringMessage() : StringMessage(nullptr) {}
  ~StringMessage() override;
  template<typename = void>
  explicit PROTOBUF_CONSTEXPR StringMessage(::google::protobuf::internal::ConstantInitialized);

  StringMessage(const StringMessage& from);
  StringMessage(StringMessage&& from) noexcept
    : StringMessage() {
    *this = ::std::move(from);
  }

  inline StringMessage& operator=(const StringMessage& from) {
    CopyFrom(from);
    return *this;
  }
  inline StringMessage& operator=(StringMessage&& from) noexcept {
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
  static const StringMessage& default_instance() {
    return *internal_default_instance();
  }
  static inline const StringMessage* internal_default_instance() {
    return reinterpret_cast<const StringMessage*>(
               &_StringMessage_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(StringMessage& a, StringMessage& b) {
    a.Swap(&b);
  }
  inline void Swap(StringMessage* other) {
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
  void UnsafeArenaSwap(StringMessage* other) {
    if (other == this) return;
    ABSL_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  StringMessage* New(::google::protobuf::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<StringMessage>(arena);
  }
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::google::protobuf::Arena* arena);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(StringMessage* other);

  private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() {
    return "grpc.gateway.service.v1.StringMessage";
  }
  protected:
  explicit StringMessage(::google::protobuf::Arena* arena);
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kValueFieldNumber = 1,
  };
  // string value = 1 [json_name = "value"];
  void clear_value() ;
  const std::string& value() const;
  template <typename Arg_ = const std::string&, typename... Args_>
  void set_value(Arg_&& arg, Args_... args);
  std::string* mutable_value();
  PROTOBUF_NODISCARD std::string* release_value();
  void set_allocated_value(std::string* ptr);

  private:
  const std::string& _internal_value() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_value(
      const std::string& value);
  std::string* _internal_mutable_value();

  public:
  // @@protoc_insertion_point(class_scope:grpc.gateway.service.v1.StringMessage)
 private:
  class _Internal;

  template <typename T> friend class ::google::protobuf::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::google::protobuf::internal::ArenaStringPtr value_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    PROTOBUF_TSAN_DECLARE_MEMBER
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_gen_2fgo_2fecho_2dservice_2fecho_2dservice_2eproto;
};

// ===================================================================




// ===================================================================


#ifdef __GNUC__
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// -------------------------------------------------------------------

// StringMessage

// string value = 1 [json_name = "value"];
inline void StringMessage::clear_value() {
  _impl_.value_.ClearToEmpty();
}
inline const std::string& StringMessage::value() const {
  // @@protoc_insertion_point(field_get:grpc.gateway.service.v1.StringMessage.value)
  return _internal_value();
}
template <typename Arg_, typename... Args_>
inline PROTOBUF_ALWAYS_INLINE void StringMessage::set_value(Arg_&& arg,
                                                     Args_... args) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.value_.Set(static_cast<Arg_&&>(arg), args..., GetArenaForAllocation());
  // @@protoc_insertion_point(field_set:grpc.gateway.service.v1.StringMessage.value)
}
inline std::string* StringMessage::mutable_value() {
  std::string* _s = _internal_mutable_value();
  // @@protoc_insertion_point(field_mutable:grpc.gateway.service.v1.StringMessage.value)
  return _s;
}
inline const std::string& StringMessage::_internal_value() const {
  PROTOBUF_TSAN_READ(&_impl_._tsan_detect_race);
  return _impl_.value_.Get();
}
inline void StringMessage::_internal_set_value(const std::string& value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  _impl_.value_.Set(value, GetArenaForAllocation());
}
inline std::string* StringMessage::_internal_mutable_value() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  ;
  return _impl_.value_.Mutable( GetArenaForAllocation());
}
inline std::string* StringMessage::release_value() {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  // @@protoc_insertion_point(field_release:grpc.gateway.service.v1.StringMessage.value)
  return _impl_.value_.Release();
}
inline void StringMessage::set_allocated_value(std::string* value) {
  PROTOBUF_TSAN_WRITE(&_impl_._tsan_detect_race);
  _impl_.value_.SetAllocated(value, GetArenaForAllocation());
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        if (_impl_.value_.IsDefault()) {
          _impl_.value_.Set("", GetArenaForAllocation());
        }
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:grpc.gateway.service.v1.StringMessage.value)
}

#ifdef __GNUC__
#pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace service
}  // namespace gateway
}  // namespace grpc


// @@protoc_insertion_point(global_scope)

#include "google/protobuf/port_undef.inc"

#endif  // GOOGLE_PROTOBUF_INCLUDED_gen_2fgo_2fecho_2dservice_2fecho_2dservice_2eproto_2epb_2eh
