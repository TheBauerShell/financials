-- MSSQL Datenbankstruktur für die Projektverwaltung

CREATE TABLE Projects (
    ID INT PRIMARY KEY IDENTITY(1,1),
    Name NVARCHAR(255) NOT NULL,
    Description NVARCHAR(MAX),
    Owner NVARCHAR(255) NOT NULL,
    Status NVARCHAR(50) NOT NULL,
    Department NVARCHAR(255) NOT NULL,
    CreatedAt DATETIME DEFAULT GETDATE(),
    UpdatedAt DATETIME DEFAULT GETDATE()
);

CREATE TABLE Users (
    ID INT PRIMARY KEY IDENTITY(1,1),
    Name NVARCHAR(255) NOT NULL,
    Email NVARCHAR(255) NOT NULL UNIQUE,
    Department NVARCHAR(255) NOT NULL,
    CreatedAt DATETIME DEFAULT GETDATE(),
    UpdatedAt DATETIME DEFAULT GETDATE()
);

CREATE TABLE ProjectComponents (
    ID INT PRIMARY KEY IDENTITY(1,1),
    ProjectID INT NOT NULL,
    ComponentType NVARCHAR(255) NOT NULL,
    Cost DECIMAL(18, 2) NOT NULL,
    FOREIGN KEY (ProjectID) REFERENCES Projects(ID)
);

CREATE TABLE ProjectStatus (
    ID INT PRIMARY KEY IDENTITY(1,1),
    ProjectID INT NOT NULL,
    Status NVARCHAR(50) NOT NULL,
    FOREIGN KEY (ProjectID) REFERENCES Projects(ID)
);