/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	authmethod "github.com/fire-ant/provider-nomad/internal/controller/acl/authmethod"
	bindingrule "github.com/fire-ant/provider-nomad/internal/controller/acl/bindingrule"
	policy "github.com/fire-ant/provider-nomad/internal/controller/acl/policy"
	role "github.com/fire-ant/provider-nomad/internal/controller/acl/role"
	token "github.com/fire-ant/provider-nomad/internal/controller/acl/token"
	config "github.com/fire-ant/provider-nomad/internal/controller/nomad/config"
	job "github.com/fire-ant/provider-nomad/internal/controller/nomad/job"
	nomadnamespace "github.com/fire-ant/provider-nomad/internal/controller/nomad/nomadnamespace"
	policynomad "github.com/fire-ant/provider-nomad/internal/controller/nomad/policy"
	specification "github.com/fire-ant/provider-nomad/internal/controller/nomad/specification"
	volume "github.com/fire-ant/provider-nomad/internal/controller/nomad/volume"
	providerconfig "github.com/fire-ant/provider-nomad/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authmethod.Setup,
		bindingrule.Setup,
		policy.Setup,
		role.Setup,
		token.Setup,
		config.Setup,
		job.Setup,
		nomadnamespace.Setup,
		policynomad.Setup,
		specification.Setup,
		volume.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
