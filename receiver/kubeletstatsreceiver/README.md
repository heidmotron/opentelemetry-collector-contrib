# Kubelet Stats Receiver

<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [beta]: metrics   |
| Distributions | [contrib], [observiq], [splunk], [sumo] |
| Issues        | [![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Areceiver%2Fkubeletstats%20&label=open&color=orange&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aopen+is%3Aissue+label%3Areceiver%2Fkubeletstats) [![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Areceiver%2Fkubeletstats%20&label=closed&color=blue&logo=opentelemetry)](https://github.com/open-telemetry/opentelemetry-collector-contrib/issues?q=is%3Aclosed+is%3Aissue+label%3Areceiver%2Fkubeletstats) |
| [Code Owners](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/CONTRIBUTING.md#becoming-a-code-owner)    | [@dmitryax](https://www.github.com/dmitryax), [@TylerHelmuth](https://www.github.com/TylerHelmuth) |

[beta]: https://github.com/open-telemetry/opentelemetry-collector#beta
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
[observiq]: https://github.com/observIQ/observiq-otel-collector
[splunk]: https://github.com/signalfx/splunk-otel-collector
[sumo]: https://github.com/SumoLogic/sumologic-otel-collector
<!-- end autogenerated section -->

The Kubelet Stats Receiver pulls node, pod, container, and volume metrics from the API server on a kubelet
and sends it down the metric pipeline for further processing.

## Metrics

Details about the metrics produced by this receiver can be found in [metadata.yaml](./metadata.yaml) with further documentation in [documentation.md](./documentation.md)

## Configuration

A kubelet runs on a kubernetes node and has an API server to which this
receiver connects. To configure this receiver, you have to tell it how
to connect and authenticate to the API server and how often to collect data
and send it to the next consumer.

Kubelet Stats Receiver supports both secure Kubelet endpoint exposed at port 10250 by default and read-only
Kubelet endpoint exposed at port 10255. If `auth_type` set to `none`, the read-only endpoint will be used. The secure 
endpoint will be used if `auth_type` set to any of the following values:

- `tls` tells the receiver to use TLS for auth and requires that the fields
`ca_file`, `key_file`, and `cert_file` also be set.
- `serviceAccount` tells this receiver to use the default service account token
to authenticate to the kubelet API.
- `kubeConfig` tells this receiver to use the kubeconfig file (KUBECONFIG env variable or ~/.kube/config)
to authenticate and use API server proxy to access the kubelet API.
- `initial_delay` (default = `1s`): defines how long this receiver waits before starting.

### TLS Example

```yaml
receivers:
  kubeletstats:
    collection_interval: 20s
    initial_delay: 1s
    auth_type: "tls"
    ca_file: "/path/to/ca.crt"
    key_file: "/path/to/apiserver.key"
    cert_file: "/path/to/apiserver.crt"
    endpoint: "https://192.168.64.1:10250"
    insecure_skip_verify: true
exporters:
  file:
    path: "fileexporter.txt"
service:
  pipelines:
    metrics:
      receivers: [kubeletstats]
      exporters: [file]
```

### Service Account Authentication Example

Although it's possible to use kubernetes' hostNetwork feature to talk to the
kubelet api from a pod, the preferred approach is to use the downward API.

Make sure the pod spec sets the node name as follows:

```yaml
env:
  - name: K8S_NODE_NAME
    valueFrom:
      fieldRef:
        fieldPath: spec.nodeName
```

Then the otel config can reference the `K8S_NODE_NAME` environment variable:

```yaml
receivers:
  kubeletstats:
    collection_interval: 20s
    auth_type: "serviceAccount"
    endpoint: "https://${env:K8S_NODE_NAME}:10250"
    insecure_skip_verify: true
exporters:
  file:
    path: "fileexporter.txt"
service:
  pipelines:
    metrics:
      receivers: [kubeletstats]
      exporters: [file]
```

Note: a missing or empty `endpoint` will cause the hostname on which the
collector is running to be used as the endpoint. If the hostNetwork flag is
set, and the collector is running in a pod, this hostname will resolve to the
node's network namespace.

### Read Only Endpoint Example

The following config can be used to collect Kubelet metrics from read-only endpoint:

```yaml
receivers:
  kubeletstats:
    collection_interval: 20s
    auth_type: "none"
    endpoint: "http://${env:K8S_NODE_NAME}:10255"
exporters:
  file:
    path: "fileexporter.txt"
service:
  pipelines:
    metrics:
      receivers: [kubeletstats]
      exporters: [file]
```

### Kubeconfig example

The following config can be used to collect Kubelet metrics from read-only endpoint, proxied by the API server:

```yaml
receivers:
  kubeletstats:
    collection_interval: 20s
    auth_type: "kubeConfig"
    context: "my-context"
    insecure_skip_verify: true
    endpoint: "${env:K8S_NODE_NAME}"
exporters:
  file:
    path: "fileexporter.txt"
service:
  pipelines:
    metrics:
      receivers: [kubeletstats]
      exporters: [file]
```
Note that using `auth_type` `kubeConfig`, the endpoint should only be the node name as the communication to the kubelet is proxied by the API server configured in the `kubeConfig`.
`insecure_skip_verify` still applies by overriding the `kubeConfig` settings.
If no `context` is specified, the current context or the default context is used.

### Extra metadata labels

By default, all produced metrics get resource labels based on what kubelet /stats/summary endpoint provides.
For some use cases it might be not enough. So it's possible to leverage other endpoints to fetch
additional metadata entities and set them as extra labels on metric resource. Currently supported metadata
include the following:

- `container.id` - to augment metrics with Container ID label obtained from container statuses exposed via `/pods`.
- `k8s.volume.type` - to collect volume type from the Pod spec exposed via `/pods` and have it as a label on volume metrics.
If there's more information available from the endpoint than just volume type, those are sycned as well depending on
the available fields and the type of volume. For example, `aws.volume.id` would be synced from `awsElasticBlockStore`
and `gcp.pd.name` is synced for `gcePersistentDisk`.

If you want to have `container.id` label added to your metrics, use `extra_metadata_labels` field to enable
it, for example:

```yaml
receivers:
  kubeletstats:
    collection_interval: 10s
    auth_type: "serviceAccount"
    endpoint: "${env:K8S_NODE_NAME}:10250"
    insecure_skip_verify: true
    extra_metadata_labels:
      - container.id
```

If `extra_metadata_labels` is not set, no additional API calls is done to fetch extra metadata.

#### Collecting Additional Volume Metadata

When dealing with Persistent Volume Claims, it is possible to optionally sync metdadata from the underlying
storage resource rather than just the volume claim. This is achieved by talking to the Kubernetes API. Below
is an example, configuration to achieve this.

```yaml
receivers:
  kubeletstats:
    collection_interval: 10s
    auth_type: "serviceAccount"
    endpoint: "${env:K8S_NODE_NAME}:10250"
    insecure_skip_verify: true
    extra_metadata_labels:
      - k8s.volume.type
    k8s_api_config:
      auth_type: serviceAccount
```

If `k8s_api_config` set, the receiver will attempt to collect metadata from underlying storage resources for
Persistent Volume Claims. For example, if a Pod is using a PVC backed by an EBS instance on AWS, the receiver
would set the `k8s.volume.type` label to be `awsElasticBlockStore` rather than `persistentVolumeClaim`.

### Metric Groups

A list of metric groups from which metrics should be collected. By default, metrics from containers,
pods and nodes will be collected. If `metric_groups` is set, only metrics from the listed groups
will be collected. Valid groups are `container`, `pod`, `node` and `volume`. For example, if you're
looking to collect only `node` and `pod` metrics from the receiver use the following configuration.

```yaml
receivers:
  kubeletstats:
    collection_interval: 10s
    auth_type: "serviceAccount"
    endpoint: "${env:K8S_NODE_NAME}:10250"
    insecure_skip_verify: true
    metric_groups:
      - node
      - pod
```

### Optional parameters

The following parameters can also be specified:

- `collection_interval` (default = `10s`): The interval at which to collect data.
- `insecure_skip_verify` (default = `false`): Whether or not to skip certificate verification.

The full list of settings exposed for this receiver are documented [here](./config.go)
with detailed sample configurations [here](./testdata/config.yaml).

### Role-based access control

The Kubelet Stats Receiver needs `get` permissions on the `nodes/stats` resources. Additionally, when using `extra_metadata_labels` or any of the `{request|limit}_utilization` metrics the processor also needs `get` permissions for `nodes/proxy` resources.

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: otel-collector
rules:
  - apiGroups: [""]
    resources: ["nodes/stats"]
    verbs: ["get"]
    
  # Only needed if you are using extra_metadata_labels or
  # are collecting the request/limit utilization metrics
  - apiGroups: [""]
    resources: ["nodes/proxy"]
    verbs: ["get"]
```