# go-http-template

Golang HTTP Template เป็น template สำหรับการสร้าง HTTP service ด้วยภาษา Go ที่มีโครงสร้างโปรเจคที่เป็นมาตรฐานและพร้อมใช้งาน

> 📝 **หมายเหตุ**: README นี้ถูกสร้างโดย AI Assistant เพื่อเป็นแนวทางในการใช้งาน template นี้ คุณสามารถปรับแต่งเนื้อหาให้เหมาะสมกับโปรเจคของคุณได้

## คุณสมบัติ

- 🚀 โครงสร้างโปรเจคที่เป็นมาตรฐาน
- 🔄 Dependency Injection
- 📦 การจัดการ dependencies ที่ดี
- 🛠️ พร้อมสำหรับการพัฒนา
- 📝 API Documentation
- 🧪 Unit Testing
- 🔍 Linting และ Code Quality

## ข้อกำหนด

- Go 1.21 หรือสูงกว่า
- Git
- Docker (ถ้าต้องการรันผ่าน Docker)

## การปรับแต่ง Template

### 1. เปลี่ยนชื่อ Package

1. แก้ไขไฟล์ `go.mod`:
```bash
# เปลี่ยน module name
go mod edit -module github.com/yourusername/your-project-name
```

2. แก้ไข import paths ในทุกไฟล์:
```bash
# ใช้คำสั่ง find และ sed เพื่อแก้ไข import paths ทั้งหมด
# สำหรับ Linux/macOS
find . -type f -name "*.go" -exec sed -i 's|github.com/yourusername/go-http-template|github.com/yourusername/your-project-name|g' {} +

# สำหรับ Windows PowerShell
Get-ChildItem -Recurse -Filter "*.go" | ForEach-Object {
    (Get-Content $_.FullName) |
    ForEach-Object { $_ -replace 'github.com/yourusername/go-http-template', 'github.com/yourusername/your-project-name' } |
    Set-Content $_.FullName
}
```

3. อัปเดต dependencies:
```bash
go mod tidy
```

### 2. ตรวจสอบการเปลี่ยนแปลง

1. ตรวจสอบว่าไม่มี import paths เก่าค้างอยู่:
```bash
# ค้นหา import paths เก่า
grep -r "github.com/yourusername/go-http-template" .
```

2. ทดสอบว่าโปรเจคยังทำงานได้:
```bash
go test ./...
go build ./...
```

## โครงสร้างโปรเจค

```
.
├── cmd/            # จุดเริ่มต้นของแอปพลิเคชัน
│   └── main.go     # จุดเริ่มต้นหลัก
├── internal/       # โค้ดส่วนตัวที่ไม่ต้องการให้ภายนอกเข้าถึง
│   ├── adapter/    # Adapters สำหรับ external services
│   └── core/       # Core business logic
├── pkg/           # โค้ดที่สามารถนำไปใช้ภายนอกได้
├── di/            # Dependency Injection
├── scripts/       # สคริปต์สำหรับการ build และ deploy
├── docker-compose.yaml  # Docker Compose configuration
└── Makefile       # Make commands สำหรับการพัฒนา
```

## การติดตั้ง

```bash
# Clone repository
git clone https://github.com/yourusername/go-http-template.git

# เข้าไปที่โฟลเดอร์โปรเจค
cd go-http-template

# ติดตั้ง dependencies
go mod download
```

## การรัน

```bash
# รันในโหมด development
go run cmd/main.go

# รันผ่าน Docker
docker-compose up
```

## การทดสอบ

```bash
# รันการทดสอบทั้งหมด
go test ./...

# รันการทดสอบพร้อม coverage
go test ./... -cover
```

## การ Build

```bash
# Build binary
go build -o bin/app cmd/main.go

# Build ผ่าน Docker
docker build -t go-http-template .
```

## การพัฒนา

### การเพิ่ม Dependencies

```bash
# เพิ่ม dependency
go get github.com/example/package

# อัปเดต dependencies
go mod tidy
```

### การ Format Code

```bash
# Format code ทั้งหมด
go fmt ./...

# ตรวจสอบ code style
go vet ./...
```

## การ Contribute

1. Fork repository
2. สร้าง branch ใหม่ (`git checkout -b feature/amazing-feature`)
3. Commit การเปลี่ยนแปลง (`git commit -m 'Add some amazing feature'`)
4. Push ไปยัง branch (`git push origin feature/amazing-feature`)
5. เปิด Pull Request

## License

โปรเจคนี้อยู่ภายใต้ลิขสิทธิ์ MIT - ดูรายละเอียดเพิ่มเติมได้ใน [LICENSE](LICENSE) file
