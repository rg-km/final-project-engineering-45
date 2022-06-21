package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/rightway.db")
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	sqlStmt :=
		`CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		fullname VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL,
		username VARCHAR(64) NOT NULL,
		password VARCHAR(64) NOT NULL,
		role VARCHAR(64) NOT NULL,
		logged_in BOOLEAN NOT NULL
	);
		
	INSERT INTO users (fullname, email, username, password, role, logged_in) 
	VALUES ('admin rightway','admin@admin.com', 'admin', 'admin', 'admin', false);

	CREATE TABLE IF NOT EXISTS fakultas (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		fakultas_name VARCHAR(64) NOT NULL,
		created_at DATETIME NOT NULL
	);

	INSERT or IGNORE INTO fakultas (id, fakultas_name, created_at)
	VALUES
	(1001, 'Fakultas Teknik', '2022-06-14 14:20:00'),
	(1002, 'Fakultas Seni Rupa dan Desain', '2022-06-14 14:20:00'),
	(1003, 'Fakultas Psikologi', '2022-06-14 14:20:00'),
	(1004, 'Fakultas Peternakan dan Pertanian', '2022-06-14 14:20:00'),
	(1005, 'Fakultas Perikanan dan Kelautan', '2022-06-14 14:20:00'),
	(1006, 'Fakultas Matematika dan Ilmu Pengetahuan Alam', '2022-06-14 14:20:00'),
	(1007, 'Fakultas Kesehatan Masyarakat', '2022-06-14 14:20:00'),
	(1008, 'Fakultas Kehutanan', '2022-06-14 14:20:00'),
	(1009, 'Fakultas Kedokteran Hewan', '2022-06-14 14:20:00'),
	(1010, 'Fakultas Kedokteran Gigi', '2022-06-14 14:20:00'),
	(1011, 'Fakultas Kedokteran', '2022-06-14 14:20:00'),
	(1012, 'Fakultas Ilmu Sosial dan Ilmu Politik', '2022-06-14 14:20:00'),
	(1013, 'Fakultas Ilmu Budaya', '2022-06-14 14:20:00'),
	(1014, 'Fakultas Hukum', '2022-06-14 14:20:00'),
	(1015, 'Fakultas Filsafat', '2022-06-14 14:20:00'),
	(1016, 'Fakultas Farmasi', '2022-06-14 14:20:00'),
	(1017, 'Fakultas Ekonomi dan Bisnis', '2022-06-14 14:20:00'),
	(1018, 'Fakultas Ekologi Manusia', '2022-06-14 14:20:00')
	;

	CREATE TABLE IF NOT EXISTS program_studi (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		prodi_name VARCHAR(64) NOT NULL,
		fakultas_id INTEGER NOT NULL,
		created_at DATETIME NOT NULL,
		FOREIGN KEY (fakultas_id) REFERENCES fakultas(id)
	);
	
	INSERT INTO program_studi (id, prodi_name, fakultas_id, created_at)
	VALUES
	(100101, 'Arsitektur', 1001, '2022-06-14 15:09:00'),
	(100102, 'Arsitektur Lansekap', 1001, '2022-06-14 15:09:00'),
	(100103, 'Ilmu Komputer', 1001, '2022-06-14 15:09:00'),
	(100104, 'Manajemen Rekayasa Industri', 1001, '2022-06-14 15:09:00'),
	(100105, 'Perencanaan Tata Ruang dan Pertanahan', 1001, '2022-06-14 15:09:00'),
	(100106, 'Perencanaan Wilayah dan Kota', 1001, '2022-06-14 15:09:00'),
	(100107, 'Teknik Metalurgi dan Material', 1001, '2022-06-14 15:09:00'),
	(100108, 'Rekayasa Infrastruktur Lingkungan', 1001, '2022-06-14 15:09:00'),
	(100109, 'Rekayasa Perancangan Mekanik', 1001, '2022-06-14 15:09:00'),
	(100110, 'Sistem dan Teknologi Informasi', 1001, '2022-06-14 15:09:00'),
	(100111, 'Teknik Arsitektur Interior', 1001, '2022-06-14 15:09:00'),
	(100112, 'Teknik Bioenergi dan Kemurgi', 1001, '2022-06-14 15:09:00'),
	(100113, 'Teknik Biomedis', 1001, '2022-06-14 15:09:00'),
	(100114, 'Teknik Bioproses', 1001, '2022-06-14 15:09:00'),
	(100115, 'Teknik dan Pengelolaan Sumber Daya Air', 1001, '2022-06-14 15:09:00'),
	(100116, 'Teknik Dirgantara', 1001, '2022-06-14 15:09:00'),
	(100117, 'Teknik Elektro', 1001, '2022-06-14 15:09:00'),
	(100118, 'Teknik Fisika', 1001, '2022-06-14 15:09:00'),
	(100119, 'Teknik Geodesi', 1001, '2022-06-14 15:09:00'),
	(100120, 'Teknik Geomatika', 1001, '2022-06-14 15:09:00'),
	(100121, 'Teknik Geofisika', 1001, '2022-06-14 15:09:00'),
	(100122, 'Teknik Geologi', 1001, '2022-06-14 15:09:00'),
	(100123, 'Teknik Industri', 1001, '2022-06-14 15:09:00'),
	(100124, 'Teknik Industri Pertanian', 1001, '2022-06-14 15:09:00'),
	(100125, 'Teknik Infrastruktur Lingkungan', 1001, '2022-06-14 15:09:00'),
	(100126, 'Teknik Infrastruktur Sipil dan Perancangan Arsitektur', 1001, '2022-06-14 15:09:00'),
	(100127, 'Teknik Kelautan', 1001, '2022-06-14 15:09:00'),
	(100128, 'Teknik Kimia', 1001, '2022-06-14 15:09:00'),
	(100129, 'Teknik Lingkungan', 1001, '2022-06-14 15:09:00'),
	(100130, 'Teknik Listrik Industri', 1001, '2022-06-14 15:09:00'),
	(100131, 'Teknik Material', 1001, '2022-06-14 15:09:00'),
	(100132, 'Teknik Mesin', 1001, '2022-06-14 15:09:00'),
	(100133, 'Teknik Nuklir', 1001, '2022-06-14 15:09:00'),
	(100134, 'Teknik Pangan', 1001, '2022-06-14 15:09:00'),
	(100135, 'Teknik Perkapalan', 1001, '2022-06-14 15:09:00'),
	(100136, 'Teknik Perminyakan', 1001, '2022-06-14 15:09:00'),
	(100137, 'Teknik Pertambangan', 1001, '2022-06-14 15:09:00'),
	(100138, 'Teknik Sipil', 1001, '2022-06-14 15:09:00'),
	(100139, 'Teknik Telekomunikasi', 1001, '2022-06-14 15:09:00'),
	(100140, 'Teknik Tenaga Listrik', 1001, '2022-06-14 15:09:00'),
	(100141, 'Teknologi dan Manajemen Perikanan Budidaya', 1001, '2022-06-14 15:09:00'),
	(100142, 'Teknologi dan Manajemen Perikanan Tangkap', 1001, '2022-06-14 15:09:00'),
	(100143, 'Teknologi Hasil Perairan', 1001, '2022-06-14 15:09:00'),
	(100144, 'Teknologi Hasil Perikanan', 1001, '2022-06-14 15:09:00'),
	(100145, 'Teknologi Hasil Ternak', 1001, '2022-06-14 15:09:00'),
	(100146, 'Teknologi Industri Pertanian', 1001, '2022-06-14 15:09:00'),
	(100147, 'Teknologi Pascapanen', 1001, '2022-06-14 15:09:00'),
	(100148, 'Teknologi Produksi Ternak', 1001, '2022-06-14 15:09:00'),
	(100149, 'Teknologi Rekayasa Kimia Industri', 1001, '2022-06-14 15:09:00'),
	(100150, 'Teknologi Rekayasa Otomasi', 1001, '2022-06-14 15:09:00'),

	(100201, 'Desain Interior', 1002, '2022-06-14 15:09:00'),
	(100202, 'Desain Komunikasi Visual', 1002, '2022-06-14 15:09:00'),
	(100203, 'Desain Produk', 1002, '2022-06-14 15:09:00'),
	(100204, 'Kriya', 1002, '2022-06-14 15:09:00'),
	(100205, 'Seni Rupa', 1002, '2022-06-14 15:09:00'),

	(100301, 'Psikologi', 1003, '2022-06-14 15:09:00'),

	(100401, 'Agribisnis', 1004, '2022-06-14 15:09:00'),
	(100402, 'Agroekoteknologi', 1004, '2022-06-14 15:09:00'),
	(100403, 'Agronomi dan Hortikultura', 1004, '2022-06-14 15:09:00'),
	(100404, 'Ilmu dan Industri Peternakan', 1004, '2022-06-14 15:09:00'),
	(100405, 'Ilmu Tanah', 1004, '2022-06-14 15:09:00'),
	(100406, 'Kartografi dan Penginderaan Jauh', 1004, '2022-06-14 15:09:00'),
	(100407, 'Manajemen Sumberdaya Lahan', 1004, '2022-06-14 15:09:00'),
	(100408, 'Nutrisi dan Teknologi Pakan', 1004, '2022-06-14 15:09:00'),
	(100409, 'Penyuluhan dan Komunikasi Pertanian', 1004, '2022-06-14 15:09:00'),
	(100410, 'Peternakan', 1004, '2022-06-14 15:09:00'),
	(100411, 'Proteksi Tanaman', 1004, '2022-06-14 15:09:00'),
	(100412, 'Rekayasa Pertanian', 1004, '2022-06-14 15:09:00'),

	(100501, 'Akuakultur', 1005, '2022-06-14 15:09:00'),
	(100502, 'Ilmu dan Teknologi Kelautan', 1005, '2022-06-14 15:09:00'),
	(100503, 'Manajemen Sumberdaya Akuatik', 1005, '2022-06-14 15:09:00'),
	(100504, 'Manajemen Sumberdaya Perairan', 1005, '2022-06-14 15:09:00'),
	(100505, 'Oseanografi', 1005, '2022-06-14 15:09:00'),

	(100601, 'Aktuaria', 1006, '2022-06-14 15:09:00'),
	(100602, 'Astronomi', 1006, '2022-06-14 15:09:00'),
	(100603, 'Biokimia', 1006, '2022-06-14 15:09:00'),
	(100604, 'Biologi', 1006, '2022-06-14 15:09:00'),
	(100605, 'Bioteknologi', 1006, '2022-06-14 15:09:00'),
	(100606, 'Elektronika dan Instrumentasi', 1006, '2022-06-14 15:09:00'),
	(100607, 'Fisika', 1006, '2022-06-14 15:09:00'),
	(100608, 'Geofisika', 1006, '2022-06-14 15:09:00'),
	(100609, 'Geografi', 1006, '2022-06-14 15:09:00'),
	(100610, 'Geografi Lingkungan', 1006, '2022-06-14 15:09:00'),
	(100611, 'Informatika', 1006, '2022-06-14 15:09:00'),
	(100612, 'Kimia', 1006, '2022-06-14 15:09:00'),
	(100613, 'Matematika', 1006, '2022-06-14 15:09:00'),
	(100614, 'Meteorologi', 1006, '2022-06-14 15:09:00'),
	(100615, 'Mikrobiologi', 1006, '2022-06-14 15:09:00'),
	(100616, 'Mikrobiologi Pertanian', 1006, '2022-06-14 15:09:00'),
	(100617, 'Rekayasa Hayati', 1006, '2022-06-14 15:09:00'),
	(100618, 'Statistika', 1006, '2022-06-14 15:09:00'),

	(100701, 'Kesehatan Lingkungan', 1007, '2022-06-14 15:09:00'),
	(100702, 'Kesehatan Masyarakat', 1007, '2022-06-14 15:09:00'),

	(100801, 'Hasil Hutan', 1008, '2022-06-14 15:09:00'),
	(100802, 'Kehutanan', 1008, '2022-06-14 15:09:00'),
	(100803, 'Konservasi Sumberdaya Hutan dan Ekowisata', 1008, '2022-06-14 15:09:00'),
	(100804, 'Manajemen Hutan', 1008, '2022-06-14 15:09:00'),
	(100805, 'Rekayasa Kehutanan', 1008, '2022-06-14 15:09:00'),
	(100806, 'Silvikultur', 1008, '2022-06-14 15:09:00'),

	(100901, 'Kedokteran Hewan', 1009, '2022-06-14 15:09:00'),

	(101001, 'Higiene Gigi', 1010, '2022-06-14 15:09:00'),
	(101002, 'Kedokteran Gigi', 1010, '2022-06-14 15:09:00'),

	(101101, 'Gizi', 1011, '2022-06-14 15:09:00'),
	(101102, 'Kedokteran', 1011, '2022-06-14 15:09:00'),
	(101103, 'Keperawatan', 1011, '2022-06-14 15:09:00'),

	(101201, 'Administrasi Bisnis', 1012, '2022-06-14 15:09:00'),
	(101202, 'Administrasi Publik', 1012, '2022-06-14 15:09:00'),
	(101203, 'Hubungan Internasional', 1012, '2022-06-14 15:09:00'),
	(101204, 'Ilmu Komunikasi', 1012, '2022-06-14 15:09:00'),
	(101205, 'Ilmu Pemerintahan', 1012, '2022-06-14 15:09:00'),
	(101206, 'Ilmu Perpustakaan', 1012, '2022-06-14 15:09:00'),
	(101207, 'Ilmu Politik', 1012, '2022-06-14 15:09:00'),
	(101208, 'Informasi dan Humas', 1012, '2022-06-14 15:09:00'),
	(101209, 'Kriminologi', 1012, '2022-06-14 15:09:00'),
	(101210, 'Manajemen dan Kebijakan Publik', 1012, '2022-06-14 15:09:00'),
	(101211, 'Pembangunan Sosial dan Kesejahteraan', 1012, '2022-06-14 15:09:00'),
	(101212, 'Pembangunan Wilayah', 1012, '2022-06-14 15:09:00'),
	(101213, 'Sosiologi', 1012, '2022-06-14 15:09:00'),

	(101301, 'Antropologi Budaya', 1013, '2022-06-14 15:09:00'),
	(101302, 'Antropologi Sosial', 1013, '2022-06-14 15:09:00'),
	(101303, 'Arkeologi', 1013, '2022-06-14 15:09:00'),
	(101304, 'Bahasa Asing Terapan', 1013, '2022-06-14 15:09:00'),
	(101305, 'Bahasa dan Kebudayaan Jepang', 1013, '2022-06-14 15:09:00'),
	(101306, 'Bahasa dan Kebudayaan Korea', 1013, '2022-06-14 15:09:00'),
	(101307, 'Sastra Arab', 1013, '2022-06-14 15:09:00'),
	(101308, 'Sastra Belanda', 1013, '2022-06-14 15:09:00'),
	(101309, 'Sastra Cina', 1013, '2022-06-14 15:09:00'),
	(101310, 'Sastra Indonesia', 1013, '2022-06-14 15:09:00'),
	(101311, 'Sastra Inggris', 1013, '2022-06-14 15:09:00'),
	(101312, 'Sastra Jawa', 1013, '2022-06-14 15:09:00'),
	(101313, 'Sastra Jepang', 1013, '2022-06-14 15:09:00'),
	(101314, 'Sastra Jerman', 1013, '2022-06-14 15:09:00'),
	(101315, 'Sastra Perancis', 1013, '2022-06-14 15:09:00'),
	(101316, 'Sastra Slavia', 1013, '2022-06-14 15:09:00'),
	(101317, 'Sejarah', 1013, '2022-06-14 15:09:00'),
	(101401, 'Hukum', 1014, '2022-06-14 15:09:00'),

	(101501, 'Filsafat', 1015, '2022-06-14 15:09:00'),

	(101601, 'Farmasi', 1016, '2022-06-14 15:09:00'),

	(101701, 'Akuntansi', 1017, '2022-06-14 15:09:00'),
	(101702, 'Akuntansi Perpajakan', 1017, '2022-06-14 15:09:00'),
	(101703, 'Bisnis', 1017, '2022-06-14 15:09:00'),
	(101704, 'Ekonomi', 1017, '2022-06-14 15:09:00'),
	(101705, 'Ekonomi Islam', 1017, '2022-06-14 15:09:00'),
	(101706, 'Ekonomi Pembangunan', 1017, '2022-06-14 15:09:00'),
	(101707, 'Ekonomi Pertanian dan Agribisnis', 1017, '2022-06-14 15:09:00'),
	(101708, 'Ekonomi Sumberdaya dan Lingkungan', 1017, '2022-06-14 15:09:00'),
	(101709, 'Ekonomi Syariah', 1017, '2022-06-14 15:09:00'),
	(101710, 'Ilmu Kesejahteraan Sosial', 1017, '2022-06-14 15:09:00'),
	(101711, 'Kewirausahaan', 1017, '2022-06-14 15:09:00'),
	(101712, 'Manajemen', 1017, '2022-06-14 15:09:00'),
	(101713, 'Manajemen dan Administrasi Logistik', 1017, '2022-06-14 15:09:00'),
	
	(101801, 'Ilmu Keluarga dan Konsumen', 1018, '2022-06-14 15:09:00'),
	(101802, 'Komunikasi dan Pengembangan Masyarakat', 1018, '2022-06-14 15:09:00')
	;`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}
