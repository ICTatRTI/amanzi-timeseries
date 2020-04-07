# protoc -I=./proto -I=./proto/vendor/ --python_out=python --java_out=java --go_out=go ./proto/amanzi.proto
# This script doesn't work but pasting the code into the terminal seems to work.
protoc -I=proto -I=proto/vendor/ --python_out=python --java_out=java --go_out=go --csharp_out=csharp proto/amanzi.proto
#git add -all
#git commit -m "updated generated code"