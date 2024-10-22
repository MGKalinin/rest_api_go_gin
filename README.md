В данном пет-проекте реализован RESTful API сервер с двумя эндпоинтами   
======================================================================   
В качестве примера используются данные записей музыкальных произведений
----------------------------------------------
### Запустить сервер
```
go run .
```   

### Запросить перечень альбомов 
```
$ curl http://localhost:8080/albums
```

### Запросить определённый альбом   
```
curl http://localhost:8080/albums/2
```

### Добавить запись в перечень альбомов   
```
$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

 