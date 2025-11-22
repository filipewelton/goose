-- Inserir feature flags
INSERT INTO
  feature_flags
VALUES
  ('administration');

INSERT INTO
  feature_flags
VALUES
  ('supervision');

INSERT INTO
  feature_flags
VALUES
  ('chat');

INSERT INTO
  feature_flags
VALUES
  ('expiration_date');

INSERT INTO
  feature_flags
VALUES
  ('notes');

INSERT INTO
  feature_flags
VALUES
  ('reminders');

-- Inserir usuário administrador
INSERT INTO
  users
VALUES
  (
    '9ffb8542-0505-4a62-a11a-8826bd8735c5',
    '5663091',
    'Filipe W. Silva',
    '$2a$12$jlinVpu2zZ3pkjk41NFazu76yZz1lCnds1tV3ESlHB5YpHUxM1K8i'
  );

-- Atribuir feature flags ao usuário administrador
INSERT INTO
  users_feature_flags
VALUES
  (
    '9ffb8542-0505-4a62-a11a-8826bd8735c5',
    'administration'
  );

INSERT INTO
  users_feature_flags
VALUES
  (
    '9ffb8542-0505-4a62-a11a-8826bd8735c5',
    'supervision'
  );

INSERT INTO
  users_feature_flags
VALUES
  (
    '9ffb8542-0505-4a62-a11a-8826bd8735c5',
    'chat'
  );

INSERT INTO
  users_feature_flags
VALUES
  (
    '9ffb8542-0505-4a62-a11a-8826bd8735c5',
    'expiration_date'
  );

INSERT INTO
  users_feature_flags
VALUES
  (
    '9ffb8542-0505-4a62-a11a-8826bd8735c5',
    'notes'
  );

INSERT INTO
  users_feature_flags
VALUES
  (
    '9ffb8542-0505-4a62-a11a-8826bd8735c5',
    'reminders'
  );
