
-- name: CreateProfile :execresult
INSERT INTO profiles (
  user_id,
  bio,
  profile_name,
  profile_picture
) VALUES (
  (SELECT id FROM users WHERE email = ?), ?, ?, ?
);

-- name: GetProfileByEmail :one
SELECT sqlc.embed(profiles)
FROM profiles
JOIN users ON profiles.user_id = users.id
WHERE users.email = ?;

-- name: GetProfileByID :one
SELECT *
FROM profiles
WHERE id = ?;

-- name: UpdateProfileBio :execresult
UPDATE profiles
JOIN users ON profiles.user_id = users.id
SET profiles.bio = ?
WHERE users.email = ?;

-- name: UpdateProfileProfilePicture :execresult
UPDATE profiles
JOIN users ON profiles.user_id = users.id
SET profiles.profile_picture = ?
WHERE users.email = ?;

-- name: UpdateProfileName :execresult
UPDATE profiles
JOIN users ON profiles.user_id = users.id
SET profiles.profile_name = ?
WHERE users.email = ?;

-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE user_id = (SELECT id FROM users WHERE email = ?);

-- name: Listprofiles :many
SELECT * 
FROM profiles
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: Countprofiles :one
SELECT count(*) 
FROM profiles;
