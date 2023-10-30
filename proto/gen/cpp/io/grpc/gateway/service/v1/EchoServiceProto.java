// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gen/go/echo-service/echo-service.proto

package io.grpc.gateway.service.v1;

public final class EchoServiceProto {
  private EchoServiceProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface StringMessageOrBuilder extends
      // @@protoc_insertion_point(interface_extends:grpc.gateway.service.v1.StringMessage)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string value = 1 [json_name = "value"];</code>
     * @return The value.
     */
    java.lang.String getValue();
    /**
     * <code>string value = 1 [json_name = "value"];</code>
     * @return The bytes for value.
     */
    com.google.protobuf.ByteString
        getValueBytes();
  }
  /**
   * Protobuf type {@code grpc.gateway.service.v1.StringMessage}
   */
  public static final class StringMessage extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:grpc.gateway.service.v1.StringMessage)
      StringMessageOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use StringMessage.newBuilder() to construct.
    private StringMessage(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private StringMessage() {
      value_ = "";
    }

    @java.lang.Override
    @SuppressWarnings({"unused"})
    protected java.lang.Object newInstance(
        UnusedPrivateParameter unused) {
      return new StringMessage();
    }

    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.grpc.gateway.service.v1.EchoServiceProto.internal_static_grpc_gateway_service_v1_StringMessage_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.grpc.gateway.service.v1.EchoServiceProto.internal_static_grpc_gateway_service_v1_StringMessage_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.grpc.gateway.service.v1.EchoServiceProto.StringMessage.class, io.grpc.gateway.service.v1.EchoServiceProto.StringMessage.Builder.class);
    }

    public static final int VALUE_FIELD_NUMBER = 1;
    @SuppressWarnings("serial")
    private volatile java.lang.Object value_ = "";
    /**
     * <code>string value = 1 [json_name = "value"];</code>
     * @return The value.
     */
    @java.lang.Override
    public java.lang.String getValue() {
      java.lang.Object ref = value_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        value_ = s;
        return s;
      }
    }
    /**
     * <code>string value = 1 [json_name = "value"];</code>
     * @return The bytes for value.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getValueBytes() {
      java.lang.Object ref = value_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        value_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }

    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(io.grpc.gateway.service.v1.EchoServiceProto.StringMessage prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code grpc.gateway.service.v1.StringMessage}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:grpc.gateway.service.v1.StringMessage)
        io.grpc.gateway.service.v1.EchoServiceProto.StringMessageOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return io.grpc.gateway.service.v1.EchoServiceProto.internal_static_grpc_gateway_service_v1_StringMessage_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return io.grpc.gateway.service.v1.EchoServiceProto.internal_static_grpc_gateway_service_v1_StringMessage_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                io.grpc.gateway.service.v1.EchoServiceProto.StringMessage.class, io.grpc.gateway.service.v1.EchoServiceProto.StringMessage.Builder.class);
      }

      // Construct using io.grpc.gateway.service.v1.EchoServiceProto.StringMessage.newBuilder()
      private Builder() {

      }

      private Builder(
          com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
        super(parent);

      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        bitField0_ = 0;
        value_ = "";
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return io.grpc.gateway.service.v1.EchoServiceProto.internal_static_grpc_gateway_service_v1_StringMessage_descriptor;
      }

      @java.lang.Override
      public io.grpc.gateway.service.v1.EchoServiceProto.StringMessage getDefaultInstanceForType() {
        return io.grpc.gateway.service.v1.EchoServiceProto.StringMessage.getDefaultInstance();
      }

      @java.lang.Override
      public io.grpc.gateway.service.v1.EchoServiceProto.StringMessage build() {
        io.grpc.gateway.service.v1.EchoServiceProto.StringMessage result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public io.grpc.gateway.service.v1.EchoServiceProto.StringMessage buildPartial() {
        io.grpc.gateway.service.v1.EchoServiceProto.StringMessage result = new io.grpc.gateway.service.v1.EchoServiceProto.StringMessage(this);
        if (bitField0_ != 0) { buildPartial0(result); }
        onBuilt();
        return result;
      }

      private void buildPartial0(io.grpc.gateway.service.v1.EchoServiceProto.StringMessage result) {
        int from_bitField0_ = bitField0_;
        if (((from_bitField0_ & 0x00000001) != 0)) {
          result.value_ = value_;
        }
      }

      @java.lang.Override
      public Builder clone() {
        return super.clone();
      }
      @java.lang.Override
      public Builder setField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.setField(field, value);
      }
      @java.lang.Override
      public Builder clearField(
          com.google.protobuf.Descriptors.FieldDescriptor field) {
        return super.clearField(field);
      }
      @java.lang.Override
      public Builder clearOneof(
          com.google.protobuf.Descriptors.OneofDescriptor oneof) {
        return super.clearOneof(oneof);
      }
      @java.lang.Override
      public Builder setRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          int index, java.lang.Object value) {
        return super.setRepeatedField(field, index, value);
      }
      @java.lang.Override
      public Builder addRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.addRepeatedField(field, value);
      }
      private int bitField0_;

      private java.lang.Object value_ = "";
      /**
       * <code>string value = 1 [json_name = "value"];</code>
       * @return The value.
       */
      public java.lang.String getValue() {
        java.lang.Object ref = value_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          value_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string value = 1 [json_name = "value"];</code>
       * @return The bytes for value.
       */
      public com.google.protobuf.ByteString
          getValueBytes() {
        java.lang.Object ref = value_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          value_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string value = 1 [json_name = "value"];</code>
       * @param value The value to set.
       * @return This builder for chaining.
       */
      public Builder setValue(
          java.lang.String value) {
        if (value == null) { throw new NullPointerException(); }
        value_ = value;
        bitField0_ |= 0x00000001;
        onChanged();
        return this;
      }
      /**
       * <code>string value = 1 [json_name = "value"];</code>
       * @return This builder for chaining.
       */
      public Builder clearValue() {
        value_ = getDefaultInstance().getValue();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
        return this;
      }
      /**
       * <code>string value = 1 [json_name = "value"];</code>
       * @param value The bytes for value to set.
       * @return This builder for chaining.
       */
      public Builder setValueBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) { throw new NullPointerException(); }
        checkByteStringIsUtf8(value);
        value_ = value;
        bitField0_ |= 0x00000001;
        onChanged();
        return this;
      }
      @java.lang.Override
      public final Builder setUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.setUnknownFields(unknownFields);
      }

      @java.lang.Override
      public final Builder mergeUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.mergeUnknownFields(unknownFields);
      }


      // @@protoc_insertion_point(builder_scope:grpc.gateway.service.v1.StringMessage)
    }

    // @@protoc_insertion_point(class_scope:grpc.gateway.service.v1.StringMessage)
    private static final io.grpc.gateway.service.v1.EchoServiceProto.StringMessage DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new io.grpc.gateway.service.v1.EchoServiceProto.StringMessage();
    }

    public static io.grpc.gateway.service.v1.EchoServiceProto.StringMessage getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<StringMessage>
        PARSER = new com.google.protobuf.AbstractParser<StringMessage>() {
      @java.lang.Override
      public StringMessage parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        Builder builder = newBuilder();
        try {
          builder.mergeFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw e.setUnfinishedMessage(builder.buildPartial());
        } catch (com.google.protobuf.UninitializedMessageException e) {
          throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
        } catch (java.io.IOException e) {
          throw new com.google.protobuf.InvalidProtocolBufferException(e)
              .setUnfinishedMessage(builder.buildPartial());
        }
        return builder.buildPartial();
      }
    };

    public static com.google.protobuf.Parser<StringMessage> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<StringMessage> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public io.grpc.gateway.service.v1.EchoServiceProto.StringMessage getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_grpc_gateway_service_v1_StringMessage_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_grpc_gateway_service_v1_StringMessage_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n&gen/go/echo-service/echo-service.proto" +
      "\022\027grpc.gateway.service.v1\032\034google/api/an" +
      "notations.proto\032;grpc-gateway/protoc-gen" +
      "-openapiv2/options/annotations.proto\"%\n\r" +
      "StringMessage\022\024\n\005value\030\001 \001(\tR\005value2\242\003\n\013" +
      "EchoService\022\300\001\n\004Echo\022&.grpc.gateway.serv" +
      "ice.v1.StringMessage\032&.grpc.gateway.serv" +
      "ice.v1.StringMessage\"h\222A=\n\014echo-service\022" +
      "\016echo a message\032\004desc*\ngetMessageJ\013\n\003200" +
      "\022\004\n\002OK\202\323\344\223\002\"\"\035/echo/service/v1/example/e" +
      "cho:\001*\022\317\001\n\014EchoCustomer\022&.grpc.gateway.s" +
      "ervice.v1.StringMessage\032&.grpc.gateway.s" +
      "ervice.v1.StringMessage\"o\222A=\n\014echo-servi" +
      "ce\022\016echo a message\032\004desc*\ngetMessageJ\013\n\003" +
      "200\022\004\n\002OK\202\323\344\223\002):\001*B$\n\001*\022\037/echo/service/v" +
      "1/example/echo/cB\344\003\n\032io.grpc.gateway.ser" +
      "vice.v1B\020EchoServiceProtoH\002Z@github.com/" +
      "grpc-cloud/proto/gen/go/gen/go/echo-serv" +
      "ice;servicev1\242\002\003GGS\252\002\027Grpc.Gateway.Servi" +
      "ce.V1\312\002\027Grpc\\Gateway\\Service\\V1\342\002#Grpc\\G" +
      "ateway\\Service\\V1\\GPBMetadata\352\002\032Grpc::Ga" +
      "teway::Service::V1\222A\361\001\022\267\001\n\010Echo API\"O\n\024e" +
      "cho_service project\022%https://github.com/" +
      "LCY2013/grpc-cloud\032\020none@example.com*U\n\032" +
      "Apache License Version 2.0\0227https://gith" +
      "ub.com/LCY2013/grpc-cloud/blob/main/LICE" +
      "NSE2\0031.0*\001\0022\020application/json:\020applicati" +
      "on/jsonj\016\n\014echo-serviceb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          io.google.api.AnnotationsProto.getDescriptor(),
          io.grpc.gateway.protoc_gen_openapiv2.options.AnnotationsProto.getDescriptor(),
        });
    internal_static_grpc_gateway_service_v1_StringMessage_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_grpc_gateway_service_v1_StringMessage_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_grpc_gateway_service_v1_StringMessage_descriptor,
        new java.lang.String[] { "Value", });
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(io.google.api.AnnotationsProto.http);
    registry.add(io.grpc.gateway.protoc_gen_openapiv2.options.AnnotationsProto.openapiv2Operation);
    registry.add(io.grpc.gateway.protoc_gen_openapiv2.options.AnnotationsProto.openapiv2Swagger);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    io.google.api.AnnotationsProto.getDescriptor();
    io.grpc.gateway.protoc_gen_openapiv2.options.AnnotationsProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
