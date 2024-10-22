DROP TABLE IF EXISTS MasterSupplier CASCADE;
DROP TABLE IF EXISTS MasterCustomer CASCADE;
DROP TABLE IF EXISTS MasterProduct CASCADE;
DROP TABLE IF EXISTS MasterWarehouse CASCADE;
DROP TABLE IF EXISTS TransaksiPenerimaanBarangHeader CASCADE;
DROP TABLE IF EXISTS TransaksiPenerimaanBarangDetail CASCADE;
DROP TABLE IF EXISTS TransaksiPengeluaranBarangHeader CASCADE;
DROP TABLE IF EXISTS TransaksiPengeluaranBarangDetail CASCADE;

CREATE TABLE MasterSupplier (
    SupplierPK SERIAL PRIMARY KEY,
    SupplierName VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE MasterCustomer (
    CustomerPK SERIAL PRIMARY KEY,
    CustomerName VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE MasterProduct (
    ProductPK SERIAL PRIMARY KEY,
    ProductName VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE MasterWarehouse (
    WhsPK SERIAL PRIMARY KEY,
    WhsName VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE TransaksiPenerimaanBarangHeader (
    TrxInPK SERIAL PRIMARY KEY,
    TrxInNo VARCHAR(50) NOT NULL,
    WhsIdf INT NOT NULL,
    TrxInDate DATE NOT NULL,
    TrxInSuppIdf INT NOT NULL,
    TrxInNotes VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    FOREIGN KEY (WhsIdf) REFERENCES MasterWarehouse(WhsPK),
    FOREIGN KEY (TrxInSuppIdf) REFERENCES MasterSupplier(SupplierPK)
);

CREATE TABLE TransaksiPenerimaanBarangDetail (
    TrxInDPK SERIAL PRIMARY KEY,
    TrxInIDF INT NOT NULL,
    TrxInDProductIdf INT NOT NULL,
    TrxInDQtyDus INT NOT NULL,
    TrxInDQtyPcs INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    FOREIGN KEY (TrxInIDF) REFERENCES TransaksiPenerimaanBarangHeader(TrxInPK),
    FOREIGN KEY (TrxInDProductIdf) REFERENCES MasterProduct(ProductPK)
);

CREATE TABLE TransaksiPengeluaranBarangHeader (
    TrxOutPK SERIAL PRIMARY KEY,
    TrxOutNo VARCHAR(50) NOT NULL,
    WhsIdf INT NOT NULL,
    TrxOutDate DATE NOT NULL,
    TrxOutSuppIdf INT NOT NULL,
    TrxOutNotes VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    FOREIGN KEY (WhsIdf) REFERENCES MasterWarehouse(WhsPK),
    FOREIGN KEY (TrxOutSuppIdf) REFERENCES MasterSupplier(SupplierPK)
);

CREATE TABLE TransaksiPengeluaranBarangDetail (
    TrxOutDPK SERIAL PRIMARY KEY,
    TrxOutIDF INT NOT NULL,
    TrxOutDProductIdf INT NOT NULL,
    TrxOutDQtyDus INT NOT NULL,
    TrxOutDQtyPcs INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    FOREIGN KEY (TrxOutIDF) REFERENCES TransaksiPengeluaranBarangHeader(TrxOutPK),
    FOREIGN KEY (TrxOutDProductIdf) REFERENCES MasterProduct(ProductPK)
);

INSERT INTO MasterSupplier (SupplierName, created_at, updated_at) 
VALUES 
    ('Supplier A', '2024-09-10', '2024-09-10'),
    ('Supplier B', '2024-09-15', '2024-09-15'),
    ('Supplier C', '2024-10-01', '2024-10-01'),
    ('Supplier D', '2024-10-05', '2024-10-05'),
    ('Supplier E', '2024-10-10', '2024-10-10');

INSERT INTO MasterCustomer (CustomerName, created_at, updated_at) 
VALUES 
    ('Customer A', '2024-09-12', '2024-09-12'),
    ('Customer B', '2024-09-18', '2024-09-18'),
    ('Customer C', '2024-10-02', '2024-10-02'),
    ('Customer D', '2024-10-06', '2024-10-06'),
    ('Customer E', '2024-10-11', '2024-10-11');

INSERT INTO MasterProduct (ProductName, created_at, updated_at) 
VALUES 
    ('Product A', '2024-09-13', '2024-09-13'),
    ('Product B', '2024-09-20', '2024-09-20'),
    ('Product C', '2024-10-03', '2024-10-03'),
    ('Product D', '2024-10-07', '2024-10-07'),
    ('Product E', '2024-10-12', '2024-10-12');

INSERT INTO MasterWarehouse (WhsName, created_at, updated_at) 
VALUES 
    ('Warehouse A', '2024-09-14', '2024-09-14'),
    ('Warehouse B', '2024-09-21', '2024-09-21'),
    ('Warehouse C', '2024-10-04', '2024-10-04'),
    ('Warehouse D', '2024-10-08', '2024-10-08'),
    ('Warehouse E', '2024-10-13', '2024-10-13');

INSERT INTO TransaksiPenerimaanBarangHeader (TrxInNo, WhsIdf, TrxInDate, TrxInSuppIdf, TrxInNotes, created_at, updated_at)
VALUES 
    ('IN-001', 1, '2024-09-10', 1, 'Penerimaan pertama dari Supplier A', '2024-09-10', '2024-09-10'),
    ('IN-002', 2, '2024-09-15', 2, 'Penerimaan kedua dari Supplier B', '2024-09-15', '2024-09-15'),
    ('IN-003', 3, '2024-10-02', 3, 'Penerimaan ketiga dari Supplier C', '2024-10-02', '2024-10-02'),
    ('IN-004', 4, '2024-10-06', 4, 'Penerimaan keempat dari Supplier D', '2024-10-06', '2024-10-06'),
    ('IN-005', 5, '2024-10-10', 5, 'Penerimaan kelima dari Supplier E', '2024-10-10', '2024-10-10');

INSERT INTO TransaksiPenerimaanBarangDetail (TrxInIDF, TrxInDProductIdf, TrxInDQtyDus, TrxInDQtyPcs, created_at, updated_at)
VALUES 
    (1, 1, 10, 100, '2024-09-10', '2024-09-10'),
    (2, 2, 5, 50, '2024-09-15', '2024-09-15'),
    (3, 3, 8, 80, '2024-10-02', '2024-10-02'),
    (4, 4, 7, 70, '2024-10-06', '2024-10-06'),
    (5, 5, 6, 60, '2024-10-10', '2024-10-10');

INSERT INTO TransaksiPengeluaranBarangHeader (TrxOutNo, WhsIdf, TrxOutDate, TrxOutSuppIdf, TrxOutNotes, created_at, updated_at)
VALUES 
    ('OUT-001', 1, '2024-09-12', 1, 'Pengeluaran pertama ke Customer A', '2024-09-12', '2024-09-12'),
    ('OUT-002', 2, '2024-09-18', 2, 'Pengeluaran kedua ke Customer B', '2024-09-18', '2024-09-18'),
    ('OUT-003', 3, '2024-10-03', 3, 'Pengeluaran ketiga ke Customer C', '2024-10-03', '2024-10-03'),
    ('OUT-004', 4, '2024-10-07', 4, 'Pengeluaran keempat ke Customer D', '2024-10-07', '2024-10-07'),
    ('OUT-005', 5, '2024-10-11', 5, 'Pengeluaran kelima ke Customer E', '2024-10-11', '2024-10-11');

INSERT INTO TransaksiPengeluaranBarangDetail (TrxOutIDF, TrxOutDProductIdf, TrxOutDQtyDus, TrxOutDQtyPcs, created_at, updated_at)
VALUES 
    (1, 1, 3, 30, '2024-09-12', '2024-09-12'),
    (2, 2, 2, 20, '2024-09-18', '2024-09-18'),
    (3, 3, 4, 40, '2024-10-03', '2024-10-03'),
    (4, 4, 5, 50, '2024-10-07', '2024-10-07'),
    (5, 5, 6, 60, '2024-10-11', '2024-10-11');
