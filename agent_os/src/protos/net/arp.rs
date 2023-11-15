// This file is generated by rust-protobuf 3.3.0. Do not edit
// .proto file is parsed by protoc 3.12.4
// @generated

// https://github.com/rust-lang/rust-clippy/issues/702
#![allow(unknown_lints)]
#![allow(clippy::all)]

#![allow(unused_attributes)]
#![cfg_attr(rustfmt, rustfmt::skip)]

#![allow(box_pointers)]
#![allow(dead_code)]
#![allow(missing_docs)]
#![allow(non_camel_case_types)]
#![allow(non_snake_case)]
#![allow(non_upper_case_globals)]
#![allow(trivial_casts)]
#![allow(unused_results)]
#![allow(unused_mut)]

//! Generated file from `arp.proto`

/// Generated files are compatible only with the same version
/// of protobuf runtime.
const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_3_3_0;

// @@protoc_insertion_point(message:net.ArpInfo)
#[derive(PartialEq,Clone,Default,Debug)]
pub struct ArpInfo {
    // message fields
    // @@protoc_insertion_point(field:net.ArpInfo.ip_address)
    pub ip_address: ::std::string::String,
    // @@protoc_insertion_point(field:net.ArpInfo.hw_type)
    pub hw_type: u32,
    // @@protoc_insertion_point(field:net.ArpInfo.flags)
    pub flags: ::std::string::String,
    // @@protoc_insertion_point(field:net.ArpInfo.device)
    pub device: ::std::string::String,
    // special fields
    // @@protoc_insertion_point(special_field:net.ArpInfo.special_fields)
    pub special_fields: ::protobuf::SpecialFields,
}

impl<'a> ::std::default::Default for &'a ArpInfo {
    fn default() -> &'a ArpInfo {
        <ArpInfo as ::protobuf::Message>::default_instance()
    }
}

impl ArpInfo {
    pub fn new() -> ArpInfo {
        ::std::default::Default::default()
    }

    fn generated_message_descriptor_data() -> ::protobuf::reflect::GeneratedMessageDescriptorData {
        let mut fields = ::std::vec::Vec::with_capacity(4);
        let mut oneofs = ::std::vec::Vec::with_capacity(0);
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "ip_address",
            |m: &ArpInfo| { &m.ip_address },
            |m: &mut ArpInfo| { &mut m.ip_address },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "hw_type",
            |m: &ArpInfo| { &m.hw_type },
            |m: &mut ArpInfo| { &mut m.hw_type },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "flags",
            |m: &ArpInfo| { &m.flags },
            |m: &mut ArpInfo| { &mut m.flags },
        ));
        fields.push(::protobuf::reflect::rt::v2::make_simpler_field_accessor::<_, _>(
            "device",
            |m: &ArpInfo| { &m.device },
            |m: &mut ArpInfo| { &mut m.device },
        ));
        ::protobuf::reflect::GeneratedMessageDescriptorData::new_2::<ArpInfo>(
            "ArpInfo",
            fields,
            oneofs,
        )
    }
}

impl ::protobuf::Message for ArpInfo {
    const NAME: &'static str = "ArpInfo";

    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::Result<()> {
        while let Some(tag) = is.read_raw_tag_or_eof()? {
            match tag {
                10 => {
                    self.ip_address = is.read_string()?;
                },
                16 => {
                    self.hw_type = is.read_uint32()?;
                },
                26 => {
                    self.flags = is.read_string()?;
                },
                34 => {
                    self.device = is.read_string()?;
                },
                tag => {
                    ::protobuf::rt::read_unknown_or_skip_group(tag, is, self.special_fields.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u64 {
        let mut my_size = 0;
        if !self.ip_address.is_empty() {
            my_size += ::protobuf::rt::string_size(1, &self.ip_address);
        }
        if self.hw_type != 0 {
            my_size += ::protobuf::rt::uint32_size(2, self.hw_type);
        }
        if !self.flags.is_empty() {
            my_size += ::protobuf::rt::string_size(3, &self.flags);
        }
        if !self.device.is_empty() {
            my_size += ::protobuf::rt::string_size(4, &self.device);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.special_fields.unknown_fields());
        self.special_fields.cached_size().set(my_size as u32);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::Result<()> {
        if !self.ip_address.is_empty() {
            os.write_string(1, &self.ip_address)?;
        }
        if self.hw_type != 0 {
            os.write_uint32(2, self.hw_type)?;
        }
        if !self.flags.is_empty() {
            os.write_string(3, &self.flags)?;
        }
        if !self.device.is_empty() {
            os.write_string(4, &self.device)?;
        }
        os.write_unknown_fields(self.special_fields.unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn special_fields(&self) -> &::protobuf::SpecialFields {
        &self.special_fields
    }

    fn mut_special_fields(&mut self) -> &mut ::protobuf::SpecialFields {
        &mut self.special_fields
    }

    fn new() -> ArpInfo {
        ArpInfo::new()
    }

    fn clear(&mut self) {
        self.ip_address.clear();
        self.hw_type = 0;
        self.flags.clear();
        self.device.clear();
        self.special_fields.clear();
    }

    fn default_instance() -> &'static ArpInfo {
        static instance: ArpInfo = ArpInfo {
            ip_address: ::std::string::String::new(),
            hw_type: 0,
            flags: ::std::string::String::new(),
            device: ::std::string::String::new(),
            special_fields: ::protobuf::SpecialFields::new(),
        };
        &instance
    }
}

impl ::protobuf::MessageFull for ArpInfo {
    fn descriptor() -> ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::Lazy::new();
        descriptor.get(|| file_descriptor().message_by_package_relative_name("ArpInfo").unwrap()).clone()
    }
}

impl ::std::fmt::Display for ArpInfo {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for ArpInfo {
    type RuntimeType = ::protobuf::reflect::rt::RuntimeTypeMessage<Self>;
}

// @@protoc_insertion_point(message:net.ArpInfos)
#[derive(PartialEq,Clone,Default,Debug)]
pub struct ArpInfos {
    // message fields
    // @@protoc_insertion_point(field:net.ArpInfos.infos)
    pub infos: ::std::vec::Vec<ArpInfo>,
    // special fields
    // @@protoc_insertion_point(special_field:net.ArpInfos.special_fields)
    pub special_fields: ::protobuf::SpecialFields,
}

impl<'a> ::std::default::Default for &'a ArpInfos {
    fn default() -> &'a ArpInfos {
        <ArpInfos as ::protobuf::Message>::default_instance()
    }
}

impl ArpInfos {
    pub fn new() -> ArpInfos {
        ::std::default::Default::default()
    }

    fn generated_message_descriptor_data() -> ::protobuf::reflect::GeneratedMessageDescriptorData {
        let mut fields = ::std::vec::Vec::with_capacity(1);
        let mut oneofs = ::std::vec::Vec::with_capacity(0);
        fields.push(::protobuf::reflect::rt::v2::make_vec_simpler_accessor::<_, _>(
            "infos",
            |m: &ArpInfos| { &m.infos },
            |m: &mut ArpInfos| { &mut m.infos },
        ));
        ::protobuf::reflect::GeneratedMessageDescriptorData::new_2::<ArpInfos>(
            "ArpInfos",
            fields,
            oneofs,
        )
    }
}

impl ::protobuf::Message for ArpInfos {
    const NAME: &'static str = "ArpInfos";

    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::Result<()> {
        while let Some(tag) = is.read_raw_tag_or_eof()? {
            match tag {
                10 => {
                    self.infos.push(is.read_message()?);
                },
                tag => {
                    ::protobuf::rt::read_unknown_or_skip_group(tag, is, self.special_fields.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u64 {
        let mut my_size = 0;
        for value in &self.infos {
            let len = value.compute_size();
            my_size += 1 + ::protobuf::rt::compute_raw_varint64_size(len) + len;
        };
        my_size += ::protobuf::rt::unknown_fields_size(self.special_fields.unknown_fields());
        self.special_fields.cached_size().set(my_size as u32);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::Result<()> {
        for v in &self.infos {
            ::protobuf::rt::write_message_field_with_cached_size(1, v, os)?;
        };
        os.write_unknown_fields(self.special_fields.unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn special_fields(&self) -> &::protobuf::SpecialFields {
        &self.special_fields
    }

    fn mut_special_fields(&mut self) -> &mut ::protobuf::SpecialFields {
        &mut self.special_fields
    }

    fn new() -> ArpInfos {
        ArpInfos::new()
    }

    fn clear(&mut self) {
        self.infos.clear();
        self.special_fields.clear();
    }

    fn default_instance() -> &'static ArpInfos {
        static instance: ArpInfos = ArpInfos {
            infos: ::std::vec::Vec::new(),
            special_fields: ::protobuf::SpecialFields::new(),
        };
        &instance
    }
}

impl ::protobuf::MessageFull for ArpInfos {
    fn descriptor() -> ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::Lazy::new();
        descriptor.get(|| file_descriptor().message_by_package_relative_name("ArpInfos").unwrap()).clone()
    }
}

impl ::std::fmt::Display for ArpInfos {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for ArpInfos {
    type RuntimeType = ::protobuf::reflect::rt::RuntimeTypeMessage<Self>;
}

static file_descriptor_proto_data: &'static [u8] = b"\
    \n\tarp.proto\x12\x03net\"o\n\x07ArpInfo\x12\x1d\n\nip_address\x18\x01\
    \x20\x01(\tR\tipAddress\x12\x17\n\x07hw_type\x18\x02\x20\x01(\rR\x06hwTy\
    pe\x12\x14\n\x05flags\x18\x03\x20\x01(\tR\x05flags\x12\x16\n\x06device\
    \x18\x04\x20\x01(\tR\x06device\".\n\x08ArpInfos\x12\"\n\x05infos\x18\x01\
    \x20\x03(\x0b2\x0c.net.ArpInfoR\x05infosb\x06proto3\
";

/// `FileDescriptorProto` object which was a source for this generated file
fn file_descriptor_proto() -> &'static ::protobuf::descriptor::FileDescriptorProto {
    static file_descriptor_proto_lazy: ::protobuf::rt::Lazy<::protobuf::descriptor::FileDescriptorProto> = ::protobuf::rt::Lazy::new();
    file_descriptor_proto_lazy.get(|| {
        ::protobuf::Message::parse_from_bytes(file_descriptor_proto_data).unwrap()
    })
}

/// `FileDescriptor` object which allows dynamic access to files
pub fn file_descriptor() -> &'static ::protobuf::reflect::FileDescriptor {
    static generated_file_descriptor_lazy: ::protobuf::rt::Lazy<::protobuf::reflect::GeneratedFileDescriptor> = ::protobuf::rt::Lazy::new();
    static file_descriptor: ::protobuf::rt::Lazy<::protobuf::reflect::FileDescriptor> = ::protobuf::rt::Lazy::new();
    file_descriptor.get(|| {
        let generated_file_descriptor = generated_file_descriptor_lazy.get(|| {
            let mut deps = ::std::vec::Vec::with_capacity(0);
            let mut messages = ::std::vec::Vec::with_capacity(2);
            messages.push(ArpInfo::generated_message_descriptor_data());
            messages.push(ArpInfos::generated_message_descriptor_data());
            let mut enums = ::std::vec::Vec::with_capacity(0);
            ::protobuf::reflect::GeneratedFileDescriptor::new_generated(
                file_descriptor_proto(),
                deps,
                messages,
                enums,
            )
        });
        ::protobuf::reflect::FileDescriptor::new_generated_2(generated_file_descriptor)
    })
}
