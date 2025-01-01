-- migrate:up
CREATE TABLE IF NOT EXISTS posts (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  post_title varchar(255) NOT NULL,
  post_content TEXT NOT NULL,
  post_hashtag TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255) NOT NULL DEFAULT 'system',
  updated_by VARCHAR(255) NOT NULL DEFAULT 'system',
  deleted_at TIMESTAMP
);

-- migrate:down
DROP TABLE IF EXISTS posts;
