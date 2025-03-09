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
        '[]'::json AS options, -- Cast NULL to empty JSON array
        points
    FROM tryout_essay_questions
)
SELECT json_agg(
    json_build_object(
        'id', q.id,
        'tryoutId', q.tryout_id,
        'text', q.text,
        'type', q.type,
        'options', q.options,
        'points', q.points
    )
) AS questions
FROM (
    SELECT * FROM mcq_questions
    UNION ALL
    SELECT * FROM essay_questions
) q
WHERE q.tryout_id = $1;