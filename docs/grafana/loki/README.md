### Connect Loki to Grafana

#### How it works?
![Design](./img/design.png)

#### 1. Add data source connection to loki
```bash
http://localhost:3030/connections/datasources
```
![Datasource](./img/datasource.png)


#### 2. Go to Explore Logs data
```bash
http://localhost:3030/explore
```
![Explore](./img/explore.png)

#### 3. Generate some logs
```bash
make gen-logs
```
