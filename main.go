package main

import (
	"errors"
	"fmt"
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
	config.AllowOrigins = []string{"http://localhost:8080"} // Add your Svelte app's origin
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
	case "+", "-", "*", "/":
		result, err := performCalculation(request.N1, request.N2, request.Operacion)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Guardar en historial
		historyEntry := HistoryEntry{
			Fecha:     time.Now().Format("2006-01-02 15:04:05"),
			Operacion: fmt.Sprintf("%f %s %f", request.N1, request.Operacion, request.N2),
			Resultado: result,
		}
		history = append(history, historyEntry)

		c.JSON(http.StatusOK, CalculadoraResponse{Resultado: result})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operación no válida"})
	}
}

func performCalculation(n1, n2 float64, operacion string) (float64, error) {
	switch operacion {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		if n2 != 0 {
			return n1 / n2, nil
		} else {
			// Manejar división por cero con un error
			return 0, errors.New("División por cero")
		}
	default:
		return 0, errors.New("Operación no válida")
	}
}

func historyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, history)
}
