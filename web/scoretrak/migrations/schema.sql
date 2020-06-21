CREATE TABLE images (
	id UUID NOT NULL,
	image BYTES NULL,
	image_type VARCHAR(255) NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	FAMILY "primary" (id, image, image_type, created_at, updated_at)
);

CREATE TABLE schema_migration (
	version VARCHAR(14) NOT NULL,
	UNIQUE INDEX schema_migration_version_idx (version ASC),
	FAMILY "primary" (version, rowid)
);

CREATE TABLE teams (
	id UUID NOT NULL,
	name VARCHAR(255) NOT NULL,
	role VARCHAR(255) NOT NULL,
	image_id UUID NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	INDEX teams_auto_index_teams_images_id_fk (image_id ASC),
	FAMILY "primary" (id, name, role, image_id, created_at, updated_at)
);

CREATE TABLE users (
	id UUID NOT NULL,
	username VARCHAR(255) NOT NULL,
	password_hash VARCHAR(255) NOT NULL,
	team_id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	CONSTRAINT "primary" PRIMARY KEY (id ASC),
	INDEX users_auto_index_users_teams_id_fk (team_id ASC),
	FAMILY "primary" (id, username, password_hash, team_id, created_at, updated_at)
);

ALTER TABLE teams ADD CONSTRAINT teams_images_id_fk FOREIGN KEY (image_id) REFERENCES images(id);
ALTER TABLE users ADD CONSTRAINT users_teams_id_fk FOREIGN KEY (team_id) REFERENCES teams(id);

-- Validate foreign key constraints. These can fail if there was unvalidated data during the dump.
ALTER TABLE teams VALIDATE CONSTRAINT teams_images_id_fk;
ALTER TABLE users VALIDATE CONSTRAINT users_teams_id_fk;
