# k8s-dashboard-plugin

Branched out of [181192/ops-cli](https://github.com/181192/ops-cli) to provide only the dashboard commands

Small utility for portforwarding into applications running in k8s and open default browser.

## Features

- Dynamic port allocation if target port is not available
- Opens default browser on all plattforms for handling URL
  - Darwin - `open`
  - Linux - `xdg-open`
  - Windows - `start`
- Configuration file to add own dashboard to populate the commands automatically
- Auto completion for Bash, Fish, Zsh and PowerShell (also for user provided dashboards via config file)
