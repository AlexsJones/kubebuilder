version=`echo "$(cat VERSION)"`
newversion=`/bin/bash ./scripts/increment.sh $version`
cp docker/Dockerfile.tmpl Dockerfile
cp docker/run.sh.tmpl run.sh
docker build --no-cache=true -t kubebuilder:$newversion .
rm Dockerfile
rm run.sh
docker tag kubebuilder:$newversion tibbar/kubebuilder:$newversion
docker push tibbar/kubebuilder:$newversion
echo "Bumping version from $version to $newversion"
echo "$newversion" > VERSION
