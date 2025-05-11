-- CREATE TABLE IF NOT EXISTS notifier_settings (
--      id SERIAL PRIMARY KEY,
--      user_id TEXT NOT NULL,
--      channel TEXT NOT NULL,
--      token TEXT NOT NULL
--     );

CREATE TABLE notifications (
       id SERIAL PRIMARY KEY,
       message TEXT NOT NULL,
       subject TEXT NOT NULL,
       metadata JSONB,
       images TEXT[]
    -- Можно добавить индекс для поиска по полям JSONB, если это необходимо
    -- CREATE INDEX idx_metadata ON notifications USING gin(metadata);
);

CREATE TABLE notifications_users (
     notification_id INTEGER REFERENCES notifications(id) ON DELETE CASCADE,
     user_id TEXT NOT NULL,
     PRIMARY KEY (notification_id, user_id)
);
