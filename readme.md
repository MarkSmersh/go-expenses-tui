## GO EXPENSES TUI
Aplikacja z interfejsem terminałowowym do zapisanya wytrat

### Instalacja

Aby rozpocząć pracę z aplikacją jest potrzebny server i sama aplikacja.

1. Server może być skompilowany za pomocą komandy:

`go run ./cmd/server/main.go`

Po kompilacji stworzy aplikację o nazwie server(.exe)

2. Aplikacja kompiluje się za pomocą

`go run ./cmd/tui/main.go`

Po kompilacji stworzy aplikację o nazwie tui(.exe)

Przed startem samej aplikacji należy być uruchomiony server. Środowisko musi zawierać zmiane:
- DB_URI: np. `postgres://username:password@addr:port/db_name` - to jest adres do bazy danych na której będzie przechowywana informacja
- JWT_SIGN_KEY: np. `qwertyuiopasdfghjklzxcvbnmqwerty` - to kod z 32 symboli, który będzie używany jako klusz szufrujący dla JWT kluczy używanych dla autentyfikacji użytkowników

### Uwagi

- Aplikacja NIE pracuja bez serwera i nie zawiera lokalnego przechowywania danych, oprócz samych ustawień użytkowanika np. adresu serwera, nazwy użytkowina i JWT klucza, jaki jest przechowiwany dla autentyfikacji z serwerem. Więc przy pojawieniu błedów powiązanych z aplikacją - nie ma potrzeby matrwić się. To nie państwa problem. A problem programisty.

- W niektórych wypadkach, aplikacja może nie adaptować do rozmiaru okna. To jest raczej problem ze strony programista (w go glowie), ale to nie nastrajnieszy problem i szybko eliminuje się przy zmianie rozmiaru okna.

- Jeżeli nie uda się urochomić server - jest możliwość korzystania z domenu `goexpensestui.qekkk.xyz` (wprowadzić jako server). Najprowdopodopniej będzie pracować do 31.12.2025, bierąc pod warunek te, że: nie będzie brak światla w pokoju serwera z powodu rosyjskiego drona (dury), rosyjsky dron nie znisczy serwera i serwer nie wybuchnie.

### Plany rozwijęcia aplikacji
_jęzeli pollub nie przywatyzuje tą aplikację. W takim razie nie mam chęci._

- Program będzie zawierać nie tylko wytraty, ale i zarobki. Nie jest to wielki problem to zrobić. Ale potrzebuje czasu.
- Dopracować moduł kalendarza. On pracuje, ale nie pozwala do wyberiania dni, jeżeli ustawić kursor zgóry. Pracuje jęzeli tylko jeżeli kursor z dołu. Przyczyna pochodzi z tego, że kalendarz był napisany od zero w czasie limitowanym.
- Więcej aniż 1 waluta, wybór waluty podstawowej, konwertacja waluty wedłuj kursu (z odstępem w 1 minutę) oraz zapis kursu do USA, CHF oraz... złota. Dlaczego? Do porownania wartości waluty podstawowej w przysłosci według kązdej tranzakcji (zapisu wytraty/zaróbku)
- Prechowywanie tranzakcji lokalne przy braku połaczenie z internetem (takie bywa?) i ich syncronizacja z serwerem.
