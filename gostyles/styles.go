package gostyles

const FILENAME string = "mpstyles.css"

const CSS string = `
@import url('https://fonts.googleapis.com/css2?family=Roboto&family=Roboto+Mono&display=swap');

* {
    font-family: 'Roboto', sans-serif;
}

body {
    margin: 0;
}

.mp-reserved-word {
    color: #569cd6;
}

.mp-string {
    color: #ce9178;
}

.mp-symbol {
    color: #ffd602;
}

.mp-number {
    color: #94a688;
}

.mp-func {
    color: #dcdcaa;
}

pre {
    background-color: #1e1e1e;
    padding: 0 1rem 1rem 1rem;
    border-radius: 0.2rem;
    cursor: pointer;
}

code {
    font-family: 'Roboto Mono', monospace;
    color: #ffffff;
}

code::selection,
code *::selection {
    background-color: transparent;
}
`