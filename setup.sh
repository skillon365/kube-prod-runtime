#!/usr/bin/bash

source env.prod

az account set --subscription $SUBSCRIPTION

az group create \
    --name $RESOURCE_GROUP \
    --subscription $SUBSCRIPTION \
    --location $LOCATION

az aks create \
	--subscription $SUBSCRIPTION \
    --resource-group $RESOURCE_GROUP \
    --name $AKS_CLUSTER \
    --generate-ssh-keys \
    --kubernetes-version $AKS_K8S_VERSION \
    --nodepool-name system \
    --vm-set-type VirtualMachineScaleSets \
    --load-balancer-sku standard \
    --enable-cluster-autoscaler \
    --node-vm-size $VM_SIZE \
    --node-osdisk-size $DISK_SIZE \
    --node-count 1 \
    --min-count 1 \
    --max-count 1

az aks nodepool add \
	--subscription $SUBSCRIPTION \
    --resource-group $RESOURCE_GROUP \
    --cluster-name $AKS_CLUSTER \
    --name spot \
    --priority Spot \
    --spot-max-price -1 \
    --eviction-policy Delete \
    --node-vm-size $VM_SIZE \
    --node-osdisk-size $DISK_SIZE \
    --enable-cluster-autoscaler \
    --node-count 1 \
    --min-count 1 \
    --max-count 3

az aks get-credentials \
	--subscription $SUBSCRIPTION \
	--resource-group $RESOURCE_GROUP \
    --name $AKS_CLUSTER \
    --overwrite-existing

./kubeprod/bin/kubeprod install aks \
  --email ${ADMIN} \
  --dns-zone "${DNS_ZONE}" \
  --dns-resource-group "${RESOURCE_GROUP}" \
  --manifests ./manifests

az network dns zone show \
  --name ${DNS_ZONE} \
  --resource-group ${RESOURCE_GROUP} \
  --query nameServers -o tsv

az group delete \
	--subscription $SUBSCRIPTION \
   	--name $RESOURCE_GROUP \
    --yes
