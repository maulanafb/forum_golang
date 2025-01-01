-- migrate:up
ALTER TABLE posts 
MODIFY COLUMN user_id BIGINT;
ALTER TABLE posts ADD CONSTRAINT fk_user_id_posts FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

-- migrate:down
ALTER TABLE posts DROP FOREIGN KEY fk_user_id_posts;
