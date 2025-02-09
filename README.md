<h3 align="center">
  <div align="center">
    <h1>Parcel Tracking Service </h1>
  </div>
  <a href="https://github.com/goroutiner/parcel_tracking_service">
    <img src="https://static.tildacdn.com/tild3332-3638-4866-a234-336133636235/lgiui.png" width="600" height="400"/>
  </a>
</h3>

---

## 📋 Описание проекта

**Parcel Tracking Service** — это удобное веб-приложение для отслеживания посылок, разработанное для упрощения процесса логистики. С помощью этого сервиса клиенты могут регистрировать посылки, отслеживать их статус, обновлять адрес доставки и удалять посылки при необходимости. \

---

## Что реализовано в приложении?

- ✔️ Регистрация посылки с привязкой к клиенту \
  ![Register](https://github.com/goroutiner/parcel_tracking_service/raw/dev/images/register_exemple.png)

- ✔️ Возможность изменения статуса посылки \
  ![Change Status](https://github.com/goroutiner/parcel_tracking_service/raw/dev/images/change-status_exemple.png)

- ✔️ Обновление адреса доставки \
  ![Update Address](https://github.com/goroutiner/parcel_tracking_service/raw/dev/images/update-address_exemple.png)

- ✔️ Удаление посылок \
  ![Delete ](https://github.com/goroutiner/parcel_tracking_service/raw/dev/images/delete_exemple.png)

- ✔️ Динамичная таблица с посылками \
  ![Table](https://github.com/goroutiner/parcel_tracking_service/raw/dev/images/table_exemple.png)

- ✔️ Интеграция с базой данных 📇

---

## 📦 Инструкция по сборке и запуску приложения через Docker

Если вы хотите запустить проект через Docker, следуйте этим шагам:

1. Убедитесь, что у вас установлен Docker.
2. Для сборки и запуска приложения необходимо находиться в корне проекта.
3. Соберем приложение, выполнив эту команду в терминале:

```
docker compose up 
```
Если нужно, чтобы приложение работало фоном и без логирования, то выполните команду:
```
docker compose up -d
```

#### Теперь вы можете открыть приложение в браузере по адресу: [http://localhost:8080](http://localhost:8080/ "Порт указываете тот, который укзан в Port")

---

## 🛠️ Технические ресурсы

- **Библиотеки для взаимодействия с БД:** [jmoiron/sqlx](https://github.com/jmoiron/sqlx) и [ackc/pgx](https://github.com/jackc/pgx)
- **Библиотека для написания тестов:** [stretchr/testify](https://github.com/stretchr/testify)

---

# Заключение

Спасибо за использование **Parcel Tracking Service**! 🚀 Надеемся, что сервис поможет вам эффективно управлять доставками и логистикой. 😊