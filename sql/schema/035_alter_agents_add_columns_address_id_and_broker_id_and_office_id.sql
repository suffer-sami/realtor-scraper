-- +goose Up
CREATE TABLE new_agents (
    id TEXT PRIMARY KEY,                    
    created_at TIMESTAMP NOT NULL,          
    updated_at TIMESTAMP NOT NULL,          
    profile_url TEXT,                              
    
    -- Personal Information Section
    first_name TEXT,
    last_name TEXT,
    nick_name TEXT,
    person_name TEXT,
    title TEXT,
    slogan TEXT,
    email TEXT,
    description TEXT,
    video TEXT,
    photo TEXT,
    website TEXT,
    
    -- Agent Performance Section
    agent_rating INTEGER,                   
    recommendations_count INTEGER,           
    review_count INTEGER,                    
    
    -- Date Information Section
    first_month INTEGER,                     
    first_year INTEGER,    
    last_updated TIMESTAMP,    
           

    -- Foreign Key References Section
    address_id INTEGER,                     
    broker_id INTEGER,                      
    office_id INTEGER,                      

    -- Foreign Key Constraints
    CONSTRAINT fk_addresses
        FOREIGN KEY (address_id) 
        REFERENCES addresses(id),
    CONSTRAINT fk_offices
        FOREIGN KEY (office_id) 
        REFERENCES offices(id),
    CONSTRAINT fk_brokers
        FOREIGN KEY (broker_id)
        REFERENCES brokers(id)
);

INSERT INTO new_agents (
    id,
    created_at,
    updated_at,
    first_name,
    last_name,
    nick_name,
    person_name,
    title,
    slogan,
    email,
    agent_rating,
    description,
    recommendations_count,
    review_count,
    last_updated,
    first_month,
    first_year,
    photo,
    video,
    profile_url,
    website,
    photo
)
SELECT 
    id,
    created_at,
    updated_at,
    first_name,
    last_name,
    nick_name,
    person_name,
    title,
    slogan,
    email,
    agent_rating,
    description,
    recommendations_count,
    review_count,
    last_updated,
    first_month,
    first_year,
    photo,
    video,
    profile_url,
    website,
    photo
FROM agents;

DROP TABLE agents;

ALTER TABLE new_agents RENAME TO agents;

-- +goose Down
CREATE TABLE new_agents (
    id TEXT PRIMARY KEY,                    
    created_at TIMESTAMP NOT NULL,          
    updated_at TIMESTAMP NOT NULL,          
    profile_url TEXT,                              
    
    -- Personal Information Section
    first_name TEXT,
    last_name TEXT,
    nick_name TEXT,
    person_name TEXT,
    title TEXT,
    slogan TEXT,
    email TEXT,
    description TEXT,
    video TEXT,
    photo TEXT,
    website TEXT,
    
    -- Agent Performance Section
    agent_rating INTEGER,                   
    recommendations_count INTEGER,           
    review_count INTEGER,                    
    
    -- Date Information Section
    first_month INTEGER,                     
    first_year INTEGER,    
    last_updated TIMESTAMP
);

INSERT INTO new_agents (
    id,
    created_at,
    updated_at,
    first_name,
    last_name,
    nick_name,
    person_name,
    title,
    slogan,
    email,
    agent_rating,
    description,
    recommendations_count,
    review_count,
    last_updated,
    first_month,
    first_year,
    photo,
    video,
    profile_url,
    website,
    photo
)
SELECT
    id,
    created_at,
    updated_at,
    first_name,
    last_name,
    nick_name,
    person_name,
    title,
    slogan,
    email,
    agent_rating,
    description,
    recommendations_count,
    review_count,
    last_updated,
    first_month,
    first_year,
    photo,
    video,
    profile_url,
    website,
    photo
FROM agents;

DROP TABLE agents;

ALTER TABLE new_agents RENAME TO agents;