---
name: jelasin
description: Menjelaskan perubahan kode terbaru di Git, arsitektur aplikasi, dan konsep dasar Golang secara mendalam dengan analogi Laravel dan Next.js/TypeScript. Aktifkan skill ini ketika user mengetik '/jelasin' atau meminta penjelasan perubahan kode untuk belajar Golang.
---

# Skill: Jelasin (Belajar Golang via Git Diff & Analogi Laravel/Next.js)

Skill ini dirancang khusus untuk memandu developer dengan latar belakang **Laravel (PHP)** dan **Next.js (TypeScript)** dalam memahami codebase **Golang** secara mendalam berbasis perubahan kode (*learning by doing via git diff*).

---

## 1. Trigger Kondisi

Skill ini diaktifkan ketika:
- User mengetik command `/jelasin`
- User meminta penjelasan mengenai perubahan kode yang di-track oleh git
- User meminta penjelasan konsep Golang dari kode yang sedang dikerjakan

---

## 2. Alur Eksekusi Agen (Step-by-Step)

Ketika skill ini dijalankan, agen **HARUS** melakukan langkah-langkah berikut secara berurutan:

### Langkah 1: Inspeksi Perubahan Git
1. Jalankan perintah git untuk memeriksa status workspace:
   ```bash
   git status
   git diff
   ```
2. Jika ada *untracked files*, buka dan baca file tersebut menggunakan `view_file` untuk memahami konteks penuh dari fitur yang baru dibuat.

### Langkah 2: Identifikasi & Pemetaan Arsitektur
Kelompokkan setiap perubahan ke dalam layer arsitektur:
- **Entity/Model**: Representasi tabel database (misal: `entity.go`).
- **Request/DTO**: Input data dari client dan validasi (misal: `request.go`).
- **Repository**: Data Access Layer / query database (misal: `repository.go`).
- **Service**: Business logic & data orchestration (misal: `service.go`).
- **Handler/Controller**: HTTP transport layer, request parsing & response handling (misal: `handler/`).
- **Main / Bootstrap**: Entry point & dependency injection wiring (misal: `main.go`).

### Langkah 3: Sajikan Tabel Analogi (Laravel & TypeScript)
Berikan tabel komparasi agar pengguna langsung memahami konsep tersebut dengan analogi teknologi yang sudah dikuasainya:

| Konsep di Fitur Ini | Padanan di Laravel | Padanan di Next.js / TypeScript | Fungsi |
| :--- | :--- | :--- | :--- |
| Struct Model | Eloquent Model | Prisma Model / Drizzle Schema | Skema tabel DB |
| Struct Request | FormRequest | Zod Schema / Type Body | Validasi & DTO |
| Repository | Repository / Query Scope | Prisma Client queries | Akses database |
| Service | Service Class / Action | Server Action / Lib function | Business logic |
| Handler | Controller Method | Route Handler (`route.ts`) | HTTP endpoint |
| `main.go` Wiring | Service Container / Providers | App Router / Server bootstrap | Dependency Injection |

### Langkah 4: Bedah Kode File per File
Jelaskan setiap file yang berubah dengan format:
1. **Apa yang Berubah?** (Tampilkan snippet perubahan penting).
2. **Kenapa Diubah?** (Alasan arsitektural / best practice).
3. **Detail Kodingan & Logika Aplikasi**: Cara data mengalir dan diproses.

### Langkah 5: Deep-Dive Konsep Dasar Golang yang Muncul
Jelaskan konsep Go yang relevan dari perubahan tersebut secara mendalam:
- **Exported vs Unexported Identifiers**: Aturan huruf kapital (public) vs huruf kecil (private/package-internal).
- **Struct & Method Receiver**: Perbedaan OOP berbasis class vs struct dengan receiver `(r *repository) MethodName()`.
- **Pointer (`*` dan `&`)**: Kapan harus pass-by-value vs pass-by-reference/pointer, kenapa GORM butuh `&books`.
- **Interface & Duck Typing**: Mengapa tidak ada keyword `implements` di Go, dan bagaimana Inversion of Control (IoC) bekerja.
- **Constructor Pattern (`New...`)**: Konvensi pembuatan instance struct di Go pengganti `__construct()` / `new Class()`.
- **Data Types & Struct Tags**: Penjelasan tag backtick `` `json:"..." binding:"..."` `` serta tipe khusus seperti `json.Number`.
- **Explicit Error Handling**: Pola `val, err := func()` dan `if err != nil`, mengapa Go tidak memakai `try-catch`.

### Langkah 6: Visualisasi Alur Data (Flow Diagram)
Buat visualisasi alur data sederhana (menggunakan Mermaid diagram atau teks diagram alur) dari incoming request sampai ke database dan response kembali ke client.

### Langkah 7: Tawarkan Penyimpanan Dokumentasi
Tanyakan atau tawarkan kepada user untuk mendokumentasikan penjelasan ke dalam folder pembelajaran:
`docs/learn/penjelasan_<nomor>.md`
agar menjadi log catatan belajar yang rapi dan terorganisir.

---

## 3. Gaya Bahasa dan Komunikasi

- **Bahasa**: Bahasa Indonesia santai, teknis, lugas, dan bersahabat (gaya rekan kerja/senior pair programmer).
- **Tone**: Edukatif dan aplikatif. Fokus pada *"Aha! Moment"* dengan menghubungkan fitur Go ke analogi Laravel & TypeScript.
- **Link File**: Selalu gunakan link markdown yang dapat diklik ke file kode terkait, misalnya `[repository.go](file:///...)`.
