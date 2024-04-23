package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	complatado bool
}

type ListaTareas struct {
	tareas []Tarea
}

// Metodo para agregar tarea
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// Metodo para marcar completado
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].complatado = true
}

// Metodo para editar tarea
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// Metodo para eliminar tarea
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}

func main() {
	// Crear una instancia
	lista := ListaTareas{}
	leer := bufio.NewReader(os.Stdin)
	for {
		var opcion int
		fmt.Println("seleccione una opcion \n",
			"1. Agregar Tarea\n",
			"2.Marcar tarea como completada\n",
			"3. Editar Tarea\n",
			"4. Eliminar Tarea\n",
			"5. Salir")

		fmt.Print("Ingrese opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n') // string y error
			fmt.Println("Ingrese el nombre de la Descripcion: ")
			t.desc, _ = leer.ReadString('\n') // string y error

			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea marcar como completada ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente")
		case 3:
			var index int
			var t Tarea
			fmt.Println("Ingrese el indice de la tarea: ")
			fmt.Scanln(&index)
			fmt.Println("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n') // string y error
			fmt.Println("Ingrese el descripcion de la tarea: ")
			t.desc, _ = leer.ReadString('\n') // string y error
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada correctamente")

		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea Eliminar ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente")
		case 5:
			fmt.Println("Saliendo del programa")
			return
		default:
			fmt.Println("Opcion invalida")
		}

		// Listar todas las tareas
		fmt.Println("Lista de tareas:")
		fmt.Println("============================================")

		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.complatado)
		}
		fmt.Println("============================================")
	}
}
