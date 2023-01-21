# lazytg

> simple telegram bot api

### Example

#### Run reading messages
```go
bot, err := tgapi.NewBot(monitoringToken)
if err != nil {
    log.Fatal("New TG bot error: ", err)
}

bot.ProductionMode()

bot.RunAnswer(func(ctx message.IContext) string {
    fmt.Println("chat id:", ctx.Message().ChatID)

    return ctx.Message().Text + ". this is answer!!!"
}, 60)
```

#### Send message
```go
if err = bot.SendChat("*bold text*", chatID); err != nil {
    log.Println("send message error: ", err)
}

if err = bot.SendChat("__underline text__", chatID); err != nil {
    log.Println("send message error: ", err)
}

if err = bot.SendChat("_italic text_", chatID); err != nil {
    log.Println("send message error: ", err)
}

if err = bot.SendChat("*multi line header*\n\n_some description:_ __hello__", chatID); err != nil {
    log.Println("send message error: ", err)
}
```