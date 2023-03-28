package main

type Insertable interface {
	Insertar()
}

type Consultable interface {
	Consultar()
}

type Database interface {
	Conectar()
	Desconectar()
}

type MiBaseDeDatos struct{}

func (db *MiBaseDeDatos) Conectar() {
	// lógica de conexión a la base de datos
}

func (db *MiBaseDeDatos) Desconectar() {
	// lógica de desconexión de la base de datos
}

type Cliente struct{}

func (c *Cliente) Insertar(db Insertable) {
	db.Insertar()
}

func (c *Cliente) Consultar(db Consultable) {
	db.Consultar()
}

type MiTabla struct{}

func (t *MiTabla) Insertar() {
	// lógica de inserción en la tabla
}

func (t *MiTabla) Consultar() {
	// lógica de consulta en la tabla
}

func main() {
	//db := &MiBaseDeDatos{}
	tabla := &MiTabla{}

	cliente := &Cliente{}
	cliente.Insertar(tabla)
	cliente.Consultar(tabla)
}
