# Dyplomat

Dyplomat is an example Envoy control plane implementation built on top of go-control-plane that demonstrates how you can use kubernetes informers to build a multi Kubernetes cluster Envoy serivce mesh.

## AWS Setup Instructions

This example comes with sample terraform to run a Kubernetes-aware control plane in an AWS VPC with 2 EKS clusters serving traffic across two hello world apps.

### Prerequisites

You will need:
1. [terraform cli](https://www.terraform.io/downloads.html) installed (>v0.12)
1. An AWS account with credentials set up locally in `~/.aws/credentials`

### Setup

1. Initialize terraform: 
    ```
    $ cd terraform/
    $ terraform init
    ```
1. Configure `eks/variables.tf`. Fill in your AWS profile to use for auth (from `~/.aws/config`), and two subnets in your default VPC for that account that EKS will use to bootstrap the clusters.
1. Create EKS clusters:
    ```
    $ terraform apply -target=module.eks
    ```
1. Use the AWS UI to add the security group of the other cluster to each cluster to allow traffic across clusters.
1. Set up [local kubectl](https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html) for your new EKS clusters.
1. Build `dyplomat` docker image and push to your registry:
    ```
    $ cd examples/dyplomat
    $ docker build . -t $REGISTRY/dyplomat:latest && docker push $REGISTRY/dyplomat:latest
    ```
1. Fill in `rbac/variables.tf`, then configure IAM role for dyplomat by running the `rbac` terraform module.
    ```
    $ terraform apply -target=module.rbac
    ```
1. Fill in your cluster api server URL's and cluster CA's in `bootstrap.yaml`
1. Run `apps` terraform against each EKS cluster:
    ```
    $ kubectl config use-context demo1-eks
    $ terraform apply -target=module.apps
    $ kubectl config use-context demo2-eks
    $ terraform apply -target=module.apps
    ```
    This will bring up a front envoy, dyplomat, and two hello world apps (nginx web servers) in each cluster.

## Adding other cloud providers

The example dyplomat implementation can be extended to be compatible with other cloud providers by implementing a new `TokenKind` in `dyplomat/auth.go`
