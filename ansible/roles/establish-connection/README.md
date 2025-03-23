establish-connection
====================

Эта роль добавляет сгенерированный публичный ключ в `authorized_keys` на нодах `workers` для дальнейшего подключения по `ssh`.

Требования
----------

- Необходимо запускать плэйбук с ключом `--ask-pass` (`-k`), так как `ssh`-ключ пока не добавлен в `authorized_keys`

Переменные
----------

- `ssh_public_key_path` - путь до публичного ключа `ssh`

Зависимости
-----------

- Необходимо установить пакет `sshpass` на хост, с которого происходит запуск (`apt install sshpass` для `Ubuntu 22.04`)

Пример `playbook`-а
-------------------

    - hosts: workers
      roles:
         - { role: username.establish-connection }

Лицензия
--------

MIT
