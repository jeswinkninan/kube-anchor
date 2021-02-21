# Kube-Anchor

Kube-Anchor is a Kubernetes Admission Controller of type ValidatingWebhook.

It is implementing an admission controller for making sure that actions such as "CREATE", "UPDATE", "DELETE", "PATCH" against the Workload resources such as deployment, statefulset, pods etc are validated and accepted/denied gracefully.  This helps the the SREs, Devops, Kubernetes Cluster Admins in making sure that the Production Freeze is logically enabled in a cluster and doesn't need to worry about accidential resource modification(s) during a Production Freeze Window.

# Installation

Please run the following commands to deploy Kube-Anchor to your cluster

```
git clone https://github.com/JESWINKNINAN/kube-anchor.git
```
Now Run Script From Where You have cloned as follows:

```
]# sh kube-anchor/deploy/kube-anchor-install.sh 
---Generating SSL SAN Certificate for Kube-Anchor Implementations---
Generating RSA private key, 2048 bit long modulus (2 primes)
.....................................+++++
....................+++++
e is 65537 (0x010001)
Generating RSA private key, 2048 bit long modulus (2 primes)
.....+++++
......+++++
e is 65537 (0x010001)
Signature ok
subject=O = Kube-Anchor, OU = Kube-Anchor, CN = kube-anchor.kube-system.svc
Getting CA Private Key
---Creating secrets using kubectl---
secret/kube-anchor-webhook-server-tls created
---Converting CA Cert to Base64 Encoded Format and Deploying Kube-Anchor---
validatingwebhookconfiguration.admissionregistration.k8s.io/kube-anchor created
deployment.apps/kube-anchor created
service/kube-anchor created
]#
```
# Getting Started Using Kube-Anchor

Label Namespaces Which Needs Kube-Anchor

```
]# kubectl label ns test0-kube-anchor kube-anchor=enabled
namespace/test0-kube-anchor labeled
```
```
]# kubectl describe ns test0-kube-anchor
Name:         test0-kube-anchor
Labels:       kube-anchor=enabled
```

Now try to make a deployment to the namespace

```
]# kubectl apply -f https://k8s.io/examples/controllers/nginx-deployment.yaml -n test0-kube-anchor
Error from server: error when creating "https://k8s.io/examples/controllers/nginx-deployment.yaml": 
admission webhook "kube-anchor.kube-system.svc" denied the request: Cluster Freeze Window Enabled via Kube-Anchor ☸ 
```

Try deleting already running deployment in a different namespace which is labelled with kube-anchor=enabled and it got denied by Kube-Anchor

```
kubectl delete deploy nginx-deployment -n test
Error from server: admission webhook "kube-anchor.kube-system.svc" denied the request: Cluster Freeze Window Enabled via Kube-Anchor ☸ 
```

For Disabling the Kube-Anchor Freeze

```
kubectl label ns test0-kube-anchor kube-anchor=disabled
```