# Kubernetes-Secrets-Store-CSI-Driver

[![Build status](https://prow.k8s.io/badge.svg?jobs=secrets-store-csi-driver-e2e-vault-postsubmit)](https://testgrid.k8s.io/sig-auth-secrets-store-csi-driver#secrets-store-csi-driver-e2e-vault-postsubmit)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/kubernetes-sigs/secrets-store-csi-driver)
[![Go Report Card](https://goreportcard.com/badge/kubernetes-sigs/secrets-store-csi-driver)](https://goreportcard.com/report/kubernetes-sigs/secrets-store-csi-driver)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kubernetes-sigs/secrets-store-csi-driver)

Secrets Store CSI driver for Kubernetes secrets - Integrates secrets stores with Kubernetes via a [Container Storage Interface (CSI)](https://kubernetes-csi.github.io/docs/) volume.

The Secrets Store CSI driver `secrets-store.csi.k8s.io` allows Kubernetes to mount multiple secrets, keys, and certs stored in enterprise-grade external secrets stores into their pods as a volume. Once the Volume is attached, the data in it is mounted into the container's file system.

## Want to help?

Join us to help define the direction and implementation of this project!
- Join the [#csi-secrets-store](https://kubernetes.slack.com/messages/csi-secrets-store) channel on [Kubernetes Slack](https://kubernetes.slack.com/).
- Use [GitHub Issues](https://github.com/kubernetes-sigs/secrets-store-csi-driver/issues) to file bugs, request features, or ask questions asynchronously.
- Join [biweekly community meetings](https://docs.google.com/document/d/1q74nboAg0GSPcom3kLWCIoWg43Qg3mr306KNL58f2hg/edit?usp=sharing) to discuss development, issues, use cases, etc.

## Features

- Mounts secrets/keys/certs to pod using a CSI volume
- Supports CSI Inline volume (Kubernetes version v1.15+)
- Supports mounting multiple secrets store objects as a single volume
- Supports multiple secrets stores as providers. Multiple providers can run in the same cluster simultaneously.
- Supports pod portability with the SecretProviderClass CRD
- Supports windows containers (Kubernetes version v1.18+)
- Supports sync with Kubernetes Secrets (Secrets Store CSI Driver v0.0.10+)

#### Table of Contents

- [Kubernetes-Secrets-Store-CSI-Driver](#kubernetes-secrets-store-csi-driver)
  - [Want to help?](#want-to-help)
  - [Features](#features)
      - [Table of Contents](#table-of-contents)
  - [How It Works](#how-it-works)
  - [Demo](#demo)
  - [Usage](#usage)
    - [Prerequisites](#prerequisites)
      - [Supported kubernetes versions](#supported-kubernetes-versions)
    - [Install the Secrets Store CSI Driver](#install-the-secrets-store-csi-driver)
    - [Use the Secrets Store CSI Driver with a Provider](#use-the-secrets-store-csi-driver-with-a-provider)
    - [Create your own SecretProviderClass Object](#create-your-own-secretproviderclass-object)
    - [Update your Deployment Yaml](#update-your-deployment-yaml)
    - [Secret Content is Mounted on Pod Start](#secret-content-is-mounted-on-pod-start)
    - [[OPTIONAL] Sync with Kubernetes Secrets](#optional-sync-with-kubernetes-secrets)
    - [[OPTIONAL] Set ENV VAR](#optional-set-env-var)
    - [[OPTIONAL] Enable Auto Rotation of Secrets](#optional-enable-auto-rotation-of-secrets)
  - [Providers](#providers)
    - [Criteria for Supported Providers](#criteria-for-supported-providers)
    - [Removal from Supported Providers](#removal-from-supported-providers)
  - [Testing](#testing)
    - [Unit Tests](#unit-tests)
    - [End-to-end Tests](#end-to-end-tests)
  - [Known Limitations](#known-limitations)
  - [Troubleshooting](#troubleshooting)
  - [Code of conduct](#code-of-conduct)

## How It Works

The diagram below illustrates how Secrets Store CSI Volume works.

![diagram](img/diagram.png)

## Demo

![Secrets Store CSI Driver Demo](img/demo.gif "Secrets Store CSI Driver Azure Key Vault Provider Demo")

## Usage

### Prerequisites

#### Supported kubernetes versions

Recommended Kubernetes version: v1.16.0+

> NOTE: The CSI Inline Volume feature was introduced in Kubernetes v1.15.x. Version 1.15.x will require the `CSIInlineVolume` feature gate to be updated in the cluster. Version 1.16+ does not require any feature gate.

<details>
<summary><strong> For v1.15.x, update CSI Inline Volume feature gate </strong></summary>

The CSI Inline Volume feature was introduced in Kubernetes v1.15.x. We need to make the following updates to include the `CSIInlineVolume` feature gate:

- Update the API Server manifest to append the following feature gate:

```yaml
--feature-gates=CSIInlineVolume=true
```

- Update Kubelet manifest on each node to append the `CSIInlineVolume` feature gate:

```yaml
--feature-gates=CSIInlineVolume=true
```
</details>

### Install the Secrets Store CSI Driver

**Using Helm Chart**

Follow the [guide to install driver using Helm](charts/secrets-store-csi-driver/README.md)


<details>
<summary><strong>[ALTERNATIVE DEPLOYMENT OPTION] Using Deployment Yamls</strong></summary>

```bash
kubectl apply -f deploy/rbac-secretproviderclass.yaml # update the namespace of the secrets-store-csi-driver ServiceAccount
kubectl apply -f deploy/csidriver.yaml
kubectl apply -f deploy/secrets-store.csi.x-k8s.io_secretproviderclasses.yaml
kubectl apply -f deploy/secrets-store.csi.x-k8s.io_secretproviderclasspodstatuses.yaml
kubectl apply -f deploy/secrets-store-csi-driver.yaml --namespace $NAMESPACE

# If using the driver to sync secrets-store content as Kubernetes Secrets, deploy the additional RBAC permissions
# required to enable this feature
kubectl apply -f deploy/rbac-secretprovidersyncing.yaml

# [OPTIONAL] For kubernetes version < 1.16 running `kubectl apply -f deploy/csidriver.yaml` will fail. To install the driver run
kubectl apply -f deploy/csidriver-1.15.yaml

# [OPTIONAL] To deploy driver on windows nodes
kubectl apply -f deploy/secrets-store-csi-driver-windows.yaml --namespace $NAMESPACE
```

To validate the installer is running as expected, run the following commands:

```bash
kubectl get po --namespace $NAMESPACE
```

You should see the Secrets Store CSI driver pods running on each agent node:

```bash
csi-secrets-store-qp9r8         3/3     Running   0          4m
csi-secrets-store-zrjt2         3/3     Running   0          4m
```

You should see the following CRDs deployed:

```bash
kubectl get crd
NAME                                               
secretproviderclasses.secrets-store.csi.x-k8s.io    
```

</details>

### Use the Secrets Store CSI Driver with a Provider

Select a provider from the following list, then follow the installation steps for the provider:
-  [Azure Provider](https://github.com/Azure/secrets-store-csi-driver-provider-azure#usage)
-  [Vault Provider](https://github.com/hashicorp/secrets-store-csi-driver-provider-vault)

### Create your own SecretProviderClass Object

To use the Secrets Store CSI driver, create a `SecretProviderClass` custom resource to provide driver configurations and provider-specific parameters to the CSI driver.

A `SecretProviderClass` custom resource should have the following components:
```yaml
apiVersion: secrets-store.csi.x-k8s.io/v1alpha1
kind: SecretProviderClass
metadata:
  name: my-provider
spec:
  provider: vault                             # accepted provider options: azure or vault
  parameters:                                 # provider-specific parameters
```

Here is a sample [`SecretProviderClass` custom resource](test/bats/tests/vault/vault_v1alpha1_secretproviderclass.yaml)

### Update your Deployment Yaml

To ensure your application is using the Secrets Store CSI driver, update your deployment yaml to use the `secrets-store.csi.k8s.io` driver and reference the `SecretProviderClass` resource created in the previous step.

```yaml
volumes:
  - name: secrets-store-inline
    csi:
      driver: secrets-store.csi.k8s.io
      readOnly: true
      volumeAttributes:
        secretProviderClass: "my-provider"
```

Here is a sample [deployment yaml](test/bats/tests/vault/nginx-pod-vault-inline-volume-secretproviderclass.yaml) using the Secrets Store CSI driver.

### Secret Content is Mounted on Pod Start
On pod start and restart, the driver will call the provider binary to retrieve the secret content from the external Secrets Store you have specified in the `SecretProviderClass` custom resource. Then the content will be mounted to the container's file system. 

To validate, once the pod is started, you should see the new mounted content at the volume path specified in your deployment yaml.

```bash
kubectl exec -it nginx-secrets-store-inline ls /mnt/secrets-store/
foo
```

### [OPTIONAL] Sync with Kubernetes Secrets

In some cases, you may want to create a Kubernetes Secret to mirror the mounted content. Use the optional `secretObjects` field to define the desired state of the synced Kubernetes secret objects. **The volume mount is required for the Sync With Kubernetes Secrets** 
> NOTE: If the provider supports object alias for the mounted file, then make sure the `objectName` in `secretObjects` matches the name of the mounted content. This could be the object name or the object alias.

A `SecretProviderClass` custom resource should have the following components:
```yaml
apiVersion: secrets-store.csi.x-k8s.io/v1alpha1
kind: SecretProviderClass
metadata:
  name: my-provider
spec:
  provider: vault                             # accepted provider options: azure or vault
  secretObjects:                              # [OPTIONAL] SecretObject defines the desired state of synced K8s secret objects
  - data:
    - key: username                           # data field to populate
      objectName: foo1                        # name of the mounted content to sync. this could be the object name or the object alias
    secretName: foosecret                     # name of the Kubernetes Secret object
    type: Opaque                              # type of the Kubernetes Secret object e.g. Opaque, kubernetes.io/tls
```
> NOTE: Here is the list of supported Kubernetes Secret types: `Opaque`, `kubernetes.io/basic-auth`, `bootstrap.kubernetes.io/token`, `kubernetes.io/dockerconfigjson`, `kubernetes.io/dockercfg`, `kubernetes.io/ssh-auth`, `kubernetes.io/service-account-token`, `kubernetes.io/tls`.  

Here is a sample [`SecretProviderClass` custom resource](test/bats/tests/vault/vault_synck8s_v1alpha1_secretproviderclass.yaml) that syncs Kubernetes secrets.

### [OPTIONAL] Set ENV VAR

Once the secret is created, you may wish to set an ENV VAR in your deployment to reference the new Kubernetes secret.

```yaml
spec:
  containers:
  - image: nginx
    name: nginx
    env:
    - name: SECRET_USERNAME
      valueFrom:
        secretKeyRef:
          name: foosecret
          key: username
```
Here is a sample [deployment yaml](test/bats/tests/vault/nginx-deployment-synck8s.yaml) that creates an ENV VAR from the synced Kubernetes secret.

### [OPTIONAL] Enable Auto Rotation of Secrets

You can setup the Secrets Store CSI Driver to periodically update the pod mount and Kubernetes Secret with the latest content from external secrets-store. Refer to [doc](docs/README.rotation.md) for steps on enabling auto rotation.

**NOTE** The CSI driver **does not restart** the application pods. It only handles updating the pod mount and Kubernetes secret similar to how Kubernetes handles updates to Kubernetes secret mounted as volumes.


## Providers

This project features a pluggable provider interface developers can implement that defines the actions of the Secrets Store CSI driver. This enables retrieval of sensitive objects stored in an enterprise-grade external secrets store into Kubernetes while continue to manage these objects outside of Kubernetes.

### Criteria for Supported Providers

Here is a list of criteria for supported provider:
1. Code audit of the provider implementation to ensure it adheres to the required provider-driver interface - [Implementing a Provider for Secrets Store CSI Driver](docs/README.new-provider.md)
2. Add provider to the e2e test suite to demonstrate it functions as expected https://github.com/kubernetes-sigs/secrets-store-csi-driver/tree/master/test/bats Please use existing providers e2e tests as a reference.
3. If any update is made by a provider (not limited to security updates), the provider is expected to update the provider's e2e test in this repo

### Removal from Supported Providers

Failure to adhere to the [Criteria for Supported Providers](#criteria-for-supported-providers) will result in the removal of the provider from the supported list and subject to another review before it can be added back to the list of supported providers.

When a provider's e2e tests are consistently failing with the latest version of the driver, the driver maintainers will coordinate with the provider maintainers to provide a fix. If the test failures are not resolved within 4 weeks, then the provider will be removed from the list of supported providers. 

## Testing

### Unit Tests

Run unit tests locally with `make test`.

### End-to-end Tests

End-to-end tests automatically runs on Prow when a PR is submitted. If you want to run using a local or remote Kubernetes cluster, make sure to have `kubectl`, `helm` and `bats` set up in your local environment and then run `make e2e-azure` or `make e2e-vault` with custom images.

Job config for test jobs run for each PR in prow can be found [here](https://github.com/kubernetes/test-infra/blob/master/config/jobs/kubernetes-sigs/secrets-store-csi-driver/secrets-store-csi-driver-config.yaml)

## Known Limitations

- [Mounted content and Kubernetes Secret not updated after secret is updated in external secrets-store](docs/README.limitations.md#mounted-content-and-kubernetes-secret-not-updated-after-secret-is-updated-in-external-secrets-store)

## Troubleshooting

- To troubleshoot issues with the csi driver, you can look at logs from the `secrets-store` container of the csi driver pod running on the same node as your application pod:
  ```bash
  kubectl get pod -o wide
  # find the secrets store csi driver pod running on the same node as your application pod

  kubectl logs csi-secrets-store-secrets-store-csi-driver-7x44t secrets-store
  ```

## Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).
