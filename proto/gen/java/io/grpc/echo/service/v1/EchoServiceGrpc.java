package io.grpc.echo.service.v1;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.59.0)",
    comments = "Source: prototpl/echo-service/echo-service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class EchoServiceGrpc {

  private EchoServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "grpc.echo.service.v1.EchoService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<io.grpc.echo.service.v1.EchoServiceProto.StringMessage,
      io.grpc.echo.service.v1.EchoServiceProto.StringMessage> getEchoMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Echo",
      requestType = io.grpc.echo.service.v1.EchoServiceProto.StringMessage.class,
      responseType = io.grpc.echo.service.v1.EchoServiceProto.StringMessage.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.grpc.echo.service.v1.EchoServiceProto.StringMessage,
      io.grpc.echo.service.v1.EchoServiceProto.StringMessage> getEchoMethod() {
    io.grpc.MethodDescriptor<io.grpc.echo.service.v1.EchoServiceProto.StringMessage, io.grpc.echo.service.v1.EchoServiceProto.StringMessage> getEchoMethod;
    if ((getEchoMethod = EchoServiceGrpc.getEchoMethod) == null) {
      synchronized (EchoServiceGrpc.class) {
        if ((getEchoMethod = EchoServiceGrpc.getEchoMethod) == null) {
          EchoServiceGrpc.getEchoMethod = getEchoMethod =
              io.grpc.MethodDescriptor.<io.grpc.echo.service.v1.EchoServiceProto.StringMessage, io.grpc.echo.service.v1.EchoServiceProto.StringMessage>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Echo"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.echo.service.v1.EchoServiceProto.StringMessage.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.echo.service.v1.EchoServiceProto.StringMessage.getDefaultInstance()))
              .setSchemaDescriptor(new EchoServiceMethodDescriptorSupplier("Echo"))
              .build();
        }
      }
    }
    return getEchoMethod;
  }

  private static volatile io.grpc.MethodDescriptor<io.grpc.echo.service.v1.EchoServiceProto.StringMessage,
      io.grpc.echo.service.v1.EchoServiceProto.StringMessage> getEchoCustomerMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "EchoCustomer",
      requestType = io.grpc.echo.service.v1.EchoServiceProto.StringMessage.class,
      responseType = io.grpc.echo.service.v1.EchoServiceProto.StringMessage.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<io.grpc.echo.service.v1.EchoServiceProto.StringMessage,
      io.grpc.echo.service.v1.EchoServiceProto.StringMessage> getEchoCustomerMethod() {
    io.grpc.MethodDescriptor<io.grpc.echo.service.v1.EchoServiceProto.StringMessage, io.grpc.echo.service.v1.EchoServiceProto.StringMessage> getEchoCustomerMethod;
    if ((getEchoCustomerMethod = EchoServiceGrpc.getEchoCustomerMethod) == null) {
      synchronized (EchoServiceGrpc.class) {
        if ((getEchoCustomerMethod = EchoServiceGrpc.getEchoCustomerMethod) == null) {
          EchoServiceGrpc.getEchoCustomerMethod = getEchoCustomerMethod =
              io.grpc.MethodDescriptor.<io.grpc.echo.service.v1.EchoServiceProto.StringMessage, io.grpc.echo.service.v1.EchoServiceProto.StringMessage>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "EchoCustomer"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.echo.service.v1.EchoServiceProto.StringMessage.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  io.grpc.echo.service.v1.EchoServiceProto.StringMessage.getDefaultInstance()))
              .setSchemaDescriptor(new EchoServiceMethodDescriptorSupplier("EchoCustomer"))
              .build();
        }
      }
    }
    return getEchoCustomerMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static EchoServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EchoServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EchoServiceStub>() {
        @java.lang.Override
        public EchoServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EchoServiceStub(channel, callOptions);
        }
      };
    return EchoServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static EchoServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EchoServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EchoServiceBlockingStub>() {
        @java.lang.Override
        public EchoServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EchoServiceBlockingStub(channel, callOptions);
        }
      };
    return EchoServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static EchoServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EchoServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EchoServiceFutureStub>() {
        @java.lang.Override
        public EchoServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EchoServiceFutureStub(channel, callOptions);
        }
      };
    return EchoServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void echo(io.grpc.echo.service.v1.EchoServiceProto.StringMessage request,
        io.grpc.stub.StreamObserver<io.grpc.echo.service.v1.EchoServiceProto.StringMessage> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getEchoMethod(), responseObserver);
    }

    /**
     */
    default void echoCustomer(io.grpc.echo.service.v1.EchoServiceProto.StringMessage request,
        io.grpc.stub.StreamObserver<io.grpc.echo.service.v1.EchoServiceProto.StringMessage> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getEchoCustomerMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service EchoService.
   */
  public static abstract class EchoServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return EchoServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service EchoService.
   */
  public static final class EchoServiceStub
      extends io.grpc.stub.AbstractAsyncStub<EchoServiceStub> {
    private EchoServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EchoServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EchoServiceStub(channel, callOptions);
    }

    /**
     */
    public void echo(io.grpc.echo.service.v1.EchoServiceProto.StringMessage request,
        io.grpc.stub.StreamObserver<io.grpc.echo.service.v1.EchoServiceProto.StringMessage> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getEchoMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void echoCustomer(io.grpc.echo.service.v1.EchoServiceProto.StringMessage request,
        io.grpc.stub.StreamObserver<io.grpc.echo.service.v1.EchoServiceProto.StringMessage> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getEchoCustomerMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service EchoService.
   */
  public static final class EchoServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<EchoServiceBlockingStub> {
    private EchoServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EchoServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EchoServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public io.grpc.echo.service.v1.EchoServiceProto.StringMessage echo(io.grpc.echo.service.v1.EchoServiceProto.StringMessage request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getEchoMethod(), getCallOptions(), request);
    }

    /**
     */
    public io.grpc.echo.service.v1.EchoServiceProto.StringMessage echoCustomer(io.grpc.echo.service.v1.EchoServiceProto.StringMessage request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getEchoCustomerMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service EchoService.
   */
  public static final class EchoServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<EchoServiceFutureStub> {
    private EchoServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EchoServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EchoServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.grpc.echo.service.v1.EchoServiceProto.StringMessage> echo(
        io.grpc.echo.service.v1.EchoServiceProto.StringMessage request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getEchoMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<io.grpc.echo.service.v1.EchoServiceProto.StringMessage> echoCustomer(
        io.grpc.echo.service.v1.EchoServiceProto.StringMessage request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getEchoCustomerMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_ECHO = 0;
  private static final int METHODID_ECHO_CUSTOMER = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_ECHO:
          serviceImpl.echo((io.grpc.echo.service.v1.EchoServiceProto.StringMessage) request,
              (io.grpc.stub.StreamObserver<io.grpc.echo.service.v1.EchoServiceProto.StringMessage>) responseObserver);
          break;
        case METHODID_ECHO_CUSTOMER:
          serviceImpl.echoCustomer((io.grpc.echo.service.v1.EchoServiceProto.StringMessage) request,
              (io.grpc.stub.StreamObserver<io.grpc.echo.service.v1.EchoServiceProto.StringMessage>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getEchoMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.grpc.echo.service.v1.EchoServiceProto.StringMessage,
              io.grpc.echo.service.v1.EchoServiceProto.StringMessage>(
                service, METHODID_ECHO)))
        .addMethod(
          getEchoCustomerMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              io.grpc.echo.service.v1.EchoServiceProto.StringMessage,
              io.grpc.echo.service.v1.EchoServiceProto.StringMessage>(
                service, METHODID_ECHO_CUSTOMER)))
        .build();
  }

  private static abstract class EchoServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    EchoServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return io.grpc.echo.service.v1.EchoServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("EchoService");
    }
  }

  private static final class EchoServiceFileDescriptorSupplier
      extends EchoServiceBaseDescriptorSupplier {
    EchoServiceFileDescriptorSupplier() {}
  }

  private static final class EchoServiceMethodDescriptorSupplier
      extends EchoServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    EchoServiceMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (EchoServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new EchoServiceFileDescriptorSupplier())
              .addMethod(getEchoMethod())
              .addMethod(getEchoCustomerMethod())
              .build();
        }
      }
    }
    return result;
  }
}
