# tiny-go-sample

## Wio Terminal

## macropad
### TinyGoのインストール
https://tinygo.org/getting-started/install/macos/

### 書き込み
```bash
tinygo flash -target=macropad-rp2040 --size short .
```

### つまりポイント

書き込みエラー
```bash
tinygo flash -target=macropad-rp2040 --size short .
   code    data     bss |   flash     ram
   7748     108    3152 |    7856    3260
error: failed to flash /var/folders/d3/r0lrql114fzcwpsl0qdnn0br0000gn/T/tinygo1246708819/main.uf2: unable to locate device: RPI-RP2
```

### BOOTSELモードへの切り替え
https://learn.adafruit.com/adafruit-macropad-rp2040/circuitpython
