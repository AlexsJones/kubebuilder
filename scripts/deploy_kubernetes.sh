#!/bin/bash

set +xe
ENVIRONMENT="build"
REGISTRY="us.gcr.io/beamery-trials"
function deploy {
  if hash vortex 2>/dev/null; then
       echo "vortex installed..."
   else
       echo "fetching vortex..."
       go get github.com/AlexsJones/vortex
       echo "Please make sure $GOPATH/bin is on the path and re-run this script"
       exit
   fi

   kubectl describe ns kubebuilder 2>/dev/null
   if [ $? -ne 0 ]
   then
     kubectl create ns kubebuilder
     echo "Created namespace kubebuilder..."
   else
     echo "Found namespace OKAY"
   fi

   kubectl get secret/auth --namespace=kubebuilder 2>/dev/null
   if [ $? -ne 0 ]
   then
     kubectl create secret generic auth --from-file=k8s/GOOGLE_APPLICATION_CREDENTIALS.json --namespace=kubebuilder
   else
     kubectl delete secret/auth --namespace=kubebuilder
     echo "Writing over exiting secret..."
     kubectl create secret generic auth --from-file=k8s/GOOGLE_APPLICATION_CREDENTIALS.json --namespace=kubebuilder
   fi

   # Do generation
   rm -rf k8s/deployment 2>/dev/null
   mkdir -p k8s/deployment

   for d in ./k8s/template/**/*; do
    filename=$(dirname $d)
    foldername=`echo basename $filename`
    folderaltered=$(echo $foldername | sed 's/template/deployment/g')
    mkdir -p deployment/$folderaltered

   newpath=$(echo $d | sed 's/template/deployment/g')
   echo "Creating $newpath"
   vortex --template $d --output $newpath -varpath ./k8s/environments/$ENVIRONMENT.yaml
   cat $newpath
   done

   #Do deployment

   kubectl apply -f ./k8s/deployment/kubebuilder_deployment --namespace=kubebuilder
}

if [ -z "$1" ]
  then
    echo "No argument supplied"
fi


while true; do
    read -p "Requires GOOGLE_APPLICATION_CREDENTIALS.json placed in the k8s folder. [y/n]" yn
    case $yn in
        [Yy]* ) deploy ; break;;
        [Nn]* ) exit;;
        * ) echo "Please answer yes or no.";;
    esac
done
