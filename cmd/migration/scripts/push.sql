CREATE TABLE feature_flags(flag VARCHAR(50) PRIMARY KEY NOT NULL);

CREATE TABLE users(
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  employee_id VARCHAR(7) UNIQUE NOT NULL,
  name VARCHAR(100) NOT NULL,
  password VARCHAR(255) NOT NULL
);

CREATE INDEX idx_employee_id ON users(employee_id);

CREATE TABLE users_feature_flags(
  user_id VARCHAR(255) NOT NULL,
  flag VARCHAR(50) NOT NULL,
  PRIMARY KEY (user_id, flag),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (flag) REFERENCES feature_flags(flag)
);
