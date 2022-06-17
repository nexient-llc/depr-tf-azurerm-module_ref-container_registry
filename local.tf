locals {

  resource_group = {
    location = var.resource_group.location
    name = local.resource_group_name
  }

  resource_group_name          = module.resource_name["resource_group"].recommended_per_length_restriction
  container_registry_name      = module.resource_name["container_registry"].lower_case
}