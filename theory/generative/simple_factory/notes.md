# Simple Factory

**Simple Factory** - одна функция (или статический метод) создает нужный продукт по параметру: строка, enum, тип.

Не GoF-паттерн, но частый прием. Логика выбора - в одном месте (`switch`).

## Схема

```
Client
  |
  v
NewDocument("pdf")  --switch-->  PDFDocument / MarkdownDocument
```

## Когда уместно

- мало вариантов, набор стабилен
- выбор по конфигу / ключу в runtime (`"pdf"`, `"md"`)
- не нужно расширять через новые типы без правки фабрики

## Не путать

| | Simple Factory | Factory Method |
|---|----------------|----------------|
| Кто создает | одна функция | метод у каждого Creator |
| Расширение | правка switch | новый Creator |

См. также: [Factory Method](../factory_method/notes.md)

Код (Go): [dps_go/generative/simple_factory/v1/](../../../dps_go/generative/simple_factory/v1/)
