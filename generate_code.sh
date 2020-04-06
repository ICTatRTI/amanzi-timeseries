# protoc -I=./proto -I=./proto/vendor/ --python_out=python --java_out=java --go_out=go ./proto/amanzi.proto
protoc -I=proto -I=proto/vendor/ --python_out=python --java_out=java --go_out=go proto/amanzi.proto
#git add -all
#git commit -m "updated generated code"