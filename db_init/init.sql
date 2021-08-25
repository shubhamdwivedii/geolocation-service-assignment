USE dockertest; 

CREATE TABLE geolocation (
    ip varchar(20),
    ccode varchar(3),
    country varchar(20),
    city varchar(20),
    latitude double, 
    longitude double, 
    mystery bigint, 
    PRIMARY KEY (ip)
);