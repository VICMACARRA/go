package main

import "fmt"

type Extras struct {
	airbag        bool
	airco         bool
	alarma        bool
	cierreCentral bool
	frenosABS     bool
	controlTC     bool
}

type Coche struct {
	marca        string
	modelo       string
	color        string
	velocidadMax int
	potencia     int
	cilindrada   int
	peso         int
	precio       int
	Extras       Extras
}

type PersonaService interface {
	getMarca() string
	setMarca(m string)
	getModelo() string
	setModelo(m string)
	getColor() string
	setColor(c string)
	getVelocidadMax() int
	setVelocidadMax(v int)
	getPotencia() int
	setPotencia(p int)
	getCilindrada() int
	setCilindrada(c int)
	getPeso() int
	setPeso(p int)
	getPrecio() int
	setPrecio(p int)
	getAirbag() bool
	setAirbag(a bool)
	getAirco() bool
	setAirco(a bool)
	getAlarma() bool
	setAlarma(a bool)
	getCierreCentral() bool
	setCierreCentral(c bool)
	getFrenosABS() bool
	setFrenosABS(f bool)
	getControlTC() bool
	setControlTC(c bool)
}

func (c Coche) getMarca() string {
	return c.marca
}
func (c *Coche) setMarca(m string) {
	c.marca = m
}
func (c Coche) getModelo() string {
	return c.modelo
}
func (c *Coche) setModelo(m string) {
	c.modelo = m
}
func (c Coche) getColor() string {
	return c.color
}
func (c2 *Coche) setColor(c string) {
	c2.color = c
}
func (c Coche) getVelocidadMax() int {
	return c.velocidadMax
}
func (c *Coche) setVelocidadMax(v int) {
	c.velocidadMax = v
}
func (c Coche) getPotencia() int {
	return c.potencia
}
func (c *Coche) setPotencia(p int) {
	c.potencia = p
}
func (c Coche) getCilindrada() int {
	return c.cilindrada
}
func (c3 *Coche) setCilindrada(c int) {
	c3.cilindrada = c
}
func (c Coche) getPeso() int {
	return c.peso
}
func (c *Coche) setPeso(p int) {
	c.peso = p
}
func (c Coche) getPrecio() int {
	return c.precio
}
func (c *Coche) setPrecio(p int) {
	c.precio = p
}
func (c Coche) getAirbag() bool {
	return c.Extras.airbag
}
func (c *Coche) setAirbag(a bool) {
	c.Extras.airbag = a
}
func (c Coche) getAirco() bool {
	return c.Extras.airco
}
func (c *Coche) setAirco(a bool) {
	c.Extras.airco = a
}
func (c Coche) getAlarma() bool {
	return c.Extras.alarma
}
func (c *Coche) setAlarma(a bool) {
	c.Extras.alarma = a
}
func (c Coche) getCierreCentral() bool {
	return c.Extras.cierreCentral
}
func (c4 *Coche) setCierreCentral(c bool) {
	c4.Extras.cierreCentral = c
}
func (c Coche) getFrenosABS() bool {
	return c.Extras.frenosABS
}
func (c *Coche) setFrenosABS(f bool) {
	c.Extras.frenosABS = f
}
func (c Coche) getControlTC() bool {
	return c.Extras.controlTC
}
func (c5 *Coche) setControlTC(c bool) {
	c5.Extras.controlTC = c
}

func NewCarService() *Coche {
	return &Coche{}
}

func NewCarService2(marca string, modelo string,color string,velocidadMax int,potencia int,cilindrada int,peso int,precio int) *Coche {
	return &Coche{marca, modelo, color, velocidadMax, potencia, cilindrada, peso, precio, Extras{false, false, false, false, false, false}}
}

func main() {

	mazda := NewCarService()

	mazda.setMarca("Mazda")
	mazda.setModelo("MX-5")
	mazda.setColor("Negro")
	mazda.setAirbag(true)
	fmt.Println(mazda)
	fmt.Println(mazda.getMarca())

		opel := NewCarService2("Opel","Corsa","rosa",100,222,1600,1200,20000)
		fmt.Println(opel)
		opel.setAirbag(true)
		fmt.Println(opel)


}