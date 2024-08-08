CREATE TABLE if not exists forum_posts (
    id uuid PRIMARY KEY not null DEFAULT GEN_RANDOM_UUID(),
    community_id UUID REFERENCES communities(id),
    user_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);