# Contributing to System Onboarding Auditor

First off, thank you for taking the time to contribute! This tool is designed to make our team more autonomous and reduce manual onboarding toil.

## 🚀 How to Add a New System Check
To keep the tool "pickable" and maintainable, we use a simple map-based configuration. 

### Step 1: Add your tool to the map
Open `main.go` and find the `tools` map. Add the binary name and the flag used to check its version.

```go
tools := map[string][]string{
    "Git":      {"version"},
    "Docker":   {"--version"},
    "Go":       {"version"},
    "kubectl":  {"version", "--client"}, // New check added here
}

Step 2: Test your changes
Run the auditor locally to ensure your new check doesn't hang the system:

go run main.go

🛠 Standards for Contributions
To ensure the tool remains resilient (as per our SecDevOps standards):

Standard Library Only: Avoid adding external dependencies to keep the binary lightweight and secure.

Context-Aware: All checks must obey the context.WithTimeout to prevent "zombie" processes.

Clean Output: Ensure the tool name fits within the %-10s padding for visual alignment.

🐛 Reporting Issues
If a check is giving a "false negative" (reporting missing when installed), please open an issue with:

Your Operating System (Windows/Linux/macOS)

The output of echo $PATH

“DevOps is an opportunity in leadership. Let’s build tools that empower everyone.”
