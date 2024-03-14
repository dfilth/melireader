# README

## Procesos, Hilos y Corrutinas

### Procesos

Los procesos son útiles cuando necesitas ejecutar programas separados e independientes entre sí. Son ideales para situaciones donde es necesario manejar múltiples solicitudes simultáneamente, como en servidores web. Cada proceso tiene su propio espacio de memoria, lo que permite aislar recursos y aprovechar múltiples núcleos de CPU en sistemas multiprocesador.

### Hilos

Los hilos son útiles cuando se quiere realizar múltiples tareas concurrentemente dentro de un mismo programa. Son ideales para compartir recursos de manera eficiente y realizar operaciones en paralelo, como en aplicaciones de edición de imágenes. Los hilos comparten el mismo espacio de memoria y son más livianos que los procesos, lo que los hace adecuados para una comunicación rápida entre tareas.

### Corrutinas

Las corrutinas, son ideales para manejar operaciones de entrada/salida concurrentemente. Son eficientes en el uso de recursos y pueden gestionarse con un menor número de hilos del sistema operativo. Son útiles para manejar muchas solicitudes de red o lectura/escritura de archivos, facilitando la programación concurrente debido a su sintaxis sencilla y la facilidad para comunicarse entre ellas mediante canales, un ejemplo claro es el challenge de leer archivos y construir una salida mediante consulta de apis externas.

## Paralelismo y Concurrencia

En lugar de realizar consultas de manera secuencial, se puede utilizar el paralelismo y la concurrencia para realizar múltiples consultas simultáneamente. Esto se puede lograr utilizando gorutinas en Go, hilos en otros lenguajes de programación o bibliotecas que admitan solicitudes HTTP concurrentes. EL mismo ejemplo de procesar archivos, transacciones bancarias, e-commerce promociones.

## Análisis de complejidad

### AlfaDB

AlfaDB ofrece una complejidad de consulta de O(1) y una complejidad de escritura de O(n^2). Es adecuada para casos de uso donde la consulta es frecuente y la escritura es menos frecuente, o donde se necesita un acceso rápido a los datos sin importar el tamaño del conjunto de datos. Por ejemplo, podría ser útil en aplicaciones donde se necesita recuperar rápidamente información específica del usuario o datos de configuración.

### BetaDB

BetaDB ofrece una complejidad de consulta y escritura de O(log n). Es más adecuada para casos de uso donde se requiere un equilibrio entre la eficiencia en la consulta y la escritura, y donde el conjunto de datos tiende a ser más grande. Por ejemplo, sería útil en aplicaciones que manejan grandes volúmenes de datos transaccionales o en sistemas de búsqueda donde se necesita un acceso eficiente a grandes conjuntos de datos ordenados.

## Optimización de recursos del sistema operativo

- **Limitación de solicitudes concurrentes:** Controlar el número de solicitudes concurrentes para evitar la sobrecarga del servidor y las restricciones de ancho de banda.
- **Caché de resultados:** Implementar una caché para almacenar los resultados de las consultas previas y evitar consultas repetidas a la API.
- **Paginación:** Dividir las consultas en lotes más pequeños y procesarlos de manera incremental para mejorar la eficiencia en el manejo de grandes cantidades de datos.
- **Optimización de consultas:** Analizar y optimizar las consultas realizadas a la API HTTP para mejorar el rendimiento del sistema.
- **Uso de asincronismo, paralelismo, procesamiento en batches en conjunto con corutinas
- **Monitoreo y ajuste:** Monitorear el rendimiento del sistema y ajustar los parámetros según sea necesario para optimizar el rendimiento.

