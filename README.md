# Nginx Ingress Controller Component Operator

[![REUSE status](https://api.reuse.software/badge/github.com/SAP/nginx-ingress-controller-cop)](https://api.reuse.software/info/github.com/SAP/nginx-ingress-controller-cop)

## About this project

Component Operator for [https://github.com/kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx).

After installing this component operator into a Kubernetes cluster, nginx ingress controllers may be installed by deploying

```yaml
apiVersion: operator.kyma-project.io/v1alpha1
kind: NginxIngressController
metadata:
  name: nginx-ingress-controller
# spec:
  # optional spec attributes
```

In `spec`, all values of the [upstream Helm chart](https://github.com/kubernetes/ingress-nginx/tree/main/charts/ingress-nginx) are allowed. Caveats:
- the component operator does not perform any validation, it just passes the provided spec to the helm chart
- the supported/allowed spec format might change, when the included upstream chart changes
- deploying multiple `NginxIngressController` resources may or may not work, depending on the supplied values, i.e. it may fail if the configuration leads to deployment of clashing singleton resources such as custom resource definitions, webhook definitions, ...

In addition, the following attributes can be supplied in `spec`:
- `namespace`: target namespace for the deployed ingress controller (if not specified, the namespace of the owning `NginxIngressController` resource will be used)
- `name`: target name for the deployed ingress controller (if not specified, generated resources will be prefixed with the name of the owning `NginxIngressController` resource)
- `additionalResources`: array of additional resource manifests that will be deployed along with the ingress controller.

## Support, Feedback, Contributing

This project is open to feature requests/suggestions, bug reports etc. via [GitHub issues](https://github.com/SAP/nginx-ingress-controller-cop/issues). Contribution and feedback are encouraged and always welcome. For more information about how to contribute, the project structure, as well as additional contribution information, see our [Contribution Guidelines](CONTRIBUTING.md).

## Code of Conduct

We as members, contributors, and leaders pledge to make participation in our community a harassment-free experience for everyone. By participating in this project, you agree to abide by its [Code of Conduct](https://github.com/SAP/.github/blob/main/CODE_OF_CONDUCT.md) at all times.

## Licensing

Copyright 2023 SAP SE or an SAP affiliate company and nginx-ingress-controller-cop contributors. Please see our [LICENSE](LICENSE) for copyright and license information. Detailed information including third-party components and their licensing/copyright information is available [via the REUSE tool](https://api.reuse.software/info/github.com/SAP/nginx-ingress-controller-cop).
