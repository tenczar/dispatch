///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/vmware/dispatch/pkg/dispatchcli/i18n"
	"github.com/vmware/dispatch/pkg/event-manager/gen/client/drivers"
	"github.com/vmware/dispatch/pkg/event-manager/gen/models"
)

var (
	updateEventDriverLong    = ""
	updateEventDriverExample = ""

	driverType string
)

// NewCmdUpdateEventDriver creates command responsible for updating dispatch function eventDriver.
func NewCmdUpdateEventDriver(out io.Writer, errOut io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "event-driver DRIVER_NAME [--driver-type DRIVER_TYPE] [--set KEY=VALUE] [--secret SECRET_NAME]",
		Short:   i18n.T("update event driver"),
		Long:    updateEventDriverLong,
		Example: updateEventDriverExample,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := createEventDriver(out, errOut, cmd, args)
			CheckErr(err)
		},
	}

	cmd.Flags().StringVar(&driverType, "driver-type", "", "driver type, currently supports 'vcenter'")
	cmd.Flags().StringVarP(&cmdFlagApplication, "application", "a", "", "associate with an application")
	cmd.Flags().StringArrayVarP(&createEventDriverConfig, "set", "s", []string{}, "set event driver configurations, default: empty")
	cmd.Flags().StringArrayVar(&createEventDriverSecrets, "secret", []string{}, "Configuration passed via secrets, can be specified multiple times or a comma-delimited string")
	return cmd
}

func updateEventDriver(out io.Writer, errOut io.Writer, cmd *cobra.Command, args []string) error {
	driverName := args[0]

	params := drivers.NewGetDriverParams()
	params.DriverName = driverName

	eventDriverResp, err := eventManagerClient().Drivers.GetDriver(params, GetAuthInfoWriter())
	if err != nil {
		return formatAPIError(err, "")
	}

	eventDriver := *eventDriverResp.Payload

	changed := false
	if cmd.Flags().Changed("driver-type") {
		eventDriver.Type = &driverType
		changed = true
	}

	if cmd.Flags().Changed("set") {
		driverConfig := models.DriverConfig{}
		for _, conf := range createEventDriverConfig {
			result := strings.Split(conf, "=")
			if len(result) != 2 {
				return formatCliError(errors.New("invalid input"), "Invalid Configuration Format, should be --set key=value")
			}
			driverConfig = append(driverConfig, &models.Config{
				Key:   result[0],
				Value: result[1],
			})
		}

		eventDriver.Config = driverConfig
		changed = true
	}

	if cmd.Flags().Changed("secret") {
		eventDriver.Secrets = createEventDriverSecrets
		changed = true
	}

	if !changed {
		fmt.Fprintf(out, "Nothing to update")
		return nil
	}

	return nil
}
