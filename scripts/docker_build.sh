version="v1"
pushd files
docker build --no-cache=true -t kubebuilder:$version .
popd
pushd k8s
sed -ie "s/kubebuilder:v1/kuberbuilder:$version/g" ./environments/build.yaml
popd
