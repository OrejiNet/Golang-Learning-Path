package gentleman

import (
	"fmt"
	"time"
)

// go -> GoRoutine

// GoRoutine -> Es un hilo de ejecución ligero que se utiliza para realizar tareas concurrentes en Go.
// Es una función que se ejecuta de manera asíncrona y puede ser creada utilizando la palabra clave "go".
// Las GoRoutines son gestionadas por el runtime de Go y pueden comunicarse entre sí a través de canales (channels)
// para sincronizar su ejecución y compartir datos. Son una parte fundamental del modelo de concurrencia de Go,
// permitiendo a los desarrolladores escribir código concurrente de manera sencilla y eficiente.

// Chanel -> Es una estructura de datos que se utiliza para comunicar y sincronizar GoRoutines en Go.

func CallHello(canal chan <- string) {
	// Enviar un mensaje al canal
	time.Sleep(1 * time.Second) // Simular una tarea que tarda un poco
	canal <- "Hola desde la go routine!"
}  

func PrintMessage(canal <- chan string) {
	// Recibir un mensaje del canal
	message := <- canal
	fmt.Println(message)
}


