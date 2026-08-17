CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ NULL
);

INSERT INTO users(id, username, email, password) VALUES 
('bd006d41-aded-4040-9934-2ba4e909ef9a', 'admin', 'admin@gmail.com', '$2a$10$PZHhUlxVZk09ZsguiOyeeuIm31UZYSuZrWtq6CfVBg0/.Vm4BVm9m'),
('d8099de3-453b-49de-91bd-2dc498b852ff','kite', 'the_world@gmail.com','$2a$10$Hom1chCiVhVVFMD0fbcd8OtsdwK1IuiC2WH5E6p5jefZEs/7m4aqe'),
('447b4bb7-659f-423f-a6aa-098fdaee186e','orca', 'twilight@gmail.com',' $2a$10$6eE/WncwkNUN7ZH3H1fUM.U4gEudRKJUQeMrrClmmnxZCUhsl4bqu');