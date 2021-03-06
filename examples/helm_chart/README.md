# Helm chart - DEPRECATED

[helm chart repo](https://github.com/armosec/armo-helm)


## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| configMap | object | `{"create":true,"params":{"clusterName":"<MyK8sClusterName>","customerGUID":"<MyGUID>,"}}` | ARMO customer information |
| fullnameOverride | string | `""` |  |
| image | object | `{"imageName":"kubescape","pullPolicy":"IfNotPresent","repository":"quay.io/armosec","tag":"latest"}` | Image and version to deploy |
| imagePullSecrets | list | `[]` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| resources | object | `{"limits":{"cpu":"500m","memory":"512Mi"},"requests":{"cpu":"200m","memory":"256Mi"}}` | Default resources for running the service in cluster |
| schedule | string | `"0 0 * * *"` | Frequency of running the scan |
| securityContext | object | `{}` |  |
| serviceAccount | object | `{"annotations":{},"create":true,"name":"kubescape-discovery"}` | Service account that runs the scan and has permissions to view the cluster |
| tolerations | list | `[]` |  |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.5.0](https://github.com/norwoodj/helm-docs/releases/v1.5.0)
