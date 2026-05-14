# provider-jfrog-project

`provider-jfrog-project` is a [Crossplane](https://crossplane.io/) provider built with [Upjet](https://github.com/crossplane/upjet) that exposes XRM-conformant managed resources for the [JFrog Project](https://jfrog.com/help/r/jfrog-platform-administration-documentation/projects) API.

It is generated from the [`jfrog/project`](https://registry.terraform.io/providers/jfrog/project/latest) Terraform provider (v1.9.6) and covers all resources **except** the top-level `project` resource.

## Managed Resources

Both a **cluster-scoped** and a **namespace-scoped** variant of the provider are included. Each exposes the same seven managed resource kinds:

| Kind | Terraform Resource |
|---|---|
| `Environment` | `project_environment` |
| `Group` | `project_group` |
| `Repository` | `project_repository` |
| `Role` | `project_role` |
| `ShareRepository` | `project_share_repository` |
| `ShareRepositoryWithAll` | `project_share_repository_with_all` |
| `User` | `project_user` |

API groups:

- Cluster-scoped: `*.jfrogproject.jfrog.com`
- Namespace-scoped: `*.jfrogproject.m.jfrog.com`

## Installation

Install the provider package from GHCR:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-jfrog-project
spec:
  package: ghcr.io/mikegio27/provider-jfrog-project:<version>
```

Or apply the example manifest directly:

```console
kubectl apply -f examples/install.yaml
```

## Configuration

Create a Kubernetes Secret with your JFrog credentials, then create a `ProviderConfig` that references it.

**Secret:**

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: jfrog-credentials
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "access_token": "<your-jfrog-access-token>",
      "url": "https://<your-platform>.jfrog.io"
    }
```

**ProviderConfig (cluster-scoped):**

```yaml
apiVersion: jfrogproject.jfrog.com/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: jfrog-credentials
      key: credentials
```

See [examples/cluster/providerconfig/](examples/cluster/providerconfig/) and [examples/namespaced/providerconfig/](examples/namespaced/providerconfig/) for full examples.

## Developing

### Prerequisites

- Go 1.24+
- Terraform 1.5.7 (versions ≥ 1.6 are BSL-licensed and excluded)
- A Kubernetes cluster (or `make dev-env` to create a local KIND cluster)

### Code Generation

To regenerate types and controllers from the Terraform provider schema:

```console
go run cmd/generator/main.go "$PWD"
```

### Running Locally

Run the provider against a Kubernetes cluster (uses `~/.kube/config`):

```console
make run
```

### Build

Build the provider binary:

```console
make build
```

Build and push the provider image and package, then install into the cluster:

```console
make all
```

### Publish

The [publish-provider-package](.github/workflows/publish-provider-package.yml) GitHub Actions workflow builds multi-arch images (linux/amd64, linux/arm64) and pushes the xpkg to GHCR. Trigger it via `workflow_dispatch` in the GitHub UI.

## Report a Bug

Open an [issue](https://github.com/mikegio27/provider-jfrog-project/issues) for bugs, improvements, or feature requests.
