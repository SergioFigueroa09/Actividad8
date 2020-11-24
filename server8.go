package main

import (
	//"errors"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)

type Server struct{
	Materias map[string]map[string]float64
	Alumnos map[string]map[string]float64
	Lista_alumnos []string
	Lista_materias []string
}
var S Server

//AgregarMateria es la función que le comenté.
func (this *Server) AgregarMateria(nombre_materia string, reply *string) error {
	fmt.Println("-------------------AgregarMateria")
	 // hacer global esta variable porque cada vez que entre se va a sobre escribir
	 //Materias := make(map[string]map[string]float64)
	 //S.Materias = Materias
	//S.Materias = make(map[string]map[string]float64) //Si comento esta linea, no pasa el problema, pero obviamente no se genera el Map.
	alumno := make(map[string]float64)
	//alumno[nombre_materia] = 90 // falta obtener nombre de alumnos y la calificacion para esa materia
	S.Materias[nombre_materia] = alumno
	S.Lista_materias = append(S.Lista_materias, nombre_materia)
	*reply = "Se agregó " + nombre_materia + " como nueva materia"
	fmt.Println("Nueva Lista de materias: ",S.Lista_materias)
	return nil
}

func (this *Server) AgregarAlumno(nombre_alumno string, reply *string) error {
	fmt.Println("-------------------AgregarAlumno")
   	materia := make(map[string]float64)
   	S.Alumnos[nombre_alumno] = materia
   	S.Lista_alumnos = append(S.Lista_alumnos, nombre_alumno)
   	*reply = "Se agregó a " + nombre_alumno + " como nuev@ alumn@"
   	fmt.Println("Nueva Lista de alumnos: ",S.Lista_alumnos)
   	return nil
}

func (this *Server) AgregarCalifMateria(args []string, reply *string) error{
	fmt.Println("-------------------AgregarCalifMateria")
	var calif float64
	fmt.Println("nombre_materia: ",args[0])
	fmt.Println("nombre_alumno: ",args[1])
	fmt.Println("calif: ",args[2])
	if s, err := strconv.ParseFloat(args[2], 64); err == nil {
		//fmt.Println(s)
		calif = s
	}

	fmt.Println("S.Materias["+args[0]+"]["+args[1]+"] = "+args[2])
	S.Materias[args[0]][args[1]] = calif
	fmt.Println("PAST MATERIAS")

	/* materia := make(map[string]float64)
	materia[args[1]] = calif
	S.Alumnos[args[0]] = materia */

	fmt.Println("S.Alumnos["+args[1]+"]["+args[0]+"] = "+args[2])
	S.Alumnos[args[1]][args[0]] = calif
	fmt.Println("PAST ALUMNOS")
	

	*reply = "Se calificó a " + args[1] + " con " + fmt.Sprintf("%f", calif) + " en la materia " + args[0]
	return nil
}

func (this *Server) MostrarMaterias(nombre_materia string, reply *[]string) error{
	fmt.Println("-------------------MostrarMaterias")
	fmt.Println("map materias: ",S.Materias)
	//fmt.Println("", S.Lista_materias)
	*reply = S.Lista_materias
	return nil
}

func (this *Server) MostrarAlumnos(nombre_alumno string, reply *[]string) error{
	fmt.Println("-------------------MostrarAlumnos")
	fmt.Println("map alumnos: ",S.Alumnos)
	//fmt.Println("", S.Lista_materias)
	*reply = S.Lista_alumnos
	return nil
}

func (this *Server) PromedioAlumno(nombre_alumno string, reply *string) error{
	fmt.Println("-------------------PromedioAlumno")
	var i float64
	i = 0
	var promedio float64
	promedio = 0
	fmt.Println(nombre_alumno)
	for materia, calificion := range S.Alumnos[nombre_alumno] {
		fmt.Println(materia + ":", calificion)
		i++
		promedio = promedio + calificion
		} 
	promedio = promedio/i
	str_prom := fmt.Sprintf("%f", promedio)
	*reply = "El promedio del alumno es: " + str_prom
	return nil
}

func (this *Server)PromedioGeneral(nombre_alumno string, reply *string)error{
	fmt.Println("-------------------PromedioGeneral")
	var i float64
	i = 0
	var promedio_Alumno float64
	var num_clases float64
	var promedio_General float64
	for k:=0;k < len(S.Lista_alumnos);k++{
		promedio_Alumno = 0
		num_clases = 0
		for _, calificion := range S.Alumnos[S.Lista_alumnos[int(i)]]{
			promedio_Alumno = promedio_Alumno + calificion
			num_clases++
		}
		i++
		promedio_Alumno = promedio_Alumno/num_clases
		promedio_General = promedio_General + promedio_Alumno
	}
	promedio_General = promedio_General/float64(len(S.Lista_alumnos))
	fmt.Println("El promedio general es de: ", promedio_General)
	*reply = "El promedio general es de: " + fmt.Sprintf("%f", promedio_General)
	return nil
}

func (this *Server) PromedioMateria (nombre_materia string, reply *string)error{
	fmt.Println("-------------------PromedioMateria")
	var i float64
	i = 0
	var promedio float64
	promedio = 0
	for alumno, calificion := range S.Materias[nombre_materia] {
		fmt.Println(alumno + ":", calificion)
		i++
		promedio = promedio + calificion
	} 
	promedio = promedio/i
	fmt.Println("El promedio de la clase es: ", promedio)
	*reply = "El promedio de la Materia de " +nombre_materia+ " es de: " + fmt.Sprintf("%f", promedio)
	return nil
}

func (this *Server) Hello(name string, reply *string) error {
	*reply = "Hello " + name
	return nil
}

func server(){
	
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp",":9999")
	if err != nil{
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main(){
	S.Materias = make(map[string]map[string]float64)
	S.Alumnos = make(map[string]map[string]float64)
	go server()

	var input string
	fmt.Scanln(&input)
}