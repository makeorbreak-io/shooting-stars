CREATE TABLE matches
(
  user1 integer,
  user2 integer,
  start_time timestamp without time zone,
  player1_shoot_time timestamp without time zone,
  player2_shoot_time timestamp without time zone,
  id integer NOT NULL,
  CONSTRAINT matches_pkey PRIMARY KEY (id)
)

CREATE TABLE users
(
  password_hash character varying,
  email character varying,
  id integer NOT NULL,
  CONSTRAINT users_pkey PRIMARY KEY (id)
)