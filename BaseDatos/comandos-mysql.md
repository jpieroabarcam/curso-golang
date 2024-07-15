# MySql
## Ejecutar
Desde la terminal de windows "CMD" ingresar desde la instacion de MySql y entrar a la carpeta bin

Ejecutar  para verificar si esta instalado

```bash
mysql --version
``` 

ingresar a mysql para entrar a usuario root
```bash
mysql -u root -p
``` 

Ingresara y permitira continuar creando un nuevo usuario
```bash
mysql> CREATE USER 'piero'@'localhost' IDENTIFIED BY 'admin'
mysql> GRANT ALL PRIVILEGES ON *.* TO 'piero'@'localhost' WITH GRANT OPTION
``` 

Salimos e ingresamos nuevamente a SQL con el usuario creado
```bash
mysql -u piero -p
``` 
Probamos si podemos ver todas las bases de datos
```bash
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sakila             |
| sys                |
| world              |
+--------------------+
7 rows in set (0.00 sec)
``` 

## Crear una base de datos

Usando el usuario creado y crear una base datos y tablas

```bash
mysql> create database if not exists db_contacts;
mysql> use db_contacts
Database changed
mysql> create table if not exists contact(id int auto_increment primary key,
    -> name varchar(255) not null,
    -> email varchar(255),
    -> phone varchar(20));
Query OK, 0 rows affected (0.09 sec)

mysql> show tables
    -> ;
+-----------------------+
| Tables_in_db_contacts |
+-----------------------+
| contact               |
+-----------------------+
1 row in set (0.01 sec)

``` 

Para realizar consulta a los registros registrando antes algunos

```bash
mysql> insert into contact(name, email, phone) values('Juan Perez', 'jperez@soaint.com', '967345183'),
    -> ('Piero Abarca', 'jabarca@soaint.com', '958279652'),
    -> ('Carlos Lopez', 'clopez@soaint.com', '963341554');
Query OK, 3 rows affected (0.02 sec)
Records: 3  Duplicates: 0  Warnings: 0

mysql> select *  from contact
    -> ;
+----+--------------+--------------------+-----------+
| id | name         | email              | phone     |
+----+--------------+--------------------+-----------+
|  1 | Juan Perez   | jperez@soaint.com  | 967345183 |
|  2 | Piero Abarca | jabarca@soaint.com | 958279652 |
|  3 | Carlos Lopez | clopez@soaint.com  | 963341554 |
+----+--------------+--------------------+-----------+
3 rows in set (0.00 sec)

``` 


