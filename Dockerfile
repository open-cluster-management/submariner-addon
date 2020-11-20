FROM docker.io/openshift/origin-release:golang-1.14 AS builder
WORKDIR /go/src/github.com/open-cluster-management/submariner-addon
COPY . .
ENV GO_PACKAGE github.com/open-cluster-management/submariner-addon
RUN make build --warn-undefined-variables
RUN make build-e2e --warn-undefined-variables

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
# expose env vars for runtime
ENV KUBECONFIG "/.kube/config"
ENV OPTIONS "/resources/options.yaml"
ENV REPORT_FILE "/results/results.xml"

COPY --from=builder /go/src/github.com/open-cluster-management/submariner-addon/submariner /
COPY --from=builder /go/src/github.com/open-cluster-management/submariner-addon/e2e.test /
RUN microdnf update && microdnf clean all
