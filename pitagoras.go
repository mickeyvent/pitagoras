package main

import(
	"fmt"
	"time"
	"os"
	"math"
	"os/exec"
	"github.com/briandowns/spinner"
)
var(
	dt,ps string
	hipotenusa,alado,frontal float64
)
func main(){
	for a:=0; a>=0; a++{
		clear()
		table()
		fmt.Printf("   [exit]  [help]  [run] \n")
		fmt.Printf("Pitagoras:")
		fmt.Scanf("%v",&dt)
		if dt == "a" || dt == "A"{
			fmt.Printf("valor de A:")
			fmt.Scanf("%v",&alado)
		}else if dt == "b" || dt == "B"{
			fmt.Printf("valor de B:")
			fmt.Scanf("%v",&frontal)
		}else if dt == "c" || dt == "C"{
			fmt.Printf("valor de C:")
			fmt.Scanf("%v",&hipotenusa)
		}else if dt == "salir" || dt == "exit"{
			os.Exit(0)
		}else if dt == "help"{
			help()
		}else if dt == "run"{
			ejecucion()
		}else{
			fmt.Printf("Erro no se reconoce:%v",dt)
			time.Sleep(1*time.Second)
		}
	}
}
func ejecucion(){
	for b:=0; b>=0; b++{
		if hipotenusa == 0{
			alado = alado * alado
			frontal = frontal * frontal
			hipotenusa = frontal + alado
			hipotenusa = math.Sqrt(hipotenusa)
			alado = math.Sqrt(alado)
			frontal = math.Sqrt(frontal)
			break
		}else if alado == 0{
			frontal = frontal * frontal
			hipotenusa = hipotenusa * hipotenusa
			alado = frontal - hipotenusa
			alado = math.Sqrt(alado)
			frontal = math.Sqrt(frontal)
			hipotenusa = math.Sqrt(hipotenusa)
			break
		}else if frontal == 0{
			hipotenusa = hipotenusa * hipotenusa
			alado = alado * alado
			frontal = alado - hipotenusa
			frontal = math.Sqrt(frontal)
			hipotenusa = math.Sqrt(hipotenusa)
			alado = math.Sqrt(alado)
			break
		}
		if b == 20{
			fmt.Printf("Llenaste todos los lados...")
			time.Sleep(2*time.Second)
			break
		}
	}
}
func animate(){
	s := spinner.New(spinner.CharSets[14],100 * time.Millisecond)
	s.Suffix = " Help !!!  Help !!!"
	s.Start()
	time.Sleep(2*time.Second)
	s.Stop()
}
func help(){
	clear()
	animate()
	fmt.Printf("\nsi quieres modificar un lado del triangulo\n")
	fmt.Printf("solo escribe la letra que quieres modificar\n")
	fmt.Printf("y luego ingresa el valor. ejemplo: \n")
	fmt.Printf("   $Pitagoras: c\n")
	fmt.Printf("      ingresa el valor C: 8 \n")
	fmt.Printf("para obtener el resultado solo escribe\n")
	fmt.Printf("   $Pitagoras: run \n")
	pause()
}
func table(){
	fmt.Printf("PITAGORAS...\n")
	fmt.Printf("      C(%v)\n",hipotenusa)
	fmt.Printf("                /|\n")
	fmt.Printf("              /  |\n")
	fmt.Printf("            /    |\n")
	fmt.Printf("          /      | A(%v)\n",alado)
	fmt.Printf("        /        |\n")
	fmt.Printf("      /          |\n")
	fmt.Printf("    /____________|\n")
	fmt.Printf("        B(%v)\n",frontal)
}
func clear(){
	cl := exec.Command("clear")
	cl.Stdout = os.Stdout
	cl.Run()
}
func pause(){
	fmt.Printf("\nENTER PARA CONTINUAR:")
	fmt.Scanf("%v",&ps)
}
