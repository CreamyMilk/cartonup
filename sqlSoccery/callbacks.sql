DROP TABLE stkCalls;
CREATE TABLE stkCalls(
    ID int NOT NULL AUTO_INCREMENT,
    houseName VARCHAR(100),
    branch   VARCHAR(100),
    phoneNo VARCHAR (100),
    checkoutRequestID VARCHAR(200),
    isDone            int DEFAULT 0,
    createdAt         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(ID)
);