# Trucode Final Challenge

Este proyecto es la solución al challenge final del programa de TrucodeV3. 
Las instrucciones del challenge junto con el detalle de los criterios de aceptación se pueden encontrar en el [Instrucciones y criterios de aceptación](./instructions.md).

La aplicación está diseñada para gestionar datos de un dataset creado por Barry Becker basado en un censo de 1994 que contiene más de 30.000 registros con datos demográficos, datos necesarios para comprender qué tipo de personas ganaban más o menos de $50,000 USD, entre otros. 

<details>
<summary>Estado de Implementación de Requerimientos</summary>

### 🔵 Historia de Usuario 1: Filtrado y Ordenamiento de Datos
- [x] Implementar filtros y paginación en el backend  
  *Incluye filtros por edad, educación, estado civil, ocupación e ingresos*
- [x] Opciones de filtrado visibles en la interfaz  
  *Implementado con v-select y v-range-slider*
- [x] Datos reactivos a los filtros seleccionados  
  *Actualización inmediata de la tabla al aplicar filtros*

### 🟣 Historia de Usuario 2: Autenticación y Persistencia
- [x] Interfaz con opciones de registro e inicio de sesión  
  *Usando JWT para manejo de sesiones*
- [x] Carga automática de configuraciones personalizadas después del login  
  *Los filtros se mantienen entre sesiones y en diferentes dispositivos*

### 🟡 Historia de Usuario 3: Diseño Responsive
- [x] Aplicación accesible en diferentes tamaños de pantalla  
  *Adaptada para móviles, tablets y desktop*  
  *Implementado con clases responsive de Vuetify (d-md-none/d-md-flex) y Tailwind*
- [x] Funcionalidades utilizables en dispositivos móviles  
  *Optimización especial para controles de filtrado en móviles*  
  *el Range slider se adapta a vista vertical en móviles usando flex-col y vista horizontal en desktop usando flex-row*

### 🟢 Historia de Usuario 4: Exportación de Datos (Opcional)
- [x] Opción para importar y actualizar datos  
  *La opción de importación recibe un .csv que debe contener el ID del registro con sus demás actualizados según se quiera modificar*  
  *Se da aviso de importación exitosa con v-snackbar*
- [x] Exportación a CSV con datos seleccionados  
  *La exportación respeta los filtros aplicados y exporta la data de todas las páginas*  
  *Backend procesa la solicitud con flag export=true y Frontend maneja la descarga del archivo mediante Blob*

### 🟠 Historia de Usuario 5: Estadísticas Resumidas (Opcional)
- [x] Estadísticas resumidas visibles en la interfaz  
- [x] Visualizaciones de distribución de ingresos por educación y ocupación  

### ⚪ Historia de Usuario 6: Visualizaciones Interactivas (Opcional)
- [x] Implementación de al menos cuatro visualizaciones interactivas  

### 🟣 Implementación y Verificaciones Finales
- [x] Despliegue de la aplicación en plataforma web  
  *Desplegado en Vercel (frontend) y AWS EC2 (backend)*
- [x] Documentación de pasos para el despliegue  
- [x] Verificación de requisitos obligatorios  
  *Todos los requisitos core HUs 1, 2, 3 completados*
</details>

## Tecnologías Principales

- **Frontend**: Vue 3, Pinia, Vue Router, Vuetify, Tailwind CSS
- **Backend**: Go, Gin, GORM
- **Base de Datos**: PostgreSQL
- **Despliegue**: Docker, AWS EC2, Vercel

## Estructura del Proyecto

El proyecto está dividido en tres partes: el frontend, el backend y el entorno de Docker.

### Frontend (trucode_data_management_app)

#### Características Principales
- **Autenticación**: Sistema completo de registro y login
- **Gestión de Sesión**: Manejo de tokens JWT y persistencia de sesión en todos los dispositivos
- **Filtrado Avanzado**: 
  - Filtros por educación, estado civil, ocupación e ingresos
  - Selector de rango de edad con slider
  - Persistencia de filtros entre sesiones
- **Visualización de Datos**: 
  - Tabla de datos paginada
  - Ordenamiento por columnas
  - Importación de CSV para actualizar registros en masivo
  - Exportación a CSV
- **Diseño Responsive**: Adaptable a dispositivos móviles y desktop

<details>
<summary>Estructura de Directorios Frontend</summary>

#### trucode_data_management_app

```
src/
├── components/
│ ├── DataTable.vue
│ ├── FilterOptions.vue
│ ├── LoginForm.vue
│ └── RegisterForm.vue
├── router/
│ └── index.ts
├── services/
│ ├── auth.ts
│ ├── filters.ts
│ ├── peopleData.ts
│ └── utils.ts
├── stores/
│ ├── auth.ts
│ └── filters.ts
├── views/
│ ├── HomeView.vue
│ ├── LoginView.vue
│ └── RegisterView.vue
└── App.vue
```
</details>

### Backend (trucode_data_management_api)

#### Funcionalidades Principales
- **Autenticación**:
  - Registro y login de usuarios
  - Generación y validación de JWT
  - Manejo y persistencia de sesiones
- **Gestión de Datos**:
  - CRUD de datos de personas
  - Filtrado múltiple
  - Ordenamiento
  - Paginación
  - Exportación e importación de datos
- **Persistencia de Filtros**:
  - Guardado de preferencias de filtros por usuario
  - Recuperación de filtros en nuevas sesiones

#### Modelo de Datos

- **User**: Modelo para la gestión de usuarios registrados.
- **PeopleData**: Modelo para la gestión de datos de personas, con campos como edad, educación, estado civil, ocupación, y más.
- **Filters**: Modelo para la gestión de filtros por edad, educación, ocupación, ingresos y estado civil para cada usuario registrado. 

<details>
<summary>Estructura de Directorios Backend</summary>
trucode_data_management_api/

```
├── auth/
│ ├── handler.go
│ └── router.go
├── database/
│ └── main.go
├── data/
│ ├── insertData.go
│ ├── loadData.go
│ └── source.data
├── models/
│ ├── peopleData.go
│ └── user.go
├── filters/
│ ├── handler.go
│ └── router.go
├── shared/
│ ├── jwt.go
│ ├── middlewares.go
│ └── sessions.go
└── main.go
```
</details>

### Docker (trucode_data_management_tools)
Configuración para el despliegue de la aplicación usando Docker Compose:
- PostgreSQL para la base de datos
- Servidor Go para la API
- Configuración de red y volúmenes

## Despliegue

- **Frontend**: Desplegado en Vercel
- **Backend**: Desplegado en AWS EC2 usando Docker

<details>
<summary>Estructura de Directorios Docker</summary>
trucode_data_management_tools/

```
├── Dockerfile
└── docker-compose.yml
```
</details>

## URL de la Aplicación

El proyecto puede visualizarse en: [https://trucode-final-challenge-gavrojas.vercel.app/](https://trucode-final-challenge-gavrojas.vercel.app/)
