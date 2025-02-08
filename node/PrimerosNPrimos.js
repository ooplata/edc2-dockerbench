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

function medirTiempoEjecucion(func) {
    return function (...args) {
        const inicio = process.hrtime.bigint();
        const resultado = func(...args);
        const fin = process.hrtime.bigint();
        const tiempoEjecucion = Number(fin - inicio) / 1000000; // Convertir a milisegundos
        console.log(`${tiempoEjecucion}ms`); // Imprime el tiempo en ms
        return resultado;
    };
}

// Medir el tiempo de la función correctamente
const tiempoTarea = medirTiempoEjecucion(sumaPrimerosPrimos);
tiempoTarea(10000);
