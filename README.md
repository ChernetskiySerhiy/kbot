# kbot# kbot

Telegram-бот на Go — практичне завдання курсу «DevOps та Kubernetes. Практичний інтенсив+» (Prometheus).

Бот у Telegram: [t.me/chernetskiy_kbot_bot](https://t.me/chernetskiy_kbot_bot)

## Що вміє

| Повідомлення | Відповідь бота |
|---|---|
| `/start` | `Hi! I'm Kbot Version. Send me hello or any text.` |
| `hello` | `Hello I'm Kbot Version!` |
| будь-який інший текст | `You wrote: <ваш текст>` |

## Технології

- Go
- [github.com/spf13/cobra](https://github.com/spf13/cobra) — команди в терміналі (`kbot`, `version`)
- [gopkg.in/telebot.v4](https://gopkg.in/telebot.v4) — робота з Telegram Bot API

## Встановлення

1. Створіть бота в Telegram через [@BotFather](https://t.me/BotFather) і отримайте токен.
2. Склонуйте репозиторій і зберіть програму:

```bash
git clone https://github.com/ChernetskiySerhiy/kbot.git
cd kbot
go build -o kbot
```

## Запуск

Токен передається через змінну середовища `TELE_TOKEN` (у коді й у репозиторії його немає):

```bash
read -s TELE_TOKEN
export TELE_TOKEN
./kbot kbot
```

## Команди

```bash
./kbot version   # показати версію
./kbot kbot      # запустити бота (те саме: ./kbot start)
./kbot --help    # довідка
```