-- 航班信息表
CREATE TABLE FLIGHTS (
                         id BIGINT AUTO_INCREMENT PRIMARY KEY,
                         flight_num VARCHAR(255) NOT NULL,
                         price INT NOT NULL,
                         num_seats INT NOT NULL,
                         num_avail INT NOT NULL,
                         from_city VARCHAR(255) NOT NULL,
                         ariv_city VARCHAR(255) NOT NULL,
                         created_at DATETIME,
                         updated_at DATETIME,
                         deleted_at DATETIME
);

-- 宾馆信息表
CREATE TABLE HOTELS (
                        id BIGINT AUTO_INCREMENT PRIMARY KEY,
                        location VARCHAR(255) NOT NULL,
                        price INT NOT NULL,
                        num_rooms INT NOT NULL,
                        num_avail INT NOT NULL,
                        created_at DATETIME,
                        updated_at DATETIME,
                        deleted_at DATETIME
);

-- 大巴信息表
CREATE TABLE BUS (
                     id BIGINT AUTO_INCREMENT PRIMARY KEY,
                     location VARCHAR(255) NOT NULL,
                     price INT NOT NULL,
                     num_bus INT NOT NULL,
                     num_avail INT NOT NULL,
                     created_at DATETIME,
                     updated_at DATETIME,
                     deleted_at DATETIME
);

-- 客户信息表
CREATE TABLE CUSTOMERS (
                           id BIGINT AUTO_INCREMENT PRIMARY KEY,
                           cust_name VARCHAR(255) NOT NULL,
                           cust_id INT UNIQUE,
                           created_at DATETIME,
                           updated_at DATETIME,
                           deleted_at DATETIME
);

-- 预订信息表
CREATE TABLE RESERVATIONS (
                              id BIGINT AUTO_INCREMENT PRIMARY KEY,
                              cust_name VARCHAR(255) NOT NULL,
                              resv_type INT NOT NULL,
                              resv_key VARCHAR(255) NOT NULL,
                              created_at DATETIME,
                              updated_at DATETIME,
                              deleted_at DATETIME
);