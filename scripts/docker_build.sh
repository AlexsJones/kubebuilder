if [ -z "$1" ]
  then
    echo "Requires google project name"
    exit 
fi

version="v1"
google_project=$1
docker build --no-cache=true -t kubebuilder:$version .
docker tag kubebuilder:$version us.gcr.io/$google_project/kubebuilder:$version
gcloud docker -- push us.gcr.io/$google_project/kubebuilder:$version
