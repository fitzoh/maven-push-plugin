# cf maven-push plugin
Want to deploy a specific version of a maven artifact to cloud foundry but don't have a local copy? Use `cf maven-push`.

`cf maven-push` is a drop in replacement for `cf push` that:
* downloads the maven artifact to a temp file based on the manifest
* calls `cf push` with the arguments passed to `cf maven-push` (with an updated path to the artifact)

A manifest file (with a single application in the manifest) is required to use this plugin.

Does not yet support [variable substitution](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html#multi-manifests) (at least for the `maven` portion of the manifest).

### usage

`$ cf maven-push -f my-manifest.yml <any other args to pass to 'cf push'>`

### example manifest

```
---
applications:
- name: my-application
  memory: 1G
  maven:
    repo-url: https://repo.maven.apache.org/maven2/
    group-id: com.group.my
    artifact-id: my-artifact
    version: 1.0.0
```

### all manifest options

`repo-url` (required)

`group-id` (required)

`artifact-id` (required)

`version` (required)

`extension` (optional, default `.jar`)

`classifier` (optional, default none)

`repo-username` (optional, default none)

`repo-password` (optional, default none)

### installation

Download the appropriate binary from the [releases page](https://github.com/fitzoh/maven-push-plugin/releases).

`$ cf install-plugin path/to/maven-push/binary`