# Go Database Connection

## อธิบายการทำงานตอนเช้า
- เมื่อเช้ามีการตกลงกันเรื่องของ Test Data และ API ให้เป็นไปในรูปแบบเดียวกัน
- เลยเริ่มที่จะแก้ API ให้ตรงกับข้อตกลงก่อน จากนั้นจึงเริ่มเขียน Test
- แต่พอวิเคราะห์ดูแล้ว จากโค้ดที่เขียน การที่จะปรับให้มี Test มันยาก
- เพราะต้อง Stub และ Mock ของขึ้นมาเยอะเกินไป เนื่องจากไม่ได้แยก Handler และ Service ออกจากกัน
- พร้อมทั้งโค้ดที่เขียนยังมี Dependency กันอยู่ด้วย

## วิธีคิด เพื่อแก้ให้เขียน Test
- เริ่มแรกให้เริ่ม Connect Database ตั้งแต่ที่ main เลย
- จากนั้น แต่ละ Function ให้รับตัว Connection เข้ามา พร้อมกับฟังก์ชันที่ต้องการเรียกใช้
- พอตอนจะเขียน Test ก็ทำการ Mock และ Stub ของที่ Handler นั้นจะ Return กลับมา
- จริง ๆ ควรจะแยก Handler และ Service ออกจากกันให้ชัดเจนไปเลย

## How to Install Project
### Run with docker image
```bash
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=sckshuhari -d mysql:5.7
```

### Install MySql Workbench
```bash
https://dev.mysql.com/downloads/workbench/?utm_source=tuicool
```

### Install go mysql driver library
```bash
go get github.com/go-sql-driver/mysql
```