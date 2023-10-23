package main

import (
    "fmt"
    "net/http"
    "strconv"
    "sync"

    "github.com/julienschmidt/httprouter"
)

type BangunDatar interface {
    HitungLuas() float64
    HitungKeliling() float64
}

type BangunRuang interface {
    HitungVolume() float64
    HitungLuasPermukaan() float64
}

type Persegi struct {
    Sisi float64
}

func (p Persegi) HitungLuas() float64 {
    return p.Sisi * p.Sisi
}

func (p Persegi) HitungKeliling() float64 {
    return 4 * p.Sisi
}

type PersegiPanjang struct {
    Panjang float64
    Lebar   float64
}

func (pp PersegiPanjang) HitungLuas() float64 {
    return pp.Panjang * pp.Lebar
}

func (pp PersegiPanjang) HitungKeliling() float64 {
    return 2 * (pp.Panjang + pp.Lebar)
}

type Lingkaran struct {
    JariJari float64
}

func (l Lingkaran) HitungLuas() float64 {
    return 3.14 * l.JariJari * l.JariJari
}

func (l Lingkaran) HitungKeliling() float64 {
    return 2 * 3.14 * l.JariJari
}

type Kubus struct {
    Sisi float64
}

func (k Kubus) HitungVolume() float64 {
    return k.Sisi * k.Sisi * k.Sisi
}

func (k Kubus) HitungLuasPermukaan() float64 {
    return 6 * k.Sisi * k.Sisi
}

type Balok struct {
    Panjang float64
    Lebar   float64
    Tinggi  float64
}

func (b Balok) HitungVolume() float64 {
    return b.Panjang * b.Lebar * b.Tinggi
}

func (b Balok) HitungLuasPermukaan() float64 {
    return 2*(b.Panjang*b.Lebar + b.Panjang*b.Tinggi + b.Lebar*b.Tinggi)
}

type Tabung struct {
    JariJari float64
    Tinggi   float64
}

func (t Tabung) HitungVolume() float64 {
    return 3.14 * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) HitungLuasPermukaan() float64 {
    return 2*3.14*t.JariJari*t.Tinggi + 2*3.14*t.JariJari*t.JariJari
}

var mu sync.Mutex

func main() {
    router := httprouter.New()

    router.GET("/bangun-datar/:bangun", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        bangun := p.ByName("bangun")

        switch bangun {
        case "persegi":
            sisiStr := r.URL.Query().Get("sisi")
            sisi, err := strconv.ParseFloat(sisiStr, 64)
            if err != nil {
                http.Error(w, "Invalid parameter 'sisi'", http.StatusBadRequest)
                return
            }

            luas := Persegi{Sisi: sisi}.HitungLuas()
            fmt.Fprintf(w, "Luas: %f", luas)
        case "persegi-panjang":
            panjangStr := r.URL.Query().Get("panjang")
            lebarStr := r.URL.Query().Get("lebar")
            hitung := r.URL.Query().Get("hitung")

            panjang, err1 := strconv.ParseFloat(panjangStr, 64)
            lebar, err2 := strconv.ParseFloat(lebarStr, 64)
            if err1 != nil || err2 != nil {
                http.Error(w, "Invalid parameters 'panjang' or 'lebar'", http.StatusBadRequest)
                return
            }

            switch hitung {
            case "luas":
                luas := PersegiPanjang{Panjang: panjang, Lebar: lebar}.HitungLuas()
                fmt.Fprintf(w, "Luas: %f", luas)
            case "keliling":
                keliling := PersegiPanjang{Panjang: panjang, Lebar: lebar}.HitungKeliling()
                fmt.Fprintf(w, "Keliling: %f", keliling)
            default:
                http.Error(w, "Invalid parameter 'hitung'", http.StatusBadRequest)
            }
        case "lingkaran":
            jariJariStr := r.URL.Query().Get("jariJari")
            hitung := r.URL.Query().Get("hitung")

            jariJari, err := strconv.ParseFloat(jariJariStr, 64)
            if err != nil {
                http.Error(w, "Invalid parameter 'jariJari'", http.StatusBadRequest)
                return
            }

            switch hitung {
            case "luas":
                luas := Lingkaran{JariJari: jariJari}.HitungLuas()
                fmt.Fprintf(w, "Luas: %f", luas)
            case "keliling":
                keliling := Lingkaran{JariJari: jariJari}.HitungKeliling()
                fmt.Fprintf(w, "Keliling: %f", keliling)
            default:
                http.Error(w, "Invalid parameter 'hitung'", http.StatusBadRequest)
            }
        default:
            http.Error(w, "Invalid bangun datar", http.StatusBadRequest)
        }
    })

    router.GET("/bangun-ruang/:bangun", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        bangun := p.ByName("bangun")

        switch bangun {
        case "kubus":
            sisiStr := r.URL.Query().Get("sisi")
            sisi, err := strconv.ParseFloat(sisiStr, 64)
            if err != nil {
                http.Error(w, "Invalid parameter 'sisi'", http.StatusBadRequest)
                return
            }

            volume := Kubus{Sisi: sisi}.HitungVolume()
            fmt.Fprintf(w, "Volume: %f", volume)
        case "balok":
            panjangStr := r.URL.Query().Get("panjang")
            lebarStr := r.URL.Query().Get("lebar")
            tinggiStr := r.URL.Query().Get("tinggi")
            hitung := r.URL.Query().Get("hitung")

            panjang, err1 := strconv.ParseFloat(panjangStr, 64)
            lebar, err2 := strconv.ParseFloat(lebarStr, 64)
            tinggi, err3 := strconv.ParseFloat(tinggiStr, 64)
            if err1 != nil || err2 != nil || err3 != nil {
                http.Error(w, "Invalid parameters 'panjang', 'lebar', or 'tinggi'", http.StatusBadRequest)
                return
            }

            switch hitung {
            case "volume":
                volume := Balok{Panjang: panjang, Lebar: lebar, Tinggi: tinggi}.HitungVolume()
                fmt.Fprintf(w, "Volume: %f", volume)
            case "luasPermukaan":
                luasPermukaan := Balok{Panjang: panjang, Lebar: lebar, Tinggi: tinggi}.HitungLuasPermukaan()
                fmt.Fprintf(w, "Luas Permukaan: %f", luasPermukaan)
            default:
                http.Error(w, "Invalid parameter 'hitung'", http.StatusBadRequest)
            }
        case "tabung":
            jariJariStr := r.URL.Query().Get("jariJari")
            tinggiStr := r.URL.Query().Get("tinggi")
            hitung := r.URL.Query().Get("hitung")

            jariJari, err1 := strconv.ParseFloat(jariJariStr, 64)
            tinggi, err2 := strconv.ParseFloat(tinggiStr, 64)
            if err1 != nil || err2 != nil {
                http.Error(w, "Invalid parameters 'jariJari' or 'tinggi'", http.StatusBadRequest)
                return
            }

            switch hitung {
            case "volume":
                volume := Tabung{JariJari: jariJari, Tinggi: tinggi}.HitungVolume()
                fmt.Fprintf(w, "Volume: %f", volume)
            case "luasPermukaan":
                luasPermukaan := Tabung{JariJari: jariJari, Tinggi: tinggi}.HitungLuasPermukaan()
                fmt.Fprintf(w, "Luas Permukaan: %f", luasPermukaan)
            default:
                http.Error(w, "Invalid parameter 'hitung'", http.StatusBadRequest)
            }
        default:
            http.Error(w, "Invalid bangun ruang", http.StatusBadRequest)
        }
    })

    http.ListenAndServe(":8080", router)
}
