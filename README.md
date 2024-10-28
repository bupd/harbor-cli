
![harbor-3](https://github.com/goharbor/harbor-cli/assets/70086051/835ab686-1cce-4ac7-bc57-05a35c2b73cc)

**Welcome to the Harbor CLI project! This powerful command-line tool facilitates seamless interaction with the Harbor container registry. It simplifies various tasks such as creating, updating, and managing projects, registries, and other resources in Harbor.**

# Project Scope 🧪

The Harbor CLI is designed to enhance your interaction with the Harbor container registry. Built on Golang, it offers a user-friendly interface to perform various tasks related to projects, registries, and more. Whether you're creating, updating, or managing resources, the Harbor CLI streamlines your workflow efficiently.

# Project Features 🤯

 🔹 Get details about projects, registries, repositories and more <br>
 🔹 Create new projects, registries, and other resources <br>
 🔹 Delete projects, registries, and other resources <br>
 🔹 Run commands with various flags for enhanced functionality <br>
 🔹 More features coming soon... 🚧

# Example Commands💡

```bash
➜ harbor --help
Official Harbor CLI

Usage:
  harbor [command]

Examples:

// Base command:
harbor

// Display help about the command:
harbor help


Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  login       Log in to Harbor registry
  project     Manage projects and assign resources to them
  registry    Manage registries
  repo        Manage repositories
  user        Manage users
  version     Version of Harbor CLI

Flags:
      --config string          config file (default is $HOME/.harbor/config.yaml) (default "/home/bishal/.harbor/config.yaml")
  -h, --help                   help for harbor
  -o, --output-format string   Output format. One of: json|yaml
  -v, --verbose                verbose output

Use "harbor [command] --help" for more information about a command.
```

#### Log in to Harbor Registry

```bash
harbor login demo.goharbor.io -u harbor-cli -p Harbor12345
```

#### Create a New Project

```bash
harbor project create
```

#### List all Projects

```bash
harbor project list

# output
┌──────────────────────────────────────────────────────────────────────────────────────────┐
│  Project Name  Access Level  Type          Repo Count    Creation Time                   │
│ ──────────────────────────────────────────────────────────────────────────────────────── │
│  library       public        project       0             1 hour ago                      │
└──────────────────────────────────────────────────────────────────────────────────────────┘
```

#### List all Repository in a Project

```bash
harbor repo list

# output
┌────────────────────────────────────────────────────────────────────────────────────────┐
│  Name                      Artifacts     Pulls         Last Modified Time              │
│ ────────────────────────────────────────────────────────────────────────────────────── │
│  library/harbor-cli        1             0             0 minute ago                    │
└────────────────────────────────────────────────────────────────────────────────────────┘
```

# Supported Platforms

Platform | Status
--|--
Linux | ✅
macOS | ✅
Windows | ✅

# Installation


## Linux and MacOS

Homebrew is the recommended way to install Harbor CLI on MacOS and Linux.


## Windows

```shell

winget install harbor

```



# Build From Source

Make sure you have latest [Dagger](https://docs.dagger.io/) installed in your system.

#### Using Dagger
```bash
git clone https://github.com/goharbor/harbor-cli.git && cd harbor-cli
dagger call build-dev --platform darwin/arm64 export --path=./harbor-cli
./harbor-dev --help
```

If golang is installed in your system, you can also build the project using the following commands:

```bash
git clone https://github.com/goharbor/harbor-cli.git
go build -o harbor-cli cmd/harbor/main.go
```


# Lint

Make sure you have latest [Dagger](https://docs.dagger.io/) installed in your system.

```bash
dagger call lint
```

#### Generate Lint Report
```bash
dagger call lint-report export --path=./LintReport.txt
```


# Community

* **Twitter:** [@project_harbor](https://twitter.com/project_harbor)
* **User Group:** Join Harbor user email group: [harbor-users@lists.cncf.io](https://lists.cncf.io/g/harbor-users) to get update of Harbor's news, features, releases, or to provide suggestion and feedback.
* **Developer Group:** Join Harbor developer group: [harbor-dev@lists.cncf.io](https://lists.cncf.io/g/harbor-dev) for discussion on Harbor development and contribution.
* **Slack:** Join Harbor's community for discussion and ask questions: [Cloud Native Computing Foundation](https://slack.cncf.io/), channel: [#harbor](https://cloud-native.slack.com/messages/harbor/), [#harbor-dev](https://cloud-native.slack.com/messages/harbor-dev/) and [#harbor-cli](https://cloud-native.slack.com/messages/harbor-cli/).

# License

This project is licensed under the Apache 2.0 License. See the [LICENSE](https://github.com/goharbor/harbor-cli/blob/main/LICENSE) file for details.

# Acknowledgements

This project is maintained by the Harbor community. We thank all our contributors and users for their support.

# ❤️ Show your support

For any questions or issues, please open an issue on our [GitHub Issues](https://github.com/goharbor/harbor-cli/issues) page.<br>
Give a ⭐ if this project helped you, Thank YOU!
