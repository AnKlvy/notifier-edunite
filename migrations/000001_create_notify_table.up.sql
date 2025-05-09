CREATE TABLE IF NOT EXISTS notifier_settings (
                                                 id SERIAL PRIMARY KEY,
                                                 user_id TEXT NOT NULL,
                                                 channel TEXT NOT NULL,
                                                 token TEXT NOT NULL,
                                                 UNIQUE (token)
    );

CREATE TABLE IF NOT EXISTS notifiсations (
                                             id SERIAL PRIMARY KEY,
                                             user_id TEXT NOT NULL,
                                             message TEXT NOT NULL,
                                             subject TEXT NOT NULL,
                                             metadata JSONB
);
