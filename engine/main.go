package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Get application path, IP, action, and app name from run.go
	appPath, ipToBlock, action, appName := GetAppDetails()

	// Validate action (if not already validated in run.go)
	if action != "allow" && action != "deny" {
		log.Fatalf("Invalid action: %s. Use 'allow' or 'deny'.\n", action)
	}

	// Modify the firewall rule
	if err := modifyFirewallRule(appPath, ipToBlock, action, appName); err != nil {
		log.Fatalf("Error modifying firewall rule: %v\n", err)
	}

	fmt.Printf("Firewall rule %sed successfully for IP %s and app %s.\n", action, ipToBlock, appName)
}

func modifyFirewallRule(appPath, ipToBlock, action, appName string) error {
	// Translate action to netsh action command (block or allow)
	netshAction := "block"
	if action == "allow" {
		netshAction = "allow"
	}

	// Define the netsh command for modifying the firewall
	ruleName := fmt.Sprintf("%s IP %s for %s", action, ipToBlock, appName)
	cmd := exec.Command("netsh", "advfirewall", "firewall", "add", "rule",
		"name="+ruleName,
		"dir=out",
		"action="+netshAction,
		"program="+appPath,
		"remoteip="+ipToBlock,
		"enable=yes")

	// Execute the command and capture the output and any errors
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to execute command: %w\nOutput: %s", err, string(output))
	}

	return nil
}
