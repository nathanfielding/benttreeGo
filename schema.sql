CREATE TABLE IF NOT EXISTS Apartments (
    id SERIAL PRIMARY KEY,
    number VARCHAR(50) NOT NULL UNIQUE,
    property VARCHAR(50) NOT NULL,
    bedrooms INT NOT NULL,
    occupancy INT NOT NULL,
    rented_as INT
);

CREATE TABLE IF NOT EXISTS Tenants (
    id SERIAL PRIMARY KEY,
    apartment_id INT NOT NULL REFERENCES Apartments(id) ON DELETE RESTRICT,
    lease_id INT UNIQUE,
    name VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    phone_number VARCHAR(50) NOT NULL UNIQUE,
    home_address VARCHAR(50) NOT NULL,
    is_renewing BOOLEAN
);

CREATE TABLE IF NOT EXISTS Leases (
    id SERIAL PRIMARY KEY,
    apartment_id INT NOT NULL REFERENCES Apartments(id) ON DELETE RESTRICT,
    tenant_id INT UNIQUE,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    monthly_rent DECIMAL(10, 2) NOT NULL,
    deposit_amount DECIMAL(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS MaintenanceRequests (
    id SERIAL PRIMARY KEY,
    apartment_id INT NOT NULL REFERENCES Apartments(id) ON DELETE RESTRICT,
    open_date DATE NOT NULL,
    close_date DATE,
    description VARCHAR(100)
);

ALTER TABLE Tenants ADD FOREIGN KEY (lease_id) REFERENCES Leases(id) ON DELETE RESTRICT;
ALTER TABLE Leases ADD FOREIGN KEY (tenant_id) REFERENCES Tenants(id) ON DELETE CASCADE;
