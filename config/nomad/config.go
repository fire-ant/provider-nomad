package nomad

import "github.com/upbound/upjet/pkg/config"

const (
	grp1 = "acl"
	grp2 = "nomad"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nomad_acl_auth_method", func(r *config.Resource) {
		r.ShortGroup = grp1
		r.ExternalName = config.NameAsIdentifier
	})
	p.AddResourceConfigurator("nomad_acl_binding_rule", func(r *config.Resource) {
		r.ShortGroup = grp1
		r.ExternalName = config.NameAsIdentifier
	})
	p.AddResourceConfigurator("nomad_acl_policy", func(r *config.Resource) {
		r.ShortGroup = grp1
		r.ExternalName = config.NameAsIdentifier
	})
	p.AddResourceConfigurator("nomad_acl_role", func(r *config.Resource) {
		r.ShortGroup = grp1
		r.ExternalName = config.NameAsIdentifier
	})
	p.AddResourceConfigurator("nomad_acl_token", func(r *config.Resource) {
		r.ShortGroup = grp1
		r.ExternalName = config.NameAsIdentifier
	})
	p.AddResourceConfigurator("nomad_acl_external_volume", func(r *config.Resource) {
		r.ShortGroup = grp1
		r.ExternalName = config.NameAsIdentifier
	})
	p.AddResourceConfigurator("nomad_job", func(r *config.Resource) {
		r.ShortGroup = grp2
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("nomad_namespace", func(r *config.Resource) {
		r.ShortGroup = grp2
		r.Kind = "NomadNamespace"
		r.ExternalName = config.NameAsIdentifier
	})
	p.AddResourceConfigurator("nomad_quota_specification", func(r *config.Resource) {
		r.ShortGroup = grp2
		r.ExternalName = config.NameAsIdentifier
	})
	p.AddResourceConfigurator("nomad_scheduler_config", func(r *config.Resource) {
		r.ShortGroup = grp2
		r.ExternalName = config.NameAsIdentifier
	})
	p.AddResourceConfigurator("nomad_sentinel_policy", func(r *config.Resource) {
		r.ShortGroup = grp2
		r.ExternalName = config.NameAsIdentifier
	})
	p.AddResourceConfigurator("nomad_volume", func(r *config.Resource) {
		r.ShortGroup = grp2
		r.ExternalName = config.NameAsIdentifier
	})
}
