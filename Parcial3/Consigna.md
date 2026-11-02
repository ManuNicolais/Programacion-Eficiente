# Proyecto: Documentación e Integración Continua (CI)

Este repositorio contiene los lineamientos y requisitos para la implementación de la documentación externa y el pipeline de Integración Continua (CI) del proyecto.

---

## 1. Documentación externa del proyecto

El equipo debe seleccionar e implementar una herramienta de documentación externa al código fuente que centralice la información funcional, técnica y operativa del sistema. Esta documentación debe mantenerse actualizada, con control de acceso y posibilidades de colaboración.

**Herramientas sugeridas:** Confluence, Microsoft SharePoint, GitHub Wiki, GitBook, MkDocs / Material for MkDocs, Read the Docs, etc.

### Contenido mínimo de la documentación
Cada equipo deberá entregar:
- **Descripción funcional:** Problema que resuelve el sistema, alcance y requisitos.
- **Arquitectura:** Módulos, responsabilidades, dependencias y diagrama de componentes.
- **Instalación y ejecución:** Dependencias, configuración y comandos necesarios.
- **Diseño técnico:** Estructuras de datos, algoritmos importantes y decisiones de implementación.
- **API y contratos:** Entradas, salidas, excepciones y ejemplos de uso (si corresponde).
- **Pruebas:** Estrategia, casos de prueba y resultados.
- **Rendimiento:** Benchmarks, profiling, complejidad temporal y consumo de memoria.
- **Seguridad:** Riesgos identificados, dependencias vulnerables y medidas correctivas.
- **Mantenimiento:** Cómo contribuir, ejecutar los controles y resolver fallos del pipeline.

> **Requisito adicional:** La documentación debe explicar de forma clara cómo ejecutar localmente los mismos controles de calidad y eficiencia que se ejecutan en el entorno de CI.

---

## 2. Integración Continua (CI): Herramientas y controles obligatorios

Se deberá configurar un pipeline de CI que se ejecute automáticamente ante cada `push` y cada `pull request` hacia la rama principal. El código no debe integrarse si incumple las condiciones mínimas de calidad definidas.

**Plataformas aceptadas:** GitHub Actions, GitLab CI/CD, Jenkins, Azure Pipelines, entre otras.

### 2.1. Matriz de controles del proyecto

Cada control debe contar con una herramienta, un comando o configuración reproducible, un reporte y un criterio de aprobación/rechazo.

| Control | Qué deben verificar | Herramientas propuestas |
| :--- | :--- | :--- |
| **Pruebas automatizadas** | Correctitud de las funciones y comportamiento esperado | `pytest`, `unittest` |
| **Benchmark** | Tiempo de ejecución de operaciones representativas | `pytest-benchmark`, `pyperf` |
| **Profiling de CPU** | Funciones y líneas que consumen más tiempo | `cProfile`, `pyinstrument`, `py-spy` |
| **Análisis estático** | Errores potenciales, estilo, construcciones problemáticas y mantenibilidad | `Ruff`, `Pylint` |
| **Cobertura** | Porcentaje de código ejecutado por las pruebas | `coverage.py`, `pytest-cov` |
| **Complejidad** | Complejidad ciclomática y tamaño de las funciones | `Radon`, `Lizard` |
| **Consumo de memoria** | Memoria máxima y asignaciones relevantes | `tracemalloc`, `memory-profiler`, `memray` |
| **Vulnerabilidades** | Patrones de código inseguro y dependencias vulnerables | `Bandit`, `pip-audit`, `Semgrep`, `OSV-Scanner` |
| **Tipado estático** | Inconsistencias en tipos y contratos | `mypy`, `Pyright` |
| **Dependencias** | Versiones reproducibles y actualizaciones de seguridad | `pip-audit`, `Dependabot`, `Renovate` |
| **Documentación** | Construcción del sitio y validez de enlaces | `MkDocs`, `Sphinx`, `markdownlint-cli2`, `Lychee` |

*(Nota: Es conveniente combinar controles complementarios, como `Bandit` para código Python y `pip-audit` para dependencias).*

### 2.2. Requisitos de Seguridad y Análisis de Vulnerabilidades

Herramientas a investigar: **Dependabot, Renovate, CodeQL, OWASP Dependency-Check**.

Se debe cumplir con lo siguiente:
1. Analizar el código fuente.
2. Analizar las dependencias.
3. Detectar secretos expuestos.
4. Generar al menos un reporte legible (y en JSON o SARIF cuando sea posible).
5. Definir una política de bloqueo por severidad (ej. bloquear críticas, revisar altas).
6. Documentar falsos positivos, excepciones y medidas correctivas.

---

## 3. Diseño del pipeline de CI

El pipeline debe organizar los controles en etapas, ejecutando primero los más rápidos y dejando las mediciones costosas para el final. (Las etapas sin dependencias pueden ejecutarse en paralelo).

**Secuencia propuesta:**
1. **Integración del código:** Push / Pull request · Revisión de cambios.
2. **Preparación del entorno:** Instalación reproducible · Dependencias fijadas.
3. **Calidad estática:** Ruff · mypy · Radon · Bandit · Secretos.
4. **Pruebas y cobertura:** pytest · coverage.py · Umbrales mínimos.
5. **Seguridad de dependencias y documentación:** pip-audit · MkDocs · Enlaces.
6. **Rendimiento y memoria:** Benchmark · Profiling · Consumo de memoria.
7. **Resultado final:** Aprobado / Rechazado · Reportes · Artefactos.

### 3.1. Qué debe ocurrir si un control falla

| Control | Condición de rechazo propuesta |
| :--- | :--- |
| **Pruebas** | Una o más pruebas fallan. |
| **Análisis estático** | Existen errores bloqueantes según la configuración. |
| **Cobertura** | Se encuentra por debajo del umbral definido. |
| **Complejidad** | Se supera un límite de complejidad establecido. |
| **Tipado** | Existen errores que la política considera bloqueantes. |
| **Seguridad** | Se detecta una vulnerabilidad que supera el umbral de rechazo. |
| **Documentación** | No se puede generar correctamente o falla una validación obligatoria. |
| **Benchmark** | Se detecta una regresión significativa respecto de una línea base confiable. |
| **Memoria** | Se supera un límite justificado y medido en condiciones reproducibles. |

*Nota sobre tolerancias: Mientras que las pruebas funcionales son deterministas, los benchmarks y las mediciones de memoria pueden requerir tolerancias, repetición y comparación con una línea base para evitar falsos positivos por el entorno.*
