# op6
[Resource](https://dev.to/mxglt/how-to-create-a-kubernetes-operator--2g6h)
[app/v1 package](https://pkg.go.dev/k8s.io/api/apps/v1)

## Description
```
err = r.Get(context.TODO(), types.NamespacedName{Name: myProxy.Spec.Name, Namespace: req.Namespace}, existingDeployments)
```
is very different from
```
err = r.Get(context.TODO(), req.Namespace, existingDeployments)
```
where the second way will keep causing
```
ERROR   Reconciler error        {"controller": "op6apikind", "controllerGroup": "op6apigroup.op6domain", "controllerKind": "Op6ApiKind", "Op6ApiKind": {"name":"op6apikind-sample","namespace":"default"}, "namespace": "default", "name": "op6apikind-sample", "reconcileID": "b211ce53-b408-48e5-bb2e-f0df6f9ba1c6", "error": "deployments.apps \"app-name\" already exists"}
```

## Getting Started

### Prerequisites
- go version v1.20.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/op6:tag
```

**NOTE:** This image ought to be published in the personal registry you specified. 
And it is required to have access to pull the image from the working environment. 
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/op6:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin 
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

