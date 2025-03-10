-- name: GetTryoutQuestionsByTryoutId :many
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
    SELECT * FROM mcq_questions
    UNION ALL
    SELECT * FROM essay_questions
) q
WHERE q.tryout_id = $1;

-- name: GetQuestionByID :one
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
    SELECT * FROM mcq_question
    UNION ALL
    SELECT * FROM essay_question
) q;

-- name: InsertMCQQuestion :one
INSERT INTO tryout_mcq_questions (tryout_id, question, correct_answer, points)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: InsertEssayQuestion :exec
INSERT INTO tryout_essay_questions (tryout_id, question, points)
VALUES ($1, $2, $3);

-- name: InsertOption :one
INSERT INTO tryout_mcq_options (question_id, option)
VALUES ($1, $2)
RETURNING id;

-- name: UpdateMCQQuestion :exec
UPDATE tryout_mcq_questions
SET question = $2,
    correct_answer = $3,
    points = $4
WHERE id = $1;

-- name: UpdateEssayQuestion :exec
UPDATE tryout_essay_questions
SET question = $2,
    points = $3
WHERE id = $1;

-- name: UpdateOption :exec
UPDATE tryout_mcq_options
SET option = $2
WHERE id = $1;

-- name: DeleteQuestion :exec
DELETE FROM tryout_mcq_questions
WHERE id = $1;

-- name: DeleteEssayQuestion :exec
DELETE FROM tryout_essay_questions
WHERE id = $1;

-- name: DeleteOption :exec
DELETE FROM tryout_mcq_options
WHERE id = $1;
