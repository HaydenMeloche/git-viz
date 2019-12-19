# git-viz
git-viz is a simple GO application that creates the GitHub contribution graph in your terminal using data from your local repositories.

> This is my first golang project. Feedback is greatly appreciated!
### Example
```
$ go run * ../../ meloche.hayden@gmail.com
```
![Picture of git-viz in terminal](./img/screenshot.png?raw=true)

The first parameter provided is the location to the various git folders. The app will scan recersivley so providing the path to a folder containing all your git projections will work the best.

The second parameter is an email that will be used to generate the graph shown above.