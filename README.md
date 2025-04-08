# 🚀 GoFiber API Starter

A simple REST API built with [GoFiber](https://gofiber.io/) and [GORM](https://gorm.io/) using PostgreSQL.  
This project demonstrates basic CRUD, database connection, validation, and structuring Go apps.

---

## 🛠 Features

- RESTful routing with GoFiber
- GORM ORM with PostgreSQL
- Auto migration
- Auto restart when save file in development using Air
- Environment variables using `godotenv`
- Input validation with `go-playground/validator`
- Clean code structure (modular)
- Ready for deployment

---

## 🚀 Getting Started

### 📋 Prerequisites

- Go 1.20+ installed
- PostgreSQL or any DB you configure
- Git

---

### 📥 Clone the project

```bash
git clone https://github.com/yasminadelia/gofiber-app.git
cd gofiber-app
```
---

### 📦 Install dependencies

```bash
go mod tidy
```
---

### ⚙️ Setup environment variables

```bash
cp .env.example .env
```
---

### 🗃️ Run the app

```bash
go run main.go
```
or
```bash
./dev.sh 
```
App will run by default on: `http://localhost:3000`

---

## 👩‍💻 Author
Made with 💙 by Yasmin Adelia

