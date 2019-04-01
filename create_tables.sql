-- -----------------------------------------------------
-- Table "tools"
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS  "tools" (
  "id" BIGSERIAL,
  "title" TEXT NOT NULL,
  "link" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "tags" TEXT[] NOT NULL,
  PRIMARY KEY ("id"));