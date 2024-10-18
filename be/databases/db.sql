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
    SupplierName VARCHAR(255) NOT NULL
);

CREATE TABLE MasterCustomer (
    CustomerPK SERIAL PRIMARY KEY,
    CustomerName VARCHAR(255) NOT NULL
);

CREATE TABLE MasterProduct (
    ProductPK SERIAL PRIMARY KEY,
    ProductName VARCHAR(255) NOT NULL
);

CREATE TABLE MasterWarehouse (
    WhsPK SERIAL PRIMARY KEY,
    WhsName VARCHAR(255) NOT NULL
);

CREATE TABLE TransaksiPenerimaanBarangHeader (
    TrxInPK SERIAL PRIMARY KEY,
    TrxInNo VARCHAR(50) NOT NULL,
    WhsIdf INT NOT NULL,
    TrxInDate DATE NOT NULL,
    TrxInSuppIdf INT NOT NULL,
    TrxInNotes VARCHAR(255),
    FOREIGN KEY (WhsIdf) REFERENCES MasterWarehouse(WhsPK),
    FOREIGN KEY (TrxInSuppIdf) REFERENCES MasterSupplier(SupplierPK)
);

CREATE TABLE TransaksiPenerimaanBarangDetail (
    TrxInDPK SERIAL PRIMARY KEY,
    TrxInIDF INT NOT NULL,
    TrxInDProductIdf INT NOT NULL,
    TrxInDQtyDus INT NOT NULL,
    TrxInDQtyPcs INT NOT NULL,
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
    FOREIGN KEY (WhsIdf) REFERENCES MasterWarehouse(WhsPK),
    FOREIGN KEY (TrxOutSuppIdf) REFERENCES MasterSupplier(SupplierPK)
);

CREATE TABLE TransaksiPengeluaranBarangDetail (
    TrxOutDPK SERIAL PRIMARY KEY,
    TrxOutIDF INT NOT NULL,
    TrxOutDProductIdf INT NOT NULL,
    TrxOutDQtyDus INT NOT NULL,
    TrxOutDQtyPcs INT NOT NULL,
    FOREIGN KEY (TrxOutIDF) REFERENCES TransaksiPengeluaranBarangHeader(TrxOutPK),
    FOREIGN KEY (TrxOutDProductIdf) REFERENCES MasterProduct(ProductPK)
);

INSERT INTO MasterSupplier (SupplierName) 
VALUES 
    ('Supplier A'),
    ('Supplier B'),
    ('Supplier C');

INSERT INTO MasterCustomer (CustomerName) 
VALUES 
    ('Customer A'),
    ('Customer B'),
    ('Customer C');

INSERT INTO MasterProduct (ProductName) 
VALUES 
    ('Product A'),
    ('Product B'),
    ('Product C');

INSERT INTO MasterWarehouse (WhsName) 
VALUES 
    ('Warehouse A');


INSERT INTO TransaksiPenerimaanBarangHeader (TrxInNo, WhsIdf, TrxInDate, TrxInSuppIdf, TrxInNotes)
VALUES 
    ('IN-001', 1, '2024-10-01', 1, 'Penerimaan pertama dari Supplier A'),
    ('IN-002', 1, '2024-10-05', 2, 'Penerimaan kedua dari Supplier B');

INSERT INTO TransaksiPenerimaanBarangDetail (TrxInIDF, TrxInDProductIdf, TrxInDQtyDus, TrxInDQtyPcs)
VALUES 
    (1, 1, 10, 100),
    (1, 2, 5, 50),
    (2, 3, 8, 80);

INSERT INTO TransaksiPengeluaranBarangHeader (TrxOutNo, WhsIdf, TrxOutDate, TrxOutSuppIdf, TrxOutNotes)
VALUES 
    ('OUT-001', 1, '2024-10-10', 1, 'Pengeluaran pertama ke Customer A'),
    ('OUT-002', 1, '2024-10-12', 2, 'Pengeluaran kedua ke Customer B');

INSERT INTO TransaksiPengeluaranBarangDetail (TrxOutIDF, TrxOutDProductIdf, TrxOutDQtyDus, TrxOutDQtyPcs)
VALUES 
    (1, 1, 3, 30),
    (1, 2, 2, 20),
    (2, 3, 4, 40);
