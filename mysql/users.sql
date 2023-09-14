CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

INSERT INTO users (name, password, email) VALUES
('Taro Yamada', 'password123', 'taro.yamada@example.com'),
('Hanako Suzuki', 'hanakoPass', 'hanako.suzuki@example.com'),
('Jiro Tanaka', 'jiroSecret', 'jiro.tanaka@example.net');

DELETE FROM users WHERE name = 'Taro Yamada';
