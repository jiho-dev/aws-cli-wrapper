package cmd

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/jiho-dev/aws-cli-wrapper/config"
	"github.com/spf13/cobra"
)

var apiGroupSubCmds = map[string][]string{
	"admin-vpc": []string{
		"allocate-random-ip-pool",
		"blackpearl-health",
		"create-public-ipv4-pool",
		"delete-public-ipv4-pool",
		"deregister-public-ipv4-pool",
		"disable-public-ipv4-pool",
		"disassociate-public-ip",
		"enable-public-ipv4-pool",
		"list-address-associations",
		"list-blackpearl",
		"list-network-acl",
		"list-network-interface",
		"list-public-ips",
		"list-public-ipv4-pool",
		"list-route-table",
		"list-security-group",
		"list-vrouters",
		"register-public-ipv4-pool",
		"release-ip-pool",
		"release-public-ip",
		"request-ip-pool",
		"show-dataversion",
		"show-flowlog",
		"show-network-interface",
		"show-papyrus-flowlog",
		"show-papyrus-summary",
		"show-revision",
		"show-snat",
		"show-summary",
		"show-vrevision",
		"show-vrouter-flowlog",
		"show-vrouter-flow",
		"show-vrouter-network-acl",
		"show-vrouter-network-interface",
		"show-vrouter-port",
		"show-vrouter-route",
		"show-vrouter-security-group",
		"show-vrouter-subnet",
		"show-vrouter-summary",
		"show-vrouter-table",
		"update-network-interface",
	},
	"ec2": []string{
		/*
			"describe-account-attributes",
			"describe-addresses",
			"describe-addresses-attribute",
			"describe-aggregate-id-format",
			"describe-auto-scaling-group-associations",
			"describe-availability-zones",
			"describe-bundle-tasks",
			"describe-byoip-cidrs",
			"describe-capacity-reservations",
			"describe-carrier-gateways",
			"describe-classic-link-instances",
			"describe-client-vpn-authorization-rules",
			"describe-client-vpn-connections",
			"describe-client-vpn-endpoints",
			"describe-client-vpn-routes",
			"describe-client-vpn-target-networks",
			"describe-coip-pools",
			"describe-compute-hosts",
			"describe-conversion-tasks",
			"describe-customer-gateways",
			"describe-dhcp-options",
			"describe-egress-only-internet-gateways",
			"describe-elastic-gpus",
			"describe-export-image-tasks",
			"describe-export-tasks",
			"describe-fast-snapshot-restores",
			"describe-fleet-history",
			"describe-fleet-instances",
			"describe-fleets",
			"describe-flow-logs",
			"describe-fpga-image-attribute",
			"describe-fpga-images",
			"describe-host-reservation-offerings",
			"describe-host-reservations",
			"describe-hosts",
			"describe-iam-instance-profile-associations",
			"describe-id-format",
			"describe-identity-id-format",
		*/
		"describe-image-attribute",
		"describe-images",
		"describe-import-image-tasks",
		"describe-import-snapshot-tasks",
		"describe-instance-attribute",
		"describe-instance-credit-specifications",
		"describe-instance-event-notification-attributes",
		"describe-instance-event-windows",
		"describe-instance-placements",
		"describe-instance-status",
		"describe-instance-type-offerings",
		"describe-instance-types",
		"describe-instances",
		"describe-internet-gateways",
		//"describe-ipv6-pools",
		"describe-key-pairs",
		/*
			"describe-launch-template-versions",
			"describe-launch-templates",
			"describe-local-gateway-route-table-virtual-interface-group-associations",
			"describe-local-gateway-route-table-vpc-associations",
			"describe-local-gateway-route-tables",
			"describe-local-gateway-virtual-interface-groups",
			"describe-local-gateway-virtual-interfaces",
			"describe-local-gateways",
			"describe-managed-prefix-lists",
			"describe-moving-addresses",
		*/
		"describe-nat-gateways",
		"describe-network-acls",
		"describe-network-insights-analyses",
		"describe-network-insights-paths",
		"describe-network-interface-attribute",
		"describe-network-interface-permissions",
		"describe-network-interfaces",
		"describe-network-interfaces-spc",
		"describe-notifications",
		"describe-placement-groups",
		"describe-prefix-lists",
		//"describe-principal-id-format",
		"describe-public-ipv4-pools",
		/*
			"describe-regions",
			"describe-replace-root-volume-tasks",
			"describe-reserved-instances",
			"describe-reserved-instances-listings",
			"describe-reserved-instances-modifications",
			"describe-reserved-instances-offerings",
			"describe-resources",
		*/
		"describe-route-tables",
		/*
			"describe-scheduled-instance-availability",
			"describe-scheduled-instances",
		*/
		"describe-security-group-references",
		"describe-security-group-rules",
		"describe-security-groups",
		/*
			"describe-service-components",
			"describe-snapshot-attribute",
			"describe-snapshots",
			"describe-spot-datafeed-subscription",
			"describe-spot-fleet-instances",
			"describe-spot-fleet-request-history",
			"describe-spot-fleet-requests",
			"describe-spot-instance-requests",
			"describe-spot-price-history",
			"describe-stale-security-groups",
			"describe-store-image-tasks",
		*/
		"describe-subnets",
		"describe-tags",
		/*
			"describe-traffic-mirror-filters",
			"describe-traffic-mirror-sessions",
			"describe-traffic-mirror-targets",
			"describe-transit-gateway-attachments",
			"describe-transit-gateway-connect-peers",
			"describe-transit-gateway-connects",
			"describe-transit-gateway-multicast-domains",
			"describe-transit-gateway-peering-attachments",
			"describe-transit-gateway-route-tables",
			"describe-transit-gateway-vpc-attachments",
			"describe-transit-gateways",
			"describe-trunk-interface-associations",
		*/
		"describe-volume-attribute",
		"describe-volume-status",
		"describe-volumes",
		"describe-volumes-modifications",
		"describe-vpc-attribute",
		"describe-vpc-classic-link",
		"describe-vpc-classic-link-dns-support",
		"describe-vpc-endpoint-connection-notifications",
		"describe-vpc-endpoint-connections",
		"describe-vpc-endpoint-service-configurations",
		"describe-vpc-endpoint-service-permissions",
		"describe-vpc-endpoint-services",
		"describe-vpc-endpoints",
		"describe-vpc-peering-connections",
		"describe-vpcs",
		/*
			"describe-vpn-connections",
			"describe-vpn-gateways",
			"get-associated-enclave-certificate-iam-roles",
			"get-associated-ipv6-pool-cidrs",
			"get-capacity-reservation-usage",
			"get-coip-pool-usage",
		*/
		"get-console-output",
		"get-console-screenshot",
		/*
			"get-default-credit-specification",
			"get-ebs-default-kms-key-id",
			"get-ebs-encryption-by-default",
			"get-flow-logs-integration-template",
			"get-groups-for-capacity-reservation",
			"get-host-reservation-purchase-preview",
			"get-launch-template-data",
			"get-managed-prefix-list-associations",
			"get-managed-prefix-list-entries",
			"get-password-data",
			"get-reserved-instances-exchange-quote",
			"get-serial-console-access-status",
			"get-subnet-cidr-reservations",
			"get-transit-gateway-attachment-propagations",
			"get-transit-gateway-multicast-domain-associations",
			"get-transit-gateway-prefix-list-references",
			"get-transit-gateway-route-table-associations",
			"get-transit-gateway-route-table-propagations",
		*/
	},
}

func InitGenerateCmd(apiGroup string) *cobra.Command {
	var genCmd = &cobra.Command{
		Use:   "generate-cmds",
		Short: "Generate commands",
		//Hidden:                true,
		//DisableFlagsInUseLine: true,
		//ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
		//Args: cobra.ExactValidArgs(1),
		Run:               generateCmdMain,
		ValidArgsFunction: getApiArgs,
	}

	if apiGroup == CMD_ADMIN_VPC {
		addProfileCmd(genCmd)
	}

	return genCmd
}

func generateCmdMain(cobraCmd *cobra.Command, args []string) {
	var group string

	if cobraCmd.HasParent() {
		p := cobraCmd.Parent()
		group = p.Use
	}

	switch group {
	case CMD_ADMIN_VPC:
		generateAdminVpcCmdMain(cobraCmd, args)
	case CMD_EC2:
		generateEc2CmdMain(cobraCmd, args)
	}
}

func generateAdminVpcCmdMain(cobraCmd *cobra.Command, args []string) {
	flags := cobraCmd.Flags()
	flags.Bool(CMD_SHOW_HELP, true, "")
	apiGroup, _ := AcwConfig.ApiGroup[CMD_ADMIN_VPC]
	groupCmds, _ := apiGroupSubCmds[CMD_ADMIN_VPC]

	for _, api := range groupCmds {
		inCmds := []string{api}

		output, err := RunCmd(inCmds, nil, true, flags)
		if err != nil {
			if output != "" {
				fmt.Printf("Output: %s \n", output)
			}

			fmt.Printf("ERR: %s \n", err)
			continue
		}

		if output == "" {
			fmt.Printf("No Output\n")
			continue
		}

		output1 := ParseOutput(output, "Result")
		if output1 == "" {
			output1 = output
		}

		output2 := FormatJson(output1)
		if output2 == "" || output2 == "{}" {
			output2 = output1
		}

		fmt.Printf("%s\n", output2)

		newOpts := config.AcwConfigApiOpt{}
		newOpts.OutputField = "Result"

		oldOpts, ok := apiGroup[api]
		if ok {
			if oldOpts.OutputField != "" {
				newOpts.OutputField = oldOpts.OutputField
			}

			newOpts.Required = oldOpts.Required
		}

		args := strings.Split(output2, "\n")
		var seeParams bool
		for _, arg := range args {
			if strings.HasPrefix(arg, "Parameters:") {
				seeParams = true
				continue
			}

			if seeParams {
				required := strings.Contains(arg, "(required)")

				arg = strings.TrimSpace(arg)
				tmp := strings.Split(arg, " ")
				key := tmp[0]
				key = strings.TrimSpace(key)

				if required {
					if !Contains(newOpts.Required, key) {
						newOpts.Required = append(newOpts.Required, key)
					}
				} else if !Contains(newOpts.Required, key) {
					newOpts.Args = append(newOpts.Args, key)
				}
			}
		}

		apiGroup[api] = newOpts
	}

	confFile := path.Join(awsDir, acwConf)
	config.WriteConfig(AcwConfig, confFile)
}

func Contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func ContainPrefixs(prefix []string, searchterm string) bool {
	for _, p := range prefix {
		if strings.HasPrefix(searchterm, p) {
			return true
		}
	}

	return false
}

func getEc2Cmds() []string {
	var ec2Cmds []string

	var ec2CmdPrefixs = []string{"describe-", "get-"}
	//out := aws ec2 help | grep -E "o "

	out, err := ExecuteAwsCli("aws", "ec2", "help")
	if err != nil {
		fmt.Printf("Err: %s \n", err)
		return nil
	}

	//fmt.Printf("out: [%s] \n", out)
	cmds := strings.Split(out, "\n")
	for _, tmp := range cmds {
		tmp = strings.TrimSpace(tmp)
		if tmp == "" {
			continue
		}

		tmps := strings.Split(tmp, " ")
		if len(tmps) < 2 {
			continue
		}

		cmd := tmps[1]
		if !ContainPrefixs(ec2CmdPrefixs, cmd) {
			continue
		}

		if cmd == "describe-local-gateway-route-table-virtual-interface-group-associa-" {
			cmd = "describe-local-gateway-route-table-virtual-interface-group-associations"
		}

		//fmt.Printf("\"%s\", \n", cmd)
		ec2Cmds = append(ec2Cmds, cmd)
	}

	return ec2Cmds
}

func generateEc2CmdMain(cobraCmd *cobra.Command, args []string) {
	flags := cobraCmd.Flags()
	flags.Bool(CMD_SHOW_HELP, true, "")
	apiGroup, _ := AcwConfig.ApiGroup[CMD_EC2]
	groupCmds, _ := apiGroupSubCmds[CMD_EC2]

	//cmds := getEc2Cmds()
	//fmt.Printf("ec2Cmds: %+v \n", cmds)

	for _, api := range groupCmds {
		fmt.Printf("api: %s \n", api)
		out, err := ExecuteAwsCli("aws", "ec2", api, "help")
		if err != nil {
			fmt.Printf("Err: %s \n", err)
		}

		//fmt.Printf("api help: %s \n", out)

		newOpts := config.AcwConfigApiOpt{}
		newOpts.OutputField = "Output"

		oldOpts, ok := apiGroup[api]
		if ok {
			if oldOpts.OutputField != "" {
				newOpts.OutputField = oldOpts.OutputField
			}

			newOpts.Required = oldOpts.Required
		}

		args := strings.Split(out, "\n")
		var seeOpts, seeSyn bool
		for _, arg := range args {
			//fmt.Printf("arg: %s \n", arg)
			if strings.Contains(arg, "SYNOPSIS") {
				seeSyn = true
				continue
			} else if strings.Contains(arg, "OPTIONS") {
				seeOpts = true
			}

			if seeSyn && seeOpts {
				break
			}

			arg = strings.TrimSpace(arg)

			if strings.Contains(arg, "--dry-run") {
				continue
			} else if !strings.Contains(arg, "[--") {
				continue
			} else if !strings.Contains(arg, "<value>") {
				// XXX
				continue
			}

			if seeSyn {
				tmp := strings.Split(arg, " ")
				key := tmp[0][3:]

				if !Contains(newOpts.Required, key) {
					newOpts.Args = append(newOpts.Args, key)
				}
			}
		}

		apiGroup[api] = newOpts
	}

	confFile := path.Join(awsDir, acwConf)
	config.WriteConfig(AcwConfig, confFile)
}
