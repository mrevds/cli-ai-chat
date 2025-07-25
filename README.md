# 🤖 CLI AI Chat [![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8.svg?style=flat&logo=go)](https://golang.org/) [![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Консольный клиент для взаимодействия с AI-моделями через AIML API с поддержкой истории диалога и интуитивным управлением.

## ✨ Key Features
- 🚀 **Поддержка GPT-4o** - Общайтесь с передовой AI-моделью
- 📚 **Контекстный диалог** - Полноценная история сообщений
- 🧼 **Управление историей** - `/clear` для очистки контекста
- ⚡ **Мгновенные ответы** - Оптимизированная работа с API
- 🔒 **Безопасность** - Ключ API через переменные окружения

## 🛠️ Установка и настройка

### Предварительные требования
- [Go](https://golang.org/dl/) 1.21 или новее
- Действующий API-ключ от [AIMLAPI](https://aimlapi.com)

### Быстрый старт
```bash
# Установка через go install
go install github.com/mrevds/cli-ai-chat@latest

# Установка переменной окружения (Linux/macOS)
export AIMLAPI_KEY="your_api_key_here"

# Для Windows
set AIMLAPI_KEY=your_api_key_here
