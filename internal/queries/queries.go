package queries

const InsertRegEntrance = `INSERT INTO Visit(id_Client, Client, Plate, Car, St, 
    date_reg, date_ent, date_out, Type, Location) VALUES (
    1, 'John Doe', 'ABC123', 'Toyota Camry', 'ACTIVE', NOW(), NOW(), NULL, 'VISIT', 'Main');`
