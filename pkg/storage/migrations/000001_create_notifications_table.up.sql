CREATE TABLE notifications (
    "id" serial NOT NULL PRIMARY KEY,
    "title" varchar(255) NOT NULL,
    "body" text,

    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "token" varchar(16) NOT NULL, /* by */
    
    "rules" json NOT NULL
);