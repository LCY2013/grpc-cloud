// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: prototpl/person-service/person-service.proto

#include "prototpl/person-service/person-service.pb.h"
#include "prototpl/person-service/person-service.grpc.pb.h"

#include <functional>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/impl/channel_interface.h>
#include <grpcpp/impl/client_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/rpc_service_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/sync_stream.h>
namespace grpc {
namespace person {
namespace service {
namespace v1 {

static const char* PersonService_method_names[] = {
  "/grpc.person.service.v1.PersonService/GetPerson",
};

std::unique_ptr< PersonService::Stub> PersonService::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< PersonService::Stub> stub(new PersonService::Stub(channel, options));
  return stub;
}

PersonService::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_GetPerson_(PersonService_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status PersonService::Stub::GetPerson(::grpc::ClientContext* context, const ::grpc::person::service::v1::PersonMessage& request, ::grpc::person::service::v1::PersonMessage* response) {
  return ::grpc::internal::BlockingUnaryCall< ::grpc::person::service::v1::PersonMessage, ::grpc::person::service::v1::PersonMessage, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_GetPerson_, context, request, response);
}

void PersonService::Stub::async::GetPerson(::grpc::ClientContext* context, const ::grpc::person::service::v1::PersonMessage* request, ::grpc::person::service::v1::PersonMessage* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::grpc::person::service::v1::PersonMessage, ::grpc::person::service::v1::PersonMessage, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetPerson_, context, request, response, std::move(f));
}

void PersonService::Stub::async::GetPerson(::grpc::ClientContext* context, const ::grpc::person::service::v1::PersonMessage* request, ::grpc::person::service::v1::PersonMessage* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_GetPerson_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::grpc::person::service::v1::PersonMessage>* PersonService::Stub::PrepareAsyncGetPersonRaw(::grpc::ClientContext* context, const ::grpc::person::service::v1::PersonMessage& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::grpc::person::service::v1::PersonMessage, ::grpc::person::service::v1::PersonMessage, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_GetPerson_, context, request);
}

::grpc::ClientAsyncResponseReader< ::grpc::person::service::v1::PersonMessage>* PersonService::Stub::AsyncGetPersonRaw(::grpc::ClientContext* context, const ::grpc::person::service::v1::PersonMessage& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetPersonRaw(context, request, cq);
  result->StartCall();
  return result;
}

PersonService::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      PersonService_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< PersonService::Service, ::grpc::person::service::v1::PersonMessage, ::grpc::person::service::v1::PersonMessage, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](PersonService::Service* service,
             ::grpc::ServerContext* ctx,
             const ::grpc::person::service::v1::PersonMessage* req,
             ::grpc::person::service::v1::PersonMessage* resp) {
               return service->GetPerson(ctx, req, resp);
             }, this)));
}

PersonService::Service::~Service() {
}

::grpc::Status PersonService::Service::GetPerson(::grpc::ServerContext* context, const ::grpc::person::service::v1::PersonMessage* request, ::grpc::person::service::v1::PersonMessage* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace grpc
}  // namespace person
}  // namespace service
}  // namespace v1
