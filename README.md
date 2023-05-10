# Konwerter faktur do systemu iHurt
dokumentacja pliku wynikowego:

https://edokumenty.insignum.pl/dn/spec/INSIGNUM%20eDokumenty%20-%20specyfikacja%20pliku%20komunikacyjnego%20Faktura.pdf

# development

Zainstalować go https://go.dev/doc/install
```
go run main.go --input ./demo/WAR-MIX_FV_VAT__F.02807.23_1\ 14.15.49.xml --output demo.xml --mode warmix
```

# budowanie
```
go build
```

# uruchamianie

```
./xmlconverter --input ./demo/WAR-MIX_FV_VAT__F.02807.23_1\ 14.15.49.xml --output demo.xml --mode warmix
```

# Parametry
 * input - ścieżka do pliku z plikem do przekonwertowania
 * output - plik wyjściowy
 * mode - format konwertera, obecnie obsługiwany tylko "warmix"
