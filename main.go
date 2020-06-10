package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	var spBot spanishBot
	var enBot englishBot

	printGreetingSP(spBot)
	printGreetingEB(enBot)
}

func (englishBot) getGreeting() string { // (en englishBot) - en убрали так как en не используется в функции
	return ("Hi trhere!")
}

func (spanishBot) getGreeting() string { // (sp englishBot) - sp убрали так как sp не используется в функции
	return ("Hola!")
}

func printGreetingEB(eb englishBot) {
	fmt.Println(eb.getGreeting()) // Сразу вызывается метод getGreeting для eb
}

func printGreetingSP(sb spanishBot) {
	fmt.Println(sb.getGreeting()) // Сразу вызывается метод getGreeting для sb
}
