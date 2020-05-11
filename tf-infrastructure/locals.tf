locals {
  project_name = "bamboo-tracker"

  dev_tags = {
    PROJECT     = "${local.project_name}-dev"
    ENVIRONMENT = "dev"
  }
}
