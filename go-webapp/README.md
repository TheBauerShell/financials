# Go Webapp

Dieses Projekt ist eine Webanwendung, die mit Go für das Backend und Vue.js für das Frontend entwickelt wurde. Die Anwendung ermöglicht die Verwaltung von IT-Projekten und deren Kosten, unterstützt verschiedene Datenbanken und integriert Keycloak für die Authentifizierung.

## Projektstruktur

- **backend/**: Enthält den Go-Code für das Backend.
  - **main.go**: Einstiegspunkt der Anwendung.
  - **controllers/**: Logik zur Verwaltung von Projekten.
  - **models/**: Definition der Datenmodelle.
  - **routes/**: Definition der API-Routen.
  - **services/**: Datenbankoperationen und -verbindungen.
  - **config/**: Konfigurationseinstellungen.
  - **migrations/**: Datenbankmigrationen.
  - **utils/**: Hilfsfunktionen, z.B. für Keycloak.

- **frontend/**: Enthält den Vue.js-Code für das Frontend.
  - **src/**: Quellcode der Anwendung.
    - **App.vue**: Hauptkomponente.
    - **main.js**: Einstiegspunkt für Vue.js.
    - **components/**: Wiederverwendbare Komponenten.
    - **views/**: Ansichten für die Benutzeroberfläche.
    - **router/**: Routen-Konfiguration.
    - **store/**: Zustandverwaltung mit Vuex.
  - **public/**: Öffentliche Dateien, einschließlich HTML.
  - **package.json**: NPM-Konfiguration.
  - **vue.config.js**: Vue.js-Konfiguration.

- **database/**: Datenbankdateien und -skripte.
  - **sqlite.db**: SQLite-Datenbank.
  - **postgres.sql, mssql.sql, mysql.sql, oracle.sql**: SQL-Skripte für andere Datenbanken.

- **docs/**: Dokumentation für Benutzer und Administratoren.
  - **user-documentation.md**: Benutzerdokumentation.
  - **admin-documentation.md**: Administrationsdokumentation.

- **Dockerfile**: Anweisungen zum Erstellen des Docker-Images.

- **docker-compose.yml**: Konfiguration für Docker-Container.

- **.env**: Umgebungsvariablen für die Anwendung.

## Installation

1. Klone das Repository:
   ```
   git clone <repository-url>
   cd go-webapp
   ```

2. Backend-Abhängigkeiten installieren:
   ```
   cd backend
   go mod tidy
   ```

3. Frontend-Abhängigkeiten installieren:
   ```
   cd frontend
   npm install
   ```

4. Umgebungsvariablen konfigurieren:
   - Erstelle eine `.env`-Datei basierend auf der `.env.example`.

5. Anwendung starten:
   - Backend:
     ```
     cd backend
     go run main.go
     ```
   - Frontend:
     ```
     cd frontend
     npm run serve
     ```

## Nutzung

- Die Anwendung ermöglicht es Benutzern, Projekte zu erstellen, zu verwalten und die Kosten zu schätzen.
- Benutzer können sich über Keycloak authentifizieren.
- Die Anwendung unterstützt die Migration von SQLite zu anderen Datenbanken.

## Lizenz

Dieses Projekt steht unter der MIT-Lizenz. Weitere Informationen finden Sie in der LICENSE-Datei.