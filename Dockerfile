FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.20-openshift-4.15 AS builder
WORKDIR /go/src/github.com/openshift/cluster-api-provider-openshift
COPY . .
RUN make build

FROM registry.ci.openshift.org/ocp/4.15:base
COPY --from=builder /go/src/github.com/openshift/cluster-api-provider-openshift/bin/manager .
COPY --from=builder /go/src/github.com/openshift/cluster-api-provider-openshift/manifests manifests
