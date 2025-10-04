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

create table difficulties (
    Id char(36),
    Alias varchar(255),
    RelativeLevel int,
    PRIMARY KEY (Id)
);

create table styles(
    Id char(36),
    Alias varchar(255),
    PRIMARY KEY (Id)
);

create table sessions(
    Id char(36),
    StartTime datetime,
    EndTime datetime,
    BoulderedSolo boolean,
    SessionState int,
    IsDeleted boolean,
    UserId char(36),
    PRIMARY KEY (Id),
    FOREIGN KEY (UserId) REFERENCES users(Id) ON DELETE CASCADE
);

create table boulders(
    Id char(36),
    Attempts int,
    SessionsTried int,
    Exhausting boolean,
    SessionId char(36),
    Likes boolean,
    ScrewedDifficultyId char(36),
    FeltLikeDifficultyId char(36),
    PRIMARY KEY (Id),
    FOREIGN KEY (SessionId) REFERENCES sessions(Id) ON DELETE CASCADE,
    FOREIGN KEY (ScrewedDifficultyId) REFERENCES difficulties(Id),
    FOREIGN KEY (FeltLikeDifficultyId) REFERENCES difficulties(Id)
);

create table boulder_styles(
    BoulderId char(36),
    StyleId char(36),
    PRIMARY KEY(BoulderId, StyleId),
    FOREIGN KEY (BoulderId) REFERENCES boulders(Id) ON DELETE CASCADE,
    FOREIGN KEY (StyleId) REFERENCES styles(Id)
);
