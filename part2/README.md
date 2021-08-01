# part2

`app` - directory contains golang code of simple http server which returns content of s3 bucket.

`helm-chart` - directory contains helm chart which install s3listobjects docker image to helm chart.

`Dockerfile` - dockerfile to build s3listobejects docker image.

## Prerequisites

- Kubernetes 1.16+
- Helm 3+
- Docker 20.10.7+
- Go 1.15.6+
- kind 0.11.1+
- curl 7.64.1+

## Step by step guide

1. You are on the root of pasrt2 directory:
```bash
Dockerfile
README.md
app
helm-chart
```
2. Build docker image with following command:
```bash
docker build -t s3listobjects:scratch-1.0 .
```
3. Run kind cluster with following command:
```bash
kind create cluster --name s3listobjects
```
4. Make sure you have access to cluster:
```bash
kubectl cluster-info
```
5. Upload docker image to kind cluster:
```bash
kind load docker-image s3listobjects:scratch-1.0 --name s3listobjects
```
5. Edit `values.yaml` in helm-chart directory(`config` values) to provide AWS env variables for application, eg:
```yaml
config:
  AWS_REGION: "us-east-2"
  AWS_ACCESS_KEY_ID: "access_key"
  AWS_SECRET_ACCESS_KEY: "secret_access_key"
```
6. Install helm chart:
```bash
helm install s3listobjects helm-chart/
```
7. Expose pod to local host via command:
```bash
export POD_NAME=$(kubectl get pods --namespace s3listobjects -l "app.kubernetes.io/name=s3listobjects,app.kubernetes.io/instance=s3listobjects" -o jsonpath="{.items[0].metadata.name}")
export CONTAINER_PORT=$(kubectl get pod --namespace s3listobjects $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
kubectl --namespace s3listobjects port-forward $POD_NAME 8080:$CONTAINER_PORT
```
8. In different console run `curl` command to list objects in s3 bucket:
```bash
curl localhost:8080/?bucket=bucket_name
```
Where `bucket_name` - is your bucket name.

If you want to deploy helm-chart to real kubernetes cluster, you need to upload docker image to some docker registry and provide image/tag and repository credentials in `values.yaml` helm chart file.

## Cleanup env

1. Remove kind cluster:
```bash
kind delete cluster --name s3listobjects
```
2. Delete docker image:
```bash
docker rmi s3listobjects:scratch-1.0
```