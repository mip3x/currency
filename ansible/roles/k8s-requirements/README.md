k8s-requirements
================

This role prepares environment for `cri-o` and `k8s` utils installation. It installs requirements (packages) and adds `apt` repository of `k8s`

Variables `defaults`
--------------------

- `k8s_version` - version of `k8s`
- `k8s_public_signing_key` - `url` to `release`-key of `apt` `k8s` repository
- `k8s_repo_src_string` - `apt` `k8s` signing key
- `k8s_filename` - name of the `k8s` file in the list `/etc/apt/sources.list.d`

Variables `vars`
----------------

- `keyrings_directory_path` - path to `keyrings` directory
- `k8s_dest_signing_key_path` - path to `gpg`-key on the remote host

Example of `playbook`
---------------------

```yaml
- hosts: k8s
  roles:
    - { role: username.k8s-requirements }
```

Лицензия
--------

MIT
