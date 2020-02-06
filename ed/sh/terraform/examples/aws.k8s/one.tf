provider "aws" {
  shared_credentials_file = "/Users/k/.aws/credentials"
  region                  = "eu-central-1"
  profile                 = "x"
}

module "my-cluster" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = "my-cluster"
  cluster_version = "1.14"
  subnets         = ["subnet-*", "subnet-*b"]
  vpc_id          = "vpc-*"

  worker_groups = [
    {
      instance_type = "m4.large"
      asg_max_size  = 2
    }
  ]
}

data "aws_eks_cluster" "cluster" {
  name = module.my-cluster.cluster_id
}

data "aws_eks_cluster_auth" "cluster" {
  name = module.my-cluster.cluster_id
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.cluster.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
  token                  = data.aws_eks_cluster_auth.cluster.token
  load_config_file       = false
  version                = "~> 1.9"
}
