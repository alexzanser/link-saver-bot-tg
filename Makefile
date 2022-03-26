build		:
			go build -o bot && ./bot -tg-bot-token ${token}

clean		:
			rm bot