# NewsWebApp 

**NewsWebApp** - Это примитивный сервис для добавления, редактирования статей, также добавлена возможность получать статьи 
с возмозжностью фильтрации по категориям.
Для филтрации запрос должен содержать:
```
http://localhost:8080/news?categories=journey&categories=family
```
## Технологии:
- **Язык**: Go
- **Фреймворк**: Fiber
- **Контейнеризация**: Docker
- **База данных**: PostgreSql

Команда для деплоя в докер:
```
docker compose -f deployed/docker-compose.yml --env-file .env up -d
```
