"use strict";
console.log("Hello from TypeScript!");
fetch("http://localhost:8080")
    .then(response => response.text())
    .then(data => console.log(data))
    .catch(err => console.error(err));
