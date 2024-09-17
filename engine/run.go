package main

import (
    "log"
    "os"
)

// GetAppDetails returns the app path, IP, action, and app name
func GetAppDetails() (string, string, string, string) {
    // Predefined list of known file paths
    appPaths := map[string]string{
        "chrome":  `C:\Program Files\Google\Chrome\Application\chrome.exe`,
        "firefox": `C:\Program Files\Mozilla Firefox\firefox.exe`,
    }

    // Ensure we have three arguments: <app-name>, <ip-address>, <action>
    if len(os.Args) != 4 {
        log.Fatalf("Usage: %s <app-name> <ip-address> <action(allow|deny)>\n", os.Args[0])
    }

    appName := os.Args[1]   // The key for the known app (e.g., "chrome")
    ipToBlock := os.Args[2] // The IP address to block/allow
    action := os.Args[3]    // Action: "allow" or "deny"

    // Check if the app path exists in our predefined list
    appPath, exists := appPaths[appName]
    if !exists {
        log.Fatalf("Application '%s' not found in known paths. Available options are: chrome, firefox\n", appName)
    }

    // Validate action
    if action != "allow" && action != "deny" {
        log.Fatalf("Invalid action: %s. Use 'allow' or 'deny'.\n", action)
    }

    return appPath, ipToBlock, action, appName
}
