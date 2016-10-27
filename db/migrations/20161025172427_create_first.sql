-- +goose Up                                                                                                                                                  
CREATE TABLE entries (
    id int NOT NULL auto_increment,
    name varchar(128),
    entry text NOT NULL,
    is_show tinyint(1) NOT NULL DEFAULT '1',
    created_at timestamp NOT NULL default current_timestamp,
    CONSTRAINT pk_entries PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE entries;
