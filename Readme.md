# Alfred

Alfred is a DevOps Assistant for helping devops engineers to work more efficient

### Install
You can download the binary file from the Release page 

### Usage
Download the binary file from the release page then copy the **config.sample.json** to **config.json** and set the configuration or use **alfred config** command to set configuration

```
root@example:/$ alfred
Alfred is an assistant for devops engineers

Usage:
  alfred [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Set configuration for Alfred
  gitlab      Gitlab Operations
  help        Help about any command

Flags:
  -h, --help     help for alfred

Use "alfred [command] --help" for more information about a command.
```

#### Gitlab Operations
Use **alfred gitlab --help** command to see what you can do with alfred
```
root@example:/$ alfred gitlab --help

Gitlab Operations

Usage:
  alfred gitlab [flags]
  alfred gitlab [command]

Available Commands:
  addEnv                 Add environment variables from CLI to a specific project in gitlab
  addEnvFromFile         Add environment variables from a file to a specific project in gitlab
  deleteEnv              Delete environment variables from CLI or from a file to a specific project in gitlab
  generateDockerfile     A brief description of your command
  generateEnvsSpringBoot Generate Environment Variables for Java Spring Boot Project

Flags:
  -h, --help             help for gitlab
  -p, --project string   Gitlab Project Name

Use "alfred gitlab [command] --help" for more information about a command.


```

### Author
Pooria Shahi (PooriaPro@gmail.com)

### License
Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0