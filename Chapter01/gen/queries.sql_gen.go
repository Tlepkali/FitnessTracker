// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: queries.sql

package chapter1

import (
	"context"
	"time"
)

const createExercise = `-- name: CreateExercise :one
INSERT INTO gowebapp.exercises (Exercise_Name)
values ($1) RETURNING Exercise_ID
`

// insert a new exercise
func (q *Queries) CreateExercise(ctx context.Context, exerciseName string) (int64, error) {
	row := q.db.QueryRowContext(ctx, createExercise, exerciseName)
	var exercise_id int64
	err := row.Scan(&exercise_id)
	return exercise_id, err
}

const createSet = `-- name: CreateSet :one
INSERT INTO gowebapp.sets (Exercise_Id, Weight)
values ($1,
        $2) RETURNING set_id, exercise_id, weight
`

type CreateSetParams struct {
	ExerciseID int64 `db:"exercise_id" json:"exerciseID"`
	Weight     int32 `db:"weight" json:"weight"`
}

// insert new exercise sets
func (q *Queries) CreateSet(ctx context.Context, arg CreateSetParams) (GowebappSet, error) {
	row := q.db.QueryRowContext(ctx, createSet, arg.ExerciseID, arg.Weight)
	var i GowebappSet
	err := row.Scan(&i.SetID, &i.ExerciseID, &i.Weight)
	return i, err
}

const createUserImage = `-- name: CreateUserImage :one
INSERT INTO gowebapp.images (User_ID, Content_Type, Image_Data)
values ($1,
        $2,
        $3) RETURNING image_id, user_id, content_type, image_data
`

type CreateUserImageParams struct {
	UserID      int64  `db:"user_id" json:"userID"`
	ContentType string `db:"content_type" json:"contentType"`
	ImageData   []byte `db:"image_data" json:"imageData"`
}

// insert a new image
func (q *Queries) CreateUserImage(ctx context.Context, arg CreateUserImageParams) (GowebappImage, error) {
	row := q.db.QueryRowContext(ctx, createUserImage, arg.UserID, arg.ContentType, arg.ImageData)
	var i GowebappImage
	err := row.Scan(
		&i.ImageID,
		&i.UserID,
		&i.ContentType,
		&i.ImageData,
	)
	return i, err
}

const createUsers = `-- name: CreateUsers :one
INSERT INTO gowebapp.users (User_Name, Pass_Word_Hash, name)
VALUES ($1,
        $2,
        $3) RETURNING user_id, user_name, pass_word_hash, name, config, created_at, is_enabled
`

type CreateUsersParams struct {
	UserName     string `db:"user_name" json:"userName"`
	PassWordHash string `db:"pass_word_hash" json:"passWordHash"`
	Name         string `db:"name" json:"name"`
}

// insert new user
func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (GowebappUser, error) {
	row := q.db.QueryRowContext(ctx, createUsers, arg.UserName, arg.PassWordHash, arg.Name)
	var i GowebappUser
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.PassWordHash,
		&i.Name,
		&i.Config,
		&i.CreatedAt,
		&i.IsEnabled,
	)
	return i, err
}

const createWorkout = `-- name: CreateWorkout :one
INSERT INTO gowebapp.workouts (User_ID, Set_ID, Start_Date)
values ($1,
        $2,
        $3) RETURNING workout_id, set_id, user_id, exercise_id, start_date
`

type CreateWorkoutParams struct {
	UserID    int64     `db:"user_id" json:"userID"`
	SetID     int64     `db:"set_id" json:"setID"`
	StartDate time.Time `db:"start_date" json:"startDate"`
}

// insert new workouts
func (q *Queries) CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (GowebappWorkout, error) {
	row := q.db.QueryRowContext(ctx, createWorkout, arg.UserID, arg.SetID, arg.StartDate)
	var i GowebappWorkout
	err := row.Scan(
		&i.WorkoutID,
		&i.SetID,
		&i.UserID,
		&i.ExerciseID,
		&i.StartDate,
	)
	return i, err
}

const deleteExercise = `-- name: DeleteExercise :exec
DELETE
FROM gowebapp.exercises e
WHERE e.exercise_id = $1
`

// delete a particular exercise
func (q *Queries) DeleteExercise(ctx context.Context, exerciseID int64) error {
	_, err := q.db.ExecContext(ctx, deleteExercise, exerciseID)
	return err
}

const deleteSets = `-- name: DeleteSets :exec
DELETE
FROM gowebapp.sets s
WHERE s.set_id = $1
`

// delete a particular exercise sets
func (q *Queries) DeleteSets(ctx context.Context, setID int64) error {
	_, err := q.db.ExecContext(ctx, deleteSets, setID)
	return err
}

const deleteUserImage = `-- name: DeleteUserImage :exec
DELETE
FROM gowebapp.images i
WHERE i.user_id = $1
`

// delete a particular user's image
func (q *Queries) DeleteUserImage(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUserImage, userID)
	return err
}

const deleteUserWorkouts = `-- name: DeleteUserWorkouts :exec
DELETE
FROM gowebapp.workouts w
WHERE w.user_id = $1
`

// delete a particular user's workouts
func (q *Queries) DeleteUserWorkouts(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUserWorkouts, userID)
	return err
}

const deleteUsers = `-- name: DeleteUsers :exec
DELETE
FROM gowebapp.users
WHERE user_id = $1
`

// delete a particular user
func (q *Queries) DeleteUsers(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUsers, userID)
	return err
}

const getUser = `-- name: GetUser :one
SELECT user_id, user_name, pass_word_hash, name, config, created_at, is_enabled
FROM gowebapp.users
WHERE user_id = $1
`

// get users of a particular user_id
func (q *Queries) GetUser(ctx context.Context, userID int64) (GowebappUser, error) {
	row := q.db.QueryRowContext(ctx, getUser, userID)
	var i GowebappUser
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.PassWordHash,
		&i.Name,
		&i.Config,
		&i.CreatedAt,
		&i.IsEnabled,
	)
	return i, err
}

const getUserImage = `-- name: GetUserImage :one
SELECT u.name, u.user_id, i.image_data
FROM gowebapp.users u,
     gowebapp.images i
WHERE u.user_id = i.user_id
  AND u.user_id = $1
`

type GetUserImageRow struct {
	Name      string `db:"name" json:"name"`
	UserID    int64  `db:"user_id" json:"userID"`
	ImageData []byte `db:"image_data" json:"imageData"`
}

// get a particular user image
func (q *Queries) GetUserImage(ctx context.Context, userID int64) (GetUserImageRow, error) {
	row := q.db.QueryRowContext(ctx, getUserImage, userID)
	var i GetUserImageRow
	err := row.Scan(&i.Name, &i.UserID, &i.ImageData)
	return i, err
}

const getUserSets = `-- name: GetUserSets :many
SELECT u.user_id, w.workout_id, w.start_date, s.set_id, s.weight
FROM gowebapp.users u,
     gowebapp.workouts w,
     gowebapp.sets s
WHERE u.user_id = w.user_id
  AND w.set_id = s.set_id
  AND u.user_id = $1
`

type GetUserSetsRow struct {
	UserID    int64     `db:"user_id" json:"userID"`
	WorkoutID int64     `db:"workout_id" json:"workoutID"`
	StartDate time.Time `db:"start_date" json:"startDate"`
	SetID     int64     `db:"set_id" json:"setID"`
	Weight    int32     `db:"weight" json:"weight"`
}

// get a particular user information, exercise sets and workouts
func (q *Queries) GetUserSets(ctx context.Context, userID int64) ([]GetUserSetsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserSets, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserSetsRow
	for rows.Next() {
		var i GetUserSetsRow
		if err := rows.Scan(
			&i.UserID,
			&i.WorkoutID,
			&i.StartDate,
			&i.SetID,
			&i.Weight,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserWorkout = `-- name: GetUserWorkout :many
SELECT u.user_id, w.workout_id, w.start_date, w.set_id
FROM gowebapp.users u,
     gowebapp.workouts w
WHERE u.user_id = w.user_id
  AND u.user_id = $1
`

type GetUserWorkoutRow struct {
	UserID    int64     `db:"user_id" json:"userID"`
	WorkoutID int64     `db:"workout_id" json:"workoutID"`
	StartDate time.Time `db:"start_date" json:"startDate"`
	SetID     int64     `db:"set_id" json:"setID"`
}

// get a particular user information and workouts
func (q *Queries) GetUserWorkout(ctx context.Context, userID int64) ([]GetUserWorkoutRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserWorkout, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserWorkoutRow
	for rows.Next() {
		var i GetUserWorkoutRow
		if err := rows.Scan(
			&i.UserID,
			&i.WorkoutID,
			&i.StartDate,
			&i.SetID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listExercises = `-- name: ListExercises :many
SELECT exercise_id, exercise_name
FROM gowebapp.exercises
ORDER BY exercise_name
`

// get all exercises ordered by the exercise name
func (q *Queries) ListExercises(ctx context.Context) ([]GowebappExercise, error) {
	rows, err := q.db.QueryContext(ctx, listExercises)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GowebappExercise
	for rows.Next() {
		var i GowebappExercise
		if err := rows.Scan(&i.ExerciseID, &i.ExerciseName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listImages = `-- name: ListImages :many
SELECT image_id, user_id, content_type, image_data
FROM gowebapp.images
ORDER BY image_id
`

// get all images ordered by the id
func (q *Queries) ListImages(ctx context.Context) ([]GowebappImage, error) {
	rows, err := q.db.QueryContext(ctx, listImages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GowebappImage
	for rows.Next() {
		var i GowebappImage
		if err := rows.Scan(
			&i.ImageID,
			&i.UserID,
			&i.ContentType,
			&i.ImageData,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSets = `-- name: ListSets :many
SELECT set_id, exercise_id, weight
FROM gowebapp.sets
ORDER BY weight
`

// get all exercise sets ordered by weight
func (q *Queries) ListSets(ctx context.Context) ([]GowebappSet, error) {
	rows, err := q.db.QueryContext(ctx, listSets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GowebappSet
	for rows.Next() {
		var i GowebappSet
		if err := rows.Scan(&i.SetID, &i.ExerciseID, &i.Weight); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT user_id, user_name, pass_word_hash, name, config, created_at, is_enabled
FROM gowebapp.users
ORDER BY user_name
`

// get all users ordered by the username
func (q *Queries) ListUsers(ctx context.Context) ([]GowebappUser, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GowebappUser
	for rows.Next() {
		var i GowebappUser
		if err := rows.Scan(
			&i.UserID,
			&i.UserName,
			&i.PassWordHash,
			&i.Name,
			&i.Config,
			&i.CreatedAt,
			&i.IsEnabled,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listWorkouts = `-- name: ListWorkouts :many
SELECT workout_id, set_id, user_id, exercise_id, start_date
FROM gowebapp.workouts
ORDER BY workout_id
`

// get all workouts ordered by id
func (q *Queries) ListWorkouts(ctx context.Context) ([]GowebappWorkout, error) {
	rows, err := q.db.QueryContext(ctx, listWorkouts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GowebappWorkout
	for rows.Next() {
		var i GowebappWorkout
		if err := rows.Scan(
			&i.WorkoutID,
			&i.SetID,
			&i.UserID,
			&i.ExerciseID,
			&i.StartDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSet = `-- name: UpdateSet :one
UPDATE gowebapp.sets
SET (Exercise_Id, Weight) = ($1, $2)
WHERE set_id = $3 RETURNING set_id, exercise_id, weight
`

type UpdateSetParams struct {
	ExerciseID int64 `db:"exercise_id" json:"exerciseID"`
	Weight     int32 `db:"weight" json:"weight"`
	SetID      int64 `db:"set_id" json:"setID"`
}

// insert a sets id
func (q *Queries) UpdateSet(ctx context.Context, arg UpdateSetParams) (GowebappSet, error) {
	row := q.db.QueryRowContext(ctx, updateSet, arg.ExerciseID, arg.Weight, arg.SetID)
	var i GowebappSet
	err := row.Scan(&i.SetID, &i.ExerciseID, &i.Weight)
	return i, err
}

const upsertExercise = `-- name: UpsertExercise :one
INSERT INTO gowebapp.exercises (Exercise_Name)
VALUES ($1) ON CONFLICT (Exercise_ID) DO
UPDATE
    SET Exercise_Name = EXCLUDED.Exercise_Name
    RETURNING Exercise_ID
`

// insert or update exercise of a particular id
func (q *Queries) UpsertExercise(ctx context.Context, exerciseName string) (int64, error) {
	row := q.db.QueryRowContext(ctx, upsertExercise, exerciseName)
	var exercise_id int64
	err := row.Scan(&exercise_id)
	return exercise_id, err
}

const upsertUserImage = `-- name: UpsertUserImage :one
INSERT INTO gowebapp.images (Image_Data)
VALUES ($1) ON CONFLICT (Image_ID) DO
UPDATE
    SET Image_Data = EXCLUDED.Image_Data
    RETURNING Image_ID
`

// insert or update image of a particular id
func (q *Queries) UpsertUserImage(ctx context.Context, imageData []byte) (int64, error) {
	row := q.db.QueryRowContext(ctx, upsertUserImage, imageData)
	var image_id int64
	err := row.Scan(&image_id)
	return image_id, err
}

const upsertWorkout = `-- name: UpsertWorkout :one
INSERT INTO gowebapp.workouts (User_ID, Set_ID, Start_Date)
values ($1,
        $2,
        $3) ON CONFLICT (Workout_ID) DO
UPDATE
    SET User_ID = EXCLUDED.User_ID,
    Set_ID = EXCLUDED.Set_ID,
    Start_Date = EXCLUDED.Start_Date
    RETURNING Workout_ID
`

type UpsertWorkoutParams struct {
	UserID    int64     `db:"user_id" json:"userID"`
	SetID     int64     `db:"set_id" json:"setID"`
	StartDate time.Time `db:"start_date" json:"startDate"`
}

// insert or update workouts based of a particular ID
func (q *Queries) UpsertWorkout(ctx context.Context, arg UpsertWorkoutParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, upsertWorkout, arg.UserID, arg.SetID, arg.StartDate)
	var workout_id int64
	err := row.Scan(&workout_id)
	return workout_id, err
}
