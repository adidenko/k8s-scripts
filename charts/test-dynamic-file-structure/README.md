# test-dynamic-file-structure

Helm chart to test building `GrafanaFolder` and `GrafanaDashboard` dynamically based on the actual
file and directory structure with plain Grafana JSON files.

## Example structure

```
files/
└── dashboards
    └── Production
        ├── Core
        │   ├── FolderOne
        │   │   ├── one-one.json
        │   │   └── one-two.json
        │   └── FolderTwo
        │       └── two-one.json
        └── Services
            └── FolderThree
                ├── three-one.json
                └── three-two.json
```

## Generated GrafanaFolder

```yaml
# Source: test-dynamic-file-structure/templates/grafana-folders.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaFolder
metadata:
  name: production
spec:
  title: Production
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
---
# Source: test-dynamic-file-structure/templates/grafana-folders.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaFolder
metadata:
  name: production-core
spec:
  title: Core
  parentFolderRef: production
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
---
# Source: test-dynamic-file-structure/templates/grafana-folders.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaFolder
metadata:
  name: production-core-folderone
spec:
  title: FolderOne
  parentFolderRef: production-core
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
---
# Source: test-dynamic-file-structure/templates/grafana-folders.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaFolder
metadata:
  name: production-core-foldertwo
spec:
  title: FolderTwo
  parentFolderRef: production-core
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
---
# Source: test-dynamic-file-structure/templates/grafana-folders.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaFolder
metadata:
  name: production-services
spec:
  title: Services
  parentFolderRef: production
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
---
# Source: test-dynamic-file-structure/templates/grafana-folders.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaFolder
metadata:
  name: production-services-folderthree
spec:
  title: FolderThree
  parentFolderRef: production-services
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
```

## Generated GrafanaDashboard

```yaml
---
# Source: test-dynamic-file-structure/templates/grafana-dashboards.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaDashboard
metadata:
  name: one-one
spec:
  folderRef: production-core-folderone
  resyncPeriod: 60s
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
  json: >
    {

      "foo": "bar",
      "name": "something"
    }
---
# Source: test-dynamic-file-structure/templates/grafana-dashboards.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaDashboard
metadata:
  name: one-two
spec:
  folderRef: production-core-folderone
  resyncPeriod: 60s
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
  json: >
    {

      "foo": "bar",
      "name": "something"
    }
---
# Source: test-dynamic-file-structure/templates/grafana-dashboards.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaDashboard
metadata:
  name: two-one
spec:
  folderRef: production-core-foldertwo
  resyncPeriod: 60s
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
  json: >
    {

      "foo": "bar",
      "name": "something"
    }
---
# Source: test-dynamic-file-structure/templates/grafana-dashboards.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaDashboard
metadata:
  name: three-one
spec:
  folderRef: production-services-folderthree
  resyncPeriod: 60s
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
  json: >
    {

      "foo": "bar",
      "name": "something"
    }
---
# Source: test-dynamic-file-structure/templates/grafana-dashboards.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaDashboard
metadata:
  name: three-two
spec:
  folderRef: production-services-folderthree
  resyncPeriod: 60s
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
  json: >
    {

      "foo": "bar",
      "name": "something"
    }
```
