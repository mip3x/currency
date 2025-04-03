crio-installation
=================

This role installs `cri-o`

Variables `defaults`
--------------------

- `crio_version` - version of `cri-o`
- `crio_public_signing_key` - `url` to `release`-key of `apt` `cri-o` repository
- `crio_repo_src_string` - signing key of `apt` `cri-o` repository
- `crio_filename` - file name of `cri-o` in the list `/etc/apt/sources.list.d`

Variables `vars`
----------------

- `crio_dest_signing_key_path` - path to `gpg`-key on the remote host

Example of `playbook`
---------------------

```yaml
- hosts: k8s
  roles:
    - { role: username.crio-installation }
```

License
-------

MIT

