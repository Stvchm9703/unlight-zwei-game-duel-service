fn main()  {
    // tonic_build::configure()
    //     .build_server(true)
    //     .build_client(false)
    //     .out_dir("src/proto")
    //     .compile(
    //         &[
    //             "proto/common.proto" , 
    //             "proto/EventHookPhase.proto",
    //             "proto/Data.proto",
    //             "proto/message.proto",
    //             "proto/service.proto",
    //         ],
    //         &["proto"] 
    //     );
    tonic_build::compile_protos("proto/common.proto").unwrap();
    tonic_build::compile_protos("proto/EventHookPhase.proto").unwrap();
    tonic_build::compile_protos("proto/Data.proto").unwrap();
    tonic_build::compile_protos("proto/message.proto").unwrap();
    tonic_build::compile_protos("proto/service.proto").unwrap();
}
