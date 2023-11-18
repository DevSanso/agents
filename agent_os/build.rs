use std::env;
use std::path::PathBuf;

use protobuf_codegen;

use lazy_static::lazy_static;

lazy_static! {
    static ref PROJECT_PATH : String = String::from(env!("CARGO_MANIFEST_DIR"));
    static ref PROTOBUF_PATH : String = (|| {
        let mut buf = PathBuf::from(env!("CARGO_MANIFEST_DIR"));
        buf.pop();
        buf.push("protobuf/agent_os");
        let s = String::from(buf.as_os_str().to_str().unwrap());
        return s;
    })();
}

fn clone_path(str1 : &'_ str, str2 :&'_ str) -> PathBuf {
    return PathBuf::from_iter(vec![str1, str2]).clone();
}

fn agent_os_root_proto_gen(builder : &mut protobuf_codegen::Codegen) -> &mut protobuf_codegen::Codegen  {
    let protobuf_path = PROTOBUF_PATH.as_str();

    let gen  = builder.protoc()
    .include(clone_path(protobuf_path, ""))
    .input(clone_path(protobuf_path, "message.proto"))
    .input(clone_path(protobuf_path, "net.proto"))
    .cargo_out_dir("src/protos");

    gen
}

fn create_proto_file() {
    let mut root_codegen = protobuf_codegen::Codegen::new();
    let root_gen = agent_os_root_proto_gen(&mut root_codegen);

    root_gen.run_from_script();
}

fn main() {
    let project_path =PROJECT_PATH.as_str();
    env::set_var("OUT_DIR",project_path);

    create_proto_file();
    //add_pub_path_in_root_mod();
}