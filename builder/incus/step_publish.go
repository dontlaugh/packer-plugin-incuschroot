// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incus

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type stepPublish struct{}

func (s *stepPublish) Run(_ctx context.Context, state multistep.StateBag) multistep.StepAction {

	var remote string

	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packersdk.Ui)

	if config.SkipPublish {
		ui.Say("skip_publish is true. Skipping step for publish")
		return multistep.ActionContinue
	}

	name := config.ContainerName
	stopArgs := []string{
		// We created the container with "--ephemeral=false" so we know it is safe to stop.
		"stop", name,
	}

	ui.Say("Stopping instance...")
	_, err := IncusCommand(stopArgs...)
	if err != nil {
		err := fmt.Errorf("Error stopping instance: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if config.PublishRemoteName != "" {
		remote = config.PublishRemoteName + ":"
	}

	publishArgs := []string{
		"publish", name, remote, "--alias", config.OutputImage,
	}

	for k, v := range config.PublishProperties {
		publishArgs = append(publishArgs, fmt.Sprintf("%s=%s", k, v))
	}

	ui.Say("Publishing image...")
	stdoutString, err := IncusCommand(publishArgs...)
	if err != nil {
		err := fmt.Errorf("Error publishing image: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	r := regexp.MustCompile("([0-9a-fA-F]+)$")
	fingerprint := r.FindAllStringSubmatch(stdoutString, -1)[0][0]

	ui.Say(fmt.Sprintf("Created image: %s", fingerprint))

	state.Put("imageFingerprint", fingerprint)

	return multistep.ActionContinue
}

func (s *stepPublish) Cleanup(_state multistep.StateBag) {
	/* TODO: do we need something here? */
}
