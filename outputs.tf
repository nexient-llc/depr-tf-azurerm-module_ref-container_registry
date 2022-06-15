output "resource_group_name" {
    value = local.resource_group_name
}

output "acr_name" {
    value = local.container_registry_name
}

output "acr_id" {
    value = module.container_registry.container_registry.id
    sensitive = true
}

output "acr_login_server" {
    value = module.container_registry.container_registry.login_server
    sensitive = true
}

output "acr_admin_username" {
    value = module.container_registry.container_registry.admin_username
    sensitive = true
}

output "acr_login_password" {
    value = module.container_registry.container_registry.admin_password
    sensitive = true
}

output "public_network_access_enabled" {
    value = module.container_registry.container_registry.public_network_access_enabled
    sensitive = true
}