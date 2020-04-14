# protoc -I=./proto -I=./proto/vendor/ --python_out=python --java_out=java --go_out=go-amanzi ./proto/amanzi.proto
# This script doesn't work but pasting the code into the terminal seems to work.
protoc -I=protobuf -I=protobuf/vendor/ --go_out=go-amanzi/ptypes  protobuf/amanzi.proto
#git add -all
#git commit -m "updated generated code"