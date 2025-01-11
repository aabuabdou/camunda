The changelog is automatically generated and it follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) format.
<!-- generated by git-cliff -->
### Release Info

Supported versions:

- Camunda applications: [8.2](https://github.com/camunda/camunda-platform/releases?q=tag%3A8.2&expanded=true)
- Helm values: [8.2.34](https://artifacthub.io/packages/helm/camunda/camunda-platform/8.2.34#parameters)
- Helm CLI: [3.16.3](https://github.com/helm/helm/releases/tag/v3.16.3)

Camunda images:

- docker.io/camunda/connectors-bundle:0.23.2
- docker.io/camunda/identity:8.2.31
- docker.io/camunda/operate:8.2.31
- docker.io/camunda/optimize:8.2.14
- docker.io/camunda/tasklist:8.2.32
- docker.io/camunda/zeebe:8.2.32
- registry.camunda.cloud/web-modeler-ee/modeler-restapi:8.2.20
- registry.camunda.cloud/web-modeler-ee/modeler-webapp:8.2.20
- registry.camunda.cloud/web-modeler-ee/modeler-websockets:8.2.20

Non-Camunda images:

- docker.elastic.co/elasticsearch/elasticsearch:7.17.26
- docker.io/bitnami/keycloak:19.0.3
- docker.io/bitnami/postgresql:14.5.0-debian-11-r35
- docker.io/bitnami/postgresql:15.10.0

### Verification

To verify the integrity of the Helm chart using [Cosign](https://docs.sigstore.dev/signing/quickstart/):

```shell
cosign verify-blob camunda-platform-8.2.34.tgz \
  --bundle camunda-platform-8.2.34.cosign.bundle \
  --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
  --certificate-identity "https://github.com/camunda/camunda-platform-helm/.github/workflows/chart-release-chores.yml@refs/pull/2560/merge"
```