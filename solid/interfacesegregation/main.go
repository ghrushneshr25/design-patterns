package main

// breaking big interfaces to smaller interfaces

type Document struct{}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m *MultiFunctionPrinter) Print(d Document) {}

func (m *MultiFunctionPrinter) Fax(d Document) {}

func (m *MultiFunctionPrinter) Scan(d Document) {}

type OldFunctionPrinter struct{}

func (m *OldFunctionPrinter) Print(d Document) {}

// Deprecated
func (m *OldFunctionPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Deprecated
func (m *OldFunctionPrinter) Scan(d Document) {
	panic("operation not supported")
}

// We have created a problem by implementing many functions

func main() {

}

/// Interface Segregation Principle

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Fax interface {
	Fax(d Document)
}

type MyPrinter struct {
}

func (printer *MyPrinter) Print(d Document) {

}

type PhotoCopier struct {
}

func (p *PhotoCopier) Scan(d Document) {

}

func (p *PhotoCopier) Print(d Document) {

}

type MultiFunctionDevice interface { /// ability to combine interfaces
	Printer
	Scanner
}

type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (p *MultiFunctionMachine) Scan(d Document) {
	p.scanner.Scan(d)
}

func (p *MultiFunctionMachine) Print(d Document) {
	p.printer.Print(d)
}
