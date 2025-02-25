-- name: CreateFeedFollow :one

WITH inserted_data AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT inserted_data.*, users.name AS username, feeds.name AS feedname
FROM inserted_data
INNER JOIN users ON users.id = inserted_data.user_id
INNER JOIN feeds ON feeds.id = inserted_data.feed_id;

-- name: GetFeedFollowsForUser :many

SELECT feeds.name AS Feedname, users.name AS username
FROM feed_follows
INNER JOIN feeds ON feeds.id = feed_follows.feed_id
INNER JOIN users ON users.id = feeds.user_id
WHERE feed_follows.user_id = $1;