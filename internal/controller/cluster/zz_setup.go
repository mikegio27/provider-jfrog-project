// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	environment "github.com/mikegio27/provider-jfrog-project/internal/controller/cluster/project/environment"
	group "github.com/mikegio27/provider-jfrog-project/internal/controller/cluster/project/group"
	project "github.com/mikegio27/provider-jfrog-project/internal/controller/cluster/project/project"
	repository "github.com/mikegio27/provider-jfrog-project/internal/controller/cluster/project/repository"
	role "github.com/mikegio27/provider-jfrog-project/internal/controller/cluster/project/role"
	sharerepository "github.com/mikegio27/provider-jfrog-project/internal/controller/cluster/project/sharerepository"
	sharerepositorywithall "github.com/mikegio27/provider-jfrog-project/internal/controller/cluster/project/sharerepositorywithall"
	user "github.com/mikegio27/provider-jfrog-project/internal/controller/cluster/project/user"
	providerconfig "github.com/mikegio27/provider-jfrog-project/internal/controller/cluster/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		environment.Setup,
		group.Setup,
		project.Setup,
		repository.Setup,
		role.Setup,
		sharerepository.Setup,
		sharerepositorywithall.Setup,
		user.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		environment.SetupGated,
		group.SetupGated,
		project.SetupGated,
		repository.SetupGated,
		role.SetupGated,
		sharerepository.SetupGated,
		sharerepositorywithall.SetupGated,
		user.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
