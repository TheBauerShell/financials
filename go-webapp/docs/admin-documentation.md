# Administrationsdokumentation für die Go-Webanwendung

## Einführung
Diese Dokumentation richtet sich an Administratoren, die für die Einrichtung und Verwaltung der Go-Webanwendung verantwortlich sind. Sie beschreibt die erforderlichen Schritte zur Konfiguration der Anwendung, zur Verwaltung von Benutzern und zur Durchführung von Datenmigrationen.

## Systemanforderungen
- Go 1.16 oder höher
- Node.js 14 oder höher
- Docker (optional, für Containerisierung)
- SQLite, PostgreSQL, MSSQL, MySQL oder Oracle (je nach Bedarf)

## Installation
1. **Repository klonen**
   ```
   git clone <repository-url>
   cd go-webapp
   ```

2. **Backend-Abhängigkeiten installieren**
   ```
   cd backend
   go mod tidy
   ```

3. **Frontend-Abhängigkeiten installieren**
   ```
   cd frontend
   npm install
   ```

4. **Umgebungsvariablen konfigurieren**
   - Erstellen Sie eine `.env`-Datei im Hauptverzeichnis und fügen Sie die erforderlichen Variablen hinzu, z.B. Datenbankverbindungsdetails und Keycloak-Konfiguration.

## Anwendung starten
- **Backend starten**
   ```
   cd backend
   go run main.go
   ```

- **Frontend starten**
   ```
   cd frontend
   npm run serve
   ```

## Keycloak-Integration
- Konfigurieren Sie die Keycloak-Details in der `backend/config/config.go`-Datei.
- Stellen Sie sicher, dass die Keycloak-Instanz läuft und die erforderlichen Clients und Benutzer eingerichtet sind.

## Datenbankverwaltung
### Standarddatenbank: SQLite
- Die Anwendung verwendet standardmäßig SQLite. Die Datenbankdatei befindet sich im `database`-Verzeichnis.

### Wechsel zu anderen Datenbanken
- Um zu PostgreSQL, MSSQL, MySQL oder Oracle zu wechseln, passen Sie die Verbindungsdetails in der `.env`-Datei an und führen Sie die entsprechenden SQL-Skripte im `database`-Verzeichnis aus.

### Datenmigration
- Die Migration von SQLite zu anderen Datenbanken wird in der `backend/migrations/migration.go`-Datei behandelt. Führen Sie die Migration mit den bereitgestellten Funktionen durch.

## Benutzerverwaltung
- Administratoren können Benutzer über die Keycloak-Oberfläche verwalten. Stellen Sie sicher, dass die Benutzer den entsprechenden Abteilungen zugeordnet sind.

## Projektverwaltung
- Administratoren können Projekte über die API verwalten. Die Endpunkte sind in der `backend/routes/routes.go`-Datei definiert.

## Dokumentation
- Für weitere Informationen zur Benutzerdokumentation, siehe `docs/user-documentation.md`.

## Support
- Bei Fragen oder Problemen wenden Sie sich bitte an den technischen Support oder konsultieren Sie die Community-Foren.

## Fazit
Diese Dokumentation bietet eine Übersicht über die administrativen Aufgaben zur Verwaltung der Go-Webanwendung. Stellen Sie sicher, dass alle Schritte sorgfältig befolgt werden, um eine reibungslose Funktion der Anwendung zu gewährleisten.