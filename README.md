# FILE_SYS

---

## psql

sudo -i -u postgres
psql

```
CREATE DATABASE filemanagement;
CREATE USER fileuser WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE filemanagement TO fileuser;
\q
```

psql -U fileuser -d filemanagement -W

```
CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL,
    md5 VARCHAR(32) NOT NULL
);
\q
```

```
go get github.com/lib/pq
```

