
-- name: CreateProfile :execresult
INSERT INTO Profiles (
  user_id,
  bio,
  profile_name,
  profile_picture
) VALUES (
  (SELECT id FROM Users WHERE email = ?), ?, ?, ?
);

-- name: GetProfileByEmail :one
SELECT sqlc.embed(Profiles)
FROM Profiles
JOIN Users ON Profiles.user_id = Users.id
WHERE Users.email = ?;

-- name: GetProfileByID :one
SELECT *
FROM Profiles
WHERE id = ?;

-- name: UpdateProfileBio :execresult
UPDATE Profiles
JOIN Users ON Profiles.user_id = Users.id
SET Profiles.bio = ?
WHERE Users.email = ?;

-- name: UpdateProfileProfilePicture :execresult
UPDATE Profiles
JOIN Users ON Profiles.user_id = Users.id
SET Profiles.profile_picture = ?
WHERE Users.email = ?;

-- name: UpdateProfileName :execresult
UPDATE Profiles
JOIN Users ON Profiles.user_id = Users.id
SET Profiles.profile_name = ?
WHERE Users.email = ?;

-- name: DeleteProfile :exec
DELETE FROM Profiles
WHERE user_id = (SELECT id FROM Users WHERE email = ?);

-- name: ListProfiles :many
SELECT * 
FROM Profiles
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CountProfiles :one
SELECT count(*) 
FROM Profiles;
