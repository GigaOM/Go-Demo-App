mysql --host=benchmarktest01.mysql.database.azure.com --user=mike -p
create database song_explorer_models;
use song_explorer_models;
create table song_explorer_models (Year VARCHAR(100), song_name VARCHAR(100), ID VARCHAR(100), created_at VARCHAR(100));