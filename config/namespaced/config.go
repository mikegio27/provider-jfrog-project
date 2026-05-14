package namespaced

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const shortGroup = "project"

// Configure configures all jfrog/project resources for the namespaced provider.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("project", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Project"
	})

	p.AddResourceConfigurator("project_environment", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Environment"
		r.References["project_key"] = ujconfig.Reference{
			TerraformName: shortGroup,
		}
	})

	p.AddResourceConfigurator("project_group", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Group"
		r.References["project_key"] = ujconfig.Reference{
			TerraformName: shortGroup,
		}
	})

	p.AddResourceConfigurator("project_repository", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Repository"
		r.References["project_key"] = ujconfig.Reference{
			TerraformName: shortGroup,
		}
	})

	p.AddResourceConfigurator("project_role", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Role"
		r.References["project_key"] = ujconfig.Reference{
			TerraformName: shortGroup,
		}
	})

	p.AddResourceConfigurator("project_share_repository", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ShareRepository"
		r.References["target_project_key"] = ujconfig.Reference{
			TerraformName: shortGroup,
		}
	})

	p.AddResourceConfigurator("project_share_repository_with_all", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ShareRepositoryWithAll"
	})

	p.AddResourceConfigurator("project_user", func(r *ujconfig.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "User"
		r.References["project_key"] = ujconfig.Reference{
			TerraformName: shortGroup,
		}
	})
}
