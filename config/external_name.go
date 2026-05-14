package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// project_environment is imported as project_key:environment_name
	"project_environment": config.TemplatedStringAsIdentifier("name",
		"{{ .parameters.project_key }}:{{ .external_name }}"),

	// project_group is imported as project_key:group_name
	"project_group": config.TemplatedStringAsIdentifier("name",
		"{{ .parameters.project_key }}:{{ .external_name }}"),

	// project_repository is imported as project_key:repository_key
	"project_repository": config.TemplatedStringAsIdentifier("key",
		"{{ .parameters.project_key }}:{{ .external_name }}"),

	// project_role is imported as project_key:role_name
	"project_role": config.TemplatedStringAsIdentifier("name",
		"{{ .parameters.project_key }}:{{ .external_name }}"),

	// project_share_repository is imported as repo_key:target_project_key
	"project_share_repository": config.TemplatedStringAsIdentifier("repo_key",
		"{{ .external_name }}:{{ .parameters.target_project_key }}"),

	// project_share_repository_with_all is imported by its repo_key
	"project_share_repository_with_all": config.ParameterAsIdentifier("repo_key"),

	// project_user is imported as project_key:username
	"project_user": config.TemplatedStringAsIdentifier("name",
		"{{ .parameters.project_key }}:{{ .external_name }}"),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
