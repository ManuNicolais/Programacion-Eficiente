# Practico 2 - Estructuras y tecnicas de programacion eficiente

## Introduccion

Este practico analiza distintas formas de resolver problemas frecuentes de
programacion eficiente. Se comparan algoritmos que producen el mismo
resultado, pero que utilizan diferentes estructuras de datos o estrategias de
calculo.

Los ejercicios muestran cinco ideas centrales:

- busqueda en una lista frente a busqueda en un set;
- costo de comparar todos los elementos frente a consultar un mapa;
- generacion de datos completos frente a generacion bajo demanda;
- busqueda lineal frente a busqueda binaria;
- recursion directa frente a memoizacion.

El objetivo no es solamente obtener una respuesta correcta. Tambien se busca
entender como crece el tiempo de ejecucion, cuanta memoria necesita cada
solucion y en que situaciones conviene elegir una alternativa sobre otra.

## Indice

1. [Metodologia](#1-metodologia)
2. [Conceptos teoricos](#2-conceptos-teoricos)
3. [Ejercicio 1: Usuarios autorizados](#3-ejercicio-1-usuarios-autorizados)
4. [Ejercicio 2: Quien compro que](#4-ejercicio-2-quien-compro-que)
5. [Ejercicio 3: El sensor espacial](#5-ejercicio-3-el-sensor-espacial)
6. [Ejercicio 4: Encontrando una coordenada perdida](#6-ejercicio-4-encontrando-una-coordenada-perdida)
7. [Ejercicio 5: El robot que repite calculos](#7-ejercicio-5-el-robot-que-repite-calculos)
8. [Conclusiones](#8-conclusiones)
9. [Como ejecutar](#9-como-ejecutar)

## 1. Metodologia

### 1.1 Entorno de ejecucion

- **Sistema operativo:** Windows (PowerShell)
- **Lenguaje:** Go
- **Modulo:** `practico2`

Cada ejercicio esta implementado en un archivo `.go` independiente. Todos
pertenecen al paquete `main` y tienen su propia funcion `main`, por lo que se
ejecutan individualmente.

### 1.2 Medicion y analisis

Los ejercicios que comparan rendimiento utilizan `time.Now()` y
`time.Since()` para medir el tiempo de ejecucion. Estas mediciones dependen del
hardware, del sistema operativo y del estado del proceso, por lo que sirven
principalmente para observar tendencias.

Para comparar algoritmos de forma mas rigurosa se debe mantener constante el
conjunto de datos, repetir varias veces cada medicion y evitar que la salida
por pantalla forme parte del intervalo medido. La complejidad algoritmica
permite complementar las mediciones y predecir como se comportara cada
solucion cuando crece el tamaño de la entrada.

## 2. Conceptos teoricos

### 2.1 Complejidad temporal y espacial

La notacion Big O describe como crece el costo de un algoritmo en funcion del
tamaño de la entrada. No indica un tiempo exacto, sino el comportamiento
dominante cuando la entrada se hace grande.

- **O(1):** el costo no depende del tamaño de la entrada;
- **O(log n):** el problema se reduce, aproximadamente, a la mitad en cada paso;
- **O(n):** se recorre la entrada una vez;
- **O(n log n):** combina recorridos lineales con divisiones del problema;
- **O(n^2):** se comparan pares de elementos o se anidan dos recorridos;
- **O(2^n):** el trabajo crece exponencialmente por ramificaciones repetidas.

El espacio auxiliar mide la memoria adicional que necesita el algoritmo,
ademas de los datos originales. Un algoritmo puede ser rapido y utilizar mas
memoria, o ahorrar memoria a cambio de tardar mas.

### 2.2 Listas, mapas y canales

Buscar un valor en un slice sin informacion adicional requiere recorrer sus
elementos y cuesta O(n) en el peor caso. En cambio, un `map` de Go puede usarse
como set: las claves representan los valores presentes y una consulta cuesta
O(1) en promedio. Construir ese mapa agrega un costo inicial O(n) y un espacio
extra O(n).

Un slice guarda todos sus elementos en memoria. Un canal con una goroutine
permite producir valores progresivamente y consumirlos a medida que llegan,
evitando guardar la secuencia completa.

### 2.3 Ordenamiento, recursion y memoizacion

La busqueda binaria necesita datos ordenados: cada comparacion con el elemento
central permite descartar la mitad de los candidatos y lograr O(log n). Sin
orden, esa decision no es valida y la busqueda binaria puede devolver un
resultado incorrecto.

La recursion directa puede resolver correctamente un problema, pero repetir
las mismas llamadas hace que su costo crezca rapidamente. La memoizacion
guarda resultados ya calculados en un mapa. De esta forma, cada subproblema se
resuelve una sola vez y el costo total puede bajar de exponencial a lineal.

## 3. Ejercicio 1 - Usuarios autorizados

## Consigna 1 y 2

En `ej1.go` estan las dos soluciones:

- autorizadosConList: recorre cada intento y por cada uno recorre toda la lista de usuarios buscando coincidencia.
- autorizadosConSet: primero carga los usuarios en un map (set) y despues recorre los intentos consultando el map.

## Consigna 3 - Complejidad

- Solucion con list: O(n * m), donde n es la cantidad de intentos y m la cantidad de usuarios autorizados. Por cada intento se recorre toda la lista.
- Solucion con set: O(n + m). Construir el set cuesta O(m) y luego cada consulta al map es O(1) en promedio, entonces revisar los n intentos cuesta O(n).

## Consigna 4 - Que usar segun el caso

a. 100 usuarios: con tan pocos usuarios la diferencia es minima, cualquiera de las dos sirve. La list es simple y no hace falta optimizar.

b. 1.000.000 de usuarios: conviene el set, porque la busqueda en la list se vuelve muy lenta al crecer m.

c. millones de intentos: tambien conviene el set, porque con la list el costo se multiplica (n * m) y se vuelve impracticable. Con el set el costo crece linealmente.

En resumen: mientras mas grande sea la cantidad de usuarios y/o de intentos, mas conviene el set.

## Consigna 5 - Medicion de tiempos

El programa mide el tiempo de ejecucion de ambas soluciones con el paquete time y lo imprime al final.

Ejemplo de salida con 100.000 usuarios y 20.000 intentos:

```
Tiempo con list: 11.55s
Tiempo con set: 33ms
```

Se ve claramente que el set es mucho mas rapido cuando las cantidades son grandes.

## 4. Ejercicio 2 - Quien compro que

## Consigna 1 - Dos versiones

En ej2.go estan las dos soluciones:

- compradoresConList (primera version): compara cada cliente contra todos los compradores, uno por uno.
- compradoresConSet (segunda version, optimizada): carga los compradores en un map y despues solo consulta el map por cada cliente.

## Consigna 2 - Complejidad

- Primera version (list): O(n * m), donde n es la cantidad de clientes y m la cantidad de compradores. Por cada cliente recorre toda la lista de compradores.
- Segunda version (set): O(n + m). Armar el set cuesta O(m) y luego cada consulta es O(1) en promedio, entonces revisar los n clientes cuesta O(n).

## Consigna 3 - Que mejora el rendimiento

El cambio clave es reemplazar la busqueda lineal dentro del segundo for (recorrer compradores) por una consulta directa a un map (hash set). Asi se evita comparar cada cliente contra todos los compradores y la busqueda pasa a ser O(1) en vez de O(m).

## Consigna 4 - Prueba con distintos tamanios

Se probo con 10.000, 100.000 y 1.000.000 de elementos. Con 1.000.000 la version con list se omite porque tarda demasiado (crece de forma cuadratica).

Resultados obtenidos:

```
Tamanio: 10000
Tiempo con list: 450ms
Tiempo con set: 3.5ms

Tamanio: 100000
Tiempo con list: 47.7s
Tiempo con set: 31ms

Tamanio: 1000000
Tiempo con list: (omitido, demasiado lento)
Tiempo con set: 822ms
```

Se confirma que la version con set escala mucho mejor a medida que crecen las listas.

## 5. Ejercicio 3 - El sensor espacial

Go no tiene yield como Python, el equivalente idiomatico es usar un channel con una goroutine que va enviando los valores uno por uno.

## Soluciones

- cuadradosLista: genera una lista (slice) completa con todos los cuadrados de una vez.
- cuadradosGenerador: usa un channel y una goroutine, va calculando y enviando cada cuadrado a medida que se pide, uno por uno.

## Consigna 1 - Cual usa mas memoria

La lista usa mas memoria, porque guarda todos los n valores en memoria al mismo tiempo. El generador practicamente no usa memoria extra, porque solo mantiene el valor actual, sin importar cuantos se pidan.

En la prueba con n = 1.000.000:

```
Lista: memoria usada: ~8.000.000 bytes
Generador: memoria usada: ~870 bytes
```

## Consigna 2 - Que sucede si n = 10.000.000

Con la lista, la memoria usada crece en la misma proporcion (10 veces mas que con 1.000.000), pudiendo llegar a ser un problema si el sistema tiene memoria limitada o si n sigue creciendo.

Con el generador, la memoria se mantiene practicamente constante, porque nunca se guardan todos los valores juntos, solo se va entregando uno a la vez.

## Consigna 3 - Cual elegir para procesar datos uno por uno

El generador. Si el sistema procesa los datos de a uno, no tiene sentido generar y guardar la lista completa en memoria. El generador permite arrancar a procesar el primer valor apenas esta disponible, sin esperar a que se calculen todos, y usando mucha menos memoria.

## Consigna 4 - Que significa procesamiento "bajo demanda"

Significa que los valores se calculan y entregan solo en el momento en que se necesitan, uno por uno, en vez de calcularlos todos por adelantado. El generador no hace nada hasta que alguien le pide el proximo valor; asi se ahorra memoria y tiempo cuando no se necesitan todos los datos de una vez (por ejemplo, si el programa corta el procesamiento antes de llegar al final).

## 6. Ejercicio 4 - Encontrando una coordenada perdida

## Consigna 1 y 2 - Implementaciones

En ej4.go estan las dos busquedas:

- busquedaLineal: recorre la lista de principio a fin comparando elemento por elemento hasta encontrar el valor.
- busquedaBinaria: usa el hecho de que la lista esta ordenada, comparando contra el elemento del medio y descartando la mitad de la lista en cada paso.

## Consigna 3 - Complejidad temporal

- Busqueda lineal: O(n). En el peor caso hay que recorrer todos los elementos.
- Busqueda binaria: O(log n). En cada paso se descarta la mitad de la lista restante.

## Consigna 4 - Por que la busqueda binaria necesita la lista ordenada

Porque su logica se basa en comparar el valor buscado contra el elemento del medio para decidir de que lado esta: si es mayor, descarta la mitad izquierda; si es menor, descarta la mitad derecha. Esa decision solo es valida si la lista esta ordenada, porque asume que todos los valores a la izquierda del medio son menores y todos los de la derecha son mayores. Si la lista no esta ordenada, esa suposicion es falsa y el algoritmo puede descartar la mitad donde en realidad estaba el valor buscado.

## Consigna 5 - Comparacion con 10 millones de elementos

Con una lista de 10.000.000 de elementos:

- La busqueda lineal, en el peor caso, necesitaria hasta 10.000.000 de comparaciones.
- La busqueda binaria, en el peor caso, necesitaria como mucho log2(10.000.000) ≈ 24 comparaciones.

La diferencia es enorme: la busqueda binaria es muchisimo mas eficiente a gran escala, siempre y cuando la lista este ordenada. Si la lista no esta ordenada y hay que ordenarla antes, ese costo (O(n log n)) puede no valer la pena si solo se va a hacer una busqueda una vez, pero si se van a hacer muchas busquedas sobre la misma lista, conviene ordenarla una vez y despues usar busqueda binaria.

## 7. Ejercicio 5 - El robot que repite calculos

## Consigna 1 - Solucion recursiva directa

costoRecursivo en ej5.go implementa la definicion tal cual: costo(n) = costo(n-1) + costo(n-2), con casos base costo(0) = 0 y costo(1) = 1.

## Consigna 2 - Complejidad

La version recursiva directa tiene complejidad O(2^n), porque cada llamada genera dos llamadas nuevas y se repiten muchisimos calculos ya hechos (por ejemplo costo(3) se vuelve a calcular muchas veces dentro del arbol de recursion).

## Consigna 3 - Version con memoizacion

costoMemo guarda cada resultado ya calculado en un map. Antes de calcular costo(n), primero se fija si ya esta en el map; si esta, lo devuelve directo sin volver a calcular. Con esto la complejidad baja a O(n), porque cada valor de 0 a n se calcula una sola vez.

## Consigna 4 - Comparacion conceptual

- La version recursiva directa es mas simple de escribir y entender, pero es muy ineficiente para valores de n grandes porque repite el mismo calculo miles o millones de veces.
- La version con memoizacion agrega una estructura extra (el map) para recordar resultados ya calculados, a cambio de una mejora enorme en tiempo de ejecucion. Usa un poco mas de memoria (O(n) para guardar el map), pero el ahorro en tiempo compensa ampliamente esa memoria extra.

## Consigna 5 - Que trabajo redundante elimina la memoizacion

En la version recursiva directa, para calcular costo(n) se vuelve a calcular costo(n-2), costo(n-3), etc. muchas veces, porque distintas ramas del arbol de recursion piden el mismo valor una y otra vez. La memoizacion elimina ese trabajo repetido: una vez que un valor se calcula, se guarda, y la proxima vez que se necesita se reutiliza directamente en vez de volver a calcularlo desde cero.

## 8. Conclusiones

1. **La estructura de datos influye directamente en el rendimiento.** Reemplazar busquedas lineales por consultas a un mapa cambia el costo de O(n * m) a O(n + m) en los ejercicios de usuarios y compradores.

2. **La complejidad permite anticipar el crecimiento.** Una solucion que funciona con listas pequeñas puede volverse impracticable cuando aumentan la cantidad de elementos o de consultas.

3. **El uso de memoria tambien forma parte de la decision.** Una lista conserva todos los resultados, mientras que un generador puede entregar valores progresivamente y mantener un consumo mucho menor.

4. **Las condiciones del problema determinan el algoritmo adecuado.** La busqueda binaria es muy eficiente, pero solo es correcta cuando los datos estan ordenados.

5. **Evitar trabajo repetido puede cambiar por completo el rendimiento.** La memoizacion transforma una recursion exponencial en una solucion lineal a cambio de usar memoria adicional.

6. **La notacion Big O debe complementarse con mediciones.** Los tiempos reales dependen de los datos, del hardware y de los costos constantes de cada implementacion.

## 9. Como ejecutar

### Requisitos previos

- Tener [Go](https://go.dev/dl/) instalado.
- Tener disponible Go en el `PATH` de PowerShell.
- Estar ubicado en la carpeta del proyecto.

### Paso a paso

```powershell
# Navegar a la carpeta del practico
cd "C:\Users\manue\OneDrive\Documentos\Programacion\PEF\practico2"

# Ejecutar cada ejercicio individualmente
go run ej1.go    # Usuarios autorizados
go run ej2.go    # Quien compro que
go run ej3.go    # Generador de cuadrados
go run ej4.go    # Busqueda lineal y binaria
go run ej5.go    # Recursion y memoizacion
```

### Verificacion estatica

Como cada archivo tiene su propia funcion `main`, se debe compilar o ejecutar
cada ejercicio por separado. Los comandos generales del modulo intentarian
compilar varias funciones `main` juntas.

```powershell
go vet .\ej1.go
go vet .\ej2.go
go vet .\ej3.go
go vet .\ej4.go
go vet .\ej5.go
```

### Estructura del proyecto

```
practico2/
	|- go.mod       (modulo Go: practico2)
	|- ej1.go       (Usuarios autorizados)
	|- ej2.go       (Quien compro que)
	|- ej3.go       (Generador de cuadrados)
	|- ej4.go       (Busqueda lineal y binaria)
	|- ej5.go       (Recursion y memoizacion)
	|- README.md    (Este archivo)
```
