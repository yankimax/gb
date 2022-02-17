#Компиляция под разные платформы.
Компилируем под Win: ```GOOS=windows go build main.go```

Смотрим информацию: ```file main.exe```

Смотрим вывод: ```file main.exe```
```
main.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows
```

Конфигурируем под Linux x64: ```GOOS=linux GOARCH=arm64 go build main.go```

Смотрим вывод: ```file main```
```
main: ELF 64-bit LSB executable, ARM aarch64, version 1 (SYSV), statically linked, Go BuildID=hRkqywg779WBmAdNEd6I/m4ur0nlYy2U6IxMqD30-/sFdLHKhu7z6BMZD4gfYk/FPclU5P2-P9l1TRnwT64, not stripped
```

Команда ```godoc -http=:6060``` вернула 
```
using GOPATH mode
```

http://127.0.0.1:6060/pkg/ - открывает страничку с документацией по всем установленным пакетам.
 