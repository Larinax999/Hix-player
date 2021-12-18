package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

var Exec sync.WaitGroup

func Randarr(arr []string) string {
	a := rand.NewSource(time.Now().UnixNano())
	b := rand.New(a)
	c := arr[b.Intn(len(arr))]
	return c
}

func conn(token string, filename string, gid string, chid string) {
	// Connect to Discord
	discord, _ := discordgo.New(token)
	discord.Open()
	dgv, _ := discord.ChannelVoiceJoin(gid, chid, false, true)
	for {
		dgvoice.PlayAudioFile(dgv, filename, make(chan bool))
	}
	// Close connections
	//dgv.Close()
	//discord.Close()
	//return true
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("[*] discord guild player by larinax999\n[!] Usage: ./Hix <token_file_name> <music_file_name> <guild id> <channel id list> \n[!] Usage: ./Hix token.txt play.mp3 000 000z000")
		os.Exit(1)
	}
	chdilist := strings.Split(os.Args[4], "z")
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("[!] file not found")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Exec.Add(1)
		token := scanner.Text()
		ch := Randarr(chdilist)
		go conn(token, os.Args[2], os.Args[3], ch)
	}
	Exec.Wait()
}
