module "ec2" {
  source              = "../../modules/ec2"
  region              = var.region
  env_name            = local.environment
  project             = "xtest"
}
