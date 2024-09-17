# Firewall Rule Manager Engine

This Go project allows you to manage Windows firewall rules for specific applications (like Chrome or Firefox) by either blocking or allowing access to a specific IP address.

## Prerequisites

- **Go** must be installed on your system. You can download it from [here](https://golang.org/dl/).
- This project assumes you're running on **Windows** because it uses `netsh` to modify firewall rules.
- Administrator privileges are required to modify firewall rules.

## Project Structure
engine/ ├── main.go # Handles firewall modification └── run.go # Defines application paths and inputs


## Running the Program

To run the program, follow these steps:

1. Open a terminal (CMD or PowerShell) with **administrator privileges** (since firewall changes require admin access).

2. Navigate to the project directory:

   ```
   cd path\to\your\project\engine
   ```
3. Execute the Go program with the following format:

   ```
   go run run.go <app-name> <ip-address> <action>
   ```
   app-name: The name of the application (chrome or firefox).
   ip-address: The IP address you want to block or allow.
   action: The action you want to perform (allow or deny).

## Example
To block the IP 192.168.1.100 for Chrome, run the following command:

```
go run run.go chrome 192.168.1.100 deny
```

To allow the IP 192.168.1.101 for Firefox, run this command:

```
go run run.go firefox 192.168.1.101 allow
```