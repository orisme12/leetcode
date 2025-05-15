package main 
import "fmt"
		
// Comentarios de una sola linea 

/*
Comentarios de varias lineas de codigos (Bloques de codigos)

*/

// Declaracion de variables de manera Global
var(
		name = "Carlos Bonet"
		edad int = 20
		direccion = "Calle 9 6a 11"
		estadoCivil = false

	)
const (
		ANGULO = 45
		PI = 3.1416
	)


func main(){
	
	// Declaracion de variable local y su valor es inferido , significa que el compilador decide el tipo de la variable, en función del valor.
	year := 1990

	// otra forma de inferido es : 
	var comunicado = "Aprobado"

	fmt.Println("Mi nombre es: ",name,"y fui afortunadamente",comunicado, "en la universidad para ingeniero de Software") // Impresión por consola 
	
	
	
	// Uso de condicional if 
	if year > 1996 {
		fmt.Println("Eres una persona relativamente Joven :3 ")
	}else if year > 1992 && year <= 1996 { 
		fmt.Println(" Ya los años te estan cyaendo encima bro :/")
	} else {
		fmt.Println("F mi rey , ya estás Viejo")
	}

	mensaje := Hello("Carlos")
	fmt.Println(mensaje)

	
	num(12)

	// Creación de Arrays 

	var array1 = [4]int{1,2,3} // Se determina el tamaño y tipo de array 
	array2 := [...]int{5,6,7,7,8,9} // Infiere en el tamaño del array usando [..]
	array3 := [5]int{1:10,2:40} // Inicializar elementos específicos en posiciones especificas
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println("el tamaño del array 1 es: ",len(array1))
	fmt.Println("el tamaño del array 2 es: ",len(array2))
	fmt.Println("el tamaño del array 2 es: ",len(array3))


	// Creación de slices 
	var slice1 = []int{1,3,4} // Declarando el typo de elementos dentro del Slice
	frutas := []string{"Bananas", "Manzanas", "Uvas", "Fresas"}
	var carrosMarcas = []string{"Mazda", "Hyundai", "Toyota"}
	slice1 = append(slice1, 20, 21) // Append Permite agregar elementos al Slice seleccionado
	slice2 := []int{5,6,7,8} 
	slice3 := append(slice1, slice2...) // De esta forma se unen dos Slices
	fmt.Println(slice3)
	fmt.Println(frutas)
	fmt.Println(carrosMarcas)

	
		// Crear un copia usando Copy 
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15} 
	takeNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(takeNumbers))
	copy(numbersCopy, takeNumbers)
   
	fmt.Printf("numbersCopy = %v\n", numbersCopy)
}

	// Función pública , el nombre de la función inicia con Mayuscula y puede usarse en otros packages dentro del proyecto
	func Hello (name string) string {
		return hi(name)
	}

	// Función privada , el nombre de la funcion empieza en minuscula y solo puede usarse dentro del package donde se creó 
	func hi (name string) string {
		return "hi " + name
	}


// Funcion para imprimir si un numero es PAR O IMPAR  Usando For y condicional if

func num (num int) { // Se crea la función y se define el paramatro y el tipo del parametro " Entero"
	fmt.Println("--- Enteros del numero ..... !" , num)

	for number:= 0; number <= num ; number++ { // Se inicializa el valor de number, se hace la comparación con el valor de num y posteriormente se aumenta el valor de number hasta que se cumpla la condición
		if number%2 ==0{ // Si el residuo de number es 0 se establece que es un par 
			fmt.Println("PAR \n ", number)
		}else if number %2 == 1{ // Si el residuo de number es 1 se establece que es un impar 
			fmt.Println("IMPAR \n ", number)
		}else{ // Caso no se realice las dos comparaciones anteriores se muestre el mensaje del else
			fmt.Println("NO ES PAR NI IMPAR %\n ", number)
		}
	}
}