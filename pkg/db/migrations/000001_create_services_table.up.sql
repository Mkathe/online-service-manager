CREATE TABLE IF NOT EXISTS services (
    Id uuid PRIMARY KEY UNIQUE,
    Name VARCHAR(255) NOT NULL,
    Price INT NOT NULL,
    UserId uuid NOT NULL ,
    StartDate DATE NOT NULL,
    EndDate DATE
)