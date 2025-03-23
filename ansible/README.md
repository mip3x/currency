# Ansible
`Ansible` block documentation

## Содержание
1. [Общие настройки](#общие-настройки)
2. [Плейбуки](#плейбуки)
    - [Плейбук `main`](#плейбук-main)
    - [Плейбук `install-package`](#плейбук-install-package)
    - [Плейбук `establish-connection`](#плейбук-establish-connection)
3. [Роли](#роли)
    - [Роль `establish-connection`](#роль-establish-connection)

---

## Общие настройки
![`inventory.yml`](./inventory.yml) содержит адреса хостов и порты, по которым происходит подключение к этим хостам. `master` описывает главную ноду, `workers` - все остальные ноды:
```yaml
all:
  children:
    master:
      hosts:
        master-node:
	  # описание мастер-ноды
    workers:
      hosts:
        worker-node1:
	  # описание worker-ноды (1)
        worker-node2:
	  # описание worker-ноды (2)
```

Далее была сгенерирована пара `ssh`-ключей:
```bash
ssh-keygen -t ed25519
```

Был написан конфиг ![`ansible.cfg`](./ansible.cfg):
```toml
[defaults]
inventory               = inventory.yml
roles_path              = roles
remote_user             = root
private_key_file        = ~/.ssh/ansible_ssh_key

[ssh_connection]
host_key_checking       = False
```

- `inventory` - путь до файла `inventory`
- `roles_path` - путь до директории, в которой находятся файлы ролей `ansible`
- `remote_user` - пользователь, через которого будет производиться подключение по `ssh`
- `private_key_file` - путь до приватного ключа `ssh`
- `host_key_checking` - параметр, определяющий, будет ли производиться проверка хоста, к которому происходит подключение по `ssh`. Выставлен в `False`, чтобы избежать появления запроса на добавление хоста в `known_hosts`

---

# Плейбуки

## Плейбук `main`
Импортирует вспомогательные плейбуки и передаёт параметры, если требуется

---

## Плейбук `install-package`
Устанавливает заданный пакет (параметр `package_name`) на заданный хост (параметр `target_host`). Требуется указание ключа `--ask-become-pass` (`-K`), так как установка пакета требует прав администратора

---

## Плейбук `establish-connection`
Запуск роли `establish-connection`

---

# Роли

## Роль `establish-connection`
Эта роль добавляет `ssh`-ключ в `authorized_keys` всех хостов типа `workers`: это необходимо для осуществления авторизации по `ssh`-ключу. (Аналог `ssh-copy-id`) Если плейбук, использующий эту роль запускается *впервые*, то должен быть использован флаг `--ask-pass` (`-k`), так как из-за отсутствия `ssh`-ключей авторизация по паролю является единственный доступной возможностью авторизоваться. Более подробная документация по роли: ![ссылка](./roles/establish-connection/README.md)
