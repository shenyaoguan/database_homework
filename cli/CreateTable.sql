-- 航班信息表
CREATE TABLE FLIGHTS
(
    flightNum VARCHAR(10) PRIMARY KEY,   -- 航班编号
    price     INT,                       -- 票价
    numSeats  INT,                       -- 座位总数
    numAvail  INT CHECK (numAvail >= 0), -- 剩余可用座位（不能小于0）
    FromCity  VARCHAR(50),               -- 出发城市
    ArivCity  VARCHAR(50)                -- 到达城市
);

-- 宾馆信息表
CREATE TABLE HOTELS
(
    location VARCHAR(50) PRIMARY KEY,  -- 宾馆位置（假设唯一）
    price    INT,                      -- 房间价格
    numRooms INT,                      -- 房间总数
    numAvail INT CHECK (numAvail >= 0) -- 剩余可用房间数（不能小于0）
);

-- 大巴信息表
CREATE TABLE BUS
(
    location VARCHAR(50) PRIMARY KEY,  -- 大巴车位置（假设唯一）
    price    INT,                      -- 票价
    numBus   INT,                      -- 大巴车总数
    numAvail INT CHECK (numAvail >= 0) -- 剩余可用大巴车数量（不能小于0）
);

-- 客户信息表
CREATE TABLE CUSTOMERS
(
    custName VARCHAR(50) PRIMARY KEY, -- 客户名称
    custID   INT UNIQUE               -- 客户ID，唯一
);

-- 预订信息表
CREATE TABLE RESERVATIONS
(
    resvKey  VARCHAR(10) PRIMARY KEY,           -- 预订主键
    custName VARCHAR(50),                       -- 客户名称
    resvType INT CHECK (resvType IN (1, 2, 3)), -- 预订类型（1: 航班, 2: 宾馆, 3: 大巴）
    resvRef  VARCHAR(50)                        -- 预订的具体航班、宾馆或大巴的标识
);

-- 强制执行 FLIGHTS 的 resvType 和 resvRef 约束的触发器
CREATE TRIGGER check_resv_flights
    BEFORE INSERT ON RESERVATIONS
    FOR EACH ROW
BEGIN
    IF NEW.resvType = 1 AND NOT EXISTS (SELECT 1 FROM FLIGHTS WHERE flightNum = NEW.resvRef) THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Invalid flightNum in resvRef for resvType 1';
    END IF;
END;

-- 强制执行 HOTELS 的 resvType 和 resvRef 约束的触发器
CREATE TRIGGER check_resv_hotels
    BEFORE INSERT ON RESERVATIONS
    FOR EACH ROW
BEGIN
    IF NEW.resvType = 2 AND NOT EXISTS (SELECT 1 FROM HOTELS WHERE location = NEW.resvRef) THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Invalid location in resvRef for resvType 2';
    END IF;
END;

-- 强制执行 BUS 的 resvType 和 resvRef 约束的触发器
CREATE TRIGGER check_resv_bus
    BEFORE INSERT ON RESERVATIONS
    FOR EACH ROW
BEGIN
    IF NEW.resvType = 3 AND NOT EXISTS (SELECT 1 FROM BUS WHERE location = NEW.resvRef) THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Invalid location in resvRef for resvType 3';
    END IF;
END;