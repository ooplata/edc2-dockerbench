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
    let suma = 0;
    let contador = 0;
    let numero = 2;

    while (contador < n) {
        if (esPrimo(numero)) {
            suma += numero;
            contador++;
        }
        numero++;
    }
    //console.log(suma)

    return suma;
}

function medirTiempoEjecucion(func) {
    return async function(...args) {
        const inicio = process.hrtime.bigint();
        await func(...args);
        const fin = process.hrtime.bigint();
        const tiempoEjecucion = Number(fin - inicio) / 1000000; // Convertir nanosegundos a milisegundos
        console.log(`Tiempo de ejecución: ${tiempoEjecucion} ms`);
        return tiempoEjecucion;
    };
}

// Envuelve sumaPrimerosPrimos en una función para pasarla a medirTiempoEjecucion
const tiempoTarea = medirTiempoEjecucion(() => sumaPrimerosPrimos(10000));

async function ejecutar() {
    await tiempoTarea();
}

ejecutar();
