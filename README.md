# Amazeing
Dieses Programm löst ein als PNG-Bild dargestelltes Labyrinth. Das Labyrinth muss einen blauen Startpunkt, einen grünen Endpunkt und schwarze Pfadpixel haben. Die Lösung des Labyrinths wird als rote Pixel dargestellt. Der kürzeste Pfad wird dabei mittels eines Breitensuche-Algorithmus ermittelt.

## Beispiel
Um dieses Programm zu verwenden, erstellen Sie eine PNG-Bilddatei, die Ihr Labyrinth mit dem folgenden Farbschema darstellt:

- Blau: Startpunkt (max. ein Pixel)
- Grün: Endpunkt
- Schwarz: Pfad
- Weiß: leerer Raum

<figure>
  <img src="https://user-images.githubusercontent.com/51135157/233867966-8583743e-0a23-43ee-b509-81a5f8c0fa57.png" alt="maze6.png" style="width: 200px; height: 200px; object-fit: none; image-rendering: pixelated;">
  <figcaption>maze6.png</figcaption>
</figure>

Starten Sie das Programm und geben Sie den Dateipfad Ihres Labyrinths als Argument an:

```
go run amazeing.go maze6.png
```

Das Programm löst nun das Labyrinth und speichert das Ergebnis unter `maze6_solved.png`.

<figure>
  <img src="https://user-images.githubusercontent.com/51135157/233867989-157e9647-4d52-4102-890d-75ee1aa3912a.png" alt="maze6_solved.png" style="width: 200px; height: 200px; object-fit: none; image-rendering: pixelated;">
  <figcaption>maze6_solved.png</figcaption>
</figure>
