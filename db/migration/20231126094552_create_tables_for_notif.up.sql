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