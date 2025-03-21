// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: tryout-question-query.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteEssayQuestion = `-- name: DeleteEssayQuestion :exec
DELETE FROM tryout_essay_questions
WHERE id = $1
`

func (q *Queries) DeleteEssayQuestion(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteEssayQuestion, id)
	return err
}

const deleteOption = `-- name: DeleteOption :exec
DELETE FROM tryout_mcq_options
WHERE id = $1
`

func (q *Queries) DeleteOption(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteOption, id)
	return err
}

const deleteQuestion = `-- name: DeleteQuestion :exec
DELETE FROM tryout_mcq_questions
WHERE id = $1
`

func (q *Queries) DeleteQuestion(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteQuestion, id)
	return err
}

const getQuestionByID = `-- name: GetQuestionByID :one
WITH mcq_question AS (
    SELECT
        q.id,
        q.tryout_id,
        q.question AS text,
        'multiple_choice' AS type,
        json_agg(
            json_build_object(
                'id', o.id,
                'text', o.option
            )
        ) AS options,
        q.correct_answer,
        q.points
    FROM tryout_mcq_questions q
    LEFT JOIN tryout_mcq_options o ON q.id = o.question_id
    WHERE q.id = $1
    GROUP BY q.id
),
essay_question AS (
    SELECT
        id,
        tryout_id,
        question AS text,
        'essay' AS type,
        '[]'::json AS options,
        NULL::UUID AS correct_answer,
        points
    FROM tryout_essay_questions
    WHERE id = $1
)
SELECT json_build_object(
    'id', q.id,
    'tryout_id', q.tryout_id,
    'text', q.text,
    'type', q.type,
    'options', q.options,
    'points', q.points,
    'correct_answer', q.correct_answer
) AS question
FROM (
    SELECT id, tryout_id, text, type, options, correct_answer, points FROM mcq_question
    UNION ALL
    SELECT id, tryout_id, text, type, options, correct_answer, points FROM essay_question
) q
`

func (q *Queries) GetQuestionByID(ctx context.Context, id pgtype.UUID) ([]byte, error) {
	row := q.db.QueryRow(ctx, getQuestionByID, id)
	var question []byte
	err := row.Scan(&question)
	return question, err
}

const getTryoutCreatorByQuestionID = `-- name: GetTryoutCreatorByQuestionID :one
SELECT t.creator_id
FROM tryout_mcq_questions q
JOIN tryout t ON q.tryout_id = t.id
WHERE q.id = $1

UNION

SELECT t.creator_id
FROM tryout_essay_questions q
JOIN tryout t ON q.tryout_id = t.id
WHERE q.id = $1
`

func (q *Queries) GetTryoutCreatorByQuestionID(ctx context.Context, id pgtype.UUID) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, getTryoutCreatorByQuestionID, id)
	var creator_id pgtype.UUID
	err := row.Scan(&creator_id)
	return creator_id, err
}

const getTryoutQuestionsByTryoutId = `-- name: GetTryoutQuestionsByTryoutId :many
WITH mcq_questions AS (
    SELECT
        q.id,
        q.tryout_id,
        q.question AS text,
        'multiple_choice' AS type,
        json_agg(
            json_build_object(
                'id', o.id,
                'text', o.option
            )
        ) AS options,
        q.correct_answer,
        q.points
    FROM tryout_mcq_questions q
    LEFT JOIN tryout_mcq_options o ON q.id = o.question_id
    GROUP BY q.id
),
essay_questions AS (
    SELECT
        id,
        tryout_id,
        question AS text,
        'essay' AS type,
        '[]'::json AS options,
        NULL::UUID AS correct_answer,
        points
    FROM tryout_essay_questions
)
SELECT json_agg(
    json_build_object(
        'id', q.id,
        'tryout_id', q.tryout_id,
        'text', q.text,
        'type', q.type,
        'options', q.options,
        'points', q.points,
        'correct_answer', q.correct_answer
    )
) AS questions
FROM (
    SELECT id, tryout_id, text, type, options, correct_answer, points FROM mcq_questions
    UNION ALL
    SELECT id, tryout_id, text, type, options, correct_answer, points FROM essay_questions
) q
WHERE q.tryout_id = $1
`

func (q *Queries) GetTryoutQuestionsByTryoutId(ctx context.Context, tryoutID pgtype.UUID) ([][]byte, error) {
	rows, err := q.db.Query(ctx, getTryoutQuestionsByTryoutId, tryoutID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := [][]byte{}
	for rows.Next() {
		var questions []byte
		if err := rows.Scan(&questions); err != nil {
			return nil, err
		}
		items = append(items, questions)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertEssayQuestion = `-- name: InsertEssayQuestion :exec
INSERT INTO tryout_essay_questions (tryout_id, question, points)
VALUES ($1, $2, $3)
`

type InsertEssayQuestionParams struct {
	TryoutID pgtype.UUID
	Question string
	Points   int32
}

func (q *Queries) InsertEssayQuestion(ctx context.Context, arg *InsertEssayQuestionParams) error {
	_, err := q.db.Exec(ctx, insertEssayQuestion, arg.TryoutID, arg.Question, arg.Points)
	return err
}

const insertMCQQuestion = `-- name: InsertMCQQuestion :one
INSERT INTO tryout_mcq_questions (tryout_id, question, correct_answer, points)
VALUES ($1, $2, $3, $4)
RETURNING id
`

type InsertMCQQuestionParams struct {
	TryoutID      pgtype.UUID
	Question      string
	CorrectAnswer pgtype.UUID
	Points        int32
}

func (q *Queries) InsertMCQQuestion(ctx context.Context, arg *InsertMCQQuestionParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, insertMCQQuestion,
		arg.TryoutID,
		arg.Question,
		arg.CorrectAnswer,
		arg.Points,
	)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const insertOption = `-- name: InsertOption :one
INSERT INTO tryout_mcq_options (question_id, option)
VALUES ($1, $2)
RETURNING id
`

type InsertOptionParams struct {
	QuestionID pgtype.UUID
	Option     string
}

func (q *Queries) InsertOption(ctx context.Context, arg *InsertOptionParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, insertOption, arg.QuestionID, arg.Option)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const updateEssayQuestion = `-- name: UpdateEssayQuestion :exec
UPDATE tryout_essay_questions
SET question = $2,
    points = $3
WHERE id = $1
`

type UpdateEssayQuestionParams struct {
	ID       pgtype.UUID
	Question string
	Points   int32
}

func (q *Queries) UpdateEssayQuestion(ctx context.Context, arg *UpdateEssayQuestionParams) error {
	_, err := q.db.Exec(ctx, updateEssayQuestion, arg.ID, arg.Question, arg.Points)
	return err
}

const updateMCQQuestion = `-- name: UpdateMCQQuestion :exec
UPDATE tryout_mcq_questions
SET question = $2,
    correct_answer = $3,
    points = $4
WHERE id = $1
`

type UpdateMCQQuestionParams struct {
	ID            pgtype.UUID
	Question      string
	CorrectAnswer pgtype.UUID
	Points        int32
}

func (q *Queries) UpdateMCQQuestion(ctx context.Context, arg *UpdateMCQQuestionParams) error {
	_, err := q.db.Exec(ctx, updateMCQQuestion,
		arg.ID,
		arg.Question,
		arg.CorrectAnswer,
		arg.Points,
	)
	return err
}

const updateOption = `-- name: UpdateOption :exec
UPDATE tryout_mcq_options
SET option = $2
WHERE id = $1
`

type UpdateOptionParams struct {
	ID     pgtype.UUID
	Option string
}

func (q *Queries) UpdateOption(ctx context.Context, arg *UpdateOptionParams) error {
	_, err := q.db.Exec(ctx, updateOption, arg.ID, arg.Option)
	return err
}
