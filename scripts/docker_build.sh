version="v1"
pushd files
docker build --no-cache=true -t kubebuilder:$version .
popd
