package image

import (
	igntypes "github.com/coreos/ignition/v2/config/v3_2/types"
	"github.com/openshift/installer/pkg/agent/imagebuilder"
	"github.com/openshift/installer/pkg/asset"
)

// Ignition is an asset that generates the agent installer ignition file.
type Ignition struct {
	Config *igntypes.Config
}

// Name returns the human-friendly name of the asset.
func (a *Ignition) Name() string {
	return "Agent Installer Ignition"
}

// Dependencies returns the assets on which the Ignition asset depends.
func (a *Ignition) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Generate generates the agent installer ignition.
func (a *Ignition) Generate(dependencies asset.Parents) error {
	configBuilder, err := imagebuilder.New()
	if err != nil {
		return err
	}

	ignition, err := configBuilder.IgnitionConfig()
	if err != nil {
		return err
	}

	a.Config = &ignition
	return nil
}
