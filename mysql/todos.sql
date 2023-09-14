CREATE TABLE todos (
    todo_id INT AUTO_INCREMENT PRIMARY KEY,
    todo_text VARCHAR(255) NOT NULL,
    check_status_id INT,
    display BOOLEAN NOT NULL DEFAULT TRUE,
    display_order INT NOT NULL,
    user_id INT,
    FOREIGN KEY (check_status_id) REFERENCES check_status(check_status_id) 
        ON UPDATE CASCADE 
        ON DELETE RESTRICT,
    FOREIGN KEY (user_id) REFERENCES users(user_id) 
        ON UPDATE CASCADE 
        ON DELETE CASCADE
) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

INSERT INTO todos (todo_text, check_status_id, display, display_order, user_id) VALUES
('食料品の買い物', 1, TRUE, 1, 1),
('プロジェクト報告書を完成させる', 2, TRUE, 2, 1),
('銀行に電話する', 3, FALSE, 3, 1);