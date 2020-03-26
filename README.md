# kubernetes-native-testbed

`THIS REPOSITORY IS STILL WORK IN PROGRESS NOW. WE PLAN TO PUBLISH AT APRIL 2020`

This is fully Kubernetes-native testbed environment.
Please contribute for add additional OSS (Vitess, NATS, etc) or microservices.

# Authors

* Masaya Aoyama [@amsy810](https://twitter.com/amsy810)
* Mizuki Urushida [@zuiurs](https://twitter.com/zuiurs)

# Microservices

| microservice | datastore |   |   |   |
|--------------|-----------|---|---|---|
| cart      | TiKV          |   |   |   |
| comment         | MongoDB          |   |   |   |
| delivery-status         | Cassandra, NATS          |   |   |   |
| order         |  TiDB         |   |   |   |
| payment-info         | PostgreSQL          |   |   |   |
| point         | YugabyteDB, Kafka, Memcached(MCRouter)          |   |   |   |
| product         | MySQL          |   |   |   |
| rate         | Redis(Centinel)          |   |   |   |
| search         | Elasticsearch          |   |   |   |
| user         | MySQL          |   |   |   |
| admin        | -          |   |   |   |

# OSS
| Name                                           | Description        | Version                | CNCF Project | URL                 |
| ---------------------------------------------- | ------------------ | ---------------------- | ------------ | ------------------- |
| Vitess                                         | Relational Database | v0.0.0 (MySQL: v0.0.0) |              | https://vitess.io/<br>https://github.com/planetscale/vitess-operator|
| presslabs-mysql-operator                       | Relational Database | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/presslabs/mysql-operator |
| oracle-mysql-operator                          | Relational Database | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/oracle/mysql-operator |
| rook-yugabytedb                                | Relational Database | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/rook/rook/blob/master/cluster/examples/kubernetes/yugabytedb |
| postgres-operator                              | Relational Database | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/zalando/postgres-operator |
| tidb-operator (TiDB & TiKV)                    | RDB / KVS | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/pingcap/tidb-operator |
| mcrouter-operator                              | Key Value Store | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/geerlingguy/mcrouter-operator |
| redis-operator                                 | Key Value Store | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/spotahome/redis-operator |
| cassandra-operator (CassKop)                            | NoSQL (RDB) | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/Orange-OpenSource/casskop |
| KubeDB (MongoDB)                               | NoSQL (Document DB)| v0.0.0 (MySQL: v0.0.0) |              | https://github.com/kubedb/operator |
| nats-operator                                           | Message Queue | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/nats-io/nats-operator |
| kafka-operator                                 | Message Queue | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/strimzi/strimzi-kafka-operator |
| minio-operator                                 | Object Storage | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/minio/minio-operator |
| rook-ceph                                      | Block Storage | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/rook/rook/tree/master/cluster/examples/kubernetes/ceph |
| nginx-ingress                                  | Ingress Controller | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/kubernetes/ingress-nginx |
| Contour (Envoy)                                | L7 LoadBalancer | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/projectcontour/contour |
| Harbor                                         | Container Registry | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/goharbor/harbor |
| Kaniko                                         | Container Build Tool | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/GoogleContainerTools/kaniko |
| Tekton Pipelines & Triggers                    | CI | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/tektoncd/triggers |
| ArgoCD                                         | CD | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/argoproj/argo-cd |
| prometheus-operator                            | Monitoring | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/coreos/prometheus-operator |
| weave-scope                                    | Monitoring | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/weaveworks/scope |
| kubernetes-dashboard v2                        | Monitoring | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/kubernetes/dashboard |
| Loki                                           | Logging | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/grafana/loki |
| Telepresence                                   | Local Development | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/telepresenceio/telepresence |
| Kustomize                                      | Manifest Management | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/kubernetes-sigs/kustomize |
| conftest                                       | Manifest Management | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/instrumenta/conftest |
| gRPC                                           | Microservice Interconnect | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/grpc/grpc-go |
| gRPC-web                                       | Microservice Interconnect | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/grpc/grpc-web |
| elastic-cloud-operator (ElasticSearch, Kibana) | Search Engine | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/elastic/cloud-on-k8s |
| Clair                                          | Security | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/quay/clair |
| (TBR): OPA Gatekeeper                          | Security | v0.0.0 (MySQL: v0.0.0) |              | https://github.com/open-policy-agent/gatekeeper |
| (TBR): Knative                                 | Serverless | v0.0.0 (MySQL: v0.0.0) |              | https://knative.dev/ |

# How to use

* pre-requirement
  * "type: LoadBalancer" service provide global IP address
  * replace xxx.xxx.xxx.xxx.nip.io
    * contour LoadBalancer static IP
    * Certification settings for issuer
    * webhook url from github for tekton triggers
  * fork this repository and replace repo settings
  * set up github webhook to tekton triggers with event-listener secret [please-modify-for-high-security-here]

* run following command

```
./tools/init.sh
```

# Destroy clusters

```
./tools/destroy_cluster.sh
```

# Directory structure

* manifests/
  * Kubernetes manifests
  * infra/: system or infrastructure manifests
    * ci-manifests/: tekton pipelines manifests
    * cd-manifests/: argocd pipelines manifests
* microservices/
  * microservice application source code
* development/
  * development manifests (skaffold, kustomization patch file)
* tools/
  * tools and scripts

# Local development

```
# initializing
source ./development/initialize.sh

# start local development for admin
./development/local-development.sh admin

# access service with  whole system or only local check
https://testbed.xxx.xxx.xxx.xxx.nip.io/admin/index.html
  or
http://localhost:8080/
```

# Memo

https://docs.google.com/spreadsheets/d/18Pza74gohErR-58ib8nUFeJcMJaTr65Jalh7EKAVc7g/edit#gid=0

