// Code generated by protoc-gen-validate
// source: prototpl/echo-service/echo-service.proto
// DO NOT EDIT!!!

#pragma once

#include <algorithm>
#include <string>
#include <sstream>
#include <unordered_set>
#include <vector>

#include "validate/validate.h"
#include "prototpl/echo-service/echo-service.pb.h"


namespace grpc {
namespace echo {
namespace service {
namespace v1 {

using std::string;


extern bool Validate(const ::grpc::echo::service::v1::StringMessage& m, pgv::ValidationMsg* err);


} // namespace
} // namespace
} // namespace
} // namespace


#define X_GRPC_ECHO_SERVICE_V1_ECHO_SERVICE(X) \
X(::grpc::echo::service::v1::StringMessage) \

