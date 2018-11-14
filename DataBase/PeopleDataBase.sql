Create DATABASE People
GO
USE People
GO
CREATE TABLE Person(
	IdPerson INT PRIMARY KEY IDENTITY,
	[Name] VARCHAR(20),
	[LastName] VARCHAR(20)
)
GO
UPDATE Person
SET [Name] = 'Panfilo', [LastName] = 'Gomez'
WHERE IdPerson = 1
GO
INSERT INTO Person ([Name],[LastName]) VALUES ('Prueba','Prueba')
SELECT * FROM Person