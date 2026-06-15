-- +goose Up
CREATE TABLE IF NOT EXISTS api_keys(
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    key_id VARCHAR(16) NOT NULL UNIQUE,
    key_hash CHAR(64) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_used_at TIMESTAMP NULL,
    revoked_at TIMESTAMP NULL,


    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
)
;

-- +goose Down
DROP TABLE IF EXISTS api_keys;
