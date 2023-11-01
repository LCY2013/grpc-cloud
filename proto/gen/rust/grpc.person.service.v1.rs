// @generated
#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct PersonMessage {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
    #[prost(string, tag="2")]
    pub age: ::prost::alloc::string::String,
}
include!("grpc.person.service.v1.serde.rs");
// @@protoc_insertion_point(module)