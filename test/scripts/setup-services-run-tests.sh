#!/bin/bash

# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This shell script is used to create docker images to run the metadata service.

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace

CLUSTER_NAME="${CLUSTER_NAME}"
ZONE="${GCP_ZONE}"
PROJECT="${GCP_PROJECT}"
REGISTRY="${GCP_REGISTRY}"
NAMESPACE="${DEPLOY_NAMESPACE}"
VERSION=$(git describe --tags --always --dirty)
VERSION=${VERSION/%?/}
VERSION=750464 ### DELETE ###

echo "Activating service-account"
gcloud auth activate-service-account --key-file=${GOOGLE_APPLICATION_CREDENTIALS}

echo "Configuring kubectl"

echo "CLUSTER_NAME: ${CLUSTER_NAME}"
echo "ZONE: ${GCP_ZONE}"
echo "PROJECT: ${GCP_PROJECT}"

apt-get update
apt-get -y install software-properties-common python-software-properties
add-apt-repository ppa:jonathonf/python-3.6
apt-get update
apt-get -y install python3.6 python3-venv

gcloud --project ${PROJECT} container clusters get-credentials ${CLUSTER_NAME} \
  --zone ${ZONE}
kubectl config set-context $(kubectl config current-context) --namespace=default
USER=`gcloud config get-value account`

kubectl apply -f - << EOF
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cluster-admins
subjects:
- kind: User
  name: $USER
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: ""
EOF

kubectl create namespace $NAMESPACE

echo "REGISTRY ${REGISTRY}"
echo "REPO_NAME ${REPO_NAME}"
echo "VERSION ${VERSION}"

cd "${MANIFESTS_DIR}"

sed -i -e "s@image: gcr.io\/kubeflow-images-public\/metadata:.*@image: ${GCP_REGISTRY}\/${REPO_NAME}\/metadata:${VERSION}@" metadata/base/metadata-deployment.yaml
sed -i -e "s@--mysql_service_host=metadata-db.kubeflow@--mysql_service_host=metadata-db.${NAMESPACE}@" metadata/base/metadata-deployment.yaml

cat metadata/base/metadata-deployment.yaml

cd metadata/base

kustomize build . | kubectl apply -n $NAMESPACE -f -

TIMEOUT=120
PODNUM=$(kubectl get deployment metadata-deployment -n $NAMESPACE -o jsonpath={.spec.replicas})
echo "Expect to have $PODNUM pods of metadata-deployment."
until [[ $(kubectl get pods -n $NAMESPACE | grep "1/1" | grep metadata-deployment | wc -l) -eq $PODNUM ]]
do
    echo Pod Status $(kubectl get pods -n $NAMESPACE | grep metadata-deployment)
    sleep 10
    TIMEOUT=$(( TIMEOUT - 1 ))
    if [[ $TIMEOUT -eq 0 ]];then
        echo "FATAL: Pods of metadata-deployment are not ready after $TIMEOUT seconds!"
        kubectl get pods -n $NAMESPACE
        exit 1
    fi
done

TIMEOUT=120
PODNUM=$(kubectl get deployment metadata-db -n $NAMESPACE -o jsonpath={.spec.replicas})
echo "Expect to have $PODNUM pods of metadata-db."
until [[ $(kubectl get pods -n $NAMESPACE | grep "1/1" | grep metadata-db | wc -l) -eq $PODNUM ]]
do
    echo Pod Status $(kubectl get pods -n $NAMESPACE | grep metadata-db)
    sleep 10
    TIMEOUT=$(( TIMEOUT - 1 ))
    if [[ $TIMEOUT -eq 0 ]];then
        echo "FATAL: Pods of metadata-db are not ready after $TIMEOUT seconds!"
        kubectl get pods -n $NAMESPACE
        exit 1
    fi
done

kubectl version
kubectl -n $NAMESPACE get deploy
kubectl -n $NAMESPACE get svc
kubectl -n $NAMESPACE get pod

# Port forwarding
TARGET_POD=$(kubectl -n $NAMESPACE get pod -o=name | grep metadata-deployment | head -1)
echo "kubectl port-forward from $TARGET_POD"
kubectl -n $NAMESPACE port-forward $TARGET_POD 8080:8080 &

# Stream server logs.
kubectl -n $NAMESPACE logs -f $TARGET_POD &

# Wait at most 20 minutes for the server to be ready.
TIMEOUT=240
until curl -H "ContentType: application/json" localhost:8080/api/v1alpha1/artifact_types || [ $TIMEOUT -eq 0 ]; do
    echo "Server is not up. $TIMEOUT"
    sleep 5
    TIMEOUT=$(( TIMEOUT - 1 ))
done

# Run CURL tests
cd "${SRC_DIR}/test/e2e" && bash make_requests.sh

cd "${SRC_DIR}/sdk/python"

rm -rf .testing-env
python3.6 -m venv .testing-env
source .testing-env/bin/activate
python3 -V
# Run Python tests
bash tests/run_tests.sh

# Test demo notebook
pip3 install jupyterlab
pip3 install nbconvert
pip3 install pandas
sed -i -e "s@metadata-grpc-service.kubeflow@127.0.0.1@" sample/demo.ipynb && \
sed -i -e "s@grpc_port=8080@grpc_port=8081@" sample/demo.ipynb && \
sed -i -e "s@pip install kubeflow-metadata --user@pip install -e ${SRC_DIR}/sdk/python@" sample/demo.ipynb && \
  python3 -m nbconvert --to notebook --execute sample/demo.ipynb

cd "${SRC_DIR}"

# Test resource watcher
# Port forwarding
TARGET_GRPC_POD=$(kubectl -n $NAMESPACE get pod -o=name | grep metadata-grpc-deployment | head -1)
echo "kubectl port-forward from $TARGET_GRPC_POD"
kubectl -n $NAMESPACE port-forward $TARGET_GRPC_POD 8081:8080 &
# Stream server logs.
kubectl -n $NAMESPACE logs -f $TARGET_GRPC_POD &

cd "${SRC_DIR}/watcher" && \
  go build -o main/main main/main.go && \
  timeout --preserve-status 30 ./main/main -kubeconfig=${HOME}/.kube/config -metadata_service=localhost:8081 -resourcelist=dockerfiles/resource_list.json
exit 0
