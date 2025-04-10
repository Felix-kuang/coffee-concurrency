# coffee-concurrency

Simulasi coffee shop kecil pake Golang, goroutines, dan channels.  

##  Fitur
- Multiple baristas (goroutine) ambil order dari channel
- Simulasi waktu bikin kopi (random delay)
- Output real-time kaya suasana di coffee shop

##  Konsep
Pakai pattern **worker pool** buat nunjukkin kekuatan concurrency di Go.

## Cara jalanin
```bash
go run main.go
