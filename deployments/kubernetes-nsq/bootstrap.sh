#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

helm repo add cockroachdb https://charts.cockroachdb.com/
helm repo update
helm install my-release --values "$DIR"/cockroach-helm-values.yml cockroachdb/cockroachdb