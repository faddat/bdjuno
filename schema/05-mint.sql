CREATE TABLE inflation
(
    value  DECIMAL NOT NULL,
    height BIGINT  NOT NULL,
    PRIMARY KEY (value, height)
);
CREATE INDEX inflation_height_index ON inflation (height);