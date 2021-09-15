create table stock_info (
    id INTEGER PRIMARY KEY,
    mtime datetime default CURRENT_TIMESTAMP,
    stock_code varchar(64) NOT NULL,
    stock_name varchar(64) NOT NULL,
    market varchar(16) NOT NULL
);
create table stock_price (
    stock_id varchar(8) NOT NULL,
    cdate date default CURRENT_DATE,
    price varchar(8) NOT NULL,
    PRIMARY KEY (stock_id, cdate)
);

create table holder_info (
    id INTEGER PRIMARY KEY,
    mtime datetime default CURRENT_TIMESTAMP,
    ctime datetime default CURRENT_TIMESTAMP,
    h_name varchar(64) NOT NULL,
    h_percentage varchar(8) NOT NULL,
    h_total INTEGER default 0,
    h_portfolio INTEGER NOT NULL,
);

create unique index holder_port_idx on holder_info (h_name, h_portfolio);
create index holder_name on holder_info (h_name);
create table holder_trans (

);

create table portfolio_info (
    id INTEGER PRIMARY KEY,
    market_value INTEGER default 0,
    cash INTEGER default 0,
    total_value INTEGER default 0
);

create table portfolio_stock (
    s_code varchar(64) NOT NULL,
    s_num INTEGER NOT NULL,
    mtime datetime default CURRENT_TIMESTAMP,
    p_id INTEGER NOT NULL
);

create unique index stock_port_idx on portfolio_stock (s_code, p_id);