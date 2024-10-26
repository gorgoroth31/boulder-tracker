-- FIRST VERSION: 25.10.2024 1830

CREATE DATABASE bouldertracker CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

use bouldertracker;

create table session(
    Id char(36),
    StartTime datetime,
    EndTime datetime,
    BoulderedSolo boolean,
    PRIMARY KEY (Id)
);

create table boulder(
    Id char(36),
    attempts int,
    sessionsTried int,
    exhausting boolean,
    SessionId char(36),
    PRIMARY KEY (Id),
    FOREIGN KEY (SessionId) REFERENCES session(Id)
);

create table difficulty (
    Id char(36),
    alias varchar(255),
    relativeLevel int,
    PRIMARY KEY (Id)
);

create table style(
    Id char(36),
    alias varchar(255),
    PRIMARY KEY (Id)
);

create table boulder_style(
    boulderId char(36),
    styleId char(36),
    PRIMARY KEY(boulderId, styleId),
    FOREIGN KEY (boulderId) REFERENCES boulder(Id),
    FOREIGN KEY (styleId) REFERENCES style(Id)
);

create table boulder_feltlike_difficulty(
    boulderId char(36),
    difficultyId char(36),
    PRIMARY KEY (boulderId, difficultyId),
    FOREIGN KEY (boulderId) REFERENCES boulder(Id),
    FOREIGN KEY (difficultyId) REFERENCES difficulty(Id)
);

create table boulder_actual_difficulty(
    boulderId char(36),
    difficultyId char(36),
    PRIMARY KEY (boulderId, difficultyId),
    FOREIGN Key (boulderId) REFERENCES boulder(Id),
    FOREIGN Key (difficultyId) REFERENCES difficulty(Id)
);
