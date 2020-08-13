provider "aws" {
  version = "~> 2.62"
  region  = var.region
  profile = var.profile_name
}

resource "aws_iam_role" "demo1_cluster" {
  name = "demo1_cluster"

  assume_role_policy = jsonencode({
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "eks.amazonaws.com"
      }
    }]
    Version = "2012-10-17"
  })
}

resource "aws_iam_role" "demo1_nodes" {
  name                 = "demo1_nodes"
  assume_role_policy = jsonencode({
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "ec2.amazonaws.com"
      }
    }]
    Version = "2012-10-17"
  })
}

resource "aws_iam_role_policy_attachment" "demo1_cluster_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
  role       = aws_iam_role.demo1_cluster.name
}

resource "aws_iam_role_policy_attachment" "demo1_service_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy"
  role       = aws_iam_role.demo1_cluster.name
}

resource "aws_iam_role_policy_attachment" "demo1_worker_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
  role       = aws_iam_role.demo1_nodes.name
}

resource "aws_iam_role_policy_attachment" "demo1_cni_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
  role       = aws_iam_role.demo1_nodes.name
}

resource "aws_iam_role_policy_attachment" "demo1_ecr_read_only_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
  role       = aws_iam_role.demo1_nodes.name
}

resource "aws_eks_cluster" "demo1" {
  name     = "demo1-eks"
  role_arn = aws_iam_role.demo1_cluster.arn

  vpc_config {
    subnet_ids = var.eks_subnets
  }

  depends_on = [
    aws_iam_role_policy_attachment.demo1_cluster_policy,
    aws_iam_role_policy_attachment.demo1_service_policy,
  ]
}

resource "aws_eks_node_group" "demo1_kubelets" {
  cluster_name    = aws_eks_cluster.demo1.name
  node_group_name = "demo1-kubelet"
  node_role_arn   = aws_iam_role.demo1_nodes.arn
  subnet_ids      = var.eks_subnets
  instance_types  = ["t3.small"]

  scaling_config {
    desired_size = 1
    max_size     = 3
    min_size     = 1
  }

  depends_on = [
    aws_iam_role_policy_attachment.demo1_worker_policy,
    aws_iam_role_policy_attachment.demo1_cni_policy,
    aws_iam_role_policy_attachment.demo1_ecr_read_only_policy,
  ]
}

data "aws_eks_cluster_auth" "demo1" {
  name = "demo1"
}

output "demo1-eks-endpoint" {
  value = "${aws_eks_cluster.demo1.endpoint}"
}

output "demo1-eks-kubeconfig-certificate-authority-data" {
  value = "${aws_eks_cluster.demo1.certificate_authority.0.data}"
}
