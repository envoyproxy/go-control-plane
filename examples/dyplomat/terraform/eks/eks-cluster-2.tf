resource "aws_iam_role" "demo2_cluster" {
  name = "demo2_cluster"

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

resource "aws_iam_role" "demo2_nodes" {
  name                 = "demo2_nodes"
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

resource "aws_iam_role_policy_attachment" "demo2_cluster_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
  role       = aws_iam_role.demo2_cluster.name
}

resource "aws_iam_role_policy_attachment" "demo2_service_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy"
  role       = aws_iam_role.demo2_cluster.name
}

resource "aws_iam_role_policy_attachment" "demo2_worker_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
  role       = aws_iam_role.demo2_nodes.name
}

resource "aws_iam_role_policy_attachment" "demo2_cni_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
  role       = aws_iam_role.demo2_nodes.name
}

resource "aws_iam_role_policy_attachment" "demo2_ecr_read_only_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
  role       = aws_iam_role.demo2_nodes.name
}

resource "aws_eks_cluster" "demo2" {
  name     = "demo2-eks"
  role_arn = aws_iam_role.demo2_cluster.arn

  vpc_config {
    subnet_ids = var.eks_subnets
  }

  depends_on = [
    aws_iam_role_policy_attachment.demo2_cluster_policy,
    aws_iam_role_policy_attachment.demo2_service_policy,
  ]
}

resource "aws_eks_node_group" "demo2_kubelets" {
  cluster_name    = aws_eks_cluster.demo2.name
  node_group_name = "demo2-kubelet"
  node_role_arn   = aws_iam_role.demo2_nodes.arn
  subnet_ids      = var.eks_subnets
  instance_types  = ["t3.small"]

  scaling_config {
    desired_size = 1
    max_size     = 3
    min_size     = 1
  }

  depends_on = [
    aws_iam_role_policy_attachment.demo2_worker_policy,
    aws_iam_role_policy_attachment.demo2_cni_policy,
    aws_iam_role_policy_attachment.demo2_ecr_read_only_policy,
  ]
}

data "aws_eks_cluster_auth" "demo2" {
  name = "demo2"
}

output "demo2-eks-endpoint" {
  value = "${aws_eks_cluster.demo2.endpoint}"
}

output "demo2-eks-kubeconfig-certificate-authority-data" {
  value = "${aws_eks_cluster.demo2.certificate_authority.0.data}"
}
