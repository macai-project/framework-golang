# framework-golang

Questo è il framework che verrà utilizzato dalle lambda per l'esecuzione del codice relativo agli eventi.

Conterrà tutto il codice utile alla corretta esecuzione della lambda ma non le logiche di business.

In questo framework verranno configurate tutte le dipendenze che possono essere usate dai progetti:

* logger
* connessioni al database
* xray

Le logiche di business risiederanno nella funzione principale dei progetti.

Per capirne il funzionamento guarda lo scheletro: [https://github.com/macai-project/scheletro-golang](https://github.com/macai-project/scheletro-golang)
