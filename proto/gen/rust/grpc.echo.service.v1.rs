// @generated
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct StringMessage {
    #[prost(string, tag="1")]
    pub value: ::prost::alloc::string::String,
}
include!("grpc.echo.service.v1.serde.rs");
// @@protoc_insertion_point(module)