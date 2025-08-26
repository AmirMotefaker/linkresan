# خط فعلی: FROM golang:1.22 AS builder
# خط جدید:
FROM golang:1.25 AS builder

# پوشه کاری را داخل کانتینر تعریف می‌کنیم
WORKDIR /app

# فایل‌های go.mod و go.sum را کپی می‌کنیم تا وابستگی‌ها را دانلود کنیم
COPY go.mod go.sum ./
RUN go mod download

# تمام فایل‌های پروژه را کپی می‌کنیم
COPY . .

# برنامه را به یک فایل اجرایی (باینری) کامپایل می‌کنیم
RUN CGO_ENABLED=0 GOOS=linux go build -o /linkresan-app ./main.go

# مرحله دوم: ساخت کانتینر نهایی (برای کاهش حجم)
# از یک ایمیج سبک به نام alpine استفاده می‌کنیم
FROM alpine:latest

# نصب ssl-certs برای درخواست‌های HTTPS
RUN apk --no-cache add ca-certificates

# فایل اجرایی ساخته شده را از مرحله اول به این کانتینر کپی می‌کنیم
COPY --from=builder /linkresan-app /usr/local/bin/linkresan-app

# پورت 8080 را برای دسترسی خارجی باز می‌کنیم
EXPOSE 8080

# دستور اجرایی برنامه هنگام شروع کانتینر
ENTRYPOINT ["linkresan-app"]

# برای استفاده از متغیرهای محیطی از یک فایل
CMD ["-env", "production"]