# Programacion Eficiente (PEF)

## Materia: Programacion Eficiente

**Carrera:** Ingenieria en Informatica
**Lenguaje utilized:** Go (Golang)
**Metodologia:** Clases teorico-practicas, laboratorio de programacion, discusion grupal y estudio de casos

---

## Proposito

Identificar las distintas problematicas que afectan la performance de un software y adquirir los conocimientos necesarios para disminuir su impacto mediante tecnicas de optimizacion de codigo y mediante el aprovechamiento del hardware disponible.

## Objetivos

- Encontrar las fuentes de ineficiencias de un programa.
- Ordenar las fuentes de ineficiencias segun su impacto en la eficiencia global.
- Aplicar las facilidades del lenguaje de programacion utilizado para mejorar la eficiencia del programa.
- Optimizar los bucles criticos reescribiendolos para aprovechar al maximo las posibilidades del procesador y del sistema de memoria.
- Combinar varios algoritmos de manera balanceada, cuidando de que las optimizaciones locales no interfieran con la optimizacion global.
- Comprender las posibilidades avanzadas que ofrece el hardware para aumentar la eficiencia.
- Conocer herramientas para analizar la performance del codigo fuente.

---

## Contenidos

### Unidad I: Arquitectura de Computadoras Modernas

- Performance del software
- Arquitecturas de microprocesadores
- Tecnologias de memoria
- Memorias cache
- Precision en numeros de punto flotante
- Estandar IEEE 754

### Unidad II: Programacion y Profiling

- Rol del compilador
- Optimizacion a cargo del compilador
- Medicion de tiempo de ejecucion
- Profiling mediante herramientas provistas en las distribuciones basicas
- Profiling con herramientas mas especificas
- Deteccion de ineficiencias del uso de la memoria
- Deteccion de no liberacion de memoria

### Unidad III: Optimizacion Manual de Codigo Fuente

- Eliminacion de acciones innecesarias
- Function inlining
- Optimizacion de ciclos de repeticion
- Patrones de acceso a memoria
- Optimizaciones utilizando metodos matematicos mas eficientes
- Uso eficiente de objetos

### Unidad IV: Seguimiento del Codigo

- Estudio de llamadas en un programa
- Estudio de interaccion entre objetos
- Estudio de interaccion entre hilos
- Estudio de interaccion de procesos en un programa paralelo

---

## Practicos del Curso

### [Practico 1: Algoritmos de Busqueda, Conteo, Deteccion y Ordenamiento](practico1-algoritmos-de-busqueda/)

Analisis empirico de algoritmos fundamentales. Se implementan, miden y comparan distintos enfoques para resolver cada problema, verificando la complejidad asintotica de forma experimental.

| Ejercicio | Problema | Algoritmos evaluados | Complejidad |
|-----------|----------|---------------------|-------------|
| Ej1 | Busqueda lineal | Manual (range) vs `slices.Contains` | O(n) |
| Ej2 | Conteo de pares | Modulo, bitwise, sin bifurcacion | O(n) |
| Ej3 | Repetidos | Fuerza bruta O(n^2), ordenamiento O(n log n), mapa O(n) | O(n^2) / O(n) |
| Ej4 | Ordenamiento | Bubble Sort, Bubble optimizado, Insertion Sort, `slices.Sort` | O(n^2) / O(n log n) |

**Conceptos clave:** Medicion con `time.Now()`, `testing.Benchmark`, complejidad empirica vs teorica, influencia del hardware (predictor de saltos, cache).

**Conclusiones principales:**
- La complejidad asintotica describe el crecimiento, no el rendimiento absoluto.
- Un algoritmo O(n^2) puede ser mas rapido que O(n) cuando la terminacion temprana es probable.
- El compilador ya realiza micro-optimizaciones (sustitucion de `%2` por `&1`).
- Medir es imprescindible: varias hipotesis resultaron incorrectas.

---

### [Practico 2: Estructuras de Datos y Tecnicas de Programacion](practico2-Estructuras-y-Tecnicas/)

Seleccion y comparacion de estructuras de datos y tecnicas de programacion para resolver problemas de forma eficiente.

| Ejercicio | Problema | Enfoques comparados |
|-----------|----------|-------------------|
| Ej1 | Usuarios autorizados | Busqueda en lista O(n*m) vs mapa O(n+m) |
| Ej2 | Quienes compraron que | Lista vs set a escala (hasta 1M elementos) |
| Ej3 | Sensor espacial | Slice completo vs generador con canales/goroutines |
| Ej4 | Coordenada perdida | Busqueda lineal O(n) vs binaria O(log n) |
| Ej5 | Robot recalculando | Recursion directa O(2^n) vs memoizada O(n) |

**Conceptos clave:** Costo de buscar en listas vs maps, selectividad, lazy evaluation, canales y goroutines, busqueda binaria, memoizacion basica.

---

### [Practico 3: Memoizacion y Caching](practico3-Memoizacion-y-Caching/)

Programacion dinamica aplicada a problemas clasicos, transformando soluciones exponenciales en polinomicas mediante memoizacion.

| Ejercicio | Problema | Clave memo | Sin memo | Con memo |
|-----------|----------|------------|----------|----------|
| Ej1 | Caminos en cuadricula | `"x,y"` | O(2^(m+n)) | O(m*n) |
| Ej2 | Subconjunto suma objetivo | `"indice,objetivo"` | O(2^n) | O(n*objetivo) |
| Ej3 | Word Break | string restante | O(2^len(S)) | O(len(S)^2) |
| Ej4 | Combinaciones de monedas | `"indice,total"` | O(2^n) | O(n*total) |

**Conceptos clave:** Memoizacion como caching de funciones, patron comun (identificar subproblemas, clave unica, almacenar, consultar), transformacion de complejidad exponencial a polinomica.

---

## Practicos Anteriores (versiones previas)

| Carpeta | Contenido |
|---------|-----------|
| `TP1/` | Ejercicios introductorios: HelloWorld, busqueda lineal, conteo de pares |
| `Practica_1_Fundamentos/` | Benchmark con graficos (gonum/plot), busqueda lineal, binaria, bubble sort |
| `Practica_2/` | Version preliminar de estructuras y tecnicas |
| `Practica_3/` | Version preliminar de memoizacion |
| `PPT/` | Ejemplos de codigo de las diapositivas (Fibonacci con/sin memoizacion) |

---

## Proyecto Final

El proyecto final se encuentra scaffolded en `ProyectoFinal/Proyecto-Programacion-Eficiente/` con la estructura estandar de Go:

```
cmd/main.go         // Punto de entrada
internal/           // Paquetes internos (pendiente de implementar)
```

---

## Bibliografia

- Jack Shirazi. *Java Performance Tuning*. O'Reilly (2007)
- Kevin Dowd y Charles R. Severance. *High Performance Computing*. O'Reilly (1998)
- Kevin R. Wadleigh y Isom L. Crawford. *Software Optimization for High-performance Computing*. Prentice Hall (2000)
- David J. Lilja. *Measuring Computer Performance: A Practitioner's Guide*. Cambridge University Press (2000)
- Dov Bulka y David Mayhew. *Efficient C++: Performance Programming Techniques*. Addison-Wesley (2011)
- Scott Meyers. *Effective C++: 55 Specific Ways to Improve Your Programs and Designs*. Addison Wesley (2011)

---

## Forma de evaluacion

**Regularizacion:** 3 trabajos praticos grupales, cada uno consistente en el estudio de un software que se optimizara y corregira en base a los temas abordados, mas elaboracion de un informe escrito.

**Examen final:** Estudio de un caso, elaboracion de un informe escrito y defensa oral.

---

## Estructura del repositorio

```
PEF/
  |- practico1-algoritmos-de-busqueda/     Practico 1: Busqueda y ordenamiento
  |- practico2-Estructuras-y-Tecnicas/     Practico 2: Estructuras y tecnicas
  |- practico3-Memoizacion-y-Caching/      Practico 3: Memoizacion y caching
  |- ProyectoFinal/                        Proyecto final del curso
  |- TP1/                                  Ejercicios introductorios
  |- Practica_1_Fundamentos/               Version previa practico 1
  |- Practica_2/                           Version previa practico 2
  |- Practica_3/                           Version previa practico 3
  |- PPT/                                  Ejemplos de diapositivas
  |- README.md                             Este archivo
```
