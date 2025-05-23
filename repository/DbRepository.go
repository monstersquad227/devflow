package repository

import (
	"context"
	"database/sql"
	"devflow/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	MysqlClient *sql.DB
	RedisClient *redis.Client
)

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		config.GlobalConfig.Mysql.Username,
		config.GlobalConfig.Mysql.Password,
		config.GlobalConfig.Mysql.Addr,
		config.GlobalConfig.Mysql.Port,
		config.GlobalConfig.Mysql.Databases,
		config.GlobalConfig.Mysql.Charset)
	var err error
	MysqlClient, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("打开 MySQL 连接失败: %v", err)
	}
	if err = MysqlClient.Ping(); err != nil {
		log.Fatalf("连接 MySQL 失败: %v", err)
	}
	log.Println("MySQL 连接成功")
}

func InitRedis() {
	Addr := fmt.Sprintf("%s:%s",
		config.GlobalConfig.Redis.Addr,
		config.GlobalConfig.Redis.Port)
	RedisClient = redis.NewClient(&redis.Options{
		Addr: Addr,
		DB:   config.GlobalConfig.Redis.Db,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("连接 Redis 失败: %v", err)
	}
	log.Println("Redis 连接成功")
}
