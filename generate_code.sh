protoc --python_out=python --java_out=java protobuf/**/*.proto
git add --all
git commit -m "updated generated code"