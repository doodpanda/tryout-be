-- CreateTable
CREATE TABLE "tryout" (
    "id" UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "title" TEXT NOT NULL,
    "description" VARCHAR(255),
    "long_description" TEXT,
    "category" TEXT,
    "duration" INTEGER,
    "difficulty" TEXT,
    "passing_score" INTEGER,
    "max_attempt" INTEGER,
    "topics" TEXT[],
    "creator_id" UUID NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "is_published" BOOLEAN NOT NULL DEFAULT TRUE
);

-- CreateTable
CREATE TABLE "user" (
    "id" UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "email" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "first_name" TEXT NOT NULL,
    "last_name" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- CreateTable
CREATE TABLE "tryout_mcq_questions" (
    "id" UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "tryout_id" UUID NOT NULL,
    "question" TEXT NOT NULL,
    "correct_answer" UUID,
    "points" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY ("tryout_id") REFERENCES "tryout" ("id") ON DELETE CASCADE
);

-- CreateTable
CREATE TABLE "tryout_attempts" (
    "id" UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "tryout_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "score" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY ("tryout_id") REFERENCES "tryout" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE
);

-- CreateTable
CREATE TABLE "tryout_mcq_options" (
    "id" UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "question_id" UUID NOT NULL,
    "option" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY ("question_id") REFERENCES "tryout_mcq_questions" ("id") ON DELETE CASCADE
);

-- CreateTable
CREATE TABLE "tryout_mcq_attempts" (
    "id" UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "tryout_id" UUID NOT NULL,
    "tryout_attempt_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "question_id" UUID NOT NULL,
    "answer" UUID NOT NULL,
    "score" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY ("tryout_id") REFERENCES "tryout" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE, 
    FOREIGN KEY ("question_id") REFERENCES "tryout_mcq_questions" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("answer") REFERENCES "tryout_mcq_options" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("tryout_attempt_id") REFERENCES "tryout_attempts" ("id") ON DELETE CASCADE
);

-- CreateTable
CREATE TABLE "tryout_essay_questions" (
    "id" UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "tryout_id" UUID NOT NULL,
    "question" TEXT NOT NULL,
    "points" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY ("tryout_id") REFERENCES "tryout" ("id") ON DELETE CASCADE
);

-- CreateTable
CREATE TABLE "tryout_essay_attempts" (
    "id" UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "tryout_id" UUID NOT NULL,
    "tryout_attempt_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "question_id" UUID NOT NULL,
    "answer" TEXT NOT NULL,
    "score" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY ("tryout_id") REFERENCES "tryout" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("question_id") REFERENCES "tryout_essay_questions" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("tryout_attempt_id") REFERENCES "tryout_attempts" ("id") ON DELETE CASCADE
);