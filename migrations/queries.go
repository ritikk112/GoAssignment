package migrations

const (
	createEmployeeTable = "CREATE TABLE IF NOT EXISTS employee (ID INT AUTO_INCREMENT PRIMARY KEY, NAME VARCHAR(255) NOT NULL, DEPT VARCHAR(255) NOT NULL);"
	dropEmployeeTable   = "DROP TABLE IF EXISTS employee"
)
