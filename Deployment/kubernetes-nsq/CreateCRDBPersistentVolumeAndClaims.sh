#!/bin/bash
set -exuo pipefail

# Clean up anything from a prior run:
kubectl delete statefulsets,persistentvolumes,persistentvolumeclaims,services,poddisruptionbudget -l app.kubernetes.io/component=cockroachdb

# Make persistent volumes and (correctly named) claims. We must create the
# claims here manually even though that sounds counter-intuitive. For details
# see https://github.com/kubernetes/contrib/pull/1295#issuecomment-230180894.
# Note that we make an extra volume here so you can manually test scale-up.
for i in $(seq 0 2); do
  cat <<EOF | kubectl create -f -
kind: PersistentVolume
apiVersion: v1
metadata:
  name: pv-cockroach${i}
  labels:
    type: local
    app.kubernetes.io/component: cockroachdb
    app.kubernetes.io/instance: scoretrak
    app.kubernetes.io/name: cockroachdb
spec:
  capacity:
    storage: 100Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/tmp/cockroachdb/${i}"
EOF
done;
