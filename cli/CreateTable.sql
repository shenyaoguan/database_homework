create table buses
(
    location  varchar(191) not null
        primary key,
    price     bigint       not null,
    num_bus   bigint       not null,
    num_avail bigint       not null
);

create table customers
(
    cust_name longtext not null,
    cust_id   bigint auto_increment
        primary key
);

create table flights
(
    flight_num varchar(191) not null
        primary key,
    price      bigint       not null,
    num_seats  bigint       not null,
    num_avail  bigint       not null,
    from_city  longtext     not null,
    ariv_city  longtext     not null
);

create table hotels
(
    location  varchar(191) not null
        primary key,
    price     bigint       not null,
    num_rooms bigint       not null,
    num_avail bigint       not null
);

create table reservations
(
    cust_name longtext     not null,
    resv_type bigint       not null,
    resv_key  varchar(191) not null
        primary key
);

