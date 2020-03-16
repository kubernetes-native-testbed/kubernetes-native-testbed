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

| Name | Description | Version | URL |
|--------------|-----------|---|---|
| Vitess | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| NATS | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| presslabs-mysql-operator | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| oracle-mysql-operator | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| cassandra-operator | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| elastic-cloud-operator (ElasticSearch, Kibana) | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| kafka-operator | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| mcrouter-operator | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| minio-operator | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| postgres-operator | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| prometheus-operator | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| redis-operator | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| tidb-operator (TiDB & TiKV) | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| rook-ceph | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| rook-yugabytedb | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| KubeDB | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| nginx-ingress | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| Harbor | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| Clair | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| kubernetes-dashboard v2 | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| Loki | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| weave-scope | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| Contour (Envoy) | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| Tekton Pipelines & Triggers | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| Kaniko | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| Telepresence | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| Kustomize | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| ArgoCD | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| conftest | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| gRPC | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| gRPC-web | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| (TBR): Knative | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |
| (TBR): OPA Gatekeeper | MySQL as a Service | v0.0.0 (MySQL: v0.0.0)| https://example.com |

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
./init.sh
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

