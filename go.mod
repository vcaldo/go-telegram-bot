module github.com/vcaldo/go-telegram-bot


require github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1 // indirect
require github.com/vcaldo/go-telegram-bot/pkg/qbitorrent v1.0.0
replace github.com/vcaldo/go-telegram-bot/pkg/qbitorrent v1.0.0 => ./pkg/qbitorrent

go 1.19

