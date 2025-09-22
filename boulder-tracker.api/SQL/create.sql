-- FIRST VERSION: 25.10.2024 1830
-- UPDATED: 30.07.2025 1440
-- UPDATED (User.Email to Principal): 22.09.2025 2030

DROP DATABASE IF EXISTS bouldertracker;

CREATE DATABASE bouldertracker CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

use bouldertracker;

create table users(
    Id char(36),
    UserName varchar(255) UNIQUE,
    Principal varchar(255) UNIQUE,
    IsDeleted boolean,
    PRIMARY KEY (Id)
);

create table sessions(
    Id char(36),
    StartTime datetime,
    EndTime datetime,
    BoulderedSolo boolean,
    IsDeleted boolean,
    UserId char(36),
    PRIMARY KEY (Id),
    FOREIGN KEY (UserId) REFERENCES users(Id) ON DELETE CASCADE
);

create table boulders(
    Id char(36),
    attempts int,
    sessionsTried int,
    exhausting boolean,
    SessionId char(36),
    PRIMARY KEY (Id),
    FOREIGN KEY (SessionId) REFERENCES sessions(Id) ON DELETE CASCADE
);

create table difficulties (
    Id char(36),
    alias varchar(255),
    relativeLevel int,
    PRIMARY KEY (Id)
);

create table styles(
    Id char(36),
    alias varchar(255),
    PRIMARY KEY (Id)
);

create table boulder_styles(
    boulderId char(36),
    styleId char(36),
    PRIMARY KEY(boulderId, styleId),
    FOREIGN KEY (boulderId) REFERENCES boulders(Id) ON DELETE CASCADE,
    FOREIGN KEY (styleId) REFERENCES styles(Id)
);

create table boulder_feltlike_difficulty(
    boulderId char(36),
    difficultyId char(36),
    PRIMARY KEY (boulderId, difficultyId),
    FOREIGN KEY (boulderId) REFERENCES boulders(Id) ON DELETE CASCADE,
    FOREIGN KEY (difficultyId) REFERENCES difficulties(Id)
);

create table boulder_actual_difficulty(
    boulderId char(36),
    difficultyId char(36),
    PRIMARY KEY (boulderId, difficultyId),
    FOREIGN Key (boulderId) REFERENCES boulders(Id) ON DELETE CASCADE,
    FOREIGN Key (difficultyId) REFERENCES difficulties(Id)
);
