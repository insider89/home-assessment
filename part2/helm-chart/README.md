# s3listobjects

`s3listobjects` is simple go utility which return objects in s3 bucket.

This chart bootstraps a `s3listobjects` deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+
- Helm 3+

## Install Chart

```console
# Helm
$ helm install [RELEASE_NAME] helm-chart/
```

_See [helm install](https://helm.sh/docs/helm/helm_install/) for command documentation._


## Uninstall Chart

```console
# Helm
$ helm uninstall [RELEASE_NAME]
```

This removes all the Kubernetes components associated with the chart and deletes the release.

_See [helm uninstall](https://helm.sh/docs/helm/helm_uninstall/) for command documentation._

## Upgrading Chart

```console
# Helm
$ helm upgrade [RELEASE_NAME] [CHART] --install
```

_See [helm upgrade](https://helm.sh/docs/helm/helm_upgrade/) for command documentation._

### ConfigMap File

S3listobjects is configured through configmap. Configmap mounted to deployment pod as env variables. To provide configmap values, edit `config` variable from `values.yaml`. For example you want to provide AWS_REGION and secrets for s3listobjects:
```yaml
config:
  AWS_REGION: "us-east-2"
  AWS_ACCESS_KEY_ID: "access_key"
  AWS_SECRET_ACCESS_KEY: "secret_access_key"
```

You can skip to provide aws secrets and use annotations to attach IAM role to the pod.
Deployment check sha256 sum of configmap and automatically restart pod in case of changes.

### Ingress TLS

If your cluster allows automatic creation/retrieval of TLS certificates (e.g. [cert-manager](https://github.com/jetstack/cert-manager)), please refer to the documentation for that mechanism.

To manually configure TLS, first create/retrieve a key & certificate pair for the address(es) you wish to protect. Then create a TLS secret in the namespace:

```console
kubectl create secret tls s3listobjects-tls --cert=path/to/tls.cert --key=path/to/tls.key
```

Include the secret's name, along with the desired hostnames in `values.yaml` file:

```yaml
ingress:
  enabled: true
  hosts:
    - s3listobjects.example.com
  tls:
    - secretName: s3listobjects-tls
      hosts:
        - s3listobjects.example.com
```
