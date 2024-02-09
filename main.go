package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CalculadoraRequest struct {
	N1        float64 `json:"n1"`
	N2        float64 `json:"n2"`
	Operacion string  `json:"operacion"`
}

type CalculadoraResponse struct {
	Resultado float64 `json:"resultado"`
}

type HistoryEntry struct {
	Fecha     string  `json:"fecha"`
	Operacion string  `json:"operacion"`
	Resultado float64 `json:"resultado"`
}

var history []HistoryEntry

func main() {
	router := gin.Default()

	//Cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"}
	router.Use(cors.New(config))

	router.POST("/Calcular", calculadoraHandler)
	router.GET("/Historial", historyHandler)

	fmt.Println("Servidor en ejecución en http://localhost:9090")
	router.Run(":9090")
}

func calculadoraHandler(c *gin.Context) {
	var request CalculadoraRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear la solicitud"})
		return
	}

	switch request.Operacion {
	case "+", "-", "x", "÷":
		result, err := performCalculation(request.N1, request.N2, request.Operacion)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Guardar en historial
		historyEntry := HistoryEntry{
			Fecha:     time.Now().Format("2006-01-02"),
			Operacion: fmt.Sprintf("%.2f %s %.2f", request.N1, request.Operacion, request.N2),
			Resultado: result,
		}
		history = append(history, historyEntry)

		c.JSON(http.StatusOK, CalculadoraResponse{Resultado: result})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operación no válida"})
	}
}

func performCalculation(n1, n2 float64, operacion string) (float64, error) {
	result := 0.0
	switch operacion {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "x":
		result = n1 * n2
	case "÷":
		if n2 != 0 {
			result = n1 / n2
		} else {
			// Manejar división por cero con un error
			return 0, errors.New("División por cero")
		}
	default:
		return 0, errors.New("Operación no válida")
	}

	// Redondeo a 2 decimales
	result = math.Round(result*100) / 100

	return result, nil
}

func historyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, history)
}
