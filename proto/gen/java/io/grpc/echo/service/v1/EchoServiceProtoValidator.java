// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: prototpl/echo-service/echo-service.proto

package io.grpc.echo.service.v1;


@SuppressWarnings("all")
public class EchoServiceProtoValidator {
	public static io.envoyproxy.pgv.ValidatorImpl validatorFor(Class clazz) {
		
		if (clazz.equals(io.grpc.echo.service.v1.EchoServiceProto.StringMessage.class)) return new StringMessageValidator();
		return null;
	}


/**
	 * Validates {@code StringMessage} protobuf objects.
	 */
	public static class StringMessageValidator implements io.envoyproxy.pgv.ValidatorImpl<io.grpc.echo.service.v1.EchoServiceProto.StringMessage> {
		
	
	
	

	public void assertValid(io.grpc.echo.service.v1.EchoServiceProto.StringMessage proto, io.envoyproxy.pgv.ValidatorIndex index) throws io.envoyproxy.pgv.ValidationException {
	
			io.envoyproxy.pgv.StringValidation.minLength(".grpc.echo.service.v1.StringMessage.value", proto.getValue(), 5);
	
	
	}
}
}

