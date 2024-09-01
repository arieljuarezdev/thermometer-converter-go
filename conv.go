package main

import (
	"fmt"
)

func main() {

	var sca_conv string
	var sca_ini string
	var temp_init float64

	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-")
	fmt.Println("Conversor termométrico")
	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-")
	fmt.Println("Escala termométrica de entrada:")
	fmt.Println("K -> Kelvin")
	fmt.Println("°F -> Fahrenheit")
	fmt.Println("°C -> Celsius")
	fmt.Print("Digite a LETRA da opção selecionada:")
	fmt.Scanf("%s", &sca_ini)

	if sca_ini == "f" || sca_ini == "F" || sca_ini == "c" || sca_ini == "C" || sca_ini == "k" || sca_ini == "K" {

		fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-")
		fmt.Println("Escala termométrica de saída:")
		fmt.Println("K -> Kelvin")
		fmt.Println("°F -> Fahrenheit")
		fmt.Println("°C -> Celsius")
		fmt.Print("Digite a LETRA da opção selecionada:")
		fmt.Scanf("%s", &sca_conv)

		if sca_conv == "f" || sca_conv == "F" || sca_conv == "c" || sca_conv == "C" || sca_conv == "k" || sca_conv == "K" {
			switch {
			case sca_ini == "f" || sca_ini == "F":
				sca_ini := "f"
				fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-")
				fmt.Print("Digite a temperatura a ser convertida em °F: ")
				fmt.Scanf("%f", &temp_init)
				// enviar por meio do paramentro o valor de escala inicial, temperatura inicial e escala a converter.
				conv_temp(sca_ini, sca_conv, temp_init)
				break
			case sca_ini == "c" || sca_ini == "C":
				sca_ini := "c"
				fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-")
				fmt.Print("Digite a temperatura a ser convertida em °C: ")
				fmt.Scanf("%f", &temp_init)
				// enviar por meio do paramentro o valor de escala inicial, temperatura inicial e escala a converter.
				conv_temp(sca_ini, sca_conv, temp_init)
				break
			case sca_ini == "k" || sca_ini == "K":
				sca_ini := "k"
				fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-")
				fmt.Print("Digite a temperatura a ser convertida em K: ")
				fmt.Scanf("%f", &temp_init)
				// enviar por meio do paramentro o valor de escala inicial, temperatura inicial e escala a converter.
				conv_temp(sca_ini, sca_conv, temp_init)
				break
			}
		}

	} else {
		fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-")
		fmt.Println("erro: digite o LETRA CORRESPONDENTE A ESCALA SOLICITADA")
		return
	}
}

func conv_temp(s_i, s_c string, t_i float64) {
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	//tratamento da escala inicial (to chapadasso, isso não bom)
	switch {
	case s_c == "f" || s_c == "F":
		if s_i == "k" {
			// conversão Kelvin p/ Fahrenheit
			t_c := (((t_i - 273.15) * 9) / 5) + 32
			fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
			fmt.Printf("%.2f K em °F, consiste em: %.2f", t_i, t_c)
		} else {
			// conversão Celsius p/ Fahrenheit
			t_c := ((t_i * 9) / 5) + 32
			fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
			fmt.Printf("%.2f °C em °F, consiste em: %.2f", t_i, t_c)
		}
		break
	case s_c == "c" || s_c == "C":
		// K || F
		if s_i == "k" {
			// Conversão de Kelvin p/ Celsius
			t_c := t_i - 273.15
			fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
			fmt.Printf("%.2f K em °C, consiste em: %2.f", t_i, t_c)
		} else {
			// Conversão de Fahrenheit p/ Celsius
			t_c := (t_i - 32) * 5 / 9
			fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
			fmt.Printf("%.2f °F em °C, consiste em: %2.f", t_i, t_c)
		}
		break
	case s_c == "k" || s_c == "K":
		// F || C
		if s_i == "f" {
			// Conversão de Fahrenheit p/ Kelvin
			t_c := (((t_i - 32) * 5) / 9) + 273.15
			fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
			fmt.Printf("%.2f °F em K, consiste em: %.2f", t_i, t_c)
		} else {
			// Conversão de Fahrenheit p/ Celsius
			t_c := t_i + 273.15
			fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
			fmt.Printf("%.2f °C em K, consiste em: %.2f", t_i, t_c)
		}
		break
	}
}
