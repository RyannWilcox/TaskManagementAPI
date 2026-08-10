
INSERT INTO roles(id, name) VALUES
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', 'admin'),
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', 'user');

INSERT INTO user_roles(user_id, role_id) VALUES
('bd006d41-aded-4040-9934-2ba4e909ef9a', 'a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d');