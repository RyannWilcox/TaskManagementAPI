
INSERT INTO permissions(id, resource, action) VALUES
('4b677009-73a2-4faa-abcb-c01d043246d7', 'profile', 'view'),
('8ef9a912-88c1-4880-a324-7162cc62dc00', 'profile', 'update'),
('ce460ef8-0605-412e-9544-b95870409654', 'profile', 'delete'),
('0a62f09c-4569-4c5e-ab60-cfde195c58d6', 'task', 'create'),
('3d13adf7-7e8b-4c17-80df-2001ca91c598', 'task', 'view'),
('4a904b7f-cfb4-4a65-930c-47e81de0895e', 'task', 'update'),
('39ca6fff-424d-487c-897c-6f920bb8ad5e', 'task', 'delete'),
('f65cd2d0-0275-4411-b5aa-828a2165d614', 'user', 'view'),
('d04b1601-b21a-4704-a88c-ee7d796b2f75', 'user', 'update'),
('6b51336d-e46a-4a8c-aab9-7e5d060e38a0', 'user', 'delete');

INSERT INTO roles(id, name) VALUES
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', 'admin'),
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', 'user');

INSERT INTO user_roles(user_id, role_id) VALUES
('bd006d41-aded-4040-9934-2ba4e909ef9a', 'a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d');

INSERT INTO role_permissions(role_id, permission_id) VALUES
-- admin permissions
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', '4b677009-73a2-4faa-abcb-c01d043246d7'),
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', '8ef9a912-88c1-4880-a324-7162cc62dc00'),
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', 'ce460ef8-0605-412e-9544-b95870409654'),
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', '0a62f09c-4569-4c5e-ab60-cfde195c58d6'),
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', '3d13adf7-7e8b-4c17-80df-2001ca91c598'),
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', '4a904b7f-cfb4-4a65-930c-47e81de0895e'),
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', '39ca6fff-424d-487c-897c-6f920bb8ad5e'),
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', 'f65cd2d0-0275-4411-b5aa-828a2165d614'),
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', 'd04b1601-b21a-4704-a88c-ee7d796b2f75'),
('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', '6b51336d-e46a-4a8c-aab9-7e5d060e38a0'),
-- user permissions
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', '4b677009-73a2-4faa-abcb-c01d043246d7'),
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', '8ef9a912-88c1-4880-a324-7162cc62dc00'),
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', '0a62f09c-4569-4c5e-ab60-cfde195c58d6'),
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', '3d13adf7-7e8b-4c17-80df-2001ca91c598'),
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', '4a904b7f-cfb4-4a65-930c-47e81de0895e'),
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', 'f65cd2d0-0275-4411-b5aa-828a2165d614'),
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', 'd04b1601-b21a-4704-a88c-ee7d796b2f75');
