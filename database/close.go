package database

func (d *Todo) Close() {
	d.db.Close()
}
