CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE="Europe/Moscow";

CREATE TABLE IF NOT EXISTS services (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    service_title VARCHAR (255) NOT NULL,
    service_description VARCHAR (255) NOT NULL,
    service_status VARCHAR (128) NOT NULL,
    detail_model VARCHAR (255) NOT NULL
);

CREATE TABLE IF NOT EXISTS vehicle (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    vehicle_price VARCHAR (128) NOT NULL,
    category VARCHAR (255) NOT NULL,
    model VARCHAR (255) NOT NULL,
    model_description VARCHAR (255) NOT NULL,
    model_characteristics JSONB NOT NULL,
    title VARCHAR (255) NOT NULL,
    vehicle_status VARCHAR (128) NOT NULL
);