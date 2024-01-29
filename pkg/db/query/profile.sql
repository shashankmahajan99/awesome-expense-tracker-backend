
-- name: CreateProfile :execresult
INSERT INTO profiles (
  user_id,
  bio,
  profile_name,
  profile_picture
) VALUES (
  (SELECT id FROM Users WHERE email = ?), ?, ?, ?
);

-- name: GetProfileByEmail :one
SELECT sqlc.embed(profiles)
FROM profiles
JOIN Users ON profiles.user_id = Users.id
WHERE Users.email = ?;

-- name: GetProfileByID :one
SELECT *
FROM profiles
WHERE id = ?;

-- name: UpdateProfileBio :execresult
UPDATE profiles
JOIN Users ON profiles.user_id = Users.id
SET profiles.bio = ?
WHERE Users.email = ?;

-- name: UpdateProfileProfilePicture :execresult
UPDATE profiles
JOIN Users ON profiles.user_id = Users.id
SET profiles.profile_picture = ?
WHERE Users.email = ?;

-- name: UpdateProfileName :execresult
UPDATE profiles
JOIN Users ON profiles.user_id = Users.id
SET profiles.profile_name = ?
WHERE Users.email = ?;

-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE user_id = (SELECT id FROM Users WHERE email = ?);

-- name: Listprofiles :many
SELECT * 
FROM profiles
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: Countprofiles :one
SELECT count(*) 
FROM profiles;
