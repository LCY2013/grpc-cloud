// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: prototpl/person-service/person-service.proto

// Protobuf Java Version: 3.25.1
package io.grpc.person.service.v1;

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
      // @@protoc_insertion_point(interface_extends:grpc.person.service.v1.PersonMessage)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
     * @return The name.
     */
    java.lang.String getName();
    /**
     * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
     * @return The bytes for name.
     */
    com.google.protobuf.ByteString
        getNameBytes();

    /**
     * <code>int32 age = 2 [json_name = "age", (.buf.validate.field) = { ... }</code>
     * @return The age.
     */
    int getAge();
  }
  /**
   * Protobuf type {@code grpc.person.service.v1.PersonMessage}
   */
  public static final class PersonMessage extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:grpc.person.service.v1.PersonMessage)
      PersonMessageOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use PersonMessage.newBuilder() to construct.
    private PersonMessage(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private PersonMessage() {
      name_ = "";
    }

    @java.lang.Override
    @SuppressWarnings({"unused"})
    protected java.lang.Object newInstance(
        UnusedPrivateParameter unused) {
      return new PersonMessage();
    }

    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.grpc.person.service.v1.PersonServiceProto.internal_static_grpc_person_service_v1_PersonMessage_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.grpc.person.service.v1.PersonServiceProto.internal_static_grpc_person_service_v1_PersonMessage_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.grpc.person.service.v1.PersonServiceProto.PersonMessage.class, io.grpc.person.service.v1.PersonServiceProto.PersonMessage.Builder.class);
    }

    public static final int NAME_FIELD_NUMBER = 1;
    @SuppressWarnings("serial")
    private volatile java.lang.Object name_ = "";
    /**
     * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
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
     * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
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
    private int age_ = 0;
    /**
     * <code>int32 age = 2 [json_name = "age", (.buf.validate.field) = { ... }</code>
     * @return The age.
     */
    @java.lang.Override
    public int getAge() {
      return age_;
    }

    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }

    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage parseFrom(
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
    public static Builder newBuilder(io.grpc.person.service.v1.PersonServiceProto.PersonMessage prototype) {
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
     * Protobuf type {@code grpc.person.service.v1.PersonMessage}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:grpc.person.service.v1.PersonMessage)
        io.grpc.person.service.v1.PersonServiceProto.PersonMessageOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return io.grpc.person.service.v1.PersonServiceProto.internal_static_grpc_person_service_v1_PersonMessage_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return io.grpc.person.service.v1.PersonServiceProto.internal_static_grpc_person_service_v1_PersonMessage_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                io.grpc.person.service.v1.PersonServiceProto.PersonMessage.class, io.grpc.person.service.v1.PersonServiceProto.PersonMessage.Builder.class);
      }

      // Construct using io.grpc.person.service.v1.PersonServiceProto.PersonMessage.newBuilder()
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
        age_ = 0;
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return io.grpc.person.service.v1.PersonServiceProto.internal_static_grpc_person_service_v1_PersonMessage_descriptor;
      }

      @java.lang.Override
      public io.grpc.person.service.v1.PersonServiceProto.PersonMessage getDefaultInstanceForType() {
        return io.grpc.person.service.v1.PersonServiceProto.PersonMessage.getDefaultInstance();
      }

      @java.lang.Override
      public io.grpc.person.service.v1.PersonServiceProto.PersonMessage build() {
        io.grpc.person.service.v1.PersonServiceProto.PersonMessage result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public io.grpc.person.service.v1.PersonServiceProto.PersonMessage buildPartial() {
        io.grpc.person.service.v1.PersonServiceProto.PersonMessage result = new io.grpc.person.service.v1.PersonServiceProto.PersonMessage(this);
        if (bitField0_ != 0) { buildPartial0(result); }
        onBuilt();
        return result;
      }

      private void buildPartial0(io.grpc.person.service.v1.PersonServiceProto.PersonMessage result) {
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
       * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
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
       * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
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
       * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
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
       * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
       * @return This builder for chaining.
       */
      public Builder clearName() {
        name_ = getDefaultInstance().getName();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
        return this;
      }
      /**
       * <code>string name = 1 [json_name = "name", (.buf.validate.field) = { ... }</code>
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

      private int age_ ;
      /**
       * <code>int32 age = 2 [json_name = "age", (.buf.validate.field) = { ... }</code>
       * @return The age.
       */
      @java.lang.Override
      public int getAge() {
        return age_;
      }
      /**
       * <code>int32 age = 2 [json_name = "age", (.buf.validate.field) = { ... }</code>
       * @param value The age to set.
       * @return This builder for chaining.
       */
      public Builder setAge(int value) {

        age_ = value;
        bitField0_ |= 0x00000002;
        onChanged();
        return this;
      }
      /**
       * <code>int32 age = 2 [json_name = "age", (.buf.validate.field) = { ... }</code>
       * @return This builder for chaining.
       */
      public Builder clearAge() {
        bitField0_ = (bitField0_ & ~0x00000002);
        age_ = 0;
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


      // @@protoc_insertion_point(builder_scope:grpc.person.service.v1.PersonMessage)
    }

    // @@protoc_insertion_point(class_scope:grpc.person.service.v1.PersonMessage)
    private static final io.grpc.person.service.v1.PersonServiceProto.PersonMessage DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new io.grpc.person.service.v1.PersonServiceProto.PersonMessage();
    }

    public static io.grpc.person.service.v1.PersonServiceProto.PersonMessage getDefaultInstance() {
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
    public io.grpc.person.service.v1.PersonServiceProto.PersonMessage getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_grpc_person_service_v1_PersonMessage_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_grpc_person_service_v1_PersonMessage_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n,prototpl/person-service/person-service" +
      ".proto\022\026grpc.person.service.v1\032\033buf/vali" +
      "date/validate.proto\032\034google/api/annotati" +
      "ons.proto\032.protoc-gen-openapiv2/options/" +
      "annotations.proto\"G\n\rPersonMessage\022\033\n\004na" +
      "me\030\001 \001(\tB\007\272H\004r\002\020\005R\004name\022\031\n\003age\030\002 \001(\005B\007\272H" +
      "\004\032\002 \024R\003age2\322\001\n\rPersonService\022\300\001\n\tGetPers" +
      "on\022%.grpc.person.service.v1.PersonMessag" +
      "e\032%.grpc.person.service.v1.PersonMessage" +
      "\"e\222AA\n\016person_service\022\021get a person info" +
      "\032\004desc*\tgetPersonJ\013\n\003200\022\004\n\002OK\202\323\344\223\002\033\"\026/p" +
      "erson/service/v1/get:\001*B\275\003\n\031io.grpc.pers" +
      "on.service.v1B\022PersonServiceProtoH\002Z\026grp" +
      "c/person/service/v1\242\002\003GPS\252\002\026Grpc.Person." +
      "Service.V1\312\002\026Grpc\\Person\\Service\\V1\342\002\"Gr" +
      "pc\\Person\\Service\\V1\\GPBMetadata\352\002\031Grpc:" +
      ":Person::Service::V1\222A\367\001\022\273\001\n\nPerson API\"" +
      "Q\n\026person_service project\022%https://githu" +
      "b.com/LCY2013/grpc-cloud\032\020none@example.c" +
      "om*U\n\032Apache License Version 2.0\0227https:" +
      "//github.com/LCY2013/grpc-cloud/blob/mai" +
      "n/LICENSE2\0031.0*\001\0022\020application/json:\020app" +
      "lication/jsonj\020\n\016person_serviceb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          io.buf.validate.ValidateProto.getDescriptor(),
          io.google.api.AnnotationsProto.getDescriptor(),
          io.grpc.gateway.protoc_gen_openapiv2.options.AnnotationsProto.getDescriptor(),
        });
    internal_static_grpc_person_service_v1_PersonMessage_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_grpc_person_service_v1_PersonMessage_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_grpc_person_service_v1_PersonMessage_descriptor,
        new java.lang.String[] { "Name", "Age", });
    com.google.protobuf.ExtensionRegistry registry =
        com.google.protobuf.ExtensionRegistry.newInstance();
    registry.add(io.buf.validate.ValidateProto.field);
    registry.add(io.google.api.AnnotationsProto.http);
    registry.add(io.grpc.gateway.protoc_gen_openapiv2.options.AnnotationsProto.openapiv2Operation);
    registry.add(io.grpc.gateway.protoc_gen_openapiv2.options.AnnotationsProto.openapiv2Swagger);
    com.google.protobuf.Descriptors.FileDescriptor
        .internalUpdateFileDescriptor(descriptor, registry);
    io.buf.validate.ValidateProto.getDescriptor();
    io.google.api.AnnotationsProto.getDescriptor();
    io.grpc.gateway.protoc_gen_openapiv2.options.AnnotationsProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
