# Was soll  am Ende rauskommen
- Eine App, die beim Start ein Dashboard anzeigt, das einzelne Statistiken zeigt, wie ich in der letzten Zeit gebouldert bin
- Der Nutzer soll Sessions hinzufügen können, die anhand gewisser properties in den Statistiken dann ausgewertet werden
- Eventuell soll auch eingeblendet werden, wann welche Farbe in der Blochütte umgeschraubt wurde (ist aber noch nicht zu erwarten)

# Tech-Stack
### Frontend
- Angular
- Apex Charts für Diagramme
- Material als UI-Library

### Backend
- Go

### Datenspeicherung
- vorerst als CSV-Datei, irgendwann vielleicht Richtung SQL-Datenbank


# Schritte
1. Entity-Model erschaffen. Was soll wie in welcher Relation zu was abgespeichert werden?
2. (Backend) Datenmodell mit structs aufbauen. Erstellen und auslesen von Testdaten in Memory
3. (Backend, CSV) Testdaten in Memory in CSV abspeichern und von CSV auslesen
4. (Backend) API-Endpoints erstellen um über z.B. Postman Daten abzuspeichern und auszulesen (nur Create und Read)
5. (Frontend) Einfache Frontend Menüführung erstellen (seitliche Navbar (eingeklappt Icons sichtbar, ausgeklappt auch Text), für Handy unten; Dashboard, Hinzufügen, Sessions)
6. (Frontend) Einfaches Frontend, das es erlaubt, Daten abzuspeichern (Hinzufügen-Reiter; Wizard-Style?) und alle Daten auf einmal auszulesen (Session-Reiter; einfache Darstellung reicht)
7. (Frontend) Mit vorhandenen Daten ein Dashboard mit Apex Charts erstellen, in dem mehrere Diagramme drin sind
8. (Backend) API-Endpunkt und Business-Logic für die Bearbeitung von Daten erstellen
9. (Frontend) Bearbeiten einer einzelnen Session
10. (Frontend) schickere Darstellung der einzelnen Sessions mit eigenen Diagrammen 
