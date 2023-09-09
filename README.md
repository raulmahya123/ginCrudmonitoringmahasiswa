# firstconnected
## LIST
1. Percobaan Dokumentasi API Menggunakan Swagger : <b> TO DO </b>

## DOKUMENTASI API BY POSTMAN
1. Dokumentasi Endpoint Fosil : import folder postman


{
	"info": {
		"_postman_id": "65c5fb6e-53cf-401c-ac12-992e567d67c0",
		"name": "jokii",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "23896798",
		"_collection_link": "https://interstellar-crescent-246765.postman.co/workspace/Team-Workspace~713a812c-5aa1-417d-8077-7ef0db5e8011/collection/23896798-65c5fb6e-53cf-401c-ac12-992e567d67c0?action=share&source=collection_link&creator=23896798"
	},
	"item": [
		{
			"name": "Mahasiswa",
			"item": [
				{
					"name": "Mahasiswa",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Nama_mhs\":\"zzzzoya\",\r\n    \"NPM\":\"123\",\r\n    \"Jurusan\":\"123\",\r\n    \"Email\":\"123\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/mahasiswa",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"mahasiswa"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get all",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/mahasiswas",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"mahasiswas"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Id",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/mahasiswa/64fc670a32cf6bd6e284d7e2",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"mahasiswa",
								"64fc670a32cf6bd6e284d7e2"
							]
						}
					},
					"response": []
				},
				{
					"name": "Put Id",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Nama_mhs\":\"adeeeeeaee\",\r\n    \"NPM\":\"asda12aadasadsqqdsa\",\r\n    \"Jurusan\":\"12asdasadww\",\r\n    \"Email\":\"1qasdaasq12\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/mahasiswa/64fc670a32cf6bd6e284d7e2",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"mahasiswa",
								"64fc670a32cf6bd6e284d7e2"
							]
						}
					},
					"response": []
				},
				{
					"name": "Deleted Id",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "localhost:8080/mahasiswa/64fc670a32cf6bd6e284d7e2",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"mahasiswa",
								"64fc670a32cf6bd6e284d7e2"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "OrangTua",
			"item": [
				{
					"name": "Orangtua",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Nama_ortu\":\"ZOYA\",\r\n    \"Phone_number\":\"127835612785638172637812\",\r\n    \"Email\":\"asd12yu3f1uy2f312ghj\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/orangtua",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"orangtua"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get all",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/orangtuas",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"orangtuas"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Id",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/orangtua/64fc593c442bf8e2094d4e8c",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"orangtua",
								"64fc593c442bf8e2094d4e8c"
							]
						}
					},
					"response": []
				},
				{
					"name": "Put Id",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Nama_ortu\":\"asd\",\r\n    \"Phone_number\":\"asd\",\r\n    \"Email\":\"asd\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/orangtua/64fc55214ba8282c2d8eb9ec",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"orangtua",
								"64fc55214ba8282c2d8eb9ec"
							]
						}
					},
					"response": []
				},
				{
					"name": "Deleted Id",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "localhost:8080/orangtua/64fc593c442bf8e2094d4e8c",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"orangtua",
								"64fc593c442bf8e2094d4e8c"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "matakuliah",
			"item": [
				{
					"name": "Matakuliah",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Nama_matkul\":\"asdasdasasdasdsadads\",\r\n    \"SKS\":\"aaas\",\r\n    \"Dosen_pengampu\":\"aasdsdssssa\",\r\n    \"Email\":\"asdasda\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/matakuliah",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"matakuliah"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get all",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"method": "GET",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": ""
						},
						"url": {
							"raw": "localhost:8080/matakuliahs",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"matakuliahs"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Id",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/matakuliah/64fc57c7f640877d7e1cdd89",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"matakuliah",
								"64fc57c7f640877d7e1cdd89"
							]
						}
					},
					"response": []
				},
				{
					"name": "Put Id",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Nama_matkul\":\"asdaasdasdsdasasdasdsadads\",\r\n    \"SKS\":\"aaaasds\",\r\n    \"Dosen_pengampu\":\"aasdasdsdssssa\",\r\n    \"Email\":\"asdasdasda\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/matakuliah/64fc55264ba8282c2d8eb9ee",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"matakuliah",
								"64fc55264ba8282c2d8eb9ee"
							]
						}
					},
					"response": []
				},
				{
					"name": "Deleted Id",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "localhost:8080/matakuliah/64fc55264ba8282c2d8eb9ee",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"matakuliah",
								"64fc55264ba8282c2d8eb9ee"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "nilai",
			"item": [
				{
					"name": "Nilai",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"NPM_ms\":\"123\",\r\n    \"Presensi\":\"1\",\r\n    \"Nilai_akhir\":\"2\",\r\n    \"Grade\":\"3\",\r\n    \"Tahun_ajaran\":\"4\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/nilai",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"nilai"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get all",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/nilais",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"nilais"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Id",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/nilai/64fc563e79d6ccf4fe44d09a",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"nilai",
								"64fc563e79d6ccf4fe44d09a"
							]
						}
					},
					"response": []
				},
				{
					"name": "Put Id",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"NPM_ms\":\"123\",\r\n    \"Presensi\":\"1\",\r\n    \"Nilai_akhir\":\"2\",\r\n    \"Grade\":\"3\",\r\n    \"Tahun_ajaran\":\"4\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/nilai/64fc563e79d6ccf4fe44d09a",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"nilai",
								"64fc563e79d6ccf4fe44d09a"
							]
						}
					},
					"response": []
				},
				{
					"name": "Deleted Id",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "localhost:8080/nilai/64fc5865d560b4dfa2cf9274",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"nilai",
								"64fc5865d560b4dfa2cf9274"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "absensi",
			"item": [
				{
					"name": "absensi",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Nama_mk\":\"ZOYA\",\r\n    \"Tanggal\":\"aaas\",\r\n    \"Checkin\":\"aasdsdssssa\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/asbensi",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"asbensi"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get all",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/absensis",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"absensis"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Id",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "localhost:8080/absensi/64fc58aad560b4dfa2cf9276",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"absensi",
								"64fc58aad560b4dfa2cf9276"
							]
						}
					},
					"response": []
				},
				{
					"name": "Put Id",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Nama_mk\":\"ZOYasdA\",\r\n    \"Tanggal\":\"aaaasds\",\r\n    \"Checkin\":\"aasdasdsdssssa\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "localhost:8080/absensi/64fc58aad560b4dfa2cf9276",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"absensi",
								"64fc58aad560b4dfa2cf9276"
							]
						}
					},
					"response": []
				},
				{
					"name": "Deleted Id",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "localhost:8080/absensi/64fc58aad560b4dfa2cf9276",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"absensi",
								"64fc58aad560b4dfa2cf9276"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "Login",
			"item": [
				{
					"name": "Login",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Password\":\"1231madepssp\",\r\n    \"Email\":\"raulade1212333@gmail.com\"\r\n}"
						},
						"url": {
							"raw": "localhost:9000/user/login",
							"host": [
								"localhost"
							],
							"port": "9000",
							"path": [
								"user",
								"login"
							]
						}
					},
					"response": []
				},
				{
					"name": "Login SignUp",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"Password\":\"1231madepssp\",\r\n    \"Email\":\"raulade1212333@gmail.com\"\r\n}"
						},
						"url": {
							"raw": "localhost:9000/user/login",
							"host": [
								"localhost"
							],
							"port": "9000",
							"path": [
								"user",
								"login"
							]
						}
					},
					"response": []
				}
			]
		}
	]
}
