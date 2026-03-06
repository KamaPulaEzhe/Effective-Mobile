-- Справочник сервисов (чтобы не было опечаток в названиях)
CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

-- Основная таблица подписок
CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_id INTEGER REFERENCES services(id) ON DELETE RESTRICT,
    price INTEGER NOT NULL CHECK (price > 0),
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Индексы
CREATE INDEX idx_subscriptions_user_id ON subscriptions(user_id);
CREATE INDEX idx_subscriptions_service_id ON subscriptions(service_id);
CREATE INDEX idx_subscriptions_dates ON subscriptions(start_date, end_date);