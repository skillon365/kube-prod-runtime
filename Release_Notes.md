Remove AKS 1.19 from compatability matrix

At this moment, BKPR only supports k8s systems that use the docker CRI.
As of AKS 1.19, containerd is the only container runtime option.

Until further analysis on how we want to support containerd, we're
reducing the scope of compability for now.

# Installation

Follow the [installation guide](https://github.com/bitnami/kube-prod-runtime/blob/master/docs/install.md) to install the Bitnami Kubernetes Production Runtime (BKPR) on your cluster.

# Supported Platforms

- [Azure Kubernetes Service (AKS)](https://github.com/bitnami/kube-prod-runtime/blob/master/docs/quickstart-aks.md)
- [Google Kubernetes Engine (GKE)](https://github.com/bitnami/kube-prod-runtime/blob/master/docs/quickstart-gke.md)
- [Amazon Elastic Container Service for Kubernetes (EKS)](https://github.com/bitnami/kube-prod-runtime/blob/master/docs/quickstart-eks.md)
- [BKPR on Generic Kubernetes Cluster (experimental)](https://github.com/bitnami/kube-prod-runtime/blob/master/docs/quickstart-generic.md)

# Migration Guide

- [BKPR v1.4 Migration Guide](https://github.com/bitnami/kube-prod-runtime/blob/master/docs/migration-guides/bkpr-1.4-migration-guide.md)

# Breaking Changes in Components

- Kibana 6.7 index (BKPR 1.3) format is not compatible with Kibana 7.5

# Changelog

- Remove AKS 1.19 from compatability matrix
- Update test matrix for v1.9.x
- Update README to reflect current BKPR versions
- Update component list
- [maintenance/master] 'mariadb-galera' updated '10.5.9' -> '10.5.10' (#1231)
- [maintenance/master] 'grafana' updated '7.5.4' -> '7.5.7' (#1266)
- elasticsearch-curator: component image updated to 'bitnami/elasticsearch-curator:5.8.4-debian-10-r31' (#1304)
- Mark TLS tests as Pending
- Update custodian to clean correct Azure subscription
- Explain how cert-manager assertions can be flaky
- Version tests support local builds and release versions
- Use new AZ resource group in BKPR subscription
- Assert on v2 alert manager API path in monitoring tests
- Update Let's Encrypt Staging Root CA
- Removing k8s 1.17 from AKS test matrix
- Reverting nginx-ingress-controller back to 0.34.1
- Replace credential in jenkins with new service principal
- Update Azure DNS info for new subscription
- Version matrix update for 1.8
- Reordered components in alphabetical order
- grafana: component image updated to 'bitnami/grafana:7.5.4-debian-10-r0' (#1186)
- elasticsearch-curator: component image updated to 'bitnami/elasticsearch-curator:5.8.3-debian-10-r117' (#1185)
- [maintenance/master] 'kibana' updated '7.9.0' -> '7.12.0' (#1173)
- [maintenance/master] 'elasticsearch' updated '7.9.0' -> '7.12.0' (#1171)
- [maintenance/master] 'prometheus' updated '2.22.1' -> '2.26.0' (#1170)
- [maintenance/master] 'fluentd' updated '1.11.2' -> '1.12.2' (#1160)
- [maintenance/master] 'node-exporter' updated '1.0.1' -> '1.1.2' (#1128)
- [maintenance/master] 'kube-state-metrics' updated '1.9.7' -> '1.9.8' (#1114)
- [maintenance/master] 'mariadb-galera' updated '10.4.12' -> '10.5.9' (#1108)
- [maintenance/master] 'nginx-ingress-controller' updated '0.35.0' -> '0.44.0' (#1100)
- [maintenance/master] 'configmap-reload' updated '0.4.0' -> '0.5.0' (#1070)
- [maintenance/master] 'external-dns' updated '0.7.3' -> '0.7.6' (#1051)
- Remove dependency on squid deployment (#1124)
- Update install.md (#1021)
- Merge #974
- prometheus: component image updated to 'bitnami/prometheus:2.22.1-debian-10-r7'
- prometheus: component image updated to 'bitnami/prometheus:2.22.1-debian-10-r6'
- prometheus: component image updated to 'bitnami/prometheus:2.22.1-debian-10-r5'
- prometheus: component image updated to 'bitnami/prometheus:2.22.1-debian-10-r4'
- prometheus: component image updated to 'bitnami/prometheus:2.22.1-debian-10-r3'
- prometheus: component image updated to 'bitnami/prometheus:2.22.1-debian-10-r2'
- prometheus: component image updated to 'bitnami/prometheus:2.22.1-debian-10-r1'
- prometheus: component image updated to 'bitnami/prometheus:2.22.1-debian-10-r0'
- Fix: Remove jjo from project maintainer
- hotfix: Fix GKE Custodian pipelines
- Merge #952
- Feat: Update GKE versions
- Fix: Update Custodian version to custom Bitnami container and fix/improve AWS / Google policies (#935)
- Merge #932
- Merge #933
- Merge #929
- Update Jenkinsfile
- Update Jenkinsfile
- nginx-ingress-controller: component image updated to 'bitnami/nginx-ingress-controller:0.35.0-debian-10-r10'
- Fix: remove subnets
- Fix cloud resources in custodian
- Feat: Improve the AWS Cloud-Custodian policies
- Feat: Update the default Cloud Custodian version for BKPR
- Update Jenkinsfile
- Update Jenkinsfile
- nginx-ingress-controller: component image updated to 'bitnami/nginx-ingress-controller:0.35.0-debian-10-r8'
- Increase timeout of the Jenkins job
- nginx-ingress-controller: component image updated to 'bitnami/nginx-ingress-controller:0.35.0-debian-10-r7'
- nginx-ingress-controller: component image updated to 'bitnami/nginx-ingress-controller:0.35.0-debian-10-r6'
- nginx-ingress-controller: component image updated to 'bitnami/nginx-ingress-controller:0.35.0-debian-10-r5'
- nginx-ingress-controller: component image updated to 'bitnami/nginx-ingress-controller:0.35.0-debian-10-r4'
- nginx-ingress-controller: component image updated to 'bitnami/nginx-ingress-controller:0.35.0-debian-10-r3'
- Update images.json
- Update images.json
- nginx-ingress-controller: component image updated to 'bitnami/nginx-ingress-controller:0.35.0-debian-10-r2'
- nginx-ingress-controller: component image updated to 'bitnami/nginx-ingress-controller:0.35.0-debian-10-r1'
- nginx-ingress-controller: component image updated to 'bitnami/nginx-ingress-controller:0.35.0-debian-10-r0'
- Merge #916
- grafana: component image updated to 'bitnami/grafana:7.1.5-debian-10-r7'
- no-change
- Update images.json
- Update images.json
- grafana: component image updated to 'bitnami/grafana:7.1.5-debian-10-r6'
- grafana: component image updated to 'bitnami/grafana:7.1.5-debian-10-r5'
- grafana: component image updated to 'bitnami/grafana:7.1.5-debian-10-r4'
- grafana: component image updated to 'bitnami/grafana:7.1.5-debian-10-r3'
- Merge #921
- no-change
- no-change commit after updating Jenkins's eksctl for 1.17
- Merge #910
- also update jenkins/cloud-custodian/Jenkinsfile
- Update images.json
- grafana: component image updated to 'bitnami/grafana:7.1.5-debian-10-r2'
- elasticsearch: component image updated to 'bitnami/elasticsearch:7.9.0-debian-10-r4'
- [jjo] move support window to 1.16,1.17 on all clouds
- Merge #918
- grafana: component image updated to 'bitnami/grafana:7.1.5-debian-10-r1'
- 
- 
- 
- Update images.json
- [jjo] fix: oauth2-proxy >= 6.0.0 doesn't provide "X-Auth-Request-User" header on AKS, use "X-Auth-Request-Email" instead
- grafana: component image updated to 'bitnami/grafana:7.1.5-debian-10-r0'
- elasticsearch: component image updated to 'bitnami/elasticsearch:7.9.0-debian-10-r3'
- elasticsearch: component image updated to 'bitnami/elasticsearch:7.9.0-debian-10-r2'
- elasticsearch: component image updated to 'bitnami/elasticsearch:7.9.0-debian-10-r1'
- elasticsearch: component image updated to 'bitnami/elasticsearch:7.9.0-debian-10-r0'
