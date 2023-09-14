CREATE TABLE check_status (
    check_status_id INT AUTO_INCREMENT PRIMARY KEY,
    check_status VARCHAR(255) NOT NULL
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

INSERT INTO check_status (check_status) VALUES
('まだ開始していない'),
('進行中'),
('完了');