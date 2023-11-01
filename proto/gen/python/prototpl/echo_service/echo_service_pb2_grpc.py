# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from prototpl.echo_service import echo_service_pb2 as prototpl_dot_echo__service_dot_echo__service__pb2


class EchoServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Echo = channel.unary_unary(
                '/grpc.echo.service.v1.EchoService/Echo',
                request_serializer=prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.SerializeToString,
                response_deserializer=prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.FromString,
                )
        self.EchoCustomer = channel.unary_unary(
                '/grpc.echo.service.v1.EchoService/EchoCustomer',
                request_serializer=prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.SerializeToString,
                response_deserializer=prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.FromString,
                )


class EchoServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Echo(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def EchoCustomer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_EchoServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Echo': grpc.unary_unary_rpc_method_handler(
                    servicer.Echo,
                    request_deserializer=prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.FromString,
                    response_serializer=prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.SerializeToString,
            ),
            'EchoCustomer': grpc.unary_unary_rpc_method_handler(
                    servicer.EchoCustomer,
                    request_deserializer=prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.FromString,
                    response_serializer=prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'grpc.echo.service.v1.EchoService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class EchoService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Echo(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.echo.service.v1.EchoService/Echo',
            prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.SerializeToString,
            prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def EchoCustomer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.echo.service.v1.EchoService/EchoCustomer',
            prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.SerializeToString,
            prototpl_dot_echo__service_dot_echo__service__pb2.StringMessage.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
