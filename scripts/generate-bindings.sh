#!/bin/bash
set -e

# Имя контракта (без расширения)
CONTRACT_NAME="TelegramRussianRoulette"

# Папка, где лежат ABI-файлы
ABI_DIR="./abi"

# Папка для сгенерированного Go-файла
OUT_DIR="./bindings"

# Создаём выходную папку, если её нет
mkdir -p ${OUT_DIR}

# Запуск abigen через Docker (теперь используем правильный образ)
docker run --rm -v "$(pwd)":/src ethereum/client-go:alltools-latest abigen \
    --abi /src/${ABI_DIR}/${CONTRACT_NAME}.abi \
    --pkg bindings \
    --out /src/${OUT_DIR}/${CONTRACT_NAME}.go

echo "Биндинги для ${CONTRACT_NAME} сгенерированы в ${OUT_DIR}/${CONTRACT_NAME}.go"
