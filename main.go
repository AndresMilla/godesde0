package main

import "github.com/AndresMilla/godesde0/webserver"

func main() {
	/* estado, texto := variables.ConviertoaTexto(1588)
	   fmt.Println(estado)
	   fmt.Println(texto)
	   if os := runtime.GOOS; os == "linux" || os == "OS X" {
	   	fmt.Println("Esto no es windows ", os)
	   } else {
	   	fmt.Println("Esto es:", os)
	   }

	   switch os := runtime.GOOS; os {
	   case "linux":
	   	fmt.Println("Esto es linux")
	   case "darwin":
	   	fmt.Println("Esto es darwin")
	   default:
	   	fmt.Printf("%s \n", os)
	   }

	   numero, texto := ejercicios.ConvNumerico("500")
	   fmt.Println(numero)
	   fmt.Println(texto)*/

	//teclado.IngresoNumeros()

	//fmt.Println(ejercicios.TablaMultiplicar())

	//files.GrabaTabla()

	//files.SumaTabla()

	//files.LeoArchivo()

	//funciones.LlamarClosure()

	//funciones.Exponencia(2)

	//arreglos_slices.Capacidad()

	//mapas.MostrarMapas()

	/*Pedro := new(modelos.Hombre)
	  ejer_interfaces.HumanosRespirando(Pedro)

	  Maria := new(modelos.Mujer)
	  ejer_interfaces.HumanosRespirando(Maria)*/

	//defer_panic.EjemploPanic()

	/*canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Sergio Andres Milla Cano", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aqui")*/

	webserver.MiWebServer()
}
