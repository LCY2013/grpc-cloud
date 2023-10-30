// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gen/proto/person-service/person-service.proto

package io.grpc.gateway.service.v1;

public final class PersonServiceProto {
  private PersonServiceProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface PersonMessageOrBuilder extends
      // @@protoc_insertion_point(interface_extends:grpc.gateway.service.v1.PersonMessage)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string name = 1 [json_name = "name"];</code>
     * @return The name.
     */
    java.lang.String getName();
    /**
     * <code>string name = 1 [json_name = "name"];</code>
     * @return The bytes for name.
     */
    com.google.protobuf.ByteString
        getNameBytes();

    /**
     * <code>string age = 2 [json_name = "age"];</code>
     * @return The age.
     */
    java.lang.String getAge();
    /**
     * <code>string age = 2 [json_name = "age"];</code>
     * @return The bytes for age.
     */
    com.google.protobuf.ByteString
        getAgeBytes();
  }
  /**
   * Protobuf type {@code grpc.gateway.service.v1.PersonMessage}
   */
  public static final class PersonMessage extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:grpc.gateway.service.v1.PersonMessage)
      PersonMessageOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use PersonMessage.newBuilder() to construct.
    private PersonMessage(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private PersonMessage() {
      name_ = "";
      age_ = "";
    }

    @java.lang.Override
    @SuppressWarnings({"unused"})
    protected java.lang.Object newInstance(
        UnusedPrivateParameter unused) {
      return new PersonMessage();
    }

    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.grpc.gateway.service.v1.PersonServiceProto.internal_static_grpc_gateway_service_v1_PersonMessage_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.grpc.gateway.service.v1.PersonServiceProto.internal_static_grpc_gateway_service_v1_PersonMessage_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage.class, io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage.Builder.class);
    }

    public static final int NAME_FIELD_NUMBER = 1;
    @SuppressWarnings("serial")
    private volatile java.lang.Object name_ = "";
    /**
     * <code>string name = 1 [json_name = "name"];</code>
     * @return The name.
     */
    @java.lang.Override
    public java.lang.String getName() {
      java.lang.Object ref = name_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        name_ = s;
        return s;
      }
    }
    /**
     * <code>string name = 1 [json_name = "name"];</code>
     * @return The bytes for name.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getNameBytes() {
      java.lang.Object ref = name_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        name_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int AGE_FIELD_NUMBER = 2;
    @SuppressWarnings("serial")
    private volatile java.lang.Object age_ = "";
    /**
     * <code>string age = 2 [json_name = "age"];</code>
     * @return The age.
     */
    @java.lang.Override
    public java.lang.String getAge() {
      java.lang.Object ref = age_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        age_ = s;
        return s;
      }
    }
    /**
     * <code>string age = 2 [json_name = "age"];</code>
     * @return The bytes for age.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getAgeBytes() {
      java.lang.Object ref = age_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        age_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }

    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage parseFrom(
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
    public static Builder newBuilder(io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage prototype) {
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
     * Protobuf type {@code grpc.gateway.service.v1.PersonMessage}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:grpc.gateway.service.v1.PersonMessage)
        io.grpc.gateway.service.v1.PersonServiceProto.PersonMessageOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return io.grpc.gateway.service.v1.PersonServiceProto.internal_static_grpc_gateway_service_v1_PersonMessage_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return io.grpc.gateway.service.v1.PersonServiceProto.internal_static_grpc_gateway_service_v1_PersonMessage_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage.class, io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage.Builder.class);
      }

      // Construct using io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage.newBuilder()
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
        name_ = "";
        age_ = "";
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return io.grpc.gateway.service.v1.PersonServiceProto.internal_static_grpc_gateway_service_v1_PersonMessage_descriptor;
      }

      @java.lang.Override
      public io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage getDefaultInstanceForType() {
        return io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage.getDefaultInstance();
      }

      @java.lang.Override
      public io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage build() {
        io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage buildPartial() {
        io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage result = new io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage(this);
        if (bitField0_ != 0) { buildPartial0(result); }
        onBuilt();
        return result;
      }

      private void buildPartial0(io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage result) {
        int from_bitField0_ = bitField0_;
        if (((from_bitField0_ & 0x00000001) != 0)) {
          result.name_ = name_;
        }
        if (((from_bitField0_ & 0x00000002) != 0)) {
          result.age_ = age_;
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

      private java.lang.Object name_ = "";
      /**
       * <code>string name = 1 [json_name = "name"];</code>
       * @return The name.
       */
      public java.lang.String getName() {
        java.lang.Object ref = name_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          name_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string name = 1 [json_name = "name"];</code>
       * @return The bytes for name.
       */
      public com.google.protobuf.ByteString
          getNameBytes() {
        java.lang.Object ref = name_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          name_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string name = 1 [json_name = "name"];</code>
       * @param value The name to set.
       * @return This builder for chaining.
       */
      public Builder setName(
          java.lang.String value) {
        if (value == null) { throw new NullPointerException(); }
        name_ = value;
        bitField0_ |= 0x00000001;
        onChanged();
        return this;
      }
      /**
       * <code>string name = 1 [json_name = "name"];</code>
       * @return This builder for chaining.
       */
      public Builder clearName() {
        name_ = getDefaultInstance().getName();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
        return this;
      }
      /**
       * <code>string name = 1 [json_name = "name"];</code>
       * @param value The bytes for name to set.
       * @return This builder for chaining.
       */
      public Builder setNameBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) { throw new NullPointerException(); }
        checkByteStringIsUtf8(value);
        name_ = value;
        bitField0_ |= 0x00000001;
        onChanged();
        return this;
      }

      private java.lang.Object age_ = "";
      /**
       * <code>string age = 2 [json_name = "age"];</code>
       * @return The age.
       */
      public java.lang.String getAge() {
        java.lang.Object ref = age_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          age_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string age = 2 [json_name = "age"];</code>
       * @return The bytes for age.
       */
      public com.google.protobuf.ByteString
          getAgeBytes() {
        java.lang.Object ref = age_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          age_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string age = 2 [json_name = "age"];</code>
       * @param value The age to set.
       * @return This builder for chaining.
       */
      public Builder setAge(
          java.lang.String value) {
        if (value == null) { throw new NullPointerException(); }
        age_ = value;
        bitField0_ |= 0x00000002;
        onChanged();
        return this;
      }
      /**
       * <code>string age = 2 [json_name = "age"];</code>
       * @return This builder for chaining.
       */
      public Builder clearAge() {
        age_ = getDefaultInstance().getAge();
        bitField0_ = (bitField0_ & ~0x00000002);
        onChanged();
        return this;
      }
      /**
       * <code>string age = 2 [json_name = "age"];</code>
       * @param value The bytes for age to set.
       * @return This builder for chaining.
       */
      public Builder setAgeBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) { throw new NullPointerException(); }
        checkByteStringIsUtf8(value);
        age_ = value;
        bitField0_ |= 0x00000002;
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


      // @@protoc_insertion_point(builder_scope:grpc.gateway.service.v1.PersonMessage)
    }

    // @@protoc_insertion_point(class_scope:grpc.gateway.service.v1.PersonMessage)
    private static final io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage();
    }

    public static io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<PersonMessage>
        PARSER = new com.google.protobuf.AbstractParser<PersonMessage>() {
      @java.lang.Override
      public PersonMessage parsePartialFrom(
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

    public static com.google.protobuf.Parser<PersonMessage> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<PersonMessage> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public io.grpc.gateway.service.v1.PersonServiceProto.PersonMessage getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_grpc_gateway_service_v1_PersonMessage_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_grpc_gateway_service_v1_PersonMessage_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n-gen/proto/person-service/person-servic" +
      "e.proto\022\027grpc.gateway.service.v1\032\034google" +
      "/api/annotations.proto\032;grpc-gateway/pro" +
      "toc-gen-openapiv2/options/annotations.pr" +
      "oto\"5\n\rPersonMessage\022\022\n\004name\030\001 \001(\tR\004name" +
      "\022\020\n\003age\030\002 \001(\tR\003age2\324\001\n\rPersonService\022\302\001\n" +
      "\tGetPerson\022&.grpc.gateway.service.v1.Per" +
      "sonMessage\032&.grpc.gateway.service.v1.Per" +
      "sonMessage\"e\222AA\n\016person_service\022\021get a p" +
      "erson info\032\004desc*\tgetPersonJ\013\n\003200\022\004\n\002OK" +
      "\202\323\344\223\002\033\"\026/person/service/v1/get:\001*B\321\003\n\032io" +
      ".grpc.gateway.service.v1B\022PersonServiceP" +
      "rotoH\002Z%go/gen/proto/person-service;serv" +
      "icev1\242\002\003GGS\252\002\027Grpc.Gateway.Service.V1\312\002\027" +
      "Grpc\\Gateway\\Service\\V1\342\002#Grpc\\Gateway\\S" +
      "ervice\\V1\\GPBMetadata\352\002\032Grpc::Gateway::S" +
      "ervice::V1\222A\367\001\022\273\001\n\nPerson API\"Q\n\026person_" +
      "service project\022%https://github.com/LCY2" +
      "013/grpc-cloud\032\020none@example.com*U\n\032Apac" +
      "he License Version 2.0\0227https://github.c" +
      "om/LCY2013/grpc-cloud/blob/main/LICENSE2" +
      "\0031.0*\001\0022\020application/json:\020application/j" +
      "sonj\020\n\016person_serviceb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          io.google.api.AnnotationsProto.getDescriptor(),
          io.grpc.gateway.protoc_gen_openapiv2.options.AnnotationsProto.getDescriptor(),
        });
    internal_static_grpc_gateway_service_v1_PersonMessage_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_grpc_gateway_service_v1_PersonMessage_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_grpc_gateway_service_v1_PersonMessage_descriptor,
        new java.lang.String[] { "Name", "Age", });
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
