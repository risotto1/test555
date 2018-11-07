#!/bin/sh
BINARY=$1
CLUSTER=$2
CONFIG=$3
"$BINARYy" &
/usr/local/bin/envoy -c "${CONFIG}" --service-cluster "$CLUSTER"
