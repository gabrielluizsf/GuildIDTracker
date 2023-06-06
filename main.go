package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		return
	}

	// Obter o token do Discord do arquivo .env
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		fmt.Println("Token do Discord não encontrado no arquivo .env")
		return
	}

	// Criar uma nova sessão do DiscordGo
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Erro ao criar a sessão do DiscordGo:", err)
		return
	}

	// Adicionar um handler para o evento de pronto (ready)
	dg.AddHandlerOnce(func(s *discordgo.Session, r *discordgo.Ready) {
		// Iterar sobre as guildas (grupos) em que o bot está presente
		for _, guild := range r.Guilds {
			fmt.Println("Guild ID:", guild.ID)
		}

		// Fechar a sessão do DiscordGo
		dg.Close()
	})

	// Abrir a sessão do DiscordGo
	err = dg.Open()
	if err != nil {
		fmt.Println("Erro ao abrir a sessão do DiscordGo:", err)
		return
	}

	// Esperar indefinidamente até que o bot seja desconectado
	select {}
}
