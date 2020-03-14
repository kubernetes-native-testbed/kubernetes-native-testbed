# kubernetes-native-testbed

This is fully Kubernetes-native testbed environment.
Please contribute for add additional OSS (Vitess, NATS, etc) or microservices.

# Authors

* Masaya Aoyama [@amsy810](https://twitter.com/amsy810)
* Mizuki Urushida [@zuiurs](https://twitter.com/zuiurs)

# Microservices
  
| microservice | datastore |   |   |   |
|--------------|-----------|---|---|---|
| product      |           |   |   |   |
| user         |           |   |   |   |
| admin        |           |   |   |   |

# OSS

* Vitess
* NATS
* presslabs-mysql-operator
* oracle-mysql-operator

# Local development

```
# initializing
source ./development/initialize.sh

# start local development for admin
./development/local-development.sh adminE

# access service with  whole system or only local check
https://testbed.xxx.xxx.xxx.xxx.nip.io/admin/index.html
  or
http://localhost:8080/
```

# Memo

https://docs.google.com/spreadsheets/d/18Pza74gohErR-58ib8nUFeJcMJaTr65Jalh7EKAVc7g/edit#gid=0

