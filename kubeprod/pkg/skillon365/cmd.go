/*
 * Bitnami Kubernetes Production Runtime - A collection of services that makes it
 * easy to run production workloads in Kubernetes.
 *
 * Copyright 2020 Bitnami
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package skillon365

import (
	"os"

	azcli "github.com/Azure/go-autorest/autorest/azure/cli"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	kubeprodcmd "github.com/bitnami/kube-prod-runtime/kubeprod/cmd"
)

const (
	flagEmail            = "email"
	flagDNSSuffix        = "dns-zone"
	flagAuthzDomain      = "authz-domain"
	flagKeycloakPassword = "keycloak-password"
	flagKeycloakGroup    = "keycloak-group"
	flagKeycloakRealm    = "keycloak-realm"
	flagCloudflareToken  = "cloudflare-token"
	flagSubID            = "subscription-id"
	flagTenantID         = "tenant-id"
)

func defaultSubscription() *azcli.Subscription {
	path, err := azcli.ProfilePath()
	if err != nil {
		log.Debugf("Unable to find azure-cli profile: %v", err)
		return nil
	}
	profile, err := azcli.LoadProfile(path)
	if err != nil {
		log.Debugf("Unable to load azure-cli profile: %v", err)
		return nil
	}

	for _, s := range profile.Subscriptions {
		if s.IsDefault {
			return &s
		}
	}
	return nil
}

var skillon365Cmd = &cobra.Command{
	Use:   "skillon365",
	Short: "Install Bitnami Production Runtime for SkillON365 Kubernetes cluster",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := kubeprodcmd.NewInstallSubcommand(cmd, "skillon365", &SkillON365Config{flags: cmd.Flags()})
		if err != nil {
			return err
		}

		return c.Run(cmd.OutOrStdout())
	},
}

func init() {
	kubeprodcmd.InstallCmd.AddCommand(skillon365Cmd)

	var defSubID, defTenantID string
	if defSub := defaultSubscription(); defSub != nil {
		defSubID = defSub.ID
		defTenantID = defSub.TenantID
	}

	skillon365Cmd.PersistentFlags().String(flagEmail, os.Getenv("EMAIL"), "Contact email for cluster admin")
	skillon365Cmd.PersistentFlags().String(flagAuthzDomain, "*", "Restrict authorized users to this email domain.  Default '*' allows all users in Azure tenant.")
	skillon365Cmd.PersistentFlags().String(flagSubID, defSubID, "Azure subscription ID")
	skillon365Cmd.PersistentFlags().String(flagTenantID, defTenantID, "Azure tenant ID")
	skillon365Cmd.PersistentFlags().String(flagDNSSuffix, "", "External DNS zone for public endpoints")
	skillon365Cmd.PersistentFlags().String(flagKeycloakGroup, "", "Restrict authorized users to this Keycloak group")
	skillon365Cmd.PersistentFlags().String(flagKeycloakRealm, "SkillON365", "Realm for kubeprod services")
	skillon365Cmd.PersistentFlags().String(flagKeycloakPassword, "", "Password for Keycloak admin user")
	skillon365Cmd.PersistentFlags().String(flagCloudflareToken, "", "Cloudflare API token")
	skillon365Cmd.MarkPersistentFlagRequired(flagKeycloakGroup)
	skillon365Cmd.MarkPersistentFlagRequired(flagKeycloakPassword)
	skillon365Cmd.MarkPersistentFlagRequired(flagCloudflareToken)
}
