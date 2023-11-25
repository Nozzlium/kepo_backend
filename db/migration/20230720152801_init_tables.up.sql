CREATE TABLE users (
  id int NOT NULL AUTO_INCREMENT,
  username varchar(25) NOT NULL,
  email varchar(50) NOT NULL,
  password varchar(200) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY unique_username_index (username),
  UNIQUE KEY unique_email_index (email)
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

CREATE TABLE notification_types (
  notif_type varchar(10) NOT NULL PRIMARY KEY
) ENGINE=InnoDB;

CREATE TABLE notifications (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id int NOT NULL,
  question_id int NOT NULL,
  notif_type varchar(10) NOT NULL,
  headline varchar(100) NOT NULL,
  preview varchar(100) NOT NULL,
  created_at timestamp NOT NULL,
  KEY fk_users_notification (user_id),
  KEY fk_questions_notification (question_id),
  KEY fk_notification_types_notification (notif_type),
  CONSTRAINT fk_users_notification FOREIGN KEY (user_id) REFERENCES users (id),
  CONSTRAINT fk_questions_notification FOREIGN KEY (question_id) REFERENCES questions (id),
  CONSTRAINT fk_notification_types_notification FOREIGN KEY (notif_type) REFERENCES notification_types (notif_type)
) ENGINE=InnoDB;