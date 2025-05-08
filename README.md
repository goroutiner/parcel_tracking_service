[![codecov](https://codecov.io/gh/goroutiner/parcel_tracking_service/graph/badge.svg)](https://codecov.io/gh/goroutiner/parcel_tracking_service)

## 📖 Translations
- [Read in Russian](/README_RU.md)

---

<h3 align="center">
  <div align="center">
    <h1>Parcel Tracking Service</h1>
  </div>
  <a href="https://github.com/goroutiner/parcel_tracking_service">
    <img src="https://static.tildacdn.com/tild3332-3638-4866-a234-336133636235/lgiui.png" width="600" height="400"/>
  </a>
</h3>

---

## 📋 Project Description

**Parcel Tracking Service** is a convenient web application for tracking parcels, designed to simplify the logistics process. With this service, customers can register parcels, track their status, update delivery addresses, and delete parcels if necessary.

---

## What is implemented in the application?

- ✔️ Parcel registration linked to a customer \
  ![Register](https://github.com/goroutiner/parcel_tracking_service/raw/dev/images/register_exemple.png)

- ✔️ Ability to change parcel status \
  ![Change Status](https://github.com/goroutiner/parcel_tracking_service/raw/dev/images/change-status_exemple.png)

- ✔️ Update delivery address \
  ![Update Address](https://github.com/goroutiner/parcel_tracking_service/raw/dev/images/update-address_exemple.png)

- ✔️ Delete parcels \
  ![Delete](https://github.com/goroutiner/parcel_tracking_service/raw/dev/images/delete_exemple.png)

- ✔️ Dynamic parcel table \
  ![Table](https://github.com/goroutiner/parcel_tracking_service/raw/dev/images/table_exemple.png)

- ✔️ Database integration 📇

---

## 📦 Instructions for Building and Running the Application with Docker

If you want to run the project using Docker, follow these steps:

1. Make sure Docker is installed.
2. To build and run the application, navigate to the project's root directory.
3. Build the application by running the following command in the terminal:

```
docker compose up
```

If you want the application to run in the background without logging, execute the command:
```
docker compose up -d
```

#### Now you can open the application in your browser at: [http://localhost:8080](http://localhost:8080/ "Specify the port you configured in Port")

---

## 🛠️ Technical Resources

- **Libraries for Database Interaction:** [jmoiron/sqlx](https://github.com/jmoiron/sqlx) and [jackc/pgx](https://github.com/jackc/pgx)

- **Library for Writing Tests:** [stretchr/testify](https://github.com/stretchr/testify)

---

# Conclusion

Thank you for using **Parcel Tracking Service**! 🚀 We hope the service helps you efficiently manage deliveries and logistics. 😊

---