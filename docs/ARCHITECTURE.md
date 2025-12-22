# Архитектура мониторинга: Prometheus vs Loki

## Поток данных для логов (Loki)

```
Go приложение → logs/loki_udemy.log → Promtail (читает файл) → Loki (хранит) → Grafana (визуализирует)
```

### Детали потока:

1. **log-generator** (Go приложение)
   - Запускается через `make gen-logs` или `docker compose run --rm log-generator`
   - Пишет логи в файл `./logs/loki_udemy.log`
   - Формат: `timestamp level=LEVEL app=myapp component=COMPONENT message`

2. **Promtail** (агент сбора логов)
   - **Читает** файлы логов из `./logs/` (монтируется как `/var/log` в контейнере)
   - **Scrape** (скрапит) - периодически читает новые строки из файлов
   - **Push** (пушит) - отправляет логи в Loki через HTTP API
   - Отслеживает позицию чтения в файле (positions.yaml)

3. **Loki** (хранилище логов)
   - Принимает логи через Push API (`/loki/api/v1/push`)
   - Хранит логи индексированно (по labels)
   - Предоставляет Query API для Grafana

4. **Grafana**
   - Подключается к Loki как datasource
   - Позволяет делать запросы через LogQL (язык запросов Loki)

## Prometheus vs Loki - ключевые различия

### Prometheus (метрики)
- **Тип данных**: Метрики (числа, счетчики, гистограммы)
- **Модель**: **Pull (Scrape)** - Prometheus сам запрашивает данные
- **Источники**: 
  - Exporters (node-exporter, application exporters)
  - Приложения с `/metrics` endpoint
- **Формат**: Prometheus format (текстовый формат метрик)
- **Хранение**: TSDB (Time Series Database)
- **Использование**: CPU, память, запросы в секунду, латентность

### Loki (логи)
- **Тип данных**: Логи (текстовые строки)
- **Модель**: **Push** - агенты (Promtail) отправляют логи в Loki
- **Источники**:
  - Файлы логов (через Promtail)
  - Docker контейнеры
  - Syslog
  - Приложения напрямую (через API)
- **Формат**: Любой текстовый формат логов
- **Хранение**: Индексированные чанки (по labels, не по содержимому)
- **Использование**: Ошибки, события, трассировка запросов

## Нужен ли Prometheus для Loki?

**НЕТ!** Prometheus и Loki - это **независимые системы**:

- ✅ Loki работает **без** Prometheus
- ✅ Prometheus работает **без** Loki
- ✅ Они **дополняют** друг друга в полном стеке мониторинга

### Когда использовать вместе:

```
Приложение
├── Метрики → Prometheus (CPU, память, RPS)
└── Логи → Loki (ошибки, события, трассировка)
         ↓
      Grafana (объединяет метрики и логи в одном дашборде)
```

## Базовый минимум для понимания

### 1. Prometheus Stack
```
Application/Exporter → Prometheus (scrape) → Grafana
```
- **Scrape interval**: как часто Prometheus запрашивает метрики
- **Targets**: список endpoints для сбора метрик
- **PromQL**: язык запросов для метрик

### 2. Loki Stack
```
Log files → Promtail (scrape файлы) → Loki (push) → Grafana
```
- **Scrape configs**: какие файлы читать
- **Labels**: метки для индексации (job, component, level)
- **LogQL**: язык запросов для логов

### 3. Важные концепции

**Labels (метки)**
- Используются для индексации и фильтрации
- В Loki: `job=loki_udemy`, `component=database`, `level=ERROR`
- В Prometheus: `job=node-exporter`, `instance=localhost:9100`

**Scrape vs Push**
- **Scrape (Pull)**: сервер сам запрашивает данные
  - Prometheus → exporters
  - Promtail → файлы логов
- **Push**: клиент сам отправляет данные
  - Promtail → Loki
  - Application → Loki (напрямую)

**Retention (хранение)**
- Prometheus: настраивается в конфиге (сколько хранить метрики)
- Loki: настраивается в конфиге (сколько хранить логи)

## Текущая архитектура проекта

```
┌─────────────────┐
│  log-generator  │ (Go приложение)
│  (контейнер)    │
└────────┬────────┘
         │ пишет
         ↓
┌─────────────────┐
│  logs/*.log     │ (файлы на хосте)
└────────┬────────┘
         │ читает (scrape)
         ↓
┌─────────────────┐
│    Promtail     │ (агент сбора)
└────────┬────────┘
         │ отправляет (push)
         ↓
┌─────────────────┐
│      Loki       │ (хранилище)
└────────┬────────┘
         │ запросы
         ↓
┌─────────────────┐
│    Grafana      │ (визуализация)
└─────────────────┘

┌─────────────────┐
│ node-exporter   │ (метрики системы)
└────────┬────────┘
         │ scrape
         ↓
┌─────────────────┐
│   Prometheus    │ (метрики)
└────────┬────────┘
         │ запросы
         ↓
┌─────────────────┐
│    Grafana      │ (визуализация)
└─────────────────┘
```
 