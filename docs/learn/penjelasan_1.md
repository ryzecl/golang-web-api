# Penjelasan Kode & Konsep Dasar Golang (Pustaka API)

Dokumen ini merangkum penjelasan arsitektur, detail perubahan kode yang di-track oleh Git, serta konsep dasar Golang dengan komparasi ke ekosistem **Laravel** dan **Next.js / TypeScript**.

---

## 1. Peta Analogi: Golang vs Laravel vs TypeScript / Next.js

Perubahan kode ini adalah proses refactoring dari gaya scripting langsung di `main.go` menuju **Layered Architecture (3-Tier / Clean Architecture)**.

| Konsep di Project Ini | Laravel | Next.js / TypeScript | Fungsi |
| :--- | :--- | :--- | :--- |
| **`entity.go` (`Book`)** | Eloquent Model (`Book`) | Prisma Model / Type definition | Representasi tabel database. |
| **`request.go` (`BookRequest`)** | FormRequest / DTO | Zod Schema / Interface Request Body | Format input data dari user & validasi. |
| **`repository.go`** | Repository Pattern / Eloquent Query | Prisma client query functions | Khusus urusan query ke database (CRUD). |
| **`service.go`** | Service Class / Action | Server Action / Business logic layer | Tempat business logic & validasi bisnis. |
| **`handler/book.go`** | Controller (`BookController`) | Route Handler (`app/api/.../route.ts`) | Menerima HTTP request, parsing JSON, kirim response. |
| **`main.go`** | Service Container bootstrap | Entry point server | Tempat inisialisasi DB, routing, dan dependency wiring. |

---

## 2. Bedah Detail Perubahan Kode per File

### A. `book/input.go` Dihapus $\rightarrow$ Diganti `book/request.go`
```go
package book

import "encoding/json"

type BookRequest struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}
```
* **Apa yang berubah?** Struct diubah namanya menjadi `BookRequest` (lebih deskriptif seperti `FormRequest` di Laravel), dan tipe data `Price` diubah dari `int` menjadi `json.Number`.
* **Kenapa `json.Number`?** Di JSON/TypeScript, angka kadang bisa dikirim sebagai string `"99000"` atau number `99000`. Jika menggunakan `int` biasa, parsing bisa error saat payload berupa string. `json.Number` menampung angka secara fleksibel dan bisa dikonversi secara eksplisit lewat `.Int64()` di Service.
* **Struct Tag (`json:"..." binding:"..."`)**:
  * `json:"title"`: Menentukan mapping field JSON dari body HTTP.
  * `binding:"required"`: Validasi bawaan validator Gin (mirip `'title' => 'required'` di Laravel).

---

### B. `book/repository.go`
Sebelumnya interface `Repository` memakai nama method huruf kecil (`findAll`, `findByID`, `create`) dan belum ada implementasinya.

Perubahan sekarang:
1. **Export Interface**:
   ```go
   type Repository interface {
       FindAll() ([]Book, error)
       FindByID(ID int) (Book, error)
       Create(book Book) (Book, error)
   }
   ```
2. **Struct & Constructor**:
   ```go
   type repository struct {
       db *gorm.DB
   }

   func NewRepository(db *gorm.DB) *repository {
       return &repository{db}
   }
   ```
3. **Implementasi Method dengan Pointer Receiver**:
   ```go
   func (r *repository) FindAll() ([]Book, error) {
       var books []Book
       err := r.db.Find(&books).Error
       if err != nil {
           return nil, err
       }
       return books, nil
   }
   ```
   * Mengambil data DB via GORM. Perhatikan `&books` (pointer): GORM langsung mengisi data ke dalam memory variable `books`.

---

### C. `book/service.go` (File Baru)
Layer ini bertindak sebagai jembatan antara Controller/Handler dan Database Repository.

```go
type Service interface {
    FindAll() ([]Book, error)
    FindByID(ID int) (Book, error)
    Create(bookRequest BookRequest) (Book, error)
}

type service struct {
    repository Repository
}

func NewService(repository Repository) *service {
    return &service{repository}
}
```
* **Dependency Injection**: `service` menerima interface `Repository`. Ini sama persis dengan Laravel:
  ```php
  public function __construct(private BookRepositoryInterface $repo) {}
  ```
* **Konversi Data di `Create()`**:
  ```go
  func (s *service) Create(bookRequest BookRequest) (Book, error) {
      price, err := bookRequest.Price.Int64()
      if err != nil {
          return Book{}, err
      }

      book := Book{
          Title: bookRequest.Title,
          Price: int(price),
      }

      newBook, err := s.repository.Create(book)
      return newBook, err
  }
  ```
  Di sini `BookRequest` (DTO/Input user) divalidasi/dikonversi tipenya, lalu dipetakan ke struct entity `Book` sebelum disimpan ke repository.

---

### D. `handler/book.go`
```diff
- var bookInput book.BookInput
+ var bookInput book.BookRequest
```
Mengikuti perubahan nama struct dari `BookInput` ke `BookRequest`.

---

### E. `main.go`
* Kode CRUD eksperimen langsung via raw `db` (seperti `db.Where(...)`, `db.Save(...)`, `db.Delete(...)`) dibersihkan.
* Kode diganti dengan menghubungkan antar layer (wiring dependencies):
  ```go
  // 1. Inisialisasi Repository (butuh koneksi DB)
  bookRepository := book.NewRepository(db)

  // 2. Inisialisasi Service (di-inject Repository)
  bookService := book.NewService(bookRepository)

  // 3. Test panggil Service
  bookRequest := book.BookRequest{
      Title: "From Zero to Hero",
      Price: "99000",
  }
  createBook, err := bookService.Create(bookRequest)
  if err != nil {
      log.Fatal("Failed to create book", err)
  }
  ```

---

## 3. Konsep Dasar Golang yang Wajib Dipahami

### 1. Public vs Private ditentukan oleh Huruf Kapital (Bukan keyword `public`/`private`)
* **Kapital (Exported)**: Bisa diakses dari package lain.
  * Contoh: `BookRequest`, `FindAll()`, `NewRepository()` $\rightarrow$ bisa dipanggil di `main.go` atau `handler`.
* **Huruf Kecil (Unexported)**: Hanya bisa diakses di package yang sama (`package book`).
  * Contoh: `type repository struct`, `type service struct` $\rightarrow$ tidak bisa diakses langsung dari `main.go`. Oleh karena itu disediakan fungsi `NewRepository` dan `NewService` yang huruf depan kapital.

### 2. Golang tidak punya `class`, tapi pakai `struct` + `Receiver`
Di PHP atau TypeScript:
```typescript
class Repository {
  private db: DB;
  constructor(db: DB) { this.db = db; }
  findAll() { ... }
}
```
Di Golang, dipisah antara data (`struct`) dan perilakunya (`method receiver`):
```go
// Datanya
type repository struct {
    db *gorm.DB
}

// Method-nya ( (r *repository) adalah receiver, mirip "$this" di Laravel atau "this" di TS )
func (r *repository) FindAll() ([]Book, error) {
    // r.db sama seperti $this->db atau this.db
}
```

### 3. Constructor Pattern (`New...`)
Go tidak memiliki constructor bawaan seperti `__construct()` atau `new MyClass()`.
Konvensi di Go adalah membuat fungsi biasa dengan awalan `New...` yang mengembalikan pointer struct:
```go
func NewRepository(db *gorm.DB) *repository {
    return &repository{db}
}
```

### 4. Interface di Go bersifat Implicit (Duck Typing)
Di TypeScript atau PHP, kita menulis `class Service implements ServiceInterface`.
Di Golang: **Tidak ada keyword `implements`**.
Jika `*repository` memiliki method `FindAll()`, `FindByID()`, dan `Create()` dengan signature yang sama persis dengan yang ada di `type Repository interface`, maka secara otomatis Go menganggap `*repository` sudah mengimplementasikan interface `Repository`.

### 5. Pointer (`*` dan `&`)
* `&variable` = Mengambil alamat memory (Address-of).
* `*Type` = Tipe pointer (menunjuk ke alamat memory objek tersebut, bukan copy by value).
* Mengapa `return &repository{db}`? Agar objek struct tidak di-copy ulang setiap dipindah-pindah, melainkan mereferensikan instance yang sama.
* Mengapa `r.db.Find(&books)`? Karena fungsi GORM butuh alamat memori variabel `books` agar bisa memasukkan hasil query langsung ke variabel tersebut.

### 6. Tidak Ada Exception (`try ... catch`), Semua Error adalah Return Value
Di Go, error dikembalikan sebagai nilai biasa:
```go
result, err := bookService.Create(bookRequest)
if err != nil {
    // Handle error secara manual (mirip pengecekan result.error di TypeScript)
    log.Fatal("Failed to create book", err)
}
```

---

## 4. Alur Kerja Aplikasi Saat Ini

```text
[HTTP Request / main.go]
          │
          ▼
   [BookRequest]  (DTO / Validasi Input)
          │
          ▼
    [BookService] (Business Logic & konversi BookRequest -> Book entity)
          │
          ▼
 [BookRepository] (Query GORM Create/Find ke Database)
          │
          ▼
   [MySQL Database]
```
