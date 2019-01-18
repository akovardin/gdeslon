# Где Слон?

Обертка на Go для API партнерсокй сети "Где Слон?". Страница с документацией по API: [https://www.gdeslon.ru/api_settings/xml](https://www.gdeslon.ru/api_settings/xml)

## Установка

```
go get github.com/horechek/gdeslon
```

## Методы

Только 1 основной метод для работы с API:

```go
Call(url, method string, params url.Values, result interface{}) error
```

## Пример

Создаем экземпляр клиента:

```go
client := gdeslon.NewClient(
    "https://api.gdeslon.ru/api",
    "xxx",
)
```

`xxx` - это ключ доступа

Получаем список баннеров через API

```go
type Response struct {
    Offers []struct {
        Id          uint64 `json:"id"`
        Name        string `json:"name"`
        Description string `json:"description"`
        Url         string `json:"url"`
    } `json:"offers"`
}

r := Response{}
params := url.Values{}
params.Add("q", "книги")
err := client.Call("search.json", "GET", params, &r)
if err != nil {
    log.Fatal(err)
}

for _, offer := range r.Offers {
    fmt.Println("==============")
    fmt.Println("id:", offer.Id)
    fmt.Println("name:", offer.Name)
    fmt.Println("description:", offer.Description)
    fmt.Println("url:", offer.Url)
}
```

Параметр `q` -- набор ключевых слов для поиска офферов

## Доступы

Доступы можно получить на странице с описанием API: [https://www.gdeslon.ru/api_settings/xml](https://www.gdeslon.ru/api_settings/xml)