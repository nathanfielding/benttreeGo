
-- SQL Schema for Apartment, Tenant, Lease, and MaintenanceRequest models

-- Apartment Model
CREATE TABLE IF NOT EXISTS Apartment (
    id INT AUTO_INCREMENT PRIMARY KEY,
    number VARCHAR(50) NOT NULL UNIQUE,
    property VARCHAR(50) NOT NULL,
    bedrooms INT NOT NULL,
    occupancy INT NOT NULL,
    rented_as INT
);

-- Tenant Model
CREATE TABLE IF NOT EXISTS Tenant (
    id INT AUTO_INCREMENT PRIMARY KEY,
    apartment_id INT NOT NULL,
    lease_id INT UNIQUE NOT NULL,
    name VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    phone_number VARCHAR(50) NOT NULL UNIQUE,
    home_address VARCHAR(50) NOT NULL,
    is_renewing BOOLEAN,
    FOREIGN KEY (apartment_id) REFERENCES Apartment(id) ON DELETE PROTECT,
    FOREIGN KEY (lease_id) REFERENCES Lease(id) ON DELETE CASCADE
);

-- Lease Model
CREATE TABLE IF NOT EXISTS Lease (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tenant_id INT UNIQUE,
    apartment_id INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    monthly_rent DECIMAL(10, 2) NOT NULL,
    deposit_amount DECIMAL(10, 2) NOT NULL,
    copy_of_lease VARCHAR(255),
    FOREIGN KEY (tenant_id) REFERENCES Tenant(id) ON DELETE PROTECT,
    FOREIGN KEY (apartment_id) REFERENCES Apartment(id) ON DELETE PROTECT
);

-- MaintenanceRequest Model
CREATE TABLE IF NOT EXISTS MaintenanceRequest (
    id INT AUTO_INCREMENT PRIMARY KEY,
    apartment_id INT NOT NULL,
    open_date DATE NOT NULL,
    close_date DATE,
    description VARCHAR(100),
    FOREIGN KEY (apartment_id) REFERENCES Apartment(id) ON DELETE PROTECT
);
