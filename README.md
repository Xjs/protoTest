# Gazelle protobuf failure demo

This repository demonstrates how gazelle fails to generate targets for multiple protobuf packages if they reside in one and the same directory.

## How to trigger

Run `gazelle -repo_root $(pwd) -go_prefix github.com/Xjs/protoTest`.

## Error message

`gazelle: found packages a () and b () in ~\src\github.com\Xjs\proto_test`

## Expected result

A `BUILD.bazel` file like mocked up in `BUILD.bazel.expected`. 

Although having used the default (hierarchical) mode here, I'm especially interested in correct results in `-experimental_flat` mode.
