package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	//var spBot spanishBot
	//var enBot englishBot

	enBot := englishBot{}
	spBot := spanishBot{}

	printGreeting(enBot)
	printGreeting(spBot)

	//printGreetingSP(spBot)
	//printGreetingEB(enBot)
}

//Реализация printGreeting с интерфейсом:
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string { // (en englishBot) - en убрали так как en не используется в функции
	return ("Hi trhere!")
}

func (spanishBot) getGreeting() string { // (sp englishBot) - sp убрали так как sp не используется в функции
	return ("Hola!")
}

/*---------------------------------------------
Реализация без интерфецсов с 2мя разными функциями
func printGreetingEB(eb englishBot) {
	fmt.Println(eb.getGreeting()) // Сразу вызывается метод getGreeting для eb
}

func printGreetingSP(sb spanishBot) {
	fmt.Println(sb.getGreeting()) // Сразу вызывается метод getGreeting для sb
}
-----------------------------------------------*/
