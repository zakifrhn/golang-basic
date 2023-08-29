package module

import "fmt"

type Data struct {
	p int
	l int
	t int
}

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type Hitung interface {
	Hitung2d() float64
	Hitung3d() float64
}

func CreateHitung(p, l, t int) *Data {
	return &Data{
		p: p,
		l: l,
		t: t,
	}
}

func (h Data) luas() float64 {
	lebar := h.l
	panjang := h.p

	rumusLuas := lebar * panjang
	fmt.Println("Luasnya adalah", rumusLuas)

	return float64(rumusLuas)
}

func (h Data) keliling() float64 {
	lebar := h.l
	panjang := h.p

	rumusKeliling := 2 * (lebar + panjang)
	fmt.Println("Kelilingnya adalah", rumusKeliling)

	return float64(rumusKeliling)
}

func (h Data) volume() float64 {
	lebar := h.l
	panjang := h.p
	tinggi := h.t

	rumusVolume := panjang * lebar * tinggi
	fmt.Println("Volumenya adalah", rumusVolume)

	return float64(rumusVolume)
}

func (d *Data) Hitung2d() float64 {
	//just only pick choose one and another please command it

	return d.luas()
	//return d.keliling()

}

func (d *Data) Hitung3d() float64 {
	return d.volume()
}
