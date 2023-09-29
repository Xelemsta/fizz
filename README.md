# fizz

- "-" is not allowed in str1 et str2 (car sert de délimiteur)
- errors returned properly (bad request, not found, etc...)
- faire test top request (bonus)
- faire test redis package

choix/ assumption
- incr dans handler au lieu de middleware pour être après l'authentification et la validation des params
- redis: key/value parfait pour le besoin + data type pratique pour récupérer le top hit

amélioration
- process fizzbuzz avec goroutines pour accélerer le rendu
- spawn swagger ui
