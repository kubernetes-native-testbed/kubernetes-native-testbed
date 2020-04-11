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
| Name                                                                             | Genre                     | Version | CNCF Project | URL                                                                                                                                         |
| -------------------------------------------------------------------------------- | ------------------------- | :-----: | :----------: | ------------------------------------------------------------------------------------------------------------------------------------------- |
| [Vitess](https://vitess.io/)                                                     | Relational Database       |  v0.0.0 |       ○      | [planetscale/vitess-operator](https://github.com/planetscale/vitess-operator)                                                               |
| [MySQL](https://www.mysql.com/)                                                  | Relational Database       |  v0.0.0 |              | [presslabs/mysql-operator](https://github.com/presslabs/mysql-operator)                                                                     |
| MySQL                                                                            | Relational Database       |  v0.0.0 |              | [oracle/mysql-operator](https://github.com/oracle/mysql-operator)                                                                           |
| [YugabyteDB](https://www.yugabyte.com/)                                          | Relational Database       |  v0.0.0 |              | [rook-yugabytedb](https://github.com/rook/rook/blob/master/cluster/examples/kubernetes/yugabytedb)                                          |
| [PostgreSQL](https://www.postgresql.org/)                                        | Relational Database       |  v0.0.0 |              | [zalando/postgres-operator](https://github.com/zalando/postgres-operator)                                                                   |
| [TiDB](https://pingcap.com/en)                                                   | Relational Database       |  v0.0.0 |              | [pingcap/tidb-operator](https://github.com/pingcap/tidb-operator)                                                                           |
| [TiKV](https://tikv.org/)                                                        | Key Value Store           |  v0.0.0 |       ○      | [pingcap/tidb-operator](https://github.com/pingcap/tidb-operator)                                                                           |
| [Memcached](https://memcached.org/)                                              | Key Value Store           |  v0.0.0 |              | [geerlingguy/mcrouter-operator ](https://github.com/geerlingguy/mcrouter-operator)                                                          |
| [Redis](https://redis.io/)                                                       | Key Value Store           |  v0.0.0 |              | [spotahome/redis-operator](https://github.com/spotahome/redis-operator)                                                                     |
| [Apache Cassandra](http://cassandra.apache.org/)                                 | NoSQL (RDB)               |  v0.0.0 |              | [Orange-OpenSource/casskop](https://github.com/Orange-OpenSource/casskop)                                                                   |
| [MongoDB](https://www.mongodb.com/)                                              | NoSQL (Document DB)       |  v0.0.0 |              | [kubedb/operator](https://github.com/kubedb/operator)                                                                                       |
| [NATS](https://nats.io/)                                                         | Message Queue             |  v0.0.0 |       ○      | [nats-io/nats-operator](https://github.com/nats-io/nats-operator)                                                                           |
| [Apache Kafka](https://kafka.apache.org/)                                        | Message Queue             |  v0.0.0 |    ○(\*1)    | [strimzi/strimzi-kafka-operator](https://github.com/strimzi/strimzi-kafka-operator)                                                         |
| [MinIO](https://min.io/)                                                         | Object Storage            |  v0.0.0 |              | [minio/minio-operator](https://github.com/minio/minio-operator)                                                                             |
| [Ceph](https://ceph.io/)                                                         | Block Storage             |  v0.0.0 |              | [rook-ceph](https://github.com/rook/rook/tree/master/cluster/examples/kubernetes/ceph)                                                      |
| [Rook](https://rook.io/)                                                         | Block Storage             |  v0.0.0 |       ○      | [rook/rook](https://github.com/rook/rook)                                                                                                   |
| [Nginx](https://www.nginx.com/)                                                  | Ingress Controller        |  v0.0.0 |              | [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx)                                                                     |
| [Envoy](https://www.envoyproxy.io/)                                              | L7 LoadBalancer           |  v0.0.0 |       ○      | [projectcontour/contour](https://github.com/projectcontour/contour)                                                                         |
| [Harbor](https://goharbor.io/)                                                   | Container Registry        |  v0.0.0 |       ○      | [goharbor/harbor](https://github.com/goharbor/harbor)             , [goharbor/harbor-operator](https://github.com/goharbor/harbor-operator) |
| Kaniko                                                                           | Container Build Tool      |  v0.0.0 |              | [GoogleContainerTools/kaniko](https://github.com/GoogleContainerTools/kaniko)                                                               |
| [Tekton](https://tekton.dev/) Triggers                                           | CI                        |  v0.0.0 |              | [tektoncd/triggers](https://github.com/tektoncd/triggers)                                                                                   |
| Tekton Pipelines                                                                 | CI                        |  v0.0.0 |              | [tektoncd/pipeline](https://github.com/tektoncd/pipeline)                                                                                   |
| [ArgoCD](https://argoproj.github.io/argo-cd/)                                    | CD                        |  v0.0.0 |              | [argoproj/argo-cd](https://github.com/argoproj/argo-cd)                                                                                     |
| [Prometheus](https://prometheus.io/)                                             | Monitoring                |  v0.0.0 |       ○      | [coreos/prometheus-operator](https://github.com/coreos/prometheus-operator)                                                                 |
| [Weave Scope](https://www.weave.works/docs/scope/latest/introducing/)            | Monitoring                |  v0.0.0 |              | [weaveworks/scope](https://github.com/weaveworks/scope)                                                                                     |
| Kubernetes Dashboard                                                             | Monitoring                |  v2.0.0 |              | [kubernetes/dashboard](https://github.com/kubernetes/dashboard)                                                                             |
| [Loki](https://grafana.com/oss/loki/)                                            | Logging                   |  v0.0.0 |              | [grafana/loki](https://github.com/grafana/loki)                                                                                             |
| [Telepresence](https://www.telepresence.io/)                                     | Local Development         |  v0.0.0 |       ○      | [telepresenceio/telepresence](https://github.com/telepresenceio/telepresence)                                                               |
| [Kustomize](https://kustomize.io/)                                               | Manifest Management       |  v0.0.0 |              | [kubernetes-sigs/kustomize](https://github.com/kubernetes-sigs/kustomize)                                                                   |
| conftest                                                                         | Manifest Management       |  v0.0.0 |              | [instrumenta/conftest](https://github.com/instrumenta/conftest)                                                                             |
| [gRPC](https://grpc.io/)                                                         | Microservice Interconnect |  v0.0.0 |       ○      | [grpc/grpc-go](https://github.com/grpc/grpc-go)                                                                                             |
| gRPC-web                                                                         | Microservice Interconnect |  v0.0.0 |              | [grpc/grpc-web](https://github.com/grpc/grpc-web)                                                                                           |
| [Elasticsearch](https://www.elastic.co/)/[Kibana](https://www.elastic.co/kibana) | EFKStack                  |  v0.0.0 |              | [elastic/cloud-on-k8s](https://github.com/elastic/cloud-on-k8s)                                                                             |
| Clair                                                                            | Security                  |  v0.0.0 |              | [quay/clair](https://github.com/quay/clair)                                                                                                 |
| (TBR): [OPA](https://www.openpolicyagent.org/) Gatekeeper                        | Security                  |  v0.0.0 |       ○      | [open-policy-agent/gatekeeper](https://github.com/open-policy-agent/gatekeeper)                                                             |
| (TBR): [Knative](https://knative.dev/)                                           | Serverless                |  v0.0.0 |              | [knative/serving](https://github.com/knative/serving), [knative/eventing](https://github.com/knative/eventing)                              |
\*1: strictly speaking, strimzi is just a member of CNCF project, not Apache Kafka.

# How to use

+ Fork repo to your org

from https://github.com/kubernetes-native-testbed/kubernetes-native-testbed

+ Edit environment settings（tools/env）

```
$ cat tools/env
export KUBERNETES_PLATFORM=gke
export GCP_PROJECT=replace-here
export GCP_REGION=asia-northeast1
```

+ Allocate static ip address

```
./tools/allocate_staticip.sh
```

+ Initialize repo to replace placeholder

```
./tools/init_repo.sh
```

+ Add webhook settings for forked repo

from https://github.com/GITHUB_ORG_NAME/kubernetes-native-testbed/settings/hooks

  * Payload URL: https://tekton.LOADBALANCER_IP_ADDRESS.nip.io/event-listener
  * Content type: application/json
  * Secret: sample-github-webhook-secret
    * if you want to change, please edit manifests/infra/instances/ci.yaml
  * Enable SSL verification: [*]
  * Just the push event: [*]
  * Active: [*]

+ Deploy applications and so on

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
  * cicd/: CI/CD pipeline settings
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

