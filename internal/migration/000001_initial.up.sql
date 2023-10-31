CREATE TABLE Instrument(Id int PRIMARY KEY, Name varchar(255));
CREATE TABLE Trade(Id int PRIMARY KEY , InstrumentId int, DateEn timestamptz, Open decimal, High
                      decimal, Low decimal, Close decimal);

INSERT INTO Instrument values(1,'AAPL'), (2,'GOOGL');