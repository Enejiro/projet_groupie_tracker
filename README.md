# Groupie Tracker

> Une application web pour explorer les artistes musicaux et leurs concerts à travers le monde

![Go](https://img.shields.io/badge/Go-1.25.0-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)
![Status](https://img.shields.io/badge/Status-In%20Development-blue)

---

## Table des matières

- [Objectif du projet](#objectif-du-projet)
- [Fonctionnalités](#fonctionnalités)
- [Installation et lancement](#installation-et-lancement)
- [Architecture](#architecture)
- [API utilisée](#api-utilisée)
- [Technologies](#technologies)
- [Améliorations futures](#améliorations-futures)
- [Licence](#licence)

---

## Objectif du projet

**Groupie Tracker** est une application web développée en Go permettant de visualiser et explorer les données d'artistes musicaux et leurs concerts à travers le monde.

### Objectifs pédagogiques atteints

- Manipulation de l'API HTTP avec le package `net/http`
- Récupération et parsing de données JSON depuis une API externe
- Organisation du projet Go en plusieurs packages
- Génération de pages web dynamiques avec `html/template`
- Gestion des erreurs côté serveur
- Conception d'une interface utilisateur claire et responsive

---

## Installation et lancement

### Prérequis

- Go **1.25.0+** installé sur votre machine
- Connexion internet (pour accéder à l'API Groupie Trackers)

### Démarrage rapide

```bash
# 1. Cloner le repository
git clone [URL_DU_REPOSITORY]
cd groupie-tracker

# 2. Lancer le serveur
go run .

# 3. Accéder à l'application
# Ouvrez votre navigateur et allez à :
# http://localhost:8080
```

> Le serveur sera en écoute sur le **port 8080**

---

## Fonctionnalités

### Fonctionnalités implémentées

#### Page d'accueil
- Présentation claire de l'application
- Navigation intuitive vers les artistes
- Affichage en grille de tous les artistes

#### Liste des artistes
- Affichage de tous les artistes sous forme de cartes
- Images des artistes
- Liens cliquables vers les pages détaillées

#### Page de détails d'un artiste
- Image en bannière
- Nom de l'artiste/groupe
- Liste complète des membres
- Année de création
- Date du premier album
- Navigation fluide

#### Événements interactifs
- Clic sur une carte d'artiste → requête vers `/artistes?id=X`
- Navigation entre les pages via la barre de navigation

#### Gestion d'erreurs
- Page d'erreur personnalisée (404, 400, 500)
- Gestion des erreurs côté serveur
- Affichage de messages d'erreur adaptés
- Liens de retour depuis les pages d'erreur

#### Interface utilisateur
- Design cohérent avec CSS
- Responsive design
- Navigation claire et intuitive


### À implémenter

- Système de filtres (année, nombre de membres, lieux)
- Affichage des dates et lieux de concert complets

---

## Architecture

```
groupie-tracker/
│
├── main.go              # Point d'entrée de l'application
├── server.go            # Configuration et démarrage du serveur HTTP
├── handler.go           # Gestionnaires de routes HTTP
├── API.go               # Interaction avec l'API Groupie Trackers
├── go.mod               # Dépendances Go
│
├── templates/           # Templates HTML
│   ├── index.html       # Page d'accueil
│   ├── artistes.html    # Page de détails d'un artiste
│   ├── map.html         # Page de la carte
│   ├── contact.html     # Page de contact
│   ├── error.html       # Pages d'erreur personnalisées
│   └── js/
│       └── sidebar.js
│
├── css/                 # Fichiers CSS
│   ├── global.css       # Styles globaux
│   ├── sidebar.css      # Navigation latérale
│   └── page.css         # Styles spécifiques
│
└── README.md            # Documentation du projet
```

### Branches Git

| Branche | Purpose |
|---------|---------|
| `main` | Branche principale de production |
| `dev` | Branche de développement |
| `feat/handlers` | Fonctionnalités des gestionnaires de routes |
| `feat/models` | Structures de données et modèles |
| `feat/templates` | Templates HTML |
| `server` | Configuration serveur |

---

## API utilisée

L'application consomme l'**API officielle Groupie Trackers** :

**URL de base:** `https://groupietrackers.herokuapp.com/api`

### Endpoints utilisés

| Endpoint | Description |
|----------|-------------|
| `GET /artists` | Récupération de la liste complète des artistes |
| `GET /artists/{id}` | Récupération des détails d'un artiste spécifique |
| `GET /relations/{id}` | Récupération des relations dates-lieux |

### Structure des données

<details>
<summary><b>Artist</b></summary>

```go
type Artist struct {
    ID           int                  // Identifiant unique
    Name         string               // Nom de l'artiste/groupe
    Image        string               // URL de l'image
    Members      []string             // Liste des membres
    Creation     int                  // Année de création
    FirstAlbum   string               // Date du premier album
    RelationsURL string               // URL vers les relations
    Relations    map[string][]string  // Dates par lieu
}
```

</details>

<details>
<summary><b>Relation</b></summary>

```go
type Relation struct {
    ID             int                  // Identifiant
    DatesLocations map[string][]string  // Map[lieu][]dates
}
```

</details>

### Fonctions API

```go
SearchArtist()
// Récupère tous les artistes
// Retour : []Artist, error
// Endpoint : GET /artists

getOneArtist(id int)
// Récupère un artiste spécifique avec ses relations
// Paramètres : id (identifiant de l'artiste)
// Retour : Artist, error
// Endpoints :
//   - GET /artists/{id}
//   - GET {RelationsURL} (pour les dates et lieux)
```

---

## Technologies

### Backend
- **Go 1.25.0**
  - `net/http` - Serveur HTTP et gestion des routes
  - `html/template` - Génération de pages dynamiques
  - `encoding/json` - Parsing des données JSON
  - `io` - Lecture des réponses HTTP

### Frontend
- **HTML5** - Structure des pages
- **CSS3** - Stylisation
  - `global.css` - Styles globaux
  - `sidebar.css` - Navigation
  - `page.css` - Styles spécifiques aux pages
- **JavaScript** - Interactions (sidebar)

### API externe
- **Groupie Trackers API** - Source de toutes les données

---

## Gestion des erreurs

L'application gère plusieurs types d'erreurs :

### Erreurs HTTP
| Code | Description |
|------|-------------|
| `404` | Page non trouvée |
| `400` | Requête invalide |
| `500` | Erreur serveur |

### Template d'erreur

Le fichier `error.html` affiche :
- Le code d'erreur (grand format)
- Un message explicatif adapté
- Un lien de retour à l'accueil
- Un lien vers la liste des artistes

---

## Améliorations futures

### Court terme
- [ ] Ajouter les filtres (année, membres, lieux)
- [ ] Compléter l'affichage des concerts (dates + lieux)

### Moyen terme
- [ ] Système de cache pour l'API
- [ ] Améliorer la gestion d'erreurs avec logging

### Long terme
- [ ] Thème sombre/clair
- [ ] Système de favoris (cookies)
- [ ] Comparaison d'artistes
- [ ] Optimisation des performances

---

## Notes techniques

### Configuration du serveur
- **Port:** `8080`
- **Fichiers statiques:** Servis depuis `/css/` et `/js/`
- **Templates:** Parsés dynamiquement pour chaque requête

### Points d'attention
- Les données proviennent exclusivement de l'API (pas de cache actuellement)
- Encodage **UTF-8** pour les caractères spéciaux
- Toute la logique est côté serveur (pas de JavaScript pour la logique métier)

---

## Équipe

Projet réalisé dans le cadre du cursus de développement web.

---

## Licence

Ce projet est réalisé dans un cadre pédagogique avec **Ynov Campus Toulouse**.

