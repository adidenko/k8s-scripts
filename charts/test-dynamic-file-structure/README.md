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
  name: core
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
  name: folderone
spec:
  title: FolderOne
  parentFolderRef: core
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
---
# Source: test-dynamic-file-structure/templates/grafana-folders.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaFolder
metadata:
  name: foldertwo
spec:
  title: FolderTwo
  parentFolderRef: core
  instanceSelector:
    matchLabels:
      grafana-instance: "default"
---
# Source: test-dynamic-file-structure/templates/grafana-folders.yaml
apiVersion: grafana.integreatly.org/v1beta1
kind: GrafanaFolder
metadata:
  name: services
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
  name: folderthree
spec:
  title: FolderThree
  parentFolderRef: services
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
  folderRef: folderone
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
  folderRef: folderone
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
  folderRef: foldertwo
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
  folderRef: folderthree
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
  folderRef: folderthree
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
