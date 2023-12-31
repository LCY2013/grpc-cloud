// Code generated by protoc-gen-validate
// source: prototpl/echo-service/echo-service.proto
// DO NOT EDIT!!!

#include "prototpl/echo-service/echo-service.pb.validate.h"

#include <google/protobuf/message.h>
#include <google/protobuf/util/time_util.h>
#include "re2/re2.h"

namespace pgv {

namespace protobuf = google::protobuf;
namespace protobuf_wkt = google::protobuf;

namespace validate {
using std::string;

// define the regex for a UUID once up-front
const re2::RE2 _uuidPattern("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$");

pgv::Validator<::grpc::echo::service::v1::StringMessage> validator___grpc__echo__service__v1__StringMessage(static_cast<bool(*)(const ::grpc::echo::service::v1::StringMessage&, pgv::ValidationMsg*)>(::grpc::echo::service::v1::Validate));


} // namespace validate
} // namespace pgv


namespace grpc {
namespace echo {
namespace service {
namespace v1 {


// Validate checks the field values on ::grpc::echo::service::v1::StringMessage
// with the rules defined in the proto definition for this message. If any
// rules are violated, the return value is false and an error message is
// written to the input string argument.

	

	

	

	

        

	

	

	



bool Validate(const ::grpc::echo::service::v1::StringMessage& m, pgv::ValidationMsg* err) {
	(void)m;
	(void)err;
	
	
	
	

	

	
		
			if (pgv::Utf8Len(m.value()) < 5) {
				{
std::ostringstream msg("invalid ");
msg << "StringMessageValidationError" << "." << "Value";
msg << ": " << "value length must be at least 5 characters";
*err = msg.str();
return false;
}
			}
		
	

	

	

	

	

	

        

	
	

		
	return true;
}


} // namespace
} // namespace
} // namespace
} // namespace

