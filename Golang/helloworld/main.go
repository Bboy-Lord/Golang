package main

import "fmt"



func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func multi(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	if b == 0 {
		fmt.Println("Erreur: divion par zero")
		return 0
	}
}
func main()  {

	var a, b int
	var op byte

	fmt.Println("Choisir le premier nombre entier :")
	fmt.Scanln(&a)

	fmt.Println("Choisir l'opération (+, -, *, /) :")
	fmt.Scanf(" %c", &op)

	fmt.Println("Choisir le deuxième nombre entier :")
	fmt.Scanln(&b)

	if op == '+' {
		fmt.Println("Résultat =", add(a, b))
	} else if op == '-' {
		fmt.Println("Résultat =", sub(a, b))
	} else if op == '*' {
		fmt.Println("Résultat =", multi(a, b))
	} else if op == '/' {
		fmt.Println("Résultat =", div(a, b))
	}

}
	/*func remigration(interpolList []string) {
		for i := 0; i < len(interpolList); i++ {
			fmt.Println("Arrête par interpol de: ", interpolList[i])
		}
	}*/

	// string = chaine de caracteres
	// int = nombre entier (positif ou negatif) sur 4 octets soit 32 bits
	// uint = nombre entier positif sur 4 octets soit 32 bits
	// float32 = nombre a virgule sur 4 octets soit 32 bits
	// bool = valeur booleenne (true ou false)
	// rune = caracter unicode 4 octets soit 32 bits
	// byte = caracter ascii 1 octet soit 8 bits
	// int64 = nombre entier (positif ou negatif) sur 8 octets soit 64 bits
	// int32(2147483647) + int32(1)
	// uint32(2147483647) + uint32(1)
	// tu créer deux variables de type int32, x et y et tu stock le resultat l'addition dans c et tu affiches c
	
	/*
	var name string

	fmt.Println("Votre prenom: ")
	fmt.Scanln(&name)
	fmt.Println("Bonjour ", name)
	// f(x) = x + 1
	// f(1) = 1 + 1 = 2
	*/

	// UN programme qui demande l'age d'un utilisateur qui veut rentrer en boite de nuit, et determine si il est majeur ou non si c'est le cas affiche Bienvenue sinon affiche mineur
	/*var age uint8

	fmt.Println("quel est ton âge ? ")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("tu es majeur, tu peut rentrer")
	} else {
		fmt.Println("tu es mineur, tu ne peut pas entrer")
	}*/
	/*
	var interpol []string = []string{"mohammed", "youssef", "ali"}
	remigration(interpol)*/

	/*

	Fais moi un programme qui demande à l'utilisateur dans un premier temps, un nombre entier sur le terminal, dans un deuxieme temps tu vas lui faire demander sur le terminal un operateur entre +, -, * et /, 
	qui sera un type byte, dans un troisieme temps tu vas lui faire demander un deuxieme nombre entier sur le terminal

	le but du programme est de prendre chaque valeur entrée par l'utilisateur et de faire l'operation entre les deux nombres en fonction de l'operateur entré par l'utilisateur et d'afficher le resultat sur le terminal.

	Pour chaque operations tu dois créer une fonction associé à celle ci exemple

	Si l'operation c'est 1 + 1 tu dois en effet créer une fonction exemple add(a int, b int) int qui va retourner la somme de a et b

	Si l'operation c'est 1 - 1 tu dois en effet créer une fonction exemple sub(a int, b int) int qui va retourner la soustraction de a et b

	*/

	/*var result int = add(5, 10)
	fmt.Println("Result: ", result)
	*/
	


