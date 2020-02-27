fn main() {
    tonic_build::compile_protos("proto/common.proto").unwrap();
    tonic_build::compile_protos("proto/EventHookPhase.proto").unwrap();
    tonic_build::compile_protos("proto/Data.proto").unwrap();
    tonic_build::compile_protos("proto/message.proto").unwrap();
    tonic_build::compile_protos("proto/service.proto").unwrap();
}