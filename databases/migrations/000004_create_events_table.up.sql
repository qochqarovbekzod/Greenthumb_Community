CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY NOT NULL DEFAULT GEN_RANDOM_UUID(),
    community_id UUID REFERENCES communities(id),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    type event_type,
    start_time TIMESTAMP WITH TIME ZONE,
    end_time TIMESTAMP WITH TIME ZONE,
    location VARCHAR(255)
);