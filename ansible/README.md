# Ansible
`Ansible` block documentation

## Contents
1. [General settings](#general-settings)
2. [Playbooks](#playbooks)
    - [Playbook `main`](#playbook-main)
    - [Playbook `install-package`](#playbook-install-package)
    - [Playbook `establish-connection`](#playbook-establish-connection)
3. [Roles](#roles)
    - [Role `establish-connection`](#role-establish-connection)

---

## General settings
![`inventory.yml`](./inventory.yml) contains the host addresses and the ports used for connecting to these hosts. The `master` group defines the primary node, while `workers` includes all the other nodes:
```yaml
all:
  children:
    master:
      hosts:
        master-node:
      # description of the master node
    workers:
      hosts:
        worker-node1:
	  # description of worker-node (1)
        worker-node2:
	  # description of worker-node (2)
```

A pair of SSH keys was generated:
```bash
ssh-keygen -t ed25519
```

A configuration file ![`ansible.cfg`](./ansible.cfg) was written:
```toml
[defaults]
inventory               = inventory.yml
roles_path              = roles
remote_user             = root
private_key_file        = ~/.ssh/ansible_ssh_key

[ssh_connection]
host_key_checking       = False
```

- `inventory` - path to the `inventory` file
- `roles_path` - path to the directory containing the Ansible role files
- `remote_user` - the user used to establish SSH connections
- `private_key_file` - path to the SSH private key
- `host_key_checking` - parameter that determines whether the SSH connection verifies the host key. Set to `False` to avoid prompts for adding new hosts to `known_hosts`

---

# Playbooks

## Playbook `main`
Imports auxiliary playbooks and passes parameters when needed.


---

## Playbook `install-package`
Installs the specified package (parameter `package_name`) on the designated host (parameter `target_host`). The `--ask-become-pass` (`-K`) flag is required because installing a package requires administrative privileges

---

## Playbook `establish-connection`
Executes the `establish-connection` role

---

# Roles

## Role `establish-connection`
This role adds an SSH key to the `authorized_keys` file on all hosts in the workers group. This is necessary for SSH key authentication (similar to `ssh-copy-id`). If the playbook that uses this role is run for the first time, the `--ask-pass` (`-k`) flag should be used, because in the absence of SSH keys, password authentication is the only available method. For more detailed documentation on this role, see: ![link](./roles/establish-connection/README.md)
