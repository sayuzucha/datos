package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"simulator/domain/models"
)

type Material struct {
	Tipo    string
	PesoMin float64
	PesoMax float64
}

func main() {
	rand.Seed(time.Now().UnixNano())

	
	materiales := []Material{
		{"plastico", 20, 60}, 
		{"lata", 12, 25},     
	}

	currentWeight := 0.0
	maxWeight := 1000.0 
	filling := true

	for i := 0; i < 1000; i++ {
		if filling {
			
			mat := materiales[rand.Intn(len(materiales))]

			add := mat.PesoMin + rand.Float64()*(mat.PesoMax-mat.PesoMin)
			currentWeight += add

			fmt.Printf("➕ Se agregó una %s (%.2f g)\n", mat.Tipo, add)

			if currentWeight >= maxWeight {
				filling = false
				fmt.Println("🧺 Contenedor lleno, vaciando...")
			}
		} else {
			
			currentWeight -= 150 + rand.Float64()*100
			if currentWeight <= 0 {
				currentWeight = 0
				filling = true
				fmt.Println("♻️ Contenedor vacío, reiniciando llenado...")
			}
		}

		data := models.HX711{
			Prototype_id: "HX_001",
			Weight:       fmt.Sprintf("%.2f", currentWeight),
		}

		jsonData, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println(string(jsonData))

		time.Sleep(1 * time.Second)
	}
}
