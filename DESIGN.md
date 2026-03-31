# labctl Design Document
## Problem Statement
Cloud images promise a fast path to a ready-to-use Linux environment. In reality geting one to work is a multi-step process involving many cli tools, creating configuration files and managing/generating copies and seeding images.Doing it in a terminal, linux' natural environment can be even more frustrating: 
- It doesn't suport copy and paste the same way a Desktop environment does.
- Everything needs to be typed; commands, path to directories, the name of the images.
- The configuration files are very strict with formatting, and wont accept `\t` character, trailing spaces or missing extra lines in the end. 
- Typos are bound to happen, wrong argument, missing `\` in multiline commands, wrong property in the config file
All of that potentially resulting in crashing vms, broken builds, missing credentials misconfigured network forwarding/bridging.In other way a huge waste of time.

## Target Users
This tool is not meant for a beginner or someone new to linux, for these users the struggle helps more than hurts, this tool is meant instead for users confortable with the linux environment, who knows how to install a distro, are very familiarized with the Linux command line interface: It professionals; developers, system administrators, devops engineers, system developers and security researchers looking for a straight forward workflow to help them quickly spin VMs precisely configured, manageable and ready for use rather in their preparation for a certification exam, as a lab for practice or as a test ground for a new access policy, kernel feature...


## Goals
Once a base image is chosen and downloaded, a user with `labctl` installed should be able to go from nothing to a running, SSH-accessible VM in under 2 minutes, without needing to remember the command to create the seed.iso, installing and using yaml linters to check for formatting errors in the `user-data` or deal with multi line argument lists with `virt-install`. Leveraging the already incredible potential and convenience of the `cloud-init` interface.
