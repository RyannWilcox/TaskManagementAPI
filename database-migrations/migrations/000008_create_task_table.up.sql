CREATE TABLE tasks (
  id UUID NOT NULL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL DEFAULT '',
  status VARCHAR(20) NOT NULL DEFAULT 'pending'
    CHECK (status IN ('pending', 'in_progress', 'completed')),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- CREATE SOME TASKS
INSERT INTO tasks (id, title, description, status, user_id) VALUES
('da1f6b13-f3bd-484f-9401-40adcc90dece', 'Set up CI pipeline', 'Add GitHub Actions workflow for build and test', 'pending', 'bd006d41-aded-4040-9934-2ba4e909ef9a'),
('7053d728-6460-40cf-be0d-570e3c20c021', 'Write API documentation', 'Document all v1 endpoints with example requests', 'in_progress', 'd8099de3-453b-49de-91bd-2dc498b852ff'),
('42329e42-982a-41ba-abe9-bd0671162e28', 'Fix task ownership bug', 'CreateTask should not trust client-supplied user_id', 'completed', '447b4bb7-659f-423f-a6aa-098fdaee186e');