package main

import (
	"fmt"
	"net/rpc"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func client() {
	nombre_materia := "nombre_materia"
	nombre_alumno := "nombre_alumno"
	var lista_alumnos_client []string
	var lista_materias_client []string
	//var calif float64
	c, err := rpc.Dial("tcp",":9999")
	if err != nil{
		fmt.Println(err)
		return
	}
	var op int64
	for{
		fmt.Println("------------------------MENU------------------------")
		fmt.Println("1.-Agregar nueva materia")
		fmt.Println("2.-Agregar Alumno")
		fmt.Println("3.-Agregar calificación a alumno en materia")
		fmt.Println("4.-Ver todas las materias")
		fmt.Println("5.-Ver todos los alumnos")
		fmt.Println("6.-Mostrar promedio de un alumno")
		fmt.Println("7.-Mostrar promedio General")
		fmt.Println("8.-Mostrar promedio de una materia")		

		fmt.Println("0.-salir")
		fmt.Scanln(&op)

		switch op{
		case 1: //Agregar nueva materia
			fmt.Println("AGREGAR NUEVA MATERIA----------")
			fmt.Println("Inserte el nombre de la nueva materia (MAYUSCULAS): ")
			fmt.Scanln(&nombre_materia)

			var result string
			err = c.Call("Server.AgregarMateria",nombre_materia,&result)
			if err != nil{
				fmt.Println(err)
			} else {
				lista_materias_client = append(lista_materias_client,nombre_materia)
				fmt.Println(result)
			}
		case 2: //Agregar nuevo Alumno

			fmt.Println("AGREGAR NUEVO ALUMNO----------")
			fmt.Println("Inserte el nombre del nuevo alumno (MAYUSCULAS): ")
			fmt.Scanln(&nombre_alumno)

			var result string
			err = c.Call("Server.AgregarAlumno",nombre_alumno,&result)
			if err != nil{
				fmt.Println(err)
			} else {
				lista_alumnos_client = append(lista_alumnos_client,nombre_alumno)
				fmt.Println(result)
			}
		case 3: //AGREGAR CALIF A ALUMNO
			var str_calif string
			fmt.Println("AGREGAR CALIF ALUMNO A MATERIA----------")
			fmt.Println("Inserte el nombre (mayusculas) de la materia EXISTENTE: ")
			fmt.Scanln(&nombre_materia)
			if (!contains(lista_materias_client,nombre_materia)){
				fmt.Println("REVISAR MATERIA E INTENTAR DE NUEVO... INGRESADO:"+nombre_materia)
				break;
			} else {
				fmt.Println("Inserte el nombre (mayusculas) de un alumno EXISTENTE: ")
				fmt.Scanln(&nombre_alumno)
				if (!contains(lista_alumnos_client,nombre_alumno)){
					fmt.Println("REVISAR ALUMNO E INTENTAR DE NUEVO... INGRESADO:"+nombre_alumno)
					break;
				} else {
					fmt.Println("Inserte la calificación de un alumno: ")
					fmt.Scanln(&str_calif)

					args := []string{"nombre_materia","nombre_alumno","calif"}
					args[0] = nombre_materia
					args[1] = nombre_alumno
					args[2] = str_calif

					var result string
					err = c.Call("Server.AgregarCalifMateria",args, &result)
					if err != nil{
						fmt.Println(err)
					} else {
						fmt.Println(result)
					}
				}
			}
			
		case 4: //MOSTRAR MATERIAS
			var name string
			var list_result []string
			err = c.Call("Server.MostrarMaterias",name, &list_result)
			if err != nil{
				fmt.Println(err)
			} else {
				fmt.Println("Lista de materias=\n", list_result)
			}
		case 5: //MOSTRAR ALUMNOS
			var name string
			var list_result []string
			err = c.Call("Server.MostrarAlumnos",name, &list_result)
			if err != nil{
				fmt.Println(err)
			} else {
				fmt.Println("Lista de Alumnos=\n", list_result)
			}
		case 6: //PROMEDIO ALUMNO
			var result string
			fmt.Println("Inserte nombre (mayusculas) del alumno EXISTENTE: ")
			fmt.Scanln(&nombre_alumno)
			if (!contains(lista_alumnos_client,nombre_alumno)){
				fmt.Println("REVISAR ALUMNO E INTENTAR DE NUEVO... INGRESADO:"+nombre_alumno)
				break;
			} else {
				err = c.Call("Server.PromedioAlumno",nombre_alumno, &result)
				if err != nil{
					fmt.Println(err)
				} else {
					fmt.Println(result)
				}
			}
		case 7://PROMEDIO GENERAL
			var result string
			err = c.Call("Server.PromedioGeneral",nombre_alumno, &result)
				if err != nil{
					fmt.Println(err)
				} else {
					fmt.Println(result)
				}
		case 8:
			var result string
			fmt.Println("Inserte nombre (mayusculas) de materia EXISTENTE: ")
			fmt.Scanln(&nombre_materia)
			if (!contains(lista_materias_client,nombre_materia)){
				fmt.Println("REVISAR MATERIA E INTENTAR DE NUEVO... INGRESADO:"+nombre_materia)
				break;
			} else {
				err = c.Call("Server.PromedioMateria",nombre_materia, &result)
				if err != nil{
					fmt.Println(err)
				} else {
					fmt.Println(result)
				}
			}
		case 0:
			return
		}
	}
}

func main(){
	client()
}
