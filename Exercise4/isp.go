package main

type Database interface {
	Conectar()
	Desconectar()
	Insertar()
	Eliminar()
	Actualizar()
	Consultar()
}

type MiBaseDeDatos struct{}

func (db *MiBaseDeDatos) Conectar() {
	// lógica de conexión a la base de datos
}

func (db *MiBaseDeDatos) Desconectar() {
	// lógica de desconexión de la base de datos
}

func (db *MiBaseDeDatos) Insertar() {
	// lógica de inserción en la base de datos
}

func (db *MiBaseDeDatos) Eliminar() {
	// lógica de eliminación en la base de datos
}

func (db *MiBaseDeDatos) Actualizar() {
	// lógica de actualización en la base de datos
}

func (db *MiBaseDeDatos) Consultar() {
	// lógica de consulta en la base de datos
}

type Cliente struct{}

func (c *Cliente) UtilizarBaseDeDatos(db Database) {
	// lógica de uso de la base de datos
	db.Conectar()
	db.Insertar()
	db.Consultar()
	db.Desconectar()
}

type MiTabla struct{}

func (t *MiTabla) Conectar() {
	// no implementado
}

func (t *MiTabla) Desconectar() {
	// no implementado
}

func (t *MiTabla) Insertar() {
	// lógica de inserción en la tabla
}

func (t *MiTabla) Eliminar() {
	// no implementado
}

func (t *MiTabla) Actualizar() {
	// no implementado
}

func (t *MiTabla) Consultar() {
	// lógica de consulta en la tabla
}

func main() {
	db := &MiBaseDeDatos{}
	tabla := &MiTabla{}

	cliente := &Cliente{}
	cliente.UtilizarBaseDeDatos(db)
	cliente.UtilizarBaseDeDatos(tabla)
}
