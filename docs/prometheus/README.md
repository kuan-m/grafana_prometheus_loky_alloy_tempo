

## 4-12. Data Types
![4-12. Data Types](img/4-12-data-types.png)

## 4-13. Binary Arithmatic Operators

## 4-14. Binary Comparison Operators

## 4-15. Set Binary Operators

## 4-16. Matchers and Selectors
- **Selector**
```bash
http_requests_total
```

- **Selector with labels**
```bash
http_requests_total{job="api", method="GET"}
```

- **Empty Selector onyl labels**
```bash
{job="prometheus"}
```

- **Regex match .***
```bash
handler=~"/api/.*"
```

## 4-17. Aggregation Operators
- **Sum**
```bash
sum(node_cpu_seconds_total)
```

- **Sum by (key)** 
```bash
sum(node_cpu_seconds_total) by (mode)
```

- **Sum without (key)** 
```bash
sum(node_cpu_seconds_total) without (mode)
```

- **Group by (key)** 
```bash
group(node_cpu_seconds_total) by (mode)
```
```bash
avg(prometheus_http_requests_total) by (code)
```

## 4-18. Time Offsets
- **Offset**
```bash
avg(prometheus_http_requests_total offset 10m) by (code)
```

## 4-19. Clamping and Checking Functions
- **Absent** returns empty if exists, return value if not exists
```bash
absent(node_cpu_seconds_total)
```
```bash
absent(node_cpu_seconds_total{cpu="x09d"})
```

- **Absent over time** returns an empty vector if the range vector passed to it has any elements
```bash
absent_over_time(node_cpu_seconds_total[5m])
```
```bash
absent_over_time(node_cpu_seconds_total{cpu="x09d"}[5m])
```

- **Clamp** return value which lower limit of min and an upper limit of max
```bash
clamp(node_cpu_seconds_total, 300, 15000)
```

- **Clamp min** return value which are more than limit
```bash
clamp_min(node_cpu_seconds_total, 300)
```

- **Clamp max** return value which are less than limit
```bash
clamp_max(node_cpu_seconds_total, 15000)
```

## 4-20. Delta and IDelta

- **Delta** returns the difference between the first and last value of each time series element in range
- temperature between now and 2 hours ago
```bash
delta(cpu_temp_celsius{host="zeus"}[2h])
```

## 4-21. Sorting and Timestamps

- **Sort**  returns vector elements sorted by their float sample values (asc)
```bash
sort(clamp_max(node_cpu_seconds_total, 15000))
```

- **Sort by label**

- **Timestamp**  returns the timestamp since January 1, 1970 UTC
```bash
timestamp(clamp_max(node_cpu_seconds_total, 15000))
```