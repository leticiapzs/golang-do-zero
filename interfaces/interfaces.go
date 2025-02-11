package main

import "fmt"

type Cachorro struct {
	Raca  string
	Cor   string
	Patas int
}

type Gato struct {
	Cor   string
	Patas int
}

func (c Cachorro) Barulho() string {
	return "au au"
}

func (c Cachorro) NumeroDePatas() int {
	return c.Patas
}

func (g Gato) Barulho() string {
	return "miau"
}

func (g Gato) NumeroDePatas() int {
	return g.Patas
}

type Animal interface {
	Barulho() string
	NumeroDePatas() int
}

func FazBarulho(animal Animal) {
	fmt.Println(animal.Barulho())
}

type Pessoa struct {
	Nome             string
	Idade            int
	Prof             string
	TempoDeProfissao int
}

type Crianca struct {
	Nome  string
	Idade int
}

func (c Crianca) FalaBomDia() string {
	return c.Nome + "deseja bom dia pra você"
}

func (p Pessoa) FalaBomDia() string {
	return fmt.Sprintf("%s deseja bom dia pra você!", p.Nome)
}

func (p Pessoa) Profissao() string {
	return fmt.Sprintf("%s tem %d anos como %s", p.Nome, p.TempoDeProfissao, p.Prof)
}

type Adulto interface {
	FalaBomDia() string
	Profissao() string
}

func BomDia(adulto Adulto) string {
	return adulto.FalaBomDia()
}

func main() {

	adulto := Pessoa{
		Nome:             "steph",
		Idade:            28,
		Prof:             "dev",
		TempoDeProfissao: 6,
	}

	crianca := Crianca{
		Nome:  "bento",
		Idade: 4,
	}

	crianca.FalaBomDia()
	adulto.FalaBomDia()
	BomDia(adulto)

	cachorro := Cachorro{
		Raca:  "spitz alemão",
		Cor:   "preto",
		Patas: 4,
	}

	gato := Gato{
		Cor:   "branco",
		Patas: 4,
	}

	FazBarulho(gato)
	FazBarulho(cachorro)
}
