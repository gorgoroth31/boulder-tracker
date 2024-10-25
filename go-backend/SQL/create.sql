use bouldertracker;

create table session(
    Id varchar(255),
    Date date,
    StartTime date,
    EndTime date,
    BoulderedSolo boolean,
    PRIMARY KEY (Id),
);

create table boulder(
    Id varchar(255),
    screwedDifficulty int,
    feltLikeDifficulty int,
    attempts int,
    sessionsTried int,
    exhausting boolean,
    style int,
    SessionId varchar(255),
    PRIMARY KEY (Id),
    FOREIGN KEY (SessionId) REFERENCES session(Id)
);

create table difficulty (
    Id int,
    alias varchar(255),
    PRIMARY KEY (Id)
);