package cluster

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const shortGroup = "project"

// Configure configures all jfrog/project resources for the cluster-scoped provider.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("project_environment", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Environment"
	})

	p.AddResourceConfigurator("project_group", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Group"
	})

	p.AddResourceConfigurator("project_repository", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Repository"
	})

	p.AddResourceConfigurator("project_role", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Role"
	})

	p.AddResourceConfigurator("project_share_repository", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ShareRepository"
	})

	p.AddResourceConfigurator("project_share_repository_with_all", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ShareRepositoryWithAll"
	})

	p.AddResourceConfigurator("project_user", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "User"
	})
}
