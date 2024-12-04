-- name: CreateAgent :one
INSERT INTO agents (
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
    website
)
VALUES (
    ?, -- id
    CURRENT_TIMESTAMP, -- created_at
    CURRENT_TIMESTAMP, -- updated_at
    ?, -- first_name
    ?, -- last_name
    ?, -- nick_name
    ?, -- person_name
    ?, -- title
    ?, -- slogan
    ?, -- email
    ?, -- agent_rating
    ?, -- description
    ?, -- recommendations_count
    ?, -- review_count
    ?, -- last_updated
    ?, -- first_month
    ?, -- first_year
    ?, -- photo
    ?, -- video
    ?, -- profile_url
    ?  -- website
)
RETURNING *;

-- name: GetAgent :one
SELECT * FROM agents WHERE id = ?;


-- name: UpdateAgentForeignKeys :exec
UPDATE agents
SET 
    address_id = ?,
    broker_id = ?,
    office_id = ?
WHERE id = ?;