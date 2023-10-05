CREATE TABLE users (
  id int NOT NULL AUTO_INCREMENT,
  username varchar(25) NOT NULL,
  email varchar(50) NOT NULL,
  password varchar(200) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY unique_index (username,email)
) ENGINE=InnoDB;

CREATE TABLE categories (
  id int PRIMARY KEY AUTO_INCREMENT,
  name varchar(25)
) ENGINE=InnoDB;

CREATE TABLE questions (
  id int NOT NULL AUTO_INCREMENT,
  user_id int NOT NULL,
  category_id int NOT NULL,
  content text NOT NULL,
  description text,
  PRIMARY KEY (id),
  created_at timestamp NOT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  KEY fk_user_question (user_id),
  KEY fk_category_question (category_id),
  FULLTEXT KEY questions_fulltext (content,description),
  CONSTRAINT fk_category_question FOREIGN KEY (category_id) REFERENCES categories (id),
  CONSTRAINT fk_user_question FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB;

CREATE TABLE question_likes (
  user_id int NOT NULL,
  question_id int NOT NULL,
  created_at timestamp NOT NULL,
  UNIQUE KEY unique_index (user_id,question_id),
  KEY fk_question_question_like (question_id),
  CONSTRAINT fk_question_question_like FOREIGN KEY (question_id) REFERENCES questions (id),
  CONSTRAINT fk_user_question_like FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB;

CREATE TABLE answers (
  id int NOT NULL AUTO_INCREMENT,
  question_id int NOT NULL,
  user_id int NOT NULL,
  content text NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  KEY fk_question_answer (question_id),
  KEY fk_user_answer (user_id),
  CONSTRAINT fk_question_answer FOREIGN KEY (question_id) REFERENCES questions (id),
  CONSTRAINT fk_user_answer FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB;

CREATE TABLE answer_likes (
  user_id int NOT NULL,
  answer_id int NOT NULL,
  created_at timestamp NOT NULL,
  UNIQUE KEY unique_index (user_id,answer_id),
  KEY fk_answer_answer_like (answer_id),
  CONSTRAINT fk_answer_answer_like FOREIGN KEY (answer_id) REFERENCES answers (id),
  CONSTRAINT fk_user_answer_like FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB;