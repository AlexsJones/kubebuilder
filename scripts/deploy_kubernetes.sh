#!/bin/bash

set +xe
ENVIRONMENT="build"
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
     kubectl create secret generic auth --from-file=k8s/auth.json --namespace=kubebuilder
   else
     kubectl delete secret/auth --namespace=kubebuilder
     echo "Writing over exiting secret..."
     kubectl create secret generic auth --from-file=k8s/auth.json --namespace=kubebuilder
   fi

   # Do generation
   rm -rf deployment 2>/dev/null
   mkdir -p deployment

   for d in ./templates/*; do
    filename=$(dirname $d)
    foldername=`echo basename $filename`
    folderaltered=$(echo $foldername | sed 's/templates/deployment/g')
    mkdir -p deployment/$folderaltered

  newpath=$(echo $d | sed 's/templates/deployment/g')
   echo "Creating $newpath"
   vortex --template $d --output $newpath -varpath ./k8s/environments/$ENVIRONMENT.yaml
   done

}

if [ -z "$1" ]
  then
    echo "No argument supplied"
fi


while true; do
    read -p "Requires GOOGLE_APPLICATION_CREDENTIALS.json renamed to auth.json and placed in the k8s folder. [y/n]" yn
    case $yn in
        [Yy]* ) deploy ; break;;
        [Nn]* ) exit;;
        * ) echo "Please answer yes or no.";;
    esac
done
