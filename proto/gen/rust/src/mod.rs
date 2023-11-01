// @generated
pub mod grpc {
    pub mod echo {
        pub mod service {
            #[cfg(feature = "grpc-echo-service-v1")]
            // @@protoc_insertion_point(attribute:grpc.echo.service.v1)
            pub mod v1 {
                include!("grpc.echo.service.v1.rs");
                // @@protoc_insertion_point(grpc.echo.service.v1)
            }
        }
    }
    pub mod person {
        pub mod service {
            #[cfg(feature = "grpc-person-service-v1")]
            // @@protoc_insertion_point(attribute:grpc.person.service.v1)
            pub mod v1 {
                include!("grpc.person.service.v1.rs");
                // @@protoc_insertion_point(grpc.person.service.v1)
            }
        }
    }
}