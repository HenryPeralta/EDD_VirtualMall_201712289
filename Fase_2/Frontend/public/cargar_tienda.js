//alert("Hola");

function obtener_json(){
    var contenido = document.getElementById('consola_json').value;
    console.log(contenido);

    fetch('http://localhost:3000/Tiendas', {
        method: 'POST',
        headers:{
            'Content-Type': 'application/json'
        },
        body:contenido
    })
}