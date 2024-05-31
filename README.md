TonUtils Go - Библиотека для взаимодействия с сетью TON (The Open Network)
=================================================================

Это пример кода на Go, который использует библиотеку TonUtils для взаимодействия с сетью TON. В этом примере показано, как подключиться к mainnet lite серверу, создать кошелек, получить данные о коллекции NFT и отправить транзакцию на минт новых NFT токенов.

Перед запуском убедитесь, что у вас установлен Go и что вы скачали и установили библиотеку TonUtils Go.

Установка
----------

1.  Установите Go: <https://golang.org/doc/install>
2.  Сделайте клон данного репозитория
3.  Переведите все монеты на кошелёк Tonkeeper (Chrome Extension) V3R
4.  Отредактируйте все файлы скриптов

Использование
-------------
1. go run connect_lite.go
2. go run nft_mint.go (перед запуском данной команды не забудьте collectionAddr := address.MustParseAddr("адрес созданной коллекции"))
3. Проверить вывод по адресу коллекции

Важно!
------
Следите за индексами файлов. И за ссылками. В файле connect_lite.go `collectionContent := nft.ContentOffchain{URI: "https://harlequin-decent-hoverfly-340.mypinata.cloud/ipfs/QmdZP1UYxHgNLB8f4fgyaxQzsagPqky3psQxcfLQx4iGWr"}` данная строка ведёт на файл collection.json (джсон коллекции), а `https://harlequin-decent-hoverfly-340.mypinata.cloud/ipfs/QmUYXMgu6jXEcVcZpSXxYv6LWtbz6PcaTBVZzVuaiqpZnm/` ведёт на папку со всеми метаданными коллекции. Все метаданные должны строго по индексу начинаться с 1

Вам необходимо изменить данные строки с вашими файлами метаданных.

Дополнительная информация
------------------------

*   Документация TonUtils Go: <https://godoc.org/github.com/xssnick/tonutils-go>
*   Официальный сайт TON: <https://ton.org/>
*   Документация TON: <https://docs.ton.org/>
