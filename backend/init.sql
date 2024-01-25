CREATE TABLE users (
    user_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

INSERT INTO users (user_id, name, email) VALUES
('1', '一郎', 'ichiro@example.com'),
('2', '二郎', 'jiro@example.com'),
('3', '三郎', 'saburo@example.com');

CREATE TABLE tasks (
    task_id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) REFERENCES users(user_id),
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL
);

INSERT INTO tasks (task_id, user_id, title, description, status) VALUES
('101', '1', 'タスク1', 'タスク1の本文', 'NOT_STARTED'),
('102', '2', 'タスク2', 'タスク1の本文', 'IN_PROGRESS'),
('103', '3', 'タスク3', 'タスク1の本文', 'DONE');

