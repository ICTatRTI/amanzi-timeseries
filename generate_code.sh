protoc --python_out=python --java_out=java --go_out=go amanzi.proto
git add -a
git commit -m "updated generated code"