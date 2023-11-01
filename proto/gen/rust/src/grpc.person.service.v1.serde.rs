// @generated
impl serde::Serialize for PersonMessage {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.name.is_empty() {
            len += 1;
        }
        if !self.age.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("grpc.person.service.v1.PersonMessage", len)?;
        if !self.name.is_empty() {
            struct_ser.serialize_field("name", &self.name)?;
        }
        if !self.age.is_empty() {
            struct_ser.serialize_field("age", &self.age)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for PersonMessage {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "name",
            "age",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Name,
            Age,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "name" => Ok(GeneratedField::Name),
                            "age" => Ok(GeneratedField::Age),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = PersonMessage;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct grpc.person.service.v1.PersonMessage")
            }

            fn visit_map<V>(self, mut map: V) -> std::result::Result<PersonMessage, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut name__ = None;
                let mut age__ = None;
                while let Some(k) = map.next_key()? {
                    match k {
                        GeneratedField::Name => {
                            if name__.is_some() {
                                return Err(serde::de::Error::duplicate_field("name"));
                            }
                            name__ = Some(map.next_value()?);
                        }
                        GeneratedField::Age => {
                            if age__.is_some() {
                                return Err(serde::de::Error::duplicate_field("age"));
                            }
                            age__ = Some(map.next_value()?);
                        }
                    }
                }
                Ok(PersonMessage {
                    name: name__.unwrap_or_default(),
                    age: age__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("grpc.person.service.v1.PersonMessage", FIELDS, GeneratedVisitor)
    }
}
