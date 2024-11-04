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
