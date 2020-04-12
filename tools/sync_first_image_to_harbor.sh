#!/bin/bash

CURRENT_DIR=$(cd $(dirname $0); pwd)

OP=${1:-apply}

cat << "_EOF_" | kubectl ${OP} -f -
apiVersion: v1
kind: ConfigMap
metadata:
  name: sync-image-to-harbor
  namespace: infra
data:
  sync_image_to_harbor.sh: |-
    echo "===== syncing $MICROSERVICE image ====================================================="
    docker login -u ${HARBOR_USER} -p ${HARBOR_PASS} registry-harbor-core.infra.svc.cluster.local
    docker pull k8stestbed/${MICROSERVICE}:latest-cache
    docker tag k8stestbed/${MICROSERVICE}:latest-cache registry-harbor-core.infra.svc.cluster.local/library/${MICROSERVICE}:latest-cache
    docker images | grep ${MICROSERVICE}
    docker push registry-harbor-core.infra.svc.cluster.local/library/${MICROSERVICE}:latest-cache
_EOF_

for MICROSERVICE in `find ${CURRENT_DIR}/../microservices -name Dockerfile | awk -F "/" '{print $(NF-1)}'`; do
cat << _EOF_ | kubectl ${OP} -f -
apiVersion: batch/v1
kind: Job
metadata:
  name: sync-image-to-harbor-${MICROSERVICE}
  namespace: infra
spec:
  completions: 1
  parallelism: 1
  backoffLimit: 100
  template:
    spec:
      restartPolicy: OnFailure
      containers:
      - name: docker-container
        image: docker:19.03.8-dind
        command: ["sh", "/config/sync_image_to_harbor.sh"]
        env:
        - name: MICROSERVICE
          value: ${MICROSERVICE}
        - name: HARBOR_USER
          value: admin
        - name: HARBOR_PASS
          valueFrom:
            secretKeyRef:
              name: registry-harbor-core
              key: HARBOR_ADMIN_PASSWORD
        volumeMounts:
        - name: docker-sock
          mountPath: /var/run/docker.sock
        - name: config-volume
          mountPath: /config
      volumes:
        - name: docker-sock
          hostPath:
            path: /var/run/docker.sock
        - name: config-volume
          configMap:
            name: sync-image-to-harbor
            items:
            - key: sync_image_to_harbor.sh
              path: sync_image_to_harbor.sh
              mode: 493 # 0755
_EOF_
done

