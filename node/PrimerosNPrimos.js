const fs = require("node:fs");

function esPrimo(num) {
    if (num <= 1) return false;
    if (num <= 3) return true;
    if (num % 2 === 0 || num % 3 === 0) return false;
    for (let i = 5; i * i <= num; i += 6) {
        if (num % i === 0 || num % (i + 2) === 0) return false;
    }
    return true;
}

function sumaPrimerosPrimos(n) {
    let suma = 2;
    let contador = 1;
    let numero = 3;

    while (contador < n) {
        if (esPrimo(numero)) {
            suma += numero;
            contador++;
        }
        numero += 2;
    }
    return suma;
}

function escribirPrimerosPrimos(n) {
    const primos = sumaPrimerosPrimos(n);
    fs.writeFileSync("out.txt", primos.toString());
}

function medirTiempoEjecucion(func) {
    return function (...args) {
        const inicio = process.hrtime.bigint();
        func(...args);
        const fin = process.hrtime.bigint();

        return Number(fin - inicio) / 1000000; // Convertir a milisegundos
    };
}

// Medir el tiempo de la función correctamente
const tiempoTarea = medirTiempoEjecucion(escribirPrimerosPrimos);
const tiempoEjecucion = tiempoTarea(10000);
console.log(`${tiempoEjecucion}ms`); // Imprime el tiempo en ms
