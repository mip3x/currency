establish-connection
====================

This role adds the generated public key to the `authorized_keys` on `workers` nodes for subsequent SSH connections

Requirements
------------

- The playbook must be run with the `--ask-pass` (`-k`) flag since the SSH key is not yet added to `authorized_keys`

Variables
---------

- `ssh_public_key_path` - the path to the SSH public key

Dependencies
------------

- The package `sshpass` must be installed on the host from which the playbook is run (e.g., `apt install sshpass` for Ubuntu 22.04)

Example Playbook
----------------

```yaml
- hosts: workers
  roles:
    - { role: username.establish-connection }
```

License
-------

MIT
