# **Sample Golang Webserver**

A sample webserver to enable quick tests.

## **Creating a docker build**

```bash
    docker build -t sample-golang-webserver:latest .
```

## **Running the docker**

```bash
    docker-compose up -d
```

## **Working with SQLite3**

1. Opening a database file in command line

```bash
    sqlite3 test.db
```

2. Quitting the terminal of SQLite3

```bash
    Ctrl+D
```

3. List all tables in SQLite3 DB

```sql
    SELECT name FROM sqlite_master
    WHERE type IN ('table','view') AND name NOT LIKE 'sqlite_%'
    UNION ALL
    SELECT name FROM sqlite_temp_master
    WHERE type IN ('table','view')
    ORDER BY 1;
```

4. List all columns of a table in SQLite3

```sql
    SELECT c.name FROM pragma_table_info('your_table_name') c;
```

## References

* [Gorilla MUX with SQLite3](https://rshipp.com/go-web-api/)