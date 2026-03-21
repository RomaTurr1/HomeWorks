-- 1)
CREATE INDEX idx_users_email ON users(email);

-- 2)
CREATE INDEX idx_attempts_user_id ON attemps(user_id);

-- 3)
CREATE INDEX idx_attempts_user_id_created_at 
ON attemps(user_id, created_at DESC);

-- 4)
CREATE INDEX idx_attempts_in_progress 
ON attemps(status) 
WHERE status = 'in_progress';

-- 5)
CREATE INDEX idx_subjects_title ON subjects(title);

CREATE INDEX idx_tests_subject_id ON tests(subject_id);

CREATE INDEX idx_attempts_test_id ON attemps(test_id);

CREATE INDEX idx_attempts_user_id ON attemps(user_id);

CREATE INDEX idx_attempts_score ON attemps(score DESC);