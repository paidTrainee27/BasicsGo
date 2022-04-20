package main

import "fmt"

// declaring blueprint/interface/requirement(client)
type Printer interface {
	print() string
}

type Scanner interface {
	scan() string
}

type ScanningPrinter interface {
	Printer
	Scanner
}

//adapting structs to the blueprint/interface
func MyProprinter() ScanningPrinter {
	return &proprinter{}
}

//creating empty model that will implement the requirement
type printer struct{}
type proprinter struct{}

//implementing the blueprint requirement
func (p *printer) print() string {
	return "printing ..."
}
func (p *proprinter) print() string {
	return "pro printing ..."
}

func (p *printer) scan() string {
	return "scanning ..."
}
func (p *proprinter) scan() string {
	return "pro scanning ..."
}

//testing the process...
func process(p ScanningPrinter) {
	fmt.Println(p.print())
	fmt.Println(p.scan())
}

func main() {
//Incase you need to pass some dependencies
// process(MyProprinter())
// OR
// printer := &printer{}
// process(printer)
	RunMultipleInterfaces()
}
