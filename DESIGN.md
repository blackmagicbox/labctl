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

## Command Interface Design
1. Command signatures
```shell
labctl new                    # full wizard, including image selection
labctl new <image-name>       # skips image selection, image must exist
labctl new <path/to/image>    # skips image selection, registers new image
```

If the <image-name> image doesn't exit it triggers an error:
```shell
$ labctl new rocky-99
error: image 'rocky-99' not found. Run 'labctl images' to see available images.
```

2. Choose distro
The first step of the wizard is selecting (or assigning if the user passes a path to the `new` command) the correct distro base 
    1. Distro selection/ Distro Assignment
    ```shell
    distro:
    -> arch
    -> debian/ubuntu
    -> REHL
    ```
    2. Image selection *(eg: REHL)*
    ```shell
    images:
    -> CentOS-stream-10
    -> Fedora-43
    -> Rocky-9
    -> Rocky-10
    ```
3. Wizzard flow
Configuration attributes:
    1. VM name         (default: rocky-9-20260331) # distro-version-date
    2. Hostname        (default: rocky-$USER) # distro-username
    3. Username        (default: $USER)
    4. Auth method     (SSH key ~/.ssh/id_ed2XXXX.pub detected ✓)
    5. Packages        (multi-select: vim, curl, htop, git...)
    6. Disk size       (default: 20G)
    7. Memory          (default: 2048MB)
    8. CPUs            (default: 2)

```shell
$ labctl new Rocky-9

    ✓ Please provide a **name** for the new vm: rocky-9-20012014 
    ✓ Please provide a **hostname**: rocky-user
    ✓ Please provide a **username**: user
    ? Authentication method > SSH / password
    # if the user chooses password...
    ✓ password: ******* 
    ? retype the password: _
    # if the user chooses SSH
    ? Enter the SSH Key:  (~/.ssh/id_ed2XXXX.pub detected ✓) / other... 
    [...]
    ✓ CPUs: 2

    VM 'rocky-9-20260331' created successfully.
    ? Start the VM now? > yes / no
    # if yes
    ⠋ Starting rocky-9-20260331...
    ✓ VM is running.
      IP:  192.168.122.14
      SSH: ssh rafael@192.168.122.14
    # if no
    You can start it later with:
    $ labctl start rocky-9-20260331   
```
Behaviour:
- Hitting enter on any prompt accepts the default
- If no SSH key is detected, auth method defaults to password
- If a path is passed, distro/image selection is skipped
- Image is copied to labctl's managed directory and resized on creation
- After configuration, user is prompted to start the VM immediately
- Regardless of choice, labctl prints the start command for future reference

# Project structure
```text
labctl/
├── main.go                 # entry point, calls cmd.Execute(), nothing else
├── go.mod
├── go.sum
├── DESIGN.md
├── README.md
├── cmd/                    # one file per command, Cobra lives here
│   ├── root.go             # root command, global flags, initialises Cobra
│   ├── new.go              # labctl new
│   ├── start.go            # labctl start
│   ├── stop.go             # labctl stop
│   ├── delete.go           # labctl delete
│   └── images.go           # labctl images
└── internal/               # private packages, not importable externally
    ├── vm/                 # talks to virsh/virt-install, manages VM lifecycle
    ├── image/              # copies, resizes, registers base images
    ├── cloudconfig/        # generates user-data and meta-data from templates
    ├── iso/                # runs cloud-localds, produces seed.iso
    ├── config/             # reads/writes ~/.config/labctl/config
    └── tui/                # Bubble Tea wizard and interactive lists
```
