# go-http-template

Golang HTTP Template เป็น template สำหรับการสร้าง HTTP service ด้วยภาษา Go ที่มีโครงสร้างโปรเจคที่เป็นมาตรฐานและพร้อมใช้งาน

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
