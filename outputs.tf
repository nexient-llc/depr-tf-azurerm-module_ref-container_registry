output "resource_group_name" {
    value = local.resource_group_name
}

output "acr_name" {
    value = local.container_registry_name
}

output "container_registry" {
    value = module.container_registry.container_registry
    sensitive = true
}