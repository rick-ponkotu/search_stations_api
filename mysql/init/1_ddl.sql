CREATE TABLE IF NOT EXISTS app.table_station (
    stasion_cd  INTEGER NOT NULL,
    staion_g_cd INTEGER NOT NULL,
    station_name TEXT NOT NULL,
    station_name_k TEXT,
    station_name_r TEXT,
    line_cd INTEGER NOT NULL,
    pref_cd INTEGER,
    post TEXT,
    addr TEXT,
    lon FLOAT,
    lat FLOAT,
    open_ymd DATE,
    close_ymd DATE,
    e_status INTEGER,
    e_sort  INTEGER
) DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS app.table_line (
    line_cd INTEGER UNIQUE NOT NULL,
    company_cd INTEGER NOT NULL,
    line_name TEXT NOT NULL,
    line_name_k TEXT,
    line_name_r TEXT,
    line_color_c TEXT,
    line_color_t TEXT,
    line_type INTEGER,
    lon FLOAT,
    lat FLOAT,
    zoom INTEGER,
    e_status INTEGER,
    e_sort INTEGER
) DEFAULT CHARSET=utf8;