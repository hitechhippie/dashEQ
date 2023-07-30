package eqobjects

type Spawngroup struct {
	Id            uint32
	Name          string
	Spawn_limit   uint8
	Dist          float32
	Max_x         float32
	Min_x         float32
	Max_y         float32
	Min_y         float32
	Delay         uint32
	Mindelay      uint32
	Despawn       uint8
	Despawn_timer uint32
	Wp_spawns     uint8
}

/*
+---------------+---------------------+------+-----+---------+----------------+
| Field         | Type                | Null | Key | Default | Extra          |
+---------------+---------------------+------+-----+---------+----------------+
| id            | int(11)             | NO   | PRI | NULL    | auto_increment |
| name          | varchar(50)         | NO   | UNI |         |                |
| spawn_limit   | tinyint(4)          | NO   |     | 0       |                |
| dist          | float               | NO   |     | 0       |                |
| max_x         | float               | NO   |     | 0       |                |
| min_x         | float               | NO   |     | 0       |                |
| max_y         | float               | NO   |     | 0       |                |
| min_y         | float               | NO   |     | 0       |                |
| delay         | int(11)             | NO   |     | 45000   |                |
| mindelay      | int(11)             | NO   |     | 15000   |                |
| despawn       | tinyint(3)          | NO   |     | 0       |                |
| despawn_timer | int(11)             | NO   |     | 100     |                |
| wp_spawns     | tinyint(1) unsigned | NO   |     | 0       |                |
+---------------+---------------------+------+-----+---------+----------------+
*/
